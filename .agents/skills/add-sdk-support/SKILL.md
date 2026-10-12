---
name: add-sdk-support
description: Add or extend an inference-go adapter for a provider API, including normalized contracts, capabilities, dispatch, streaming, debugging and presets.
---

# Add SDK / adapter support

Keep public contracts provider-neutral. Do not expose vendor SDK types from `package inference` or `spec`.

## Decide first

| Endpoint change                                                   | Correct scope                                       |
| ----------------------------------------------------------------- | --------------------------------------------------- |
| Existing adapter accepts it unchanged                             | Provider/model preset and capability metadata only. |
| Existing adapter needs a small compatible dialect/behavior change | Extend that adapter, capabilities, tests, and docs. |
| Different wire protocol                                           | Add a new adapter family and dispatch support.      |

- E.g.: Mistral-style compatible variant is added as a dialect. Inspect `modelpreset/provider_mistral.go`, `capabilityoverride.ParamDialectOverride`, and `internal/openaichatsdk/param_dialect.go` to check the flow.
- A new API such as an Google Gemini Interactions protocol needs to be added as a full adapter path below.

## Establish the normalized contract

1. Read `spec/*.go`, `docs/current_capability_ref.md`, and the closest existing adapter README.
2. Decide which existing normalized inputs, outputs, tools, reasoning, cache controls, and stream events can be represented faithfully.
3. Do not add arbitrary provider-specific passthrough merely to reach feature parity. If a feature must become public and portable, make a deliberate `spec` contract change and update every adapter/documentation implication.
4. If changing a file listed by `internal/sdkutil/data_contract_meta.go`, deliberately update the data-contract hash and version policy; run its validation tests.

## Implement an adapter family

Create `internal/<adapter>sdk/`, modeled on the closest of:

- `internal/anthropicsdk`
- `internal/openaichatsdk`
- `internal/openairesponsessdk`
- `internal/googlegeneratecontentsdk`

Implement `sdkutil.CompletionProvider` from `internal/sdkutil/completion.go`:

- constructor and provider-info ownership;
- API-key set/init/deinit lifecycle;
- custom origin, path prefix, headers, and API-key-header handling;
- immutable provider capability return;
- normalized fetch/convert/response path.

Add dependencies deliberately in `go.mod`; keep vendor types inside the internal adapter.

## Normalize before conversion

Every adapter must call `sdkutil.NormalizeRequestForSDK` before constructing vendor requests. Honor `FetchCompletionOptions.CapabilityResolver`, use the effective profile, preserve returned warnings, and do not mutate the caller request.

Define adapter base capabilities in `internal/<adapter>sdk/capability.go`. Model/provider differences belong in catalog or caller overrides, not a guessed universal base. Cover modalities, reasoning, stop sequences, output formats/verbosity, tools/tool-output transport, caching, and dialect fields.

Use existing conversion code as the contract map:

- input/output unions: each adapter's `api_*.go` and input-processing files;
- tool-name/policy rules: `internal/sdkutil/tool_choice_map.go`;
- cache normalization: `internal/sdkutil/capability_cache_control_normalize.go`;
- tool-output normalization: `internal/sdkutil/capability_tool_output_normalize.go`.

## Streaming, reasoning, and debugging

- Stream only when both `ModelParam.Stream` and an options stream handler are set.
- Emit normalized text/thinking events with provider, model, and completion key; use `SafeCallStreamHandler` and `ResolveStreamConfig`/buffering helpers.
- Return the final normalized response and usage after streaming, including partial response/error where possible.
- Preserve only provider-native reasoning continuation data. Inspect the Anthropic, OpenAI Responses, and Google sanitizers; never replay incompatible signatures, plaintext reasoning, or encrypted data across families.
- Integrate `CompletionDebugger`: obtain its HTTP client at initialization, start/end a span per request, and attach debug details exactly as existing adapters do.

`ProviderSetAPI` handles route fingerprints around adapter calls, but the adapter still owns native reasoning conversion and sanitization.

## Wire the public dispatcher

For a new protocol:

1. Add its `ProviderSDKType` in `spec/provider.go`.
2. Add it to `isProviderSDKTypeSupported` and `getProviderAPI` in `provider_set.go`.
3. Add it to `isCatalogProviderSDKType` in `modelpreset/catalog.go` before shipping presets.
4. Add the internal constructor/import and provider lifecycle coverage.
5. Add a catalog provider/model only after adapter dispatch works.

For a compatible-provider extension, do not add a new SDK type unless the wire behavior truly differs.

## Tests and documentation

- Add focused tests for request conversion, response conversion, capability normalization, malformed/unsupported inputs, reasoning continuation, tool round trips, cache controls, and streaming handler failure/panic behavior.
- Add credential-gated end-to-end coverage under `internal/integration` for a new public path. Run the relevant adapter tests, `task test:catalog` when presets change, and the repository suite.
- Inspect the `taskfile.yml` to add tasks for new adapter.

Update all affected documentation:

- adapter README under `internal/<adapter>sdk/`;
- root `README.md` support tables;
- `docs/current_capability_ref.md` adapter/preset matrices and caveats;
- `modelpreset/README.md` and catalog tests for shipped presets.

## Completion gate

- No vendor types escaped into public normalized APIs.
- Normalizer, warnings, capabilities, conversion, and response shapes agree.
- Custom endpoint/auth behavior is tested.
- Streaming/debugging/reasoning behavior matches existing adapter conventions.
- Catalog metadata and docs state only verified support.
