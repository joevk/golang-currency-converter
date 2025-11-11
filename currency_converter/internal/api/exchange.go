package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

// RatesResponse represents the JSON structure returned by Open Exchange Rates latest endpoint.
// Only include the fields we need.
type RatesResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

// Client wraps HTTP operations for the exchange rates API.
// Provides a customizable HTTP client (e.g., for timeouts or testing with mock server).
//
// API Docs: https://docs.openexchangerates.org/
// Free plan endpoint used: https://openexchangerates.org/api/latest.json
// Base is always USD on the free plan.

type Client struct {
	HTTP *http.Client
	APIKey string
}

const baseURL = "https://openexchangerates.org/api/latest.json"

// NewClient creates a new API client with sensible defaults.
func NewClient() *Client {
	return &Client{
		HTTP: &http.Client{Timeout: 10 * time.Second},
		APIKey: os.Getenv("OPENEXCHANGERATES_API_KEY"),
	}
}

// FetchRates retrieves the latest exchange rates. Returns an error if the API key is missing
// or any network/JSON issues occur.
func (c *Client) FetchRates() (*RatesResponse, error) {
	if c.APIKey == "" {
		return nil, errors.New("missing API key: set OPENEXCHANGERATES_API_KEY")
	}

	url := fmt.Sprintf("%s?app_id=%s", baseURL, c.APIKey)
	resp, err := c.HTTP.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var data RatesResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if data.Rates == nil || len(data.Rates) == 0 {
		return nil, errors.New("no rates returned")
	}

	return &data, nil
}
