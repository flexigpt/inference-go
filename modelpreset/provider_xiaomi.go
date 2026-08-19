package modelpreset

import "github.com/flexigpt/inference-go/spec"

const (
	ModelNameOpenRouterXiaomiMiMoV25Pro spec.ModelName = "xiaomi/mimo-v2.5-pro"
	ModelNameOpenRouterXiaomiMiMoV25    spec.ModelName = "xiaomi/mimo-v2.5"

	ModelNameMiMoV25ProDeepInfra      spec.ModelName = "XiaomiMiMo/MiMo-V2.5-Pro:deepinfra"
	ModelNameMiMoV2FlashFeatherlessAI spec.ModelName = "XiaomiMiMo/MiMo-V2-Flash:featherless-ai"
)

const (
	DisplayNameMiMoV2Flash      = "MiMo V2 Flash"
	DisplayNameMiMoV25Pro       = "MiMo V2.5 Pro"
	DisplayNameXiaomiMiMoV25    = "Xiaomi MiMo V2.5"
	DisplayNameXiaomiMiMoV25Pro = "Xiaomi MiMo V2.5 Pro"

	DisplayNameMiMoV2FlashFeatherlessAI = "MiMo V2 Flash (Featherless AI)"
	DisplayNameMiMoV25ProDeepInfra      = "MiMo V2.5 Pro (DeepInfra)"
)

const (
	PresetMiMoV2Flash      ModelPresetID = "mimov2flash"
	PresetMiMoV25Pro       ModelPresetID = "mimov25pro"
	PresetXiaomiMiMoV25    ModelPresetID = "xiaomiMiMoV25"
	PresetXiaomiMiMoV25Pro ModelPresetID = "xiaomiMiMoV25Pro"

	PresetMiMoV25ProDeepInfra      ModelPresetID = "mimov25proDeepInfra"
	PresetMiMoV2FlashFeatherlessAI ModelPresetID = "mimov2flashFeatherlessAI"
)
