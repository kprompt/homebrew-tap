package amazonbedrockcompat

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yourorg/yourrepo/types"
)

func TestGenerateText_Success(t *testing.T) {
	// Setup a test HTTP server to mock Amazon Bedrock compat API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/generate" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("unexpected method: %s", r.Method)
		}
		auth := r.Header.Get("Authorization")
		if auth != "Bearer testkey" {
			t.Errorf("unexpected auth header: %s", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"text":"Hello from Bedrock"}`))
	}))
	defer ts.Close()

	p := NewProvider("testkey", ts.URL)
	resp, err := p.GenerateText(context.Background(), &types.TextGenerationRequest{
		Prompt:      "Hello",
		MaxTokens:   10,
		Temperature: 0.7,
		Model:       "bedrock-model",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Text != "Hello from Bedrock" {
		t.Errorf("unexpected response text: %s", resp.Text)
	}
}

func TestGenerateText_EmptyPrompt(t *testing.T) {
	p := NewProvider("testkey", "http://localhost")
	_, err := p.GenerateText(context.Background(), &types.TextGenerationRequest{
		Prompt: "",
	})
	if err == nil {
		t.Fatal("expected error for empty prompt")
	}
}

func TestGenerateText_NilRequest(t *testing.T) {
	p := NewProvider("testkey", "http://localhost")
	_, err := p.GenerateText(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil request")
	}
}
