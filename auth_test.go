package auth

import (
	"os"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Set up environment variable for testing, if GetAPIKey relies on one
	os.Setenv("API_KEY", "test-api-key")
	defer os.Unsetenv("API_KEY") // clean up after the test

	// Call the function
	apiKey := GetAPIKey()

	// Define the expected result
	expectedAPIKey := "test-api-key"

	// Check if the result is what we expect
	if apiKey != expectedAPIKey {
		t.Errorf("GetAPIKey() = %v; want %v", apiKey, expectedAPIKey)
	}
}
