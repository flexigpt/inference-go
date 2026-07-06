# LLM Inference for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![lint](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml)
[![test](https://github.com/flexigpt/inference-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/test.yml)

A single normalized Go interface for LLM inference across vendor APIs, OpenAI-compatible gateways, and local runtimes, using official SDKs where available and normalized adapter compatibility elsewhere.

- [Features at a glance](#features-at-a-glance)
- [Installation](#installation)
- [Quickstart](#quickstart)
- [Examples](#examples)
- [Provider configuration](#provider-configuration)
- [Supported providers](#supported-providers)
  - [Anthropic Messages API](#anthropic-messages-api)
  - [OpenAI Responses API](#openai-responses-api)
  - [OpenAI Chat Completions API](#openai-chat-completions-api)
  - [Google Generate Content API](#google-generate-content-api)
  - [Mistral AI API](#mistral-ai-api)
  - [xAI API](#xai-api)
  - [OpenRouter](#openrouter)
  - [Hugging Face Router](#hugging-face-router)
  - [OpenAI-compatible local and self-hosted runtimes](#openai-compatible-local-and-self-hosted-runtimes)
- [Model presets](#model-presets)
- [Model capabilities and normalization](#model-capabilities-and-normalization)
  - [Capability overrides](#capability-overrides)
- [HTTP debugging](#http-debugging)
- [Notes](#notes)
- [Development](#development)
- [License](#license)

## Features at a glance

- Single normalized interface via `ProviderSetAPI`

- Normalized wire adapters today:
  - Anthropic Messages API via `github.com/anthropics/anthropic-sdk-go`
  - OpenAI Responses API via `github.com/openai/openai-go/v3`
  - OpenAI Chat Completions API via `github.com/openai/openai-go/v3`
  - Google Generate Content API via `google.golang.org/genai`

- Runtime provider presets today:
  - Anthropic
  - OpenAI Responses
  - OpenAI Chat Completions
  - Google Gemini
  - Mistral
  - xAI
  - OpenRouter
  - Hugging Face Router
  - LocalAI, LM Studio, llama.cpp, Ollama, SGLang, and vLLM

- Common preset mappings:
  - Anthropic and Ollama presets use the Anthropic-compatible adapter.
  - OpenAI Chat, Hugging Face Router, Mistral, and llama.cpp presets use the OpenAI Chat Completions-compatible adapter.
  - OpenAI Responses, xAI, OpenRouter, LocalAI, LM Studio, SGLang, and vLLM presets use the OpenAI Responses-compatible adapter.
  - Google Gemini presets use the Google Generate Content adapter.

- Normalized request/response model in `spec/`:
  - text, image, and file input content
  - assistant/user/tool/reasoning content
  - function/custom/web-search tool definitions and tool calls
  - structured output and verbosity controls
  - reasoning/thinking controls
  - streaming events for text and thinking
  - usage accounting
  - cache-control normalization where supported

- Request normalization before provider calls:
  - capability-driven validation and safe dropping of unsupported features
  - provider/model-specific parameter dialect selection where declared by capabilities
  - warnings returned in `FetchCompletionResponse.Warnings`
  - per-model capability override support through `FetchCompletionOptions.CapabilityResolver`
  - preset-based provider and model capability overrides through `modelpreset`

- Streaming:
  - text streaming for supported providers
  - thinking/reasoning streaming where the provider exposes it

- Debugging:
  - pluggable `CompletionDebugger`
  - built-in HTTP debugger in `debugclient`

## Installation

```bash
# Go 1.26+
go get github.com/flexigpt/inference-go
```

## Quickstart

Basic flow:

1. Create a `ProviderSetAPI`.
2. Register one or more providers with `AddProvider`
   - The easiest path is to use a predefined vendor specific `modelpreset`, which contains provider connection defaults, model defaults, and per-provider/per-model capability overrides.
3. Set each provider API key with `SetProviderAPIKey`
4. Call `FetchCompletion`

## Examples

Basic:

```go
ctx := context.Background()
ps, err := inference.NewProviderSetAPI()
if err != nil {
    return err
}
providerPreset, err := modelpreset.Provider(modelpreset.ProviderOpenAIResponses)
if err != nil {
    return err
}
modelPreset, err := modelpreset.Model(modelpreset.ProviderOpenAIResponses, modelpreset.PresetGPT5Mini)
if err != nil {
    return err
}
if _, err := ps.AddProviderFromPreset(ctx, providerPreset.Name, providerPreset); err != nil {
    return err
}
if err := ps.SetProviderAPIKey(ctx, providerPreset.Name, os.Getenv("OPENAI_API_KEY")); err != nil {
    return err
}

completionKey := string(modelPreset.ID)
resolver, err := ps.NewPresetCapabilityResolver(
    ctx,
    providerPreset.Name,
    providerPreset,
    modelPreset,
    completionKey,
)
if err != nil {
    return err
}

modelParam := modelPreset.ModelParam
modelParam.Stream = false
modelParam.MaxOutputLength = 2048
modelParam.SystemPrompt = "You are concise."
resp, err := ps.FetchCompletion(ctx, providerPreset.Name, &spec.FetchCompletionRequest{
    ModelParam: modelParam,
    Inputs: []spec.InputUnion{{
        Kind: spec.InputKindInputMessage,
        InputMessage: &spec.InputOutputContent{
            Role: spec.RoleUser,
            Contents: []spec.InputOutputContentItemUnion{{
                Kind: spec.ContentItemKindText,
                TextItem: &spec.ContentItemText{Text: "Say hello in one sentence."},
            }},
        },
    }},
}, &spec.FetchCompletionOptions{
    CompletionKey:      completionKey,
    CapabilityResolver: resolver,
})
if err != nil {
    return err
}
_ = resp
```

Available repository examples:

- Anthropic
  - [Basic Anthropic call](internal/integration/example_anthropic_basic_test.go)
  - [Anthropic tools + streaming + reasoning](internal/integration/example_anthropic_tools_streaming_test.go)

- OpenAI
  - [Basic OpenAI Chat Completions](internal/integration/example_openai_chat_basic_test.go)
  - [OpenAI Responses basic](internal/integration/example_openai_responses_basic_test.go)
  - [OpenAI Responses tools + attachments + streaming](internal/integration/example_openai_responses_tools_attachments_test.go)
  - [OpenAI Chat tools + JSON Schema + streaming](internal/integration/example_openai_chat_tools_websearch_stream_test.go)

- Google
  - [Google Generate Content basic](internal/integration/example_google_genai_basic_test.go)
  - [Google Generate Content function-tool round trip](internal/integration/example_google_genai_tools_roundtrip_test.go)
  - [Google Generate Content web search + thinking + streaming](internal/integration/example_google_genai_websearch_streaming_test.go)

- Preset-backed providers
  - The same `AddProviderFromPreset` and `NewPresetCapabilityResolver` flow works for Mistral, xAI, OpenRouter, Hugging Face Router, LocalAI, LM Studio, llama.cpp, Ollama, SGLang, and vLLM.
  - These providers reuse the normalized wire adapters listed above and rely on provider/model preset capability overrides for provider-specific behavior.

- [Capability override example (get provider caps, override per-model)](./internal/integration/example_capability_override_test.go)

## Provider configuration

Providers are registered dynamically with `ProviderSetAPI.AddProvider`.

```go
type AddProviderConfig struct {
    SDKType                  spec.ProviderSDKType
    Origin                   string
    ChatCompletionPathPrefix string
    APIKeyHeaderKey          string
    DefaultHeaders           map[string]string
}
```

Fields:

- `SDKType`
  - Selects the normalized wire adapter, not necessarily the public provider brand.
  - `spec.ProviderSDKTypeAnthropic`
  - `spec.ProviderSDKTypeOpenAIChatCompletions`
  - `spec.ProviderSDKTypeOpenAIResponses`
  - `spec.ProviderSDKTypeGoogleGenerateContent`

- `Origin`
  - Required
  - Base origin for the provider or gateway/proxy
  - May point at a hosted vendor API, a hosted router, or a local runtime.

- `ChatCompletionPathPrefix`
  - Optional generic path prefix
  - Historical field name, reused across providers
  - Useful when routing through a gateway path prefix
  - Adapters trim built-in endpoint suffixes when needed:
    - Anthropic: trailing `v1/messages`
    - OpenAI Chat: trailing `chat/completions`
    - OpenAI Responses: trailing `responses`

- `APIKeyHeaderKey`
  - Optional override for non-standard gateway auth headers

- `DefaultHeaders`
  - Optional extra headers added to every request

## Supported providers

`ProviderSetAPI` supports four normalized wire adapters. The `modelpreset` package then supplies ready-to-use provider presets for hosted vendors, hosted routers, and local runtimes.

| Preset provider         | Provider constant                     | Wire adapter                       | Notes                                                                                                          |
| ----------------------- | ------------------------------------- | ---------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Anthropic               | `modelpreset.ProviderAnthropic`       | Anthropic Messages                 | Official Anthropic SDK adapter                                                                                 |
| OpenAI Responses        | `modelpreset.ProviderOpenAIResponses` | OpenAI Responses                   | Official OpenAI SDK adapter                                                                                    |
| OpenAI Chat Completions | `modelpreset.ProviderOpenAIChat`      | OpenAI Chat Completions            | Official OpenAI SDK adapter                                                                                    |
| Google Gemini           | `modelpreset.ProviderGoogleGemini`    | Google Generate Content            | Official Google GenAI SDK adapter                                                                              |
| Mistral                 | `modelpreset.ProviderMistral`         | OpenAI Chat Completions-compatible | Uses Mistral API origin with model-specific capability overrides                                               |
| xAI                     | `modelpreset.ProviderXAI`             | OpenAI Responses-compatible        | Uses xAI API origin with model-specific reasoning overrides                                                    |
| OpenRouter              | `modelpreset.ProviderOpenRouter`      | OpenAI Responses-compatible        | Router presets include model-level modality, output, reasoning, and tool overrides                             |
| Hugging Face Router     | `modelpreset.ProviderHuggingFace`     | OpenAI Chat Completions-compatible | Routed backend suffixes such as `:fireworks-ai` and `:featherless-ai` are treated as distinct model identities |
| LocalAI                 | `modelpreset.ProviderLocalAI`         | OpenAI Responses-compatible        | Local/server-compatible preset with local model defaults                                                       |
| LM Studio               | `modelpreset.ProviderLMStudio`        | OpenAI Responses-compatible        | Local OpenAI-compatible preset                                                                                 |
| llama.cpp               | `modelpreset.ProviderLlamaCPP`        | OpenAI Chat Completions-compatible | Local OpenAI-compatible preset                                                                                 |
| Ollama                  | `modelpreset.ProviderOllama`          | Anthropic-compatible               | Local Anthropic-compatible preset                                                                              |
| SGLang                  | `modelpreset.ProviderSGLang`          | OpenAI Responses-compatible        | Self-hosted OpenAI-compatible preset                                                                           |
| vLLM                    | `modelpreset.ProviderVLLM`            | OpenAI Responses-compatible        | Self-hosted OpenAI-compatible preset                                                                           |

Capability support is derived from:

1. the selected wire adapter base capabilities,
2. the provider preset override,
3. the model preset override,
4. any caller-supplied override.

### Anthropic Messages API

| Area                  | Support | Notes                                                                 |
| --------------------- | ------- | --------------------------------------------------------------------- |
| Text input/output     | yes     | User/assistant messages normalized                                    |
| Streaming text        | yes     |                                                                       |
| Reasoning/thinking    | yes     | Signed thinking and redacted thinking supported                       |
| Streaming thinking    | yes     | Redacted thinking is not streamed                                     |
| Output format         | yes     | text and `jsonSchema`                                                 |
| Output verbosity      | yes     | maps to Anthropic effort                                              |
| Stop sequences        | yes     | maps to `stop_sequences`                                              |
| Images input          | yes     | base64 or URL                                                         |
| Files input           | partial | PDFs supported; plain-text file document mapping is still pending     |
| Function/custom tools | yes     |                                                                       |
| Web search            | yes     | server-side web search tool and result blocks                         |
| Tool policy           | yes     | `auto`, `any`, `tool`, `none`                                         |
| Cache control         | partial | top-level, input/output content, tool choice, tool call, tool output  |
| Citations             | partial | URL citations normalized                                              |
| Usage                 | yes     | input/output/cached; no explicit reasoning token count from Anthropic |

Normalization notes:

- reasoning input history keeps Anthropic-compatible signed/redacted reasoning only
- if an interleaved tool-result turn requires Anthropic thinking to be enabled/disabled, the adapter applies the needed override
- tool-result ordering is normalized for Anthropic’s strict tool-use/tool-result turn rules

### OpenAI Responses API

| Area                  | Support | Notes                                                           |
| --------------------- | ------- | --------------------------------------------------------------- |
| Text input/output     | yes     |                                                                 |
| Streaming text        | yes     |                                                                 |
| Reasoning/thinking    | yes     | config + reasoning output items                                 |
| Streaming thinking    | yes     | reasoning summary and reasoning text deltas                     |
| Output format         | yes     | text and `jsonSchema`                                           |
| Output verbosity      | yes     |                                                                 |
| Stop sequences        | no      | dropped with warning by normalization                           |
| Images input          | yes     | base64 or URL                                                   |
| Files input           | yes     | base64 or URL                                                   |
| Function/custom tools | yes     | custom tool definitions are currently emitted as function tools |
| Web search            | yes     | built-in web search tool                                        |
| Tool policy           | yes     | `auto`, `any`, `tool`, `none`                                   |
| Cache control         | partial | top-level prompt cache only                                     |
| Citations             | yes     | URL citations normalized                                        |
| Usage                 | yes     | input/output/cached/reasoning                                   |

Normalization notes:

- reasoning input history is sanitized to OpenAI-compatible encrypted reasoning only
- if no encrypted reasoning input exists, reasoning history items are dropped
- stateful Responses features like `previous_response_id` and provider-side storage are intentionally not normalized

### OpenAI Chat Completions API

| Area                      | Support | Notes                                                                 |
| ------------------------- | ------- | --------------------------------------------------------------------- |
| Text input/output         | yes     | first choice only is surfaced                                         |
| Streaming text            | yes     |                                                                       |
| Reasoning config          | yes     | reasoning effort only                                                 |
| Streaming thinking        | no      | API does not expose separate reasoning stream                         |
| Reasoning message history | no      | dropped by adapter                                                    |
| Output format             | yes     | text and `jsonSchema`                                                 |
| Output verbosity          | yes     | `max` maps to `high`                                                  |
| Stop sequences            | yes     | up to 4                                                               |
| Images input              | yes     | base64 data URL or remote URL                                         |
| Files input               | partial | embedded file data only                                               |
| Function/custom tools     | yes     | custom tool definitions are currently emitted as function tools       |
| Web search                | yes     | via top-level `web_search_options`, not as a normal tool call         |
| Tool policy               | yes     | `auto`, `any`, `tool`, `none`                                         |
| Cache control             | partial | top-level prompt cache only                                           |
| Citations                 | yes     | URL citations from annotations                                        |
| Usage                     | yes     | input/output/cached/reasoning                                         |
| System prompt role        | yes     | sent as `developer` for `o*` / `gpt-5*` model families, else `system` |

Normalization notes:

- reasoning message inputs are dropped because Chat Completions does not support structured reasoning history
- tool outputs are normalized back in as text-only tool messages
- web search forcing semantics differ from function tools because Chat Completions exposes web search as top-level request options, not as a standard tool call

### Google Generate Content API

| Area                  | Support | Notes                                                                                                                             |
| --------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Text input/output     | yes     | first candidate only is surfaced                                                                                                  |
| Streaming text        | yes     |                                                                                                                                   |
| Reasoning/thinking    | yes     | config + Google-native signed thought history; signatures on assistant text and function-tool-call parts are preserved for replay |
| Streaming thinking    | yes     | streams thought text when exposed by the API                                                                                      |
| Output format         | partial | text and `jsonSchema`; currently only the raw schema payload is forwarded                                                         |
| Output verbosity      | no      | dropped with warning by normalization                                                                                             |
| Stop sequences        | yes     | normalized up to capability max                                                                                                   |
| Images input          | yes     | inline bytes or URI                                                                                                               |
| Files input           | yes     | inline bytes or URI                                                                                                               |
| Function/custom tools | yes     | custom tool definitions are emitted as function declarations                                                                      |
| Web search            | yes     | Google Search grounding normalized as web-search call/output                                                                      |
| Tool policy           | partial | `auto`, `any`, `tool`, `none` for callable tools; web search cannot be forced as a callable tool                                  |
| Cache control         | no      | dropped with warning by normalization                                                                                             |
| Citations             | partial | grounding is normalized as web-search tool outputs, not attached to text citations yet                                            |
| Usage                 | yes     | input/output/cached/reasoning                                                                                                     |

Normalization notes:

- reasoning input history keeps only valid Google-native signed thoughts
- non-Google reasoning history is sanitized out before request conversion
- assistant text/tool-call signatures emitted by Gemini are preserved and passed back on follow-up turns
- function tool output history is currently text-only
- `ToolPolicy.DisableParallel` is not currently normalized for Google Generate Content

### Mistral AI API

Mistral presets use the OpenAI Chat Completions-compatible adapter with Mistral-specific connection defaults and capability overrides.

| Area               | Support | Notes                                                                         |
| ------------------ | ------- | ----------------------------------------------------------------------------- |
| Text input/output  | yes     | Via OpenAI Chat-compatible request/response shape                             |
| Streaming text     | yes     |                                                                               |
| Reasoning config   | partial | Presets advertise Mistral-supported reasoning levels where applicable         |
| Streaming thinking | no      | OpenAI Chat-compatible response path does not expose separate thinking stream |
| Output format      | yes     | text and `jsonSchema` where model/provider capabilities allow                 |
| Output verbosity   | no      | Mistral provider override disables verbosity                                  |
| Stop sequences     | yes     |                                                                               |
| Images input       | yes     | Provider preset advertises text and image input                               |
| Files input        | no      | Not enabled by the Mistral provider preset                                    |
| Function tools     | yes     |                                                                               |
| Custom tools       | no      | Provider preset only advertises function tools                                |
| Web search         | no      | Not enabled by the Mistral provider preset                                    |
| Tool policy        | yes     | `auto`, `any`, `tool`, `none`                                                 |
| Cache control      | no      | Provider preset disables automatic/top-level cache controls                   |
| Usage              | yes     | Subject to what the OpenAI Chat-compatible response exposes                   |

Normalization notes:

- Mistral uses a provider-specific parameter dialect for `max_tokens`.
- Some Mistral models expose reasoning through a provider-specific reasoning configuration; presets model this through capability overrides.

### xAI API

xAI presets use the OpenAI Responses-compatible adapter with xAI-specific connection defaults and model-level reasoning overrides.

| Area               | Support | Notes                                                                             |
| ------------------ | ------- | --------------------------------------------------------------------------------- |
| Text input/output  | yes     | Via OpenAI Responses-compatible request/response shape                            |
| Streaming text     | yes     |                                                                                   |
| Reasoning/thinking | partial | Presets declare reasoning levels and encrypted reasoning support where applicable |
| Streaming thinking | partial | Depends on what the xAI Responses-compatible endpoint emits                       |
| Output format      | yes     | text and `jsonSchema`                                                             |
| Output verbosity   | no      | xAI provider override disables verbosity                                          |
| Stop sequences     | no      | xAI provider preset disables stop sequences                                       |
| Images input       | yes     | Provider preset advertises text and image input                                   |
| Files input        | no      | Not enabled by the xAI provider preset                                            |
| Function tools     | yes     |                                                                                   |
| Web search         | yes     | Provider preset advertises web search                                             |
| Tool policy        | yes     | `auto`, `any`, `tool`, `none`                                                     |
| Cache control      | partial | Ephemeral top-level cache key support where declared                              |
| Usage              | yes     | Subject to what the Responses-compatible endpoint exposes                         |

Normalization notes:

- Some xAI model presets explicitly disable normalized reasoning config even when the provider-wide preset supports it.
- Model-level capability overrides should be used through `NewPresetCapabilityResolver`.

### OpenRouter

OpenRouter presets use the OpenAI Responses-compatible adapter. The model presets are intentionally detailed because routed OpenRouter models differ significantly in modalities, reasoning levels, JSON Schema support, and tool support.

| Area               | Support | Notes                                                                                                            |
| ------------------ | ------- | ---------------------------------------------------------------------------------------------------------------- |
| Text input/output  | yes     | Via OpenAI Responses-compatible request/response shape                                                           |
| Streaming text     | yes     |                                                                                                                  |
| Reasoning/thinking | partial | Model presets declare supported reasoning levels and summary support                                             |
| Streaming thinking | partial | Depends on what the routed model/provider emits                                                                  |
| Output format      | partial | Some model presets support text only; others support text and `jsonSchema`                                       |
| Output verbosity   | partial | Provider-wide preset allows it, many model presets disable it                                                    |
| Stop sequences     | no      | OpenRouter preset disables stop sequences                                                                        |
| Images input       | partial | Model-specific                                                                                                   |
| Files input        | partial | Provider-wide preset allows files; model-specific overrides may narrow modalities                                |
| Audio/video input  | pending | Capability metadata can represent these modalities, but cross-provider audio/video normalization remains pending |
| Function tools     | partial | Model-specific                                                                                                   |
| Custom tools       | partial | Provider-wide preset allows custom tools; many model presets narrow to function tools                            |
| Web search         | partial | Provider-wide preset advertises web search; routed model behavior can vary                                       |
| Tool policy        | yes     | `auto`, `any`, `tool`, `none` where tools are enabled                                                            |
| Cache control      | no      | Not enabled by the OpenRouter preset                                                                             |
| Usage              | yes     | Subject to what OpenRouter returns for the routed model                                                          |

Normalization notes:

- OpenRouter model presets should be treated as model-specific contracts. Do not assume provider-wide capabilities apply unchanged to every routed model.
- Models with `:free` suffixes or routed variants are distinct model identities when the model name itself includes the suffix.

### Hugging Face Router

Hugging Face Router presets use the OpenAI Chat Completions-compatible adapter. Routed backend suffixes are intentionally part of the preset identity when present in the model name.

| Area                  | Support | Notes                                                                            |
| --------------------- | ------- | -------------------------------------------------------------------------------- |
| Text input/output     | yes     | Via OpenAI Chat-compatible request/response shape                                |
| Streaming text        | yes     |                                                                                  |
| Reasoning config      | partial | Provider-wide preset exposes reasoning; model presets narrow support where known |
| Streaming thinking    | no      | OpenAI Chat-compatible response path does not expose separate thinking stream    |
| Output format         | yes     | text and `jsonSchema` where backend supports it                                  |
| Output verbosity      | yes     | Provider-wide preset advertises verbosity                                        |
| Stop sequences        | yes     | Up to provider capability max                                                    |
| Images input          | partial | Provider-wide preset allows images; backend/model support can vary               |
| Files input           | partial | Provider-wide preset allows files; backend/model support can vary                |
| Function/custom tools | partial | Provider-wide preset advertises tools; backend/model support can vary            |
| Web search            | partial | Provider-wide preset advertises web search; backend/model support can vary       |
| Tool policy           | yes     | `auto`, `any`, `tool`, `none` where tools are enabled                            |
| Cache control         | no      | Not enabled by the Hugging Face preset                                           |
| Usage                 | yes     | Subject to what the router/backend returns                                       |

Normalization notes:

- Routed backend suffixes such as `:fireworks-ai`, `:deepinfra`, `:novita`, `:featherless-ai`, and `:cerebras` are treated as distinct preset/model identities.
- Display names for routed Hugging Face presets include the backend name to make that distinction visible to users.

### OpenAI-compatible local and self-hosted runtimes

The preset catalog includes local and self-hosted runtimes for common development and deployment setups.

| Preset provider | Wire adapter                       | Default origin           | Notes                                                                      |
| --------------- | ---------------------------------- | ------------------------ | -------------------------------------------------------------------------- |
| LocalAI         | OpenAI Responses-compatible        | `http://127.0.0.1:8080`  | Local runtime with text/image/file provider preset and per-model overrides |
| LM Studio       | OpenAI Responses-compatible        | `http://127.0.0.1:1234`  | Local OpenAI-compatible server preset                                      |
| llama.cpp       | OpenAI Chat Completions-compatible | `http://127.0.0.1:8080`  | Local OpenAI-compatible server preset                                      |
| Ollama          | Anthropic-compatible               | `http://127.0.0.1:11434` | Local Anthropic-compatible preset with constrained tool policy             |
| SGLang          | OpenAI Responses-compatible        | `http://127.0.0.1:30000` | Self-hosted OpenAI-compatible server preset                                |
| vLLM            | OpenAI Responses-compatible        | `http://127.0.0.1:8000`  | Self-hosted OpenAI-compatible server preset                                |

Normalization notes:

- Local runtimes vary widely by model and server version. The presets provide useful defaults, not a guarantee that every server build supports every declared feature.
- Most local/self-hosted presets rely heavily on model-level overrides for reasoning, modalities, output format support, and stop-sequence behavior.
- Callers should pass a preset capability resolver per completion so model-level overrides are applied.

## Model presets

Package `modelpreset` provides a runtime catalog of common providers and models.

It includes:

- provider names
- model preset IDs
- model names
- provider connection defaults
- model default `spec.ModelParam`
- provider-level capability overrides
- model-level capability overrides

Included preset providers:

- `ProviderAnthropic`
- `ProviderOpenAIResponses`
- `ProviderOpenAIChat`
- `ProviderGoogleGemini`
- `ProviderHuggingFace`
- `ProviderMistral`
- `ProviderOpenRouter`
- `ProviderXAI`
- `ProviderLocalAI`
- `ProviderLMStudio`
- `ProviderLlamaCPP`
- `ProviderOllama`
- `ProviderSGLang`
- `ProviderVLLM`

Preset model IDs, display names, and model names are provider-agnostic where the underlying model identity is the same. Routed or backend-specific models keep distinct IDs and display names when the backend is part of the effective model identity.

Typical use:

```go
providerPreset, err := modelpreset.Provider(modelpreset.ProviderAnthropic)
modelPreset, err := modelpreset.Model(modelpreset.ProviderAnthropic, modelpreset.PresetClaudeSonnet46)
_, err = ps.AddProviderFromPreset(ctx, providerPreset.Name, providerPreset)
```

- The returned presets are cloned. Callers may mutate/customize returned values safely as required.
- Apps that need persistence should treat `modelpreset` as immutable base data and store their own overlay/preference fields separately.
- Apps that persist preset IDs should expect catalog IDs to be stable within a release line, but should still be prepared to migrate IDs when model identities are renamed or provider routing becomes part of the identity.

## Model capabilities and normalization

Capabilities are described by `spec.ModelCapabilities` in [`spec/capability.go`](./spec/capability.go).

Default provider capability profiles live in:

- Anthropic: [`internal/anthropicsdk/capability.go`](./internal/anthropicsdk/capability.go)
- OpenAI Responses: [`internal/openairesponsessdk/capability.go`](./internal/openairesponsessdk/capability.go)
- OpenAI Chat: [`internal/openaichatsdk/capability.go`](./internal/openaichatsdk/capability.go)
- Google Generate Content: [`internal/googlegeneratecontentsdk/capability.go`](./internal/googlegeneratecontentsdk/capability.go)

Preset capability overrides live in:

- provider presets in [`modelpreset`](./modelpreset)
- provider-wide override fields on `ProviderPreset`
- model-level override fields on `ModelPreset`

Hosted routers and local runtimes generally reuse one of the default provider capability profiles and then patch it with preset overrides.
For example, OpenRouter uses the OpenAI Responses-compatible adapter plus model-specific overrides, while Mistral and Hugging Face Router use the OpenAI Chat-compatible adapter plus provider/model overrides.

You can inspect the active provider-wide default via:

- `ProviderSetAPI.GetProviderCapability(ctx, providerName)`

Normalization behavior:

- unsupported contract-like features generally return an error
  - example: unsupported output format
- unsupported safe-to-drop features are removed and reported via `FetchCompletionResponse.Warnings`
  - example: unsupported verbosity or cache-control scope
- some provider-specific history items are sanitized before request conversion
  - Anthropic: only Anthropic-compatible reasoning history is retained
  - OpenAI Responses: only encrypted reasoning history is retained
  - OpenAI Chat: reasoning history is dropped
  - Google: only valid signed Google thought history is retained
- provider and model preset overrides can narrow broad adapter capabilities
  - examples: disabling stop sequences, limiting reasoning levels, changing tool support, changing output format support, or selecting parameter dialects

For per-model capability differences, pass a custom `spec.ModelCapabilityResolver` in `FetchCompletionOptions`.

- For most model-preset based callers, use:
  - `ProviderSetAPI.NewPresetCapabilityResolver`
  - `capabilityoverride.DeriveModelCapabilities`
  - `capabilityoverride.NewCompletionKeyResolver`

### Capability overrides

Provider SDKs expose broad provider-level capabilities. Real models often differ:

- one model may not support files
- one model may only allow a subset of reasoning levels
- one gateway may use a different parameter dialect
- one model may require temperature to be omitted when reasoning is enabled
- one routed model may support JSON Schema while another model from the same router only supports text output

`capabilityoverride.ModelCapabilitiesOverride` is a patch-like form of `spec.ModelCapabilities`.

Layering order is:

1. SDK/provider base capability profile
2. provider preset override
3. model preset override
4. caller/user override, if any

Use `ProviderSetAPI.NewPresetCapabilityResolver` for the common case:

```go
resolver, err := ps.NewPresetCapabilityResolver(
    ctx,
    providerPreset.Name,
    providerPreset,
    modelPreset,
    string(modelPreset.ID),
)
```

Then pass it per completion:

```go
opts := &spec.FetchCompletionOptions{
    CompletionKey:      string(modelPreset.ID),
    CapabilityResolver: resolver,
}
```

`AddProviderFromPreset` only configures the provider connection. Capability overrides are applied per completion through `FetchCompletionOptions`, because
the active model can differ from call to call.
This is especially important for gateway providers such as OpenRouter and Hugging Face Router, and for local/self-hosted runtimes where model support can vary significantly.

## HTTP debugging

The library exposes a pluggable `CompletionDebugger`:

```go
type CompletionDebugger interface {
    HTTPClient(base *http.Client) *http.Client
    StartSpan(ctx context.Context, info *spec.CompletionSpanStart) (context.Context, spec.CompletionSpan)
}
```

Package `debugclient` includes a ready-to-use implementation:

- wraps provider SDK HTTP clients
- captures scrubbed request/response metadata
- redacts secrets and sensitive content
- attaches structured debug data to `FetchCompletionResponse.DebugDetails`

Typical setup:

```go
dbg := debugclient.NewHTTPCompletionDebugger(&debugclient.DebugConfig{
    LogToSlog: false,
})

ps, _ := inference.NewProviderSetAPI(
    inference.WithDebugClientBuilder(func(p spec.ProviderParam) spec.CompletionDebugger {
        return dbg
    }),
)
```

## Notes

- Stateless focus
  - the SDK intentionally focuses on stateless request/response flows
  - provider-native conversation state, uploaded file IDs, stored responses, and similar stateful features are out of scope for the normalized interface

- Opaque provider-specific fields
  - many provider-native details remain available only through debug payloads, not the normalized response structs

- Prompt filtering
  - `ModelParam.MaxPromptLength` uses a heuristic tokenizer via `sdkutil.FilterMessagesByTokenCount`
  - it is approximate, not a provider tokenizer

- Choice/candidate handling
  - OpenAI Chat surfaces the first choice
  - Google Generate Content surfaces the first candidate

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
