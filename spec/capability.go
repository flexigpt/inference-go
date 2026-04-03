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

type ReasoningCapabilities struct {
	SupportedReasoningTypes          []ReasoningType  `json:"supportedReasoningTypes"`
	SupportedReasoningLevels         []ReasoningLevel `json:"supportedReasoningLevels"`
	SupportsSummaryStyle             bool             `json:"supportsSummaryStyle"`
	SupportsEncryptedReasoningInput  bool             `json:"supportsEncryptedReasoningInput"`
	TemperatureDisallowedWhenEnabled bool             `json:"temperatureDisallowedWhenEnabled"`
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

type ToolCapabilities struct {
	SupportedToolTypes        []ToolType       `json:"supportedToolTypes"`
	SupportedToolPolicyModes  []ToolPolicyMode `json:"supportedToolPolicyModes"`
	SupportsParallelToolCalls bool             `json:"supportsParallelToolCalls"`
	MaxForcedTools            int              `json:"maxForcedTools"`
}

type CacheControlCapabilities struct {
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
