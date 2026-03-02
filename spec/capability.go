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
	SupportedTypes                   []ReasoningType
	SupportedLevels                  []ReasoningLevel
	SupportsSummaryStyle             bool
	TemperatureDisallowedWhenEnabled bool
	SupportsEncryptedReasoningInput  bool
}

type StopSequenceCapabilities struct {
	Supported               bool
	Max                     int
	DisallowedWithReasoning bool
}

type OutputCapabilities struct {
	SupportedFormats  []OutputFormatKind
	SupportsVerbosity bool
}

type ToolCapabilities struct {
	SupportedToolTypes        []ToolType
	SupportedPolicyModes      []ToolPolicyMode
	SupportsParallelToolCalls bool
	MaxForcedTools            int
}

type ModelCapabilities struct {
	ModalitiesIn  []Modality
	ModalitiesOut []Modality

	Reasoning     *ReasoningCapabilities
	StopSequences *StopSequenceCapabilities
	Output        *OutputCapabilities
	Tools         *ToolCapabilities
}

type ResolveModelCapabilitiesRequest struct {
	SDKType       ProviderSDKType
	Model         ModelName
	CapabilityKey string
}

type ModelCapabilityResolver interface {
	ResolveModelCapabilities(ctx context.Context, req ResolveModelCapabilitiesRequest) (*ModelCapabilities, error)
}
