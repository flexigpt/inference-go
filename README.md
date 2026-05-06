# LLM Inference for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/flexigpt/inference-go)](https://goreportcard.com/report/github.com/flexigpt/inference-go)
[![lint](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/lint.yml)
[![test](https://github.com/flexigpt/inference-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/flexigpt/inference-go/actions/workflows/test.yml)

A single normalized Go interface for LLM inference across multiple providers, using their official SDKs where available.

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
- [Model capabilities and normalization](#model-capabilities-and-normalization)
- [HTTP debugging](#http-debugging)
- [Notes](#notes)
- [Development](#development)
- [License](#license)

## Features at a glance

- Single normalized interface via `ProviderSetAPI`
- Provider support today:
  - Anthropic Messages API via `github.com/anthropics/anthropic-sdk-go`
  - OpenAI Responses API via `github.com/openai/openai-go/v3`
  - OpenAI Chat Completions API via `github.com/openai/openai-go/v3`
  - Google Generate Content API via `google.golang.org/genai`

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
  - warnings returned in `FetchCompletionResponse.Warnings`
  - per-model capability override support through `FetchCompletionOptions.CapabilityResolver`

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

1. Create a `ProviderSetAPI`
2. Register one or more providers with `AddProvider`
3. Set each provider API key with `SetProviderAPIKey`
4. Call `FetchCompletion`

## Examples

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
  - `spec.ProviderSDKTypeAnthropic`
  - `spec.ProviderSDKTypeOpenAIChatCompletions`
  - `spec.ProviderSDKTypeOpenAIResponses`
  - `spec.ProviderSDKTypeGoogleGenerateContent`

- `Origin`
  - Required
  - Base origin for the provider or gateway/proxy

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

## Model capabilities and normalization

Capabilities are described by `spec.ModelCapabilities` in [`spec/capability.go`](./spec/capability.go).

Default provider capability profiles live in:

- Anthropic: [`internal/anthropicsdk/capability.go`](./internal/anthropicsdk/capability.go)
- OpenAI Responses: [`internal/openairesponsessdk/capability.go`](./internal/openairesponsessdk/capability.go)
- OpenAI Chat: [`internal/openaichatsdk/capability.go`](./internal/openaichatsdk/capability.go)
- Google Generate Content: [`internal/googlegeneratecontentsdk/capability.go`](./internal/googlegeneratecontentsdk/capability.go)

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

For per-model capability differences, pass a custom `spec.ModelCapabilityResolver` in `FetchCompletionOptions`.

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
