package constant

type TransactionTypeWallet string

const (
	TopUpTransactionTypeWallet    TransactionTypeWallet = "TOPUP"
	WithdrawTransactionTypeWallet TransactionTypeWallet = "WITHDRAW"
)

const (
	XenditChannelCode    = "PH_GCASH"
	XenditCurrency       = "IDR"
	XenditOutletNo       = 24
	XenditWithdrawalDesc = "WITHDRAWAL_BALANCE"
)
