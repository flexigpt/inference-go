package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetLlamaCPPBehemoth ModelPresetID = "llama4Behemoth"
	PresetLlamaCPPMaverick ModelPresetID = "llama4Maverick"
	PresetLlamaCPPScout    ModelPresetID = "llama4Scout"
)

const (
	ModelNameLlamaCPPBehemoth spec.ModelName = "llama4-behemoth"
	ModelNameLlamaCPPMaverick spec.ModelName = "llama4-maverick"
	ModelNameLlamaCPPScout    spec.ModelName = "llama4-scout"
)

var modelLlamaCPPBehemoth = ModelPreset{
	ID:          PresetLlamaCPPBehemoth,
	Name:        ModelNameLlamaCPPBehemoth,
	DisplayName: "LLama 4 Behemoth",
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlamaCPPBehemoth,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelLlamaCPPMaverick = ModelPreset{
	ID:          PresetLlamaCPPMaverick,
	Name:        ModelNameLlamaCPPMaverick,
	DisplayName: "LLama 4 Maverick",
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlamaCPPMaverick,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelLlamaCPPScout = ModelPreset{
	ID:          PresetLlamaCPPScout,
	Name:        ModelNameLlamaCPPScout,
	DisplayName: "LLama 4 Scout",
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlamaCPPScout,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerLlamaCPP = ProviderPreset{
	Name:                     ProviderLlamaCPP,
	DisplayName:              "llama.cpp",
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   "http://127.0.0.1:8080",
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
		PresetLlamaCPPBehemoth: modelLlamaCPPBehemoth,
		PresetLlamaCPPMaverick: modelLlamaCPPMaverick,
		PresetLlamaCPPScout:    modelLlamaCPPScout,
	},
}
