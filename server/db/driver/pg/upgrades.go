// This code is available on the terms of the project LICENSE.md file,
// also available online at https://blueoakcouncil.org/license/1.0.0.

package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"

	"decred.org/dcrdex/dex"
	"decred.org/dcrdex/dex/calc"
	"decred.org/dcrdex/server/asset"
	"decred.org/dcrdex/server/db/driver/pg/internal"
)

const dbVersion = 7

// The number of upgrades defined MUST be equal to dbVersion.
var upgrades = []func(db *sql.Tx) error{
	// v1 upgrade adds the schema_version column to the meta table, possibly
	// creating the table if it was missing.
	v1Upgrade,

	// v2 upgrade creates epochs_report table, if it does not exist, and
	// populates the table with partial historical data from the epochs and
	// matches table. This includes match volumes, high/low/start/end rates, but
	// does not include the booked volume statistics in the book_buys* and
	// book_sells* columns since this data requires a book snapshot at the time
	// of matching to generate.
	v2Upgrade,

	// v3 upgrade adds the fee_asset column to the accounts table.
	v3Upgrade,

	// v4 upgrade updates the markets tables to use a integer type that can
	// accommodate a 32-bit unsigned integer.
	v4Upgrade,

	// v5 upgrade adds an epoch_gap column to the cancel order tables to
	// facilitate free cancels.
	v5Upgrade,

	// v6 upgrade creates the bonds table. A future upgrade may add a new
	// old_fee_coin column to the accounts table for when a manual refund is
	// processed.
	v6Upgrade,

	// v7 upgrade adds a reputation_ver column to the accounts table. This
	// facilitates a rolling upgrade of reputation tracking to address an issue
	// with the DB design.
	v7Upgrade,
}

// v1Upgrade adds the schema_version column and removes the state_hash column
// from the meta table.
func v1Upgrade(tx *sql.Tx) error {
	// Create the meta table with the v0 scheme. Even if the table does not
	// exists, we should not create it fresh with the current scheme since one
	// or more subsequent upgrades may alter the meta scheme.
	metaV0Stmt := `CREATE TABLE IF NOT EXISTS %s (state_hash BYTEA)`
	metaCreated, err := createTableStmt(tx, metaV0Stmt, publicSchema, metaTableName)
	if err != nil {
		return fmt.Errorf("failed to create meta table: %w", err)
	}
	if metaCreated {
		log.Infof("Created new %q table", metaTableName)    // from 0.2+pre master
		_, err = tx.Exec(`INSERT INTO meta DEFAULT VALUES`) // might be CreateMetaRow, but pin to the v0 stmt
		if err != nil {
			return fmt.Errorf("failed to create row for meta table: %w", err)
		}
	} else {
		log.Infof("Existing %q table", metaTableName) // from release-0.1
	}

	// Create the schema_version column. The caller must set the version to 1.
	_, err = tx.Exec(`ALTER TABLE ` + metaTableName + ` ADD COLUMN IF NOT EXISTS schema_version INT4 DEFAULT 0;`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`ALTER TABLE ` + metaTableName + ` DROP COLUMN IF EXISTS state_hash;`)
	return err
}

// matchStatsForMarketEpoch is used by v2Upgrade to retrieve match rates and
// quantities for a given epoch.
func matchStatsForMarketEpoch(stmt *sql.Stmt, epochIdx, epochDur uint64) (rates, quantities []uint64, sell []bool, err error) {
	var rows *sql.Rows
	rows, err = stmt.Query(epochIdx, epochDur)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rate, quantity fastUint64
		var takerSell bool
		err = rows.Scan(&quantity, &rate, &takerSell)
		if err != nil {
			return nil, nil, nil, err
		}
		rates = append(rates, uint64(rate))
		quantities = append(quantities, uint64(quantity))
		sell = append(sell, takerSell)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, nil, err
	}
	return
}

// v2Upgrade populates the epoch_reports table with historical data from the
// matches table.
func v2Upgrade(tx *sql.Tx) error {
	mkts, err := loadMarkets(tx, marketsTableName)
	if err != nil {
		return fmt.Errorf("failed to read markets table: %w", err)
	}

	unitInfo := func(assetID uint32) dex.UnitInfo {
		ui, err := asset.UnitInfo(assetID)
		if err != nil {
			log.Errorf("no unit info found for %d (%q)", assetID, dex.BipIDSymbol(assetID))
			ui.Conventional.ConversionFactor = 1e8
		}
		return ui
	}

	doMarketMatches := func(mkt *dex.MarketInfo) error {
		log.Infof("Populating %s with volume data for market %q matches...", epochsTableName, mkt.Name)

		baseUnitInfo, quoteUnitInfo := unitInfo(mkt.Base), unitInfo(mkt.Quote)

		// Create the epochs_report table if it does not already exist.
		_, err := createTable(tx, mkt.Name, epochReportsTableName)
		if err != nil {
			return err
		}

		// For each unique epoch duration, get the first and last epoch index.
		fullEpochsTableName := mkt.Name + "." + epochsTableName
		stmt := fmt.Sprintf(`SELECT epoch_dur, MIN(epoch_idx), MAX(epoch_idx)
			FROM %s GROUP BY epoch_dur;`, fullEpochsTableName)
		rows, err := tx.Query(stmt)
		if err != nil {
			return err
		}
		defer rows.Close()

		var durs, starts, ends []uint64
		for rows.Next() {
			var dur, first, last uint64
			if err = rows.Scan(&dur, &first, &last); err != nil {
				return err
			}
			durs = append(durs, dur)
			starts = append(starts, first)
			ends = append(ends, last)
		}

		if err = rows.Err(); err != nil {
			return err
		}

		// epoch_reports INSERT statement
		mktEpochReportsTablename := mkt.Name + "." + epochReportsTableName
		reportStmt := fmt.Sprintf(internal.InsertPartialEpochReport, mktEpochReportsTablename)
		reportStmtPrep, err := tx.Prepare(reportStmt)
		if err != nil {
			return err
		}
		defer reportStmtPrep.Close()

		// Create a temporary matches index on (epochidx, epochdur).
		fullMatchesTableName := mkt.Name + "." + matchesTableName
		matchIndexName := "matches_epidxdur_temp_idx"
		_, err = tx.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (epochidx, epochdur);",
			matchIndexName, fullMatchesTableName))
		if err != nil {
			return err
		}
		defer func() {
			if errors.Is(err, sql.ErrTxDone) {
				return // whole transaction including index creation is rolled back
			}
			// Success or other error - drop the index explicitly.
			fullIndexName := mkt.Name + "." + matchIndexName
			_, errDrop := tx.Exec(fmt.Sprintf("DROP INDEX %s;", fullIndexName))
			if errDrop != nil {
				log.Warnf("Failed to drop index %v: %v", fullIndexName, errDrop)
			}
		}()

		// matches(qty,rate,takerSell) SELECT statement
		matchStatsStmt := fmt.Sprintf(internal.RetrieveMatchStatsByEpoch, fullMatchesTableName)
		matchStatsStmtPrep, err := tx.Prepare(matchStatsStmt)
		if err != nil {
			return err
		}
		defer matchStatsStmtPrep.Close()

		var startRate, endRate uint64
		var totalMatches uint64
		var totalVolume, totalQVolume uint64
		for i, dur := range durs {
			log.Infof("Processing all %d of the %d ms %q epochs from idx %d to %d...",
				ends[i]-starts[i]+1, dur, mkt.Name, starts[i], ends[i])
			endIdx := ends[i]
			for idx := starts[i]; idx <= endIdx; idx++ {
				if idx%50000 == 0 {
					to := idx + 50000
					if to > endIdx+1 {
						to = endIdx + 1
					}
					log.Infof(" - Processing epochs [%d, %d)...", idx, to)
				}
				var rates, quantities []uint64 // don't shadow err from outer scope
				rates, quantities, _, err = matchStatsForMarketEpoch(matchStatsStmtPrep, idx, dur)
				if err != nil {
					return err
				}
				epochEnd := (idx + 1) * dur
				if len(rates) == 0 {
					// No trade matches in this epoch.
					_, err = reportStmtPrep.Exec(epochEnd, dur, 0, 0, 0, 0, startRate, startRate)
					if err != nil {
						return err
					}
					continue
				}

				var matchVolume, quoteVolume, highRate uint64
				lowRate := uint64(math.MaxInt64)
				for i, qty := range quantities {
					matchVolume += qty
					rate := rates[i]
					quoteVolume += calc.BaseToQuote(rate, qty)
					if rate > highRate {
						highRate = rate
					}
					if rate < lowRate {
						lowRate = rate
					}
				}
				totalVolume += matchVolume
				totalQVolume += quoteVolume
				totalMatches += uint64(len(quantities))

				// In the absence of a book snapshot, ballpark the rates. Note
				// that cancel order matches that change the mid market book
				// rate are not captured so start/end rates can be inaccurate
				// given long periods with no trades but book changes.
				midRate := (lowRate + highRate) / 2 // maybe average instead
				if startRate == 0 {
					startRate = midRate
				} else {
					startRate = endRate // from previous epoch with matches
				}
				endRate = midRate

				// No book buy / sell depth (see bookVolumes in server/matcher).
				_, err = reportStmtPrep.Exec(epochEnd, dur, matchVolume, quoteVolume,
					highRate, lowRate, startRate, endRate)
				if err != nil {
					return err
				}
			}
		} // range durs
		log.Debugf("Processed %d matches doing %s in %s volume (%s in %s volume)", totalMatches,
			baseUnitInfo.ConventionalString(totalVolume), strings.ToUpper(dex.BipIDSymbol(mkt.Base)),
			quoteUnitInfo.ConventionalString(totalQVolume), strings.ToUpper(dex.BipIDSymbol(mkt.Quote)))
		return nil
	}

	for _, mkt := range mkts {
		err = doMarketMatches(mkt)
		if err != nil {
			return err
		}
	}
	return nil
}

func v3Upgrade(tx *sql.Tx) error {
	// Create the fee_asset column.
	_, err := tx.Exec(`ALTER TABLE ` + accountsTableName + ` ADD COLUMN IF NOT EXISTS fee_asset INT4;`)
	if err != nil {
		return err
	}
	// Set existing rows fee_asset to 42, Decred's asset ID, since prior to this
	// upgrade, only DCR was accepted for registration.
	_, err = tx.Exec(`UPDATE ` + accountsTableName + ` SET fee_asset = 42;`) // not as default in ALTER
	return err
}

func v4Upgrade(tx *sql.Tx) (err error) {
	if _, err = tx.Exec("ALTER TABLE markets ALTER COLUMN base TYPE INT8;"); err != nil {
		return
	}
	_, err = tx.Exec("ALTER TABLE markets ALTER COLUMN quote TYPE INT8;")
	return err
}

func v5Upgrade(tx *sql.Tx) (err error) {
	mkts, err := loadMarkets(tx, marketsTableName)
	if err != nil {
		return fmt.Errorf("failed to read markets table: %w", err)
	}

	doTable := func(tableName string) error {
		_, err = tx.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN epoch_gap INT4 DEFAULT -1;", tableName))
		return err
	}

	log.Infof("Adding epoch_gap column to cancel tables for %d markets", len(mkts))

	for _, mkt := range mkts {
		if err := doTable(mkt.Name + "." + cancelsArchivedTableName); err != nil {
			return err
		}
		if err := doTable(mkt.Name + "." + cancelsActiveTableName); err != nil {
			return err
		}
	}
	return nil
}

// v6Upgrade creates the bonds table and its indexes on account_id and lockTime.
func v6Upgrade(tx *sql.Tx) error {
	bondsCreated, err := createTableStmt(tx, internal.CreateBondsTableV0, publicSchema, bondsTableName)
	if err != nil {
		return fmt.Errorf("failed to create bonds table: %w", err)
	}
	if bondsCreated {
		log.Infof("Created new %q table", bondsTableName)
	} else {
		log.Warnf("Unexpected existing %q table!", bondsTableName)
	}

	namespacedBondsTable := publicSchema + "." + bondsTableName
	err = createIndexStmt(tx, internal.CreateBondsAcctIndexV0, indexBondsOnAccountName, namespacedBondsTable)
	if err != nil {
		return fmt.Errorf("failed to index bonds table on account: %w", err)
	}

	err = createIndexStmt(tx, internal.CreateBondsLockTimeIndexV0, indexBondsOnLockTimeName, namespacedBondsTable)
	if err != nil {
		return fmt.Errorf("failed to index bonds table on lock time: %w", err)
	}

	// drop the accounts.broken_rule column
	namespacedAccountsTable := publicSchema + "." + accountsTableName
	_, err = tx.Exec(fmt.Sprintf("ALTER TABLE %s DROP COLUMN IF EXISTS broken_rule;", namespacedAccountsTable))
	if err != nil {
		return fmt.Errorf("failed to drop the accounts.broken_rule column: %w", err)
	}

	return nil
}

func v7Upgrade(tx *sql.Tx) error {
	const columnName = "reputation_ver"
	const tableName = publicSchema + "." + accountsTableName
	// Create the column, setting existing entries to false.
	query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s INT2 DEFAULT 0;", tableName, columnName)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("error adding reputation_ver column: %w", err)
	}
	// New entries should be true.
	query = fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET DEFAULT 1;", tableName, columnName)
	if _, err := tx.Exec(query); err != nil {
		return fmt.Errorf("error updating reputation_ver default value: %w", err)
	}
	return nil
}

// DBVersion retrieves the database version from the meta table.
func DBVersion(db *sql.DB) (ver uint32, err error) {
	err = db.QueryRow(internal.SelectDBVersion).Scan(&ver)
	return
}

func setDBVersion(db sqlExecutor, ver uint32) error {
	res, err := db.Exec(internal.SetDBVersion, ver)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("set the DB version in %d rows instead of 1", n)
	}
	return nil
}

func upgradeDB(ctx context.Context, db *sql.DB) error {
	// Get the DB version from the meta table. Nonexistent meta table or
	// meta.schema_version column implies v0, the upgrade from which adds the
	// table and schema_version column.
	var current uint32
	found, err := tableExists(db, metaTableName)
	if err != nil {
		return err
	}
	if found {
		found, err = columnExists(db, "public", metaTableName, "schema_version")
		if err != nil {
			return err
		}
		if found {
			current, err = DBVersion(db)
			if err != nil {
				return fmt.Errorf("failed to get DB version: %w", err)
			}
		} // else v1 upgrade creates meta.schema_version column
	} // else v1 upgrade creates meta table

	if current == dbVersion {
		log.Infof("DCRDEX database ready at version %d", dbVersion)
		return nil // all upgraded
	}

	if current > dbVersion {
		return fmt.Errorf("current DB version %d is newer than highest recognized version %d",
			current, dbVersion)
	}

	runUpgradeTx := func(targetVer uint32, up func(db *sql.Tx) error) error {
		// Canceling the context automatically rolls back the transaction.
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			// On error, rollback the transaction unless ctx was canceled
			// (sql.ErrTxDone) because then rollback is automatic. See the
			// (*sql.DB).BeginTx docs.
			if err == nil || errors.Is(err, sql.ErrTxDone) {
				return
			}
			log.Warnf("Rolling back upgrade to version %d", targetVer-1)
			errRollback := tx.Rollback()
			if errRollback != nil {
				log.Errorf("Rollback failed: %v", errRollback)
			}
		}()

		if err = up(tx); err != nil {
			return fmt.Errorf("failed to upgrade to db version %d: %w", targetVer, err)
		}

		if err = setDBVersion(tx, targetVer); err != nil {
			return fmt.Errorf("failed to set new DB version %d: %w", targetVer, err)
		}

		err = tx.Commit() // for the defer
		return err
	}

	log.Infof("Upgrading DB scheme from %d to %d", current, len(upgrades))
	for i, up := range upgrades[current:] {
		targetVer := current + uint32(i) + 1
		log.Debugf("Upgrading DB scheme to %d...", targetVer)
		if err = runUpgradeTx(targetVer, up); err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return fmt.Errorf("upgrade cancelled (rolled back to version %d)", current+uint32(i))
			}
			return err
		}
	}

	current, err = DBVersion(db)
	if err != nil {
		return fmt.Errorf("failed to get DB version: %w", err)
	}
	log.Infof("Upgrades complete. DB is at version %d", current)
	return nil
}
