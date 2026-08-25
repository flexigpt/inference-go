# Model Presets

`modelpreset` is the built-in provider and model catalog for `inference-go`.

Use it when you want reviewed defaults for provider connection settings, model identifiers, model limits, and model-specific capabilities.

The catalog is especially useful for:

- Providers with multiple model families
- Hosted routers
- Local runtimes
- OpenAI-compatible services
- Models with different reasoning or tool support
- Models that require provider-specific parameter behavior

For general provider usage, see [`../docs/using-providers.md`](../docs/using-providers.md).

## What a preset gives you

A provider preset gives you:

- Provider identity and display name
- Adapter family
- Default origin and endpoint path
- API-key header behavior
- Default request headers
- Provider-wide capability restrictions
- The provider's model preset collection

A model preset gives you:

- Provider-specific model name
- Stable preset ID
- Display name
- Default context and output limits
- Temperature or reasoning default
- Optional cache, output, and stop defaults
- Model-specific capability restrictions

A preset is not a promise that an external provider will never change. It is a reviewed default that is validated by the catalog test suite.

## Supported provider presets

| Provider            | Constant                  | Adapter family          | Model families represented                                          |
| ------------------- | ------------------------- | ----------------------- | ------------------------------------------------------------------- |
| Anthropic           | `ProviderAnthropic`       | Anthropic Messages      | Claude                                                              |
| DeepSeek            | `ProviderDeepSeek`        | OpenAI Responses        | DeepSeek V4                                                         |
| Google Gemini       | `ProviderGoogleGemini`    | Google Generate Content | Gemini and related Google models                                    |
| Hugging Face Router | `ProviderHuggingFace`     | OpenAI Chat Completions | Routed open models and backend variants                             |
| LocalAI             | `ProviderLocalAI`         | OpenAI Responses        | Local model presets                                                 |
| LM Studio           | `ProviderLMStudio`        | OpenAI Responses        | Local model presets                                                 |
| llama.cpp           | `ProviderLlamaCPP`        | OpenAI Chat Completions | Llama and Qwen local models                                         |
| Meta                | `ProviderMeta`            | OpenAI Responses        | Muse Spark                                                          |
| MiniMax             | `ProviderMiniMax`         | OpenAI Responses        | MiniMax M2 and M3                                                   |
| Mistral             | `ProviderMistral`         | OpenAI Chat Completions | Mistral and Devstral                                                |
| Moonshot            | `ProviderMoonshot`        | Anthropic Messages      | Kimi                                                                |
| Ollama              | `ProviderOllama`          | Anthropic Messages      | Ollama-tagged local models                                          |
| OpenAI Chat         | `ProviderOpenAIChat`      | OpenAI Chat Completions | GPT 4.1 and GPT-4o                                                  |
| OpenAI Responses    | `ProviderOpenAIResponses` | OpenAI Responses        | GPT-5 and GPT-OSS families                                          |
| OpenRouter          | `ProviderOpenRouter`      | OpenAI Responses        | Routed DeepSeek, Qwen, MiniMax, Kimi, GLM, Xiaomi, and other models |
| QwenCloud           | `ProviderQwen`            | OpenAI Responses        | Qwen, Qwen Coder, and Qwen Character                                |
| SGLang              | `ProviderSGLang`          | OpenAI Responses        | Local model presets                                                 |
| vLLM                | `ProviderVLLM`            | OpenAI Responses        | Local model presets                                                 |
| xAI                 | `ProviderXAI`             | OpenAI Responses        | Grok                                                                |
| Xiaomi              | `ProviderXiaomi`          | OpenAI Responses        | MiMo                                                                |
| Z.AI                | `ProviderZAI`             | OpenAI Chat Completions | GLM                                                                 |
| Z.AI Coding Plan    | `ProviderZAICodingPlan`   | OpenAI Responses        | GLM Coding Plan models                                              |

The exact available model IDs for a provider are available through `ModelPresetIDs`.

## Why model-specific capability resolution matters

A provider may support a broad adapter surface while a model supports only part of it.

Examples:

- One model can accept images while another is text-only.
- One model can use reasoning while another cannot.
- One model can use JSON Schema while another only supports text.
- One model can call tools while another cannot.
- One model can use a different reasoning budget or level range.
- One router route can differ from another route for the same base model family.

For that reason, adding a provider preset is only half of the setup.

Applications should also create a model-specific resolver through `NewPresetCapabilityResolver` and pass it with each completion request.

## Catalog safety

The catalog returns cloned values.

Applications can safely customize a returned provider or model preset for one request or one deployment without mutating the built-in catalog.

`ValidateCatalog` validates built-in or application-owned catalog data before it is used.

Validation covers high-level data integrity, including:

- Provider identity and connection fields
- Model identity consistency
- Sampling defaults
- Reasoning defaults
- Cache/output/stop defaults
- Capability override structure

## Stable IDs and model identity

Preset IDs are intended to be stable within a release line.

Applications that persist preset IDs should still plan for migrations when:

- A provider retires a model
- A model name changes
- A router backend becomes part of the required model identity
- A provider splits one product into separate routes or credential types

Routed variants and backend suffixes are intentionally treated as separate model identities when they affect behavior.

## Local and router presets

Local and router presets are useful defaults, not hardware or deployment guarantees.

Review and test the effective behavior when:

- Running a different local server version
- Loading a custom model
- Using a gateway proxy
- Selecting a specific router backend
- Enabling tools, structured output, reasoning, or multimodal input

A caller capability override is appropriate when a deployment differs from the built-in catalog.

## Adding a provider or model preset

When adding or changing catalog data:

1. Add the provider/model declaration.
2. Register the provider in the catalog.
3. Ensure every model has a temperature or reasoning default.
4. Declare known provider/model capability differences.
5. Add or update catalog validation coverage.
6. Update this provider table if provider support changes.
7. Update application overlays that consume the built-in catalog.
8. Add an integration example if the public setup workflow changes.

The catalog tests are intended to make missing registration, invalid defaults, and incomplete application integration fail early.
