# Inference-Go Capability Reference

This document is the implementation reference for normalized capability support across:

- Wire adapters
- Provider presets
- Model presets
- Hosted routers
- Local runtimes
- Request normalization
- Response normalization

The effective capability profile for a request is determined by:

1. The base capability profile of the selected adapter.
2. The provider preset capability override.
3. The model preset capability override.
4. Any application-provided capability override.

A compatible HTTP endpoint does not automatically support the complete feature set of the adapter family. Always use `ProviderSetAPI.NewPresetCapabilityResolver` with catalog providers and models.

## Terminology

- Done
  - Implemented in the current codebase.
  - Includes request normalization, request conversion, or response conversion as applicable.

- Dropped with warning
  - Accepted by the normalized request model but removed during capability normalization.
  - A warning is returned in `FetchCompletionResponse.Warnings`.

- Sanitized
  - Provider-incompatible history or content is removed or transformed before a vendor request is built.

- Model-dependent
  - Support varies by model preset or application-specific capability override.

- Pending
  - Not implemented, partially implemented, or deliberately excluded from the normalized API.

## Capability sources

- `spec/capability.go`
- `spec/param_reasoning.go`
- `internal/anthropicsdk/capability.go`
- `internal/openairesponsessdk/capability.go`
- `internal/openaichatsdk/capability.go`
- `internal/googlegeneratecontentsdk/capability.go`
- `capabilityoverride`
- `modelpreset`

## Normalized request and response support

### Done

- Provider registration and dispatch
  - Anthropic Messages
  - OpenAI Responses
  - OpenAI Chat Completions
  - Google Generate Content
  - Anthropic-compatible provider presets
  - OpenAI Responses-compatible provider presets
  - OpenAI Chat Completions-compatible provider presets

- Normalized request and response contracts
  - `spec.FetchCompletionRequest`
  - `spec.FetchCompletionResponse`
  - `spec.ModelParam`
  - `spec.ReasoningParam`
  - `spec.InputUnion`
  - `spec.OutputUnion`
  - `spec.ToolChoice`
  - `spec.ToolPolicy`
  - `spec.CacheControl`
  - `spec.OutputParam`

- Request normalization
  - Request deep cloning
  - Input-modality validation
  - Reasoning-type validation
  - Reasoning-level validation
  - Summary-style capability handling
  - Reasoning-context capability handling
  - Reasoning-mode capability handling
  - Stop-sequence normalization
  - Output-format validation
  - Output-verbosity capability handling
  - Tool and tool-policy validation
  - Client tool-output normalization
  - Cache-control normalization
  - Capability-driven warnings

- Streaming
  - Text streaming where exposed by the upstream provider
  - Thinking/reasoning streaming for Anthropic, OpenAI Responses, and Google Generate Content where exposed upstream
  - Buffered streaming with configurable flush interval and chunk size
  - Stream-handler panic conversion to errors

- Usage normalization
  - Input-token totals
  - Cached input tokens where exposed
  - Uncached input tokens where exposed
  - Output tokens
  - Reasoning tokens where exposed

- Capability overrides
  - Adapter base capabilities
  - Provider preset overrides
  - Model preset overrides
  - Application-defined overrides
  - Completion-key capability resolvers
  - Deep-cloned resolver capabilities

### Cross-provider limitations

The following remain intentionally non-portable or only partially normalized:

- Audio input and output
- Video input and output
- Image generation and image output
- Provider-native stored conversations
- Previous-response identifiers
- Background jobs
- Prompt resources
- Uploaded-file lifecycle management
- Arbitrary provider-specific parameter passthrough
- Multi-candidate response selection
- Provider-native safety, metadata, service-tier, and user identifiers

## Adapter capability reference

### Anthropic Messages adapter

Source files:

- `internal/anthropicsdk/api_anthropic_messages.go`
- `internal/anthropicsdk/input_processing.go`
- `internal/anthropicsdk/thinking.go`
- `internal/anthropicsdk/cache_control.go`
- `internal/anthropicsdk/capability.go`

| Capability              | Support                                                 |
| ----------------------- | ------------------------------------------------------- |
| Text input/output       | Supported                                               |
| Image input             | Supported                                               |
| File input              | Partial, PDF-focused                                    |
| Text streaming          | Supported                                               |
| Thinking streaming      | Supported when upstream thinking is emitted             |
| Reasoning type          | Hybrid token budget and level-based reasoning           |
| Reasoning history       | Anthropic signed thinking and redacted thinking         |
| Summary style           | Supported through adaptive-thinking display translation |
| Reasoning context       | Unsupported and dropped with warning                    |
| Reasoning mode          | Unsupported and dropped with warning                    |
| JSON Schema output      | Supported                                               |
| Output verbosity        | Supported through Anthropic effort                      |
| Stop sequences          | Supported                                               |
| Function tools          | Supported                                               |
| Custom tools            | Supported                                               |
| Web search              | Supported                                               |
| Tool policies           | `auto`, `any`, `tool`, `none`                           |
| Parallel tools          | Supported                                               |
| Cache control           | Top-level, message, tool, and tool-output scopes        |
| Reasoning cache control | Unsupported                                             |
| URL citations           | Supported where returned by Anthropic                   |
| Output modalities       | Text only in the normalized surface                     |

#### Anthropic reasoning summary mapping

Anthropic adaptive thinking supports only two display values:

- `omitted`
- `summarized`

Normalized mapping:

| `ReasoningSummaryStyle` | Anthropic adaptive display |
| ----------------------- | -------------------------- |
| Unspecified             | `summarized`               |
| `omitted`               | `omitted`                  |
| `auto`                  | `summarized`               |
| `concise`               | `summarized`               |
| `detailed`              | `summarized`               |

This mapping applies to `ReasoningTypeSingleWithLevels`, which uses Anthropic adaptive thinking.

Fixed-budget Anthropic thinking does not expose an equivalent summary-display control.

#### Anthropic reasoning history

- Signed thinking is replayed only when it includes a valid Anthropic signature.
- Redacted thinking is replayed as Anthropic redacted-thinking content.
- Unsigned plaintext reasoning is not replayed as Anthropic thinking.
- Tool-result turn ordering is validated because Anthropic requires matching client tool results immediately after assistant tool-use turns.
- Temperature can be removed with a warning when effective Anthropic reasoning capabilities disallow temperature while thinking is enabled.

#### Anthropic pending items

- Plain-text `text/*` file conversion to Anthropic document input
- Richer citation normalization beyond current URL citation support
- Reasoning-content cache control
- Image, audio, and video output
- Safe allowlisted passthrough for advanced Anthropic parameters

### OpenAI Responses adapter

Source files:

- `internal/openairesponsessdk/api_openai_responses.go`
- `internal/openairesponsessdk/thinking.go`
- `internal/openairesponsessdk/cache_control.go`
- `internal/openairesponsessdk/capability.go`

| Capability                    | Support                                               |
| ----------------------------- | ----------------------------------------------------- |
| Text input/output             | Supported                                             |
| Image input                   | Supported                                             |
| File input                    | Supported                                             |
| Text streaming                | Supported                                             |
| Thinking streaming            | Supported when reasoning deltas are emitted           |
| Reasoning type                | Level-based reasoning                                 |
| Hybrid token-budget reasoning | Unsupported                                           |
| Reasoning history             | Encrypted reasoning content only                      |
| Summary style                 | Supported                                             |
| Reasoning context             | Supported by base adapter capability                  |
| Reasoning mode                | Supported by base adapter capability                  |
| JSON Schema output            | Supported                                             |
| Output verbosity              | Supported                                             |
| Stop sequences                | Unsupported and dropped with warning                  |
| Function tools                | Supported                                             |
| Custom tools                  | Supported through compatible function-style transport |
| Web search                    | Supported where provider/model supports it            |
| Tool policies                 | `auto`, `any`, `tool`, `none`                         |
| Parallel tools                | Supported where provider/model allows it              |
| Cache control                 | Top-level prompt cache key and retention              |
| Per-message cache control     | Unsupported and dropped with warning                  |
| URL citations                 | Supported where returned by provider                  |
| Output modalities             | Text only in the normalized surface                   |

#### OpenAI Responses summary mapping

| `ReasoningSummaryStyle` | OpenAI Responses `reasoning.summary` |
| ----------------------- | ------------------------------------ |
| Unspecified             | `auto`                               |
| `omitted`               | `concise`                            |
| `auto`                  | `auto`                               |
| `concise`               | `concise`                            |
| `detailed`              | `detailed`                           |

#### OpenAI Responses context and mode

The OpenAI Responses adapter supports:

- `ReasoningContextAuto`
- `ReasoningContextCurrentTurn`
- `ReasoningContextAllTurns`
- `ReasoningModeStandard`
- `ReasoningModePro`

The official `modelpreset.ProviderOpenAIResponses` preset enables both context and mode.

Responses-compatible presets for non-OpenAI providers intentionally disable both controls unless their provider/model capability profile explicitly enables them.

When disabled, normalization removes them with:

- `reasoning_context_dropped_unsupported`
- `reasoning_mode_dropped_unsupported`

#### OpenAI Responses reasoning history

- Encrypted reasoning content is preserved for OpenAI Responses continuation.
- If no encrypted reasoning content is present, reasoning history is dropped.
- If encrypted and plaintext reasoning are mixed, only encrypted reasoning content is retained.
- Anthropic signatures, Google thought signatures, and plaintext reasoning must not be reused as OpenAI Responses reasoning input.

#### OpenAI Responses pending items

- Provider-native stateful conversations
- `previous_response_id`
- `store`
- `background`
- Prompt objects
- Full web-search result round-trip representation
- Custom tool definitions without function-style fallback
- Safe allowlisted passthrough for `include`, `truncation`, `service_tier`, metadata, and stream options

### OpenAI Chat Completions adapter

Source files:

- `internal/openaichatsdk/api_openai_chat_completions.go`
- `internal/openaichatsdk/cache_control.go`
- `internal/openaichatsdk/capability.go`

| Capability                | Support                                               |
| ------------------------- | ----------------------------------------------------- |
| Text input/output         | Supported                                             |
| Image input               | Supported                                             |
| File input                | Supported                                             |
| Text streaming            | Supported                                             |
| Thinking streaming        | Unsupported                                           |
| Reasoning type            | Level-based effort where accepted by model/provider   |
| Reasoning history         | Unsupported                                           |
| Summary style             | Unsupported and dropped with warning                  |
| Reasoning context         | Unsupported and dropped with warning                  |
| Reasoning mode            | Unsupported and dropped with warning                  |
| JSON Schema output        | Supported                                             |
| Output verbosity          | Supported                                             |
| Stop sequences            | Supported, up to effective capability limit           |
| Function tools            | Supported                                             |
| Custom tools              | Supported through compatible function-style transport |
| Web search                | Top-level `web_search_options` behavior               |
| Tool policies             | `auto`, `any`, `tool`, `none`                         |
| Parallel tools            | `DisableParallel` maps to `parallel_tool_calls=false` |
| Cache control             | Top-level prompt cache key and retention              |
| Per-message cache control | Unsupported and dropped with warning                  |
| URL citations             | Supported where returned in annotations               |
| Output modalities         | Text only in the normalized surface                   |

#### OpenAI Chat Completions limitations

- Structured reasoning input/output messages are not supported.
- Reasoning messages are not sent as Chat Completions history.
- Thinking streaming is not available.
- Web search is not equivalent to a normal named tool call.
- Forcing web search through `toolPolicy.mode=tool` is not portable.
- Image/file tool output content is not preserved as structured tool output on return turns.

### Google Generate Content adapter

Source files:

- `internal/googlegeneratecontentsdk/api_google_genai.go`
- `internal/googlegeneratecontentsdk/input_processing.go`
- `internal/googlegeneratecontentsdk/thinking.go`
- `internal/googlegeneratecontentsdk/capability.go`

| Capability         | Support                                                 |
| ------------------ | ------------------------------------------------------- |
| Text input/output  | Supported                                               |
| Image input        | Supported                                               |
| File input         | Supported                                               |
| Text streaming     | Supported                                               |
| Thinking streaming | Supported when thought text is emitted                  |
| Reasoning type     | Hybrid token budget and level-based reasoning           |
| Reasoning history  | Valid Google thought signatures only                    |
| Summary style      | Supported through `ThinkingConfig.IncludeThoughts`      |
| Reasoning context  | Unsupported and dropped with warning                    |
| Reasoning mode     | Unsupported and dropped with warning                    |
| JSON Schema output | Supported through raw JSON Schema payload               |
| Output verbosity   | Not mapped by adapter                                   |
| Stop sequences     | Supported, subject to effective model limit             |
| Function tools     | Supported                                               |
| Custom tools       | Supported through function declarations                 |
| Web search         | Google Search grounding                                 |
| Tool policies      | `auto`, `any`, `tool`, `none` for callable tools        |
| Parallel tools     | `DisableParallel` is not normalized/enforced            |
| Cache control      | Unsupported and dropped with warning                    |
| URL citations      | Partial, grounding maps to synthetic web-search outputs |
| Output modalities  | Text only in the normalized surface                     |

#### Google reasoning summary mapping

Google maps summary visibility to `ThinkingConfig.IncludeThoughts`.

| `ReasoningSummaryStyle` | `IncludeThoughts` |
| ----------------------- | ----------------- |
| Unspecified             | `true`            |
| `omitted`               | `false`           |
| `auto`                  | `true`            |
| `concise`               | `true`            |
| `detailed`              | `true`            |

Reasoning still executes when `IncludeThoughts=false`; Google simply does not return thought text.

#### Google reasoning history

- Only valid Google thought signatures are replayed.
- Signature-only thought parts are retained.
- Anthropic signed/redacted thinking is dropped.
- OpenAI encrypted reasoning content is dropped.
- Plaintext unsigned reasoning is dropped.

#### Google pending items

- `ToolPolicy.DisableParallel` capability enforcement
- Mapping `JSONSchemaParam.Name`, `Description`, and `Strict`
- Rich image/file function-tool output history
- Full grounding option normalization
- Attaching grounding citations directly to `ContentItemText.Citations`
- Multi-candidate response handling
- Image, audio, and video output

## Preset provider catalog reference

The catalog is implemented in `modelpreset`.

Provider and model presets are static reviewed defaults. They are not runtime discovery or remote capability probing.

### Complete provider matrix

| Provider preset     | Adapter family                     | Catalog model families                         | Notable capability posture                                                                     |
| ------------------- | ---------------------------------- | ---------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| Anthropic           | Anthropic Messages                 | Claude Fable, Opus, Sonnet, Haiku              | Anthropic-native reasoning, cache controls, tools, and summary translation                     |
| DeepSeek            | OpenAI Responses-compatible        | DeepSeek V4                                    | Text-focused profile with constrained tools and level-based reasoning                          |
| Google Gemini       | Google Generate Content            | Gemini 2.5 and Gemini 3.x                      | Native thought signatures, reasoning budgets/levels, grounding, files, and summary translation |
| Hugging Face Router | OpenAI Chat Completions-compatible | Routed open models                             | Backend-qualified model IDs and model-specific overrides                                       |
| LocalAI             | OpenAI Responses-compatible        | Local model presets                            | Local deployment behavior is model/server dependent                                            |
| LM Studio           | OpenAI Responses-compatible        | Local model presets                            | Local deployment behavior is model/server dependent                                            |
| llama.cpp           | OpenAI Chat Completions-compatible | Llama and Qwen local presets                   | Local deployment behavior is model/server dependent                                            |
| Meta                | OpenAI Responses-compatible        | Muse Spark 1.1, 1.2, 1.3, Contributor variants | Summary style enabled; context/mode disabled                                                   |
| MiniMax             | OpenAI Responses-compatible        | MiniMax M2 and M3                              | Model-specific reasoning with text/tool restrictions                                           |
| Mistral             | OpenAI Chat Completions-compatible | Mistral and Devstral                           | Mistral-specific parameter dialect                                                             |
| Moonshot            | Anthropic Messages-compatible      | Kimi models                                    | Model-specific Anthropic-compatible reasoning/tool behavior                                    |
| Ollama              | Anthropic Messages-compatible      | Ollama-tagged local models                     | Local Anthropic-compatible behavior with constrained tool policy                               |
| OpenAI Chat         | OpenAI Chat Completions            | GPT-4.1 and GPT-4o                             | Chat Completions transport and model-specific no-reasoning overrides                           |
| OpenAI Responses    | OpenAI Responses                   | GPT-5 and GPT-6 families                       | Summary, context, and mode enabled                                                             |
| OpenRouter          | OpenAI Responses-compatible        | Routed hosted models                           | Model-specific modalities, output formats, tools, and reasoning levels                         |
| QwenCloud           | OpenAI Responses-compatible        | Qwen, Qwen Coder, Character models             | DashScope compatible-mode behavior with provider/model restrictions                            |
| SGLang              | OpenAI Responses-compatible        | Local model presets                            | Self-hosted behavior with model-specific overrides                                             |
| vLLM                | OpenAI Responses-compatible        | Local model presets                            | Self-hosted behavior with model-specific overrides                                             |
| xAI                 | OpenAI Responses-compatible        | Grok models                                    | Model-specific reasoning, encrypted reasoning, cache, tool, and output support                 |
| Xiaomi              | OpenAI Responses-compatible        | MiMo models                                    | Text-focused provider defaults with model-specific image support                               |
| Z.AI                | OpenAI Chat Completions-compatible | GLM models                                     | Pay-as-you-go Chat Completions endpoint                                                        |
| Z.AI Coding Plan    | OpenAI Responses-compatible        | GLM Coding Plan models                         | Separate endpoint/credential route with model-specific overrides                               |

### Provider reasoning-control matrix

| Provider            | Reasoning configuration                        | Summary style                              | Context/mode |
| ------------------- | ---------------------------------------------- | ------------------------------------------ | ------------ |
| Anthropic           | Hybrid-budget and level-based, model-dependent | Anthropic display translation              | Unsupported  |
| DeepSeek            | Level-based                                    | Unsupported                                | Unsupported  |
| Google Gemini       | Hybrid-budget and level-based                  | `IncludeThoughts` translation              | Unsupported  |
| Hugging Face Router | Model-dependent                                | Model-dependent                            | Unsupported  |
| LocalAI             | Model-dependent                                | Model-dependent                            | Unsupported  |
| LM Studio           | Model-dependent                                | Model-dependent                            | Unsupported  |
| llama.cpp           | Model-dependent                                | Model-dependent                            | Unsupported  |
| Meta                | Level-based                                    | OpenAI Responses summary mapping           | Unsupported  |
| MiniMax             | Model-dependent level-based                    | Unsupported                                | Unsupported  |
| Mistral             | Model-dependent level-based                    | Unsupported                                | Unsupported  |
| Moonshot            | Model-dependent hybrid/level-based             | Unsupported                                | Unsupported  |
| Ollama              | Model-dependent                                | Unsupported                                | Unsupported  |
| OpenAI Chat         | Model-dependent level-based                    | Unsupported                                | Unsupported  |
| OpenAI Responses    | Level-based                                    | Native, with `omitted` mapped to `concise` | Supported    |
| OpenRouter          | Model-dependent level-based                    | Model-dependent                            | Unsupported  |
| QwenCloud           | Level-based                                    | Unsupported                                | Unsupported  |
| SGLang              | Model-dependent                                | Model-dependent                            | Unsupported  |
| vLLM                | Model-dependent                                | Model-dependent                            | Unsupported  |
| xAI                 | Model-dependent level-based                    | Model-dependent                            | Unsupported  |
| Xiaomi              | Level-based                                    | Unsupported                                | Unsupported  |
| Z.AI                | Provider/model dependent                       | Unsupported                                | Unsupported  |
| Z.AI Coding Plan    | Model-dependent                                | Model-dependent                            | Unsupported  |

### Meta Muse Spark capability profile

The Meta preset uses the OpenAI Responses-compatible adapter.

Muse Spark 1.1, 1.2, 1.3, and Contributor variants are included in the catalog.

The Meta provider capability profile declares:

- Text, image, and file input
- Text output
- Level-based reasoning
  - `minimal`
  - `low`
  - `medium`
  - `high`
  - `xhigh`
- Summary-style reasoning
- Encrypted reasoning input support
- JSON Schema output
- Function and web-search tool types
- `auto` tool policy
- Top-level prompt cache key and retention support
- No normalized output verbosity
- No stop sequences
- No reasoning context
- No reasoning mode

Meta examples currently cover:

- Muse Spark 1.3 Contributor basic completion
- Muse Spark 1.3 Contributor reasoning summary style
- Muse Spark 1.3 Contributor cache controls
- Muse Spark 1.3 Contributor image/file input
- Muse Spark 1.3 Contributor JSON Schema output
- Muse Spark 1.3 Contributor function tools with `auto` policy
- Muse Spark 1.3 Contributor streamed text/reasoning output

Meta does not yet have a dedicated live web-search integration example.

### Catalog validation and test coverage

The catalog test suite verifies:

- Provider catalog identity and connection fields
- Provider names returned by `ProviderNames`
- Provider/model lookup APIs
- Model preset IDs
- Model preset names and `ModelParam.Name` consistency
- Required sampling defaults
- Reasoning defaults
- Cache defaults
- Output defaults
- Stop-sequence defaults
- Capability override structural validity
- Preset cloning behavior
- ProviderSet integration
- Effective capability derivation
- Request normalization for model defaults

Important tests include:

- `TestDefaultCatalogValidatesEveryProvider`
- `TestCatalogContainsAllRegisteredProviders`
- `TestCatalogModelMembershipIsExhaustive`
- `TestCatalogPublicLookupMethodsCoverEveryProviderAndModel`
- `TestCatalogPresetsIntegrateWithProviderSet`

`TestCatalogModelMembershipIsExhaustive` is intentionally a static manifest. Adding, removing, or renaming a provider/model preset requires an explicit test update.

A passing catalog test confirms internal metadata consistency. It does not replace live verification against provider endpoints.

## Capability warnings

The normalizer can return warnings for safe request reductions.

Important reasoning-related warning codes include:

| Warning code                            | Meaning                                                                   |
| --------------------------------------- | ------------------------------------------------------------------------- |
| `reasoning_dropped_unsupported`         | Reasoning type/configuration is unsupported                               |
| `reasoning_dropped_invalid_level`       | Requested reasoning level is unavailable for effective model capabilities |
| `reasoning_summaryStyle_dropped`        | Summary style is unsupported and was removed                              |
| `reasoning_context_dropped_unsupported` | Reasoning context is unsupported and was removed                          |
| `reasoning_mode_dropped_unsupported`    | Reasoning mode is unsupported and was removed                             |
| `temperature_dropped_reasoning_enabled` | Temperature conflicts with reasoning for the effective model              |
| `stopSequences_dropped_unsupported`     | Stop sequences are unsupported                                            |
| `stopSequences_dropped_reasoning`       | Stop sequences conflict with reasoning                                    |
| `verbosity_dropped_unsupported`         | Output verbosity is unsupported                                           |
| `cacheControl_dropped_unsupported`      | Cache control scope is unsupported                                        |
| `cacheControl_ttl_dropped_unsupported`  | Cache retention/TTL is unsupported                                        |
| `cacheControl_key_dropped_unsupported`  | Cache key is unsupported                                                  |
| `toolChoice_dropped_unsupported`        | Tool type is unsupported                                                  |
| `toolOutput_collapsed_to_string`        | Rich tool output was converted to a string-only transport                 |

## Known capability caveats

- `ToolCapabilities.SupportsParallelToolCalls` exists, but normalization does not yet enforce unsupported `ToolPolicy.DisableParallel` requests consistently.
- Google provider presets may expose output verbosity metadata, but the Google adapter currently does not map normalized output verbosity into `GenerateContentConfig`.
- A model capability override can differ from the underlying adapter default.
- A provider may expose a compatible endpoint but reject newer upstream API fields.
- Routers and local servers can change capability behavior independently of this repository.
- Provider/model capability overrides should be updated when integration verification discovers a behavior change.

## Pending cross-provider work

- Audio input/output normalization
- Video input/output normalization
- Image output normalization
- Richer citation abstraction
- Consistent grounding citations on text content
- Multi-candidate response handling
- Provider-native response metadata promotion
- Safe allowlisted `AdditionalParametersRawJSON` passthrough
- Explicit stateful conversation policy
- Previous-response IDs
- Background jobs
- Stored responses
- Prompt-resource support
- Uploaded-file lifecycle support
- Uniform enforcement for `ToolPolicy.DisableParallel`
- More complete provider-specific web-search behavior normalization
