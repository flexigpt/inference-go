---
name: inference-go-usage
description: Use inference-go end to end for normalized model inference - select a preset, register/authenticate a provider, resolve capabilities, send requests, stream, continue conversations, use tools, and diagnose results.
---

# Inference-Go end-to-end cheat sheet

Requires Go 1.26+. Start with a catalog preset unless the endpoint/model is application-defined.

## Standard preset flow

1. Create `inference.ProviderSetAPI`; optionally supply a logger and debug-client builder.
2. Look up a provider with `modelpreset.Provider` and a model with `modelpreset.Model`. Returned values are clones and may be customized for a request.
3. Register it with `AddProviderFromPreset` and set credentials with `SetProviderAPIKey`.
4. Create a model-specific resolver with `NewPresetCapabilityResolver` using a stable completion key, normally `string(model.ID)`.
5. Put that same key and resolver in `spec.FetchCompletionOptions`.
6. Copy the model's `ModelParam`, make request-specific changes, create a nonempty user-led input history, and call `FetchCompletion` with the registered provider name.
7. Handle both the returned response and error. Errors can accompany a partial response; inspect `Response.Error`, `Warnings`, `Outputs`, and `Usage`.

See the complete setup sequence in `internal/integration/example_preset_catalog_test.go` and reusable helpers in `internal/integration/helpers_test.go`.

## Request essentials

| Need                   | Use                                 | Notes                                                                        |
| ---------------------- | ----------------------------------- | ---------------------------------------------------------------------------- |
| Model/defaults         | `FetchCompletionRequest.ModelParam` | `Name` and at least one input are required. `SystemPrompt` is top-level.     |
| User/assistant history | `spec.InputUnion`                   | Use matching kind/pointer fields; history must begin with a real user input. |
| Text/image/file        | `InputOutputContent` items          | Effective model capabilities decide whether the modality is valid.           |
| Reasoning              | `ModelParam.Reasoning`              | Types, levels, summary/context/mode vary by model/adapter.                   |
| JSON Schema            | `ModelParam.OutputParam`            | Unsupported formats are errors, not silent fallbacks.                        |
| Tools                  | `ToolChoices` and `ToolPolicy`      | Preserve returned call IDs for the next turn.                                |
| Cache/stops/verbosity  | `ModelParam` fields                 | Unsupported safe settings become warnings or are narrowed.                   |

Use the preset resolver on every request. Effective capability order is adapter base, provider preset, model preset, then caller override. `docs/current_capability_ref.md` is the authoritative feature matrix.

`MaxPromptLength` causes heuristic history trimming in `ProviderSetAPI`; retain enough complete turns and tool call/output pairs in application history.

## Read responses and continue a conversation

`FetchCompletionResponse.Outputs` is a normalized union. Handle assistant text, reasoning, function/custom calls, and web-search events by kind; `Usage` normalizes input, cache, output, and reasoning token counts where upstream exposes them.

For a follow-up, retain the original user inputs, append normalized assistant outputs as their corresponding `InputUnion` forms, then append the next user input. `outputUnionsToInputs` in `internal/integration/helpers_test.go` is the reference mapping.

Reasoning continuation is provider/route-specific. Keep unmodified native reasoning only for a compatible continuation; the library fingerprints routes and filters mismatched reasoning. When switching provider families, retain ordinary text conversation history rather than provider-native reasoning.

## Streaming

Set `ModelParam.Stream` and provide `FetchCompletionOptions.StreamHandler`. The handler receives `text` and, where supported, `thinking` events. Returning an error stops the stream and surfaces it from `FetchCompletion`; configure buffering with `StreamConfig` if needed.

If `Stream` is true but no handler is supplied, adapters use their non-streaming path. See the streaming examples for Anthropic, OpenAI Responses, OpenAI Chat, and Google under `internal/integration/`.

## Tools

1. Define `ToolChoice` with a stable ID, name, type, description, and JSON-schema arguments for callable tools.
2. Choose `ToolPolicy`: `auto`, `any`, one named `tool`, or `none`.
3. On a client function/custom call, execute it, then append the call and a matching `ToolOutput` to the next request. Preserve type, choice ID, call ID, name, and order.
4. Keep Anthropic tool-result ordering strict: a matching user tool result immediately follows the assistant tool-use turn.

Use `example_anthropic_tools_streaming_test.go` and `example_google_genai_tools_roundtrip_test.go` for client tool loops. Web search is provider-native: Anthropic/Responses can return server-side events; Google uses grounding; Chat Completions uses top-level web-search options. Do not treat it as a universally forceable client tool.

## Debug and troubleshoot

Use `WithDebugClientBuilder` with `debugclient.NewHTTPCompletionDebugger`; inspect `Response.DebugDetails`. The default debugger strips model text/large content from captured details; see `debugclient/debugger.go` for controls.

When behavior differs from expectations, check in order:

1. registered provider name, endpoint path, headers, and API key;
2. exact model ID and completion-key/resolver match;
3. `Warnings` for dropped/narrowed settings;
4. effective capability metadata and the adapter-specific README;
5. normalized output/error before assuming a vendor SDK failure.

## Choose the right next action

- Add a reviewed library catalog entry: use `add-model-preset`.
- Use an application-owned compatible endpoint/model: use `custom-providers-and-models`.
- Support a new wire API: use `add-sdk-support`.
