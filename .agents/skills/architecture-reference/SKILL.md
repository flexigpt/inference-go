---
name: architecture-reference
description: Navigate inference-go architecture when locating a change, diagnosing a behavior, or deciding whether work belongs in contracts, normalization, an adapter, presets, overrides, debugging, or tests.
---

# Inference-Go architecture reference

Use this before adding cross-cutting behavior or when a request reaches an unexpected adapter, is normalized away, fails at an upstream API, or returns an unexpected normalized result. This is a navigation guide: follow the linked source rather than duplicating implementation here.

## System shape

`inference-go` is a provider-neutral contract over concrete wire adapters:

```text
application
  -> package inference (ProviderSetAPI: registration, dispatch, history boundary)
  -> internal/<adapter>sdk (normalize, translate, call vendor SDK, translate back)
  -> vendor API
```

The normalized surface is `spec`; provider SDK types must remain in `internal/*sdk`. The built-in catalog supplies reviewed endpoint/model defaults. Capabilities determine the safe subset for a particular model and route.

## Package map

| Area                     | Responsibility                                                                                                                  | Start here                                              | Do not put here                                                                     |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `spec/`                  | Public normalized request, response, provider, capability, tool, cache, and stream contracts                                    | `spec/api.go`, `spec/io_union.go`, `spec/capability.go` | Vendor SDK types or endpoint-specific fields without a deliberate portable contract |
| root `package inference` | Provider registration, lifecycle dispatch, route-bound reasoning filtering, prompt-history trimming                             | `provider_set.go`                                       | Vendor request/response conversion or catalog data                                  |
| `internal/sdkutil/`      | Shared adapter protocol, request cloning/normalization, capability enforcement, stream safety, tool naming, continuation safety | `capability_normalize.go`, `completion.go`              | Vendor-specific wire mapping                                                        |
| `internal/<adapter>sdk/` | One adapter family’s lifecycle, base capabilities, normalized-to-vendor conversion, streaming, normalized response conversion   | `api_*.go`, `capability.go`, package README             | Cross-provider catalog policy or public contract changes                            |
| `modelpreset/`           | Static provider/model connection defaults and capability patches                                                                | `catalog.go`, `provider_*.go`                           | Runtime probing or new wire protocol implementations                                |
| `capabilityoverride/`    | Patch representation, validation, cloning, layering, resolver implementation                                                    | `type.go`, `apply.go`, `resolver.go`                    | Complete adapter base profiles or vendor conversion                                 |
| `debugclient/`           | HTTP instrumentation, redaction, scrubbed debug payloads, safe streaming transport                                              | `debugger.go`, `http_transport.go`, `scrubber.go`       | Provider semantics or request normalization                                         |
| `internal/integration/`  | Credential-gated executable examples and end-to-end behavior references                                                         | `helpers_test.go`, `example_*.go`                       | Unit-only conversion edge cases                                                     |
| `docs/`                  | Current support and capability policy                                                                                           | `docs/current_capability_ref.md`                        | Historical behavior; ignore `docs/archive/`                                         |

`internal/logutil/` owns the process-wide SDK logger. `taskfile.yml` groups focused catalog and adapter test commands.

## Request lifecycle and ownership

1. Catalog/application setup: `modelpreset` supplies cloned presets. `ProviderSetAPI.AddProviderFromPreset` builds a transport configuration and dispatches to one adapter.
2. Authentication/lifecycle: `SetProviderAPIKey` stores the key and initializes that adapter’s vendor client. Adapter initialization owns origin, path-prefix, header, and debugger HTTP-client translation.
3. Capability resolution: `ProviderSetAPI.NewPresetCapabilityResolver` derives effective capabilities in this order: adapter base -> provider override -> model override. A caller can derive a further override. `CompletionKeyResolver` binds the profile to the request’s completion key.
4. Pre-adapter safety: `ProviderSetAPI.FetchCompletion` filters provider-native reasoning by endpoint fingerprint and applies optional heuristic prompt trimming.
5. Shared normalization: every adapter calls `sdkutil.NormalizeRequestForSDK` before conversion. It clones the request, sanitizes conversation start, validates modalities, errors for unsupported contract-like output formats, and warns while dropping/narrowing safe optional settings.
6. Adapter conversion: the selected adapter converts normalized inputs, model parameters, tools, and policy to its vendor SDK; it owns provider-native reasoning and stream event handling.
7. Response boundary: the adapter returns `spec.FetchCompletionResponse` with normalized outputs, usage, warnings, optional partial error, and debug details. `ProviderSetAPI` stamps its route fingerprint on reasoning outputs.

Implications:

- Treat caller requests and catalog values as immutable. Normalize/clone rather than edit them in place.
- Preserve warnings; they are the explicit record of safe capability reductions.
- Preserve provider-native reasoning only on the same compatible route. Ordinary text history is the portable fallback.
- A `StreamHandler` activates streaming only when `ModelParam.Stream` is also true.

## Put a change in the right bucket

| If the change is                                                   | Primary location                                                           | Also review                                                                           |
| ------------------------------------------------------------------ | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| A portable public field or union variant                           | `spec/`                                                                    | all adapters, cloning, normalizer, contract metadata/hash, docs, integration examples |
| A shared validation/drop/warning rule                              | `internal/sdkutil/`                                                        | each adapter capability and focused normalization tests                               |
| A vendor wire representation or response mapping                   | matching `internal/<adapter>sdk/`                                          | its `capability.go`, README, streaming and conversion tests                           |
| A provider/model endpoint, limits, or verified feature restriction | `modelpreset/provider_*.go`                                                | catalog registration/manifests, capability reference, catalog tests                   |
| A deployment-specific difference                                   | application profile plus `capabilityoverride`                              | the `custom-providers-and-models` skill; do not change built-in metadata              |
| A partial capability-patch behavior                                | `capabilityoverride/`                                                      | validation, cloning, JSON behavior, resolver tests                                    |
| HTTP capture/redaction/transport behavior                          | `debugclient/`                                                             | debugger span integration in every adapter as applicable                              |
| A new wire protocol                                                | new `internal/<adapter>sdk/` plus `spec.ProviderSDKType` and root dispatch | use `add-sdk-support`; it is not a preset-only change                                 |

For a compatible endpoint, start with a preset/override. Do not add an adapter family merely because the endpoint has a different provider name. Conversely, do not force a genuinely different protocol into an existing adapter.

## Capability model

`spec.ModelCapabilities` is the effective, complete profile consumed by normalization. `capabilityoverride.ModelCapabilitiesOverride` is intentionally only a patch:

- nil slice or scalar pointer: inherit;
- non-nil empty slice: explicitly clear inherited support;
- later patches override only supplied fields.

Base capabilities belong in `internal/<adapter>sdk/capability.go`. Provider-wide and model-specific restrictions belong in `modelpreset`. `capabilityoverride.ValidateModelCapabilitiesOverride` validates patch structure, not whether a partial patch is a complete profile. See `docs/current_capability_ref.md` for actual adapter behavior and warning/error policy.

Use a resolver whenever model or route behavior differs from an adapter baseline. The request `CompletionKey` must match the resolver’s key.

## Adapter families

| Family                  | Package                              | Native continuation data / notable boundary                                                     |
| ----------------------- | ------------------------------------ | ----------------------------------------------------------------------------------------------- |
| Anthropic Messages      | `internal/anthropicsdk/`             | Signed or redacted thinking; strict client tool-result ordering                                 |
| OpenAI Responses        | `internal/openairesponsessdk/`       | Encrypted reasoning; top-level prompt cache controls                                            |
| OpenAI Chat Completions | `internal/openaichatsdk/`            | No structured reasoning history or thinking stream; dialect support lives in `param_dialect.go` |
| Google Generate Content | `internal/googlegeneratecontentsdk/` | Google thought signatures; grounding is server-side web search                                  |

Read the adapter README before changing behavior. A shared `spec` feature may intentionally have different mapping, warning, or unsupported status across these families.

## Debugging route

Follow request order; do not begin at the vendor SDK.

1. Registration/auth/client failure: inspect `provider_set.go`, the selected preset’s origin/path/header fields, then that adapter’s `InitLLM` and `SetProviderAPIKey`.
2. Unexpected rejection, dropped setting, or missing tool: inspect response `Warnings`, then `sdkutil.NormalizeRequestForSDK`, the effective resolver profile, and relevant overrides. Check the capability reference before changing conversion code.
3. Wrong vendor payload: inspect the adapter’s conversion path and enable `debugclient.NewHTTPCompletionDebugger`. `DebugDetails` is intentionally scrubbed; do not weaken redaction casually.
4. Unexpected normalized output: inspect the adapter’s `outputsFrom...` conversion and the matching unit tests. Keep provider data provider-native until that boundary.
5. Continuation/tool-history failure: inspect `sdkutil/reasoning_continuation.go`, conversation-start sanitation, adapter reasoning sanitizer, and adapter-specific tool ordering rules.
6. Streaming-only failure: inspect the adapter `doStreaming`, `sdkutil/throttle_stream.go`, handler errors/panics, terminal-event checks, and stream transport behavior.
7. Catalog inconsistency: use `modelpreset/catalog.go`, `catalog_model_membership_test.go`, and `provider_set_preset_test.go`; a passing catalog test proves metadata consistency, not live endpoint support.

## Change invariants

- Keep `spec` provider-neutral and adapters internal.
- Normalize before vendor conversion; preserve returned warnings.
- Never mutate the caller request or shared catalog state.
- Keep model defaults compatible with their derived capabilities.
- Keep tool call IDs/order and only replay native reasoning on a matching route.
- Treat unsupported output formats and unsupported tool-policy modes as errors; only safely reducible controls become warnings.
- When a data-contract file changes, update `internal/sdkutil/data_contract_meta.go` deliberately and run its validation.
- Update docs and the closest focused tests with behavior changes; use live integration coverage to verify external claims.

## Existing task skills

- Built-in provider/model catalog work: `add-model-preset`.
- Application-owned compatible deployments: `custom-providers-and-models`.
- New or extended adapter family: `add-sdk-support`.
- Consumer setup and tool/conversation usage: `inference-go-usage`.
