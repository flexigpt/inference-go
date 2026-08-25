# OpenAI Responses Adapter

Use this adapter for OpenAI Responses and providers that expose a compatible Responses-style API.

Preset-backed providers using this adapter include:

- OpenAI Responses
- DeepSeek
- Meta
- MiniMax
- OpenRouter
- QwenCloud
- Xiaomi
- Z.AI Coding Plan
- xAI
- LocalAI
- LM Studio
- SGLang
- vLLM

## Best fit

Choose this adapter when your application benefits from:

- OpenAI Responses-style reasoning
- Encrypted reasoning continuation
- Rich image and file input
- JSON Schema output
- Built-in web search where supported by the provider/model
- Function and custom tools
- Thinking streaming where exposed by the upstream response
- Top-level prompt-cache controls

## High-level support

| Capability               | Support                                               |
| ------------------------ | ----------------------------------------------------- |
| Text input and output    | Supported                                             |
| Image input              | Supported                                             |
| File input               | Supported                                             |
| Text streaming           | Supported                                             |
| Thinking streaming       | Supported where upstream reasoning deltas are emitted |
| Reasoning configuration  | Supported for compatible models                       |
| Reasoning history        | Encrypted reasoning only                              |
| JSON Schema output       | Supported                                             |
| Output verbosity         | Supported where the effective model allows it         |
| Stop sequences           | Not supported by the base adapter                     |
| Function tools           | Supported                                             |
| Custom tools             | Supported through compatible function-style transport |
| Web search               | Supported where the provider/model exposes it         |
| Tool policy              | Supported                                             |
| Cache control            | Top-level prompt-cache behavior only                  |
| URL citations            | Supported where returned by the provider              |
| Audio/video/image output | Not normalized                                        |

## Important guidance

### Reasoning continuation

OpenAI Responses continuation is safest when the prior response contains encrypted reasoning content.

The adapter keeps encrypted reasoning and removes incompatible visible reasoning data before replaying history.

Do not expect Anthropic signed thinking or Google thought signatures to continue through this adapter.

### Stop sequences

The base Responses API does not support normalized stop sequences.

If your application depends on stop sequences, inspect warnings and consider a Chat Completions-compatible provider/model instead.

### Cache controls

This adapter supports top-level prompt-cache behavior.

Per-message, tool, and reasoning cache controls are not portable through the Responses adapter.

### Tools

Function and custom tools are supported.

Custom tools may use function-style transport underneath because that is the compatible request shape available through the adapter.

Web search is provider-native. Its exact returned detail and policy behavior can vary by provider/model.

### Stateful provider features

The normalized API is intentionally stateless.

Do not depend on provider-native features such as stored responses, previous-response IDs, background jobs, or prompt resources through this adapter.

## Working examples

- [OpenAI Responses basic completion](../integration/example_openai_responses_basic_test.go)
- [OpenAI Responses tools, attachments, and streaming](../integration/example_openai_responses_tools_attachments_test.go)
- [Capability overrides](../integration/example_capability_override_test.go)
