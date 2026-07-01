package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetHuggingFaceGPTOSS120           ModelPresetID = "hfGptOss120"
	PresetHuggingFaceGPTOSS20            ModelPresetID = "hfGptOss20"
	PresetHuggingFaceQwen3Coder30        ModelPresetID = "hfQwen3Coder30"
	PresetHuggingFaceGLM52               ModelPresetID = "hfGlm52"
	PresetHuggingFaceDeepSeekV4Flash     ModelPresetID = "hfDeepSeekV4Flash"
	PresetHuggingFaceDeepSeekV4Pro       ModelPresetID = "hfDeepSeekV4Pro"
	PresetHuggingFaceOrnith1035BFP8      ModelPresetID = "hfOrnith1035BFP8"
	PresetHuggingFaceGLM52FP8            ModelPresetID = "hfGlm52FP8"
	PresetHuggingFaceMiMoV25Pro          ModelPresetID = "hfMiMoV25Pro"
	PresetHuggingFaceNemotron3UltraNVFP4 ModelPresetID = "hfNemotron3UltraNVFP4"
	PresetHuggingFaceQwen3CoderNext      ModelPresetID = "hfQwen3CoderNext"
	PresetHuggingFaceNemotron3UltraBF16  ModelPresetID = "hfNemotron3UltraBF16"
	PresetHuggingFaceGLM51               ModelPresetID = "hfGlm51"
	PresetHuggingFaceNemotron3SuperBF16  ModelPresetID = "hfNemotron3SuperBF16"
	PresetHuggingFaceMiniMaxM27          ModelPresetID = "hfMiniMaxM27"
	PresetHuggingFaceKimiK2Thinking      ModelPresetID = "hfKimiK2Thinking"
	PresetHuggingFaceGLM5                ModelPresetID = "hfGlm5"
	PresetHuggingFaceStep35Flash         ModelPresetID = "hfStep35Flash"
	PresetHuggingFaceMiniMaxM25          ModelPresetID = "hfMiniMaxM25"
	PresetHuggingFaceKimiK2Instruct      ModelPresetID = "hfKimiK2Instruct"
	PresetHuggingFaceKimiK2Instruct0905  ModelPresetID = "hfKimiK2Instruct0905"
	PresetHuggingFaceMiMoV2Flash         ModelPresetID = "hfMiMoV2Flash"
	PresetHuggingFaceGLM47               ModelPresetID = "hfGlm47"
	PresetHuggingFaceGLM51FP8            ModelPresetID = "hfGlm51FP8"
)

const (
	ModelNameHuggingFaceGPTOSS120           spec.ModelName = "openai/gpt-oss-120b:fireworks-ai"
	ModelNameHuggingFaceGPTOSS20            spec.ModelName = "openai/gpt-oss-20b:fireworks-ai"
	ModelNameHuggingFaceQwen3Coder30        spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct:fireworks-ai"
	ModelNameHuggingFaceGLM52               spec.ModelName = "zai-org/GLM-5.2:fireworks-ai"
	ModelNameHuggingFaceDeepSeekV4Flash     spec.ModelName = "deepseek-ai/DeepSeek-V4-Flash:fireworks-ai"
	ModelNameHuggingFaceDeepSeekV4Pro       spec.ModelName = "deepseek-ai/DeepSeek-V4-Pro:fireworks-ai"
	ModelNameHuggingFaceOrnith1035BFP8      spec.ModelName = "deepreinforce-ai/Ornith-1.0-35B-FP8:deepinfra"
	ModelNameHuggingFaceGLM52FP8            spec.ModelName = "zai-org/GLM-5.2-FP8:zai-org"
	ModelNameHuggingFaceMiMoV25Pro          spec.ModelName = "XiaomiMiMo/MiMo-V2.5-Pro:deepinfra"
	ModelNameHuggingFaceNemotron3UltraNVFP4 spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4:fireworks-ai"
	ModelNameHuggingFaceQwen3CoderNext      spec.ModelName = "Qwen/Qwen3-Coder-Next:novita"
	ModelNameHuggingFaceNemotron3UltraBF16  spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16:deepinfra"
	ModelNameHuggingFaceGLM51               spec.ModelName = "zai-org/GLM-5.1:fireworks-ai"
	ModelNameHuggingFaceNemotron3SuperBF16  spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16:featherless-ai"
	ModelNameHuggingFaceMiniMaxM27          spec.ModelName = "MiniMaxAI/MiniMax-M2.7:fireworks-ai"
	ModelNameHuggingFaceKimiK2Thinking      spec.ModelName = "moonshotai/Kimi-K2-Thinking:featherless-ai"
	ModelNameHuggingFaceGLM5                spec.ModelName = "zai-org/GLM-5:novita"
	ModelNameHuggingFaceStep35Flash         spec.ModelName = "stepfun-ai/Step-3.5-Flash:featherless-ai"
	ModelNameHuggingFaceMiniMaxM25          spec.ModelName = "MiniMaxAI/MiniMax-M2.5:novita"
	ModelNameHuggingFaceKimiK2Instruct      spec.ModelName = "moonshotai/Kimi-K2-Instruct:novita"
	ModelNameHuggingFaceKimiK2Instruct0905  spec.ModelName = "moonshotai/Kimi-K2-Instruct-0905:novita"
	ModelNameHuggingFaceMiMoV2Flash         spec.ModelName = "XiaomiMiMo/MiMo-V2-Flash:featherless-ai"
	ModelNameHuggingFaceGLM47               spec.ModelName = "zai-org/GLM-4.7:cerebras"
	ModelNameHuggingFaceGLM51FP8            spec.ModelName = "zai-org/GLM-5.1-FP8:fireworks-ai"
)

const (
	huggingFaceDefaultMaxOutputLength = 8192
	huggingFaceDefaultTimeout         = 1800
)

func huggingFaceModelPreset(
	id ModelPresetID,
	name spec.ModelName,
	displayName string,
	maxPromptLength int,
	temperature float64,
) ModelPreset {
	return ModelPreset{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		ModelParam: spec.ModelParam{
			Name:            name,
			Stream:          true,
			MaxPromptLength: maxPromptLength,
			MaxOutputLength: huggingFaceDefaultMaxOutputLength,
			Temperature:     &temperature,
			SystemPrompt:    "",
			Timeout:         huggingFaceDefaultTimeout,
		},
	}
}

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

var modelHuggingFaceGLM52 = huggingFaceModelPreset(
	PresetHuggingFaceGLM52,
	ModelNameHuggingFaceGLM52,
	"HF GLM 5.2",
	1048576,
	1.0,
)

var modelHuggingFaceDeepSeekV4Flash = huggingFaceModelPreset(
	PresetHuggingFaceDeepSeekV4Flash,
	ModelNameHuggingFaceDeepSeekV4Flash,
	"HF DeepSeek V4 Flash",
	1048576,
	1.0,
)

var modelHuggingFaceDeepSeekV4Pro = huggingFaceModelPreset(
	PresetHuggingFaceDeepSeekV4Pro,
	ModelNameHuggingFaceDeepSeekV4Pro,
	"HF DeepSeek V4 Pro",
	1048576,
	1.0,
)

var modelHuggingFaceOrnith1035BFP8 = huggingFaceModelPreset(
	PresetHuggingFaceOrnith1035BFP8,
	ModelNameHuggingFaceOrnith1035BFP8,
	"HF Ornith 1.0 35B FP8",
	262144,
	1.0,
)

var modelHuggingFaceGLM52FP8 = huggingFaceModelPreset(
	PresetHuggingFaceGLM52FP8,
	ModelNameHuggingFaceGLM52FP8,
	"HF GLM 5.2 FP8",
	1048576,
	1.0,
)

var modelHuggingFaceMiMoV25Pro = huggingFaceModelPreset(
	PresetHuggingFaceMiMoV25Pro,
	ModelNameHuggingFaceMiMoV25Pro,
	"HF MiMo V2.5 Pro",
	131272,
	1.0,
)

var modelHuggingFaceNemotron3UltraNVFP4 = huggingFaceModelPreset(
	PresetHuggingFaceNemotron3UltraNVFP4,
	ModelNameHuggingFaceNemotron3UltraNVFP4,
	"HF NVIDIA Nemotron 3 Ultra 550B A55B NVFP4",
	262144,
	1.0,
)

var modelHuggingFaceQwen3CoderNext = huggingFaceModelPreset(
	PresetHuggingFaceQwen3CoderNext,
	ModelNameHuggingFaceQwen3CoderNext,
	"HF Qwen3 Coder Next",
	1048576,
	1.0,
)

var modelHuggingFaceNemotron3UltraBF16 = huggingFaceModelPreset(
	PresetHuggingFaceNemotron3UltraBF16,
	ModelNameHuggingFaceNemotron3UltraBF16,
	"HF NVIDIA Nemotron 3 Ultra 550B A55B BF16",
	262144,
	1.0,
)

var modelHuggingFaceGLM51 = huggingFaceModelPreset(
	PresetHuggingFaceGLM51,
	ModelNameHuggingFaceGLM51,
	"HF GLM 5.1",
	202752,
	1.0,
)

var modelHuggingFaceNemotron3SuperBF16 = huggingFaceModelPreset(
	PresetHuggingFaceNemotron3SuperBF16,
	ModelNameHuggingFaceNemotron3SuperBF16,
	"HF NVIDIA Nemotron 3 Super 120B A12B BF16",
	262144,
	1.0,
)

var modelHuggingFaceMiniMaxM27 = huggingFaceModelPreset(
	PresetHuggingFaceMiniMaxM27,
	ModelNameHuggingFaceMiniMaxM27,
	"HF MiniMax M2.7",
	204800,
	1.0,
)

var modelHuggingFaceKimiK2Thinking = huggingFaceModelPreset(
	PresetHuggingFaceKimiK2Thinking,
	ModelNameHuggingFaceKimiK2Thinking,
	"HF Kimi K2 Thinking",
	262144,
	1.0,
)

var modelHuggingFaceGLM5 = huggingFaceModelPreset(
	PresetHuggingFaceGLM5,
	ModelNameHuggingFaceGLM5,
	"HF GLM 5",
	202752,
	1.0,
)

var modelHuggingFaceStep35Flash = huggingFaceModelPreset(
	PresetHuggingFaceStep35Flash,
	ModelNameHuggingFaceStep35Flash,
	"HF Step 3.5 Flash",
	131072,
	1.0,
)

var modelHuggingFaceMiniMaxM25 = huggingFaceModelPreset(
	PresetHuggingFaceMiniMaxM25,
	ModelNameHuggingFaceMiniMaxM25,
	"HF MiniMax M2.5",
	196608,
	1.0,
)

var modelHuggingFaceKimiK2Instruct = huggingFaceModelPreset(
	PresetHuggingFaceKimiK2Instruct,
	ModelNameHuggingFaceKimiK2Instruct,
	"HF Kimi K2 Instruct",
	131072,
	1.0,
)

var modelHuggingFaceKimiK2Instruct0905 = huggingFaceModelPreset(
	PresetHuggingFaceKimiK2Instruct0905,
	ModelNameHuggingFaceKimiK2Instruct0905,
	"HF Kimi K2 Instruct 0905",
	262144,
	1.0,
)

var modelHuggingFaceMiMoV2Flash = huggingFaceModelPreset(
	PresetHuggingFaceMiMoV2Flash,
	ModelNameHuggingFaceMiMoV2Flash,
	"HF MiMo V2 Flash",
	262144,
	1.0,
)

var modelHuggingFaceGLM47 = huggingFaceModelPreset(
	PresetHuggingFaceGLM47,
	ModelNameHuggingFaceGLM47,
	"HF GLM 4.7",
	128000,
	1.0,
)

var modelHuggingFaceGLM51FP8 = huggingFaceModelPreset(
	PresetHuggingFaceGLM51FP8,
	ModelNameHuggingFaceGLM51FP8,
	"HF GLM 5.1 FP8",
	202752,
	1.0,
)

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
		PresetHuggingFaceGPTOSS120:           modelHuggingFaceGPTOSS120,
		PresetHuggingFaceGPTOSS20:            modelHuggingFaceGPTOSS20,
		PresetHuggingFaceQwen3Coder30:        modelHuggingFaceQwen3Coder30,
		PresetHuggingFaceGLM52:               modelHuggingFaceGLM52,
		PresetHuggingFaceDeepSeekV4Flash:     modelHuggingFaceDeepSeekV4Flash,
		PresetHuggingFaceDeepSeekV4Pro:       modelHuggingFaceDeepSeekV4Pro,
		PresetHuggingFaceOrnith1035BFP8:      modelHuggingFaceOrnith1035BFP8,
		PresetHuggingFaceGLM52FP8:            modelHuggingFaceGLM52FP8,
		PresetHuggingFaceMiMoV25Pro:          modelHuggingFaceMiMoV25Pro,
		PresetHuggingFaceNemotron3UltraNVFP4: modelHuggingFaceNemotron3UltraNVFP4,
		PresetHuggingFaceQwen3CoderNext:      modelHuggingFaceQwen3CoderNext,
		PresetHuggingFaceNemotron3UltraBF16:  modelHuggingFaceNemotron3UltraBF16,
		PresetHuggingFaceGLM51:               modelHuggingFaceGLM51,
		PresetHuggingFaceNemotron3SuperBF16:  modelHuggingFaceNemotron3SuperBF16,
		PresetHuggingFaceMiniMaxM27:          modelHuggingFaceMiniMaxM27,
		PresetHuggingFaceKimiK2Thinking:      modelHuggingFaceKimiK2Thinking,
		PresetHuggingFaceGLM5:                modelHuggingFaceGLM5,
		PresetHuggingFaceStep35Flash:         modelHuggingFaceStep35Flash,
		PresetHuggingFaceMiniMaxM25:          modelHuggingFaceMiniMaxM25,
		PresetHuggingFaceKimiK2Instruct:      modelHuggingFaceKimiK2Instruct,
		PresetHuggingFaceKimiK2Instruct0905:  modelHuggingFaceKimiK2Instruct0905,
		PresetHuggingFaceMiMoV2Flash:         modelHuggingFaceMiMoV2Flash,
		PresetHuggingFaceGLM47:               modelHuggingFaceGLM47,
		PresetHuggingFaceGLM51FP8:            modelHuggingFaceGLM51FP8,
	},
}
