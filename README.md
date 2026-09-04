# LLM Inference for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![lint](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml)
[![test](https://github.com/flexigpt/inference-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/test.yml)

`inference-go` is a normalized Go interface for language-model inference across vendor APIs, OpenAI/Anthropic/Google-compatible services, routers, and local runtimes.

It gives applications one request and response model while preserving meaningful provider and model differences through capability profiles and model presets.

- [Installation](#installation)
- [What you get](#what-you-get)
- [Supported provider presets](#supported-provider-presets)
- [Feature overview](#feature-overview)
- [What is not yet a portable feature](#what-is-not-yet-a-portable-feature)
- [Start with a preset](#start-with-a-preset)
- [Working examples](#working-examples)
- [Capability-aware behavior](#capability-aware-behavior)
- [Detailed guides](#detailed-guides)
- [Development](#development)
- [License](#license)

## Installation

Requires Go `1.26` or newer.

Install with `go get github.com/flexigpt/inference-go`.

## What you get

- One completion API through `ProviderSetAPI`
- Anthropic Messages, OpenAI Responses, OpenAI Chat Completions, and Google Generate Content API format adapters
- Provider and model presets through `modelpreset`
- Text, image, and file input where supported by the selected model
- Text and thinking streaming where the upstream API exposes it
- Reasoning controls with model-specific validation
- Structured output and JSON Schema where supported
- Function, custom, and web-search tool abstractions
- Cache-control normalization where supported
- Usage normalization for input, output, cached, and reasoning tokens where exposed
- Capability-driven warnings when a safe setting is removed
- Early errors for unsupported contract-like settings such as an unavailable output format
- Model-level capability overrides for gateways, routers, and local runtimes
- HTTP debugging through `CompletionDebugger` and `debugclient`

## Supported provider presets

The preset catalog currently includes the following providers.

| Provider preset         | Adapter family          | Typical use                                            |
| ----------------------- | ----------------------- | ------------------------------------------------------ |
| Anthropic               | Anthropic Messages      | Claude models and Anthropic-native thinking/tool flows |
| DeepSeek                | OpenAI Responses        | DeepSeek hosted models                                 |
| Google Gemini           | Google Generate Content | Gemini models, native thinking, grounding, and files   |
| Hugging Face Router     | OpenAI Chat Completions | Routed hosted models with backend-specific identities  |
| LocalAI                 | OpenAI Responses        | Local OpenAI-compatible serving                        |
| LM Studio               | OpenAI Responses        | Local desktop/server inference                         |
| llama.cpp               | OpenAI Chat Completions | Local OpenAI-compatible serving                        |
| Meta                    | OpenAI Responses        | Meta-hosted Muse Spark models                          |
| MiniMax                 | OpenAI Responses        | MiniMax hosted models                                  |
| Mistral                 | OpenAI Chat Completions | Mistral-hosted models                                  |
| Moonshot                | Anthropic Messages      | Kimi and Moonshot models                               |
| Ollama                  | Anthropic Messages      | Local Anthropic-compatible serving                     |
| OpenAI Chat Completions | OpenAI Chat Completions | OpenAI Chat models and compatible endpoints            |
| OpenAI Responses        | OpenAI Responses        | OpenAI Responses models                                |
| OpenRouter              | OpenAI Responses        | Routed models with model-specific overrides            |
| QwenCloud               | OpenAI Responses        | Qwen/DashScope compatible-mode models                  |
| SGLang                  | OpenAI Responses        | Self-hosted OpenAI-compatible serving                  |
| vLLM                    | OpenAI Responses        | Self-hosted OpenAI-compatible serving                  |
| xAI                     | OpenAI Responses        | Grok and xAI-hosted models                             |
| Xiaomi                  | OpenAI Responses        | MiMo hosted models                                     |
| Z.AI                    | OpenAI Chat Completions | Z.AI pay-as-you-go models                              |
| Z.AI Coding Plan        | OpenAI Responses        | Z.AI Coding Plan models and credentials                |

A provider adapter is not a promise that every model supports every adapter feature. The selected model preset and its effective capability resolver determine the actual request surface.

See [`modelpreset/README.md`](./modelpreset/README.md) for catalog guidance and the current provider/model grouping.

## Feature overview

| Feature               | What `inference-go` provides                                                                 |
| --------------------- | -------------------------------------------------------------------------------------------- |
| Text input and output | Supported across all adapters                                                                |
| Image input           | Supported by selected adapters and models                                                    |
| File input            | Supported by selected adapters and models                                                    |
| Text streaming        | Supported where the upstream provider streams text                                           |
| Thinking streaming    | Supported by Anthropic, OpenAI Responses, and Google Generate Content where exposed upstream |
| Reasoning controls    | Capability-aware validation and model-specific restrictions                                  |
| JSON Schema output    | Supported by selected adapters and models                                                    |
| Function tools        | Supported across the main adapters                                                           |
| Custom tools          | Supported where the selected adapter/model can represent them                                |
| Web search            | Supported through provider-native web-search or grounding features where available           |
| Cache control         | Supported where the provider/model exposes compatible cache behavior                         |
| Usage accounting      | Normalized where upstream usage data is available                                            |
| HTTP debugging        | Supported through `CompletionDebugger`                                                       |

## What is not yet a portable feature

The library intentionally does not yet provide a general cross-provider abstraction for:

- Audio input or output
- Video input or output
- Image generation or image output
- Provider-native stored conversations
- Previous-response identifiers
- Background jobs
- Prompt resources
- Uploaded-file lifecycle management
- Arbitrary provider-specific request parameters

These features may be available from individual vendors. They are not yet part of the stable normalized surface.

## Start with a preset

For most applications, start with `modelpreset`.

The typical workflow is:

1. Select a provider preset.
2. Select a model preset for that provider.
3. Add the provider connection through `AddProviderFromPreset`.
4. Set the API key through `SetProviderAPIKey`.
5. Build a model-specific resolver with `NewPresetCapabilityResolver`.
6. Pass that resolver through `FetchCompletionOptions`.

The resolver matters because providers, routers, and local servers often host models with different reasoning, tool, modality, output, or cache support.

See the executable setup example in [`internal/integration/example_preset_catalog_test.go`](./internal/integration/example_preset_catalog_test.go).

## Working examples

Use the integration examples as the current source of truth for end-to-end flows.

| Scenario                                                 | Example                                                                                                                           |
| -------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Preset lookup, provider registration, and resolver setup | [`example_preset_catalog_test.go`](./internal/integration/example_preset_catalog_test.go)                                         |
| Anthropic basic completion                               | [`example_anthropic_basic_test.go`](./internal/integration/example_anthropic_basic_test.go)                                       |
| Anthropic tools, thinking, and streaming                 | [`example_anthropic_tools_streaming_test.go`](./internal/integration/example_anthropic_tools_streaming_test.go)                   |
| OpenAI Responses basic completion                        | [`example_openai_responses_basic_test.go`](./internal/integration/example_openai_responses_basic_test.go)                         |
| OpenAI Responses tools, files, images, and streaming     | [`example_openai_responses_tools_attachments_test.go`](./internal/integration/example_openai_responses_tools_attachments_test.go) |
| Meta Muse Spark basic completion                         | [`example_meta_muse_spark_test.go`](./internal/integration/example_meta_muse_spark_test.go)                                       |
| Meta Muse Spark multi-parameter request                  | [`example_meta_muse_spark_test.go`](./internal/integration/example_meta_muse_spark_test.go)                                       |
| OpenAI Chat Completions basic completion                 | [`example_openai_chat_basic_test.go`](./internal/integration/example_openai_chat_basic_test.go)                                   |
| OpenAI Chat tools, JSON Schema, and streaming            | [`example_openai_chat_tools_websearch_stream_test.go`](./internal/integration/example_openai_chat_tools_websearch_stream_test.go) |
| Google Generate Content basic completion                 | [`example_google_genai_basic_test.go`](./internal/integration/example_google_genai_basic_test.go)                                 |
| Google function-tool round trip                          | [`example_google_genai_tools_roundtrip_test.go`](./internal/integration/example_google_genai_tools_roundtrip_test.go)             |
| Google web search, thinking, and streaming               | [`example_google_genai_websearch_streaming_test.go`](./internal/integration/example_google_genai_websearch_streaming_test.go)     |
| Capability overrides                                     | [`example_capability_override_test.go`](./internal/integration/example_capability_override_test.go)                               |

The live examples read credentials from environment variables. They remain useful implementation references even when credentials are unavailable.

## Capability-aware behavior

Each completion is validated and normalized against effective model capabilities.

Effective capabilities are usually derived from:

1. the adapter's base capability profile
2. the provider preset override
3. the model preset override
4. an optional caller override

The normalizer can:

- Reject a request when it asks for something the selected model cannot safely provide
- Remove safe optional settings and return warnings
- Narrow requests to model-specific limits
- Sanitize provider-specific reasoning history before it is replayed
- Adapt tool-output transport to the selected model

For routers and local runtimes, add a caller override only when the deployed server or model differs from catalog metadata, and test that exact deployment before relying on advanced features.

Applications own conversation history. Preserve tool-call IDs and provider-native reasoning data only for compatible continuations; prefer ordinary user and assistant text when moving between provider families.

## Detailed guides

| Topic                             | Guide                                                                                          |
| --------------------------------- | ---------------------------------------------------------------------------------------------- |
| Provider and model preset catalog | [`modelpreset/README.md`](./modelpreset/README.md)                                             |
| Anthropic Messages behavior       | [`internal/anthropicsdk/README.md`](./internal/anthropicsdk/README.md)                         |
| OpenAI Responses behavior         | [`internal/openairesponsessdk/README.md`](./internal/openairesponsessdk/README.md)             |
| OpenAI Chat Completions behavior  | [`internal/openaichatsdk/README.md`](./internal/openaichatsdk/README.md)                       |
| Google Generate Content behavior  | [`internal/googlegeneratecontentsdk/README.md`](./internal/googlegeneratecontentsdk/README.md) |

## Development

- Formatting/linting uses the repository configuration in `.golangci.yml`
- Useful scripts are available in `taskfile.yml`
- PRs are welcome
  - keep the public surface small and provider-neutral
  - avoid leaking provider SDK types into `package inference` or `spec`

## License

Copyright (c) 2026 - Present - Pankaj Pipada

All source code in this repository, unless otherwise noted, is licensed under the MIT License.
See [LICENSE](./LICENSE) for details.
