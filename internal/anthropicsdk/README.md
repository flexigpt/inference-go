# Anthropic Messages Adapter

Use this adapter when you need Anthropic Messages behavior or an Anthropic-compatible provider preset.

Preset-backed providers using this adapter:

- Anthropic
- Moonshot
- Ollama

## Best fit

Choose this adapter when your application benefits from:

- Claude-style thinking and signed reasoning continuation
- Anthropic-native tool calling
- Strict tool-use and tool-result conversation handling
- Anthropic ephemeral cache controls
- PDF input
- Server-side web search
- Thinking streaming

## High-level support

| Capability               | Support                                                        |
| ------------------------ | -------------------------------------------------------------- |
| Text input and output    | Supported                                                      |
| Image input              | Supported                                                      |
| File input               | Partial, PDF-focused                                           |
| Text streaming           | Supported                                                      |
| Thinking streaming       | Supported where upstream thinking is emitted                   |
| Reasoning configuration  | Supported                                                      |
| Reasoning history        | Supported for Anthropic-compatible signed or redacted thinking |
| JSON Schema output       | Supported                                                      |
| Output verbosity         | Supported                                                      |
| Stop sequences           | Supported                                                      |
| Function tools           | Supported                                                      |
| Custom tools             | Supported                                                      |
| Web search               | Supported                                                      |
| Tool policy              | Supported                                                      |
| Cache control            | Supported at several Anthropic-compatible scopes               |
| URL citations            | Supported where returned by Anthropic-compatible content       |
| Audio/video/image output | Not normalized                                                 |

## Important guidance

### Thinking and temperature

Some Anthropic models do not accept temperature while thinking is enabled.

Use the model preset resolver and inspect warnings. The normalizer removes temperature when the effective model capability profile requires that behavior.

### Reasoning summary style

The Anthropic adaptive-thinking API exposes only two display options:

- `omitted`
- `summarized`

`ReasoningSummaryStyleOmitted` maps to `display="omitted"`.

An unspecified summary style and the normalized `auto`, `concise`, and `detailed` styles map to `display="summarized"`.

This mapping applies to adaptive thinking, which is selected by `ReasoningTypeSingleWithLevels`.
Fixed-budget thinking does not expose an equivalent summary-display control.

`ReasoningParam.Context` and `ReasoningParam.Mode` are not supported by this adapter and are removed during normalization with warnings.

### Tool-result ordering

Anthropic tool history is stricter than many OpenAI-compatible APIs.

When a model emits a client tool call:

1. Preserve the assistant tool call.
2. Return matching tool output in the next user turn.
3. Do not leave unresolved tool calls at the end of request history.

The Anthropic tool round-trip example demonstrates the required conversation shape.

### Reasoning history

Reuse Anthropic signed or redacted reasoning only with compatible Anthropic-family continuations.

Do not expect ordinary reasoning text from another provider to work as Anthropic thinking history.

### Files

PDF input is the supported file path.

Plain-text and other file formats are not currently portable through this adapter.

### Cache control

Anthropic cache behavior is designed around ephemeral cache controls.

Use the effective model capability resolver because provider/model presets can narrow cache support.

## Working examples

- [Basic Anthropic completion](../integration/example_anthropic_basic_test.go)
- [Tools, reasoning, and streaming](../integration/example_anthropic_tools_streaming_test.go)
- [Anthropic function-tool round trip](../integration/example_anthropic_tools_streaming_test.go)
