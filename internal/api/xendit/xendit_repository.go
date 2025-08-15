package xendit

import (
	"bytes"
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

func (r *Repository) Payout(ctx context.Context, request PayoutPayload) (*PayoutResponse, error) {
	apiKey := os.Getenv("API_KEY_XENDIT")
	baseURL := os.Getenv("URL_XENDIT")
	pathURL := fmt.Sprintf("%s/v2/payouts", baseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pathURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", pkg.SetBasicAuth(apiKey, ""))
	req.Header.Set("Idempotency-key", request.ReferenceId)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("xendit API error: %s", string(body))
	}

	var response PayoutResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
