package wallet

import "withdrawal-balance/internal/constant"

type UpdateBalanceRequest struct {
	WalletID int64
	UserID   int64
	Amount   float64
	Type     constant.TransactionTypeWallet
}
