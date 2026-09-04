package integration

import (
	"slices"
	"testing"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

func TestMetaMultiParameterNormalization(t *testing.T) {
	ctx := t.Context()

	providerSet, err := newProviderSetWithDebug(0)
	if err != nil {
		t.Fatalf("newProviderSetWithDebug: %v", err)
	}

	provider, model, err := addCatalogModelProvider(
		ctx,
		providerSet,
		modelpreset.ProviderMeta,
		modelpreset.PresetMuseSpark13Contributor,
	)
	if err != nil {
		t.Fatalf("addCatalogModelProvider: %v", err)
	}

	baseCapabilities, err := providerSet.GetProviderCapability(ctx, provider.Name)
	if err != nil {
		t.Fatalf("GetProviderCapability: %v", err)
	}

	options, err := presetFetchOptions(ctx, providerSet, provider, model)
	if err != nil {
		t.Fatalf("presetFetchOptions: %v", err)
	}

	omitted := spec.ReasoningSummaryStyleOmitted
	auto := spec.ReasoningSummaryStyleAuto
	concise := spec.ReasoningSummaryStyleConcise
	detailed := spec.ReasoningSummaryStyleDetailed

	tests := []struct {
		name         string
		level        spec.ReasoningLevel
		summaryStyle *spec.ReasoningSummaryStyle
		context      spec.ReasoningContext
		mode         spec.ReasoningMode
	}{
		{
			name:    "default summary style",
			level:   spec.ReasoningLevelMinimal,
			context: spec.ReasoningContextAuto,
			mode:    spec.ReasoningModeStandard,
		},
		{
			name:         "omitted summary style",
			level:        spec.ReasoningLevelLow,
			summaryStyle: &omitted,
			context:      spec.ReasoningContextCurrentTurn,
			mode:         spec.ReasoningModePro,
		},
		{
			name:         "auto summary style",
			level:        spec.ReasoningLevelMedium,
			summaryStyle: &auto,
			context:      spec.ReasoningContextAllTurns,
			mode:         spec.ReasoningModeStandard,
		},
		{
			name:         "concise summary style",
			level:        spec.ReasoningLevelHigh,
			summaryStyle: &concise,
			context:      spec.ReasoningContextAuto,
			mode:         spec.ReasoningModePro,
		},
		{
			name:         "detailed summary style",
			level:        spec.ReasoningLevelXHigh,
			summaryStyle: &detailed,
			context:      spec.ReasoningContextAllTurns,
			mode:         spec.ReasoningModeStandard,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := metaRequest(model.ModelParam, tc)

			normalized, capabilities, warnings, err := sdkutil.NormalizeRequestForSDK(
				ctx,
				request,
				options,
				provider.SDKType,
				baseCapabilities,
			)
			if err != nil {
				t.Fatalf("NormalizeRequestForSDK: %v", err)
			}
			if capabilities == nil || capabilities.ReasoningCapabilities == nil {
				t.Fatal("expected effective reasoning capabilities")
			}

			reasoningCaps := capabilities.ReasoningCapabilities
			if !reasoningCaps.SupportsSummaryStyle {
				t.Fatal("Meta Muse Spark 1.3 Contributor must support summary style")
			}
			if reasoningCaps.SupportsReasoningContext {
				t.Fatal("Meta Muse Spark 1.3 Contributor must not support reasoning context")
			}
			if reasoningCaps.SupportsReasoningMode {
				t.Fatal("Meta Muse Spark 1.3 Contributor must not support reasoning mode")
			}

			reasoning := normalized.ModelParam.Reasoning
			if reasoning == nil {
				t.Fatal("reasoning unexpectedly dropped")
			}
			if reasoning.Level != tc.level {
				t.Fatalf("reasoning level got %q want %q", reasoning.Level, tc.level)
			}
			if reasoning.SummaryStyle != tc.summaryStyle &&
				(reasoning.SummaryStyle == nil || tc.summaryStyle == nil ||
					*reasoning.SummaryStyle != *tc.summaryStyle) {
				t.Fatalf("summaryStyle got %#v want %#v", reasoning.SummaryStyle, tc.summaryStyle)
			}
			if reasoning.Context != nil {
				t.Fatalf("reasoning context should be dropped, got %q", *reasoning.Context)
			}
			if reasoning.Mode != nil {
				t.Fatalf("reasoning mode should be dropped, got %q", *reasoning.Mode)
			}

			if normalized.ModelParam.CacheControl == nil ||
				normalized.ModelParam.CacheControl.TTL != spec.CacheControlTTL24h ||
				normalized.ModelParam.CacheControl.Key != "meta-muse-cache-key" {
				t.Fatalf("cache control unexpectedly changed: %#v", normalized.ModelParam.CacheControl)
			}
			if normalized.ModelParam.OutputParam == nil ||
				normalized.ModelParam.OutputParam.Format == nil ||
				normalized.ModelParam.OutputParam.Format.Kind != spec.OutputFormatKindJSONSchema {
				t.Fatalf("JSON Schema output unexpectedly changed: %#v", normalized.ModelParam.OutputParam)
			}
			if len(normalized.ToolChoices) != 1 ||
				normalized.ToolChoices[0].Type != spec.ToolTypeFunction {
				t.Fatalf("function tool unexpectedly changed: %#v", normalized.ToolChoices)
			}
			if normalized.ToolPolicy == nil || normalized.ToolPolicy.Mode != spec.ToolPolicyModeAuto {
				t.Fatalf("tool policy unexpectedly changed: %#v", normalized.ToolPolicy)
			}

			gotWarningCodes := make([]string, 0, len(warnings))
			for _, warning := range warnings {
				gotWarningCodes = append(gotWarningCodes, warning.Code)
			}
			wantWarningCodes := []string{
				"reasoning_context_dropped_unsupported",
				"reasoning_mode_dropped_unsupported",
			}
			if !slices.Equal(gotWarningCodes, wantWarningCodes) {
				t.Fatalf(
					"warning codes got %#v want %#v",
					gotWarningCodes,
					wantWarningCodes,
				)
			}
		})
	}
}

func metaRequest(
	modelParam spec.ModelParam,
	tc struct {
		name         string
		level        spec.ReasoningLevel
		summaryStyle *spec.ReasoningSummaryStyle
		context      spec.ReasoningContext
		mode         spec.ReasoningMode
	},
) *spec.FetchCompletionRequest {
	modelParam.Reasoning = &spec.ReasoningParam{
		Type:         spec.ReasoningTypeSingleWithLevels,
		Level:        tc.level,
		SummaryStyle: tc.summaryStyle,
		Context:      &tc.context,
		Mode:         &tc.mode,
	}
	modelParam.CacheControl = &spec.CacheControl{
		Kind: spec.CacheControlKindEphemeral,
		TTL:  spec.CacheControlTTL24h,
		Key:  "meta-muse-cache-key",
	}
	modelParam.OutputParam = &spec.OutputParam{
		Format: &spec.OutputFormat{
			Kind: spec.OutputFormatKindJSONSchema,
			JSONSchemaParam: &spec.JSONSchemaParam{
				Name: "meta_muse_response",
				Schema: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"answer": map[string]any{"type": "string"},
					},
					"required":             []any{"answer"},
					"additionalProperties": false,
				},
				Strict: true,
			},
		},
	}

	return &spec.FetchCompletionRequest{
		ModelParam: modelParam,
		Inputs: []spec.InputUnion{{
			Kind: spec.InputKindInputMessage,
			InputMessage: &spec.InputOutputContent{
				Role: spec.RoleUser,
				Contents: []spec.InputOutputContentItemUnion{
					{
						Kind:     spec.ContentItemKindText,
						TextItem: &spec.ContentItemText{Text: "Describe the attached inputs."},
					},
					{
						Kind: spec.ContentItemKindImage,
						ImageItem: &spec.ContentItemImage{
							ImageURL:  "https://example.com/example.png",
							ImageMIME: "image/png",
							Detail:    spec.ImageDetailLow,
						},
					},
					{
						Kind: spec.ContentItemKindFile,
						FileItem: &spec.ContentItemFile{
							FileURL:  "https://example.com/example.pdf",
							FileMIME: "application/pdf",
						},
					},
				},
			},
		}},
		ToolChoices: []spec.ToolChoice{
			newEchoToolChoice("meta-echo-tool", "meta_echo"),
		},
		ToolPolicy: &spec.ToolPolicy{
			Mode: spec.ToolPolicyModeAuto,
		},
	}
}
