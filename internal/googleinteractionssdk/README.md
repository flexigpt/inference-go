# Google Interactions Adapter

Use this adapter for stateless Google Gemini inference through the Google
Interactions API.

Preset-backed provider:

- Google Gemini Interactions

Gemini 4 Neon, Carbon, Hydrogen, and Argon use this adapter.

## Request boundary

Every request uses model-create interactions only.

The adapter explicitly sends:

- `store=false`
- `background=false`

The adapter does not support:

- Managed agents
- Stored interactions
- `previous_interaction_id`
- Continuation-token resumption
- Background execution
- Interaction retrieval, cancellation, or deletion
- Safety-setting configuration
- Computer use, MCP, file search, Maps, or URL-context tools

The SDK automatic-continuation option is disabled. A
`requires_continuation` response returns accumulated partial output with an
error because continuation requires stateful provider behavior.

## Supported normalized capabilities

| Capability | Support |
| --- | --- |
| Text input and output | Supported |
| Image input | Supported for documented image MIME types |
| Document input | Supported for documented document MIME types |
| Text streaming | Supported |
| Thinking-summary streaming | Supported |
| Reasoning | `minimal`, `low`, `medium`, `high`, and `xhigh` |
| Reasoning continuation | Same-route signed thought steps only |
| JSON Schema output | Supported through text response format |
| Stop sequences | Supported |
| Function tools | Supported |
| Custom tools | Supported through function declarations |
| Function-tool result content | String or text/image/document content list |
| Google Search grounding | Supported |
| Usage | Input, cached input, output, and thought tokens |
| HTTP debugging | Supported |

## Reasoning

Interactions thought signatures are opaque provider data.

The adapter:

- Replays only reasoning with a nonempty thought signature.
- Preserves summaries when present.
- Drops plaintext thinking, redacted thinking, and encrypted reasoning data.
- Relies on `ProviderSetAPI` route fingerprints to prevent replay across
  incompatible provider routes.

`ReasoningSummaryStyleOmitted` maps to `thinking_summaries=none`.
`auto`, `concise`, and `detailed` map to `thinking_summaries=auto`.

Hybrid token-budget reasoning, reasoning context, and reasoning mode are not
supported.

## Tools and grounding

Function and custom tools use Interactions function declarations.

Function tool results preserve the normalized call ID and function name.

Single text results use the SDK string union member. Rich text, image, and
document results use the SDK `FunctionResultSubcontent` array union member.
The shared normalizer still collapses to string only when a caller override
restricts transport to string-only.

Google Search is server-side grounding.

- `toolPolicy=auto` can include Google Search.
- Forced `any` and `tool` policies require function or custom tools.
- `toolPolicy=none` omits every tool.
- `DisableParallel` is not enforceable and produces a warning.

Interactions Google Search result steps expose search suggestions. The supplied
SDK model does not expose source URL, title, or character-span data. The adapter
maps suggestions to normalized web-search output content and does not invent
`URLCitation` values.

## Streaming

Streaming requires both:

- `ModelParam.Stream=true`
- A non-nil `FetchCompletionOptions.StreamHandler`

The adapter calls `Interactions.Create` with `Stream=true`. The official SDK
returns `InteractionSSEStreamEvent` on the create response. The adapter
consumes that iterator through `Next` and `Value`.

The adapter emits normalized text and thinking events from official SDK step
delta accessors. It buffers events with the shared stream configuration,
converts handler panics to errors, cancels the active stream after handler
failure, and returns accumulated partial output.

## Unsupported optional controls

- Temperature is dropped with a warning.
- Output verbosity is dropped by shared capability normalization.
- Cache control is dropped by shared capability normalization.
- `AdditionalParametersRawJSON` is not forwarded.
- JSON Schema name, description, and strict metadata are not forwarded.
