package capabilityoverride

import "github.com/flexigpt/inference-go/spec"

type ReasoningTokenBudgetCapabilitiesOverride struct {
	MinAllowed      *int  `json:"minAllowed,omitempty"`
	MaxAllowed      *int  `json:"maxAllowed,omitempty"`
	ZeroAllowed     *bool `json:"zeroAllowed,omitempty"`
	MinusOneAllowed *bool `json:"minusOneAllowed,omitempty"`
}

type ReasoningCapabilitiesOverride struct {
	SupportsReasoningConfig *bool `json:"supportsReasoningConfig,omitempty"`

	SupportedReasoningTypes       []spec.ReasoningType                      `json:"supportedReasoningTypes,omitempty"`
	SupportedReasoningLevels      []spec.ReasoningLevel                     `json:"supportedReasoningLevels,omitempty"`
	HybridTokenBudgetCapabilities *ReasoningTokenBudgetCapabilitiesOverride `json:"hybridTokenBudgetCapabilities,omitempty"`

	SupportsSummaryStyle             *bool `json:"supportsSummaryStyle,omitempty"`
	SupportsEncryptedReasoningInput  *bool `json:"supportsEncryptedReasoningInput,omitempty"`
	TemperatureDisallowedWhenEnabled *bool `json:"temperatureDisallowedWhenEnabled,omitempty"`
}

type StopSequenceCapabilitiesOverride struct {
	IsSupported             *bool `json:"isSupported,omitempty"`
	DisallowedWithReasoning *bool `json:"disallowedWithReasoning,omitempty"`
	MaxSequences            *int  `json:"maxSequences,omitempty"`
}

type OutputCapabilitiesOverride struct {
	SupportedOutputFormats []spec.OutputFormatKind `json:"supportedOutputFormats,omitempty"`
	SupportsVerbosity      *bool                   `json:"supportsVerbosity,omitempty"`
}

type ToolCapabilitiesOverride struct {
	SupportedToolTypes               []spec.ToolType             `json:"supportedToolTypes,omitempty"`
	SupportedToolPolicyModes         []spec.ToolPolicyMode       `json:"supportedToolPolicyModes,omitempty"`
	SupportsParallelToolCalls        *bool                       `json:"supportsParallelToolCalls,omitempty"`
	MaxForcedTools                   *int                        `json:"maxForcedTools,omitempty"`
	SupportedClientToolOutputFormats []spec.ToolOutputFormatKind `json:"supportedClientToolOutputFormats,omitempty"`
}

type CacheControlCapabilitiesOverride struct {
	SupportsTTL    *bool                   `json:"supportsTTL,omitempty"`
	SupportedKinds []spec.CacheControlKind `json:"supportedKinds,omitempty"`
	SupportedTTLs  []spec.CacheControlTTL  `json:"supportedTTLs,omitempty"`
	SupportsKey    *bool                   `json:"supportsKey,omitempty"`
}

type CacheCapabilitiesOverride struct {
	SupportsAutomaticCaching *bool                             `json:"supportsAutomaticCaching,omitempty"`
	TopLevel                 *CacheControlCapabilitiesOverride `json:"topLevel,omitempty"`
	InputOutputContent       *CacheControlCapabilitiesOverride `json:"inputOutputContent,omitempty"`
	ReasoningContent         *CacheControlCapabilitiesOverride `json:"reasoningContent,omitempty"`
	ToolChoice               *CacheControlCapabilitiesOverride `json:"toolChoice,omitempty"`
	ToolCall                 *CacheControlCapabilitiesOverride `json:"toolCall,omitempty"`
	ToolOutput               *CacheControlCapabilitiesOverride `json:"toolOutput,omitempty"`
}

type ParamDialectOverride struct {
	MaxOutputTokensParamName *spec.MaxOutputTokensParamName `json:"maxOutputTokensParamName,omitempty"`
	ToolChoiceParamStyle     *spec.ToolChoiceParamStyle     `json:"toolChoiceParamStyle,omitempty"`
}

// ModelCapabilitiesOverride is a patch-like version of inference-go's ModelCapabilities.
//
// Semantics:
//   - nil slice means no override provided.
//   - empty slice means override to empty.
//   - pointer scalar nil means not provided.
//   - pointer scalar non-nil means override.
//
// This struct is intended for storage and API transport as an override only.
// The effective/derived capabilities should be computed at runtime and should not be stored.
//
// An override is not a complete capability object. It is a partial patch.
// Fields may be supplied independently and then layered with provider/base/model capabilities.
// Structural validation should validate: enum values, duplicates, non-negative numbers, local numeric consistency.
// It should not require that one partial override is internally complete.
type ModelCapabilitiesOverride struct {
	ModalitiesIn  []spec.Modality `json:"modalitiesIn,omitempty"`
	ModalitiesOut []spec.Modality `json:"modalitiesOut,omitempty"`

	ReasoningCapabilities    *ReasoningCapabilitiesOverride    `json:"reasoningCapabilities,omitempty"`
	StopSequenceCapabilities *StopSequenceCapabilitiesOverride `json:"stopSequenceCapabilities,omitempty"`
	OutputCapabilities       *OutputCapabilitiesOverride       `json:"outputCapabilities,omitempty"`
	ToolCapabilities         *ToolCapabilitiesOverride         `json:"toolCapabilities,omitempty"`
	CacheCapabilities        *CacheCapabilitiesOverride        `json:"cacheCapabilities,omitempty"`
	ParamDialect             *ParamDialectOverride             `json:"paramDialect,omitempty"`
}
