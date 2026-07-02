package integration

import (
	"context"
	"testing"

	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/openaichatsdk"
	"github.com/flexigpt/inference-go/internal/openairesponsessdk"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/modelpreset"
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
		Name:    modelpreset.ProviderOpenAIChat,
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
		ModelParam: spec.ModelParam{Name: modelpreset.ModelNameGPT41Mini},
		Inputs: []spec.InputUnion{
			newUserTextInput("hi"),
		},
		ToolChoices: []spec.ToolChoice{{
			Type: spec.ToolTypeWebSearch,
			ID:   "ws",
			Name: spec.DefaultWebSearchToolName,
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
		Name:    modelpreset.ProviderOpenAIResponses,
		SDKType: spec.ProviderSDKTypeOpenAIResponses,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	baseCaps, err := api.GetProviderCapability(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	pp, err := modelpreset.Provider(modelpreset.ProviderOpenAIResponses)
	if err != nil {
		t.Fatal(err)
	}
	mp, err := modelpreset.Model(
		modelpreset.ProviderOpenAIResponses,
		modelpreset.PresetGPT5Mini,
	)
	if err != nil {
		t.Fatal(err)
	}

	caps := capabilityoverride.DeriveModelCapabilities(
		baseCaps,
		pp.CapabilitiesOverride,
		mp.CapabilitiesOverride,
		&capabilityoverride.ModelCapabilitiesOverride{
			ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
				SupportedReasoningLevels: []spec.ReasoningLevel{
					spec.ReasoningLevelLow,
					spec.ReasoningLevelMedium,
					spec.ReasoningLevelHigh,
				},
			},
		},
	)

	req := &spec.FetchCompletionRequest{
		ModelParam: spec.ModelParam{
			Name: mp.Name,
			Reasoning: &spec.ReasoningParam{
				Type:  spec.ReasoningTypeSingleWithLevels,
				Level: spec.ReasoningLevelXHigh,
			},
		},
		Inputs: []spec.InputUnion{
			newUserTextInput("hi"),
		},
	}

	nreq, _, warns, err := sdkutil.NormalizeRequestForSDK(
		t.Context(),
		req,
		&spec.FetchCompletionOptions{
			CapabilityResolver: staticCapsResolver{Caps: &caps},
			CompletionKey:      string(mp.ID),
		},
		pp.SDKType,
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
			break
		}
	}
	if !found {
		t.Fatalf("expected reasoning_dropped_invalid_level warning; got %#v", warns)
	}
}
