package integration

import (
	"testing"

	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

func TestCapabilityOverride_GetProviderCapsThenOverride(t *testing.T) {
	ctx := t.Context()

	ps, err := newProviderSetWithDebug(0)
	if err != nil {
		t.Fatal(err)
	}

	pp, mp, err := addCatalogModelProvider(
		ctx,
		ps,
		modelpreset.ProviderOpenAIResponses,
		modelpreset.PresetOpenAIResponsesGPT5Mini,
	)
	if err != nil {
		t.Fatal(err)
	}

	baseCaps, err := ps.GetProviderCapability(ctx, pp.Name)
	if err != nil {
		t.Fatal(err)
	}

	override := &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	}

	caps := capabilityoverride.DeriveModelCapabilities(
		baseCaps,
		pp.CapabilitiesOverride,
		mp.CapabilitiesOverride,
		override,
	)

	completionKey := string(mp.ID)
	resolver := capabilityoverride.NewCompletionKeyResolver(completionKey, &caps)

	t.Run("modalities: file input rejected when fileIn unsupported", func(t *testing.T) {
		req := &spec.FetchCompletionRequest{
			ModelParam: spec.ModelParam{Name: mp.Name},
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

		_, _, _, err := sdkutil.NormalizeRequestForSDK(
			ctx,
			req,
			&spec.FetchCompletionOptions{
				CompletionKey:      completionKey,
				CapabilityResolver: resolver,
			},
			pp.SDKType,
			baseCaps,
		)
		if err == nil {
			t.Fatalf("expected modality error, got nil")
		}
	})

	t.Run("reasoning: unsupported level dropped with warning", func(t *testing.T) {
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

		capped, _, warns, err := sdkutil.NormalizeRequestForSDK(
			ctx,
			req,
			&spec.FetchCompletionOptions{
				CompletionKey:      completionKey,
				CapabilityResolver: resolver,
			},
			pp.SDKType,
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

	t.Run("tool outputs: rich function output collapses to string when resolver says string-only", func(t *testing.T) {
		newOverride := override
		newOverride.ModalitiesIn = []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityFileIn}
		caps := capabilityoverride.DeriveModelCapabilities(
			baseCaps,
			pp.CapabilitiesOverride,
			mp.CapabilitiesOverride,
			newOverride,
		)
		resolver := capabilityoverride.NewCompletionKeyResolver(completionKey, &caps)

		req := &spec.FetchCompletionRequest{
			ModelParam: spec.ModelParam{Name: mp.Name},
			Inputs: []spec.InputUnion{{
				Kind: spec.InputKindFunctionToolOutput,
				FunctionToolOutput: &spec.ToolOutput{
					Type:   spec.ToolTypeFunction,
					CallID: "call_1",
					Name:   "tool",
					Contents: []spec.ToolOutputItemUnion{
						{
							Kind:     spec.ContentItemKindText,
							TextItem: &spec.ContentItemText{Text: "hello"},
						},
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

		capped, _, warns, err := sdkutil.NormalizeRequestForSDK(
			ctx,
			req,
			&spec.FetchCompletionOptions{
				CompletionKey:      completionKey,
				CapabilityResolver: resolver,
			},
			pp.SDKType,
			baseCaps,
		)
		if err != nil {
			t.Fatal(err)
		}

		out := capped.Inputs[0].FunctionToolOutput
		if out == nil || len(out.Contents) != 1 || out.Contents[0].Kind != spec.ContentItemKindText {
			t.Fatalf("expected collapsed single text tool output; got %#v", out)
		}

		found := false
		for _, w := range warns {
			if w.Code == "toolOutput_collapsed_to_string" {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected toolOutput_collapsed_to_string warning; got %#v", warns)
		}
	})
}
