package spec

import "context"

type Modality string

const (
	ModalityTextIn   Modality = "textIn"
	ModalityTextOut  Modality = "textOut"
	ModalityImageIn  Modality = "imageIn"
	ModalityImageOut Modality = "imageOut"
	ModalityFileIn   Modality = "fileIn"
	ModalityFileOut  Modality = "fileOut"
	ModalityAudioIn  Modality = "audioIn"
	ModalityAudioOut Modality = "audioOut"
	ModalityVideoIn  Modality = "videoIn"
	ModalityVideoOut Modality = "videoOut"
)

// ReasoningTokenBudgetCapabilities applies to ReasoningTypeHybridWithTokens.
// MinAllowed/MaxAllowed apply to positive token budgets only.
// Special values 0 and -1 are governed separately because some providers/models
// support them as distinct control values.
type ReasoningTokenBudgetCapabilities struct {
	MinAllowed      int  `json:"minAllowed,omitempty"`
	MaxAllowed      int  `json:"maxAllowed,omitempty"`
	ZeroAllowed     bool `json:"zeroAllowed,omitempty"`
	MinusOneAllowed bool `json:"minusOneAllowed,omitempty"`
}

type ReasoningCapabilities struct {
	// Top-level gate: whether request-side reasoning config is supported at all.
	// If false, ModelParam.Reasoning must not be sent.
	SupportsReasoningConfig bool `json:"supportsReasoningConfig"`

	SupportedReasoningTypes          []ReasoningType                   `json:"supportedReasoningTypes"`
	SupportedReasoningLevels         []ReasoningLevel                  `json:"supportedReasoningLevels"`
	HybridTokenBudgetCapabilities    *ReasoningTokenBudgetCapabilities `json:"hybridTokenBudgetCapabilities,omitempty"`
	SupportsSummaryStyle             bool                              `json:"supportsSummaryStyle"`
	SupportsEncryptedReasoningInput  bool                              `json:"supportsEncryptedReasoningInput"`
	TemperatureDisallowedWhenEnabled bool                              `json:"temperatureDisallowedWhenEnabled"`
}

type StopSequenceCapabilities struct {
	IsSupported             bool `json:"isSupported"`
	DisallowedWithReasoning bool `json:"disallowedWithReasoning"`
	MaxSequences            int  `json:"maxSequences"`
}

type OutputCapabilities struct {
	SupportedOutputFormats []OutputFormatKind `json:"supportedOutputFormats"`
	SupportsVerbosity      bool               `json:"supportsVerbosity"`
}

// ToolOutputFormatKind describes how caller-supplied tool outputs
// (FunctionToolOutput / CustomToolOutput) may be sent back to the model.
type ToolOutputFormatKind string

const (
	ToolOutputFormatKindString          ToolOutputFormatKind = "string"
	ToolOutputFormatKindContentItemList ToolOutputFormatKind = "contentItemList"
)

type ToolCapabilities struct {
	SupportedToolTypes               []ToolType             `json:"supportedToolTypes"`
	SupportedToolPolicyModes         []ToolPolicyMode       `json:"supportedToolPolicyModes"`
	SupportsParallelToolCalls        bool                   `json:"supportsParallelToolCalls"`
	MaxForcedTools                   int                    `json:"maxForcedTools"`
	SupportedClientToolOutputFormats []ToolOutputFormatKind `json:"supportedClientToolOutputFormats,omitempty"`
}

type CacheControlCapabilities struct {
	// Top-level gate: whether request-side TTL/retention is supported at all.
	SupportsTTL    bool               `json:"supportsTTL"`
	SupportedKinds []CacheControlKind `json:"supportedKinds,omitempty"`
	SupportedTTLs  []CacheControlTTL  `json:"supportedTTLs,omitempty"`
	SupportsKey    bool               `json:"supportsKey"`
}

type CacheCapabilities struct {
	SupportsAutomaticCaching bool `json:"supportsAutomaticCaching"`

	TopLevel           *CacheControlCapabilities `json:"topLevel,omitempty"`
	InputOutputContent *CacheControlCapabilities `json:"inputOutputContent,omitempty"`
	ReasoningContent   *CacheControlCapabilities `json:"reasoningContent,omitempty"`
	ToolChoice         *CacheControlCapabilities `json:"toolChoice,omitempty"`
	ToolCall           *CacheControlCapabilities `json:"toolCall,omitempty"`
	ToolOutput         *CacheControlCapabilities `json:"toolOutput,omitempty"`
}

type ModelCapabilities struct {
	ModalitiesIn  []Modality `json:"modalitiesIn"`
	ModalitiesOut []Modality `json:"modalitiesOut"`

	ReasoningCapabilities    *ReasoningCapabilities    `json:"reasoningCapabilities,omitempty"`
	StopSequenceCapabilities *StopSequenceCapabilities `json:"stopSequenceCapabilities,omitempty"`
	OutputCapabilities       *OutputCapabilities       `json:"outputCapabilities,omitempty"`
	ToolCapabilities         *ToolCapabilities         `json:"toolCapabilities,omitempty"`
	CacheCapabilities        *CacheCapabilities        `json:"cacheCapabilities,omitempty"`
}

type ResolveModelCapabilitiesRequest struct {
	ProviderSDKType ProviderSDKType `json:"providerSDKType"`
	ModelName       ModelName       `json:"modelName"`
	CompletionKey   string          `json:"completionKey"`
}

type ModelCapabilityResolver interface {
	ResolveModelCapabilities(ctx context.Context, req ResolveModelCapabilitiesRequest) (*ModelCapabilities, error)
}
