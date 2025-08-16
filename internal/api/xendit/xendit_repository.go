package xendit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"withdrawal-balance/pkg"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Payout(ctx context.Context, request PayoutPayload) (result PayoutResponse, err error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	baseURL := os.Getenv("XENDIT_URL")
	url := fmt.Sprintf("%s/v2/payouts", baseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return result, err
	}

	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	Headers["Authorization"] = pkg.SetBasicAuth(apiKey, "")
	Headers["Idempotency-key"] = request.ReferenceId

	payload := pkg.Request{
		Method:      pkg.POST,
		URL:         url,
		Headers:     Headers,
		QueryParams: map[string]string{},
		Body:        jsonData,
	}

	resp, err := pkg.SendRequestRaw(payload, "", true, "InquiryLoanInformation")
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("xendit API error: %s", string(body))
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *Repository) CancelPayout(ctx context.Context, request RestCancelPayoutPayload) (result RestCancelPayoutResponse, err error) {
	apiKey := os.Getenv("XENDIT_API_KEY")
	baseURL := os.Getenv("XENDIT_URL")
	url := fmt.Sprintf("%s/v2/payouts/%s/cancel", baseURL, request.Id)

	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	Headers["Authorization"] = pkg.SetBasicAuth(apiKey, "")
	Headers["Idempotency-key"] = request.ReferenceId

	payload := pkg.Request{
		Method:      pkg.POST,
		URL:         url,
		Headers:     Headers,
		QueryParams: map[string]string{},
		Body:        nil,
	}

	resp, err := pkg.SendRequestRaw(payload, "", true, "InquiryLoanInformation")
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("xendit API error: %s", string(body))
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}
