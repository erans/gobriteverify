package gobriteverify

import (
	"os"
	"strings"
	"testing"
)

func getAPIKeyFromEnv() string {
	env := os.Environ()
	for _, v := range env {
		if v != "" {
			parts := strings.Split(v, "=")
			if len(parts) > 1 {
				if parts[0] == "TEST_BRITEVERIFY_API_KEY" {
					return parts[1]
				}
			}
		}
	}

	return ""
}

func getClient(t *testing.T) *Client {
	apiKey := getAPIKeyFromEnv()
	if apiKey == "" {
		t.Error("Missing api key from environment. Make sure to set the 'TEST_BRITEVERIFY_API_KEY' environment variable to the BriteVerify API Key")
	}

	client := NewClient(apiKey)
	if client == nil {
		t.Error("apiKey is missing or invalid")
	}

	return client
}

func runTest(t *testing.T, email string, expectedStatus string) {
	client := getClient(t)
	var result *BriteVerifyEmailsResponse
	var err error

	// Test invalid email account
	if result, err = client.Verify(email); err == nil {
		if result.Status != expectedStatus {
			t.Errorf("Expected status=%s", expectedStatus)
		}
	} else {
		t.Error(err)
	}
}

func TestInvalidEmailAddress(t *testing.T) {
	runTest(t, "aa@gmail.com", "invalid")
}

func TestValidEmailAddress(t *testing.T) {
	runTest(t, "support@briteverify.com", "valid")
}
