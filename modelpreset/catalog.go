package modelpreset

import (
	"errors"
	"slices"

	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderAnthropic       spec.ProviderName = "anthropic"
	ProviderGoogleGemini    spec.ProviderName = "googlegemini"
	ProviderHuggingFace     spec.ProviderName = "huggingface"
	ProviderLlamaCPP        spec.ProviderName = "llamacpp"
	ProviderMistral         spec.ProviderName = "mistral"
	ProviderOpenAIChat      spec.ProviderName = "openai"
	ProviderOpenAIResponses spec.ProviderName = "openairesponses"
	ProviderOpenRouter      spec.ProviderName = "openrouter"
	ProviderXAI             spec.ProviderName = "xai"
)

var (
	ErrProviderNotFound = errors.New("provider preset not found")
	ErrModelNotFound    = errors.New("model preset not found")
)

var catalogProviders = map[spec.ProviderName]ProviderPreset{
	ProviderAnthropic:       providerAnthropic,
	ProviderGoogleGemini:    providerGoogleGemini,
	ProviderHuggingFace:     providerHuggingFace,
	ProviderLlamaCPP:        providerLlamaCPP,
	ProviderMistral:         providerMistral,
	ProviderOpenAIChat:      providerOpenAIChat,
	ProviderOpenAIResponses: providerOpenAIResponses,
	ProviderOpenRouter:      providerOpenRouter,
	ProviderXAI:             providerXAI,
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
