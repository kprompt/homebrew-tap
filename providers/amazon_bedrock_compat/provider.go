package amazonbedrockcompat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/yourorg/yourrepo/providers"
	"github.com/yourorg/yourrepo/types"
)

// Provider implements the Enterprise Deferred Amazon Bedrock compatibility path.
type Provider struct {
	client *http.Client
	apiKey string
	endpoint string
}

// NewProvider creates a new Amazon Bedrock compat provider with given config.
func NewProvider(apiKey, endpoint string) *Provider {
	return &Provider{
		client: &http.Client{Timeout: 10 * time.Second},
		apiKey: apiKey,
		endpoint: endpoint,
	}
}

// GenerateText calls the Amazon Bedrock compat API and returns generated text.
func (p *Provider) GenerateText(ctx context.Context, req *types.TextGenerationRequest) (*types.TextGenerationResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}
	if req.Prompt == "" {
		return nil, errors.New("prompt cannot be empty")
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.endpoint+"/generate", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	// Compose request body according to Amazon Bedrock compat API spec
	body := map[string]interface{}{
		"prompt": req.Prompt,
		"max_tokens": req.MaxTokens,
		"temperature": req.Temperature,
		"model": req.Model,
	}
	// Encode body to JSON
	err = providers.EncodeJSONBody(httpReq, body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp struct {
		Text string `json:"text"`
	}
	err = providers.DecodeJSONBody(resp.Body, &apiResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &types.TextGenerationResponse{
		Text: apiResp.Text,
	}, nil
}

func init() {
	providers.Register("amazon_bedrock_compat", func(cfg map[string]string) (providers.Provider, error) {
		apiKey, ok := cfg["api_key"]
		if !ok || apiKey == "" {
			return nil, errors.New("missing api_key in config")
		}
		endpoint, ok := cfg["endpoint"]
		if !ok || endpoint == "" {
			return nil, errors.New("missing endpoint in config")
		}
		return NewProvider(apiKey, endpoint), nil
	})
}
