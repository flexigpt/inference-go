package integration

import (
	"context"
	"testing"

	"github.com/flexigpt/inference-go/internal/openaichatsdk"
	"github.com/flexigpt/inference-go/internal/openairesponsessdk"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

type staticCapsResolver struct {
	Caps *spec.ModelCapabilities
}

func (r staticCapsResolver) ResolveModelCapabilities(
	ctx context.Context,
	req spec.ResolveModelCapabilitiesRequest,
) (*spec.ModelCapabilities, error) {
	return r.Caps, nil
}

func TestNormalizeRequestForSDK_OpenAIChat_PreservesWebSearchToolChoice(t *testing.T) {
	api, err := openaichatsdk.NewOpenAIChatCompletionsAPI(spec.ProviderParam{
		Name:    "openai-chat",
		SDKType: spec.ProviderSDKTypeOpenAIChatCompletions,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	caps, err := api.GetProviderCapability(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{Name: "gpt-4.1-mini"},
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "hi"},
				}},
			},
		}},
		ToolChoices: []spec.ToolChoice{{
			Type: spec.ToolTypeWebSearch,
			ID:   "ws",
			Name: "web_search",
			WebSearchArguments: &spec.WebSearchToolChoiceItem{
				SearchContextSize: spec.WebSearchContextSizeMedium,
			},
		}},
	}

	nreq, _, warns, err := sdkutil.NormalizeRequestForSDK(
		t.Context(),
		req,
		&spec.FetchCompletionOptions{},
		spec.ProviderSDKTypeOpenAIChatCompletions,
		caps,
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(warns) != 0 {
		t.Fatalf("expected no warnings; got %#v", warns)
	}
	if len(nreq.ToolChoices) != 1 || nreq.ToolChoices[0].Type != spec.ToolTypeWebSearch {
		t.Fatalf("expected webSearch toolChoice preserved; got %#v", nreq.ToolChoices)
	}
}

func TestNormalizeRequestForSDK_ResolverRestrictsReasoningLevels(t *testing.T) {
	api, err := openairesponsessdk.NewOpenAIResponsesAPI(spec.ProviderParam{
		Name:    "openai-responses",
		SDKType: spec.ProviderSDKTypeOpenAIResponses,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	baseCaps, err := api.GetProviderCapability(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	// Example: model does NOT allow reasoning level none/xhigh (per your prompt).
	custom := baseCaps
	custom.ReasoningCapabilities = &spec.ReasoningCapabilities{
		SupportsReasoningConfig: true,
		SupportedReasoningTypes: []spec.ReasoningType{spec.ReasoningTypeSingleWithLevels},
		SupportedReasoningLevels: []spec.ReasoningLevel{
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
		},
		SupportsSummaryStyle:             true,
		SupportsEncryptedReasoningInput:  false,
		TemperatureDisallowedWhenEnabled: false,
	}

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
				Contents: []spec.InputOutputContentItemUnion{{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "hi"},
				}},
			},
		}},
	}

	nreq, _, warns, err := sdkutil.NormalizeRequestForSDK(
		t.Context(),
		req,
		&spec.FetchCompletionOptions{
			CapabilityResolver: staticCapsResolver{Caps: &custom},
			CompletionKey:      "gpt5mini",
		},
		spec.ProviderSDKTypeOpenAIResponses,
		baseCaps,
	)
	if err != nil {
		t.Fatal(err)
	}
	if nreq.ModelParam.Reasoning != nil {
		t.Fatalf("expected reasoning to be dropped; got %#v", nreq.ModelParam.Reasoning)
	}
	found := false
	for _, w := range warns {
		if w.Code == "reasoning_dropped_invalid_level" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected reasoning_dropped_invalid_level warning; got %#v", warns)
	}
}
