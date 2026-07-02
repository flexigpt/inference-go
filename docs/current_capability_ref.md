# API capability, preset provider, and normalization done and pending reference

This document is the current implementation reference for normalized API capability support across wire adapters, provider presets, hosted routers, and local runtimes.

Terminology used here:

- Done
  - implemented today in code
  - includes both request normalization and response normalization
- Dropped with warning
  - accepted at the normalized layer, then removed during normalization
  - warning appears in `FetchCompletionResponse.Warnings`
- Sanitized
  - provider adapter removes or transforms provider-incompatible history/content before building the vendor request
- Pending
  - not implemented yet, partially implemented, or intentionally deferred until the normalized surface is clarified

Primary capability sources:

- `spec/capability.go`
- `internal/anthropicsdk/capability.go`
- `internal/openairesponsessdk/capability.go`
- `internal/openaichatsdk/capability.go`
- `internal/googlegeneratecontentsdk/capability.go`
- `capabilityoverride`
- `modelpreset`

## Table of contents <!-- omit from toc -->

- [Cross-provider Reference](#cross-provider-reference)
  - [Done](#done)
  - [Pending](#pending)
- [Anthropic Messages API Reference](#anthropic-messages-api-reference)
  - [Done Anthropic Reference](#done-anthropic-reference)
  - [Pending Anthropic Reference](#pending-anthropic-reference)
- [OpenAI Responses API Reference](#openai-responses-api-reference)
  - [Done OpenAI Responses Reference](#done-openai-responses-reference)
  - [Pending OpenAI Responses Reference](#pending-openai-responses-reference)
- [OpenAI Chat Completions API Reference](#openai-chat-completions-api-reference)
  - [Done OpenAI Chat Completions Reference](#done-openai-chat-completions-reference)
  - [Pending OpenAI Chat Completions Reference](#pending-openai-chat-completions-reference)
- [Google Generate Content API Reference](#google-generate-content-api-reference)
  - [Done Google Generate Content Reference](#done-google-generate-content-reference)
  - [Pending Google Generate Content Reference](#pending-google-generate-content-reference)
- [Preset Provider Catalog Reference](#preset-provider-catalog-reference)
  - [Done Presets Reference](#done-presets-reference)
  - [Pending Presets Reference](#pending-presets-reference)
- [Cross-provider Pending Backlog Reference](#cross-provider-pending-backlog-reference)

## Cross-provider Reference

### Done

- Provider registration and dispatch
  - Anthropic Messages
  - OpenAI Responses
  - OpenAI Chat Completions
  - Google Generate Content
  - Mistral through OpenAI Chat-compatible adapter presets
  - Hugging Face Router through OpenAI Chat-compatible adapter presets
  - llama.cpp through OpenAI Chat-compatible adapter presets
  - xAI through OpenAI Responses-compatible adapter presets
  - OpenRouter through OpenAI Responses-compatible adapter presets
  - LocalAI through OpenAI Responses-compatible adapter presets
  - LM Studio through OpenAI Responses-compatible adapter presets
  - SGLang through OpenAI Responses-compatible adapter presets
  - vLLM through OpenAI Responses-compatible adapter presets
  - Ollama through Anthropic-compatible adapter presets

- Normalized request/response model
  - `spec.FetchCompletionRequest`
  - `spec.FetchCompletionResponse`
  - `spec.InputUnion`
  - `spec.OutputUnion`
  - `spec.ModelParam`
  - `spec.ToolChoice`
  - `spec.ToolPolicy`

- Streaming
  - normalized text streaming
  - normalized thinking streaming where the provider exposes it

- Usage normalization
  - input tokens
  - output tokens
  - cached token accounting where exposed
  - reasoning token accounting where exposed

- Capability-driven normalization
  - request cloning
  - modality validation
  - reasoning validation
  - stop-sequence normalization
  - output-format validation
  - tool capability validation
  - cache-control normalization
  - parameter dialect overrides where declared
  - warnings returned in `FetchCompletionResponse.Warnings`

- Capability override layering
  - SDK/provider base capabilities
  - provider preset overrides
  - model preset overrides
  - caller/user overrides

- Preset catalog
  - provider connection defaults
  - provider-level capability overrides
  - model default params
  - model-level capability overrides

- Debugging
  - `CompletionDebugger`
  - `debugclient.HTTPCompletionDebugger`

### Pending

- Audio input/output normalization
- Video input/output normalization
- Image output normalization
- Richer cross-provider citation normalization beyond URL/basic grounding forms
- Safe allowlisted passthrough for `ModelParam.AdditionalParametersRawJSON`
- Full capability-vs-adapter enforcement for modalities that can be represented in presets but are not yet normalized end to end
  - examples: audio and video input advertised by some gateway model metadata
- Explicit capability signal for `ToolPolicy.DisableParallel`
  - today this is not represented separately in the capability model
  - some providers can disable parallel tool calls, some cannot, and docs must currently describe that per provider

## Anthropic Messages API Reference

Capability source:

- `internal/anthropicsdk/capability.go`

Adapter references:

- `internal/anthropicsdk/api_anthropic_messages.go`
- `internal/anthropicsdk/input_processing.go`
- `internal/anthropicsdk/thinking.go`
- `internal/anthropicsdk/cache_control.go`

### Done Anthropic Reference

- Message/input normalization
  - user input messages
  - assistant output messages
  - reasoning messages
  - function/custom tool calls
  - function/custom tool outputs
  - web-search tool calls
  - web-search tool outputs
  - top-level system prompt

- Modalities
  - text input
  - image input
  - file input
  - text output

- Streaming
  - text streaming
  - thinking streaming

- Reasoning/thinking
  - reasoning request config
  - signed thinking history pass-back
  - redacted thinking history pass-back
  - thinking override logic for tool-result turn boundaries
  - reasoning output normalization

- Output controls
  - text output format
  - JSON Schema output format
  - verbosity mapped to Anthropic effort

- Stop sequences
  - supported and normalized

- Tools
  - function tools
  - custom tools
  - web search tool
  - tool policy: `auto`, `any`, `tool`, `none`
  - forced single-tool constraint via `MaxForcedTools = 1`
  - parallel tool usage supported

- Cache control
  - top-level cache control
  - input/output content cache control
  - tool choice cache control
  - tool call cache control
  - tool output cache control

- Citations
  - URL citations normalized from supported Anthropic citation data

- Usage
  - input tokens
  - cached input tokens
  - output tokens

### Pending Anthropic Reference

- File/document handling
  - plain-text `text/*` file document mapping is still pending
  - current adapter skips non-PDF base64 file content for Anthropic documents

- Cache control
  - reasoning-content cache control is not normalized

- Citations
  - richer citation location variants beyond current URL-style normalization

- Output modalities
  - image output
  - audio/video output

- Safe passthrough candidates
  - `metadata.user_id`
  - `service_tier`
  - `top_k`

## OpenAI Responses API Reference

Capability source:

- `internal/openairesponsessdk/capability.go`

Adapter references:

- `internal/openairesponsessdk/api_openai_responses.go`
- `internal/openairesponsessdk/thinking.go`
- `internal/openairesponsessdk/cache_control.go`

### Done OpenAI Responses Reference

- Message/input normalization
  - user input messages
  - assistant output messages
  - reasoning messages
  - function/custom tool calls
  - function/custom tool outputs
  - web-search tool calls
  - system prompt via top-level instructions

- Modalities
  - text input
  - image input
  - file input
  - text output

- Streaming
  - text streaming
  - thinking/reasoning streaming

- Reasoning/thinking
  - reasoning config via effort/summary
  - reasoning output normalization
  - encrypted reasoning input history pass-back
  - reasoning history sanitization to encrypted-only form

- Output controls
  - text output format
  - JSON Schema output format
  - verbosity

- Tools
  - function tools
  - custom tools
  - web search tool
  - tool policy: `auto`, `any`, `tool`, `none`
  - parallel tool calls supported

- Cache control
  - top-level prompt cache key
  - top-level prompt cache retention (`in-memory`, `24h`)
  - per-message/per-item cache controls dropped with warning

- Citations
  - URL citations normalized

- Usage
  - input tokens
  - cached input tokens
  - output tokens
  - reasoning tokens

### Pending OpenAI Responses Reference

- Stop sequences
  - unsupported by Responses API
  - dropped with warning

- Tool definitions
  - custom tool definitions are currently emitted as function tools in the adapter

- Web-search result normalization
  - structured web-search result output is limited
  - many search results surface through text/citations rather than a full round-trippable tool-output object

- Stateful/provider-native features intentionally not normalized
  - `previous_response_id`
  - `store`
  - `background`
  - conversation state
  - prompt objects

- Safe passthrough candidates
  - `include`
  - `truncation`
  - `service_tier`
  - `metadata`
  - safety/user identifiers
  - stream options

## OpenAI Chat Completions API Reference

Capability source:

- `internal/openaichatsdk/capability.go`

Adapter references:

- `internal/openaichatsdk/api_openai_chat_completions.go`
- `internal/openaichatsdk/cache_control.go`

### Done OpenAI Chat Completions Reference

- Message/input normalization
  - user input messages
  - assistant output messages
  - function/custom tool calls
  - function/custom tool outputs
  - top-level system prompt
  - `gpt-5*` / `o*` system-prompt role normalization to `developer`

- Modalities
  - text input
  - image input
  - file input
  - text output

- Streaming
  - text streaming

- Reasoning config
  - reasoning effort config only

- Output controls
  - text output format
  - JSON Schema output format
  - verbosity

- Stop sequences
  - supported up to 4
  - normalization/truncation handled by capability layer

- Tools
  - function tools
  - custom tools
  - tool policy: `auto`, `any`, `tool`, `none`
  - `DisableParallel` mapped to `parallel_tool_calls=false`
  - web search support via top-level `web_search_options`

- Cache control
  - top-level prompt cache key
  - top-level prompt cache retention (`in-memory`, `24h`)
  - per-message/per-item cache controls dropped with warning

- Citations
  - URL citations normalized from annotations

- Usage
  - input tokens
  - cached input tokens
  - output tokens
  - reasoning tokens where exposed by the API

### Pending OpenAI Chat Completions Reference

- Reasoning history
  - structured reasoning input/output messages are not supported by Chat Completions
  - reasoning messages are dropped/sanitized out

- Streaming thinking
  - unsupported by Chat Completions API

- File/tool-output richness
  - tool outputs are effectively text-only in the normalized Chat round-trip path
  - image/file tool output items are not preserved as structured tool outputs when sent back into Chat Completions

- Tool definitions
  - custom tool definitions are currently emitted as function tools in the adapter

- Web search semantics
  - Chat Completions web search is not a normal tool call
  - forcing web search via `toolPolicy.mode=tool` is not a true cross-provider equivalent

- Safe passthrough candidates
  - `service_tier`
  - safety/user identifiers
  - message `name`
  - model-specific penalties/logprobs/seed/logit-bias where normalization is later desired

## Google Generate Content API Reference

Capability source:

- `internal/googlegeneratecontentsdk/capability.go`

Adapter references:

- `internal/googlegeneratecontentsdk/api_google_genai.go`
- `internal/googlegeneratecontentsdk/input_processing.go`
- `internal/googlegeneratecontentsdk/thinking.go`

### Done Google Generate Content Reference

- Message/input normalization
  - user input messages
  - assistant output messages
  - reasoning messages
  - function/custom tool calls
  - function/custom tool outputs
  - top-level system prompt via `SystemInstruction`

- Modalities
  - text input
  - image input
  - file input
  - text output

- Streaming
  - text streaming
  - thinking streaming

- Reasoning/thinking
  - token-budget reasoning config
  - level-based reasoning config
  - Google-native signed thought history pass-back
  - reasoning output normalization, including thought signature preservation

- Output controls
  - text output format
  - JSON Schema output format via raw schema payload

- Stop sequences
  - supported and normalized

- Tools
  - function tools
  - custom tools
  - Google Search grounding
  - tool policy: `auto`, `any`, `tool`, `none` for callable tools
  - web search grounding normalized into synthetic web-search tool call/output entries

- Cache control
  - unsupported
  - normalized request cache controls are dropped with warning

- Usage
  - input tokens
  - cached input tokens
  - output tokens
  - reasoning tokens

### Pending Google Generate Content Reference

- Tool parallelism
  - `ToolPolicy.DisableParallel` is not currently normalized/enforced for Google Generate Content
  - this should eventually be represented explicitly in capabilities or warning behavior

- JSON Schema subfield normalization
  - current adapter forwards the raw schema object
  - `JSONSchemaParam.Name`
  - `JSONSchemaParam.Description`
  - `JSONSchemaParam.Strict`
    are not currently mapped

- Tool output history richness
  - function/custom tool output history is effectively text-only
  - image/file tool output items are not preserved in Gemini function-response history

- Web search grounding options
  - grounding is supported
  - provider-specific per-tool knobs are only partially normalized today

- Citations
  - grounding is not yet attached back onto `ContentItemText.Citations`
  - current normalized form is synthetic web-search call/output items

- Candidate handling
  - first candidate only is normalized today

- Output modalities
  - image output is not normalized
  - audio/video output is not normalized

## Preset Provider Catalog Reference

This section covers provider presets that reuse one of the normalized wire adapters and then apply provider/model capability overrides through `modelpreset`.

Capability sources:

- `modelpreset`
- `capabilityoverride`
- the selected SDK adapter base capability file

Adapter mapping:

| Preset provider     | Provider constant                     | Wire adapter                       | Capability notes                                                                                              |
| ------------------- | ------------------------------------- | ---------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| Anthropic           | `modelpreset.ProviderAnthropic`       | Anthropic Messages                 | Provider preset extends Anthropic base caps with catalog defaults and model-specific reasoning/cache behavior |
| OpenAI Responses    | `modelpreset.ProviderOpenAIResponses` | OpenAI Responses                   | Provider preset plus per-model reasoning-level restrictions                                                   |
| OpenAI Chat         | `modelpreset.ProviderOpenAIChat`      | OpenAI Chat Completions            | Provider preset plus model-level no-reasoning overrides for selected non-reasoning models                     |
| Google Gemini       | `modelpreset.ProviderGoogleGemini`    | Google Generate Content            | Provider preset plus per-model reasoning and token-budget restrictions                                        |
| Mistral             | `modelpreset.ProviderMistral`         | OpenAI Chat Completions-compatible | Provider override narrows modalities/tools/cache and selects Mistral parameter dialect                        |
| Hugging Face Router | `modelpreset.ProviderHuggingFace`     | OpenAI Chat Completions-compatible | Routed backend suffixes are distinct model identities; model overrides restrict reasoning where known         |
| llama.cpp           | `modelpreset.ProviderLlamaCPP`        | OpenAI Chat Completions-compatible | Local OpenAI-compatible preset with model defaults                                                            |
| xAI                 | `modelpreset.ProviderXAI`             | OpenAI Responses-compatible        | Provider/model overrides describe xAI reasoning, encrypted reasoning, cache, tool, and output behavior        |
| OpenRouter          | `modelpreset.ProviderOpenRouter`      | OpenAI Responses-compatible        | Model overrides are central; routed models vary in modalities, output formats, tools, and reasoning levels    |
| LocalAI             | `modelpreset.ProviderLocalAI`         | OpenAI Responses-compatible        | Local preset with broad provider caps and per-model modality/reasoning overrides                              |
| LM Studio           | `modelpreset.ProviderLMStudio`        | OpenAI Responses-compatible        | Local preset with per-model modality/reasoning overrides                                                      |
| Ollama              | `modelpreset.ProviderOllama`          | Anthropic-compatible               | Local preset with constrained tool policy and Anthropic-style request adapter                                 |
| SGLang              | `modelpreset.ProviderSGLang`          | OpenAI Responses-compatible        | Self-hosted preset with per-model modality/reasoning overrides                                                |
| vLLM                | `modelpreset.ProviderVLLM`            | OpenAI Responses-compatible        | Self-hosted preset with per-model modality/reasoning overrides                                                |

### Done Presets Reference

- Provider preset registration
  - `ProviderPreset.Name`
  - `ProviderPreset.DisplayName`
  - `ProviderPreset.SDKType`
  - `ProviderPreset.Origin`
  - `ProviderPreset.ChatCompletionPathPrefix`
  - `ProviderPreset.APIKeyHeaderKey`
  - `ProviderPreset.DefaultHeaders`
  - `ProviderPreset.CapabilitiesOverride`
  - `ProviderPreset.ModelPresets`

- Model preset registration
  - `ModelPreset.ID`
  - `ModelPreset.Name`
  - `ModelPreset.DisplayName`
  - `ModelPreset.ModelParam`
  - `ModelPreset.CapabilitiesOverride`

- Capability resolver support
  - provider preset overrides
  - model preset overrides
  - cloned presets returned by catalog APIs
  - derived per-completion capabilities through `ProviderSetAPI.NewPresetCapabilityResolver`

- Distinct routed model identities
  - Hugging Face routed backend suffixes such as `:fireworks-ai`, `:deepinfra`, `:novita`, `:featherless-ai`, and `:cerebras`
  - OpenRouter routed model names with suffixes such as `:free`
  - display names that include backend/router-specific distinctions where the backend is part of the effective model identity

- Hosted API presets
  - Mistral provider preset
    - text and image input
    - text output
    - reasoning config for supported models
    - function tools
    - tool policies `auto`, `any`, `tool`, `none`
    - cache disabled
    - Mistral parameter dialect override
  - xAI provider preset
    - text and image input
    - text output
    - reasoning levels
    - encrypted reasoning support where declared
    - function and web-search tools
    - top-level ephemeral cache support where declared

- Hosted router presets
  - OpenRouter provider preset
    - OpenAI Responses-compatible adapter
    - model-level modality overrides
    - model-level reasoning-level overrides
    - model-level output-format overrides
    - model-level tool and parallel-tool-call overrides
    - stop sequences disabled at provider preset level
  - Hugging Face Router provider preset
    - OpenAI Chat-compatible adapter
    - routed backend model names
    - routed backend-specific preset IDs and display names
    - model-level reasoning overrides for known reasoning models

- Local and self-hosted runtime presets
  - LocalAI
    - OpenAI Responses-compatible adapter
    - text/image/file provider capabilities
    - per-model modality and reasoning overrides
    - stop sequences disabled
  - LM Studio
    - OpenAI Responses-compatible adapter
    - text/image provider capabilities
    - per-model modality and reasoning overrides
    - stop sequences disabled
  - llama.cpp
    - OpenAI Chat-compatible adapter
    - local origin defaults
    - model defaults for local serving
  - Ollama
    - Anthropic-compatible adapter
    - text/image provider capabilities
    - constrained tool policy
    - model-level reasoning overrides
  - SGLang
    - OpenAI Responses-compatible adapter
    - text/image provider capabilities
    - per-model modality and reasoning overrides
    - stop sequences disabled
  - vLLM
    - OpenAI Responses-compatible adapter
    - text/image provider capabilities
    - per-model modality and reasoning overrides
    - stop sequences disabled

- Shared clean catalog constants
  - provider-agnostic display names for the same base model identity
  - provider-agnostic preset IDs where the model identity is the same
  - backend-specific preset IDs and display names where routed backend identity matters

### Pending Presets Reference

- Runtime discovery
  - presets are static catalog data
  - no automatic discovery of local server model lists
  - no automatic probing of hosted router model capabilities

- Router/local capability drift
  - hosted routers can change model support without this package changing
  - local runtimes can vary by version, build flags, and loaded model
  - callers may still need user overrides for deployments with known differences

- Audio/video end-to-end support
  - capability structs can represent audio/video modalities
  - some gateway model presets may include audio/video modality metadata
  - normalized audio/video request and response conversion is still pending cross-provider work

- Full local runtime parity
  - local runtimes often implement only subsets of OpenAI-compatible or Anthropic-compatible APIs
  - behavior around tools, JSON Schema, reasoning, and streaming can vary significantly

- Provider-specific advanced parameters
  - safe allowlisted passthrough for `ModelParam.AdditionalParametersRawJSON` is still pending
  - router-specific controls are intentionally not normalized yet

## Cross-provider Pending Backlog Reference

These are the main remaining normalized-surface items that affect more than one provider. Pending:

- Richer citation abstraction
  - beyond URL citations
  - grounding/page/block/offset variants where a stable cross-provider model makes sense

- Output metadata promotion
  - small stable normalized response metadata such as response ID, model, finish/stop reason where safe

- Additional top-level controls
  - `TopP`
  - penalties
  - logprobs
  - seed/logit bias
  - only when a clean normalized model is chosen

- Safe allowlisted passthrough
  - implement `ModelParam.AdditionalParametersRawJSON` as provider-specific allowlisted merge

- Broader multimodal normalization
  - audio
  - video
  - image generation / image output

- Explicit stateful-feature policy
  - continue keeping provider-native stateful conversation/storage/file ecosystems out of scope unless a deliberate normalized design is introduced
