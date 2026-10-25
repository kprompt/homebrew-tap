package providers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAmazonBedrockProvider_Complete_Success(t *testing.T) {
	// Setup a test HTTP server to mock Amazon Bedrock API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/completions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			t.Errorf("missing or invalid Authorization header: %s", auth)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"id": "cmpl-123",
			"object": "text_completion",
			"created": 1672531200,
			"choices": [
				{
					"text": "Hello from Amazon Bedrock!",
					"finish_reason": "stop",
					"index": 0
				}
			]
		}`))
	}))
	defer ts.Close()

	provider := NewAmazonBedrockProvider(ts.URL, "test-api-key")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := provider.Complete(ctx, "test-model", "Say hello", 10, 0.5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "Hello from Amazon Bedrock!"
	if resp != expected {
		t.Errorf("expected response %q, got %q", expected, resp)
	}
}

func TestAmazonBedrockProvider_Complete_NoChoices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"id": "cmpl-123",
			"object": "text_completion",
			"created": 1672531200,
			"choices": []
		}`))
	}))
	defer ts.Close()

	provider := NewAmazonBedrockProvider(ts.URL, "test-api-key")
	ctx := context.Background()

	_, err := provider.Complete(ctx, "test-model", "Say hello", 10, 0.5)
	if err == nil || !strings.Contains(err.Error(), "no completion choices") {
		t.Errorf("expected error about no completion choices, got %v", err)
	}
}

func TestAmazonBedrockProvider_Complete_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	provider := NewAmazonBedrockProvider(ts.URL, "test-api-key")
	ctx := context.Background()

	_, err := provider.Complete(ctx, "test-model", "Say hello", 10, 0.5)
	if err == nil || !strings.Contains(err.Error(), "unexpected status code") {
		t.Errorf("expected error about unexpected status code, got %v", err)
	}
}
