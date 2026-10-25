package amazon_bedrock_compat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ourorg/primaryrepo/enterprise/deferred"
	"github.com/ourorg/primaryrepo/logging"
)

// Provider implements the Amazon Bedrock compatibility path for Enterprise Deferred.
type Provider struct {
	httpClient *http.Client
	endpoint   string
	apiKey     string
	logger     logging.Logger
}

// NewProvider creates a new Amazon Bedrock compat provider with the given config.
func NewProvider(endpoint, apiKey string, logger logging.Logger) *Provider {
	return &Provider{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		endpoint:   endpoint,
		apiKey:     apiKey,
		logger:     logger,
	}
}

// GenerateText implements the deferred.Provider interface for text generation.
func (p *Provider) GenerateText(ctx context.Context, prompt string, opts deferred.GenerationOptions) (string, error) {
	req, err := p.newRequest(ctx, prompt, opts)
	if err != nil {
		p.logger.Error("failed to create request", "error", err)
		return "", err
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		p.logger.Error("request failed", "error", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		p.logger.Error("non-200 response", "status", resp.StatusCode)
		return "", fmt.Errorf("amazon bedrock compat: unexpected status code %d", resp.StatusCode)
	}

	text, err := parseResponse(resp)
	if err != nil {
		p.logger.Error("failed to parse response", "error", err)
		return "", err
	}

	return text, nil
}

func (p *Provider) newRequest(ctx context.Context, prompt string, opts deferred.GenerationOptions) (*http.Request, error) {
	payload := map[string]interface{}{
		"prompt": prompt,
		"max_tokens": opts.MaxTokens,
		"temperature": opts.Temperature,
		"top_p": opts.TopP,
		// Add other Bedrock compat parameters as needed.
	}

	req, err := deferred.NewJSONRequest(ctx, "POST", p.endpoint+"/generate", payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "ApiKey "+p.apiKey)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func parseResponse(resp *http.Response) (string, error) {
	var parsed struct {
		Completion string `json:"completion"`
	}
	err := deferred.DecodeJSON(resp.Body, &parsed)
	if err != nil {
		return "", err
	}
	if parsed.Completion == "" {
		return "", errors.New("empty completion in response")
	}
	return parsed.Completion, nil
}
