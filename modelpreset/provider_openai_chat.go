package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetOpenAIChatGPT41     ModelPresetID = "gpt41"
	PresetOpenAIChatGPT41Mini ModelPresetID = "gpt41Mini"
	PresetOpenAIChatGPT4o     ModelPresetID = "gpt4o"
	PresetOpenAIChatGPT4oMini ModelPresetID = "gpt4oMini"
)

const (
	ModelNameOpenAIChatGPT41     spec.ModelName = "gpt-4.1"
	ModelNameOpenAIChatGPT41Mini spec.ModelName = "gpt-4.1-mini"
	ModelNameOpenAIChatGPT4o     spec.ModelName = "gpt-4o"
	ModelNameOpenAIChatGPT4oMini spec.ModelName = "gpt-4o-mini"
)

var openAIChatNoReasoningOverride = &capabilityoverride.ModelCapabilitiesOverride{
	ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
		SupportedReasoningTypes:  []spec.ReasoningType{},
		SupportedReasoningLevels: []spec.ReasoningLevel{},
	},
}

var modelOpenAIChatGPT41 = ModelPreset{
	ID:          PresetOpenAIChatGPT41,
	Name:        ModelNameOpenAIChatGPT41,
	DisplayName: "OpenAI GPT 4.1",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIChatGPT41,
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
	ID:          PresetOpenAIChatGPT41Mini,
	Name:        ModelNameOpenAIChatGPT41Mini,
	DisplayName: "OpenAI GPT 4.1 Mini",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIChatGPT41Mini,
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
	ID:          PresetOpenAIChatGPT4o,
	Name:        ModelNameOpenAIChatGPT4o,
	DisplayName: "OpenAI GPT 4o",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIChatGPT4o,
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
	ID:          PresetOpenAIChatGPT4oMini,
	Name:        ModelNameOpenAIChatGPT4oMini,
	DisplayName: "OpenAI GPT 4o Mini",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenAIChatGPT4oMini,
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
	DisplayName:              "OpenAI Chat Completions API",
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   spec.DefaultOpenAIOrigin,
	ChatCompletionPathPrefix: spec.DefaultOpenAIChatCompletionsPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(sdkutil.CloneStringMap(spec.DefaultBaseHeaders)),
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
		PresetOpenAIChatGPT41:     modelOpenAIChatGPT41,
		PresetOpenAIChatGPT41Mini: modelOpenAIChatGPT41Mini,
		PresetOpenAIChatGPT4o:     modelOpenAIChatGPT4o,
		PresetOpenAIChatGPT4oMini: modelOpenAIChatGPT4oMini,
	},
}
