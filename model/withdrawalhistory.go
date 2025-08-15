package model

import "time"

type WithdrawalHistory struct {
	ID                   int64      `json:"id" db:"id"`
	UserID               int64      `json:"user_id" db:"user_id"`
	WalletID             int64      `json:"wallet_id" db:"wallet_id"`
	Amount               float64    `json:"amount" db:"amount"`
	BankAccountNumber    string     `json:"bank_account_number" db:"bank_account_number"`
	BankAccountName      string     `json:"bank_account_name" db:"bank_account_name"`
	BankCode             string     `json:"bank_code" db:"bank_code"`
	Status               string     `json:"status" db:"status"`
	TransactionReference string     `json:"transaction_reference" db:"transaction_reference"`
	PayoutId             string     `json:"payout_id" db:"payout_id"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at" db:"updated_at"`
}
