package integration

import (
	"context"
	"errors"
	"testing"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

// overrideResolver is a minimal ModelCapabilityResolver example used in tests.
// In a real app you might:
//   - keep a per-model capability table,
//   - or derive capabilities from an upstream registry.
type overrideResolver struct {
	// key: model name
	byModel map[spec.ModelName]*spec.ModelCapabilities
}

func (r overrideResolver) ResolveModelCapabilities(
	ctx context.Context,
	req spec.ResolveModelCapabilitiesRequest,
) (*spec.ModelCapabilities, error) {
	if r.byModel == nil {
		return nil, errors.New("invalid model")
	}
	if c := r.byModel[req.ModelName]; c != nil {
		return c, nil
	}
	return nil, errors.New("model not found")
}

func TestCapabilityOverride_GetProviderCapsThenOverride(t *testing.T) {
	ctx := t.Context()

	ps, err := inference.NewProviderSetAPI()
	if err != nil {
		t.Fatal(err)
	}

	_, err = ps.AddProvider(ctx, "openai-responses", &inference.AddProviderConfig{
		SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
		Origin:                   spec.DefaultOpenAIOrigin,
		ChatCompletionPathPrefix: "/v1/responses",
		APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 1) Get SDK-wide default capabilities programmatically.
	baseCaps, err := ps.GetProviderCapability(ctx, "openai-responses")
	if err != nil {
		t.Fatal(err)
	}

	// 2) Override per-model capabilities.
	//
	// This is an example override (not necessarily reflecting OpenAI’s real limits):
	//   - disallow file input
	//   - disallow reasoning level xhigh
	//
	// The point is: *capabilities are the authoritative enforcement mechanism*.
	override := baseCaps
	override.ModalitiesIn = []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn} // drop fileIn
	if override.ReasoningCapabilities != nil {
		override.ReasoningCapabilities.SupportedReasoningLevels = []spec.ReasoningLevel{
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
		}
	}

	resolver := overrideResolver{
		byModel: map[spec.ModelName]*spec.ModelCapabilities{
			"gpt-5-mini": &override,
		},
	}

	t.Run("modalities: file input rejected when fileIn unsupported", func(t *testing.T) {
		req := &spec.FetchCompletionRequest{
			ModelParam: spec.ModelParam{Name: "gpt-5-mini"},
			Inputs: []spec.InputUnion{{
				Kind: spec.InputKindInputMessage,
				InputMessage: &spec.InputOutputContent{
					Role: spec.RoleUser,
					Contents: []spec.InputOutputContentItemUnion{
						{Kind: spec.ContentItemKindText, TextItem: &spec.ContentItemText{Text: "hi"}},
						{
							Kind: spec.ContentItemKindFile,
							FileItem: &spec.ContentItemFile{
								FileURL:  "https://example.com/a.pdf",
								FileMIME: "application/pdf",
							},
						},
					},
				},
			}},
		}

		_, _, err := sdkutil.NormalizeRequestForSDK(
			ctx,
			req,
			&spec.FetchCompletionOptions{
				CompletionKey:      "gpt5mini",
				CapabilityResolver: resolver,
			},
			spec.ProviderSDKTypeOpenAIResponses,
			baseCaps,
		)
		if err == nil {
			t.Fatalf("expected modality error, got nil")
		}
	})

	t.Run("reasoning: unsupported level dropped with warning", func(t *testing.T) {
		req := &spec.FetchCompletionRequest{
			ModelParam: spec.ModelParam{
				Name: "gpt-5-mini",
				Reasoning: &spec.ReasoningParam{
					Type:  spec.ReasoningTypeSingleWithLevels,
					Level: spec.ReasoningLevelXHigh,
				},
			},
			Inputs: []spec.InputUnion{{
				Kind: spec.InputKindInputMessage,
				InputMessage: &spec.InputOutputContent{
					Role: spec.RoleUser,
					Contents: []spec.InputOutputContentItemUnion{
						{Kind: spec.ContentItemKindText, TextItem: &spec.ContentItemText{Text: "hi"}},
					},
				},
			}},
		}

		capped, warns, err := sdkutil.NormalizeRequestForSDK(
			ctx,
			req,
			&spec.FetchCompletionOptions{
				CompletionKey:      "gpt5mini",
				CapabilityResolver: resolver,
			},
			spec.ProviderSDKTypeOpenAIResponses,
			baseCaps,
		)
		if err != nil {
			t.Fatal(err)
		}
		if capped.ModelParam.Reasoning != nil {
			t.Fatalf("expected reasoning dropped, got %#v", capped.ModelParam.Reasoning)
		}
		found := false
		for _, w := range warns {
			if w.Code == "reasoning_dropped_invalid_level" {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected reasoning_dropped_invalid_level warning; got %#v", warns)
		}
	})
}
