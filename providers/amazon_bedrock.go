package providers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// AmazonBedrockProvider implements the Enterprise Deferred provider interface for Amazon Bedrock.
type AmazonBedrockProvider struct {
	Endpoint   string
	APIKey     string
	HTTPClient *http.Client
}

// NewAmazonBedrockProvider creates a new AmazonBedrockProvider with default HTTP client.
func NewAmazonBedrockProvider(endpoint, apiKey string) *AmazonBedrockProvider {
	return &AmazonBedrockProvider{
		Endpoint:   endpoint,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// CompletionRequest represents the request payload for Amazon Bedrock completion.
type CompletionRequest struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt"`
	MaxTokens   int      `json:"max_tokens,omitempty"`
	Temperature float64  `json:"temperature,omitempty"`
	Stop        []string `json:"stop,omitempty"`
}

// CompletionResponse represents the response from Amazon Bedrock completion.
type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

// Complete sends a completion request to Amazon Bedrock and returns the generated text.
func (p *AmazonBedrockProvider) Complete(ctx context.Context, model, prompt string, maxTokens int, temperature float64, stop []string) (string, error) {
	if p.Endpoint == "" || p.APIKey == "" {
		return "", errors.New("amazon bedrock provider not properly configured")
	}

	reqBody := CompletionRequest{
		Model:       model,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: temperature,
		Stop:        stop,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.Endpoint+"/v1/completions", strings.NewReader(string(reqBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.APIKey)

	resp, err := p.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var completionResp CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&completionResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(completionResp.Choices) == 0 {
		return "", errors.New("no completion choices returned")
	}

	return completionResp.Choices[0].Text, nil
}
