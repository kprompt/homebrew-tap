package amazon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// BedrockClient provides compatibility with Amazon Bedrock API.
type BedrockClient struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

// NewBedrockClient creates a new BedrockClient with the given endpoint and API key.
func NewBedrockClient(endpoint, apiKey string) *BedrockClient {
	return &BedrockClient{
		Endpoint:   strings.TrimRight(endpoint, "/"),
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// BedrockRequest represents a request to the Bedrock API.
type BedrockRequest struct {
	Model   string                 `json:"model"`
	Input   string                 `json:"input"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// BedrockResponse represents a response from the Bedrock API.
type BedrockResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// GenerateText sends a text generation request to Amazon Bedrock and returns the generated text.
func (c *BedrockClient) GenerateText(ctx context.Context, model, input string, options map[string]interface{}) (string, error) {
	if model == "" {
		return "", errors.New("model name is required")
	}
	if input == "" {
		return "", errors.New("input text is required")
	}

	reqBody := BedrockRequest{
		Model:   model,
		Input:   input,
		Options: options,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Endpoint+"/generate", strings.NewReader(string(reqBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var bedrockResp BedrockResponse
	if err := json.NewDecoder(resp.Body).Decode(&bedrockResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(bedrockResp.Choices) == 0 {
		return "", errors.New("no choices returned from Bedrock")
	}

	return bedrockResp.Choices[0].Text, nil
}
