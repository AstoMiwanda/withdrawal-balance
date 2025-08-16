package xendit

import "time"

type CreatePayoutRequest struct {
	BankAccountNumber   string  `json:"bank_account_number"`
	BankAccountName     string  `json:"bank_account_name"`
	Amount              float64 `json:"amount"`
	ReceiptNotification string  `json:"receipt_notification"`
}

type CreatePayoutResponse struct {
	Id          string `json:"id"`
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
}

type CancelPayoutRequest struct {
	Id string `json:"id"`
}

type CancelPayoutResponse struct {
	Id          string `json:"id"`
	ReferenceId string `json:"reference_id"`
	Status      string `json:"status"`
}

type (
	PayoutPayload struct {
		ReferenceId         string              `json:"reference_id"`
		ChannelCode         string              `json:"channel_code"`
		ChannelProperties   ChannelProperties   `json:"channel_properties"`
		Amount              float64             `json:"amount"`
		Description         string              `json:"description"`
		Currency            string              `json:"currency"`
		ReceiptNotification ReceiptNotification `json:"receipt_notification"`
		Metadata            Metadata            `json:"metadata"`
	}

	ChannelProperties struct {
		AccountNumber     string `json:"account_number"`
		AccountHolderName string `json:"account_holder_name"`
	}

	ReceiptNotification struct {
		EmailTo  []string `json:"email_to"`
		EmailCc  []string `json:"email_cc"`
		EmailBcc []string `json:"email_bcc"`
	}

	Metadata struct {
		OutletNo int `json:"outlet_no"`
	}

	PayoutResponse struct {
		Id                   string    `json:"id"`
		Amount               float64   `json:"amount"`
		ChannelCode          string    `json:"channel_code"`
		Currency             string    `json:"currency"`
		Description          string    `json:"description"`
		ReferenceId          string    `json:"reference_id"`
		Status               string    `json:"status"`
		Created              time.Time `json:"created"`
		Updated              time.Time `json:"updated"`
		EstimatedArrivalTime time.Time `json:"estimated_arrival_time"`
		BusinessId           string    `json:"business_id"`
		ChannelProperties    struct {
			AccountNumber     string `json:"account_number"`
			AccountHolderName string `json:"account_holder_name"`
		} `json:"channel_properties"`
		ReceiptNotification struct {
			EmailTo  []string `json:"email_to"`
			EmailCc  []string `json:"email_cc"`
			EmailBcc []string `json:"email_bcc"`
		} `json:"receipt_notification"`
		Metadata struct {
			OutletNo int `json:"outlet_no"`
		} `json:"metadata"`
	}
)

type (
	RestCancelPayoutPayload struct {
		Id          string `json:"id"`
		ReferenceId string `json:"reference_id"`
	}

	RestCancelPayoutResponse struct {
		Id                   string    `json:"id"`
		Amount               int       `json:"amount"`
		ChannelCode          string    `json:"channel_code"`
		Currency             string    `json:"currency"`
		Description          string    `json:"description"`
		ReferenceId          string    `json:"reference_id"`
		Status               string    `json:"status"`
		Created              time.Time `json:"created"`
		Updated              time.Time `json:"updated"`
		EstimatedArrivalTime time.Time `json:"estimated_arrival_time"`
		BusinessId           string    `json:"business_id"`
		ChannelProperties    struct {
			PayoutCode          string    `json:"payout_code"`
			RecipientGivenNames string    `json:"recipient_given_names"`
			RecipientSurname    string    `json:"recipient_surname"`
			ExpiresAt           time.Time `json:"expires_at"`
		} `json:"channel_properties"`
	}
)
