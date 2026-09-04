package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

const (
	ProviderQwen spec.ProviderName = "qwen"

	DisplayNameProviderQwen = "QwenCloud"
)

const (
	ModelNameQwen38Max           spec.ModelName = "qwen3.8-max"
	ModelNameQwen3824TA95B       spec.ModelName = "qwen3.8-2.4t-a95b"
	ModelNameQwen3827B           spec.ModelName = "qwen3.8-27b"
	ModelNameQwen37Max           spec.ModelName = "qwen3.7-max"
	ModelNameQwen37Max20260608   spec.ModelName = "qwen3.7-max-2026-06-08"
	ModelNameQwen37Max20260520   spec.ModelName = "qwen3.7-max-2026-05-20"
	ModelNameQwen3Max            spec.ModelName = "qwen3-max"
	ModelNameQwen3Max20260123    spec.ModelName = "qwen3-max-2026-01-23"
	ModelNameQwen37Plus          spec.ModelName = "qwen3.7-plus"
	ModelNameQwen37Plus20260526  spec.ModelName = "qwen3.7-plus-2026-05-26"
	ModelNameQwen37Flash         spec.ModelName = "qwen3.7-flash"
	ModelNameQwen37Flash20260715 spec.ModelName = "qwen3.7-flash-2026-07-15"
	ModelNameQwen36Plus          spec.ModelName = "qwen3.6-plus"
	ModelNameQwen36Plus20260402  spec.ModelName = "qwen3.6-plus-2026-04-02"
	ModelNameQwen36Flash         spec.ModelName = "qwen3.6-flash"
	ModelNameQwen36Flash20260416 spec.ModelName = "qwen3.6-flash-2026-04-16"
	ModelNameQwen3635BA3B        spec.ModelName = "qwen3.6-35b-a3b"
	ModelNameQwen3627B           spec.ModelName = "qwen3.6-27b"
	ModelNameQwen35Plus          spec.ModelName = "qwen3.5-plus"
	ModelNameQwen35Plus20260420  spec.ModelName = "qwen3.5-plus-2026-04-20"
	ModelNameQwen35Plus20260215  spec.ModelName = "qwen3.5-plus-2026-02-15"
	ModelNameQwen35Flash         spec.ModelName = "qwen3.5-flash"
	ModelNameQwen35Flash20260223 spec.ModelName = "qwen3.5-flash-2026-02-23"
	ModelNameQwen35397BA17B      spec.ModelName = "qwen3.5-397b-a17b"
	ModelNameQwen35122BA10B      spec.ModelName = "qwen3.5-122b-a10b"
	ModelNameQwen3527B           spec.ModelName = "qwen3.5-27b"
	ModelNameQwen3535BA3B        spec.ModelName = "qwen3.5-35b-a3b"
	ModelNameQwenPlus            spec.ModelName = "qwen-plus"
	ModelNameQwenFlash           spec.ModelName = "qwen-flash"
	ModelNameQwen3CoderPlus      spec.ModelName = "qwen3-coder-plus"
	ModelNameQwen3CoderFlash     spec.ModelName = "qwen3-coder-flash"
	ModelNameQwenPlusCharacter   spec.ModelName = "qwen-plus-character"
	ModelNameQwenFlashCharacter  spec.ModelName = "qwen-flash-character"

	ModelNameQwen3Coder30BA3BFireworksAI spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct:fireworks-ai"
	ModelNameQwen3CoderNextNovita        spec.ModelName = "Qwen/Qwen3-Coder-Next:novita"
	ModelNameQwen3635BA3BRepo            spec.ModelName = "Qwen/Qwen3.6-35B-A3B"
	ModelNameQwen3627BRepo               spec.ModelName = "Qwen/Qwen3.6-27B"
	ModelNameQwen3VL30BRepo              spec.ModelName = "qwen/qwen3-vl-30b"
	ModelNameQwen3VL30BA3BRepo           spec.ModelName = "Qwen/Qwen3-VL-30B-A3B-Instruct"
	ModelNameQwen3Coder30BA3BRepo        spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct"
	ModelNameQwen3635BA3BLocal           spec.ModelName = "qwen3.6-35b-a3b"
	ModelNameQwen3VL30BA3BLocal          spec.ModelName = "qwen3-vl-30b-a3b"
	ModelNameQwen3Coder30BA3BLocal       spec.ModelName = "qwen3-coder-30b-a3b"

	ModelNameOpenRouterQwen37Max  spec.ModelName = "qwen/qwen3.7-max"
	ModelNameOpenRouterQwen37Plus spec.ModelName = "qwen/qwen3.7-plus"

	ModelNameQwen3635BOllama     spec.ModelName = "qwen3.6:35b"
	ModelNameQwen3627BOllama     spec.ModelName = "qwen3.6:27b"
	ModelNameQwen3VL30BOllama    spec.ModelName = "qwen3-vl:30b"
	ModelNameQwen3Coder30BOllama spec.ModelName = "qwen3-coder:30b"
)

const (
	DisplayNameQwen38Max           = "Qwen3.8 Max"
	DisplayNameQwen3824TA95B       = "Qwen3.8 2.4T A95B"
	DisplayNameQwen3827B           = "Qwen3.8 27B"
	DisplayNameQwen37Max20260608   = "Qwen3.7 Max 2026-06-08"
	DisplayNameQwen37Max20260520   = "Qwen3.7 Max 2026-05-20"
	DisplayNameQwen3Max            = "Qwen3 Max"
	DisplayNameQwen3Max20260123    = "Qwen3 Max 2026-01-23"
	DisplayNameQwen37Plus20260526  = "Qwen3.7 Plus 2026-05-26"
	DisplayNameQwen37Flash         = "Qwen3.7 Flash"
	DisplayNameQwen37Flash20260715 = "Qwen3.7 Flash 2026-07-15"
	DisplayNameQwen36Plus          = "Qwen3.6 Plus"
	DisplayNameQwen36Plus20260402  = "Qwen3.6 Plus 2026-04-02"
	DisplayNameQwen36Flash         = "Qwen3.6 Flash"
	DisplayNameQwen36Flash20260416 = "Qwen3.6 Flash 2026-04-16"
	DisplayNameQwen35Plus          = "Qwen3.5 Plus"
	DisplayNameQwen35Plus20260420  = "Qwen3.5 Plus 2026-04-20"
	DisplayNameQwen35Plus20260215  = "Qwen3.5 Plus 2026-02-15"
	DisplayNameQwen35Flash         = "Qwen3.5 Flash"
	DisplayNameQwen35Flash20260223 = "Qwen3.5 Flash 2026-02-23"
	DisplayNameQwen35397BA17B      = "Qwen3.5 397B A17B"
	DisplayNameQwen35122BA10B      = "Qwen3.5 122B A10B"
	DisplayNameQwen3527B           = "Qwen3.5 27B"
	DisplayNameQwen3535BA3B        = "Qwen3.5 35B A3B"
	DisplayNameQwenPlus            = "Qwen Plus"
	DisplayNameQwenFlash           = "Qwen Flash"
	DisplayNameQwen3CoderPlus      = "Qwen3 Coder Plus"
	DisplayNameQwen3CoderFlash     = "Qwen3 Coder Flash"
	DisplayNameQwenPlusCharacter   = "Qwen Plus Character"
	DisplayNameQwenFlashCharacter  = "Qwen Flash Character"

	DisplayNameQwen3635BA3B     = "Qwen3.6 35B A3B"
	DisplayNameQwen3635B        = "Qwen3.6 35B"
	DisplayNameQwen3627B        = "Qwen3.6 27B"
	DisplayNameQwen37Max        = "Qwen3.7 Max"
	DisplayNameQwen37Plus       = "Qwen3.7 Plus"
	DisplayNameQwen3VL30BA3B    = "Qwen3-VL 30B A3B"
	DisplayNameQwen3VL30B       = "Qwen3-VL 30B"
	DisplayNameQwen3Coder30B    = "Qwen3-Coder 30B"
	DisplayNameQwen3Coder30BA3B = "Qwen3-Coder 30B A3B"
	DisplayNameQwen3CoderNext   = "Qwen3-Coder Next"

	DisplayNameQwen3Coder30BA3BFireworksAI = "Qwen3-Coder 30B A3B (Fireworks AI)"
	DisplayNameQwen3CoderNextNovita        = "Qwen3-Coder Next (Novita)"
)

const (
	PresetQwen38Max           ModelPresetID = "qwen38Max"
	PresetQwen3824TA95B       ModelPresetID = "qwen3824tA95B"
	PresetQwen3827B           ModelPresetID = "qwen3827B"
	PresetQwen37Max20260608   ModelPresetID = "qwen37Max20260608"
	PresetQwen37Max20260520   ModelPresetID = "qwen37Max20260520"
	PresetQwen3Max            ModelPresetID = "qwen3Max"
	PresetQwen3Max20260123    ModelPresetID = "qwen3Max20260123"
	PresetQwen37Plus20260526  ModelPresetID = "qwen37Plus20260526"
	PresetQwen37Flash         ModelPresetID = "qwen37Flash"
	PresetQwen37Flash20260715 ModelPresetID = "qwen37Flash20260715"
	PresetQwen36Plus          ModelPresetID = "qwen36Plus"
	PresetQwen36Plus20260402  ModelPresetID = "qwen36Plus20260402"
	PresetQwen36Flash         ModelPresetID = "qwen36Flash"
	PresetQwen36Flash20260416 ModelPresetID = "qwen36Flash20260416"
	PresetQwen3635BA3BDirect  ModelPresetID = "qwen3635bA3B"
	PresetQwen3627BDirect     ModelPresetID = "qwen3627B"
	PresetQwen35Plus          ModelPresetID = "qwen35Plus"
	PresetQwen35Plus20260420  ModelPresetID = "qwen35Plus20260420"
	PresetQwen35Plus20260215  ModelPresetID = "qwen35Plus20260215"
	PresetQwen35Flash         ModelPresetID = "qwen35Flash"
	PresetQwen35Flash20260223 ModelPresetID = "qwen35Flash20260223"
	PresetQwen35397BA17B      ModelPresetID = "qwen35397bA17B"
	PresetQwen35122BA10B      ModelPresetID = "qwen35122bA10B"
	PresetQwen3527B           ModelPresetID = "qwen3527B"
	PresetQwen3535BA3B        ModelPresetID = "qwen3535bA3B"
	PresetQwenPlus            ModelPresetID = "qwenPlus"
	PresetQwenFlash           ModelPresetID = "qwenFlash"
	PresetQwen3CoderPlus      ModelPresetID = "qwen3CoderPlus"
	PresetQwen3CoderFlash     ModelPresetID = "qwen3CoderFlash"
	PresetQwenPlusCharacter   ModelPresetID = "qwenPlusCharacter"
	PresetQwenFlashCharacter  ModelPresetID = "qwenFlashCharacter"

	PresetQwen3635BA3B     ModelPresetID = "qwen3635ba3b"
	PresetQwen3635B        ModelPresetID = "qwen3635b"
	PresetQwen3627B        ModelPresetID = "qwen3627b"
	PresetQwen37Max        ModelPresetID = "qwen37Max"
	PresetQwen37Plus       ModelPresetID = "qwen37Plus"
	PresetQwen3VL30BA3B    ModelPresetID = "qwen3vl30ba3b"
	PresetQwen3VL30B       ModelPresetID = "qwen3vl30b"
	PresetQwen3Coder30B    ModelPresetID = "qwen3coder30b"
	PresetQwen3Coder30BA3B ModelPresetID = "qwen3coder30ba3b"
	PresetQwen3CoderNext   ModelPresetID = "qwen3codernext"

	PresetQwen3Coder30BA3BFireworksAI ModelPresetID = "qwen3coder30ba3bFireworksAI"
	PresetQwen3CoderNextNovita        ModelPresetID = "qwen3codernextNovita"
)

var modelQwenQwen38Max = ModelPreset{
	ID:          PresetQwen38Max,
	Name:        ModelNameQwen38Max,
	DisplayName: DisplayNameQwen38Max,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen38Max,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen3824TA95B = ModelPreset{
	ID:          PresetQwen3824TA95B,
	Name:        ModelNameQwen3824TA95B,
	DisplayName: DisplayNameQwen3824TA95B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3824TA95B,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen3827B = ModelPreset{
	ID:          PresetQwen3827B,
	Name:        ModelNameQwen3827B,
	DisplayName: DisplayNameQwen3827B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3827B,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 131072,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen37Max = ModelPreset{
	ID:          PresetQwen37Max,
	Name:        ModelNameQwen37Max,
	DisplayName: DisplayNameQwen37Max,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Max,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen37Max20260608 = ModelPreset{
	ID:          PresetQwen37Max20260608,
	Name:        ModelNameQwen37Max20260608,
	DisplayName: DisplayNameQwen37Max20260608,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Max20260608,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen37Max20260520 = ModelPreset{
	ID:          PresetQwen37Max20260520,
	Name:        ModelNameQwen37Max20260520,
	DisplayName: DisplayNameQwen37Max20260520,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Max20260520,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen3Max = ModelPreset{
	ID:          PresetQwen3Max,
	Name:        ModelNameQwen3Max,
	DisplayName: DisplayNameQwen3Max,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Max,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen3Max20260123 = ModelPreset{
	ID:          PresetQwen3Max20260123,
	Name:        ModelNameQwen3Max20260123,
	DisplayName: DisplayNameQwen3Max20260123,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3Max20260123,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
	},
}

var modelQwenQwen37Plus = ModelPreset{
	ID:          PresetQwen37Plus,
	Name:        ModelNameQwen37Plus,
	DisplayName: DisplayNameQwen37Plus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Plus,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen37Plus20260526 = ModelPreset{
	ID:          PresetQwen37Plus20260526,
	Name:        ModelNameQwen37Plus20260526,
	DisplayName: DisplayNameQwen37Plus20260526,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Plus20260526,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen37Flash = ModelPreset{
	ID:          PresetQwen37Flash,
	Name:        ModelNameQwen37Flash,
	DisplayName: DisplayNameQwen37Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Flash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen37Flash20260715 = ModelPreset{
	ID:          PresetQwen37Flash20260715,
	Name:        ModelNameQwen37Flash20260715,
	DisplayName: DisplayNameQwen37Flash20260715,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen37Flash20260715,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen36Plus = ModelPreset{
	ID:          PresetQwen36Plus,
	Name:        ModelNameQwen36Plus,
	DisplayName: DisplayNameQwen36Plus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen36Plus,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen36Plus20260402 = ModelPreset{
	ID:          PresetQwen36Plus20260402,
	Name:        ModelNameQwen36Plus20260402,
	DisplayName: DisplayNameQwen36Plus20260402,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen36Plus20260402,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen36Flash = ModelPreset{
	ID:          PresetQwen36Flash,
	Name:        ModelNameQwen36Flash,
	DisplayName: DisplayNameQwen36Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen36Flash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen36Flash20260416 = ModelPreset{
	ID:          PresetQwen36Flash20260416,
	Name:        ModelNameQwen36Flash20260416,
	DisplayName: DisplayNameQwen36Flash20260416,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen36Flash20260416,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen3635BA3B = ModelPreset{
	ID:          PresetQwen3635BA3BDirect,
	Name:        ModelNameQwen3635BA3B,
	DisplayName: DisplayNameQwen3635BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3635BA3B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen3627B = ModelPreset{
	ID:          PresetQwen3627BDirect,
	Name:        ModelNameQwen3627B,
	DisplayName: DisplayNameQwen3627B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3627B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
				spec.ToolPolicyModeTool,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelQwenQwen35Plus = ModelPreset{
	ID:          PresetQwen35Plus,
	Name:        ModelNameQwen35Plus,
	DisplayName: DisplayNameQwen35Plus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35Plus,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35Plus20260420 = ModelPreset{
	ID:          PresetQwen35Plus20260420,
	Name:        ModelNameQwen35Plus20260420,
	DisplayName: DisplayNameQwen35Plus20260420,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35Plus20260420,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35Plus20260215 = ModelPreset{
	ID:          PresetQwen35Plus20260215,
	Name:        ModelNameQwen35Plus20260215,
	DisplayName: DisplayNameQwen35Plus20260215,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35Plus20260215,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35Flash = ModelPreset{
	ID:          PresetQwen35Flash,
	Name:        ModelNameQwen35Flash,
	DisplayName: DisplayNameQwen35Flash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35Flash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35Flash20260223 = ModelPreset{
	ID:          PresetQwen35Flash20260223,
	Name:        ModelNameQwen35Flash20260223,
	DisplayName: DisplayNameQwen35Flash20260223,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35Flash20260223,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35397BA17B = ModelPreset{
	ID:          PresetQwen35397BA17B,
	Name:        ModelNameQwen35397BA17B,
	DisplayName: DisplayNameQwen35397BA17B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35397BA17B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen35122BA10B = ModelPreset{
	ID:          PresetQwen35122BA10B,
	Name:        ModelNameQwen35122BA10B,
	DisplayName: DisplayNameQwen35122BA10B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen35122BA10B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen3527B = ModelPreset{
	ID:          PresetQwen3527B,
	Name:        ModelNameQwen3527B,
	DisplayName: DisplayNameQwen3527B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3527B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwen3535BA3B = ModelPreset{
	ID:          PresetQwen3535BA3B,
	Name:        ModelNameQwen3535BA3B,
	DisplayName: DisplayNameQwen3535BA3B,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3535BA3B,
		Stream:          true,
		MaxPromptLength: 262144,
		MaxOutputLength: 65536,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
}

var modelQwenQwenPlus = ModelPreset{
	ID:          PresetQwenPlus,
	Name:        ModelNameQwenPlus,
	DisplayName: DisplayNameQwenPlus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwenPlus,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 32768,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
				spec.ToolPolicyModeTool,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelQwenQwenFlash = ModelPreset{
	ID:          PresetQwenFlash,
	Name:        ModelNameQwenFlash,
	DisplayName: DisplayNameQwenFlash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwenFlash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 32768,
		Reasoning:       reasoningSingle(spec.ReasoningLevelHigh),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ToolCapabilities: &capabilityoverride.ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
				spec.ToolPolicyModeTool,
			},
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelQwenQwen3CoderPlus = ModelPreset{
	ID:          PresetQwen3CoderPlus,
	Name:        ModelNameQwen3CoderPlus,
	DisplayName: DisplayNameQwen3CoderPlus,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3CoderPlus,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:  new(false),
			SupportedReasoningTypes:  []spec.ReasoningType{},
			SupportedReasoningLevels: []spec.ReasoningLevel{},
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelQwenQwen3CoderFlash = ModelPreset{
	ID:          PresetQwen3CoderFlash,
	Name:        ModelNameQwen3CoderFlash,
	DisplayName: DisplayNameQwen3CoderFlash,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwen3CoderFlash,
		Stream:          true,
		MaxPromptLength: 1048576,
		MaxOutputLength: 65536,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:  new(false),
			SupportedReasoningTypes:  []spec.ReasoningType{},
			SupportedReasoningLevels: []spec.ReasoningLevel{},
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
			SupportsParallelToolCalls: new(false),
			MaxForcedTools:            new(1),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
	},
}

var modelQwenQwenPlusCharacter = ModelPreset{
	ID:          PresetQwenPlusCharacter,
	Name:        ModelNameQwenPlusCharacter,
	DisplayName: DisplayNameQwenPlusCharacter,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwenPlusCharacter,
		Stream:          true,
		MaxPromptLength: 32768,
		MaxOutputLength: 4096,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:  new(false),
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

var modelQwenQwenFlashCharacter = ModelPreset{
	ID:          PresetQwenFlashCharacter,
	Name:        ModelNameQwenFlashCharacter,
	DisplayName: DisplayNameQwenFlashCharacter,
	ModelParam: spec.ModelParam{
		Name:            ModelNameQwenFlashCharacter,
		Stream:          true,
		MaxPromptLength: 8192,
		MaxOutputLength: 4096,
		Temperature:     new(1.0),
		SystemPrompt:    "",
		Timeout:         1800,
	},
	CapabilitiesOverride: &capabilityoverride.ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{spec.ModalityTextIn},
		ModalitiesOut: []spec.Modality{spec.ModalityTextOut},
		ReasoningCapabilities: &capabilityoverride.ReasoningCapabilitiesOverride{
			SupportsReasoningConfig:  new(false),
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

var providerQwen = ProviderPreset{
	Name:                     ProviderQwen,
	DisplayName:              DisplayNameProviderQwen,
	SDKType:                  spec.ProviderSDKTypeOpenAIResponses,
	Origin:                   "https://dashscope-intl.aliyuncs.com",
	ChatCompletionPathPrefix: "/compatible-mode/v1/responses",
	APIKeyHeaderKey:          spec.DefaultAuthorizationHeaderKey,
	DefaultHeaders: map[string]string{
		spec.DefaultContentTypeHeaderKey: spec.DefaultContentTypeHeader,
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
				spec.ToolTypeWebSearch,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeNone,
				spec.ToolPolicyModeTool,
			},
			SupportsParallelToolCalls: new(false),
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
		PresetQwen38Max:           modelQwenQwen38Max,
		PresetQwen3824TA95B:       modelQwenQwen3824TA95B,
		PresetQwen3827B:           modelQwenQwen3827B,
		PresetQwen37Max:           modelQwenQwen37Max,
		PresetQwen37Max20260608:   modelQwenQwen37Max20260608,
		PresetQwen37Max20260520:   modelQwenQwen37Max20260520,
		PresetQwen3Max:            modelQwenQwen3Max,
		PresetQwen3Max20260123:    modelQwenQwen3Max20260123,
		PresetQwen37Plus:          modelQwenQwen37Plus,
		PresetQwen37Plus20260526:  modelQwenQwen37Plus20260526,
		PresetQwen37Flash:         modelQwenQwen37Flash,
		PresetQwen37Flash20260715: modelQwenQwen37Flash20260715,
		PresetQwen36Plus:          modelQwenQwen36Plus,
		PresetQwen36Plus20260402:  modelQwenQwen36Plus20260402,
		PresetQwen36Flash:         modelQwenQwen36Flash,
		PresetQwen36Flash20260416: modelQwenQwen36Flash20260416,
		PresetQwen3635BA3BDirect:  modelQwenQwen3635BA3B,
		PresetQwen3627BDirect:     modelQwenQwen3627B,
		PresetQwen35Plus:          modelQwenQwen35Plus,
		PresetQwen35Plus20260420:  modelQwenQwen35Plus20260420,
		PresetQwen35Plus20260215:  modelQwenQwen35Plus20260215,
		PresetQwen35Flash:         modelQwenQwen35Flash,
		PresetQwen35Flash20260223: modelQwenQwen35Flash20260223,
		PresetQwen35397BA17B:      modelQwenQwen35397BA17B,
		PresetQwen35122BA10B:      modelQwenQwen35122BA10B,
		PresetQwen3527B:           modelQwenQwen3527B,
		PresetQwen3535BA3B:        modelQwenQwen3535BA3B,
		PresetQwenPlus:            modelQwenQwenPlus,
		PresetQwenFlash:           modelQwenQwenFlash,
		PresetQwen3CoderPlus:      modelQwenQwen3CoderPlus,
		PresetQwen3CoderFlash:     modelQwenQwen3CoderFlash,
		PresetQwenPlusCharacter:   modelQwenQwenPlusCharacter,
		PresetQwenFlashCharacter:  modelQwenQwenFlashCharacter,
	},
}
