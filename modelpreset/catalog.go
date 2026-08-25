package modelpreset

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"math"
	"net/url"
	"slices"
	"strings"
	"unicode"

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
	ProviderMoonshot:        providerMoonshot,
	ProviderOllama:          providerOllama,
	ProviderOpenAIChat:      providerOpenAIChat,
	ProviderOpenAIResponses: providerOpenAIResponses,
	ProviderOpenRouter:      providerOpenRouter,
	ProviderQwen:            providerQwen,
	ProviderSGLang:          providerSGLang,
	ProviderVLLM:            providerVLLM,
	ProviderXAI:             providerXAI,
	ProviderXiaomi:          providerXiaomi,
	ProviderZAI:             providerZAI,
	ProviderZAICodingPlan:   providerZAICodingPlan,
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

// ValidateCatalog validates the complete preset catalog without requiring API
// keys, network access, or adapter initialization.
//
// It intentionally validates catalog data rather than remote-provider behavior.
// Remote provider compatibility remains covered by integration tests and normal
// request normalization.
func ValidateCatalog(c Catalog) error {
	if len(c.Providers) == 0 {
		return errors.New("catalog has no providers")
	}

	providerNames := make([]spec.ProviderName, 0, len(c.Providers))
	for name := range c.Providers {
		providerNames = append(providerNames, name)
	}
	slices.Sort(providerNames)

	var errs []error
	for _, providerName := range providerNames {
		if err := validateCatalogProviderPreset(providerName, c.Providers[providerName]); err != nil {
			errs = append(errs, fmt.Errorf("provider %q: %w", providerName, err))
		}
	}

	return errors.Join(errs...)
}

func validateCatalogProviderPreset(
	mapKey spec.ProviderName,
	provider ProviderPreset,
) error {
	var errs []error

	if err := validateCatalogIdentifier("catalog map key", string(mapKey)); err != nil {
		errs = append(errs, err)
	}
	if err := validateCatalogIdentifier("name", string(provider.Name)); err != nil {
		errs = append(errs, err)
	}
	if mapKey != provider.Name {
		errs = append(errs, fmt.Errorf(
			"catalog map key %q does not match provider.name %q",
			mapKey,
			provider.Name,
		))
	}
	if strings.TrimSpace(provider.DisplayName) == "" {
		errs = append(errs, errors.New("displayName is empty"))
	}
	if !isCatalogProviderSDKType(provider.SDKType) {
		errs = append(errs, fmt.Errorf("unsupported sdkType %q", provider.SDKType))
	}
	if err := validateCatalogOrigin(provider.Origin); err != nil {
		errs = append(errs, err)
	}
	if err := validateCatalogPathPrefix(provider.ChatCompletionPathPrefix); err != nil {
		errs = append(errs, err)
	}
	if strings.TrimSpace(provider.APIKeyHeaderKey) == "" {
		errs = append(errs, errors.New("apiKeyHeaderKey is empty"))
	}
	if err := validateCatalogHeaders(provider.DefaultHeaders); err != nil {
		errs = append(errs, err)
	}
	if err := capabilityoverride.ValidateModelCapabilitiesOverride(provider.CapabilitiesOverride); err != nil {
		errs = append(errs, fmt.Errorf("capabilitiesOverride: %w", err))
	}

	if len(provider.ModelPresets) == 0 {
		errs = append(errs, errors.New("modelPresets is empty"))
		return errors.Join(errs...)
	}

	modelIDs := make([]ModelPresetID, 0, len(provider.ModelPresets))
	for id := range provider.ModelPresets {
		modelIDs = append(modelIDs, id)
	}
	slices.Sort(modelIDs)

	for _, modelID := range modelIDs {
		if err := validateCatalogModelPreset(modelID, provider.ModelPresets[modelID]); err != nil {
			errs = append(errs, fmt.Errorf("model %q: %w", modelID, err))
		}
	}

	return errors.Join(errs...)
}

func validateCatalogModelPreset(
	mapKey ModelPresetID,
	model ModelPreset,
) error {
	var errs []error

	if err := validateCatalogIdentifier("catalog map key", string(mapKey)); err != nil {
		errs = append(errs, err)
	}
	if err := validateCatalogIdentifier("id", string(model.ID)); err != nil {
		errs = append(errs, err)
	}
	if mapKey != model.ID {
		errs = append(errs, fmt.Errorf(
			"catalog map key %q does not match model.id %q",
			mapKey,
			model.ID,
		))
	}
	if strings.TrimSpace(string(model.Name)) == "" {
		errs = append(errs, errors.New("name is empty"))
	}
	if strings.TrimSpace(model.DisplayName) == "" {
		errs = append(errs, errors.New("displayName is empty"))
	}
	if model.Name != model.ModelParam.Name {
		errs = append(errs, fmt.Errorf(
			"name %q does not match modelParam.name %q",
			model.Name,
			model.ModelParam.Name,
		))
	}
	if model.ModelParam.Temperature == nil && model.ModelParam.Reasoning == nil {
		errs = append(errs, errors.New("either reasoning or temperature must be set"))
	}
	if model.ModelParam.Temperature != nil &&
		(math.IsNaN(*model.ModelParam.Temperature) ||
			math.IsInf(*model.ModelParam.Temperature, 0)) {
		errs = append(errs, errors.New("temperature must be finite"))
	}
	if model.ModelParam.MaxPromptLength < 0 {
		errs = append(errs, errors.New("maxPromptLength must be >= 0"))
	}
	if model.ModelParam.MaxOutputLength < 0 {
		errs = append(errs, errors.New("maxOutputLength must be >= 0"))
	}
	if model.ModelParam.Timeout < 0 {
		errs = append(errs, errors.New("timeout must be >= 0"))
	}
	if err := validateCatalogReasoning(model.ModelParam.Reasoning); err != nil {
		errs = append(errs, fmt.Errorf("reasoning: %w", err))
	}
	if err := validateCatalogCacheControl(model.ModelParam.CacheControl); err != nil {
		errs = append(errs, fmt.Errorf("cacheControl: %w", err))
	}
	if err := validateCatalogOutputParam(model.ModelParam.OutputParam); err != nil {
		errs = append(errs, fmt.Errorf("outputParam: %w", err))
	}
	if err := validateCatalogStopSequences(model.ModelParam.StopSequences); err != nil {
		errs = append(errs, fmt.Errorf("stopSequences: %w", err))
	}
	if raw := model.ModelParam.AdditionalParametersRawJSON; raw != nil &&
		strings.TrimSpace(*raw) != "" &&
		!json.Valid([]byte(*raw)) {
		errs = append(errs, errors.New("additionalParametersRawJSON is not valid JSON"))
	}
	if err := capabilityoverride.ValidateModelCapabilitiesOverride(model.CapabilitiesOverride); err != nil {
		errs = append(errs, fmt.Errorf("capabilitiesOverride: %w", err))
	}

	return errors.Join(errs...)
}

func validateCatalogIdentifier(field, value string) error {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return fmt.Errorf("%s is empty", field)
	}
	if trimmed != value {
		return fmt.Errorf("%s has leading or trailing whitespace", field)
	}
	if strings.IndexFunc(value, unicode.IsSpace) >= 0 {
		return fmt.Errorf("%s contains whitespace", field)
	}
	return nil
}

func validateCatalogOrigin(origin string) error {
	trimmed := strings.TrimSpace(origin)
	if trimmed == "" {
		return errors.New("origin is empty")
	}
	if trimmed != origin {
		return errors.New("origin has leading or trailing whitespace")
	}

	u, err := url.Parse(origin)
	if err != nil {
		return fmt.Errorf("origin is invalid: %w", err)
	}
	if u.Host == "" {
		return errors.New("origin must include a host")
	}
	switch u.Scheme {
	case "http", "https":
	default:
		return fmt.Errorf("origin has unsupported scheme %q", u.Scheme)
	}
	if u.RawQuery != "" || u.Fragment != "" {
		return errors.New("origin must not include a query or fragment")
	}

	return nil
}

func validateCatalogPathPrefix(pathPrefix string) error {
	trimmed := strings.TrimSpace(pathPrefix)
	if trimmed == "" {
		return errors.New("chatCompletionPathPrefix is empty")
	}
	if trimmed != pathPrefix {
		return errors.New("chatCompletionPathPrefix has leading or trailing whitespace")
	}
	if !strings.HasPrefix(pathPrefix, "/") {
		return errors.New("chatCompletionPathPrefix must start with '/'")
	}
	if strings.ContainsAny(pathPrefix, "?#") {
		return errors.New("chatCompletionPathPrefix must not include a query or fragment")
	}
	return nil
}

func validateCatalogHeaders(headers map[string]string) error {
	if len(headers) == 0 {
		return nil
	}

	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var errs []error
	for _, key := range keys {
		if strings.TrimSpace(key) == "" {
			errs = append(errs, errors.New("defaultHeaders contains an empty header name"))
			continue
		}
		if strings.TrimSpace(key) != key {
			errs = append(errs, fmt.Errorf(
				"defaultHeaders header name %q has leading or trailing whitespace",
				key,
			))
		}
	}

	return errors.Join(errs...)
}

func validateCatalogReasoning(reasoning *spec.ReasoningParam) error {
	if reasoning == nil {
		return nil
	}

	switch reasoning.Type {
	case spec.ReasoningTypeHybridWithTokens:
		if reasoning.Tokens <= 0 {
			return errors.New("tokens must be > 0 for hybridWithTokens")
		}

	case spec.ReasoningTypeSingleWithLevels:
		switch reasoning.Level {
		case spec.ReasoningLevelNone,
			spec.ReasoningLevelMinimal,
			spec.ReasoningLevelLow,
			spec.ReasoningLevelMedium,
			spec.ReasoningLevelHigh,
			spec.ReasoningLevelXHigh,
			spec.ReasoningLevelMax:
		default:
			return fmt.Errorf("invalid level %q for singleWithLevels", reasoning.Level)
		}

	default:
		return fmt.Errorf("unknown type %q", reasoning.Type)
	}

	if reasoning.SummaryStyle == nil {
		return nil
	}

	switch *reasoning.SummaryStyle {
	case spec.ReasoningSummaryStyleAuto,
		spec.ReasoningSummaryStyleConcise,
		spec.ReasoningSummaryStyleDetailed:
		return nil
	default:
		return fmt.Errorf("unknown summaryStyle %q", *reasoning.SummaryStyle)
	}
}

func validateCatalogCacheControl(cacheControl *spec.CacheControl) error {
	if cacheControl == nil {
		return nil
	}

	switch cacheControl.Kind {
	case "", spec.CacheControlKindEphemeral:
	default:
		return fmt.Errorf("unknown kind %q", cacheControl.Kind)
	}

	switch cacheControl.TTL {
	case "",
		spec.CacheControlTTL5m,
		spec.CacheControlTTL1h,
		spec.CacheControlTTL24h,
		spec.CacheControlTTLInMemory:
	default:
		return fmt.Errorf("unknown ttl %q", cacheControl.TTL)
	}

	if cacheControl.Key != strings.TrimSpace(cacheControl.Key) {
		return errors.New("key has leading or trailing whitespace")
	}

	return nil
}

func validateCatalogOutputParam(output *spec.OutputParam) error {
	if output == nil {
		return nil
	}

	if output.Verbosity != nil {
		switch *output.Verbosity {
		case spec.OutputVerbosityLow,
			spec.OutputVerbosityMedium,
			spec.OutputVerbosityHigh,
			spec.OutputVerbosityXHigh,
			spec.OutputVerbosityMax:
		default:
			return fmt.Errorf("unknown verbosity %q", *output.Verbosity)
		}
	}

	if output.Format == nil {
		return nil
	}

	switch output.Format.Kind {
	case spec.OutputFormatKindText:
		if output.Format.JSONSchemaParam != nil {
			return errors.New("jsonSchemaParam must be nil when format.kind is text")
		}

	case spec.OutputFormatKindJSONSchema:
		if output.Format.JSONSchemaParam == nil {
			return errors.New("jsonSchemaParam is required when format.kind is jsonSchema")
		}
		if !isCatalogJSONSchemaName(output.Format.JSONSchemaParam.Name) {
			return fmt.Errorf(
				"invalid jsonSchemaParam.name %q",
				output.Format.JSONSchemaParam.Name,
			)
		}
		if output.Format.JSONSchemaParam.Schema == nil {
			return errors.New("jsonSchemaParam.schema is required")
		}

	default:
		return fmt.Errorf("unknown format.kind %q", output.Format.Kind)
	}

	return nil
}

func validateCatalogStopSequences(stopSequences []string) error {
	if len(stopSequences) > 4 {
		return fmt.Errorf("too many stop sequences: %d (max 4)", len(stopSequences))
	}

	seen := make(map[string]struct{}, len(stopSequences))
	for i, stop := range stopSequences {
		if strings.TrimSpace(stop) == "" {
			return fmt.Errorf("[%d] is empty", i)
		}
		if _, ok := seen[stop]; ok {
			return fmt.Errorf("[%d] duplicates %q", i, stop)
		}
		seen[stop] = struct{}{}
	}

	return nil
}

func isCatalogJSONSchemaName(name string) bool {
	if name == "" || len(name) > 64 {
		return false
	}

	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z':
		case r >= 'A' && r <= 'Z':
		case r >= '0' && r <= '9':
		case r == '_', r == '-':
		default:
			return false
		}
	}

	return true
}

func isCatalogProviderSDKType(sdkType spec.ProviderSDKType) bool {
	switch sdkType {
	case spec.ProviderSDKTypeAnthropic,
		spec.ProviderSDKTypeOpenAIChatCompletions,
		spec.ProviderSDKTypeOpenAIResponses,
		spec.ProviderSDKTypeGoogleGenerateContent:
		return true
	default:
		return false
	}
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
