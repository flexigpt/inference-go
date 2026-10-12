---
name: custom-providers-and-models
description: Configure application-owned compatible providers and model profiles with inference-go, including capability resolution, registration, validation, and safe rollout.
---

# Use application-defined providers and models

Use this when an application has its own endpoint, deployment, router route, or model ID. This does not modify the built-in catalog.

## Boundary

A consumer-defined provider must use one of the four existing adapter families in `spec/provider.go`. `ProviderSetAPI` cannot register a new wire protocol at runtime. For a different protocol, library adapter work is required.

A compatible endpoint can still differ from the adapter baseline. Supply conservative provider/model capability metadata and verify advanced features against the deployed route.

## Preferred setup: application-owned presets

Create application-owned `modelpreset.ProviderPreset` and `modelpreset.ModelPreset` values. Keep the model in the provider's `ModelPresets` map, then validate an application `modelpreset.Catalog` containing that provider with `modelpreset.ValidateCatalog`.

### Provider profile

Provide a unique `Name`, display name, supported `SDKType`, origin, endpoint path prefix, API-key header key, headers, optional provider-wide capability override, and at least one model. Reuse connection conventions from a built-in provider using the same adapter; adapters handle terminal endpoint segments differently.

### Model profile

Provide a stable `ModelPresetID`, exact served `ModelName`, display name, `ModelParam`, and optional model override. `Name` must equal `ModelParam.Name`; supply temperature or reasoning; keep limits/timeouts nonnegative. Start with known-good text defaults, then add modalities, tools, JSON Schema, cache, and reasoning only after testing.

`modelpreset.Provider` and `modelpreset.Model` return clones, so a built-in preset can be copied and adjusted for one deployment. Do not expect application profiles to appear in `DefaultCatalog`, `ProviderNames`, or built-in lookup APIs.

## Runtime sequence

1. Create `inference.ProviderSetAPI`.
2. Register the provider with `AddProviderFromPreset`. The runtime provider name must be unique; a blank name uses `ProviderPreset.Name`.
3. Call `SetProviderAPIKey`; this initializes the adapter. Use the same runtime name for every fetch.
4. Build a resolver with `NewPresetCapabilityResolver`, passing that runtime name, the application provider/model profiles, and a stable completion key (normally the model preset ID).
5. Put both that completion key and resolver in `spec.FetchCompletionOptions`.
6. Send `FetchCompletionRequest` values using the model's `ModelParam`; inspect response outputs, warnings, and errors.

The completion key passed in options must match the resolver key. A mismatch intentionally fails capability resolution. Create a distinct resolver/key for each model or deployment profile.

## Capability metadata

The resolver derives:

1. selected adapter base capability;
2. provider override;
3. model override.

Override semantics matter: nil means inherit, while an explicit empty slice clears inherited support. See `capabilityoverride/type.go`, `apply.go`, and `validate.go`.

For routers and local servers, restrict the broad adapter baseline to what the selected model and route actually accept. In particular review input modalities, reasoning types/levels, output format, tool type/policy, client tool-output format, cache scopes, and parameter dialect. `docs/current_capability_ref.md` explains normalized behavior and warning/error policy.

A resolver is not optional for model-specific behavior. Without one, normalization uses the adapter's base capability profile and may overstate your endpoint.

## Existing provider, new deployment/model

Copy the built-in provider preset, modify connection fields/headers only as needed, create an application model profile with the exact served name and restrictions, then follow the runtime sequence. Keep the runtime provider name distinct if a built-in provider is already registered in the same `ProviderSetAPI`.

## Minimal alternative: direct registration

`AddProvider` accepts `inference.AddProviderConfig` when a preset is inconvenient. It configures transport only. Immediately obtain the adapter base profile with `GetProviderCapability`, derive the desired profile with `capabilityoverride.DeriveModelCapabilities`, and use `capabilityoverride.NewCompletionKeyResolver` in fetch options. This is more manual; application-owned presets plus `NewPresetCapabilityResolver` are usually clearer and validate better.

## Rollout and diagnosis

- Validate the application catalog before registration; validation is structural, not a remote probe.
- Smoke-test text first, then each claimed advanced feature independently.
- Treat unsupported output formats/tool policy as errors and safe reductions as `FetchCompletionResponse.Warnings`.
- Use `WithDebugClientBuilder` and `debugclient.NewHTTPCompletionDebugger` to inspect scrubbed request/response metadata when endpoint setup fails.
- Preserve provider-native reasoning only on compatible continuations. The library fingerprints routes and drops mismatched reasoning; never fabricate signatures or encrypted reasoning.

## Checklist

- Compatible adapter family selected.
- Runtime provider name, resolver key, and fetch options are consistent.
- Exact endpoint, path, headers, and model ID verified.
- Capability patches are conservative and tested.
- Application catalog validation and live behavior checks pass.
