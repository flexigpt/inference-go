package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelHuggingFaceGPTOSS120 = ModelPreset{
	ID:          PresetGPTOSS120BFireworksAI,
	Name:        ModelNameGPTOSS120BFireworksAI,
	DisplayName: DisplayNameGPTOSS120BFireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPTOSS120BFireworksAI,
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
	ID:          PresetGPTOSS20BFireworksAI,
	Name:        ModelNameGPTOSS20BFireworksAI,
	DisplayName: DisplayNameGPTOSS20BFireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPTOSS20BFireworksAI,
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
	ID:          PresetQwen3Coder30BA3BFireworksAI,
	Name:        ModelNameQwen3Coder30BA3BFireworksAI,
	DisplayName: DisplayNameQwen3Coder30BA3BFireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Coder30BA3BFireworksAI,
		Stream:          true,
		MaxPromptLength: 8192,
		MaxOutputLength: 8192,
		Temperature:     new(0.7),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM52 = ModelPreset{
	ID:          PresetGLM52FireworksAI,
	Name:        ModelNameGLM52FireworksAI,
	DisplayName: DisplayNameGLM52FireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM52FireworksAI,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceDeepSeekV4Flash = ModelPreset{
	ID:          PresetDeepSeekV4FlashFireworksAI,
	Name:        ModelNameDeepSeekV4FlashFireworksAI,
	DisplayName: DisplayNameDeepSeekV4FlashFireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekV4FlashFireworksAI,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceDeepSeekV4Pro = ModelPreset{
	ID:          PresetDeepSeekV4ProFireworksAI,
	Name:        ModelNameDeepSeekV4ProFireworksAI,
	DisplayName: DisplayNameDeepSeekV4ProFireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekV4ProFireworksAI,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceOrnith1035BFP8 = ModelPreset{
	ID:          PresetOrnith1035BFP8DeepInfra,
	Name:        ModelNameOrnith1035BFP8DeepInfra,
	DisplayName: DisplayNameOrnith1035BFP8DeepInfra,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOrnith1035BFP8DeepInfra,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM52FP8 = ModelPreset{
	ID:          PresetGLM52FP8ZAI,
	Name:        ModelNameGLM52FP8ZAI,
	DisplayName: DisplayNameGLM52FP8ZAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM52FP8ZAI,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceMiMoV25Pro = ModelPreset{
	ID:          PresetMiMoV25ProDeepInfra,
	Name:        ModelNameMiMoV25ProDeepInfra,
	DisplayName: DisplayNameMiMoV25ProDeepInfra,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV25ProDeepInfra,
		Stream:          true,
		MaxPromptLength: 131272,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceNemotron3UltraNVFP4 = ModelPreset{
	ID:          PresetNemotron3UltraNVFP4FireworksAI,
	Name:        ModelNameNemotron3UltraNVFP4FireworksAI,
	DisplayName: DisplayNameNemotron3UltraNVFP4FireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameNemotron3UltraNVFP4FireworksAI,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceQwen3CoderNext = ModelPreset{
	ID:          PresetQwen3CoderNextNovita,
	Name:        ModelNameQwen3CoderNextNovita,
	DisplayName: DisplayNameQwen3CoderNextNovita,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3CoderNextNovita,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceNemotron3UltraBF16 = ModelPreset{
	ID:          PresetNemotron3UltraBF16DeepInfra,
	Name:        ModelNameNemotron3UltraBF16DeepInfra,
	DisplayName: DisplayNameNemotron3UltraBF16DeepInfra,
	ModelParam: spec.ModelParam{
		Name:            ModelNameNemotron3UltraBF16DeepInfra,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM51 = ModelPreset{
	ID:          PresetGLM51FireworksAI,
	Name:        ModelNameGLM51FireworksAI,
	DisplayName: DisplayNameGLM51FireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM51FireworksAI,
		Stream:          true,
		MaxPromptLength: 202752,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceNemotron3SuperBF16 = ModelPreset{
	ID:          PresetNemotron3SuperBF16FeatherlessAI,
	Name:        ModelNameNemotron3SuperBF16FeatherlessAI,
	DisplayName: DisplayNameNemotron3SuperBF16FeatherlessAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameNemotron3SuperBF16FeatherlessAI,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceMiniMaxM27 = ModelPreset{
	ID:          PresetMiniMaxM27FireworksAI,
	Name:        ModelNameMiniMaxM27FireworksAI,
	DisplayName: DisplayNameMiniMaxM27FireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM27FireworksAI,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceKimiK2Thinking = ModelPreset{
	ID:          PresetKimiK2ThinkingFeatherlessAI,
	Name:        ModelNameKimiK2ThinkingFeatherlessAI,
	DisplayName: DisplayNameKimiK2ThinkingFeatherlessAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameKimiK2ThinkingFeatherlessAI,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM5 = ModelPreset{
	ID:          PresetGLM5Novita,
	Name:        ModelNameGLM5Novita,
	DisplayName: DisplayNameGLM5Novita,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM5Novita,
		Stream:          true,
		MaxPromptLength: 202752,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceStep35Flash = ModelPreset{
	ID:          PresetStep35FlashFeatherlessAI,
	Name:        ModelNameStep35FlashFeatherlessAI,
	DisplayName: DisplayNameStep35FlashFeatherlessAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameStep35FlashFeatherlessAI,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceMiniMaxM25 = ModelPreset{
	ID:          PresetMiniMaxM25Novita,
	Name:        ModelNameMiniMaxM25Novita,
	DisplayName: DisplayNameMiniMaxM25Novita,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiniMaxM25Novita,
		Stream:          true,
		MaxPromptLength: 196608,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceKimiK2Instruct = ModelPreset{
	ID:          PresetKimiK2InstructNovita,
	Name:        ModelNameKimiK2InstructNovita,
	DisplayName: DisplayNameKimiK2InstructNovita,
	ModelParam: spec.ModelParam{
		Name:            ModelNameKimiK2InstructNovita,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceKimiK2Instruct0905 = ModelPreset{
	ID:          PresetKimiK2Instruct0905Novita,
	Name:        ModelNameKimiK2Instruct0905Novita,
	DisplayName: DisplayNameKimiK2Instruct0905Novita,
	ModelParam: spec.ModelParam{
		Name:            ModelNameKimiK2Instruct0905Novita,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceMiMoV2Flash = ModelPreset{
	ID:          PresetMiMoV2FlashFeatherlessAI,
	Name:        ModelNameMiMoV2FlashFeatherlessAI,
	DisplayName: DisplayNameMiMoV2FlashFeatherlessAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMiMoV2FlashFeatherlessAI,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM47 = ModelPreset{
	ID:          PresetGLM47Cerebras,
	Name:        ModelNameGLM47Cerebras,
	DisplayName: DisplayNameGLM47Cerebras,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47Cerebras,
		Stream:          true,
		MaxPromptLength: 128000,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelHuggingFaceGLM51FP8 = ModelPreset{
	ID:          PresetGLM51FP8FireworksAI,
	Name:        ModelNameGLM51FP8FireworksAI,
	DisplayName: DisplayNameGLM51FP8FireworksAI,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM51FP8FireworksAI,
		Stream:          true,
		MaxPromptLength: 202752,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerHuggingFace = ProviderPreset{
	Name:                     ProviderHuggingFace,
	DisplayName:              DisplayNameProviderHuggingFace,
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
		PresetGPTOSS120BFireworksAI:           modelHuggingFaceGPTOSS120,
		PresetGPTOSS20BFireworksAI:            modelHuggingFaceGPTOSS20,
		PresetQwen3Coder30BA3BFireworksAI:     modelHuggingFaceQwen3Coder30,
		PresetGLM52FireworksAI:                modelHuggingFaceGLM52,
		PresetDeepSeekV4FlashFireworksAI:      modelHuggingFaceDeepSeekV4Flash,
		PresetDeepSeekV4ProFireworksAI:        modelHuggingFaceDeepSeekV4Pro,
		PresetOrnith1035BFP8DeepInfra:         modelHuggingFaceOrnith1035BFP8,
		PresetGLM52FP8ZAI:                     modelHuggingFaceGLM52FP8,
		PresetMiMoV25ProDeepInfra:             modelHuggingFaceMiMoV25Pro,
		PresetNemotron3UltraNVFP4FireworksAI:  modelHuggingFaceNemotron3UltraNVFP4,
		PresetQwen3CoderNextNovita:            modelHuggingFaceQwen3CoderNext,
		PresetNemotron3UltraBF16DeepInfra:     modelHuggingFaceNemotron3UltraBF16,
		PresetGLM51FireworksAI:                modelHuggingFaceGLM51,
		PresetNemotron3SuperBF16FeatherlessAI: modelHuggingFaceNemotron3SuperBF16,
		PresetMiniMaxM27FireworksAI:           modelHuggingFaceMiniMaxM27,
		PresetKimiK2ThinkingFeatherlessAI:     modelHuggingFaceKimiK2Thinking,
		PresetGLM5Novita:                      modelHuggingFaceGLM5,
		PresetStep35FlashFeatherlessAI:        modelHuggingFaceStep35Flash,
		PresetMiniMaxM25Novita:                modelHuggingFaceMiniMaxM25,
		PresetKimiK2InstructNovita:            modelHuggingFaceKimiK2Instruct,
		PresetKimiK2Instruct0905Novita:        modelHuggingFaceKimiK2Instruct0905,
		PresetMiMoV2FlashFeatherlessAI:        modelHuggingFaceMiMoV2Flash,
		PresetGLM47Cerebras:                   modelHuggingFaceGLM47,
		PresetGLM51FP8FireworksAI:             modelHuggingFaceGLM51FP8,
	},
}
