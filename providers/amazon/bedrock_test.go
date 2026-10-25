package amazon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBedrockClient_GenerateText_Success(t *testing.T) {
	// Mock server to simulate Bedrock API
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/generate" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("missing or wrong Authorization header")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"text":"Hello from Bedrock"}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewBedrockClient(server.URL, "test-api-key")
	ctx := context.Background()
	text, err := client.GenerateText(ctx, "test-model", "Hello", 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if text != "Hello from Bedrock" {
		t.Errorf("unexpected text: %s", text)
	}
}

func TestBedrockClient_GenerateText_Errors(t *testing.T) {
	client := NewBedrockClient("http://invalid", "key")
	ctx := context.Background()

	_, err := client.GenerateText(ctx, "", "prompt", 10)
	if err == nil {
		t.Error("expected error for empty model")
	}

	_, err = client.GenerateText(ctx, "model", "", 10)
	if err == nil {
		t.Error("expected error for empty prompt")
	}
}
