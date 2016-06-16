package gobriteverify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BriteVerifyEmailsResponse provides a strcuture for the API call response
type BriteVerifyEmailsResponse struct {
	Address     string  `json:"address"`
	Account     string  `json:"account"`
	Domain      string  `json:"domain"`
	Status      string  `json:"status"`
	ErrorCode   string  `json:"error_code"`
	Error       string  `json:"error"`
	Disposable  bool    `json:"disposable"`
	RoleAddress bool    `json:"role_address"`
	Duration    float64 `json:"duration"`
}

// Client is the BriteVerify client struct
type Client struct {
	apiKey string
}

// NewClient creates a new BriteVerify Client
func NewClient(apiKey string) *Client {
	if apiKey == "" {
		return nil
	}
	return &Client{apiKey: apiKey}
}

// Verify calls the BriteVerfiy email verification API
func (client *Client) Verify(email string) (*BriteVerifyEmailsResponse, error) {
	url := fmt.Sprintf("https://bpi.briteverify.com/emails.json?address=%s&apikey=%s", email, client.apiKey)

	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err == nil {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Response failed with status code %d", resp.StatusCode)
		}

		decoder := json.NewDecoder(resp.Body)
		result := &BriteVerifyEmailsResponse{}
		if err = decoder.Decode(result); err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, err
}
