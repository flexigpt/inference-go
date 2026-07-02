package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var openAIChatNoReasoningOverride = &capabilityoverride.ModelCapabilitiesOverride{
	ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
		SupportedReasoningTypes:  []spec.ReasoningType{},
		SupportedReasoningLevels: []spec.ReasoningLevel{},
	},
}

var modelOpenAIChatGPT41 = ModelPreset{
	ID:          PresetGPT41,
	Name:        ModelNameGPT41,
	DisplayName: DisplayNameGPT41,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT41,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIChatNoReasoningOverride,
}

var modelOpenAIChatGPT41Mini = ModelPreset{
	ID:          PresetGPT41Mini,
	Name:        ModelNameGPT41Mini,
	DisplayName: DisplayNameGPT41Mini,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT41Mini,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIChatNoReasoningOverride,
}

var modelOpenAIChatGPT4o = ModelPreset{
	ID:          PresetGPT4o,
	Name:        ModelNameGPT4o,
	DisplayName: DisplayNameGPT4o,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT4o,
		Stream:          true,
		MaxPromptLength: 64000,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIChatNoReasoningOverride,
}

var modelOpenAIChatGPT4oMini = ModelPreset{
	ID:          PresetGPT4oMini,
	Name:        ModelNameGPT4oMini,
	DisplayName: DisplayNameGPT4oMini,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPT4oMini,
		Stream:          true,
		MaxPromptLength: 64000,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openAIChatNoReasoningOverride,
}

var providerOpenAIChat = ProviderPreset{
	Name:                     ProviderOpenAIChat,
	DisplayName:              DisplayNameProviderOpenAIChat,
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   spec.DefaultOpenAIOrigin,
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
		PresetGPT41:     modelOpenAIChatGPT41,
		PresetGPT41Mini: modelOpenAIChatGPT41Mini,
		PresetGPT4o:     modelOpenAIChatGPT4o,
		PresetGPT4oMini: modelOpenAIChatGPT4oMini,
	},
}
