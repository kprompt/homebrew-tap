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
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/v1/completions" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		if r.Header.Get("x-api-key") != "test-key" {
			t.Errorf("missing or wrong API key header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"id": "test-id",
			"object": "text_completion",
			"created": 1234567890,
			"choices": [
				{
					"text": "Hello from Amazon Bedrock!",
					"index": 0,
					"finish_reason": "stop"
				}
			]
		}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	provider := NewAmazonBedrockProvider(server.URL, "test-key")
	provider.HTTPClient.Timeout = 2 * time.Second

	ctx := context.Background()
	text, err := provider.Complete(ctx, "bedrock-model", "Say hello", 10, 0.5, []string{"\n"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(text, "Hello from Amazon Bedrock!") {
		t.Errorf("unexpected completion text: %s", text)
	}
}

func TestAmazonBedrockProvider_Complete_BadConfig(t *testing.T) {
	provider := NewAmazonBedrockProvider("", "")
	ctx := context.Background()
	_, err := provider.Complete(ctx, "model", "prompt", 10, 0.5, nil)
	if err == nil {
		t.Fatal("expected error due to bad config, got nil")
	}
}

func TestAmazonBedrockProvider_Complete_HTTPError(t *testing.T) {
	// Setup a test HTTP server that returns 500
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	provider := NewAmazonBedrockProvider(server.URL, "key")
	ctx := context.Background()
	_, err := provider.Complete(ctx, "model", "prompt", 10, 0.5, nil)
	if err == nil || !strings.Contains(err.Error(), "unexpected status code") {
		t.Fatalf("expected status code error, got %v", err)
	}
}

func TestAmazonBedrockProvider_Complete_InvalidJSON(t *testing.T) {
	// Setup a test HTTP server that returns invalid JSON
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{invalid json}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	provider := NewAmazonBedrockProvider(server.URL, "key")
	ctx := context.Background()
	_, err := provider.Complete(ctx, "model", "prompt", 10, 0.5, nil)
	if err == nil || !strings.Contains(err.Error(), "failed to decode response") {
		t.Fatalf("expected decode error, got %v", err)
	}
}

func TestAmazonBedrockProvider_Complete_NoChoices(t *testing.T) {
	// Setup a test HTTP server that returns empty choices array
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"id": "id",
			"object": "text_completion",
			"created": 123,
			"choices": []
		}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	provider := NewAmazonBedrockProvider(server.URL, "key")
	ctx := context.Background()
	_, err := provider.Complete(ctx, "model", "prompt", 10, 0.5, nil)
	if err == nil || !strings.Contains(err.Error(), "no completion choices") {
		t.Fatalf("expected no choices error, got %v", err)
	}
}
