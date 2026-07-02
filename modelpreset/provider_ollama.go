package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

var modelOllamaGemma426B = ModelPreset{
	ID:          PresetGemma426B,
	Name:        ModelNameGemma426BOllama,
	DisplayName: DisplayNameGemma426B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemma426BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaGemma4E4B = ModelPreset{
	ID:          PresetGemma4E4B,
	Name:        ModelNameGemma4E4BOllama,
	DisplayName: DisplayNameGemma4E4B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemma4E4BOllama,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaGPTOSS20B = ModelPreset{
	ID:          PresetGPTOSS20B,
	Name:        ModelNameGPTOSS20BOllama,
	DisplayName: DisplayNameGPTOSS20B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGPTOSS20BOllama,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(false), false, true),
}

var modelOllamaQwen3635B = ModelPreset{
	ID:          PresetQwen3635B,
	Name:        ModelNameQwen3635BOllama,
	DisplayName: DisplayNameQwen3635B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3635BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaQwen3627B = ModelPreset{
	ID:          PresetQwen3627B,
	Name:        ModelNameQwen3627BOllama,
	DisplayName: DisplayNameQwen3627B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3627BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaDeepSeekR18B = ModelPreset{
	ID:          PresetDeepSeekR18B,
	Name:        ModelNameDeepSeekR18BOllama,
	DisplayName: DisplayNameDeepSeekR18B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameDeepSeekR18BOllama,
		Stream:          true,
		MaxPromptLength: 131072,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(true), false, true),
}

var modelOllamaQwen3VL30B = ModelPreset{
	ID:          PresetQwen3VL30B,
	Name:        ModelNameQwen3VL30BOllama,
	DisplayName: DisplayNameQwen3VL30B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3VL30BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaMinistral314B = ModelPreset{
	ID:          PresetMinistral314B,
	Name:        ModelNameMinistral314BOllama,
	DisplayName: DisplayNameMinistral314B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameMinistral314BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextImageReasoning(reasoningLevels(true), false, true),
}

var modelOllamaQwen3Coder30B = ModelPreset{
	ID:          PresetQwen3Coder30B,
	Name:        ModelNameQwen3Coder30BOllama,
	DisplayName: DisplayNameQwen3Coder30B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Coder30BOllama,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     new(0.1),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnly(),
}

var modelOllamaPhi4Reasoning14B = ModelPreset{
	ID:          PresetPhi4Reasoning14B,
	Name:        ModelNamePhi4Reasoning14BOllama,
	DisplayName: DisplayNamePhi4Reasoning14B,
	ModelParam: spec.ModelParam{
		Name:            ModelNamePhi4Reasoning14BOllama,
		Stream:          true,
		MaxPromptLength: 32768,
		MaxOutputLength: 8192,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: capTextOnlyReasoning(reasoningLevels(false), false, true),
}

var providerOllama = ProviderPreset{
	Name:                     ProviderOllama,
	DisplayName:              DisplayNameProviderOllama,
	SDKType:                  spec.ProviderSDKTypeAnthropic,
	Origin:                   "http://127.0.0.1:11434",
	ChatCompletionPathPrefix: spec.DefaultAnthropicChatCompletionPrefix,
	APIKeyHeaderKey:          spec.DefaultAnthropicAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
		"accept":                         spec.DefaultContentTypeHeader,
		"anthropic-version":              "2023-06-01",
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		StopSequenceCapabilities: &capabilityoverride.StopSequenceCapabilitiesOverride{
			IsSupported:             new(true),
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
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(0),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
	ModelPresets: map[ModelPresetID]ModelPreset{
		PresetGemma426B:        modelOllamaGemma426B,
		PresetGemma4E4B:        modelOllamaGemma4E4B,
		PresetGPTOSS20B:        modelOllamaGPTOSS20B,
		PresetQwen3635B:        modelOllamaQwen3635B,
		PresetQwen3627B:        modelOllamaQwen3627B,
		PresetDeepSeekR18B:     modelOllamaDeepSeekR18B,
		PresetQwen3VL30B:       modelOllamaQwen3VL30B,
		PresetMinistral314B:    modelOllamaMinistral314B,
		PresetQwen3Coder30B:    modelOllamaQwen3Coder30B,
		PresetPhi4Reasoning14B: modelOllamaPhi4Reasoning14B,
	},
}
