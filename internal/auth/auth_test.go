package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("no auth header included", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})

	t.Run("malformed auth header", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "Bearer")
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("correct auth header", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey 123")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if apiKey != "123" {
			t.Errorf("expected api key 123, got %v", apiKey)
		}
	})
}
