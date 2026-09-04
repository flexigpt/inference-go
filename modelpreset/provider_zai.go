package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderZAI spec.ProviderName = "zai"

	DisplayNameProviderZAI = "Z.AI"
)

const (
	ProviderZAICodingPlan spec.ProviderName = "zaiCodingPlan"

	DisplayNameProviderZAICodingPlan = "Z.AI Coding Plan"
)

const (
	ModelNameGLM53        spec.ModelName = "glm-5.3"
	ModelNameGLM52        spec.ModelName = "glm-5.2"
	ModelNameGLM51        spec.ModelName = "glm-5.1"
	ModelNameGLM5         spec.ModelName = "glm-5"
	ModelNameGLM5Turbo    spec.ModelName = "glm-5-turbo"
	ModelNameGLM47        spec.ModelName = "glm-4.7"
	ModelNameGLM47FlashX  spec.ModelName = "glm-4.7-flashx"
	ModelNameGLM47Flash   spec.ModelName = "glm-4.7-flash"
	ModelNameGLM46        spec.ModelName = "glm-4.6"
	ModelNameGLM45        spec.ModelName = "glm-4.5"
	ModelNameGLM45X       spec.ModelName = "glm-4.5-x"
	ModelNameGLM45Air     spec.ModelName = "glm-4.5-air"
	ModelNameGLM45AirX    spec.ModelName = "glm-4.5-airx"
	ModelNameGLM45Flash   spec.ModelName = "glm-4.5-flash"
	ModelNameGLM5VTurbo   spec.ModelName = "glm-5v-turbo"
	ModelNameGLM46V       spec.ModelName = "glm-4.6v"
	ModelNameGLM46VFlash  spec.ModelName = "glm-4.6v-flash"
	ModelNameGLM46VFlashX spec.ModelName = "glm-4.6v-flashx"
	ModelNameGLM45V       spec.ModelName = "glm-4.5v"

	ModelNameGLM47Flash30BA3BRepo  spec.ModelName = "zai-org/GLM-4.7-Flash"
	ModelNameGLM47Flash30BA3BLocal spec.ModelName = "glm-4.7-flash"

	ModelNameGLM52FireworksAI    spec.ModelName = "zai-org/GLM-5.2:fireworks-ai"
	ModelNameGLM51FireworksAI    spec.ModelName = "zai-org/GLM-5.1:fireworks-ai"
	ModelNameGLM51FP8FireworksAI spec.ModelName = "zai-org/GLM-5.1-FP8:fireworks-ai"

	ModelNameGLM52FP8ZAI spec.ModelName = "zai-org/GLM-5.2-FP8:zai-org"

	ModelNameGLM5Novita    spec.ModelName = "zai-org/GLM-5:novita"
	ModelNameGLM47Cerebras spec.ModelName = "zai-org/GLM-4.7:cerebras"

	ModelNameOpenRouterZAIGLM51 spec.ModelName = "z-ai/glm-5.1"
	ModelNameOpenRouterZAIGLM52 spec.ModelName = "z-ai/glm-5.2"
)

const (
	DisplayNameGLM53        = "GLM 5.3"
	DisplayNameGLM5Turbo    = "GLM 5 Turbo"
	DisplayNameGLM47FlashX  = "GLM 4.7 FlashX"
	DisplayNameGLM47Flash   = "GLM 4.7 Flash"
	DisplayNameGLM46        = "GLM 4.6"
	DisplayNameGLM45        = "GLM 4.5"
	DisplayNameGLM45X       = "GLM 4.5 X"
	DisplayNameGLM45Air     = "GLM 4.5 Air"
	DisplayNameGLM45AirX    = "GLM 4.5 AirX"
	DisplayNameGLM45Flash   = "GLM 4.5 Flash"
	DisplayNameGLM5VTurbo   = "GLM 5V Turbo"
	DisplayNameGLM46V       = "GLM 4.6V"
	DisplayNameGLM46VFlash  = "GLM 4.6V Flash"
	DisplayNameGLM46VFlashX = "GLM 4.6V FlashX"
	DisplayNameGLM45V       = "GLM 4.5V"

	DisplayNameGLM47Flash30BA3B = "GLM-4.7 Flash 30B A3B"
	DisplayNameGLM47            = "GLM 4.7"
	DisplayNameGLM5             = "GLM 5"
	DisplayNameGLM51            = "GLM 5.1"
	DisplayNameGLM51FP8         = "GLM 5.1 FP8"
	DisplayNameGLM52            = "GLM 5.2"
	DisplayNameGLM52FP8         = "GLM 5.2 FP8"

	DisplayNameZAIGLM51    = "Z.AI GLM 5.1"
	DisplayNameZAIGLM52    = "Z.AI GLM 5.2"
	DisplayNameGLM52FP8ZAI = "GLM 5.2 FP8 (Z.AI)"

	DisplayNameGLM51FireworksAI    = "GLM 5.1 (Fireworks AI)"
	DisplayNameGLM51FP8FireworksAI = "GLM 5.1 FP8 (Fireworks AI)"
	DisplayNameGLM52FireworksAI    = "GLM 5.2 (Fireworks AI)"

	DisplayNameGLM47Cerebras = "GLM 4.7 (Cerebras)"
	DisplayNameGLM5Novita    = "GLM 5 (Novita)"
)

const (
	PresetGLM53        ModelPresetID = "glm53"
	PresetGLM5Turbo    ModelPresetID = "glm5Turbo"
	PresetGLM47FlashX  ModelPresetID = "glm47FlashX"
	PresetGLM47Flash   ModelPresetID = "glm47Flash"
	PresetGLM46        ModelPresetID = "glm46"
	PresetGLM45        ModelPresetID = "glm45"
	PresetGLM45X       ModelPresetID = "glm45X"
	PresetGLM45Air     ModelPresetID = "glm45Air"
	PresetGLM45AirX    ModelPresetID = "glm45AirX"
	PresetGLM45Flash   ModelPresetID = "glm45Flash"
	PresetGLM5VTurbo   ModelPresetID = "glm5VTurbo"
	PresetGLM46V       ModelPresetID = "glm46V"
	PresetGLM46VFlash  ModelPresetID = "glm46VFlash"
	PresetGLM46VFlashX ModelPresetID = "glm46VFlashX"
	PresetGLM45V       ModelPresetID = "glm45V"

	PresetGLM47Flash30BA3B ModelPresetID = "glm47flash30ba3b"
	PresetGLM47            ModelPresetID = "glm47"
	PresetGLM5             ModelPresetID = "glm5"
	PresetGLM51            ModelPresetID = "glm51"
	PresetGLM51FP8         ModelPresetID = "glm51fp8"
	PresetGLM52            ModelPresetID = "glm52"
	PresetGLM52FP8         ModelPresetID = "glm52fp8"

	PresetZAIGLM51    ModelPresetID = "zaiglm51"
	PresetZAIGLM52    ModelPresetID = "zaiglm52"
	PresetGLM52FP8ZAI ModelPresetID = "glm52fp8ZAI"

	PresetGLM52FireworksAI    ModelPresetID = "glm52FireworksAI"
	PresetGLM51FireworksAI    ModelPresetID = "glm51FireworksAI"
	PresetGLM51FP8FireworksAI ModelPresetID = "glm51fp8FireworksAI"

	PresetGLM5Novita    ModelPresetID = "glm5Novita"
	PresetGLM47Cerebras ModelPresetID = "glm47Cerebras"
)

var modelZAIPayAsYouGoGLM53 = ModelPreset{
	ID:          PresetGLM53,
	Name:        ModelNameGLM53,
	DisplayName: DisplayNameGLM53,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM53,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMax),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelZAIPayAsYouGoGLM52 = ModelPreset{
	ID:          PresetGLM52,
	Name:        ModelNameGLM52,
	DisplayName: DisplayNameGLM52,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM52,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMax),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
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
	},
}

var modelZAIPayAsYouGoGLM51 = ModelPreset{
	ID:          PresetGLM51,
	Name:        ModelNameGLM51,
	DisplayName: DisplayNameGLM51,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM51,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelZAIPayAsYouGoGLM5 = ModelPreset{
	ID:          PresetGLM5,
	Name:        ModelNameGLM5,
	DisplayName: DisplayNameGLM5,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM5,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM5Turbo = ModelPreset{
	ID:          PresetGLM5Turbo,
	Name:        ModelNameGLM5Turbo,
	DisplayName: DisplayNameGLM5Turbo,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM5Turbo,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM47 = ModelPreset{
	ID:          PresetGLM47,
	Name:        ModelNameGLM47,
	DisplayName: DisplayNameGLM47,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM47FlashX = ModelPreset{
	ID:          PresetGLM47FlashX,
	Name:        ModelNameGLM47FlashX,
	DisplayName: DisplayNameGLM47FlashX,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47FlashX,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM47Flash = ModelPreset{
	ID:          PresetGLM47Flash,
	Name:        ModelNameGLM47Flash,
	DisplayName: DisplayNameGLM47Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47Flash,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM46 = ModelPreset{
	ID:          PresetGLM46,
	Name:        ModelNameGLM46,
	DisplayName: DisplayNameGLM46,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM46,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45 = ModelPreset{
	ID:          PresetGLM45,
	Name:        ModelNameGLM45,
	DisplayName: DisplayNameGLM45,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 98304,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45X = ModelPreset{
	ID:          PresetGLM45X,
	Name:        ModelNameGLM45X,
	DisplayName: DisplayNameGLM45X,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45X,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 98304,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45Air = ModelPreset{
	ID:          PresetGLM45Air,
	Name:        ModelNameGLM45Air,
	DisplayName: DisplayNameGLM45Air,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45Air,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 98304,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45AirX = ModelPreset{
	ID:          PresetGLM45AirX,
	Name:        ModelNameGLM45AirX,
	DisplayName: DisplayNameGLM45AirX,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45AirX,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 98304,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45Flash = ModelPreset{
	ID:          PresetGLM45Flash,
	Name:        ModelNameGLM45Flash,
	DisplayName: DisplayNameGLM45Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45Flash,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 98304,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM5VTurbo = ModelPreset{
	ID:          PresetGLM5VTurbo,
	Name:        ModelNameGLM5VTurbo,
	DisplayName: DisplayNameGLM5VTurbo,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM5VTurbo,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM46V = ModelPreset{
	ID:          PresetGLM46V,
	Name:        ModelNameGLM46V,
	DisplayName: DisplayNameGLM46V,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM46V,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 32768,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM46VFlash = ModelPreset{
	ID:          PresetGLM46VFlash,
	Name:        ModelNameGLM46VFlash,
	DisplayName: DisplayNameGLM46VFlash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM46VFlash,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 32768,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM46VFlashX = ModelPreset{
	ID:          PresetGLM46VFlashX,
	Name:        ModelNameGLM46VFlashX,
	DisplayName: DisplayNameGLM46VFlashX,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM46VFlashX,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 32768,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAIPayAsYouGoGLM45V = ModelPreset{
	ID:          PresetGLM45V,
	Name:        ModelNameGLM45V,
	DisplayName: DisplayNameGLM45V,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM45V,
		Stream:          true,
		MaxPromptLength: 65536,
		MaxOutputLength: 16384,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes:  []spec.ReasoningType{},
			SupportedReasoningLevels: []spec.ReasoningLevel{},
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes:        []spec.ToolType{},
			SupportedToolPolicyModes:  []spec.ToolPolicyMode{},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelZAICodingPlanGLM53 = ModelPreset{
	ID:          PresetGLM53,
	Name:        ModelNameGLM53,
	DisplayName: DisplayNameGLM53,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM53,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMax),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelMax,
			},
			SupportsSummaryStyle:             new(true),
			SupportsEncryptedReasoningInput:  new(false),
			TemperatureDisallowedWhenEnabled: new(false),
		},
	},
}

var modelZAICodingPlanGLM5Turbo = ModelPreset{
	ID:          PresetGLM5Turbo,
	Name:        ModelNameGLM5Turbo,
	DisplayName: DisplayNameGLM5Turbo,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM5Turbo,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var modelZAICodingPlanGLM47 = ModelPreset{
	ID:          PresetGLM47,
	Name:        ModelNameGLM47,
	DisplayName: DisplayNameGLM47,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47,
		Stream:          true,
		MaxPromptLength: 204800,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: noReasoningOverride,
}

var providerZAI = ProviderPreset{
	Name:        ProviderZAI,
	DisplayName: DisplayNameProviderZAI,

	// Z.AI documents pay-as-you-go through its OpenAI-compatible Chat
	// Completions endpoint. The OpenAI Chat adapter strips the terminal
	// chat/completions suffix before configuring the SDK base URL.
	SDKType:                  spec.ProviderSDKTypeOpenAIChatCompletions,
	Origin:                   "https://api.z.ai",
	ChatCompletionPathPrefix: "/api/paas/v4/chat/completions",
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
		"accept-language":                "en-US,en",
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:          new(false),
			SupportedReasoningTypes:          []spec.ReasoningType{},
			SupportedReasoningLevels:         []spec.ReasoningLevel{},
			SupportsSummaryStyle:             new(false),
			SupportsEncryptedReasoningInput:  new(false),
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
			},
			SupportsVerbosity: new(false),
		},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
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
		ParamDialect: &capabilityoverride.ParamDialectOverride{
			MaxOutputTokensParamName: new(spec.MaxOutputTokensParamNameMaxTokens),
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGLM53:        modelZAIPayAsYouGoGLM53,
		PresetGLM52:        modelZAIPayAsYouGoGLM52,
		PresetGLM51:        modelZAIPayAsYouGoGLM51,
		PresetGLM5:         modelZAIPayAsYouGoGLM5,
		PresetGLM5Turbo:    modelZAIPayAsYouGoGLM5Turbo,
		PresetGLM47:        modelZAIPayAsYouGoGLM47,
		PresetGLM47FlashX:  modelZAIPayAsYouGoGLM47FlashX,
		PresetGLM47Flash:   modelZAIPayAsYouGoGLM47Flash,
		PresetGLM46:        modelZAIPayAsYouGoGLM46,
		PresetGLM45:        modelZAIPayAsYouGoGLM45,
		PresetGLM45X:       modelZAIPayAsYouGoGLM45X,
		PresetGLM45Air:     modelZAIPayAsYouGoGLM45Air,
		PresetGLM45AirX:    modelZAIPayAsYouGoGLM45AirX,
		PresetGLM45Flash:   modelZAIPayAsYouGoGLM45Flash,
		PresetGLM5VTurbo:   modelZAIPayAsYouGoGLM5VTurbo,
		PresetGLM46V:       modelZAIPayAsYouGoGLM46V,
		PresetGLM46VFlash:  modelZAIPayAsYouGoGLM46VFlash,
		PresetGLM46VFlashX: modelZAIPayAsYouGoGLM46VFlashX,
		PresetGLM45V:       modelZAIPayAsYouGoGLM45V,
	},
}

var providerZAICodingPlan = ProviderPreset{
	Name:        ProviderZAICodingPlan,
	DisplayName: DisplayNameProviderZAICodingPlan,

	// Coding Plan keys and pay-as-you-go/prepaid keys are not interchangeable.
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://api.z.ai",
	ChatCompletionPathPrefix: "/api/v1/responses",
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
		"accept-language":                "en-US,en",
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:          new(false),
			SupportedReasoningTypes:          []spec.ReasoningType{},
			SupportedReasoningLevels:         []spec.ReasoningLevel{},
			SupportsSummaryStyle:             new(false),
			SupportsReasoningContext:         new(false),
			SupportsReasoningMode:            new(false),
			SupportsEncryptedReasoningInput:  new(false),
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
				spec.ToolPolicyModeTool,
			},
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
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
		PresetGLM53:     modelZAICodingPlanGLM53,
		PresetGLM5Turbo: modelZAICodingPlanGLM5Turbo,
		PresetGLM47:     modelZAICodingPlanGLM47,
	},
}
