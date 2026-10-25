package providers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAmazonBedrockClient_CallBedrock_Success(t *testing.T) {
	// Setup a test HTTP server to mock Amazon Bedrock API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.Header.Get("x-api-key") != "test-api-key" {
			t.Errorf("expected x-api-key header 'test-api-key', got %s", r.Header.Get("x-api-key"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"output":"Hello from Bedrock"}`))
	}))
	defer ts.Close()

	client := NewAmazonBedrockClient(ts.URL, "test-api-key")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	output, err := client.CallBedrock(ctx, BedrockRequest{
		Model: "bedrock-model",
		Input: "test input",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != "Hello from Bedrock" {
		t.Errorf("expected output 'Hello from Bedrock', got %s", output)
	}
}

func TestAmazonBedrockClient_CallBedrock_HTTPError(t *testing.T) {
	client := NewAmazonBedrockClient("http://invalid-host", "key")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := client.CallBedrock(ctx, BedrockRequest{
		Model: "model",
		Input: "input",
	})
	if err == nil {
		t.Fatal("expected error but got none")
	}
}

func TestAmazonBedrockClient_CallBedrock_Non200Status(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	client := NewAmazonBedrockClient(ts.URL, "key")
	ctx := context.Background()

	_, err := client.CallBedrock(ctx, BedrockRequest{
		Model: "model",
		Input: "input",
	})
	if err == nil || err.Error() != "unexpected status code 400 from Bedrock" {
		t.Fatalf("expected status code error, got %v", err)
	}
}

func TestAmazonBedrockClient_CallBedrock_EmptyOutput(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"output":""}`))
	}))
	defer ts.Close()

	client := NewAmazonBedrockClient(ts.URL, "key")
	ctx := context.Background()

	_, err := client.CallBedrock(ctx, BedrockRequest{
		Model: "model",
		Input: "input",
	})
	if err == nil || err.Error() != "empty output from Bedrock" {
		t.Fatalf("expected empty output error, got %v", err)
	}
}
