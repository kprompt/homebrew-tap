package amazon

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/yourorg/yourrepo/providers"
)

// BedrockClient is a client for Amazon Bedrock service.
type BedrockClient struct {
	httpClient *http.Client
	endpoint   string
	apiKey     string
}

// NewBedrockClient creates a new BedrockClient with the given endpoint and API key.
func NewBedrockClient(endpoint, apiKey string) *BedrockClient {
	return &BedrockClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		endpoint:   endpoint,
		apiKey:     apiKey,
	}
}

// GenerateText sends a prompt to Amazon Bedrock and returns the generated text.
func (c *BedrockClient) GenerateText(ctx context.Context, model string, prompt string, maxTokens int) (string, error) {
	if model == "" {
		return "", errors.New("model name is required")
	}
	if prompt == "" {
		return "", errors.New("prompt is required")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint+"/generate", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	// Add headers and query parameters as needed for Bedrock API compatibility
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	q := req.URL.Query()
	q.Add("model", model)
	q.Add("prompt", prompt)
	q.Add("max_tokens", fmt.Sprintf("%d", maxTokens))
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad response status: %s", resp.Status)
	}

	var result struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Text, nil
}

// BedrockProvider implements the providers.Provider interface for Amazon Bedrock.
type BedrockProvider struct {
	client *BedrockClient
}

// NewBedrockProvider creates a new BedrockProvider.
func NewBedrockProvider(endpoint, apiKey string) *BedrockProvider {
	return &BedrockProvider{
		client: NewBedrockClient(endpoint, apiKey),
	}
}

// Generate implements providers.Provider interface.
func (p *BedrockProvider) Generate(ctx context.Context, model string, prompt string, maxTokens int) (string, error) {
	return p.client.GenerateText(ctx, model, prompt, maxTokens)
}

// Ensure BedrockProvider implements providers.Provider interface.
var _ providers.Provider = (*BedrockProvider)(nil)
