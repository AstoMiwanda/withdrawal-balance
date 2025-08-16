package wallet

import "withdrawal-balance/internal/constant"

type UpdateBalanceRequest struct {
	WalletID int64
	UserID   int64
	Amount   float64
	Type     constant.TransactionTypeWallet
}

type InquiryBalanceRequest struct {
	UserId   int64 `json:"user_id"`
	WalletId int64 `json:"wallet_id"`
}

type InquiryBalanceResponse struct {
	UserId   int64   `json:"user_id"`
	WalletId int64   `json:"wallet_id"`
	Balance  float64 `json:"balance"`
}
