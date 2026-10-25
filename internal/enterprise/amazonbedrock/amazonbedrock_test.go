package amazonbedrock

import (
	"context"
	"testing"
)

type mockClient struct {
	response string
	err      error
}

func (m *mockClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	return m.response, m.err
}

func TestProvider_Generate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		client     Client
		prompt     string
		want       string
		expectErr  bool
		errMessage string
	}{
		{
			name:   "successful generation",
			client: &mockClient{response: "response text", err: nil},
			prompt: "hello",
			want:   "response text",
		},
		{
			name:       "client returns error",
			client:     &mockClient{response: "", err: errTest},
			prompt:     "hello",
			expectErr:  true,
			errMessage: "test error",
		},
		{
			name:       "nil client",
			client:     nil,
			prompt:     "hello",
			expectErr:  true,
			errMessage: "client is not initialized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.client)
			got, err := p.Generate(ctx, tt.prompt)
			if (err != nil) != tt.expectErr {
				t.Fatalf("Generate() error = %v, wantErr %v", err, tt.expectErr)
			}
			if err != nil && tt.errMessage != "" && err.Error() != tt.errMessage {
				t.Errorf("Generate() error message = %v, want %v", err.Error(), tt.errMessage)
			}
			if got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

var errTest = &testError{"test error"}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
