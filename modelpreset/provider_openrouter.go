package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetOpenRouterDeepSeekV4Flash          ModelPresetID = "deepseekV4Flash"
	PresetOpenRouterXiaomiMiMoV25            ModelPresetID = "xiaomiMiMoV25"
	PresetOpenRouterTencentHy3Preview        ModelPresetID = "tencentHy3Preview"
	PresetOpenRouterMiniMaxM3                ModelPresetID = "minimaxM3"
	PresetOpenRouterZAIGLM52                 ModelPresetID = "zaiglm52"
	PresetOpenRouterDeepSeekV4Pro            ModelPresetID = "deepseekV4Pro"
	PresetOpenRouterStepFunStep37Flash       ModelPresetID = "stepfunStep37Flash"
	PresetOpenRouterNVIDIANemotron3UltraFree ModelPresetID = "nvidiaNemotron3UltraFree"
	PresetOpenRouterPoolsideLagunaM1Free     ModelPresetID = "poolsideLagunaM1Free"
	PresetOpenRouterXiaomiMiMoV25Pro         ModelPresetID = "xiaomiMiMoV25Pro"
	PresetOpenRouterNVIDIANemotron3SuperFree ModelPresetID = "nvidiaNemotron3SuperFree"
	PresetOpenRouterMoonshotKimiK26          ModelPresetID = "moonshotKimiK26"
	PresetOpenRouterQwen37Max                ModelPresetID = "qwen37Max"
	PresetOpenRouterZAIGLM51                 ModelPresetID = "zaiglm51"
	PresetOpenRouterMoonshotKimiK27Code      ModelPresetID = "moonshotKimiK27Code"
	PresetOpenRouterQwen37Plus               ModelPresetID = "qwen37Plus"

	PresetOpenRouterMiniMaxM27     ModelPresetID = "minimaxm27"
	PresetOpenRouterMiniMaxM25Free ModelPresetID = "minimaxm25free"
)

const (
	ModelNameOpenRouterDeepSeekV4Flash          spec.ModelName = "deepseek/deepseek-v4-flash"
	ModelNameOpenRouterXiaomiMiMoV25            spec.ModelName = "xiaomi/mimo-v2.5"
	ModelNameOpenRouterTencentHy3Preview        spec.ModelName = "tencent/hy3-preview"
	ModelNameOpenRouterMiniMaxM3                spec.ModelName = "minimax/minimax-m3"
	ModelNameOpenRouterZAIGLM52                 spec.ModelName = "z-ai/glm-5.2"
	ModelNameOpenRouterDeepSeekV4Pro            spec.ModelName = "deepseek/deepseek-v4-pro"
	ModelNameOpenRouterStepFunStep37Flash       spec.ModelName = "stepfun/step-3.7-flash"
	ModelNameOpenRouterNVIDIANemotron3UltraFree spec.ModelName = "nvidia/nemotron-3-ultra-550b-a55b:free"
	ModelNameOpenRouterPoolsideLagunaM1Free     spec.ModelName = "poolside/laguna-m.1:free"
	ModelNameOpenRouterXiaomiMiMoV25Pro         spec.ModelName = "xiaomi/mimo-v2.5-pro"
	ModelNameOpenRouterNVIDIANemotron3SuperFree spec.ModelName = "nvidia/nemotron-3-super-120b-a12b:free"
	ModelNameOpenRouterMoonshotKimiK26          spec.ModelName = "moonshotai/kimi-k2.6"
	ModelNameOpenRouterQwen37Max                spec.ModelName = "qwen/qwen3.7-max"
	ModelNameOpenRouterZAIGLM51                 spec.ModelName = "z-ai/glm-5.1"
	ModelNameOpenRouterMoonshotKimiK27Code      spec.ModelName = "moonshotai/kimi-k2.7-code"
	ModelNameOpenRouterQwen37Plus               spec.ModelName = "qwen/qwen3.7-plus"

	ModelNameOpenRouterMiniMaxM27     spec.ModelName = "minimax/minimax-m2.7"
	ModelNameOpenRouterMiniMaxM25Free spec.ModelName = "minimax/minimax-m2.5:free"
)

var modelOpenRouterDeepSeekV4Flash = openRouterModelPreset(
	PresetOpenRouterDeepSeekV4Flash,
	ModelNameOpenRouterDeepSeekV4Flash,
	"OpenRouter DeepSeek V4 Flash",
	1048575,
	1048575,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelHigh, spec.ReasoningLevelXHigh},
		true,
		false,
	),
)

var modelOpenRouterXiaomiMiMoV25 = openRouterModelPreset(
	PresetOpenRouterXiaomiMiMoV25,
	ModelNameOpenRouterXiaomiMiMoV25,
	"OpenRouter Xiaomi MiMo V2.5",
	32000,
	32000,
	new(1.0),
	nil,
	openRouterModelOverride(openRouterTextAudioImageVideoModalities(), openRouterDefaultReasoningLevels(), true, false),
)

var modelOpenRouterTencentHy3Preview = openRouterModelPreset(
	PresetOpenRouterTencentHy3Preview,
	ModelNameOpenRouterTencentHy3Preview,
	"OpenRouter Tencent Hy3 Preview",
	262144,
	262144,
	new(0.9),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelNone, spec.ReasoningLevelLow, spec.ReasoningLevelHigh},
		false,
		false,
	),
)

var modelOpenRouterMiniMaxM3 = openRouterModelPreset(
	PresetOpenRouterMiniMaxM3,
	ModelNameOpenRouterMiniMaxM3,
	"OpenRouter MiniMax M3",
	524288,
	512000,
	new(1.0),
	nil,
	openRouterModelOverride(openRouterTextImageVideoModalities(), openRouterDefaultReasoningLevels(), true, false),
)

var modelOpenRouterZAIGLM52 = openRouterModelPreset(
	PresetOpenRouterZAIGLM52,
	ModelNameOpenRouterZAIGLM52,
	"OpenRouter Z.AI GLM 5.2",
	1048576,
	32768,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelHigh, spec.ReasoningLevelXHigh},
		true,
		true,
	),
)

var modelOpenRouterDeepSeekV4Pro = openRouterModelPreset(
	PresetOpenRouterDeepSeekV4Pro,
	ModelNameOpenRouterDeepSeekV4Pro,
	"OpenRouter DeepSeek V4 Pro",
	1048576,
	384000,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelHigh, spec.ReasoningLevelXHigh},
		true,
		false,
	),
)

var modelOpenRouterStepFunStep37Flash = openRouterModelPreset(
	PresetOpenRouterStepFunStep37Flash,
	ModelNameOpenRouterStepFunStep37Flash,
	"OpenRouter StepFun Step 3.7 Flash",
	256000,
	256000,
	nil,
	reasoningSingle(spec.ReasoningLevelMedium),
	openRouterModelOverride(
		openRouterTextImageVideoModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelLow, spec.ReasoningLevelMedium, spec.ReasoningLevelHigh},
		true,
		false,
	),
)

var modelOpenRouterNVIDIANemotron3UltraFree = openRouterModelPreset(
	PresetOpenRouterNVIDIANemotron3UltraFree,
	ModelNameOpenRouterNVIDIANemotron3UltraFree,
	"OpenRouter NVIDIA Nemotron 3 Ultra Free",
	1000000,
	65536,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelMedium, spec.ReasoningLevelHigh},
		false,
		false,
	),
)

var modelOpenRouterPoolsideLagunaM1Free = openRouterModelPreset(
	PresetOpenRouterPoolsideLagunaM1Free,
	ModelNameOpenRouterPoolsideLagunaM1Free,
	"OpenRouter Poolside Laguna M.1 Free",
	262144,
	32768,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextOnlyModalities(), openRouterDefaultReasoningLevels(), false, false),
)

var modelOpenRouterXiaomiMiMoV25Pro = openRouterModelPreset(
	PresetOpenRouterXiaomiMiMoV25Pro,
	ModelNameOpenRouterXiaomiMiMoV25Pro,
	"OpenRouter Xiaomi MiMo V2.5 Pro",
	1048576,
	131072,
	new(1.0),
	nil,
	openRouterModelOverride(openRouterTextOnlyModalities(), openRouterDefaultReasoningLevels(), true, false),
)

var modelOpenRouterNVIDIANemotron3SuperFree = openRouterModelPreset(
	PresetOpenRouterNVIDIANemotron3SuperFree,
	ModelNameOpenRouterNVIDIANemotron3SuperFree,
	"OpenRouter NVIDIA Nemotron 3 Super Free",
	262144,
	262144,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelMedium),
	openRouterModelOverride(
		openRouterTextOnlyModalities(),
		[]spec.ReasoningLevel{spec.ReasoningLevelLow, spec.ReasoningLevelMedium},
		true,
		false,
	),
)

var modelOpenRouterMoonshotKimiK26 = openRouterModelPreset(
	PresetOpenRouterMoonshotKimiK26,
	ModelNameOpenRouterMoonshotKimiK26,
	"OpenRouter MoonshotAI Kimi K2.6",
	262144,
	262144,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextImageModalities(), openRouterDefaultReasoningLevels(), true, true),
)

var modelOpenRouterQwen37Max = openRouterModelPreset(
	PresetOpenRouterQwen37Max,
	ModelNameOpenRouterQwen37Max,
	"OpenRouter Qwen3.7 Max",
	1000000,
	65536,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextOnlyModalities(), openRouterDefaultReasoningLevels(), true, false),
)

var modelOpenRouterZAIGLM51 = openRouterModelPreset(
	PresetOpenRouterZAIGLM51,
	ModelNameOpenRouterZAIGLM51,
	"OpenRouter Z.AI GLM 5.1",
	65536,
	65536,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextOnlyModalities(), openRouterDefaultReasoningLevels(), true, true),
)

var modelOpenRouterMoonshotKimiK27Code = openRouterModelPreset(
	PresetOpenRouterMoonshotKimiK27Code,
	ModelNameOpenRouterMoonshotKimiK27Code,
	"OpenRouter MoonshotAI Kimi K2.7 Code",
	262144,
	16384,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextImageModalities(), openRouterDefaultReasoningLevels(), true, true),
)

var modelOpenRouterQwen37Plus = openRouterModelPreset(
	PresetOpenRouterQwen37Plus,
	ModelNameOpenRouterQwen37Plus,
	"OpenRouter Qwen3.7 Plus",
	1000000,
	65536,
	nil,
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterModelOverride(openRouterTextImageModalities(), openRouterDefaultReasoningLevels(), true, false),
)

var modelOpenRouterMiniMaxM27 = openRouterModelPreset(
	PresetOpenRouterMiniMaxM27,
	ModelNameOpenRouterMiniMaxM27,
	"OpenRouter MiniMax M2.7",
	180000,
	125000,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterReasoningOverride(),
)

var modelOpenRouterMiniMaxM25Free = openRouterModelPreset(
	PresetOpenRouterMiniMaxM25Free,
	ModelNameOpenRouterMiniMaxM25Free,
	"OpenRouter MiniMax M2.5 free",
	100000,
	8000,
	new(1.0),
	reasoningSingle(spec.ReasoningLevelHigh),
	openRouterReasoningOverride(),
)

func openRouterModelPreset(
	id ModelPresetID,
	name spec.ModelName,
	displayName string,
	maxPromptLength int,
	maxOutputLength int,
	temperature *float64,
	reasoning *spec.ReasoningParam,
	capabilitiesOverride *capabilityoverride.ModelCapabilitiesOverride,
) ModelPreset {
	return ModelPreset{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		ModelParam: spec.ModelParam{
			Name:            name,
			Stream:          true,
			MaxPromptLength: maxPromptLength,
			MaxOutputLength: maxOutputLength,
			Temperature:     temperature,
			Reasoning:       reasoning,
			SystemPrompt:    "",
			Timeout:         1800,
		},
		CapabilitiesOverride: capabilitiesOverride,
	}
}

func openRouterTextOnlyModalities() []spec.Modality {
	return []spec.Modality{spec.ModalityTextIn}
}

func openRouterTextImageModalities() []spec.Modality {
	return []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn}
}

func openRouterTextImageVideoModalities() []spec.Modality {
	return []spec.Modality{spec.ModalityTextIn, spec.ModalityImageIn, spec.ModalityVideoIn}
}

func openRouterTextAudioImageVideoModalities() []spec.Modality {
	return []spec.Modality{spec.ModalityTextIn, spec.ModalityAudioIn, spec.ModalityImageIn, spec.ModalityVideoIn}
}

func openRouterDefaultReasoningLevels() []spec.ReasoningLevel {
	return []spec.ReasoningLevel{
		spec.ReasoningLevelLow,
		spec.ReasoningLevelMedium,
		spec.ReasoningLevelHigh,
	}
}

func openRouterModelOverride(
	modalitiesIn []spec.Modality,
	reasoningLevels []spec.ReasoningLevel,
	supportsJSONSchema bool,
	supportsParallelToolCalls bool,
) *capabilityoverride.ModelCapabilitiesOverride {
	outputFormats := []spec.OutputFormatKind{
		spec.OutputFormatKindText,
	}
	if supportsJSONSchema {
		outputFormats = append(outputFormats, spec.OutputFormatKindJSONSchema)
	}

	return &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: modalitiesIn,
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: openRouterReasoningCapabilitiesOverride(reasoningLevels...),
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: outputFormats,
			SupportsVerbosity:      new(false),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(supportsParallelToolCalls),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	}
}

func openRouterReasoningOverride(levels ...spec.ReasoningLevel) *capabilityoverride.ModelCapabilitiesOverride {
	return &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: openRouterReasoningCapabilitiesOverride(levels...),
	}
}

func openRouterReasoningCapabilitiesOverride(
	levels ...spec.ReasoningLevel,
) *capabilityoverride.ReasoningCapabilitiesOverride {
	if len(levels) == 0 {
		levels = openRouterDefaultReasoningLevels()
	}

	return &capabilityoverride.ReasoningCapabilitiesOverride{
		TemperatureDisallowedWhenEnabled: new(false),
		SupportedReasoningTypes: []spec.ReasoningType{
			spec.ReasoningTypeSingleWithLevels,
		},
		SupportedReasoningLevels: levels,
		SupportsSummaryStyle:     new(true),
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
		PresetOpenRouterDeepSeekV4Flash:          modelOpenRouterDeepSeekV4Flash,
		PresetOpenRouterXiaomiMiMoV25:            modelOpenRouterXiaomiMiMoV25,
		PresetOpenRouterTencentHy3Preview:        modelOpenRouterTencentHy3Preview,
		PresetOpenRouterMiniMaxM3:                modelOpenRouterMiniMaxM3,
		PresetOpenRouterZAIGLM52:                 modelOpenRouterZAIGLM52,
		PresetOpenRouterDeepSeekV4Pro:            modelOpenRouterDeepSeekV4Pro,
		PresetOpenRouterStepFunStep37Flash:       modelOpenRouterStepFunStep37Flash,
		PresetOpenRouterNVIDIANemotron3UltraFree: modelOpenRouterNVIDIANemotron3UltraFree,
		PresetOpenRouterPoolsideLagunaM1Free:     modelOpenRouterPoolsideLagunaM1Free,
		PresetOpenRouterXiaomiMiMoV25Pro:         modelOpenRouterXiaomiMiMoV25Pro,
		PresetOpenRouterNVIDIANemotron3SuperFree: modelOpenRouterNVIDIANemotron3SuperFree,
		PresetOpenRouterMoonshotKimiK26:          modelOpenRouterMoonshotKimiK26,
		PresetOpenRouterQwen37Max:                modelOpenRouterQwen37Max,
		PresetOpenRouterZAIGLM51:                 modelOpenRouterZAIGLM51,
		PresetOpenRouterMoonshotKimiK27Code:      modelOpenRouterMoonshotKimiK27Code,
		PresetOpenRouterQwen37Plus:               modelOpenRouterQwen37Plus,
		PresetOpenRouterMiniMaxM27:               modelOpenRouterMiniMaxM27,
		PresetOpenRouterMiniMaxM25Free:           modelOpenRouterMiniMaxM25Free,
	},
}
