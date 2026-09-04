# Google Generate Content Adapter

Use this adapter for Google Gemini through the Google Generate Content API.

Preset-backed provider:

- Google Gemini

## Best fit

Choose this adapter when you need:

- Gemini-native thinking controls
- Google thought-signature continuation
- Google Search grounding
- Image and file input
- Thinking streaming
- Function-tool workflows with Gemini models

## High-level support

| Capability               | Support                                                    |
| ------------------------ | ---------------------------------------------------------- |
| Text input and output    | Supported                                                  |
| Image input              | Supported                                                  |
| File input               | Supported                                                  |
| Text streaming           | Supported                                                  |
| Thinking streaming       | Supported where upstream thought text is emitted           |
| Reasoning configuration  | Supported                                                  |
| Reasoning history        | Supported for valid Google thought signatures              |
| JSON Schema output       | Supported                                                  |
| Output verbosity         | Not supported                                              |
| Stop sequences           | Supported, subject to effective model limits               |
| Function tools           | Supported                                                  |
| Custom tools             | Supported through function declarations                    |
| Web search               | Supported through Google Search grounding                  |
| Tool policy              | Supported for callable tools                               |
| Cache control            | Not supported through the normalized cache-control surface |
| URL citations            | Partial, grounding is represented as web-search outputs    |
| Audio/video/image output | Not normalized                                             |

## Important guidance

### Preserve Google thought signatures

When continuing a Gemini conversation with reasoning, preserve the normalized assistant output and reasoning items returned by the previous response.

Google thought signatures are provider-issued continuation data. Do not alter or fabricate them.

Only valid Google thought signatures are replayed through this adapter.

### Reasoning summary style

Google controls visibility of reasoning through `ThinkingConfig.IncludeThoughts`.

- `ReasoningSummaryStyleOmitted` maps to `IncludeThoughts=false`.
- An unspecified summary style and `auto`, `concise`, or `detailed` map to `IncludeThoughts=true`.

Reasoning still executes when `IncludeThoughts=false`; the provider simply does not return thought text.

`ReasoningParam.Context` and `ReasoningParam.Mode` are not supported by Google Generate Content and are removed during normalization with warnings.

### Web search is grounding

Google Search is server-side grounding.

It is not a client function-tool round trip.

Applications can enable grounding, but should not expect web search to be forced exactly like a named function tool.

### Tool parallelism

The normalized `DisableParallel` request is not currently enforced by this adapter.

If your application requires single-tool-call behavior, validate it at the application level.

### Structured output

Text and JSON Schema output are supported.

The schema payload is forwarded to Google, but some normalized schema metadata is not currently portable.

### Cache control

Google context caching is not represented through the normalized `CacheControl` model.

Normalized cache-control settings are removed with warnings.

### Candidate behavior

The adapter normalizes the first returned candidate.

Applications that need multi-candidate selection should use provider-native behavior outside the current normalized surface.

## Working examples

- [Google basic completion](../integration/example_google_genai_basic_test.go)
- [Google function-tool round trip](../integration/example_google_genai_tools_roundtrip_test.go)
- [Google web search, thinking, and streaming](../integration/example_google_genai_websearch_streaming_test.go)
- [Google multi-turn tool loop](../integration/example_google_generate_content_loop_test.go)
