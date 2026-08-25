package inference

import (
	"slices"
	"testing"

	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

func TestCatalogPresetsIntegrateWithProviderSet(t *testing.T) {
	ctx := t.Context()
	catalog := modelpreset.DefaultCatalog()

	providerSet, err := NewProviderSetAPI()
	if err != nil {
		t.Fatalf("NewProviderSetAPI: %v", err)
	}

	type providerCase struct {
		name   spec.ProviderName
		preset modelpreset.ProviderPreset
	}

	providerNames := modelpreset.ProviderNames()
	providerCases := make([]providerCase, 0, len(providerNames))

	for _, providerName := range providerNames {
		providerCases = append(providerCases, providerCase{
			name:   providerName,
			preset: catalog.Providers[providerName],
		})
	}

	for _, tc := range providerCases {
		t.Run("add/"+string(tc.name), func(t *testing.T) {
			got, err := providerSet.AddProviderFromPreset(ctx, tc.name, tc.preset)
			if err != nil {
				t.Fatalf("AddProviderFromPreset(%q): %v", tc.name, err)
			}

			if got.Name != tc.name {
				t.Fatalf("provider name got %q want %q", got.Name, tc.name)
			}
			if got.SDKType != tc.preset.SDKType {
				t.Fatalf("sdkType got %q want %q", got.SDKType, tc.preset.SDKType)
			}
			if got.Origin != tc.preset.Origin {
				t.Fatalf("origin got %q want %q", got.Origin, tc.preset.Origin)
			}
			if got.ChatCompletionPathPrefix != tc.preset.ChatCompletionPathPrefix {
				t.Fatalf(
					"path prefix got %q want %q",
					got.ChatCompletionPathPrefix,
					tc.preset.ChatCompletionPathPrefix,
				)
			}
			if got.APIKeyHeaderKey != tc.preset.APIKeyHeaderKey {
				t.Fatalf(
					"apiKeyHeaderKey got %q want %q",
					got.APIKeyHeaderKey,
					tc.preset.APIKeyHeaderKey,
				)
			}
		})
	}

	for _, providerCase := range providerCases {
		modelIDs := make([]modelpreset.ModelPresetID, 0, len(providerCase.preset.ModelPresets))
		for modelID := range providerCase.preset.ModelPresets {
			modelIDs = append(modelIDs, modelID)
		}
		slices.Sort(modelIDs)

		for _, modelID := range modelIDs {
			model := providerCase.preset.ModelPresets[modelID]

			t.Run(
				"model/"+string(providerCase.name)+"/"+string(modelID),
				func(t *testing.T) {
					baseCapabilities, err := providerSet.GetProviderCapability(ctx, providerCase.name)
					if err != nil {
						t.Fatalf(
							"GetProviderCapability(%q): %v",
							providerCase.name,
							err,
						)
					}

					completionKey := string(model.ID)
					resolver, err := providerSet.NewPresetCapabilityResolver(
						ctx,
						providerCase.name,
						providerCase.preset,
						model,
						completionKey,
					)
					if err != nil {
						t.Fatalf("NewPresetCapabilityResolver: %v", err)
					}

					effectiveCapabilities, err := resolver.ResolveModelCapabilities(
						ctx,
						spec.ResolveModelCapabilitiesRequest{
							ProviderSDKType: providerCase.preset.SDKType,
							ModelName:       model.Name,
							CompletionKey:   completionKey,
						},
					)
					if err != nil {
						t.Fatalf("ResolveModelCapabilities: %v", err)
					}
					if effectiveCapabilities == nil {
						t.Fatal("ResolveModelCapabilities returned nil capabilities")
					}

					assertCatalogPresetDefaultCompatible(
						t,
						providerCase.name,
						modelID,
						model,
						effectiveCapabilities,
					)

					normalized, normalizedCapabilities, _, err := sdkutil.NormalizeRequestForSDK(
						ctx,
						&spec.FetchCompletionRequest{
							ModelParam: model.ModelParam,
							Inputs: []spec.InputUnion{
								catalogPresetTextInput(),
							},
						},
						&spec.FetchCompletionOptions{
							CompletionKey:      completionKey,
							CapabilityResolver: resolver,
						},
						providerCase.preset.SDKType,
						baseCapabilities,
					)
					if err != nil {
						t.Fatalf("NormalizeRequestForSDK: %v", err)
					}
					if normalized == nil {
						t.Fatal("NormalizeRequestForSDK returned nil request")
					}
					if normalizedCapabilities == nil {
						t.Fatal("NormalizeRequestForSDK returned nil capabilities")
					}
					if normalized.ModelParam.Name != model.ModelParam.Name {
						t.Fatalf(
							"normalized model name got %q want %q",
							normalized.ModelParam.Name,
							model.ModelParam.Name,
						)
					}
				},
			)
		}
	}
}

func assertCatalogPresetDefaultCompatible(
	t *testing.T,
	providerName spec.ProviderName,
	modelID modelpreset.ModelPresetID,
	model modelpreset.ModelPreset,
	capabilities *spec.ModelCapabilities,
) {
	t.Helper()

	if model.ModelParam.Temperature == nil && model.ModelParam.Reasoning == nil {
		t.Fatalf(
			"%s/%s has neither temperature nor reasoning",
			providerName,
			modelID,
		)
	}

	if reasoning := model.ModelParam.Reasoning; reasoning != nil {
		reasoningCapabilities := capabilities.ReasoningCapabilities
		if reasoningCapabilities == nil || !reasoningCapabilities.SupportsReasoningConfig {
			t.Fatalf(
				"%s/%s default reasoning is unsupported by effective capabilities",
				providerName,
				modelID,
			)
		}
		if !slices.Contains(reasoningCapabilities.SupportedReasoningTypes, reasoning.Type) {
			t.Fatalf(
				"%s/%s reasoning type %q is unsupported by effective capabilities",
				providerName,
				modelID,
				reasoning.Type,
			)
		}

		switch reasoning.Type {
		case spec.ReasoningTypeSingleWithLevels:
			if !slices.Contains(
				reasoningCapabilities.SupportedReasoningLevels,
				reasoning.Level,
			) {
				t.Fatalf(
					"%s/%s reasoning level %q is unsupported by effective capabilities",
					providerName,
					modelID,
					reasoning.Level,
				)
			}

		case spec.ReasoningTypeHybridWithTokens:
			budget := reasoningCapabilities.HybridTokenBudgetCapabilities
			if budget == nil {
				return
			}
			if budget.MinAllowed > 0 && reasoning.Tokens < budget.MinAllowed {
				t.Fatalf(
					"%s/%s reasoning tokens %d are below minAllowed %d",
					providerName,
					modelID,
					reasoning.Tokens,
					budget.MinAllowed,
				)
			}
			if budget.MaxAllowed > 0 && reasoning.Tokens > budget.MaxAllowed {
				t.Fatalf(
					"%s/%s reasoning tokens %d exceed maxAllowed %d",
					providerName,
					modelID,
					reasoning.Tokens,
					budget.MaxAllowed,
				)
			}
		}
	}

	if len(model.ModelParam.StopSequences) > 0 {
		stopCapabilities := capabilities.StopSequenceCapabilities
		if stopCapabilities == nil || !stopCapabilities.IsSupported {
			t.Fatalf(
				"%s/%s has unsupported default stop sequences",
				providerName,
				modelID,
			)
		}
		if stopCapabilities.MaxSequences > 0 &&
			len(model.ModelParam.StopSequences) > stopCapabilities.MaxSequences {
			t.Fatalf(
				"%s/%s has %d stop sequences but max is %d",
				providerName,
				modelID,
				len(model.ModelParam.StopSequences),
				stopCapabilities.MaxSequences,
			)
		}
	}

	if output := model.ModelParam.OutputParam; output != nil {
		outputCapabilities := capabilities.OutputCapabilities
		if output.Format != nil &&
			(outputCapabilities == nil ||
				!slices.Contains(
					outputCapabilities.SupportedOutputFormats,
					output.Format.Kind,
				)) {
			t.Fatalf(
				"%s/%s default output format %q is unsupported",
				providerName,
				modelID,
				output.Format.Kind,
			)
		}
		if output.Verbosity != nil &&
			(outputCapabilities == nil || !outputCapabilities.SupportsVerbosity) {
			t.Fatalf(
				"%s/%s default output verbosity is unsupported",
				providerName,
				modelID,
			)
		}
	}

	if cacheControl := model.ModelParam.CacheControl; cacheControl != nil {
		if capabilities.CacheCapabilities == nil ||
			capabilities.CacheCapabilities.TopLevel == nil {
			t.Fatalf(
				"%s/%s has unsupported top-level cache control",
				providerName,
				modelID,
			)
		}

		cacheCapabilities := capabilities.CacheCapabilities.TopLevel
		if cacheControl.Kind != "" &&
			len(cacheCapabilities.SupportedKinds) > 0 &&
			!slices.Contains(cacheCapabilities.SupportedKinds, cacheControl.Kind) {
			t.Fatalf(
				"%s/%s cache kind %q is unsupported",
				providerName,
				modelID,
				cacheControl.Kind,
			)
		}
		if cacheControl.TTL != "" &&
			(!cacheCapabilities.SupportsTTL ||
				(len(cacheCapabilities.SupportedTTLs) > 0 &&
					!slices.Contains(cacheCapabilities.SupportedTTLs, cacheControl.TTL))) {
			t.Fatalf(
				"%s/%s cache TTL %q is unsupported",
				providerName,
				modelID,
				cacheControl.TTL,
			)
		}
		if cacheControl.Key != "" && !cacheCapabilities.SupportsKey {
			t.Fatalf(
				"%s/%s cache key is unsupported",
				providerName,
				modelID,
			)
		}
	}
}

func catalogPresetTextInput() spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			Role: spec.RoleUser,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: "catalog validation"},
				},
			},
		},
	}
}
