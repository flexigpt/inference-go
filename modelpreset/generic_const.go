package modelpreset

import (
	"github.com/flexigpt/inference-go/spec"
)

const (
	ModelNameGemma426BA4BRepo  spec.ModelName = "google/gemma-4-26b-a4b"
	ModelNameGemma426BA4BLocal spec.ModelName = "gemma4-26b-a4b"
)

const (
	DisplayNameGemma426BA4B = "Gemma 4 26B A4B"
	DisplayNameGemma426B    = "Gemma 4 26B"
	DisplayNameGemma4E4B    = "Gemma 4 E4B"
)

const (
	PresetGemma426BA4B ModelPresetID = "gemma426ba4b"
	PresetGemma426B    ModelPresetID = "gemma426b"
	PresetGemma4E4B    ModelPresetID = "gemma4e4b"
)

const (
	ModelNameGPTOSS120BFireworksAI spec.ModelName = "openai/gpt-oss-120b:fireworks-ai"
	ModelNameGPTOSS20BFireworksAI  spec.ModelName = "openai/gpt-oss-20b:fireworks-ai"

	ModelNameQwen3Coder30BA3BFireworksAI    spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct:fireworks-ai"
	ModelNameGLM52FireworksAI               spec.ModelName = "zai-org/GLM-5.2:fireworks-ai"
	ModelNameDeepSeekV4FlashFireworksAI     spec.ModelName = "deepseek-ai/DeepSeek-V4-Flash:fireworks-ai"
	ModelNameDeepSeekV4ProFireworksAI       spec.ModelName = "deepseek-ai/DeepSeek-V4-Pro:fireworks-ai"
	ModelNameNemotron3UltraNVFP4FireworksAI spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4:fireworks-ai"
	ModelNameGLM51FireworksAI               spec.ModelName = "zai-org/GLM-5.1:fireworks-ai"
	ModelNameMiniMaxM27FireworksAI          spec.ModelName = "MiniMaxAI/MiniMax-M2.7:fireworks-ai"
	ModelNameGLM51FP8FireworksAI            spec.ModelName = "zai-org/GLM-5.1-FP8:fireworks-ai"

	ModelNameOrnith1035BFP8DeepInfra     spec.ModelName = "deepreinforce-ai/Ornith-1.0-35B-FP8:deepinfra"
	ModelNameMiMoV25ProDeepInfra         spec.ModelName = "XiaomiMiMo/MiMo-V2.5-Pro:deepinfra"
	ModelNameNemotron3UltraBF16DeepInfra spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16:deepinfra"

	ModelNameGLM52FP8ZAI spec.ModelName = "zai-org/GLM-5.2-FP8:zai-org"

	ModelNameQwen3CoderNextNovita     spec.ModelName = "Qwen/Qwen3-Coder-Next:novita"
	ModelNameGLM5Novita               spec.ModelName = "zai-org/GLM-5:novita"
	ModelNameMiniMaxM25Novita         spec.ModelName = "MiniMaxAI/MiniMax-M2.5:novita"
	ModelNameKimiK2InstructNovita     spec.ModelName = "moonshotai/Kimi-K2-Instruct:novita"
	ModelNameKimiK2Instruct0905Novita spec.ModelName = "moonshotai/Kimi-K2-Instruct-0905:novita"

	ModelNameNemotron3SuperBF16FeatherlessAI spec.ModelName = "nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16:featherless-ai"
	ModelNameKimiK2ThinkingFeatherlessAI     spec.ModelName = "moonshotai/Kimi-K2-Thinking:featherless-ai"
	ModelNameStep35FlashFeatherlessAI        spec.ModelName = "stepfun-ai/Step-3.5-Flash:featherless-ai"
	ModelNameMiMoV2FlashFeatherlessAI        spec.ModelName = "XiaomiMiMo/MiMo-V2-Flash:featherless-ai"

	ModelNameGLM47Cerebras spec.ModelName = "zai-org/GLM-4.7:cerebras"
)

const (
	ModelNameGPTOSS20BRepo        spec.ModelName = "openai/gpt-oss-20b"
	ModelNameQwen3635BA3BRepo     spec.ModelName = "Qwen/Qwen3.6-35B-A3B"
	ModelNameQwen3627BRepo        spec.ModelName = "Qwen/Qwen3.6-27B"
	ModelNameDeepSeekR18BRepo     spec.ModelName = "deepseek-ai/DeepSeek-R1-0528-Qwen3-8B"
	ModelNameQwen3VL30BRepo       spec.ModelName = "qwen/qwen3-vl-30b"
	ModelNameQwen3VL30BA3BRepo    spec.ModelName = "Qwen/Qwen3-VL-30B-A3B-Instruct"
	ModelNameMinistral314BRepo    spec.ModelName = "mistralai/Ministral-3-14B-Instruct-2512"
	ModelNameQwen3Coder30BA3BRepo spec.ModelName = "Qwen/Qwen3-Coder-30B-A3B-Instruct"
	ModelNameGLM47Flash30BA3BRepo spec.ModelName = "zai-org/GLM-4.7-Flash"
	ModelNamePhi4Reasoning14BRepo spec.ModelName = "microsoft/Phi-4-reasoning"
	ModelNameDevstral224BRepo     spec.ModelName = "mistralai/Devstral-Small-2-24B-Instruct-2512"
)

const (
	ModelNameGPTOSS20BLocal        spec.ModelName = "gpt-oss-20b"
	ModelNameQwen3635BA3BLocal     spec.ModelName = "qwen3.6-35b-a3b"
	ModelNameDeepSeekR18BLocal     spec.ModelName = "deepseek-r1-8b"
	ModelNameQwen3VL30BA3BLocal    spec.ModelName = "qwen3-vl-30b-a3b"
	ModelNameMinistral314BLocal    spec.ModelName = "ministral-3-14b"
	ModelNameQwen3Coder30BA3BLocal spec.ModelName = "qwen3-coder-30b-a3b"
	ModelNameGLM47Flash30BA3BLocal spec.ModelName = "glm-4.7-flash"
	ModelNamePhi4Reasoning14BLocal spec.ModelName = "phi-4-reasoning"
	ModelNameDevstral224BLocal     spec.ModelName = "devstral-small-2-24b"
)

const (
	DisplayNameGPTOSS120B = "gpt-oss 120B"
	DisplayNameGPTOSS20B  = "gpt-oss 20B"
)

const (
	DisplayNameLlama4Behemoth = "Llama 4 Behemoth"
	DisplayNameLlama4Maverick = "Llama 4 Maverick"
	DisplayNameLlama4Scout    = "Llama 4 Scout"
)

const (
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
)

const (
	DisplayNameDeepSeekR18B    = "DeepSeek-R1 8B"
	DisplayNameDeepSeekV4Flash = "DeepSeek V4 Flash"
	DisplayNameDeepSeekV4Pro   = "DeepSeek V4 Pro"
)

const (
	DisplayNameMinistral314B = "Ministral 3 14B"
)

const (
	DisplayNameGLM47Flash30BA3B = "GLM-4.7 Flash 30B A3B"
	DisplayNameGLM47            = "GLM 4.7"
	DisplayNameGLM5             = "GLM 5"
	DisplayNameGLM51            = "GLM 5.1"
	DisplayNameGLM51FP8         = "GLM 5.1 FP8"
	DisplayNameGLM52            = "GLM 5.2"
	DisplayNameGLM52FP8         = "GLM 5.2 FP8"
)

const (
	DisplayNameZAIGLM51             = "Z.AI GLM 5.1"
	DisplayNameZAIGLM52             = "Z.AI GLM 5.2"
	DisplayNamePhi4Reasoning14B     = "Phi-4 Reasoning 14B"
	DisplayNameDevstral224B         = "Devstral Small 2 24B"
	DisplayNameOrnith1035BFP8       = "Ornith 1.0 35B FP8"
	DisplayNameMiMoV2Flash          = "MiMo V2 Flash"
	DisplayNameMiMoV25Pro           = "MiMo V2.5 Pro"
	DisplayNameXiaomiMiMoV25        = "Xiaomi MiMo V2.5"
	DisplayNameXiaomiMiMoV25Pro     = "Xiaomi MiMo V2.5 Pro"
	DisplayNameMiniMaxM25           = "MiniMax M2.5"
	DisplayNameMiniMaxM25Free       = "MiniMax M2.5 Free"
	DisplayNameMiniMaxM27           = "MiniMax M2.7"
	DisplayNameMiniMaxM3            = "MiniMax M3"
	DisplayNameKimiK2Thinking       = "Kimi K2 Thinking"
	DisplayNameKimiK2Instruct       = "Kimi K2 Instruct"
	DisplayNameKimiK2Instruct0905   = "Kimi K2 Instruct 0905"
	DisplayNameMoonshotKimiK26      = "MoonshotAI Kimi K2.6"
	DisplayNameMoonshotKimiK27Code  = "MoonshotAI Kimi K2.7 Code"
	DisplayNameStep35Flash          = "Step 3.5 Flash"
	DisplayNameStep37Flash          = "Step 3.7 Flash"
	DisplayNameTencentHy3Preview    = "Tencent Hy3 Preview"
	DisplayNamePoolsideLagunaM1Free = "Poolside Laguna M.1 Free"
)

const (
	DisplayNameNemotron3UltraNVFP4 = "NVIDIA Nemotron 3 Ultra 550B A55B NVFP4"
	DisplayNameNemotron3UltraBF16  = "NVIDIA Nemotron 3 Ultra 550B A55B BF16"
	DisplayNameNemotron3SuperBF16  = "NVIDIA Nemotron 3 Super 120B A12B BF16"
	DisplayNameNemotron3UltraFree  = "NVIDIA Nemotron 3 Ultra Free"
	DisplayNameNemotron3SuperFree  = "NVIDIA Nemotron 3 Super Free"
)

const (
	PresetGPTOSS120B ModelPresetID = "gptoss120b"
	PresetGPTOSS20B  ModelPresetID = "gptoss20b"

	PresetLlama4Behemoth ModelPresetID = "llama4Behemoth"
	PresetLlama4Maverick ModelPresetID = "llama4Maverick"
	PresetLlama4Scout    ModelPresetID = "llama4Scout"

	PresetQwen3635BA3B         ModelPresetID = "qwen3635ba3b"
	PresetQwen3635B            ModelPresetID = "qwen3635b"
	PresetQwen3627B            ModelPresetID = "qwen3627b"
	PresetQwen37Max            ModelPresetID = "qwen37Max"
	PresetQwen37Plus           ModelPresetID = "qwen37Plus"
	PresetQwen3VL30BA3B        ModelPresetID = "qwen3vl30ba3b"
	PresetQwen3VL30B           ModelPresetID = "qwen3vl30b"
	PresetQwen3Coder30B        ModelPresetID = "qwen3coder30b"
	PresetQwen3Coder30BA3B     ModelPresetID = "qwen3coder30ba3b"
	PresetQwen3CoderNext       ModelPresetID = "qwen3codernext"
	PresetDeepSeekR18B         ModelPresetID = "deepseekr18b"
	PresetDeepSeekV4Flash      ModelPresetID = "deepseekv4flash"
	PresetDeepSeekV4Pro        ModelPresetID = "deepseekv4pro"
	PresetMinistral314B        ModelPresetID = "ministral314b"
	PresetGLM47Flash30BA3B     ModelPresetID = "glm47flash30ba3b"
	PresetGLM47                ModelPresetID = "glm47"
	PresetGLM5                 ModelPresetID = "glm5"
	PresetGLM51                ModelPresetID = "glm51"
	PresetGLM51FP8             ModelPresetID = "glm51fp8"
	PresetGLM52                ModelPresetID = "glm52"
	PresetGLM52FP8             ModelPresetID = "glm52fp8"
	PresetZAIGLM51             ModelPresetID = "zaiglm51"
	PresetZAIGLM52             ModelPresetID = "zaiglm52"
	PresetPhi4Reasoning14B     ModelPresetID = "phi4reasoning14b"
	PresetDevstral224B         ModelPresetID = "devstral224b"
	PresetOrnith1035BFP8       ModelPresetID = "ornith1035bfp8"
	PresetMiMoV2Flash          ModelPresetID = "mimov2flash"
	PresetMiMoV25Pro           ModelPresetID = "mimov25pro"
	PresetXiaomiMiMoV25        ModelPresetID = "xiaomiMiMoV25"
	PresetXiaomiMiMoV25Pro     ModelPresetID = "xiaomiMiMoV25Pro"
	PresetMiniMaxM25           ModelPresetID = "minimaxm25"
	PresetMiniMaxM25Free       ModelPresetID = "minimaxm25free"
	PresetMiniMaxM27           ModelPresetID = "minimaxm27"
	PresetMiniMaxM3            ModelPresetID = "minimaxM3"
	PresetKimiK2Thinking       ModelPresetID = "kimik2thinking"
	PresetKimiK2Instruct       ModelPresetID = "kimik2instruct"
	PresetKimiK2Instruct0905   ModelPresetID = "kimik2instruct0905"
	PresetMoonshotKimiK26      ModelPresetID = "moonshotKimiK26"
	PresetMoonshotKimiK27Code  ModelPresetID = "moonshotKimiK27Code"
	PresetStep35Flash          ModelPresetID = "step35flash"
	PresetStep37Flash          ModelPresetID = "step37Flash"
	PresetTencentHy3Preview    ModelPresetID = "tencentHy3Preview"
	PresetPoolsideLagunaM1Free ModelPresetID = "poolsideLagunaM1Free"

	PresetNemotron3UltraNVFP4 ModelPresetID = "nemotron3ultraNVFP4"
	PresetNemotron3UltraBF16  ModelPresetID = "nemotron3ultraBF16"
	PresetNemotron3SuperBF16  ModelPresetID = "nemotron3superBF16"
	PresetNemotron3UltraFree  ModelPresetID = "nvidiaNemotron3UltraFree"
	PresetNemotron3SuperFree  ModelPresetID = "nvidiaNemotron3SuperFree"
)

const (
	PresetGPTOSS120BFireworksAI ModelPresetID = "gptoss120bFireworksAI"
	PresetGPTOSS20BFireworksAI  ModelPresetID = "gptoss20bFireworksAI"

	PresetQwen3Coder30BA3BFireworksAI    ModelPresetID = "qwen3coder30ba3bFireworksAI"
	PresetGLM52FireworksAI               ModelPresetID = "glm52FireworksAI"
	PresetDeepSeekV4FlashFireworksAI     ModelPresetID = "deepseekv4flashFireworksAI"
	PresetDeepSeekV4ProFireworksAI       ModelPresetID = "deepseekv4proFireworksAI"
	PresetNemotron3UltraNVFP4FireworksAI ModelPresetID = "nemotron3ultraNVFP4FireworksAI"
	PresetGLM51FireworksAI               ModelPresetID = "glm51FireworksAI"
	PresetMiniMaxM27FireworksAI          ModelPresetID = "minimaxm27FireworksAI"
	PresetGLM51FP8FireworksAI            ModelPresetID = "glm51fp8FireworksAI"

	PresetOrnith1035BFP8DeepInfra     ModelPresetID = "ornith1035bfp8DeepInfra"
	PresetMiMoV25ProDeepInfra         ModelPresetID = "mimov25proDeepInfra"
	PresetNemotron3UltraBF16DeepInfra ModelPresetID = "nemotron3ultraBF16DeepInfra"

	PresetGLM52FP8ZAI ModelPresetID = "glm52fp8ZAI"

	PresetQwen3CoderNextNovita     ModelPresetID = "qwen3codernextNovita"
	PresetGLM5Novita               ModelPresetID = "glm5Novita"
	PresetMiniMaxM25Novita         ModelPresetID = "minimaxm25Novita"
	PresetKimiK2InstructNovita     ModelPresetID = "kimik2instructNovita"
	PresetKimiK2Instruct0905Novita ModelPresetID = "kimik2instruct0905Novita"

	PresetNemotron3SuperBF16FeatherlessAI ModelPresetID = "nemotron3superBF16FeatherlessAI"
	PresetKimiK2ThinkingFeatherlessAI     ModelPresetID = "kimik2thinkingFeatherlessAI"
	PresetStep35FlashFeatherlessAI        ModelPresetID = "step35flashFeatherlessAI"
	PresetMiMoV2FlashFeatherlessAI        ModelPresetID = "mimov2flashFeatherlessAI"

	PresetGLM47Cerebras ModelPresetID = "glm47Cerebras"
)

const (
	ProviderAnthropic       spec.ProviderName = "anthropic"
	ProviderLocalAI         spec.ProviderName = "localai"
	ProviderLMStudio        spec.ProviderName = "lmstudio"
	ProviderGoogleGemini    spec.ProviderName = "googlegemini"
	ProviderHuggingFace     spec.ProviderName = "huggingface"
	ProviderLlamaCPP        spec.ProviderName = "llamacpp"
	ProviderMistral         spec.ProviderName = "mistral"
	ProviderOllama          spec.ProviderName = "ollama"
	ProviderOpenAIChat      spec.ProviderName = "openai"
	ProviderOpenAIResponses spec.ProviderName = "openairesponses"
	ProviderOpenRouter      spec.ProviderName = "openrouter"
	ProviderSGLang          spec.ProviderName = "sglang"
	ProviderVLLM            spec.ProviderName = "vllm"
	ProviderXAI             spec.ProviderName = "xai"
)

const (
	DisplayNameProviderAnthropic       = "Anthropic"
	DisplayNameProviderGoogleGemini    = "Google Gemini API"
	DisplayNameProviderHuggingFace     = "Hugging Face"
	DisplayNameProviderLlamaCPP        = "llama.cpp"
	DisplayNameProviderLMStudio        = "LM Studio"
	DisplayNameProviderLocalAI         = "LocalAI"
	DisplayNameProviderMistral         = "Mistral AI"
	DisplayNameProviderOllama          = "Ollama"
	DisplayNameProviderOpenAIChat      = "OpenAI Chat Completions API"
	DisplayNameProviderOpenAIResponses = "OpenAI Responses API"
	DisplayNameProviderOpenRouter      = "OpenRouter"
	DisplayNameProviderSGLang          = "SGLang"
	DisplayNameProviderVLLM            = "vLLM"
	DisplayNameProviderXAI             = "xAI"
)

var catalogProviders = map[spec.ProviderName]ProviderPreset{
	ProviderAnthropic:       providerAnthropic,
	ProviderLocalAI:         providerLocalAI,
	ProviderLMStudio:        providerLMStudio,
	ProviderGoogleGemini:    providerGoogleGemini,
	ProviderHuggingFace:     providerHuggingFace,
	ProviderLlamaCPP:        providerLlamaCPP,
	ProviderMistral:         providerMistral,
	ProviderOllama:          providerOllama,
	ProviderOpenAIChat:      providerOpenAIChat,
	ProviderOpenAIResponses: providerOpenAIResponses,
	ProviderOpenRouter:      providerOpenRouter,
	ProviderSGLang:          providerSGLang,
	ProviderVLLM:            providerVLLM,
	ProviderXAI:             providerXAI,
}
