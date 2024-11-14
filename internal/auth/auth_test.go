package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: Proper Authorization header
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedAPIKey := "test-api-key"
	if apiKey != expectedAPIKey {
		t.Errorf("expected %v, got %v", expectedAPIKey, apiKey)
	}

	// Test case 2: Missing Authorization header
	headers = http.Header{}          // Reset headers with no Authorization
	apiKey, err = GetAPIKey(headers) // Call function with missing header
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	// Test case 3: Malformed Authorization header
	headers = http.Header{}
	headers.Set("Authorization", "Bearer test-api-key") // Incorrect prefix

	apiKey, err = GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Fatalf("expected malformed authorization header error, got %v", err)
	}
}
