package amazonbedrock

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateText(t *testing.T) {
	// Setup a test server to mock Amazon Bedrock API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/generate-text" {
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
		_, _ = w.Write([]byte(`{"text":"Hello from Bedrock"}`))
	}))
	defer ts.Close()

	client := NewClient(ts.URL, "test-api-key")

	resp, err := client.GenerateText(context.Background(), GenerateTextRequest{
		Model: "bedrock-test-model",
		Input: "Say hello",
	})
	if err != nil {
		t.Fatalf("GenerateText failed: %v", err)
	}

	if resp.Text != "Hello from Bedrock" {
		t.Errorf("unexpected response text: %s", resp.Text)
	}
}
