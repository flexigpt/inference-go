package modelpreset

import (
	"errors"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

func TestDefaultCatalogValidatesEveryProvider(t *testing.T) {
	catalog := DefaultCatalog()

	if err := ValidateCatalog(catalog); err != nil {
		t.Fatalf("ValidateCatalog(DefaultCatalog()): %v", err)
	}

	providerNames := ProviderNames()
	tests := make([]struct {
		name     spec.ProviderName
		provider ProviderPreset
	}, 0, len(providerNames))

	for _, providerName := range providerNames {
		tests = append(tests, struct {
			name     spec.ProviderName
			provider ProviderPreset
		}{
			name:     providerName,
			provider: catalog.Providers[providerName],
		})
	}

	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			singleProviderCatalog := Catalog{
				Providers: map[spec.ProviderName]ProviderPreset{
					tc.name: tc.provider,
				},
			}

			if err := ValidateCatalog(singleProviderCatalog); err != nil {
				t.Fatalf("provider catalog validation failed: %v", err)
			}
		})
	}
}

func TestCatalogContainsAllRegisteredProviders(t *testing.T) {
	tests := []struct {
		name string
		want []spec.ProviderName
	}{
		{
			name: "all supported provider registrations are present",
			want: []spec.ProviderName{
				ProviderAnthropic,
				ProviderDeepSeek,
				ProviderGoogleGemini,
				ProviderHuggingFace,
				ProviderLlamaCPP,
				ProviderLMStudio,
				ProviderLocalAI,
				ProviderMeta,
				ProviderMiniMax,
				ProviderMistral,
				ProviderMoonshot,
				ProviderOllama,
				ProviderOpenAIChat,
				ProviderOpenAIResponses,
				ProviderOpenRouter,
				ProviderQwen,
				ProviderSGLang,
				ProviderVLLM,
				ProviderXAI,
				ProviderXiaomi,
				ProviderZAI,
				ProviderZAICodingPlan,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ProviderNames()
			want := slices.Clone(tc.want)
			slices.Sort(want)

			if !slices.Equal(got, want) {
				t.Fatalf("ProviderNames() mismatch\n got: %#v\nwant: %#v", got, want)
			}

			catalog := DefaultCatalog()
			if len(catalog.Providers) != len(want) {
				t.Fatalf(
					"catalog provider count got %d want %d",
					len(catalog.Providers),
					len(want),
				)
			}
		})
	}
}

func TestNewProviderModelMembership(t *testing.T) {
	tests := []struct {
		name     spec.ProviderName
		modelIDs []ModelPresetID
	}{
		{
			name: ProviderDeepSeek,
			modelIDs: []ModelPresetID{
				PresetDeepSeekV4Flash,
				PresetDeepSeekV4Pro,
			},
		},
		{
			name: ProviderMeta,
			modelIDs: []ModelPresetID{
				PresetMuseSpark13,
				PresetMuseSpark13Contributor,
				PresetMuseSpark12,
				PresetMuseSpark12Contributor,
				PresetMuseSpark11,
			},
		},
		{
			name: ProviderMiniMax,
			modelIDs: []ModelPresetID{
				PresetMiniMaxM2,
				PresetMiniMaxM21,
				PresetMiniMaxM21Highspeed,
				PresetMiniMaxM25,
				PresetMiniMaxM25Highspeed,
				PresetMiniMaxM27,
				PresetMiniMaxM27Highspeed,
				PresetMiniMaxM3,
			},
		},
		{
			name: ProviderMoonshot,
			modelIDs: []ModelPresetID{
				PresetMoonshotKimiK3,
				PresetMoonshotKimiK26,
				PresetMoonshotKimiK27Code,
				PresetMoonshotKimiK27CodeHighspeed,
			},
		},
		{
			name: ProviderQwen,
			modelIDs: []ModelPresetID{
				PresetQwen38Max,
				PresetQwen3824TA95B,
				PresetQwen3827B,
				PresetQwen37Max,
				PresetQwen37Max20260520,
				PresetQwen37Max20260608,
				PresetQwen3Max,
				PresetQwen3Max20260123,
				PresetQwen37Plus,
				PresetQwen37Plus20260526,
				PresetQwen37Flash,
				PresetQwen37Flash20260715,
				PresetQwen36Plus,
				PresetQwen36Plus20260402,
				PresetQwen36Flash,
				PresetQwen36Flash20260416,
				PresetQwen3635BA3BDirect,
				PresetQwen3627BDirect,
				PresetQwen35Plus,
				PresetQwen35Plus20260215,
				PresetQwen35Plus20260420,
				PresetQwen35Flash,
				PresetQwen35Flash20260223,
				PresetQwen35397BA17B,
				PresetQwen35122BA10B,
				PresetQwen3527B,
				PresetQwen3535BA3B,
				PresetQwenPlus,
				PresetQwenFlash,
				PresetQwen3CoderPlus,
				PresetQwen3CoderFlash,
				PresetQwenPlusCharacter,
				PresetQwenFlashCharacter,
			},
		},
		{
			name: ProviderXiaomi,
			modelIDs: []ModelPresetID{
				PresetMiMoV25,
				PresetMiMoV25Pro,
			},
		},
		{
			name: ProviderZAI,
			modelIDs: []ModelPresetID{
				PresetGLM53,
				PresetGLM52,
				PresetGLM51,
				PresetGLM5,
				PresetGLM5Turbo,
				PresetGLM47,
				PresetGLM47FlashX,
				PresetGLM47Flash,
				PresetGLM46,
				PresetGLM45,
				PresetGLM45X,
				PresetGLM45Air,
				PresetGLM45AirX,
				PresetGLM45Flash,
				PresetGLM5VTurbo,
				PresetGLM46V,
				PresetGLM46VFlash,
				PresetGLM46VFlashX,
				PresetGLM45V,
			},
		},
		{
			name: ProviderZAICodingPlan,
			modelIDs: []ModelPresetID{
				PresetGLM53,
				PresetGLM5Turbo,
				PresetGLM47,
			},
		},
	}

	catalog := DefaultCatalog()

	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			provider, ok := catalog.Providers[tc.name]
			if !ok {
				t.Fatalf("provider %q is not registered", tc.name)
			}

			got := make([]ModelPresetID, 0, len(provider.ModelPresets))
			for id := range provider.ModelPresets {
				got = append(got, id)
			}
			slices.Sort(got)

			want := slices.Clone(tc.modelIDs)
			slices.Sort(want)

			if !slices.Equal(got, want) {
				t.Fatalf(
					"provider %q model membership mismatch\n got: %#v\nwant: %#v",
					tc.name,
					got,
					want,
				)
			}
		})
	}
}

func TestCatalogPublicLookupMethodsCoverEveryProviderAndModel(t *testing.T) {
	catalog := DefaultCatalog()

	for _, providerName := range ProviderNames() {
		t.Run(string(providerName), func(t *testing.T) {
			wantProvider := catalog.Providers[providerName]

			gotProvider, err := Provider(providerName)
			if err != nil {
				t.Fatalf("Provider(%q): %v", providerName, err)
			}
			if !reflect.DeepEqual(gotProvider, wantProvider) {
				t.Fatalf(
					"Provider(%q) mismatch\n got: %#v\nwant: %#v",
					providerName,
					gotProvider,
					wantProvider,
				)
			}

			gotIDs, err := ModelPresetIDs(providerName)
			if err != nil {
				t.Fatalf("ModelPresetIDs(%q): %v", providerName, err)
			}

			wantIDs := make([]ModelPresetID, 0, len(wantProvider.ModelPresets))
			for modelID := range wantProvider.ModelPresets {
				wantIDs = append(wantIDs, modelID)
			}
			slices.Sort(wantIDs)

			if !slices.Equal(gotIDs, wantIDs) {
				t.Fatalf(
					"ModelPresetIDs(%q) mismatch\n got: %#v\nwant: %#v",
					providerName,
					gotIDs,
					wantIDs,
				)
			}

			for _, modelID := range gotIDs {
				wantModel := wantProvider.ModelPresets[modelID]

				gotModel, err := Model(providerName, modelID)
				if err != nil {
					t.Fatalf("Model(%q, %q): %v", providerName, modelID, err)
				}
				if !reflect.DeepEqual(gotModel, wantModel) {
					t.Fatalf(
						"Model(%q, %q) mismatch\n got: %#v\nwant: %#v",
						providerName,
						modelID,
						gotModel,
						wantModel,
					)
				}
			}
		})
	}
}

func TestCatalogPublicLookupErrors(t *testing.T) {
	tests := []struct {
		name string
		call func() error
		want error
	}{
		{
			name: "Provider reports ErrProviderNotFound",
			call: func() error {
				_, err := Provider("missing")
				return err
			},
			want: ErrProviderNotFound,
		},
		{
			name: "Model reports ErrProviderNotFound",
			call: func() error {
				_, err := Model("missing", "missing")
				return err
			},
			want: ErrProviderNotFound,
		},
		{
			name: "Model reports ErrModelNotFound",
			call: func() error {
				_, err := Model(ProviderAnthropic, "missing")
				return err
			},
			want: ErrModelNotFound,
		},
		{
			name: "ModelPresetIDs reports ErrProviderNotFound",
			call: func() error {
				_, err := ModelPresetIDs("missing")
				return err
			},
			want: ErrProviderNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if !errors.Is(err, tc.want) {
				t.Fatalf("error got %v want errors.Is(..., %v)", err, tc.want)
			}
		})
	}
}

func TestCatalogReturnsIndependentCopies(t *testing.T) {
	providerBaseline, err := Provider(ProviderAnthropic)
	if err != nil {
		t.Fatalf("Provider baseline: %v", err)
	}
	provider, err := Provider(ProviderAnthropic)
	if err != nil {
		t.Fatalf("Provider mutable copy: %v", err)
	}

	if provider.DefaultHeaders == nil {
		t.Fatal("Anthropic default headers unexpectedly nil")
	}
	if provider.CapabilitiesOverride == nil || len(provider.CapabilitiesOverride.ModalitiesIn) == 0 {
		t.Fatal("Anthropic capabilities override unexpectedly incomplete")
	}

	provider.DefaultHeaders["x-mutated"] = "true"
	provider.CapabilitiesOverride.ModalitiesIn[0] = spec.ModalityAudioIn

	embeddedModel := provider.ModelPresets[PresetClaudeSonnet5]
	if embeddedModel.ModelParam.Temperature == nil ||
		embeddedModel.ModelParam.Reasoning == nil ||
		embeddedModel.ModelParam.CacheControl == nil {
		t.Fatal("Anthropic Sonnet 5 preset unexpectedly incomplete")
	}
	*embeddedModel.ModelParam.Temperature = 0.42
	embeddedModel.ModelParam.Reasoning.Level = spec.ReasoningLevelNone
	embeddedModel.ModelParam.CacheControl.TTL = spec.CacheControlTTL1h
	provider.ModelPresets[PresetClaudeSonnet5] = embeddedModel

	freshProvider, err := Provider(ProviderAnthropic)
	if err != nil {
		t.Fatalf("Provider fresh copy: %v", err)
	}
	if !reflect.DeepEqual(freshProvider, providerBaseline) {
		t.Fatalf(
			"Provider returned shared mutable state\n got: %#v\nwant: %#v",
			freshProvider,
			providerBaseline,
		)
	}

	modelBaseline, err := Model(ProviderAnthropic, PresetClaudeSonnet5)
	if err != nil {
		t.Fatalf("Model baseline: %v", err)
	}
	model, err := Model(ProviderAnthropic, PresetClaudeSonnet5)
	if err != nil {
		t.Fatalf("Model mutable copy: %v", err)
	}

	*model.ModelParam.Temperature = 0.73
	model.ModelParam.Reasoning.Level = spec.ReasoningLevelLow
	model.ModelParam.CacheControl.TTL = spec.CacheControlTTL1h

	freshModel, err := Model(ProviderAnthropic, PresetClaudeSonnet5)
	if err != nil {
		t.Fatalf("Model fresh copy: %v", err)
	}
	if !reflect.DeepEqual(freshModel, modelBaseline) {
		t.Fatalf(
			"Model returned shared mutable state\n got: %#v\nwant: %#v",
			freshModel,
			modelBaseline,
		)
	}
}

func TestCloneModelPresetClonesReasoningControlPointers(t *testing.T) {
	summaryStyle := spec.ReasoningSummaryStyleOmitted
	context := spec.ReasoningContextAllTurns
	mode := spec.ReasoningModePro

	model := ModelPreset{
		ID:          "test-model",
		Name:        "test-model",
		DisplayName: "Test Model",
		ModelParam: spec.ModelParam{
			Name: "test-model",
			Reasoning: &spec.ReasoningParam{
				Type:         spec.ReasoningTypeSingleWithLevels,
				Level:        spec.ReasoningLevelHigh,
				SummaryStyle: &summaryStyle,
				Context:      &context,
				Mode:         &mode,
			},
		},
	}

	cloned := CloneModelPreset(model)
	if cloned.ModelParam.Reasoning == nil {
		t.Fatal("expected cloned reasoning parameter")
	}
	if cloned.ModelParam.Reasoning.SummaryStyle == model.ModelParam.Reasoning.SummaryStyle {
		t.Fatal("summaryStyle pointer was not cloned")
	}
	if cloned.ModelParam.Reasoning.Context == model.ModelParam.Reasoning.Context {
		t.Fatal("context pointer was not cloned")
	}
	if cloned.ModelParam.Reasoning.Mode == model.ModelParam.Reasoning.Mode {
		t.Fatal("mode pointer was not cloned")
	}

	*model.ModelParam.Reasoning.SummaryStyle = spec.ReasoningSummaryStyleDetailed
	*model.ModelParam.Reasoning.Context = spec.ReasoningContextCurrentTurn
	*model.ModelParam.Reasoning.Mode = spec.ReasoningModeStandard

	if got := *cloned.ModelParam.Reasoning.SummaryStyle; got != spec.ReasoningSummaryStyleOmitted {
		t.Fatalf("summaryStyle got %q want omitted", got)
	}
	if got := *cloned.ModelParam.Reasoning.Context; got != spec.ReasoningContextAllTurns {
		t.Fatalf("context got %q want all_turns", got)
	}
	if got := *cloned.ModelParam.Reasoning.Mode; got != spec.ReasoningModePro {
		t.Fatalf("mode got %q want pro", got)
	}
}

func TestValidateCatalogReasoningControls(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Catalog)
		wantErr string
	}{
		{
			name: "valid omitted summary context and mode",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					summaryStyle := spec.ReasoningSummaryStyleOmitted
					context := spec.ReasoningContextAllTurns
					mode := spec.ReasoningModePro
					model.ModelParam.Reasoning = &spec.ReasoningParam{
						Type:         spec.ReasoningTypeSingleWithLevels,
						Level:        spec.ReasoningLevelLow,
						SummaryStyle: &summaryStyle,
						Context:      &context,
						Mode:         &mode,
					}
				})
			},
		},
		{
			name: "invalid reasoning context",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					context := spec.ReasoningContext("invalid")
					model.ModelParam.Reasoning = &spec.ReasoningParam{
						Type:    spec.ReasoningTypeSingleWithLevels,
						Level:   spec.ReasoningLevelLow,
						Context: &context,
					}
				})
			},
			wantErr: "unknown context",
		},
		{
			name: "invalid reasoning mode",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					mode := spec.ReasoningMode("invalid")
					model.ModelParam.Reasoning = &spec.ReasoningParam{
						Type:  spec.ReasoningTypeSingleWithLevels,
						Level: spec.ReasoningLevelLow,
						Mode:  &mode,
					}
				})
			},
			wantErr: "unknown mode",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			catalog := CloneCatalog(catalogValidationTestCatalog())
			tc.mutate(&catalog)

			err := ValidateCatalog(catalog)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateCatalog: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error got %v want substring %q", err, tc.wantErr)
			}
		})
	}
}

func TestValidateCatalogRejectsInvalidData(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Catalog)
		wantErr string
	}{
		{
			name: "provider map key must match provider name",
			mutate: func(catalog *Catalog) {
				provider := catalog.Providers["test"]
				provider.Name = "other"
				catalog.Providers["test"] = provider
			},
			wantErr: "does not match provider.name",
		},
		{
			name: "provider path prefix is required",
			mutate: func(catalog *Catalog) {
				provider := catalog.Providers["test"]
				provider.ChatCompletionPathPrefix = ""
				catalog.Providers["test"] = provider
			},
			wantErr: "chatCompletionPathPrefix is empty",
		},
		{
			name: "model requires sampling default",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					model.ModelParam.Temperature = nil
					model.ModelParam.Reasoning = nil
				})
			},
			wantErr: "either reasoning or temperature must be set",
		},
		{
			name: "model name must match model parameter name",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					model.ModelParam.Name = "different-model"
				})
			},
			wantErr: "does not match modelParam.name",
		},
		{
			name: "hybrid reasoning requires positive token budget",
			mutate: func(catalog *Catalog) {
				mutateCatalogValidationTestModel(catalog, func(model *ModelPreset) {
					model.ModelParam.Temperature = nil
					model.ModelParam.Reasoning = &spec.ReasoningParam{
						Type:   spec.ReasoningTypeHybridWithTokens,
						Tokens: 0,
					}
				})
			},
			wantErr: "tokens must be > 0",
		},
		{
			name: "capability override values are validated",
			mutate: func(catalog *Catalog) {
				provider := catalog.Providers["test"]
				provider.CapabilitiesOverride = &capabilityoverride.ModelCapabilitiesOverride{
					ModalitiesIn: []spec.Modality{"invalid"},
				}
				catalog.Providers["test"] = provider
			},
			wantErr: "invalid input modality",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			catalog := CloneCatalog(catalogValidationTestCatalog())
			tc.mutate(&catalog)

			err := ValidateCatalog(catalog)
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf(
					"error got %q, want substring %q",
					err.Error(),
					tc.wantErr,
				)
			}
		})
	}
}

func TestZAIPayAsYouGoUsesDocumentedChatCompletionsTransport(t *testing.T) {
	tests := []struct {
		name    spec.ProviderName
		sdkType spec.ProviderSDKType
		origin  string
		path    string
	}{
		{
			name:    ProviderZAI,
			sdkType: spec.ProviderSDKTypeOpenAIChatCompletions,
			origin:  "https://api.z.ai",
			path:    "/api/paas/v4/chat/completions",
		},
	}

	for _, tc := range tests {
		t.Run(string(tc.name), func(t *testing.T) {
			provider, err := Provider(tc.name)
			if err != nil {
				t.Fatalf("Provider(%q): %v", tc.name, err)
			}
			if provider.SDKType != tc.sdkType {
				t.Fatalf("sdkType got %q want %q", provider.SDKType, tc.sdkType)
			}
			if provider.Origin != tc.origin {
				t.Fatalf("origin got %q want %q", provider.Origin, tc.origin)
			}
			if provider.ChatCompletionPathPrefix != tc.path {
				t.Fatalf(
					"chatCompletionPathPrefix got %q want %q",
					provider.ChatCompletionPathPrefix,
					tc.path,
				)
			}
		})
	}
}

func catalogValidationTestCatalog() Catalog {
	temperature := 1.0
	const providerName spec.ProviderName = "test"
	const modelID ModelPresetID = "model"
	const modelName spec.ModelName = "test-model"

	return Catalog{
		Providers: map[spec.ProviderName]ProviderPreset{
			providerName: {
				Name:                     providerName,
				DisplayName:              "Test Provider",
				SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
				Origin:                   "https://example.test",
				ChatCompletionPathPrefix: "/v1/responses",
				APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
				DefaultHeaders: map[string]string{
					spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
				},
				ModelPresets: map[ModelPresetID]ModelPreset{
					modelID: {
						ID:          modelID,
						Name:        modelName,
						DisplayName: "Test Model",
						ModelParam: spec.ModelParam{
							Name:        modelName,
							Temperature: &temperature,
						},
					},
				},
			},
		},
	}
}

func mutateCatalogValidationTestModel(
	catalog *Catalog,
	mutate func(*ModelPreset),
) {
	provider := catalog.Providers["test"]
	model := provider.ModelPresets["model"]
	mutate(&model)
	provider.ModelPresets["model"] = model
	catalog.Providers["test"] = provider
}
