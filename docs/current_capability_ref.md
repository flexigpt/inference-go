# API capability and normalization done and pending reference

This document is the current implementation reference for normalized API capability support across providers.

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

- [Cross-provider reference](#cross-provider-reference)
  - [Done reference](#done-reference)
  - [Pending reference](#pending-reference)
- [Anthropic Messages API reference](#anthropic-messages-api-reference)
  - [Done reference](#done-reference-1)
  - [Pending reference](#pending-reference-1)
- [OpenAI Responses API reference](#openai-responses-api-reference)
  - [Done reference](#done-reference-2)
  - [Pending reference](#pending-reference-2)
- [OpenAI Chat Completions API reference](#openai-chat-completions-api-reference)
  - [Done reference](#done-reference-3)
  - [Pending reference](#pending-reference-3)
- [Google Generate Content API reference](#google-generate-content-api-reference)
  - [Done reference](#done-reference-4)
  - [Pending reference](#pending-reference-4)
- [Cross-provider pending backlog reference](#cross-provider-pending-backlog-reference)
  - [Pending](#pending)

---

## Cross-provider reference

### Done reference

- Provider registration and dispatch
  - Anthropic Messages
  - OpenAI Responses
  - OpenAI Chat Completions
  - Google Generate Content

- Normalized request/response model
  - `spec.FetchCompletionRequest`
  - `spec.FetchCompletionResponse`
  - `spec.InputUnion`
  - `spec.OutputUnion`
  - `spec.ModelParam`
  - `spec.ToolChoice`
  - `spec.ToolPolicy`

- Capability-driven normalization
  - request cloning
  - modality validation
  - reasoning validation
  - stop-sequence normalization
  - output-format validation
  - tool capability validation
  - cache-control normalization
  - warnings returned in `FetchCompletionResponse.Warnings`

- Streaming
  - normalized text streaming
  - normalized thinking streaming where the provider exposes it

- Usage normalization
  - input tokens
  - output tokens
  - cached token accounting where exposed
  - reasoning token accounting where exposed

- Debugging
  - `CompletionDebugger`
  - `debugclient.HTTPCompletionDebugger`

### Pending reference

- Audio input/output normalization
- Video input/output normalization
- Image output normalization
- Richer cross-provider citation normalization beyond URL/basic grounding forms
- Safe allowlisted passthrough for `ModelParam.AdditionalParametersRawJSON`
- Explicit capability signal for `ToolPolicy.DisableParallel`
  - today this is not represented separately in the capability model
  - some providers can disable parallel tool calls, some cannot, and docs must currently describe that per provider

---

## Anthropic Messages API reference

Capability source:

- `internal/anthropicsdk/capability.go`

Adapter references:

- `internal/anthropicsdk/api_anthropic_messages.go`
- `internal/anthropicsdk/input_processing.go`
- `internal/anthropicsdk/thinking.go`
- `internal/anthropicsdk/cache_control.go`

### Done reference

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

### Pending reference

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

---

## OpenAI Responses API reference

Capability source:

- `internal/openairesponsessdk/capability.go`

Adapter references:

- `internal/openairesponsessdk/api_openai_responses.go`
- `internal/openairesponsessdk/thinking.go`
- `internal/openairesponsessdk/cache_control.go`

### Done reference

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

### Pending reference

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

---

## OpenAI Chat Completions API reference

Capability source:

- `internal/openaichatsdk/capability.go`

Adapter references:

- `internal/openaichatsdk/api_openai_chat_completions.go`
- `internal/openaichatsdk/cache_control.go`

### Done reference

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

### Pending reference

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

---

## Google Generate Content API reference

Capability source:

- `internal/googlegeneratecontentsdk/capability.go`

Adapter references:

- `internal/googlegeneratecontentsdk/api_google_genai.go`
- `internal/googlegeneratecontentsdk/input_processing.go`
- `internal/googlegeneratecontentsdk/thinking.go`

### Done reference

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

### Pending reference

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

---

## Cross-provider pending backlog reference

These are the main remaining normalized-surface items that affect more than one provider.

### Pending

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
