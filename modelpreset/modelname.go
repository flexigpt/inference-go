package modelpreset

import "github.com/flexigpt/inference-go/spec"

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
	ModelNameGemma426BA4BRepo     spec.ModelName = "google/gemma-4-26b-a4b"
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
	ModelNameGemma426BA4BLocal     spec.ModelName = "gemma4-26b-a4b"
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
