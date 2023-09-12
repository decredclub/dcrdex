//go:build !firolive

package firo

import (
	"encoding/hex"
	"testing"

	"decred.org/dcrdex/dex"
)

var (
	regtestNoAlgoBlock, _ = hex.DecodeString(
		"01000030941ebce22649170c39611357e5c5cbc715b6a4a8ffc19f7efd0ad839abbb26" +
			"55fd11d087ba44d9f1e5fe1558b705ca1ac300b550c0a5b9f22e9f4c13c20a4ee62e0c" +
			"ea64ffff7f200000000001010000000100000000000000000000000000000000000000" +
			"00000000000000000000000000ffffffff050260010101ffffffff0600cb4c00010000" +
			"001976a914785c1c768cc6161bbfe9d00a51ba40221e84c5b088ac00e1f50500000000" +
			"1976a914296134d2415bf1f2b518b3f673816d7e603b160088ac00e1f5050000000019" +
			"76a914e1e1dc06a889c1b6d3eb00eef7a96f6a7cfb884888ac00e1f505000000001976" +
			"a914ab03ecfddee6330497be894d16c29ae341c123aa88ac00a3e111000000001976a9" +
			"144281a58a1d5b2d3285e00cb45a8492debbdad4c588ac00e1f505000000001976a914" +
			"1fd264c0bb53bd9fef18e2248ddf1383d6e811ae88ac00000000",
	)
	// block 29_000
	testnetMTPBlock, _ = hex.DecodeString(
		"00100020f7c22ae9fcb01a1f985435b671a074fd90d8f47a635a33b141a2463d483380" +
			"e3120a77bb0dd29f81195dfbae6a6c891f6296f7b3fc3595c0a6459a9a9d2afd6543ef" +
			"83605a54031f4511000000100000eea5f091d764a4ed653c4115cd3fb22be46b7a7059" +
			"ffc528443fd29bd1510200000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000103000500010000000000000000" +
			"000000000000000000000000000000000000000000000000ffffffff050248710101ff" +
			"ffffff0340be402500000000232102d15462d90b9bda92843342008b1a8badecbacbe2" +
			"f40c56e0316f5be4e1575a8cace0052d0b000000001976a914cfae8d3625968ac3230b" +
			"500a274669937627de4c88ac60b8131a000000001976a91466e5d30df1c81c80115a4c" +
			"217f63b42c407274a288ac00000000460200487100000a34e4cdee8b6da853ad822ab7" +
			"dffb8b0a48e9ea13d7a15a78f92acf209b2e27516d3ac82406daa19869646438b7a794" +
			"b28b83cef27e2cea1625b55c81c733e1",
	)
	// block 49_999
	// This block purposely has two outputs so that we are testing that we are
	// properly handling vExtraData at the end of transaction 0. If we were not
	// handling vExtraData correctly, transaction 1 would not parse.
	testnetProgPOWBlock, _ = hex.DecodeString(
		"001000205563b46ee309fc965da0b84dbee54dbad12948ffcdf33a618cff54b9073a59" +
			"f87e11a8b427e46190a9a63f8de85fe97762ba65e024e5a371010726388b3bd3ce9ab2" +
			"5261c750271c4fc300004d3945540080a04f1bbeffb1890956c41454c844b02f6cbc25" +
			"ebef986ce15b5b8f007dc70ea940e80203000500010000000000000000000000000000" +
			"000000000000000000000000000000000000ffffffff1b034fc30000000000049ab252" +
			"61024fa00a2f4d696e74506f6e642fffffffff0340be4025000000001976a91420e7c4" +
			"6d8b6678a5c1789e2d13b046a19060640288ac60b8131a000000001976a91466e5d30d" +
			"f1c81c80115a4c217f63b42c407274a288ace0052d0b000000001976a914cfae8d3625" +
			"968ac3230b500a274669937627de4c88ac000000004602004fc300001d0dfd4c026fb4" +
			"907e090386bfc1faf738dff87c825c1495cdd789fe27285d8ecca4fad477f3e165e35b" +
			"faac65779f45bb25a9e9b63c927036b883a11f96573303000600000000000000fd4901" +
			"01004fc300000100016f0ddcfbb8ef363fee4e09a2d56dccd690b46f486e9986a54c3c" +
			"c2d700bb24b63200000000000000320000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000000000000000000000000000000000000000000000" +
			"0000000000000000000000000000",
	)
	// Zeroth output of coinbase tx on testnet seems to be constanst.
	testnetValue00 int64 = 625000000
	noAlgoValue00  int64 = 4300000000
)

func TestDeserializeFiroBlock(t *testing.T) {
	// blockB, _ := hex.DecodeString("0010002043b5f66819ee245908141baefdac2dbdcd806ac250561ffff97ffda4ddf3cfce592af0eae16251b447ce5f57d61615d231b362db44138a8875ba799736e9a3e8ae512262ffff00208038010062020000000000007c92c3ff1ce62bdfdf727a56f5811a810a61a4d3be12260044e00532fe636aac0103000500010000000000000000000000000000000000000000000000000000000000000000ffffffff06038038010101ffffffff0340be4025000000001976a914eeb6d4eca74d16d4a24d97b70c4aada73bf6ce8388ace0052d0b000000001976a914cfae8d3625968ac3230b500a274669937627de4c88ac60b8131a000000001976a914839c8366952ac0acfa3c443692ea7d1b9251aeb388ac0000000046020080380100f19a7535df1c311432c54abf03715c6e18470bd0a0b9c1d30af595a68e966b44d920ea21c0b34a06c68ccba686855bfc020050efc368d6e7102fa0d2ba2b894d")
	for _, test := range []*struct {
		name            string
		serializedBlock []byte
		coinbaseValue0  int64
		net             dex.Network
	}{
		{"regtest - No algo", regtestNoAlgoBlock, noAlgoValue00, dex.Regtest},
		{"testnet - MTP", testnetMTPBlock, testnetValue00, dex.Testnet},
		{"testnet - ProgPOW", testnetProgPOWBlock, testnetValue00, dex.Testnet},
	} {
		t.Run(test.name, func(t *testing.T) {
			if blk, err := deserializeFiroBlock(test.serializedBlock, test.net); err != nil {
				t.Fatalf("deserializeFiroBlock error: %v", err)
			} else {
				if len(blk.Transactions) == 0 {
					t.Fatal("no transactions")
				}
				cbTx := blk.Transactions[0]
				if len(cbTx.TxOut) == 0 {
					t.Fatalf("no txouts")
				}
				txOut := cbTx.TxOut[0]
				if txOut.Value != test.coinbaseValue0 {
					t.Fatalf("wrong value for zeroth output of coinbase tx: wanted %d, got %d", test.coinbaseValue0, txOut.Value)
				}
			}
		})
	}
}