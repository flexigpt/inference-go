package sdkutil

import (
	"context"

	"github.com/flexigpt/inference-go/spec"
)

type CompletionProvider interface {
	InitLLM(ctx context.Context) error
	DeInitLLM(ctx context.Context) error
	GetProviderInfo(ctx context.Context) *spec.ProviderParam
	IsConfigured(ctx context.Context) bool
	SetProviderAPIKey(ctx context.Context, apiKey string) error
	GetProviderCapability(ctx context.Context) (spec.ModelCapabilities, error)

	FetchCompletion(
		ctx context.Context,
		fetchCompletionRequest *spec.FetchCompletionRequest,
		opts *spec.FetchCompletionOptions,
	) (*spec.FetchCompletionResponse, error)
}
