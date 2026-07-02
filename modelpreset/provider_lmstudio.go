package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelLMStudioGemma426BA4B = ModelPreset{
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

var modelLMStudioGPTOSS20B = ModelPreset{
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

var modelLMStudioQwen3635BA3B = ModelPreset{
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

var modelLMStudioQwen3627B = ModelPreset{
	ID:          PresetQwen3627B,
	Name:        ModelNameQwen3627BRepo,
	DisplayName: DisplayNameQwen3627B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3627BRepo,
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

var modelLMStudioDeepSeekR18B = ModelPreset{
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

var modelLMStudioQwen3VL30B = ModelPreset{
	ID:          PresetQwen3VL30B,
	Name:        ModelNameQwen3VL30BRepo,
	DisplayName: DisplayNameQwen3VL30B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3VL30BRepo,
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

var modelLMStudioMinistral314B = ModelPreset{
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

var modelLMStudioQwen3Coder30BA3B = ModelPreset{
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

var modelLMStudioGLM47Flash30BA3B = ModelPreset{
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

var modelLMStudioDevstral224B = ModelPreset{
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

var providerLMStudio = ProviderPreset{
	Name:                     ProviderLMStudio,
	DisplayName:              DisplayNameProviderLMStudio,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "http://127.0.0.1:1234",
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
		PresetGemma426BA4B:     modelLMStudioGemma426BA4B,
		PresetGPTOSS20B:        modelLMStudioGPTOSS20B,
		PresetQwen3635BA3B:     modelLMStudioQwen3635BA3B,
		PresetQwen3627B:        modelLMStudioQwen3627B,
		PresetDeepSeekR18B:     modelLMStudioDeepSeekR18B,
		PresetQwen3VL30B:       modelLMStudioQwen3VL30B,
		PresetMinistral314B:    modelLMStudioMinistral314B,
		PresetQwen3Coder30BA3B: modelLMStudioQwen3Coder30BA3B,
		PresetGLM47Flash30BA3B: modelLMStudioGLM47Flash30BA3B,
		PresetDevstral224B:     modelLMStudioDevstral224B,
	},
}
