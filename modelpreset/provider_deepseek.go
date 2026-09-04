package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderDeepSeek spec.ProviderName = "deepseek"

	DisplayNameProviderDeepSeek = "DeepSeek"
)

const (
	ModelNameDeepSeekV4Flash spec.ModelName = "deepseek-v4-flash"
	ModelNameDeepSeekV4Pro   spec.ModelName = "deepseek-v4-pro"

	ModelNameDeepSeekR18BRepo  spec.ModelName = "deepseek-ai/DeepSeek-R1-0528-Qwen3-8B"
	ModelNameDeepSeekR18BLocal spec.ModelName = "deepseek-r1-8b"

	ModelNameDeepSeekV4FlashFireworksAI spec.ModelName = "deepseek-ai/DeepSeek-V4-Flash:fireworks-ai"
	ModelNameDeepSeekV4ProFireworksAI   spec.ModelName = "deepseek-ai/DeepSeek-V4-Pro:fireworks-ai"

	ModelNameDeepSeekR18BOllama spec.ModelName = "deepseek-r1:8b"

	ModelNameOpenRouterDeepSeekV4Flash spec.ModelName = "deepseek/deepseek-v4-flash"
	ModelNameOpenRouterDeepSeekV4Pro   spec.ModelName = "deepseek/deepseek-v4-pro"
)

const (
	DisplayNameDeepSeekR18B    = "DeepSeek-R1 8B"
	DisplayNameDeepSeekV4Flash = "DeepSeek V4 Flash"
	DisplayNameDeepSeekV4Pro   = "DeepSeek V4 Pro"

	DisplayNameDeepSeekV4FlashFireworksAI = "DeepSeek V4 Flash (Fireworks AI)"
	DisplayNameDeepSeekV4ProFireworksAI   = "DeepSeek V4 Pro (Fireworks AI)"
)

const (
	PresetDeepSeekR18B    ModelPresetID = "deepseekr18b"
	PresetDeepSeekV4Flash ModelPresetID = "deepseekv4flash"
	PresetDeepSeekV4Pro   ModelPresetID = "deepseekv4pro"

	PresetDeepSeekV4FlashFireworksAI ModelPresetID = "deepseekv4flashFireworksAI"
	PresetDeepSeekV4ProFireworksAI   ModelPresetID = "deepseekv4proFireworksAI"
)

var modelDeepSeekV4Flash = ModelPreset{
	ID:          PresetDeepSeekV4Flash,
	Name:        ModelNameDeepSeekV4Flash,
	DisplayName: DisplayNameDeepSeekV4Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekV4Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 393216,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelDeepSeekV4Pro = ModelPreset{
	ID:          PresetDeepSeekV4Pro,
	Name:        ModelNameDeepSeekV4Pro,
	DisplayName: DisplayNameDeepSeekV4Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekV4Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 393216,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var providerDeepSeek = ProviderPreset{
	Name:                     ProviderDeepSeek,
	DisplayName:              DisplayNameProviderDeepSeek,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.deepseek.com",
	ChatCompletionPathPrefix: "/responses",
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(false),
			SupportsReasoningContext:         new(false),
			SupportsReasoningMode:            new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(false),
			DisallowedWithReasoning: new(false),
			MaxSequences:            new(0),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
			},
			SupportsVerbosity: new(false),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
		CacheCapabilities: &capabilityoverride.CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(true),
			TopLevel: &capabilityoverride.CacheControlCapabilitiesOverride{
				SupportsTTL:    new(false),
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
				SupportsKey:    new(false),
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetDeepSeekV4Flash: modelDeepSeekV4Flash,
		PresetDeepSeekV4Pro:   modelDeepSeekV4Pro,
	},
}
