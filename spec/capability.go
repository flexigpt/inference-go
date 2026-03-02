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
	SupportedTypes                   []ReasoningType  `json:"supportedTypes"`
	SupportedLevels                  []ReasoningLevel `json:"supportedLevels"`
	SupportsSummaryStyle             bool             `json:"supportsSummaryStyle"`
	TemperatureDisallowedWhenEnabled bool             `json:"temperatureDisallowedWhenEnabled"`
	SupportsEncryptedReasoningInput  bool             `json:"supportsEncryptedReasoningInput"`
}

type StopSequenceCapabilities struct {
	Supported               bool `json:"supported"`
	DisallowedWithReasoning bool `json:"disallowedWithReasoning"`
	Max                     int  `json:"max"`
}

type OutputCapabilities struct {
	SupportedFormats  []OutputFormatKind `json:"supportedFormats"`
	SupportsVerbosity bool               `json:"supportsVerbosity"`
}

type ToolCapabilities struct {
	SupportedToolTypes        []ToolType       `json:"supportedToolTypes"`
	SupportedPolicyModes      []ToolPolicyMode `json:"supportedPolicyModes"`
	SupportsParallelToolCalls bool             `json:"supportsParallelToolCalls"`
	MaxForcedTools            int              `json:"maxForcedTools"`
}

type ModelCapabilities struct {
	ModalitiesIn  []Modality `json:"modalitiesIn"`
	ModalitiesOut []Modality `json:"modalitiesOut"`

	Reasoning     *ReasoningCapabilities    `json:"reasoning,omitempty"`
	StopSequences *StopSequenceCapabilities `json:"stopSequences,omitempty"`
	Output        *OutputCapabilities       `json:"output,omitempty"`
	Tools         *ToolCapabilities         `json:"tools,omitempty"`
}

type ResolveModelCapabilitiesRequest struct {
	SDKType       ProviderSDKType `json:"sdkType"`
	Model         ModelName       `json:"model"`
	CapabilityKey string          `json:"capabilityKey"`
}

type ModelCapabilityResolver interface {
	ResolveModelCapabilities(ctx context.Context, req ResolveModelCapabilitiesRequest) (*ModelCapabilities, error)
}
