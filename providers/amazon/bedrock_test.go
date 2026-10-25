package amazon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGenerateText_Success(t *testing.T) {
	// Setup a test server to mock Bedrock API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/generate" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-api-key" {
			t.Errorf("unexpected Authorization header: %s", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "resp-123",
			"object": "text_completion",
			"created": 1234567890,
			"choices": [{"text": "Hello from Bedrock!", "finish_reason": "stop", "index": 0}]
		}`))
	}))
	defer ts.Close()

	client := NewBedrockClient(ts.URL, "test-api-key")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	output, err := client.GenerateText(ctx, "test-model", "Hello", nil)
	if err != nil {
		t.Fatalf("GenerateText failed: %v", err)
	}
	if !strings.Contains(output, "Bedrock") {
		t.Errorf("unexpected output: %s", output)
	}
}

func TestGenerateText_EmptyModel(t *testing.T) {
	client := NewBedrockClient("http://example.com", "")
	_, err := client.GenerateText(context.Background(), "", "input", nil)
	if err == nil || !strings.Contains(err.Error(), "model name is required") {
		t.Errorf("expected model name error, got %v", err)
	}
}

func TestGenerateText_EmptyInput(t *testing.T) {
	client := NewBedrockClient("http://example.com", "")
	_, err := client.GenerateText(context.Background(), "model", "", nil)
	if err == nil || !strings.Contains(err.Error(), "input text is required") {
		t.Errorf("expected input text error, got %v", err)
	}
}

func TestGenerateText_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer ts.Close()

	client := NewBedrockClient(ts.URL, "")
	_, err := client.GenerateText(context.Background(), "model", "input", nil)
	if err == nil || !strings.Contains(err.Error(), "unexpected status code") {
		t.Errorf("expected status code error, got %v", err)
	}
}

func TestGenerateText_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer ts.Close()

	client := NewBedrockClient(ts.URL, "")
	_, err := client.GenerateText(context.Background(), "model", "input", nil)
	if err == nil || !strings.Contains(err.Error(), "failed to decode response") {
		t.Errorf("expected decode error, got %v", err)
	}
}

func TestGenerateText_NoChoices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"choices":[]}`))
	}))
	defer ts.Close()

	client := NewBedrockClient(ts.URL, "")
	_, err := client.GenerateText(context.Background(), "model", "input", nil)
	if err == nil || !strings.Contains(err.Error(), "no choices returned") {
		t.Errorf("expected no choices error, got %v", err)
	}
}
