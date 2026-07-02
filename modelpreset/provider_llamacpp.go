package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelLlamaCPPBehemoth = ModelPreset{
	ID:          PresetLlama4Behemoth,
	Name:        ModelNameLlama4BehemothLocal,
	DisplayName: DisplayNameLlama4Behemoth,
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlama4BehemothLocal,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelLlamaCPPMaverick = ModelPreset{
	ID:          PresetLlama4Maverick,
	Name:        ModelNameLlama4MaverickLocal,
	DisplayName: DisplayNameLlama4Maverick,
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlama4MaverickLocal,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelLlamaCPPScout = ModelPreset{
	ID:          PresetLlama4Scout,
	Name:        ModelNameLlama4ScoutLocal,
	DisplayName: DisplayNameLlama4Scout,
	ModelParam: spec.ModelParam{
		Name:            ModelNameLlama4ScoutLocal,
		Stream:          true,
		MaxPromptLength: 4096,
		MaxOutputLength: 4096,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelLlamaCPPQwen3635BA3B = ModelPreset{
	ID:          PresetQwen3635BA3B,
	Name:        ModelNameQwen3635BA3BLocal,
	DisplayName: DisplayNameQwen3635BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3635BA3BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(0.1),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerLlamaCPP = ProviderPreset{
	Name:                     ProviderLlamaCPP,
	DisplayName:              DisplayNameProviderLlamaCPP,
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
		PresetLlama4Behemoth: modelLlamaCPPBehemoth,
		PresetLlama4Maverick: modelLlamaCPPMaverick,
		PresetLlama4Scout:    modelLlamaCPPScout,
		PresetQwen3635BA3B:   modelLlamaCPPQwen3635BA3B,
	},
}
