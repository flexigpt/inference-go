package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelGoogleGemini31Pro = ModelPreset{
	ID:          PresetGemini31Pro,
	Name:        ModelNameGemini31Pro,
	DisplayName: DisplayNameGemini31Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini31Pro,
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

var modelGoogleGemini35Flash = ModelPreset{
	ID:          PresetGemini35Flash,
	Name:        ModelNameGemini35Flash,
	DisplayName: DisplayNameGemini35Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini35Flash,
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

var modelGoogleGemini3Flash = ModelPreset{
	ID:          PresetGemini3Flash,
	Name:        ModelNameGemini3Flash,
	DisplayName: DisplayNameGemini3Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini3Flash,
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
	ID:          PresetGemini31FlashLite,
	Name:        ModelNameGemini31FlashLite,
	DisplayName: DisplayNameGemini31FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini31FlashLite,
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
	ID:          PresetGemini25Flash,
	Name:        ModelNameGemini25Flash,
	DisplayName: DisplayNameGemini25Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini25Flash,
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
	ID:          PresetGemini25FlashLite,
	Name:        ModelNameGemini25FlashLite,
	DisplayName: DisplayNameGemini25FlashLite,
	ModelParam: spec.ModelParam{
		Name:            ModelNameGemini25FlashLite,
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
	DisplayName:              DisplayNameProviderGoogleGemini,
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
		PresetGemini31Pro:       modelGoogleGemini31Pro,
		PresetGemini35Flash:     modelGoogleGemini35Flash,
		PresetGemini3Flash:      modelGoogleGemini3Flash,
		PresetGemini31FlashLite: modelGoogleGemini31FlashLite,
		PresetGemini25Flash:     modelGoogleGemini25Flash,
		PresetGemini25FlashLite: modelGoogleGemini25FlashLite,
	},
}
