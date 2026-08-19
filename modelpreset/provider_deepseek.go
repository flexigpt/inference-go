package modelpreset

import "github.com/flexigpt/inference-go/spec"

const (
	ModelNameDeepSeekR18BRepo  spec.ModelName = "deepseek-ai/DeepSeek-R1-0528-Qwen3-8B"
	ModelNameDeepSeekR18BLocal spec.ModelName = "deepseek-r1-8b"

	ModelNameDeepSeekV4FlashFireworksAI spec.ModelName = "deepseek-ai/DeepSeek-V4-Flash:fireworks-ai"
	ModelNameDeepSeekV4ProFireworksAI   spec.ModelName = "deepseek-ai/DeepSeek-V4-Pro:fireworks-ai"

	ModelNameDeepSeekR18BOllama spec.ModelName = "deepseek-r1:8b"

	ModelNameOpenRouterDeepSeekV4Flash spec.ModelName = "deepseek/deepseek-v4-flash"
	ModelNameOpenRouterDeepSeekV4Pro   spec.ModelName = "deepseek/deepseek-v4-pro"
)

const (
	DisplayNameDeepSeekR18B    = "DeepSeek-R1 8B"
	DisplayNameDeepSeekV4Flash = "DeepSeek V4 Flash"
	DisplayNameDeepSeekV4Pro   = "DeepSeek V4 Pro"

	DisplayNameDeepSeekV4FlashFireworksAI = "DeepSeek V4 Flash (Fireworks AI)"
	DisplayNameDeepSeekV4ProFireworksAI   = "DeepSeek V4 Pro (Fireworks AI)"
)

const (
	PresetDeepSeekR18B    ModelPresetID = "deepseekr18b"
	PresetDeepSeekV4Flash ModelPresetID = "deepseekv4flash"
	PresetDeepSeekV4Pro   ModelPresetID = "deepseekv4pro"

	PresetDeepSeekV4FlashFireworksAI ModelPresetID = "deepseekv4flashFireworksAI"
	PresetDeepSeekV4ProFireworksAI   ModelPresetID = "deepseekv4proFireworksAI"
)
