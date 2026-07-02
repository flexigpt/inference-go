package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var modelOpenRouterDeepSeekV4Flash = ModelPreset{
	ID:          PresetDeepSeekV4Flash,
	Name:        ModelNameOpenRouterDeepSeekV4Flash,
	DisplayName: DisplayNameDeepSeekV4Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterDeepSeekV4Flash,
		Stream:          true,
		MaxPromptLength: 1048575,
		MaxOutputLength: 1048575,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle: new(true),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterXiaomiMiMoV25 = ModelPreset{
	ID:          PresetXiaomiMiMoV25,
	Name:        ModelNameOpenRouterXiaomiMiMoV25,
	DisplayName: DisplayNameXiaomiMiMoV25,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterXiaomiMiMoV25,
		Stream:          true,
		MaxPromptLength: 32000,
		MaxOutputLength: 32000,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityAudioIn,
			spec.ModalityImageIn,
			spec.ModalityVideoIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterTencentHy3Preview = ModelPreset{
	ID:          PresetTencentHy3Preview,
	Name:        ModelNameOpenRouterTencentHy3Preview,
	DisplayName: DisplayNameTencentHy3Preview,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterTencentHy3Preview,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 262144,
		Temperature:     new(0.9),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
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
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterMiniMaxM3 = ModelPreset{
	ID:          PresetMiniMaxM3,
	Name:        ModelNameOpenRouterMiniMaxM3,
	DisplayName: DisplayNameMiniMaxM3,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMiniMaxM3,
		Stream:          true,
		MaxPromptLength: 524288,
		MaxOutputLength: 512000,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityVideoIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterZAIGLM52 = ModelPreset{
	ID:          PresetZAIGLM52,
	Name:        ModelNameOpenRouterZAIGLM52,
	DisplayName: DisplayNameZAIGLM52,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterZAIGLM52,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 32768,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle: new(true),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterDeepSeekV4Pro = ModelPreset{
	ID:          PresetDeepSeekV4Pro,
	Name:        ModelNameOpenRouterDeepSeekV4Pro,
	DisplayName: DisplayNameDeepSeekV4Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterDeepSeekV4Pro,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 384000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
			},
			SupportsSummaryStyle: new(true),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterStep37Flash = ModelPreset{
	ID:          PresetStep37Flash,
	Name:        ModelNameOpenRouterStep37Flash,
	DisplayName: DisplayNameStep37Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterStep37Flash,
		Stream:          true,
		MaxPromptLength: 256000,
		MaxOutputLength: 256000,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityVideoIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterNemotron3UltraFree = ModelPreset{
	ID:          PresetNemotron3UltraFree,
	Name:        ModelNameOpenRouterNemotron3UltraFree,
	DisplayName: DisplayNameNemotron3UltraFree,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterNemotron3UltraFree,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
			},
			SupportsSummaryStyle: new(true),
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
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterPoolsideLagunaM1Free = ModelPreset{
	ID:          PresetPoolsideLagunaM1Free,
	Name:        ModelNameOpenRouterPoolsideLagunaM1Free,
	DisplayName: DisplayNamePoolsideLagunaM1Free,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterPoolsideLagunaM1Free,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 32768,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterXiaomiMiMoV25Pro = ModelPreset{
	ID:          PresetXiaomiMiMoV25Pro,
	Name:        ModelNameOpenRouterXiaomiMiMoV25Pro,
	DisplayName: DisplayNameXiaomiMiMoV25Pro,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterXiaomiMiMoV25Pro,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Temperature:     new(1.0),
		Reasoning:       nil,
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterNemotron3SuperFree = ModelPreset{
	ID:          PresetNemotron3SuperFree,
	Name:        ModelNameOpenRouterNemotron3SuperFree,
	DisplayName: DisplayNameNemotron3SuperFree,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterNemotron3SuperFree,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 262144,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelMedium),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
			},
			SupportsSummaryStyle: new(true),
		},
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterMoonshotKimiK26 = ModelPreset{
	ID:          PresetMoonshotKimiK26,
	Name:        ModelNameOpenRouterMoonshotKimiK26,
	DisplayName: DisplayNameMoonshotKimiK26,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMoonshotKimiK26,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 262144,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
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
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterQwen37Max = ModelPreset{
	ID:          PresetQwen37Max,
	Name:        ModelNameOpenRouterQwen37Max,
	DisplayName: DisplayNameQwen37Max,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterQwen37Max,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterZAIGLM51 = ModelPreset{
	ID:          PresetZAIGLM51,
	Name:        ModelNameOpenRouterZAIGLM51,
	DisplayName: DisplayNameZAIGLM51,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterZAIGLM51,
		Stream:          true,
		MaxPromptLength: 65536,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterMoonshotKimiK27Code = ModelPreset{
	ID:          PresetMoonshotKimiK27Code,
	Name:        ModelNameOpenRouterMoonshotKimiK27Code,
	DisplayName: DisplayNameMoonshotKimiK27Code,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMoonshotKimiK27Code,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 16384,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
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
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterQwen37Plus = ModelPreset{
	ID:          PresetQwen37Plus,
	Name:        ModelNameOpenRouterQwen37Plus,
	DisplayName: DisplayNameQwen37Plus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterQwen37Plus,
		Stream:          true,
		MaxPromptLength: 1000000,
		MaxOutputLength: 65536,
		Temperature:     nil,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
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
			TemperatureDisallowedWhenEnabled: new(false),
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
		OutputCapabilities: &capabilityoverride.OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(false),
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
	},
}

var modelOpenRouterMiniMaxM27 = ModelPreset{
	ID:          PresetMiniMaxM27,
	Name:        ModelNameOpenRouterMiniMaxM27,
	DisplayName: DisplayNameMiniMaxM27,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMiniMaxM27,
		Stream:          true,
		MaxPromptLength: 180000,
		MaxOutputLength: 125000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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

var modelOpenRouterMiniMaxM25Free = ModelPreset{
	ID:          PresetMiniMaxM25Free,
	Name:        ModelNameOpenRouterMiniMaxM25Free,
	DisplayName: DisplayNameMiniMaxM25Free,
	ModelParam: spec.ModelParam{
		Name:            ModelNameOpenRouterMiniMaxM25Free,
		Stream:          true,
		MaxPromptLength: 100000,
		MaxOutputLength: 8000,
		Temperature:     new(1.0),
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			TemperatureDisallowedWhenEnabled: new(false),
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

var providerOpenRouter = ProviderPreset{
	Name:                     ProviderOpenRouter,
	DisplayName:              DisplayNameProviderOpenRouter,
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
		PresetDeepSeekV4Flash:      modelOpenRouterDeepSeekV4Flash,
		PresetXiaomiMiMoV25:        modelOpenRouterXiaomiMiMoV25,
		PresetTencentHy3Preview:    modelOpenRouterTencentHy3Preview,
		PresetMiniMaxM3:            modelOpenRouterMiniMaxM3,
		PresetZAIGLM52:             modelOpenRouterZAIGLM52,
		PresetDeepSeekV4Pro:        modelOpenRouterDeepSeekV4Pro,
		PresetStep37Flash:          modelOpenRouterStep37Flash,
		PresetNemotron3UltraFree:   modelOpenRouterNemotron3UltraFree,
		PresetPoolsideLagunaM1Free: modelOpenRouterPoolsideLagunaM1Free,
		PresetXiaomiMiMoV25Pro:     modelOpenRouterXiaomiMiMoV25Pro,
		PresetNemotron3SuperFree:   modelOpenRouterNemotron3SuperFree,
		PresetMoonshotKimiK26:      modelOpenRouterMoonshotKimiK26,
		PresetQwen37Max:            modelOpenRouterQwen37Max,
		PresetZAIGLM51:             modelOpenRouterZAIGLM51,
		PresetMoonshotKimiK27Code:  modelOpenRouterMoonshotKimiK27Code,
		PresetQwen37Plus:           modelOpenRouterQwen37Plus,
		PresetMiniMaxM27:           modelOpenRouterMiniMaxM27,
		PresetMiniMaxM25Free:       modelOpenRouterMiniMaxM25Free,
	},
}
