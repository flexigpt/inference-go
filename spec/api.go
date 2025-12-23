package spec

import (
	"context"
)

type FetchCompletionResponse struct {
	Outputs      []OutputUnion  `json:"outputs,omitempty"`
	Usage        *Usage         `json:"usage,omitempty"`
	Error        *Error         `json:"error,omitempty"`
	DebugDetails map[string]any `json:"debugDetails,omitempty"`
}

type FetchCompletionRequest struct {
	ModelParam  ModelParam   `json:"modelParam"`
	Inputs      []InputUnion `json:"inputs"`
	ToolChoices []ToolChoice `json:"toolChoices,omitempty"`
}

type CompletionProvider interface {
	InitLLM(ctx context.Context) error
	DeInitLLM(ctx context.Context) error
	GetProviderInfo(ctx context.Context) *ProviderParam
	IsConfigured(ctx context.Context) bool
	SetProviderAPIKey(ctx context.Context, apiKey string) error
	FetchCompletion(
		ctx context.Context,
		fetchCompletionRequest *FetchCompletionRequest,
		onStreamTextData func(textData string) error,
		onStreamThinkingData func(thinkingData string) error,
	) (*FetchCompletionResponse, error)
}
