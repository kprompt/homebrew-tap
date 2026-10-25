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

// AmazonBedrockClient is a client for Amazon Bedrock API.
type AmazonBedrockClient struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

// NewAmazonBedrockClient creates a new AmazonBedrockClient with default HTTP client.
func NewAmazonBedrockClient(endpoint, apiKey string) *AmazonBedrockClient {
	return &AmazonBedrockClient{
		Endpoint:   endpoint,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// BedrockRequest represents a request to Amazon Bedrock.
type BedrockRequest struct {
	Model   string                 `json:"model"`
	Input   string                 `json:"input"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// BedrockResponse represents a response from Amazon Bedrock.
type BedrockResponse struct {
	Output string `json:"output"`
}

// CallBedrock sends a request to Amazon Bedrock and returns the output or error.
func (c *AmazonBedrockClient) CallBedrock(ctx context.Context, req BedrockRequest) (string, error) {
	url := strings.TrimRight(c.Endpoint, "/") + "/invoke"
	payload, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(payload)))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d from Bedrock", resp.StatusCode)
	}

	var bedrockResp BedrockResponse
	if err := json.NewDecoder(resp.Body).Decode(&bedrockResp); err != nil {
		return "", fmt.Errorf("failed to decode Bedrock response: %w", err)
	}

	if bedrockResp.Output == "" {
		return "", errors.New("empty output from Bedrock")
	}

	return bedrockResp.Output, nil
}
