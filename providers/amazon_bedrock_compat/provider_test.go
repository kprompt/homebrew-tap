package amazon_bedrock_compat

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ourorg/primaryrepo/enterprise/deferred"
	"github.com/ourorg/primaryrepo/logging/logtest"
)

func TestProvider_GenerateText_Success(t *testing.T) {
	// Setup a test HTTP server that simulates Amazon Bedrock compat API.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/generate" {
			t.Errorf("expected /generate path, got %s", r.URL.Path)
		}
		auth := r.Header.Get("Authorization")
		if auth != "ApiKey testkey" {
			t.Errorf("expected Authorization header 'ApiKey testkey', got %s", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"completion":"Hello from Bedrock!"}`))
	}))
	defer ts.Close()

	logger := logtest.NewTestLogger(t)
	provider := NewProvider(ts.URL, "testkey", logger)

	text, err := provider.GenerateText(context.Background(), "Hello", deferred.GenerationOptions{
		MaxTokens:   5,
		Temperature: 0.7,
		TopP:        1.0,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if text != "Hello from Bedrock!" {
		t.Errorf("unexpected completion text: %q", text)
	}
}

func TestProvider_GenerateText_HTTPError(t *testing.T) {
	// Simulate server returning 500 error.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	logger := logtest.NewTestLogger(t)
	provider := NewProvider(ts.URL, "testkey", logger)

	_, err := provider.GenerateText(context.Background(), "Hello", deferred.GenerationOptions{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestProvider_GenerateText_BadResponse(t *testing.T) {
	// Simulate server returning invalid JSON.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{invalid json}`))
	}))
	defer ts.Close()

	logger := logtest.NewTestLogger(t)
	provider := NewProvider(ts.URL, "testkey", logger)

	_, err := provider.GenerateText(context.Background(), "Hello", deferred.GenerationOptions{})
	if err == nil {
		t.Fatal("expected error due to invalid JSON, got nil")
	}
}
