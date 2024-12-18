package locales

import "decred.org/dcrdex/client/intl"

var ZhCN = map[string]*intl.Translation{
	"Language":                       {T: "zh-CN"},
	"Markets": {T: "市场"}, // unused
	"Wallets": {T: "钱包"}, // unused
	"Notifications": {T: "通知"},
	"Recent Activity": {T: "近期活动"},
	"Sign Out": {T: "退出"},
	"Order History": {T: "历史订单"},
	"load from file": {T: "从文件加载"},
	"loaded from file": {T: "从文件加载"},
	"defaults": {T: "默认"},
	"Wallet Password": {T: "钱包密码"},
	"w_password_helper": {T: "这是您使用钱包软件配置的密码。"},
	"w_password_tooltip": {T: "如果您的钱包不需要密码，请将密码清空。"},
	"App Password": {T: "应用程序密码"},
	"Add": {T: "添加"},
	"Unlock": {T: "解锁"},
	"Rescan": {T: "重新扫描"},
	"Wallet": {T: "钱包"},
	"app_password_reminder": {T: "执行敏感钱包操作时，始终需要您的应用程序密码。"},
	"DEX Address": {T: "DEX 地址"},
	"TLS Certificate": {T: "TLS 证书"},
	"remove": {T: "移除"},
	"add a file": {T: "添加文件"},
	"Submit": {T: "提交"},
	"Skip Registration": {T: "跳过注册"},
	"Confirm Registration": {Version: 1,T: "确认注册与绑定"},
	"app_pw_reg": {Version: 1,T: "请输入您的应用密码以确认 DEX 注册与债券创建。"},
	"reg_confirm_submit": {Version: 1,T: "当您提交此表单时，您钱包中的资金将暂时锁定于忠诚债券合约，未来您可以赎回。"},
	"bond_strength": {T: "债券强度"},
	"target_tier": {T: "目标等级"},
	"target_tier_tooltip": {T: "这是您希望维持的账户等级。如果您希望禁用等级维持，请设置为零（不再发布新债券）。"},
	"compensation_tooltip": {T: "启用发布额外债券以弥补受到惩罚的等级。"},
	"Actual Tier": {T: "当前等级"},
	"Penalties": {T: "惩罚"},
	"Change Tier": {T: "更改等级"},
	"Limit Bonus": {T: "限制奖金"},
	"Score": {T: "分数"},
	"Confirm Bond Options": {T: "确认债券选项"},
	"provided_markets": {T: "此 DEX 提供以下市场："},
	"accepted_fee_assets": {T: "此 DEX 接受以下债券资产："},
	"base_header": {T: "基准"},
	"quote_header": {T: "报价"},
	"lot_size_headsup": {T: "所有交易必须是倍数。"},
	"Password": {T: "密码"},
	"Register": {T: "注册"},
	"Authorize Export": {T: "授权导出"},
	"export_app_pw_msg": {T: "请输入您的应用密码以确认导出账户"},
	"Disable Account": {T: "禁用账户"},
	"disable_dex_server": {T: "此 DEX 服务器可以在未来任何时间在设置页面重新启用。", Version: 1},
	"Authorize Import": {T: "授权导入"},
	"app_pw_import_msg": {T: "请输入您的应用密码以确认账户导入"},
	"Account File": {T: "账户文件"},
	"Change Application Password": {Version: 1,T: "更改密码"},
	"Current Password": {T: "当前密码"},
	"New Password": {T: "新密码"},
	"Confirm New Password": {T: "确认新密码"},
	"cancel_no_pw": {T: "提交取消订单以处理剩余"},
	"cancel_remain": {T: "剩余金额可能会在取消订单匹配之前发生变化。"},
	"Log In": {Version: 1,T: "解锁"},
	"epoch": {T: "纪元"},
	"price": {T: "价格"},
	"volume": {T: "交易量"},
	"volume_24": {T: "24小时交易量"},
	"High": {T: "最高"},
	"Low": {T: "最低"},
	"buys": {T: "买入"},
	"sells": {T: "卖出"},
	"Buy Orders": {T: "买单"},
	"Quantity": {T: "数量"},
	"Rate": {T: "汇率"},
	"Limit Order": {T: "限价单"},
	"Market Order": {T: "市价单"},
	"create_account_to_trade": {T: "创建账户以进行交易"},
	"need_to_register_msg": {T: "您需要在 <span id=\"unregisteredDex\"></span> 创建账户才能交易。"},
	"Create Account": {T: "创建账户"},
	"reg_status_msg": {T: "为了在 <span id=\"regStatusDex\" class=\"text-break\"></span> 进行交易，您的待处理债券需要确认。"},
	"posting_bonds_shortly": {T: "正在创建债券..."},
	"bond_creation_pending_msg": {T: "为了在 <span id=\"postingBondsDex\" class=\"text-break\"></span> 进行交易，债券将很快被创建。"},
	"action_required_to_trade": {T: "交易需要操作"},
	"acct_tier_post_bond": {T: "您的账户等级是 <span id=\"acctTier\"></span>。您需要发布新的债券才能交易。"},
	"enable_bond_maintenance": {T: "在 DEX 设置页面启用债券维护。"},
	"Buy": {T: "买入"},
	"Sell": {T: "卖出"},
	"lot_size": {T: "批量大小"},
	"Rate Step": {T: "汇率步骤"},
	"Max": {T: "最大"},
	"lot": {T: "批次"},
	"min trade is about": {T: "最低交易大约是"},
	"immediate_explanation": {T: "如果订单在下一个匹配周期内没有完全匹配，任何未匹配的数量将不会被预定或再次匹配。仅限于买方订单。"},
	"Immediate or cancel": {T: "立即或取消"},
	"Balances": {T: "余额"},
	"outdated_tooltip": {T: "余额可能已过时。请连接钱包以刷新。"},
	"available": {T: "可用"},
	"connect_refresh_tooltip": {T: "点击连接并刷新"},
	"add_a_wallet": {T: `添加一个 <span data-tmpl="addWalletSymbol"></span> 钱包`},
	"locked": {T: "已锁定"},
	"immature": {T: "未成熟"},
	"fee balance": {T: "费用余额"},
	"Sell Orders": {T: "卖单"},
	"Your Orders": {T: "您的订单"},
	"sweep_orders": {T: "隐藏已完全执行的订单"},
	"sweep_order": {T: "隐藏此已完全执行的订单"},
	"Recent Matches": {T: "近期匹配"},
	"Type": {T: "类型"},
	"Side": {T: "方向"},
	"Age": {T: "年龄"},
	"Filled": {T: "已填充"},
	"Settled": {T: "已结算"},
	"Status": {T: "状态"},
	"view order history": {T: "查看订单历史"},
	"cancel_order": {T: "取消订单"},
	"order details": {T: "订单详情"},
	"verify_order": {T: `验证 <span id="vSideHeader"></span> 订单`},
	"prevent_temporary_overlocking": {T: "防止资金暂时过度锁定"},
	"You are submitting an order to": {T: "您正在提交一个订单到"},
	"at a rate of": {T: "以汇率"},
	"for a total of": {T: "总共"},
	"verify_market": {T: "这是一个市场订单，将与订单簿上最佳可用订单匹配。根据当前市场中间价格，您可能会收到大约"},
	"auth_order_app_pw": {T: "使用您的应用密码授权此订单。"},
	"lots": {T: "批次"},
	"order_disclaimer": {Version: 1,T: `<span class="red">重要</span>: 交易需要时间来结算，您不能在结算完成之前关闭 <span class=brand></span> 软件，
	或 <span data-quote-ticker></span> 或 <span data-base-ticker></span> 区块链和/或钱包软件。
	结算可能需要几分钟，也可能需要几小时。`},
	"acknowledge_and_hide": {T: "确认并隐藏"},
	"show_disclaimer": {T: "查看警告"},
	"Order": {T: "订单"},
	"see all orders": {T: "查看所有订单"},
	"Exchange": {T: "交易所"},
	"Market": {T: "市场"},
	"Offering": {T: "报价"},
	"Asking": {T: "请求"},
	"Fees": {T: "费用"},
	"order_fees_tooltip": {T: "链上交易费用，通常由矿工收取。Decred DEX 不收取交易费用。"},
	"Matches": {T: "匹配"},
	"Match ID": {T: "匹配 ID"},
	"Time": {T: "时间"},
	"ago": {T: "之前"},
	"Cancellation": {T: "取消"},
	"Order Portion": {T: "订单部分"},
	"you": {T: "您"},
	"them": {T: "他们"},
	"Redemption": {T: "赎回"},
	"Refund": {T: "退款"},
	"Funding Coins": {T: "资金币种"},
	"Exchanges": {T: "交易所"},
	"apply": {T: "应用"},
	"Assets": {T: "资产"},
	"Trade": {T: "交易"},
	"Set App Password": {Version: 1,T: "设置新密码"},
	"reg_set_app_pw_msg": {Version: 1,T: "设置您的密码。此密码用于解锁 <span class=brand></span> 供使用。"},
	"Password Again": {T: "再次输入密码"},
	"Add a DEX": {Version: 1,T: "添加 DCRDEX 服务器"},
	"Pick a server": {T: "选择服务器"},
	"reg_ssl_needed": {T: "看起来我们没有为此 DEX 提供 SSL 证书。请添加服务器的证书以继续。"},
	"Dark Mode": {T: "深色模式"},
	"Show pop-up notifications": {T: "显示弹出通知"},
	"Account ID": {T: "账户 ID"},
	"Export Account": {T: "导出账户"},
	"simultaneous_servers_msg": {Version: 1,T: "<span class=brand></span> 支持同时使用任意数量的 DEX 服务器。"},
	"Change App Password": {T: "更改应用密码"},
	"enable_browser_ntfn": {T: "启用桌面通知"},
	"browser_ntfn_blocked": {T: "浏览器通知当前被阻止。请在浏览器中解锁此站点以接收通知。"},
	"enable_browser_ntfn_info": {T: "桌面通知即使在窗口不活动时也会出现。当您打开其他应用程序时，这会很有帮助，因为您会收到 DCRDEX 事件的通知。下面自定义您希望接收的通知类型。"},
	"Save Notifications": {T: "保存通知"},
	"Build ID": {T: "构建 ID"},
	"Connect": {T: "连接"},
	"Send": {T: "发送"},
	"Deposit": {T: "存款"},
	"Receive": {T: "接收"},
	"Lock": {T: "锁定"},
	"New Address": {T: "新地址"},
	"New Deposit Address": {T: "新存款地址"},
	"Address": {T: "地址"},
	"Amount": {T: "金额"},
	"Authorize the transfer with your app password.": {T: "使用您的应用密码授权转账。"},
	"Reconfigure": {T: "重新配置"},
	"pw_change_instructions": {T: "下面更改密码不会更改 <span class=brand></span> 的密码。使用此表单在您通过钱包软件直接更改密码后更新 <span class=brand></span>。"},
	"New Wallet Password": {T: "新钱包密码"},
	"pw_change_warn": {T: "注意：在进行活跃交易时更换钱包可能导致资金丢失。"},
	"Show more options": {T: "显示更多选项"},
	"seed_implore_msg": {T: "您应仔细写下您的应用种子并保存副本。如果您丢失对该机器或关键应用文件的访问权限，可以使用该种子恢复您的 DEX 账户和本地钱包。一些旧账户无法从种子恢复，无论是新账户还是旧账户，最好将账户密钥与种子分开备份。"},
	"View Application Seed": {T: "查看应用种子"},
	"Remember my password": {T: "记住我的密码"},
	"pw_for_seed": {T: "输入您的应用密码以显示种子。确保没有其他人能看到您的屏幕。"},
	"Asset": {T: "资产"},
	"Balance": {T: "余额"},
	"Actions": {T: "操作"},
		"Restoration Seed": {T: "恢复种子"},
	"Restore from seed": {T: "从种子恢复"},
	"Import Account": {T: "导入账户"},
	"no_wallet": {T: "没有钱包"},
	"Token on": {T: "代币开启"},
	"create_a_x_wallet": {T: "创建 <span data-asset-name=1></span> 钱包"},
	"dont_share": {T: "不要分享。不要丢失它。"},
	"Show Me": {T: "给我看"},
	"Wallet Settings": {T: "钱包设置"},
	"add_a_x_wallet": {T: `添加 <img data-tmpl="assetLogo" class="small-icon mx-1"> <span data-tmpl="assetName"></span> 钱包`},
	"ready": {T: "准备好"},
	"off": {T: "关闭"},
	"Export Trades": {T: "导出交易"},
	"change the wallet type": {T: "更改钱包类型"},
	"confirmations": {T: "确认"},
	"pick a different asset": {T: "选择不同的资产"},
	"Create": {T: "创建"},
	"1 Sync the Blockchain": {T: "1：同步区块链"},
	"Progress": {T: "进度"},
	"remaining": {T: "剩余"},
	"2 Fund your Wallet": {T: "2：为您的钱包充值"},
	"bond_lock": {T: "保证金锁定："},
	"bond_definition": {T: "保证金是暂时锁定在链上合同中的资金。合同到期后，您的钱包将收回这些资金。在主网中，资金被锁定 2 个月。"},
	"bonds_can_be_revoked": {T: "如果账户从事持续的破坏性行为（如取消交换），保证金可以被撤销。被撤销的保证金可以通过正常的交易活动重新激活。"},
	"bond_cost_breakdown": {T: `包括 <span data-tmpl="bondLockNoFees"></span> 作为保证金和 <span data-tmpl="bondLockFees"></span> 作为交易费用`},
	"bigger_bonds_higher_limit": {T: "您可以创建更大的保证金以提高您的交易等级，允许一次交易更大的数量。更大的保证金还增加了在交易特权暂停之前，违规的容量。"},
	"limits_reputation": {T: "通过进行正常的交易活动并成功完成交易，您可以提高声誉，从而增加交易限额。"},
	"wallet_bond_reserves": {T: "您的钱包会自动保留足够的资金，以保持您选择的交易等级的保证金，并会广播新保证金以替换过期的保证金。您可以在交易所的设置面板中调整您的交易等级。将交易等级设置为零可禁用您的账户（直到现有保证金到期为止）。"},
	"Got it": {T: "明白了"},
	"Trading Limits": {T: "交易限额"},
	"What is a fidelity bond": {T: "什么是保证金？"},
	"what_s_a_bond": {T: "什么是保证金？"},
	"order_form_remaining_limit": {T: `<span id="orderLimitRemain"></span> 份剩余在等级 <span id="orderTradingTier"></span> 交易限额`},
	"Parcel Size": {T: "包裹大小"},
	"Trading Limit": {T: "交易限额"},
	"Current Usage": {T: "当前使用量"},
	"score_factors": {T: "通过成功完成交易提高您的评分。未能完成交易将降低您的评分。"},
	"Bond amount": {T: "保证金额度"},
	"Reserves for tx fees": {T: "用于交易费用的储备资金，以维持您的保证金"},
	"Tx Fee Balance": {T: "交易费用余额："},
	"Your Deposit Address": {T: "您的钱包存款地址"},
	"Send enough for bonds": {T: `请确保您发送足够的资金以支付网络费用。您可以向您的钱包存入任意数量的资金，因为只有保证金金额会用于下一步。存款必须确认后才能继续。`},
	"Send enough with estimate": {T: `至少存入 <span data-tmpl="totalForBond"></span> <span class="unit">XYZ</span> 来覆盖保证金和费用。您可以向钱包存入任意数量的资金，因为只有所需金额会在下一步中使用。钱包可能需要对新存入的资金进行确认才能继续。`},
	"Send funds for token": {T: `存入至少 <span data-tmpl="tokenFees"></span> <span class="unit">XYZ</span> 和 <span data-tmpl="parentFees"></span> <span data-tmpl="parentUnit">XYZ</span> 以支付费用。您可以向钱包存入任意数量的资金，因为只有所需金额会在下一步中使用。存款必须确认后才能继续。`},
	"add a different server": {T: "添加不同的服务器"},
	"Add a custom server": {T: "添加自定义服务器"},
	"plus tx fees": {T: "+ 交易费用"},
	"Export Seed": {T: "导出种子"},
	"Total": {T: "总计"},
	"Trading": {T: "交易"},
	"Receiving Approximately": {T: "大约接收"},
	"Fee Projection": {T: "费用预测"},
	"details": {T: "详情"},
	"to": {T: "至"},
	"Options": {T: "选项"},
	"fee_projection_tooltip": {T: "如果网络条件在您的订单匹配前没有变化，总费用应在此范围内。"},
	"unlock_for_details": {T: "解锁您的钱包以获取订单详情和更多选项。"},
	"estimate_unavailable": {T: "订单估算和选项不可用"},
	"Fee Details": {T: "费用详情"},
	"estimate_market_conditions": {T: "最佳和最差情况下的估算基于当前的网络条件，可能会在订单匹配时发生变化。"},
	"Best Case Fees": {T: "最佳情况费用"},
	"best_case_conditions": {T: "最佳情况费用发生在整个订单在单次匹配中完成时。"},
	"Swap": {T: "交换"},
	"Redeem": {T: "赎回"},
	"Worst Case Fees": {T: "最差情况费用"},
	"worst_case_conditions": {T: "最差情况可能发生在订单分多次匹配时，跨越多个时期。"},
	"Maximum Possible Swap Fees": {T: "最大可能交换费用"},
	"max_fee_conditions": {T: "这是您在交换中支付的最高费用。费用通常按此费率的一个小部分计算。最大费用在订单提交后不会变化。"},
	"wallet_logs": {T: "钱包日志"},
	"accelerate_order": {T: "加速订单"},
	"acceleration_text": {T: "如果您的交换交易卡住了，您可以尝试通过额外的交易加速它们。当现有未确认交易的费用率过低，无法在下一个区块中被挖矿时，这种方法很有帮助，但如果区块只是在缓慢挖掘，则无效。提交此表单后，将创建一笔交易，使用更高的费用将您的交换启动交易的找零转给自己。您的交换交易的有效费用率将变为您选择的费用率。请选择足以被包括在下一个区块中的费用率。请通过区块浏览器确认费用率是否足够。" },
    "effective_swap_tx_rate": {T: "有效交换交易费用率" },
    "current_fee": {T: "当前建议的费用率" },
    "accelerate_success": {T: "成功提交交易：<span id=\"accelerateTxID\"></span>" },
    "accelerate": {T: "加速" },
    "acceleration_transactions": {T: "加速交易" },
    "acceleration_cost_msg": {T: "将有效费用率增加到 <span id=\"feeRateEstimate\"></span> 将花费 <span id=\"feeEstimate\"></span>" },
    "recent_acceleration_msg": {T: "您的最新加速交易只有 <span id=\"recentAccelerationTime\"></span> 分钟前！您确定要加速吗？" },
    "recent_swap_msg": {T: "您最早的未确认交换交易仅在 <span id=\"recentSwapTime\"></span> 分钟前提交！您确定要加速吗？" },
    "early_acceleration_help_msg": {T: "这对您的订单没有害处，但可能会浪费钱。加速仅在现有未确认交易的费用率过低，无法在下一个区块中被挖矿时才有帮助，而不是区块挖掘的速度较慢时。您可以通过关闭此弹窗并点击您之前的交易在区块浏览器中确认这一点。" },
    "recover": {T: "恢复" },
    "recover_wallet": {T: "恢复钱包" },
    "recover_warning": {T: "恢复钱包将把所有钱包数据移动到备份文件夹。您将需要等待钱包与网络同步，这可能会花费较长时间，然后才能再次使用钱包。" },
    "wallet_actively_used": {T: "钱包正在使用中！" },
    "confirm_force_message": {T: "此钱包正在管理订单。采取此操作后，钱包将需要很长时间才能重新同步，可能导致订单失败。仅在绝对必要时才采取此操作！" },
    "confirm": {T: "确认" },
    "cancel": {T: "取消" },
    "Update TLS Certificate": {T: "更新TLS证书" },
    "registered dexes": {T: "已注册的DEX：" },
    "successful_cert_update": {T: "成功更新证书！" },
    "update dex host": {T: "更新DEX主机" },
    "copied": {T: "已复制！" },
    "export_wallet": {T: "导出钱包" },
    "pw_for_wallet_seed": {T: "输入您的应用密码以显示钱包种子。确保没有其他人可以看到您的屏幕。如果有人获得钱包种子，他们将能够窃取您所有的资金。" },
    "export_wallet_disclaimer": {T: "<span class=\"warning-text\">在DEX中有活动交易时使用外部恢复的钱包可能导致交易失败并丧失资金。</span> 除非您是经验丰富的用户并且知道自己在做什么，否则不建议导出钱包。" },
    "export_wallet_msg": {T: "以下是需要恢复钱包到一些流行外部钱包所需的种子。不要在您有活动交易时使用外部钱包进行交易。" },
    "clipboard_warning": {T: "复制/粘贴钱包种子可能存在安全风险。请自行承担风险。" },
    "fiat_exchange_rate_sources": {T: "法币汇率来源" },
    "Synchronizing": {T: "同步中" },
    "wallet_wait_synced": {T: "钱包将在同步后创建" },
    "Create a Wallet": {T: "创建钱包" },
    "Wallet Type": {T: "钱包类型" },
    "Peer Count": {T: "节点数" },
    "Sync Progress": {T: "同步进度" },
    "Settings": {T: "设置" },
    "asset_name Markets": {T: "<span data-asset-name=1></span> 市场" },
    "Host": {T: "主机" },
    "Trading Tier": {T: "交易层级" },
    "Bond Lock": {T: "保证金锁仓" },
    "Fee Reserves": {T: "费用储备" },
    "Select your bond asset": {T: "选择您的保证资产" },
    "choose a different asset": {T: "选择其他资产" },
    "Choose your trading tier": {T: "选择您的交易层级" },
    "trading_tier_message": {T: "提高您的交易层级可以使您交易更大金额。交易限额也会随着声誉的增加而增长。" },
    "No Recent Activity": {T: "没有最近的活动" },
    "Recent asset_name Activity": {T: "最近的 <span data-asset-name=1></span> 活动" },
    "other_actions": {T: "其他操作" },
    "estimated_fee": {T: "预计费用" },
    "estimated_total_spend": {T: "预计总支出" },
    "estimated_balance": {T: "预计余额" },
    "max_estimated_send": {T: "最大预计发送" },
    "max_estimated_send_fee": {T: "最大预计发送费用" },
    "sending": {T: "发送中" },
    "transfer": {T: "转账" },
    "max_estimated_send_tooltip": {T: "这是在勾选“从发送金额中扣除费用”时，如果您提取当前余额，将接收到的预计金额。如果没有勾选扣除费用的选项，则这是您可以发送的最大预计金额。" },
    "Maker Swap": {T: "做市商交换" },
    "Taker Swap": {T: "接盘商交换" },
    "Maker Redemption": {T: "做市商赎回" },
    "Taker Redemption": {T: "接盘商赎回" },
    "Pending": {T: "待处理" },
    "Mempool": {T: "内存池" },
    "disable_wallet": {T: "禁用钱包" },
    "enable_wallet": {T: "启用钱包" },
    "disable_wallet_warning": {T: "注意：启动 <span class=brand></span> 软件时，此钱包将无法连接，且在启用之前无法使用。此操作还将禁用所有依赖此钱包的代币钱包。" },
    "enable_wallet_message": {T: "此钱包将恢复操作，可能需要一些时间进行同步。如果这是一个代币钱包，链的主资产钱包也将被启用。" },
    "Disabled": {T: "已禁用" },
    "txfee_not_available": {T: "交易费用当前不可用" },
    "Fee unavailable": {T: "费用不可用" },
    "fiat_exchange_rate_msg": {T: "不同的源可能提供不同资产子集的法币汇率。您应该选择所有可以接受的源，以便获得更多资产的法币汇率。具有多个来源法币汇率的资产将使用所有来源的平均值。注意：dcrdate只为BTC和DCR提供法币汇率。" },
    "delete_archived_records": {T: "删除已归档记录" },
    "date_time": {T: "日期和时间" },
    "delete_all_archived_records": {T: "未勾选则删除所有已归档记录。" },
    "show_archived_date_msg": {T: "指定保留最新记录的日期" },
    "archived_date_tooltip": {T: "在您指定的日期和时间之前创建的归档订单和匹配将从数据库中删除。" },
    "save_matches_to_file": {T: "保存匹配到CSV文件" },
    "save_orders_to_file": {T: "保存订单到CSV文件" },
    "save_orders_to_file_msg": {T: "可选：是否将删除的订单保存到bisonw数据目录的CSV文件中。默认值为false。" },
    "save_matches_to_file_msg": {T: "可选：是否将删除的匹配保存到bisonw数据目录的CSV文件中。默认值为false。" },
    // Market maker bot
    "Market Making": {T: "做市" },
    "Off": {T: "关闭" },
    "whats_a_market_maker": {T: "添加一个做市机器人来创建市场。做市时，您会在市场的订单簿两侧战略性地维护订单。除了潜在的利润外，做市商通过提供流动性并减少滑点来改善市场的健康。" },
    "Add a Market Maker Bot": {T: "添加做市机器人" },
    "Add another bot": {T: "添加另一个机器人" },
    "Market Maker Settings": {T: "做市商设置" },
    "strategy_percent_plus": {T: "间隙从盈亏平衡点开始，盈亏平衡点是买卖组合产生利润的最小差价，然后增加指定的差价，计算为中间差价的百分比。" },
    "strategy_percent": {T: "间隙设置为中间差价的指定百分比。" },
    "strategy_absolute_plus": {T: "间隙从盈亏平衡点开始，盈亏平衡点是买卖组合产生利润的最小差价，然后在两侧增加指定的利率差异。" },
    "strategy_absolute": {T: "差价设置为中间差价左右各指定的值，总差价为指定值的两倍。" },
    "strategy_multiplier": {T: "间隙设置为盈亏平衡点的指定倍数，盈亏平衡点是买卖组合产生利润的最小差价。" },
    "Make a Market": {T: "创建市场！" },
    "Editing Program": {T: "编辑程序" },
    "exit edit mode": {T: "退出编辑模式" },
    "per side": {T: "每侧" },
    "Start_loudly": {T: "开始！" },
    "Update": {T: "更新" },
    "Show other settings": {T: "显示其他设置" },
    "Hide settings": {T: "隐藏设置" },
    "lot_commit_bullet": {T: "总承诺为 2 x 批量" },
    "funds_split_bullet": {T: "起始资金可以在一侧或分开" },
    "target_maint_bullet": {T: "程序保持当前未成交的批量订单" },
    "no_limit_bullet": {T: "结算批量的数量没有限制" },
    "Your Programs": {T: "您的程序" },
    "No programs to display": {T: "没有可显示的程序" },
    "Running": {T: "运行中" },
    "Pause": {T: "暂停" },
    "Paused": {T: "已暂停" },
    "Start": {T: "开始" },
    "Drift tolerance": {T: "漂移容忍度" },
    "Order persistence": {T: "订单持久性" },
    "Oracle bias": {T: "预言机偏差" },
    "Multiplier": {T: "倍数" },
    "Oracle weight": {T: "预言机权重" },
    "Configure": {T: "配置" },
    "Retire": {T: "退休" },
    "Lots": {T: "批量" },
    "show advanced options": {T: "显示高级选项" },
    "hide advanced options": {T: "隐藏高级选项" },
    "manage_peers": {T: "管理对等方" },
    "enter_peer_address": {T: "输入对等方地址" },
    "add_peer": {T: "添加对等方" },
    "address": {T: "地址" },
    "source": {T: "来源" },
    "connected": {T: "已连接" },
    "Remove": {T: "移除" },
    "unready_wallets_msg": {T: "您的钱包必须连接并解锁才能处理交易。请尽快解决！" },
    "Error": {T: "错误" },
    "configuration guide": {T: "配置指南" },
	// Approve/Unapprove Tokens
	"Approve": {T: "批准"},
    "approve_token_text": {Version: 1,T: "为了交易代币，您必须批准交换合约代为处理代币。这是每个代币的一次性操作。<br/><br/>批准交易的估计费用为<b><span data-tmpl='feeEstimate'></span></b>"},
    "token_approval_tx_msg": {T: "您的代币批准已提交，交易ID为："},
    "approval_required_buy": {T: "购买之前需要代币批准"},
    "approval_required_sell": {T: "出售之前需要代币批准"},
    "approval_required_both": {T: "交易之前需要代币批准"},
    "disallow_token": {T: "移除代币授权"},
    "unapprove_token_for": {T: "移除代币授权 for <span id='tokenVersionTableAssetSymbol'></span>"},
    "unapprove_token_version": {T: "移除代币授权 for <span id='tokenAllowanceRemoveSymbol'></span>，版本 <span id='tokenAllowanceRemoveVersion'></span>"},
    "unapprove_token_text": {T: "如果您移除此版本的交换合约授权，您将无法继续交易，直到重新授权。<br/><br/>移除授权的估计交易费用为<b><span id='unapprovalFeeEstimate'></span></b>"},
    "version": {T: "版本"},
    "used_by_dex": {T: "被DEX使用"},
    "no_token_allowances": {T: "您尚未为此代币授予任何交换合约的授权。"},
    "token_unapproval_tx_msg": {T: "您的代币授权已被移除，交易ID为："},
    "approval_change_pending": {T: "授权更改待处理"},
    // Init page
    "Quick Placements": {Version: 1,T: "快速配置"},
    "quickconfig_server_header": {T: "我们可以连接到这些服务器。"},
    "quickconfig_wallet_header": {T: "我们现在可以激活这些本地钱包。从钱包视图中配置其他钱包。"},
    "quickconfig_server_error_header": {T: "连接以下服务器时遇到错误。您可以从设置视图中再次尝试。"},
    "quickconfig_wallet_error_header": {T: "创建以下钱包时遇到错误。钱包可以在钱包视图中配置。"},
    "Continue": {T: "继续"},
    "Backup App Seed": {T: "备份应用种子"},
    "seed_backup_implore": {T: "强烈建议您现在备份应用种子。应用种子对恢复您的保密债券和本地钱包至关重要。"},
    "Backup Now": {T: "立即备份"},
    "Skip this step for now": {T: "暂时跳过此步骤"},
    "save_seed_instructions": {T: "将以下种子保存到安全的地方。不要与任何人共享。"},
    "Done": {T: "完成"},
    "gap_strategy_multiplier": {T: "乘数"},
    "gap_strategy_absolute": {T: "绝对值"},
    "gap_strategy_absolute_plus": {T: "绝对值加"},
    "gap_strategy_percent": {T: "百分比"},
    "gap_strategy_percent_plus": {T: "百分比加"},
    "settings_updated": {T: "设置已更新！"},
    "add_a_bot": {T: "添加机器人"},
    "start_market_making": {T: "开始"},
    "stop_market_making": {T: "停止"},
    "enabled": {T: "已启用"},
    "running": {T: "运行中"},
    "bot_type": {T: "机器人类型"},
    "base_balance": {T: "基础余额"},
    "quote_balance": {T: "报价余额"},
    "enable_rebalance_tooltip": {T: "根据需要在钱包和CEX之间转移资金，无需提示"},
    "cex_alloc_tooltip": {T: "最初分配给此机器人的余额在交易所上的数量"},
    "arb_minbal_tooltip": {T: "在发起存款或提取之前允许的最小余额"},
    "arb_transfer_tooltip": {T: "可转移的最小金额"},
    "Select a Market": {T: "选择市场"},
    "Arbitrage Rebalance": {T: "套利再平衡"},
    "Minimum Balance": {T: "最小余额"},
    "Minimum Transfer": {T: "最小转账"},
    "update_settings": {Version: 1,T: "保存设置"},
    "create_bot": {T: "创建机器人"},
    "reset_settings": {T: "重置设置"},
    "gap_strategy": {T: "差距策略"},
    "gap_strategy_tooltip": {T: "乘数：盈亏平衡差价的倍数，\n绝对值：设置绝对买卖价差，\n绝对值加：绝对值加上盈亏平衡差价，\n百分比：差距设置为现货价格的百分比，\n百分比加：百分比加上盈亏平衡差价"},
    "buy_placements": {T: "购买配置"},
    "buy_placements_tooltip": {T: "为机器人定义购买配置。如果余额不足以下所有订单，机器人将按优先级顺序下单。"},
    "sell_placements": {T: "出售配置"},
    "sell_placements_tooltip": {T: "为机器人定义出售配置。如果余额不足以下所有订单，机器人将按优先级顺序下单。"},
    "priority": {T: "优先级"},
    "use_oracles": {T: "使用预言机"},
    "use_oracles_tooltip": {T: "如果勾选，机器人将结合其他交易所的汇率和DEX中间差价来确定订单配置汇率。"},
    "oracle_weighting_tooltip": {T: "在确定订单配置汇率时给予预言机汇率的权重。DEX市场薄弱时，使用较高的权重以避免操纵。"},
    "oracle_bias_tooltip": {T: "在使用预言机汇率确定订单配置汇率之前，按此金额调整预言机汇率。"},
    "empty_market_rate": {T: "空市场汇率"},
    "empty_market_rate_tooltip": {T: "如果DEX市场为空且没有可用的预言机，可以选择提供一个汇率。"},
    "quote_asset_balance_tooltip": {T: "为此机器人分配的报价资产数量。如果该范围的上限小于100%，意味着其余余额已分配给其他机器人。"},
    "drift_tolerance_tooltip": {T: "订单允许偏离理想价格的最大幅度，超过此范围的订单将被取消并重新预定。"},
    "order_persistence_tooltip": {T: "CEX订单未立即匹配时允许预定的时间"},
    "no_balance_available": {T: "无可用余额"},
    "wallet_settings": {T: "钱包设置"},
    "cex_rebalance": {T: "CEX再平衡设置"},
    "Oracles": {T: "预言机"},
    "loading_oracles": {T: "加载预言机..."},
    "no_oracles": {T: "此市场没有可用的预言机"},
    "fiat_rates": {T: "法币汇率"},
    "market_making_running": {T: "市场做市正在运行"},
    "cannot_manually_trade": {T: "在市场做市运行时，您无法手动下单"},
    "back": {T: "返回"},
    "current_tier_tooltip": {T: "由活跃债券表示的级别。提高目标级别以提升交易限额，并补偿任何惩罚（如果有）。"},
    "Reset App Password": {T: "重置应用密码"},
    "reset_app_pw_msg": {T: "使用您的应用种子重置应用密码。如果提供正确的应用种子，您可以使用新密码重新登录。"},
    "Forgot Password": {T: "忘记密码？"},
    "subtract_fees_from_amount": {T: "从发送金额中扣除费用。"},
    "Instructions:": {T: "说明："},
    "URL": {T: "网址"},
	
	// Staking UI
		"Price": {T: "价格" },
	"Ticket": {T: "票" },
	"Staking": {T: "质押" },
	"Active tickets": {T: "活动票" },
	"Currently Queued": {T: "当前排队" },
	"Immature tickets": {T: "未成熟票" },
	"Queued tickets": {T: "排队票" },
	"Tickets bought": {T: "购买的票" },
	"Total rewards": {T: "总奖励" },
	"Votes cast": {T: "已投票" },
	"VSP": {T: "VSP" },
	"select_vsp_from_list": {T: "选择投票服务提供商" },
	"Or add custom VSP URL:": {T: "或添加自定义VSP URL：" },
	"Custom VSP URL": {T: "自定义VSP URL" },
	"Set Votes": {T: "设置投票" },
	"agendas": {T: "议程" },
	"Agendas": {T: "议程" },
	"treasury spends": {T: "财政支出" },
	"staking_unavailable_for_rpc": {T: "RPC钱包使用SPV时，无法进行质押" },
	"Purchase Tickets": {T: "购买票" },
	"Tickets": {T: "票" },
	"Ticket Price": {T: "票价" },
	"Vote Reward": {T: "投票奖励" },
	"fee_rate_percent": {T: "费用率 (%)" },
	"Live Tickets": {T: "实时票" },
	"Current Price": {T: "当前价格" },
	"Available Balance": {T: "可用余额" },
	"ticket_qty_label": {T: "多少票？" },
	"Ticket History": {T: "票历史" },
	"No tickets to show": {T: "没有票显示" },
	"Treasury Spends": {T: "财政支出" },
	"No": {T: "否" },
	"Yes": {T: "是" },
	"Treasury Keys": {T: "财政密钥" },
	"bonds_pending_refund": {T: "待退款债券" },
	"asset_name tx_history": {T: "<span data-asset-name=1></span> 交易历史" },
	"load_earlier_transactions": {T: "加载早期交易..." },
	"ID": {T: "ID" },
	"no_tx_history": {T: "没有交易显示" },
	"tx_details": {T: "交易详情" },
	"fee": {T: "费用" },
	"tx_id": {T: "交易ID" },
	"bond_id": {T: "债券ID" },
	"locktime": {T: "锁定时间" },
	"recipient": {T: "收件人" },
	"block": {T: "区块" },
	"timestamp": {T: "时间戳" },
	"nonce": {T: "随机数" },
	"Enable Privacy": {T: "启用隐私" },
	"Privacy": {T: "隐私" },
	"Enable": {T: "启用" },
	"Disable Privacy": {T: "禁用隐私" },
	"Configure Privacy": {T: "配置隐私" },
	"Privacy active": {T: "隐私已启用" },
	"Privacy off": {T: "隐私已关闭" },
	"staking_disabled": {T: "票务购买已被<span id=\"extensionModeAppName\"></span>禁用。" },
	"loading privacy status": {T: "加载隐私状态" },
	"mixing_pw_prompt": {T: "输入密码以解锁钱包并启用隐私" },
	"cspp_addr": {T: "CSPP服务器地址" },
	"cspp_how": {T: "StakeShuffle创建的输出无法与先前的链上活动确定关联。虽然这种隐私方式下的输出金额仍然可见，但摧毁了可追溯性，意味着无法从公开数据推断出你的链上历史。" },
	"privacy_intro": {T: "启用隐私时，您的所有资金将通过使用StakeShuffle协议的地址历史模糊服务发送。" },
	"decred_privacy": {T: "Decred的隐私形式尤其强大，因为Decred钱包将隐私与质押结合，形成了一个始终较大的匿名集，这是隐私的关键特性。" },
	"privacy_optional": {T: "隐私是完全可选的，可以随时禁用。启用隐私会产生额外的交易费用，但这些费用历来相对较低。" },
	"privacy_unlocked": {T: "钱包在混合时必须保持解锁状态。" },
	"bots_running_view_only": {T: "机器人正在运行。您处于仅查看模式。" },
	"select_a_cex_prompt": {T: "选择一个交易所以启用套利" },
	"Market not available": {T: "市场不可用" },
	"bot_profit_title": {T: "选择您的利润阈值" },
	"bot_profit_explainer": {T: "套利机器人只会尝试那些能够带来至少这个水平利润的交易。" },
	"configure_cex_prompt": {T: "配置您的交易所API以启用套利功能。" },
	"API Key": {T: "API 密钥" },
	"API Secret": {T: "API 密钥" },
	"Available": {T: "可用" },
	"Locked": {T: "锁定" },
	"profit_loss": {T: "利润/亏损" },
		"base_change": {T: "基础变动" },
	"quote_change": {T: "报价变动" },
	"base_fees": {T: "基础费用" },
	"quote_fees": {T: "报价费用" },
	"run_logs": {T: "运行日志" },
	"previous_run_logs": {T: "查看之前的运行日志" },
	"previous_mm_runs": {T: "之前的市场制造运行" },
	"start_time": {T: "开始时间" },
	"end_time": {T: "结束时间" },
	"logs": {T: "日志" },
	"logs_for": {T: "日志：" },
	"view_logs": {T: "查看日志" },
	"swap_transaction": {T: "交换交易" },
	"redeem_transaction": {T: "兑换交易" },
	"refund_transaction": {T: "退款交易" },
	"funding_transaction": {T: "资金交易" },
	"transaction_id": {T: "交易ID" },
	"credited_amt": {T: "贷方金额" },
	"amt_received": {T: "收到金额" },
	"dex_order_details": {T: "DEX订单详情" },
	"cex_order_details": {T: "CEX订单详情" },
	"deposit_details": {T: "存款详情" },
	"withdrawal_details": {T: "取款详情" },
	"base_filled": {T: "基础已填充" },
	"quote_filled": {T: "报价已填充" },
	"Details": {T: "详情" },
	"Fee Allocation": {T: "费用分配" },
	"Booking reserves": {T: "预定储备" },
	"Swap funding": {T: "交换资金" },
	"swaps": {T: "交换" },
	"register_to_enable_mm": {T: `注册与<span data-tmpl="host"></span>启用交易` },
	"No markets available": {T: "没有可用市场" },
	"Arb": {T: "套利" },
	"Choose Your Bot": {T: "选择你的机器人" },
	"only_basic_mm_no_more_cexes": {T: "仅提供基础市场制造机器人。要启用套利选项，请配置您的交易所API。" },
	"only_basic_mm_more_cexes": {T: "仅为此市场提供基础市场制造机器人。" },
	"configure_more_cexes": {T: "配置其他交易所以最大化您的选择。" },
	"Basic Market Maker": {T: "基础市场制造商" },
	"mm_plus_arb": {T: "市场制造商 + 套利" },
	"Basic Arbitrage": {T: "基础套利" },
	"Fix errors": {T: "修复错误" },
	"Commited": {T: "已承诺" },
	"Required": {T: "必需" },
	"OK": {T: "确定" },
	"remote exchange gap": {T: "远程交易所差价"},
	"Match buffer": {T: "匹配缓冲区" },
	"Price increment": {T: "价格增量" },
	"Profit threshold": {T: "利润阈值" },
	"Lot Size USD": {T: "每手大小（美元）" },
	"Lot Size": {T: "每手大小" },
	"Lots per side": {T: "每边手数" },
	"Lots per level": {T: "每级手数" },
	"Price levels per side": {T: "每边价格级别" },
	"Configure manually": {T: "手动配置" },
	"Loading market data": {T: "加载市场数据" },
	"on": {T: "开启" },
	"err_with_cex_creds": {T: "连接这些凭证时发生错误" },
	"approve_token_wallet_addr": {T: `<span data-tmpl="parentName" class="me-1"></span> 地址：` },
	"Available fee balance": {T: "可用费用余额" },
	"add_provider_tooltip": {T: "您正在使用默认的公共RPC提供商集。此提供商集可能会过时或不可靠。请改为指定您自己的可信提供商。" },
	"add_custom_rpc_provider": {T: "添加自定义RPC提供商" },
	"Profit": {T: "利润" },
	"Inventory": {T: "库存" },
	"Booked orders": {T: "预定订单" },
	"Fee reserves": {T: "费用储备" },
	"Pending deposits": {T: "待处理存款" },
	"Pending withdrawals": {T: "待处理取款" },
	"Settled matches": {T: "已结算匹配" },
	"Traded": {T: "已交易" },
	"Basis price": {T: "基础价格" },
	"Fee gap": {T: "费用差距" },
	"Remote gap": {T: "远程差距" },
	"Round_trip fees": {T: "往返费用" },
	"feegap_tooltip": {T: "调整费率以考虑链上交易费用" },
	"remotegap_tooltip": {T: "链接的CEX市场上的买卖差价" },
	"max_zero_no_fees": {T: `<span id="maxZeroNoFeesTicker"></span> 余额 < 最小费用 ~<span id="maxZeroMinFees"></span>` },
	"max_zero_no_bal": {T: `低 <span id="maxZeroNoBalTicker"></span> 余额` },
	"Wallet Options": {T: "钱包选项" },
	"balance_diff": {T: "余额差异" },
	"usd_diff": {T: "美元差异" },
	"usd_rate": {T: "美元汇率" },
	"dex_sell": {T: "DEX 卖出" },
	"dex_buy": {T: "DEX 买入" },
	"cex_sell": {T: "CEX 卖出" },
	"cex_buy": {T: "CEX 买入" },
	"deposit": {T: "存款" },
	"withdrawal": {T: "取款" },
	"filters": {T: "筛选器" },
	"Apply": {T: "应用" },
	"Block Sync": {T: "区块同步" },
	"Balance Discovery": {T: "余额发现" },
	"Finding Addresses": {T: "查找地址" },
	"Hide Mixing Transactions": {T: "隐藏混合交易" },
	"Redeem game code": {T: "兑换游戏代码" },
	"Redeem Game Code": {T: "兑换游戏代码" },
	"Code": {T: "代码" },
	"Message_optional": {T: "消息（可选）" },
	"Game code redeemed": {T: "游戏代码已兑换！" },
	"Transaction": {T: "交易" },
	"Value": {T: "价值" },
	"Prepaid bond redeemed": {T: "预付债券已兑换！" },
	"Enable Account": {T: "启用账户" },
	"Show trading tier info": {T: "显示交易等级信息" },
	"Hide trading tier info": {T: "隐藏交易等级信息" },
	"Show reputation": {T: "显示声誉" },
	"Hide reputation": {T: "隐藏声誉" },
	"buy_orders_success": {T: "所有买单成功提交" },
	"sell_orders_success": {T: "所有卖单成功提交" },
	"buy_orders_failed": {T: "无法提交所有买单" },
	"sell_orders_failed": {T: "无法提交所有卖单" },
	"Order report": {T: "订单报告" },
	"Remaining": {T: "剩余" },
	"Used": {T: "已使用" },
	"Deficiency": {T: "不足" },
	"Deficiency with Pending": {T: "不足（待处理）" },
	"Standing Lots": {T: "未成交手数" },
	"Ordered Lots": {T: "已下单手数" },
	"Arb Rate": {T: "套利率" },
	"Required DEX": {T: "所需DEX" },
	"Required CEX": {T: "所需CEX" },
	"Used DEX": {T: "已用DEX" },
	"Used CEX": {T: "已用CEX" },
	"Causes Self Match": {T: "导致自成交" },
	"Priority": {T: "优先级" },
	"Wallet Balances": {T: "钱包余额" },
	"Placements": {T: "下单" },
	"delete_bot": {T: "删除机器人" },
}
