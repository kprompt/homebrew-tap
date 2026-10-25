// Package amazonbedrock provides an Enterprise Deferred provider implementation
// for Amazon Bedrock with compatibility path support.
//
// This package implements the P-008 capability.
package amazonbedrock

import (
	"context"
	"errors"
	"fmt"
)

// ProviderName is the canonical name of this provider.
const ProviderName = "amazonbedrock"

// Client defines the interface for the Amazon Bedrock client.
type Client interface {
	// GenerateText generates text from a prompt.
	GenerateText(ctx context.Context, prompt string) (string, error)
}

// Provider implements the Enterprise Deferred provider interface for Amazon Bedrock.
type Provider struct {
	client Client
}

// New returns a new Amazon Bedrock provider instance.
func New(client Client) *Provider {
	return &Provider{client: client}
}

// Generate calls the underlying client to generate text from the prompt.
func (p *Provider) Generate(ctx context.Context, prompt string) (string, error) {
	if p.client == nil {
		return "", errors.New("amazonbedrock: client is not initialized")
	}
	return p.client.GenerateText(ctx, prompt)
}

// String returns the provider name.
func (p *Provider) String() string {
	return ProviderName
}

// ExampleClient is a simple implementation of Client for demonstration.
type ExampleClient struct{}

// GenerateText returns a dummy response for the prompt.
func (c *ExampleClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}
	return fmt.Sprintf("Amazon Bedrock response to: %s", prompt), nil
}
