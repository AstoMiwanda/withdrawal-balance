package withdrawalhistory

type CreateWithdrawalBalanceRequest struct {
	UserID            int64   `json:"user_id" validate:"required"`
	WalletID          int64   `json:"wallet_id" validate:"required"`
	Amount            float64 `json:"amount" validate:"required"`
	BankCode          string  `json:"bank_code" validate:"required"`
	BankAccountNumber string  `json:"bank_account_number" validate:"required"`
	BankAccountName   string  `json:"bank_account_name" validate:"required"`
}

type CreateWithdrawalBalanceResponse struct {
	ID              string  `json:"id"`
	UserID          int64   `json:"user_id"`
	WalletID        int64   `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	ReferenceNumber string  `json:"reference_number"`
}
