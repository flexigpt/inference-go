package inference

import (
	"context"

	"github.com/ppipada/inference-go/spec"
)

type FetchCompletionResponse struct {
	Outputs      []spec.OutputUnion `json:"outputs,omitempty"`
	Usage        *spec.Usage        `json:"usage,omitempty"`
	Error        *spec.Error        `json:"error,omitempty"`
	DebugDetails map[string]any     `json:"debugDetails,omitempty"`
}

type FetchCompletionRequest struct {
	ModelParam  spec.ModelParam   `json:"modelParam"`
	Inputs      []spec.InputUnion `json:"inputs"`
	ToolChoices []spec.ToolChoice `json:"toolChoices,omitempty"`
}

type CompletionProvider interface {
	InitLLM(ctx context.Context) error
	DeInitLLM(ctx context.Context) error
	GetProviderInfo(ctx context.Context) *spec.ProviderParam
	IsConfigured(ctx context.Context) bool
	SetProviderAPIKey(ctx context.Context, apiKey string) error
	FetchCompletion(
		ctx context.Context,
		fetchCompletionRequest *FetchCompletionRequest,
		onStreamTextData func(textData string) error,
		onStreamThinkingData func(thinkingData string) error,
	) (*FetchCompletionResponse, error)
}
