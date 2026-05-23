package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

const (
	PresetGoogleGemini31Pro       ModelPresetID = "gemini31Pro"
	PresetGoogleGemini3Flash      ModelPresetID = "gemini3Flash"
	PresetGoogleGemini31FlashLite ModelPresetID = "gemini31FlashLite"
	PresetGoogleGemini25Flash     ModelPresetID = "gemini25Flash"
	PresetGoogleGemini25FlashLite ModelPresetID = "gemini25FlashLite"
)

const (
	ModelNameGoogleGemini31Pro       spec.ModelName = "gemini-3.1-pro-preview"
	ModelNameGoogleGemini3Flash      spec.ModelName = "gemini-3-flash-preview"
	ModelNameGoogleGemini31FlashLite spec.ModelName = "gemini-3.1-flash-lite-preview"
	ModelNameGoogleGemini25Flash     spec.ModelName = "gemini-2.5-flash"
	ModelNameGoogleGemini25FlashLite spec.ModelName = "gemini-2.5-flash-lite-preview-06-17"
)

var modelGoogleGemini31Pro = ModelPreset{
	ID:          PresetGoogleGemini31Pro,
	Name:        ModelNameGoogleGemini31Pro,
	DisplayName: "GoogleAI Gemini 3.1 Pro",
	ModelParam: spec.ModelParam{
		Name:            ModelNameGoogleGemini31Pro,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
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
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini3Flash = ModelPreset{
	ID:          PresetGoogleGemini3Flash,
	Name:        ModelNameGoogleGemini3Flash,
	DisplayName: "GoogleAI Gemini 3 Flash",
	ModelParam: spec.ModelParam{
		Name:            ModelNameGoogleGemini3Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini31FlashLite = ModelPreset{
	ID:          PresetGoogleGemini31FlashLite,
	Name:        ModelNameGoogleGemini31FlashLite,
	DisplayName: "GoogleAI Gemini 3.1 Flash Lite",
	ModelParam: spec.ModelParam{
		Name:            ModelNameGoogleGemini31FlashLite,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
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
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
		},
	},
}

var modelGoogleGemini25Flash = ModelPreset{
	ID:          PresetGoogleGemini25Flash,
	Name:        ModelNameGoogleGemini25Flash,
	DisplayName: "GoogleAI Gemini 2.5 Flash",
	ModelParam: spec.ModelParam{
		Name:            ModelNameGoogleGemini25Flash,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1),
				MaxAllowed:      new(24576),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
		},
	},
}

var modelGoogleGemini25FlashLite = ModelPreset{
	ID:          PresetGoogleGemini25FlashLite,
	Name:        ModelNameGoogleGemini25FlashLite,
	DisplayName: "GoogleAI Gemini 2.5 Flash Lite",
	ModelParam: spec.ModelParam{
		Name:            ModelNameGoogleGemini25FlashLite,
		Stream:          true,
		MaxPromptLength: 200000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningHybrid(1024),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(512),
				MaxAllowed:      new(24576),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
		},
	},
}

var providerGoogleGemini = ProviderPreset{
	Name:                     ProviderGoogleGemini,
	DisplayName:              "Google Gemini API",
	SDKType:                  spec.ProviderSDKTypeGoogleGenerateContent,
	Origin:                   spec.DefaultGoogleGenerateContentOrigin,
	ChatCompletionPathPrefix: spec.DefaultGoogleGenerateContentPrefix,
	APIKeyHeaderKey:          spec.DefaultGoogleGenerateContentAPIKeyHeaderKey,
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
				spec.ReasoningTypeHybridWithTokens,
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
			HybridTokenBudgetCapabilities: &capabilityoverride.ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1),
				MaxAllowed:      new(32768),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
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
		PresetGoogleGemini31Pro:       modelGoogleGemini31Pro,
		PresetGoogleGemini3Flash:      modelGoogleGemini3Flash,
		PresetGoogleGemini31FlashLite: modelGoogleGemini31FlashLite,
		PresetGoogleGemini25Flash:     modelGoogleGemini25Flash,
		PresetGoogleGemini25FlashLite: modelGoogleGemini25FlashLite,
	},
}
