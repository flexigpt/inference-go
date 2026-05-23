package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetOpenRouterNVIDIANemotron3SuperFree ModelPresetID = "nvidiaNemotron3SuperFree"
	PresetOpenRouterMiniMaxM27               ModelPresetID = "minimaxm27"
	PresetOpenRouterMiniMaxM25Free           ModelPresetID = "minimaxm25free"
	PresetOpenRouterZAIGLM51                 ModelPresetID = "zaiglm51"
)

const (
	ModelNameOpenRouterNVIDIANemotron3SuperFree spec.ModelName = "nvidia/nemotron-3-super-120b-a12b:free"
	ModelNameOpenRouterMiniMaxM27               spec.ModelName = "minimax/minimax-m2.7"
	ModelNameOpenRouterMiniMaxM25Free           spec.ModelName = "minimax/minimax-m2.5:free"
	ModelNameOpenRouterZAIGLM51                 spec.ModelName = "z-ai/glm-5.1"
)

var modelOpenRouterNVIDIANemotron3SuperFree = ModelPreset{
	ID:          PresetOpenRouterNVIDIANemotron3SuperFree,
	Name:        ModelNameOpenRouterNVIDIANemotron3SuperFree,
	DisplayName: "OpenRouter NVIDIA Nemotron 3 Super Free",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterNVIDIANemotron3SuperFree,
		Stream:          true,
		MaxPromptLength: 32768,
		MaxOutputLength: 32768,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelOpenRouterMiniMaxM27 = ModelPreset{
	ID:          PresetOpenRouterMiniMaxM27,
	Name:        ModelNameOpenRouterMiniMaxM27,
	DisplayName: "OpenRouter MiniMax M2.7",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMiniMaxM27,
		Stream:          true,
		MaxPromptLength: 180000,
		MaxOutputLength: 125000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openRouterReasoningOverride(),
}

var modelOpenRouterMiniMaxM25Free = ModelPreset{
	ID:          PresetOpenRouterMiniMaxM25Free,
	Name:        ModelNameOpenRouterMiniMaxM25Free,
	DisplayName: "OpenRouter MiniMax M2.5 free",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMiniMaxM25Free,
		Stream:          true,
		MaxPromptLength: 100000,
		MaxOutputLength: 8000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openRouterReasoningOverride(),
}

var modelOpenRouterZAIGLM51 = ModelPreset{
	ID:          PresetOpenRouterZAIGLM51,
	Name:        ModelNameOpenRouterZAIGLM51,
	DisplayName: "OpenRouter Z.AI GLM5.1",
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterZAIGLM51,
		Stream:          true,
		MaxPromptLength: 180000,
		MaxOutputLength: 125000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: openRouterReasoningOverride(),
}

func openRouterReasoningOverride() *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
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
			SupportsSummaryStyle: new(true),
		},
	}
}

var providerOpenRouter = ProviderPreset{
	Name:                     ProviderOpenRouter,
	DisplayName:              "OpenRouter",
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://openrouter.ai",
	ChatCompletionPathPrefix: "/api" + spec.DefaultOpenAIResponsesPrefix,
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
			SupportsSummaryStyle:             new(true),
			SupportsEncryptedReasoningInput:  new(true),
			TemperatureDisallowedWhenEnabled: new(false),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(false),
			DisallowedWithReasoning: new(false),
			MaxSequences:            new(0),
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
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetOpenRouterNVIDIANemotron3SuperFree: modelOpenRouterNVIDIANemotron3SuperFree,
		PresetOpenRouterMiniMaxM27:               modelOpenRouterMiniMaxM27,
		PresetOpenRouterMiniMaxM25Free:           modelOpenRouterMiniMaxM25Free,
		PresetOpenRouterZAIGLM51:                 modelOpenRouterZAIGLM51,
	},
}
