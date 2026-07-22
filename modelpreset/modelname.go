package modelpreset

import "github.com/flexigpt/inference-go/spec"

const (
	ModelNameGPT56Sol   spec.ModelName = "gpt-5.6-sol"
	ModelNameGPT56Terra spec.ModelName = "gpt-5.6-terra"
	ModelNameGPT56Luna  spec.ModelName = "gpt-5.6-luna"

	ModelNameGPT55 spec.ModelName = "gpt-5.5"

	ModelNameGPT54     spec.ModelName = "gpt-5.4"
	ModelNameGPT54Mini spec.ModelName = "gpt-5.4-mini"
	ModelNameGPT54Nano spec.ModelName = "gpt-5.4-nano"

	ModelNameGPT53Codex spec.ModelName = "gpt-5.3-codex"

	ModelNameGPT52      spec.ModelName = "gpt-5.2"
	ModelNameGPT52Codex spec.ModelName = "gpt-5.2-codex"

	ModelNameGPT51         spec.ModelName = "gpt-5.1"
	ModelNameGPT51Codex    spec.ModelName = "gpt-5.1-codex"
	ModelNameGPT51CodexMax spec.ModelName = "gpt-5.1-codex-max"

	ModelNameGPT5Mini spec.ModelName = "gpt-5-mini"

	ModelNameGPT41     spec.ModelName = "gpt-4.1"
	ModelNameGPT41Mini spec.ModelName = "gpt-4.1-mini"
	ModelNameGPT4o     spec.ModelName = "gpt-4o"
	ModelNameGPT4oMini spec.ModelName = "gpt-4o-mini"
)

const (
	ModelNameClaudeFable5   spec.ModelName = "claude-fable-5"
	ModelNameClaudeOpus48   spec.ModelName = "claude-opus-4-8"
	ModelNameClaudeOpus47   spec.ModelName = "claude-opus-4-7"
	ModelNameClaudeOpus46   spec.ModelName = "claude-opus-4-6"
	ModelNameClaudeOpus45   spec.ModelName = "claude-opus-4-5-20251101"
	ModelNameClaudeOpus41   spec.ModelName = "claude-opus-4-1-20250805"
	ModelNameClaudeSonnet5  spec.ModelName = "claude-sonnet-5"
	ModelNameClaudeSonnet46 spec.ModelName = "claude-sonnet-4-6"
	ModelNameClaudeSonnet45 spec.ModelName = "claude-sonnet-4-5-20250929"
	ModelNameClaudeSonnet4  spec.ModelName = "claude-sonnet-4-20250514"
	ModelNameClaudeHaiku45  spec.ModelName = "claude-haiku-4-5-20251001"
)

const (
	ModelNameGemini36Flash     spec.ModelName = "gemini-3.6-flash"
	ModelNameGemini35Flash     spec.ModelName = "gemini-3.5-flash"
	ModelNameGemini35FlashLite spec.ModelName = "gemini-3.5-flash-lite"
	ModelNameGemini31Pro       spec.ModelName = "gemini-3.1-pro-preview"
	ModelNameGemini31FlashLite spec.ModelName = "gemini-3.1-flash-lite"
	ModelNameGemini3Flash      spec.ModelName = "gemini-3-flash-preview"
	ModelNameGemini25Flash     spec.ModelName = "gemini-2.5-flash"
	ModelNameGemini25FlashLite spec.ModelName = "gemini-2.5-flash-lite-preview-06-17"
)

const (
	ModelNameMistralMedium35 spec.ModelName = "mistral-medium-3-5"
	ModelNameMistralSmall4   spec.ModelName = "mistral-small-2603"
	ModelNameMistralLarge3   spec.ModelName = "mistral-large-2512"
	ModelNameDevstral2       spec.ModelName = "devstral-2512"
)

const (
	ModelNameGrokBuild01        spec.ModelName = "grok-build-0.1"
	ModelNameGrok45             spec.ModelName = "grok-4.5"
	ModelNameGrok43             spec.ModelName = "grok-4.3"
	ModelNameGrok42Reasoning    spec.ModelName = "grok-4.20-0309-reasoning"
	ModelNameGrok42NonReasoning spec.ModelName = "grok-4.20-0309-non-reasoning"
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

const (
	ModelNameLlama4BehemothLocal spec.ModelName = "llama4-behemoth"
	ModelNameLlama4MaverickLocal spec.ModelName = "llama4-maverick"
	ModelNameLlama4ScoutLocal    spec.ModelName = "llama4-scout"
)

const (
	ModelNameGemma426BOllama        spec.ModelName = "gemma4:26b"
	ModelNameGemma4E4BOllama        spec.ModelName = "gemma4:e4b"
	ModelNameGPTOSS20BOllama        spec.ModelName = "gpt-oss:20b"
	ModelNameQwen3635BOllama        spec.ModelName = "qwen3.6:35b"
	ModelNameQwen3627BOllama        spec.ModelName = "qwen3.6:27b"
	ModelNameDeepSeekR18BOllama     spec.ModelName = "deepseek-r1:8b"
	ModelNameQwen3VL30BOllama       spec.ModelName = "qwen3-vl:30b"
	ModelNameMinistral314BOllama    spec.ModelName = "ministral-3:14b"
	ModelNameQwen3Coder30BOllama    spec.ModelName = "qwen3-coder:30b"
	ModelNamePhi4Reasoning14BOllama spec.ModelName = "phi4-reasoning:14b"
)

const (
	ModelNameOpenRouterDeepSeekV4Flash      spec.ModelName = "deepseek/deepseek-v4-flash"
	ModelNameOpenRouterXiaomiMiMoV25        spec.ModelName = "xiaomi/mimo-v2.5"
	ModelNameOpenRouterTencentHy3Preview    spec.ModelName = "tencent/hy3-preview"
	ModelNameOpenRouterMiniMaxM3            spec.ModelName = "minimax/minimax-m3"
	ModelNameOpenRouterZAIGLM52             spec.ModelName = "z-ai/glm-5.2"
	ModelNameOpenRouterDeepSeekV4Pro        spec.ModelName = "deepseek/deepseek-v4-pro"
	ModelNameOpenRouterStep37Flash          spec.ModelName = "stepfun/step-3.7-flash"
	ModelNameOpenRouterNemotron3UltraFree   spec.ModelName = "nvidia/nemotron-3-ultra-550b-a55b:free"
	ModelNameOpenRouterPoolsideLagunaM1Free spec.ModelName = "poolside/laguna-m.1:free"
	ModelNameOpenRouterXiaomiMiMoV25Pro     spec.ModelName = "xiaomi/mimo-v2.5-pro"
	ModelNameOpenRouterNemotron3SuperFree   spec.ModelName = "nvidia/nemotron-3-super-120b-a12b:free"
	ModelNameOpenRouterMoonshotKimiK26      spec.ModelName = "moonshotai/kimi-k2.6"
	ModelNameOpenRouterQwen37Max            spec.ModelName = "qwen/qwen3.7-max"
	ModelNameOpenRouterZAIGLM51             spec.ModelName = "z-ai/glm-5.1"
	ModelNameOpenRouterMoonshotKimiK27Code  spec.ModelName = "moonshotai/kimi-k2.7-code"
	ModelNameOpenRouterQwen37Plus           spec.ModelName = "qwen/qwen3.7-plus"
	ModelNameOpenRouterMiniMaxM27           spec.ModelName = "minimax/minimax-m2.7"
	ModelNameOpenRouterMiniMaxM25Free       spec.ModelName = "minimax/minimax-m2.5:free"
)
