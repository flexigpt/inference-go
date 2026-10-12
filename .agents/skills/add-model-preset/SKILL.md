---
name: add-model-preset
description: Add or update a built-in provider or model preset in inference-go, including catalog registration and capability metadata.
---

# Add a built-in model preset

Use this for adding support in `modelpreset`.

## Choose the change

| Need                                       | Change                                                 |
| ------------------------------------------ | ------------------------------------------------------ |
| Model on a catalog provider                | Edit that provider's `modelpreset/provider_*.go`.      |
| New provider using an existing wire family | Add a provider file, models, and catalog registration. |
| New wire protocol                          | Do not fake it as a preset; add adapter support first. |

Supported adapter families are defined in `spec/provider.go`: Anthropic Messages, OpenAI Chat Completions, OpenAI Responses, and Google Generate Content.

## Add a model to an existing provider

1. Read the target `modelpreset/provider_<provider>.go`, a nearby comparable model, and its adapter capability file under `internal/*sdk/capability.go`.
2. Add stable provider-scoped identifiers: model name, display name, and `ModelPresetID`. Put shared identities in `modelpreset/other_const.go` only when genuinely reused by multiple provider files.
3. Declare one `ModelPreset` and add it to the provider's `ModelPresets` map.
4. Set reviewed `ModelParam` defaults: served model name, streaming preference, prompt/output limits, timeout, system prompt, and either temperature or reasoning. Add cache, output, or stop defaults only when verified.
5. Add a model capability override only for differences from the provider/adapter baseline. Test the exact route before advertising a feature.

`ValidateCatalog` requires nonblank IDs/names/display names, matching map key and `ID`, matching `Name` and `ModelParam.Name`, a temperature or reasoning default, nonnegative limits/timeouts, valid optional JSON, and structurally valid capability data. See `modelpreset/catalog.go`.

## Add a provider and models

1. Choose a compatible existing `ProviderSDKType`; inspect a provider using that same adapter for URL and header behavior. A compatible endpoint is not proof of full adapter support.
2. Create `modelpreset/provider_<provider>.go` with:
   - provider name/display constants;
   - `ProviderPreset` connection fields (`SDKType`, origin, path prefix, API-key header, default headers);
   - provider-wide capability restrictions;
   - at least one model and its `ModelPresets` entry.
3. Register the provider in `catalogProviders` in `modelpreset/catalog.go`.
4. Add model declarations as above. Use the exact deployed model ID; router/backend-qualified IDs are separate identities when behavior differs.

Catalog validation requires an HTTP(S) origin, a nonempty path starting with `/`, a nonempty API-key header key, and at least one model. Copy adapter-specific path conventions from an existing compatible preset: adapters remove their terminal endpoint segment before configuring their upstream SDK.

## Capability rules

Effective capabilities are layered as:

1. adapter base capability;
2. provider override;
3. model override;
4. caller override.

`capabilityoverride.ModelCapabilitiesOverride` is a patch, not a complete profile. Nil slices/pointers inherit; an explicit empty slice disables the inherited list. Be precise for modalities, reasoning types/levels, output formats, tools, cache scopes, and parameter dialects. Review `capabilityoverride/type.go`, `apply.go`, and `validate.go`; use `docs/current_capability_ref.md` as the behavior reference.

Do not claim support merely because another model on the same provider supports it. Incorrect metadata can make normalization send fields the endpoint rejects.

## Required review updates

- Update `modelpreset/catalog_model_membership_test.go` for every added/removed/renamed model.
- For a new provider, update `TestCatalogContainsAllRegisteredProviders` and any affected explicit membership test in `modelpreset/catalog_test.go`.
- Keep catalog validation and ProviderSet integration passing; `provider_set_preset_test.go` verifies every catalog model's defaults against effective capabilities.
- Update `modelpreset/README.md`; update root `README.md` when provider support changes.
- Update relevant matrices/caveats in `docs/current_capability_ref.md` when capability posture changes.
- Add or extend an integration example when connection setup or public behavior is novel.

## Verify

Run the catalog test target in `taskfile.yml` (`task test:catalog`), then focused adapter and live integration coverage where credentials exist. Check both returned errors and normalization warnings. A passing catalog test proves metadata consistency, not vendor behavior.

## Final checklist

- Stable, whitespace-free provider/model IDs.
- Exact provider model ID and route settings.
- Defaults compatible with the derived capability profile.
- Conservative, verified overrides.
- Catalog registration, manifests, docs, and tests updated together.
