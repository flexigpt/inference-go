package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelSGLangGemma426BA4B = ModelPreset{
	ID:          PresetGemma426BA4B,
	Name:        ModelNameGemma426BA4BRepo,
	DisplayName: DisplayNameGemma426BA4B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemma426BA4BRepo,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), true, false),
}

var modelSGLangGPTOSS20B = ModelPreset{
	ID:          PresetGPTOSS20B,
	Name:        ModelNameGPTOSS20BRepo,
	DisplayName: DisplayNameGPTOSS20B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPTOSS20BRepo,
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

var modelSGLangQwen3635BA3B = ModelPreset{
	ID:          PresetQwen3635BA3B,
	Name:        ModelNameQwen3635BA3BRepo,
	DisplayName: DisplayNameQwen3635BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3635BA3BRepo,
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

var modelSGLangQwen3VL30BA3B = ModelPreset{
	ID:          PresetQwen3VL30BA3B,
	Name:        ModelNameQwen3VL30BA3BRepo,
	DisplayName: DisplayNameQwen3VL30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3VL30BA3BRepo,
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

var modelSGLangDeepSeekR18B = ModelPreset{
	ID:          PresetDeepSeekR18B,
	Name:        ModelNameDeepSeekR18BRepo,
	DisplayName: DisplayNameDeepSeekR18B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekR18BRepo,
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

var modelSGLangMinistral314B = ModelPreset{
	ID:          PresetMinistral314B,
	Name:        ModelNameMinistral314BRepo,
	DisplayName: DisplayNameMinistral314B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMinistral314BRepo,
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

var modelSGLangQwen3Coder30BA3B = ModelPreset{
	ID:          PresetQwen3Coder30BA3B,
	Name:        ModelNameQwen3Coder30BA3BRepo,
	DisplayName: DisplayNameQwen3Coder30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Coder30BA3BRepo,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnly(),
}

var modelSGLangGLM47Flash30BA3B = ModelPreset{
	ID:          PresetGLM47Flash30BA3B,
	Name:        ModelNameGLM47Flash30BA3BRepo,
	DisplayName: DisplayNameGLM47Flash30BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGLM47Flash30BA3BRepo,
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

var modelSGLangPhi4Reasoning14B = ModelPreset{
	ID:          PresetPhi4Reasoning14B,
	Name:        ModelNamePhi4Reasoning14BRepo,
	DisplayName: DisplayNamePhi4Reasoning14B,
	ModelParam: spec.ModelParam{
		Name:            ModelNamePhi4Reasoning14BRepo,
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

var modelSGLangDevstral224B = ModelPreset{
	ID:          PresetDevstral224B,
	Name:        ModelNameDevstral224BRepo,
	DisplayName: DisplayNameDevstral224B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDevstral224BRepo,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImage(),
}

var providerSGLang = ProviderPreset{
	Name:                     ProviderSGLang,
	DisplayName:              DisplayNameProviderSGLang,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "http://127.0.0.1:30000",
	ChatCompletionPathPrefix: spec.DefaultOpenAIResponsesPrefix,
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders:           sdkutil.CloneStringMap(spec.DefaultBaseHeaders),
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
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
		PresetGemma426BA4B:     modelSGLangGemma426BA4B,
		PresetGPTOSS20B:        modelSGLangGPTOSS20B,
		PresetQwen3635BA3B:     modelSGLangQwen3635BA3B,
		PresetQwen3VL30BA3B:    modelSGLangQwen3VL30BA3B,
		PresetDeepSeekR18B:     modelSGLangDeepSeekR18B,
		PresetMinistral314B:    modelSGLangMinistral314B,
		PresetQwen3Coder30BA3B: modelSGLangQwen3Coder30BA3B,
		PresetGLM47Flash30BA3B: modelSGLangGLM47Flash30BA3B,
		PresetPhi4Reasoning14B: modelSGLangPhi4Reasoning14B,
		PresetDevstral224B:     modelSGLangDevstral224B,
	},
}
