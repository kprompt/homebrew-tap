package amazonbedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// BedrockClient is a client for Amazon Bedrock API.
type BedrockClient struct {
	Endpoint   string
	HTTPClient *http.Client
	APIKey     string
}

// NewClient creates a new BedrockClient.
func NewClient(endpoint, apiKey string) *BedrockClient {
	return &BedrockClient{
		Endpoint:   strings.TrimRight(endpoint, "/"),
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
	}
}

// GenerateTextRequest represents a request to generate text.
type GenerateTextRequest struct {
	Model  string                 `json:"model"`
	Input  string                 `json:"input"`
	Params map[string]interface{} `json:"parameters,omitempty"`
}

// GenerateTextResponse represents a response from generate text.
type GenerateTextResponse struct {
	Text string `json:"text"`
}

// GenerateText calls the Amazon Bedrock text generation API.
func (c *BedrockClient) GenerateText(ctx context.Context, req GenerateTextRequest) (GenerateTextResponse, error) {
	url := fmt.Sprintf("%s/generate-text", c.Endpoint)

	payload, err := json.Marshal(req)
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(payload)))
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return GenerateTextResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GenerateTextResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var genResp GenerateTextResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		return GenerateTextResponse{}, fmt.Errorf("decode response: %w", err)
	}

	return genResp, nil
}
