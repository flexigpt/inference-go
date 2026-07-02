package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelLocalAIGemma426BA4B = ModelPreset{
	ID:          PresetGemma426BA4B,
	Name:        ModelNameGemma426BA4BLocal,
	DisplayName: DisplayNameGemma426BA4B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemma426BA4BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIGPTOSS20B = ModelPreset{
	ID:          PresetGPTOSS20B,
	Name:        ModelNameGPTOSS20BLocal,
	DisplayName: DisplayNameGPTOSS20B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPTOSS20BLocal,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(false), true, false),
}

var modelLocalAIQwen3635BA3B = ModelPreset{
	ID:          PresetQwen3635BA3B,
	Name:        ModelNameQwen3635BA3BLocal,
	DisplayName: DisplayNameQwen3635BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3635BA3BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIDeepSeekR18B = ModelPreset{
	ID:          PresetDeepSeekR18B,
	Name:        ModelNameDeepSeekR18BLocal,
	DisplayName: DisplayNameDeepSeekR18B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekR18BLocal,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIQwen3VL30BA3B = ModelPreset{
	ID:          PresetQwen3VL30BA3B,
	Name:        ModelNameQwen3VL30BA3BLocal,
	DisplayName: DisplayNameQwen3VL30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3VL30BA3BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIMinistral314B = ModelPreset{
	ID:          PresetMinistral314B,
	Name:        ModelNameMinistral314BLocal,
	DisplayName: DisplayNameMinistral314B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMinistral314BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIQwen3Coder30BA3B = ModelPreset{
	ID:          PresetQwen3Coder30BA3B,
	Name:        ModelNameQwen3Coder30BA3BLocal,
	DisplayName: DisplayNameQwen3Coder30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Coder30BA3BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnly(),
}

var modelLocalAIGLM47Flash30BA3B = ModelPreset{
	ID:          PresetGLM47Flash30BA3B,
	Name:        ModelNameGLM47Flash30BA3BLocal,
	DisplayName: DisplayNameGLM47Flash30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47Flash30BA3BLocal,
		Stream:          true,
		MaxPromptLength: 202752,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(true), true, false),
}

var modelLocalAIPhi4Reasoning14B = ModelPreset{
	ID:          PresetPhi4Reasoning14B,
	Name:        ModelNamePhi4Reasoning14BLocal,
	DisplayName: DisplayNamePhi4Reasoning14B,
	ModelParam: spec.ModelParam{
		Name:            ModelNamePhi4Reasoning14BLocal,
		Stream:          true,
		MaxPromptLength: 32768,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(false), true, false),
}

var modelLocalAIDevstral224B = ModelPreset{
	ID:          PresetDevstral224B,
	Name:        ModelNameDevstral224BLocal,
	DisplayName: DisplayNameDevstral224B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDevstral224BLocal,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImage(),
}

var providerLocalAI = ProviderPreset{
	Name:                     ProviderLocalAI,
	DisplayName:              DisplayNameProviderLocalAI,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "http://127.0.0.1:8080",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
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
		PresetGemma426BA4B:     modelLocalAIGemma426BA4B,
		PresetGPTOSS20B:        modelLocalAIGPTOSS20B,
		PresetQwen3635BA3B:     modelLocalAIQwen3635BA3B,
		PresetDeepSeekR18B:     modelLocalAIDeepSeekR18B,
		PresetQwen3VL30BA3B:    modelLocalAIQwen3VL30BA3B,
		PresetMinistral314B:    modelLocalAIMinistral314B,
		PresetQwen3Coder30BA3B: modelLocalAIQwen3Coder30BA3B,
		PresetGLM47Flash30BA3B: modelLocalAIGLM47Flash30BA3B,
		PresetPhi4Reasoning14B: modelLocalAIPhi4Reasoning14B,
		PresetDevstral224B:     modelLocalAIDevstral224B,
	},
}
