package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetHuggingFaceGPTOSS120    ModelPresetID = "hfGptOss120"
	PresetHuggingFaceGPTOSS20     ModelPresetID = "hfGptOss20"
	PresetHuggingFaceQwen3Coder30 ModelPresetID = "hfQwen3Coder30"
)

const (
	ModelNameHuggingFaceGPTOSS120    spec.ModelName = "openai/gpt-oss-120b:fireworks-ai"
	ModelNameHuggingFaceGPTOSS20     spec.ModelName = "openai/gpt-oss-20b:fireworks-ai"
	ModelNameHuggingFaceQwen3Coder30 spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct:fireworks-ai"
)

var modelHuggingFaceGPTOSS120 = ModelPreset{
	ID:          PresetHuggingFaceGPTOSS120,
	Name:        ModelNameHuggingFaceGPTOSS120,
	DisplayName: "HF GPT OSS 120B",
	ModelParam: spec.ModelParam{
		Name:            ModelNameHuggingFaceGPTOSS120,
		Stream:          true,
		MaxPromptLength: 8192,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(false),
		},
	},
}

var modelHuggingFaceGPTOSS20 = ModelPreset{
	ID:          PresetHuggingFaceGPTOSS20,
	Name:        ModelNameHuggingFaceGPTOSS20,
	DisplayName: "HF GPT OSS 20B",
	ModelParam: spec.ModelParam{
		Name:            ModelNameHuggingFaceGPTOSS20,
		Stream:          true,
		MaxPromptLength: 8192,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(false),
		},
	},
}

var modelHuggingFaceQwen3Coder30 = ModelPreset{
	ID:          PresetHuggingFaceQwen3Coder30,
	Name:        ModelNameHuggingFaceQwen3Coder30,
	DisplayName: "HF Qwen3 Coder 30B",
	ModelParam: spec.ModelParam{
		Name:            ModelNameHuggingFaceQwen3Coder30,
		Stream:          true,
		MaxPromptLength: 8192,
		MaxOutputLength: 8192,
		Temperature:     new(0.7),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerHuggingFace = ProviderPreset{
	Name:                     ProviderHuggingFace,
	DisplayName:              "Hugging Face",
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   "https://router.huggingface.co",
	ChatCompletionPathPrefix: spec.DefaultOpenAIChatCompletionsPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityFileIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(true),
			DisallowedWithReasoning: new(false),
			MaxSequences:            new(4),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(true),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
				spec.ToolTypeCustom,
				spec.ToolTypeWebSearch,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetHuggingFaceGPTOSS120:    modelHuggingFaceGPTOSS120,
		PresetHuggingFaceGPTOSS20:     modelHuggingFaceGPTOSS20,
		PresetHuggingFaceQwen3Coder30: modelHuggingFaceQwen3Coder30,
	},
}
