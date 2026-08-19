package modelpreset

import (
	"github.com/flexigpt/inference-go/spec"
)

const (
	ModelNamePhi4Reasoning14BRepo  spec.ModelName = "microsoft/Phi-4-reasoning"
	ModelNamePhi4Reasoning14BLocal spec.ModelName = "phi-4-reasoning"

	ModelNamePhi4Reasoning14BOllama spec.ModelName = "phi4-reasoning:14b"

	DisplayNamePhi4Reasoning14B = "Phi-4 Reasoning 14B"

	PresetPhi4Reasoning14B ModelPresetID = "phi4reasoning14b"
)

const (
	ModelNameOrnith1035BFP8DeepInfra spec.ModelName = "deepreinforce-ai/Ornith-1.0-35B-FP8:deepinfra"

	DisplayNameOrnith1035BFP8          = "Ornith 1.0 35B FP8"
	DisplayNameOrnith1035BFP8DeepInfra = "Ornith 1.0 35B FP8 (DeepInfra)"

	PresetOrnith1035BFP8          ModelPresetID = "ornith1035bfp8"
	PresetOrnith1035BFP8DeepInfra ModelPresetID = "ornith1035bfp8DeepInfra"
)

const (
	ModelNameStep35FlashFeatherlessAI spec.ModelName = "stepfun-ai/Step-3.5-Flash:featherless-ai"
	ModelNameOpenRouterStep37Flash    spec.ModelName = "stepfun/step-3.7-flash"

	DisplayNameStep35Flash              = "Step 3.5 Flash"
	DisplayNameStep37Flash              = "Step 3.7 Flash"
	DisplayNameStep35FlashFeatherlessAI = "Step 3.5 Flash (Featherless AI)"

	PresetStep35Flash              ModelPresetID = "step35flash"
	PresetStep37Flash              ModelPresetID = "step37Flash"
	PresetStep35FlashFeatherlessAI ModelPresetID = "step35flashFeatherlessAI"
)

const (
	ModelNameOpenRouterTencentHy3Preview spec.ModelName = "tencent/hy3-preview"

	DisplayNameTencentHy3Preview = "Tencent Hy3 Preview"

	PresetTencentHy3Preview ModelPresetID = "tencentHy3Preview"
)

const (
	ModelNameOpenRouterPoolsideLagunaM1Free spec.ModelName = "poolside/laguna-m.1:free"

	DisplayNamePoolsideLagunaM1Free = "Poolside Laguna M.1 Free"

	PresetPoolsideLagunaM1Free ModelPresetID = "poolsideLagunaM1Free"
)

const (
	ModelNameNemotron3UltraNVFP4FireworksAI  spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4:fireworks-ai"
	ModelNameNemotron3UltraBF16DeepInfra     spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16:deepinfra"
	ModelNameNemotron3SuperBF16FeatherlessAI spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16:featherless-ai"
	ModelNameOpenRouterNemotron3SuperFree    spec.ModelName = "nvidia/nemotron-3-super-120b-a12b:free"
	ModelNameOpenRouterNemotron3UltraFree    spec.ModelName = "nvidia/nemotron-3-ultra-550b-a55b:free"

	DisplayNameNemotron3UltraNVFP4             = "NVIDIA Nemotron 3 Ultra 550B A55B NVFP4"
	DisplayNameNemotron3UltraBF16              = "NVIDIA Nemotron 3 Ultra 550B A55B BF16"
	DisplayNameNemotron3SuperBF16              = "NVIDIA Nemotron 3 Super 120B A12B BF16"
	DisplayNameNemotron3UltraFree              = "NVIDIA Nemotron 3 Ultra Free"
	DisplayNameNemotron3SuperFree              = "NVIDIA Nemotron 3 Super Free"
	DisplayNameNemotron3SuperBF16FeatherlessAI = "NVIDIA Nemotron 3 Super 120B A12B BF16 (Featherless AI)"
	DisplayNameNemotron3UltraBF16DeepInfra     = "NVIDIA Nemotron 3 Ultra 550B A55B BF16 (DeepInfra)"
	DisplayNameNemotron3UltraNVFP4FireworksAI  = "NVIDIA Nemotron 3 Ultra 550B A55B NVFP4 (Fireworks AI)"

	PresetNemotron3UltraNVFP4             ModelPresetID = "nemotron3ultraNVFP4"
	PresetNemotron3UltraBF16              ModelPresetID = "nemotron3ultraBF16"
	PresetNemotron3SuperBF16              ModelPresetID = "nemotron3superBF16"
	PresetNemotron3UltraFree              ModelPresetID = "nvidiaNemotron3UltraFree"
	PresetNemotron3SuperFree              ModelPresetID = "nvidiaNemotron3SuperFree"
	PresetNemotron3UltraNVFP4FireworksAI  ModelPresetID = "nemotron3ultraNVFP4FireworksAI"
	PresetNemotron3UltraBF16DeepInfra     ModelPresetID = "nemotron3ultraBF16DeepInfra"
	PresetNemotron3SuperBF16FeatherlessAI ModelPresetID = "nemotron3superBF16FeatherlessAI"
)
