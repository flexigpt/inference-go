package modelpreset

import (
	"errors"
	"maps"
	"slices"

	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var (
	ErrProviderNotFound = errors.New("provider preset not found")
	ErrModelNotFound    = errors.New("model preset not found")
)

type ModelPresetID string

type ModelPreset struct {
	ID          ModelPresetID  `json:"id"`
	Name        spec.ModelName `json:"name"`
	DisplayName string         `json:"displayName"`

	// ModelParam is the default runtime request model configuration.
	// Callers should treat values returned by this package as immutable.
	ModelParam spec.ModelParam `json:"modelParam"`

	// CapabilitiesOverride is a runtime capability patch applied over provider/base SDK capabilities.
	// It is not the derived/effective capability profile.
	CapabilitiesOverride *capabilityoverride.ModelCapabilitiesOverride `json:"capabilitiesOverride,omitempty"`
}

type ProviderPreset struct {
	Name        spec.ProviderName    `json:"name"`
	DisplayName string               `json:"displayName"`
	SDKType     spec.ProviderSDKType `json:"sdkType"`

	Origin                   string            `json:"origin"`
	ChatCompletionPathPrefix string            `json:"chatCompletionPathPrefix"`
	APIKeyHeaderKey          string            `json:"apiKeyHeaderKey"`
	DefaultHeaders           map[string]string `json:"defaultHeaders,omitempty"`

	// CapabilitiesOverride is a provider-wide runtime capability patch.
	// Model preset overrides are applied after this.
	CapabilitiesOverride *capabilityoverride.ModelCapabilitiesOverride `json:"capabilitiesOverride,omitempty"`

	ModelPresets map[ModelPresetID]ModelPreset `json:"modelPresets"`
}

type Catalog struct {
	Providers map[spec.ProviderName]ProviderPreset `json:"providers"`
}

var catalogProviders = map[spec.ProviderName]ProviderPreset{
	ProviderAnthropic:       providerAnthropic,
	ProviderDeepSeek:        providerDeepSeek,
	ProviderLocalAI:         providerLocalAI,
	ProviderLMStudio:        providerLMStudio,
	ProviderGoogleGemini:    providerGoogleGemini,
	ProviderHuggingFace:     providerHuggingFace,
	ProviderLlamaCPP:        providerLlamaCPP,
	ProviderMeta:            providerMeta,
	ProviderMiniMax:         providerMiniMax,
	ProviderMistral:         providerMistral,
	ProviderOllama:          providerOllama,
	ProviderOpenAIChat:      providerOpenAIChat,
	ProviderOpenAIResponses: providerOpenAIResponses,
	ProviderOpenRouter:      providerOpenRouter,
	ProviderSGLang:          providerSGLang,
	ProviderVLLM:            providerVLLM,
	ProviderXAI:             providerXAI,
	ProviderXiaomiMiMo:      providerXiaomiMiMo,
}

func DefaultCatalog() Catalog {
	return CloneCatalog(Catalog{
		Providers: catalogProviders,
	})
}

func ProviderNames() []spec.ProviderName {
	names := make([]spec.ProviderName, 0, len(catalogProviders))
	for name := range catalogProviders {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

func Provider(name spec.ProviderName) (ProviderPreset, error) {
	pp, ok := catalogProviders[name]
	if !ok {
		return ProviderPreset{}, ErrProviderNotFound
	}
	return CloneProviderPreset(pp), nil
}

func Model(provider spec.ProviderName, modelPresetID ModelPresetID) (ModelPreset, error) {
	pp, ok := catalogProviders[provider]
	if !ok {
		return ModelPreset{}, ErrProviderNotFound
	}
	mp, ok := pp.ModelPresets[modelPresetID]
	if !ok {
		return ModelPreset{}, ErrModelNotFound
	}
	return CloneModelPreset(mp), nil
}

func ModelPresetIDs(provider spec.ProviderName) ([]ModelPresetID, error) {
	pp, ok := catalogProviders[provider]
	if !ok {
		return nil, ErrProviderNotFound
	}

	ids := make([]ModelPresetID, 0, len(pp.ModelPresets))
	for id := range pp.ModelPresets {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids, nil
}

func CloneCatalog(in Catalog) Catalog {
	out := Catalog{
		Providers: make(map[spec.ProviderName]ProviderPreset, len(in.Providers)),
	}
	for k, v := range in.Providers {
		out.Providers[k] = CloneProviderPreset(v)
	}
	return out
}

func CloneProviderPreset(in ProviderPreset) ProviderPreset {
	out := in
	out.DefaultHeaders = maps.Clone(in.DefaultHeaders)
	out.CapabilitiesOverride = capabilityoverride.CloneModelCapabilitiesOverride(in.CapabilitiesOverride)
	out.ModelPresets = make(map[ModelPresetID]ModelPreset, len(in.ModelPresets))
	for k, v := range in.ModelPresets {
		out.ModelPresets[k] = CloneModelPreset(v)
	}
	return out
}

func CloneModelPreset(in ModelPreset) ModelPreset {
	out := in
	out.ModelParam = cloneModelParam(in.ModelParam)
	out.CapabilitiesOverride = capabilityoverride.CloneModelCapabilitiesOverride(in.CapabilitiesOverride)
	return out
}

func cloneModelParam(in spec.ModelParam) spec.ModelParam {
	out := in
	out.Temperature = sdkutil.CloneFloat64Ptr(in.Temperature)
	out.Reasoning = cloneReasoningParam(in.Reasoning)
	out.CacheControl = cloneCacheControl(in.CacheControl)
	out.OutputParam = cloneOutputParam(in.OutputParam)
	out.StopSequences = slices.Clone(in.StopSequences)
	out.AdditionalParametersRawJSON = sdkutil.CloneStringPtr(in.AdditionalParametersRawJSON)
	return out
}

func cloneCacheControl(in *spec.CacheControl) *spec.CacheControl {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func cloneReasoningParam(in *spec.ReasoningParam) *spec.ReasoningParam {
	if in == nil {
		return nil
	}
	out := *in
	if in.SummaryStyle != nil {
		v := *in.SummaryStyle
		out.SummaryStyle = &v
	}
	return &out
}

func cloneOutputParam(in *spec.OutputParam) *spec.OutputParam {
	if in == nil {
		return nil
	}
	out := *in
	if in.Verbosity != nil {
		v := *in.Verbosity
		out.Verbosity = &v
	}
	if in.Format != nil {
		f := *in.Format
		if f.JSONSchemaParam != nil {
			j := *f.JSONSchemaParam
			j.Schema = maps.Clone(j.Schema)
			f.JSONSchemaParam = &j
		}
		out.Format = &f
	}
	return &out
}

func reasoningLevels(includeNone bool) []spec.ReasoningLevel {
	levels := make([]spec.ReasoningLevel, 0, 4)
	if includeNone {
		levels = append(levels, spec.ReasoningLevelNone)
	}
	levels = append(levels,
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	)
	return levels
}

func capTextOnly() *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
	}
}

func capTextImage() *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
	}
}

func capTextOnlyReasoning(
	levels []spec.ReasoningLevel,
	summaryStyle bool,
	temperatureDisallowed bool,
) *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels:         levels,
			SupportsSummaryStyle:             new(summaryStyle),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(temperatureDisallowed),
		},
	}
}

func capTextImageReasoning(
	levels []spec.ReasoningLevel,
	summaryStyle bool,
	temperatureDisallowed bool,
) *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels:         levels,
			SupportsSummaryStyle:             new(summaryStyle),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(temperatureDisallowed),
		},
	}
}

func reasoningHybrid(tokens int) *spec.ReasoningParam {
	return &spec.ReasoningParam{
		Type:   spec.ReasoningTypeHybridWithTokens,
		Tokens: tokens,
	}
}

func reasoningSingle(level spec.ReasoningLevel) *spec.ReasoningParam {
	return &spec.ReasoningParam{
		Type:  spec.ReasoningTypeSingleWithLevels,
		Level: level,
	}
}

func cacheEphemeral5m() *spec.CacheControl {
	return &spec.CacheControl{
		Kind: spec.CacheControlKindEphemeral,
		TTL:  spec.CacheControlTTL5m,
	}
}
