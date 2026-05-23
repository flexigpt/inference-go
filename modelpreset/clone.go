package modelpreset

import (
	"maps"
	"slices"

	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

func CloneCatalog(in Catalog) Catalog {
	out := Catalog{
		Providers: make(map[spec.ProviderName]ProviderPreset, len(in.Providers)),
	}
	for k, v := range in.Providers {
		out.Providers[k] = CloneProviderPreset(v)
	}
	return out
}

func CloneProviderPreset(in ProviderPreset) ProviderPreset {
	out := in
	out.DefaultHeaders = maps.Clone(in.DefaultHeaders)
	out.CapabilitiesOverride = capabilityoverride.CloneModelCapabilitiesOverride(in.CapabilitiesOverride)
	out.ModelPresets = make(map[ModelPresetID]ModelPreset, len(in.ModelPresets))
	for k, v := range in.ModelPresets {
		out.ModelPresets[k] = CloneModelPreset(v)
	}
	return out
}

func CloneModelPreset(in ModelPreset) ModelPreset {
	out := in
	out.ModelParam = cloneModelParam(in.ModelParam)
	out.CapabilitiesOverride = capabilityoverride.CloneModelCapabilitiesOverride(in.CapabilitiesOverride)
	return out
}

func cloneModelParam(in spec.ModelParam) spec.ModelParam {
	out := in
	out.Temperature = sdkutil.CloneFloat64Ptr(in.Temperature)
	out.Reasoning = cloneReasoningParam(in.Reasoning)
	out.CacheControl = cloneCacheControl(in.CacheControl)
	out.OutputParam = cloneOutputParam(in.OutputParam)
	out.StopSequences = slices.Clone(in.StopSequences)
	out.AdditionalParametersRawJSON = sdkutil.CloneStringPtr(in.AdditionalParametersRawJSON)
	return out
}

func cloneCacheControl(in *spec.CacheControl) *spec.CacheControl {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func cloneReasoningParam(in *spec.ReasoningParam) *spec.ReasoningParam {
	if in == nil {
		return nil
	}
	out := *in
	if in.SummaryStyle != nil {
		v := *in.SummaryStyle
		out.SummaryStyle = &v
	}
	return &out
}

func cloneOutputParam(in *spec.OutputParam) *spec.OutputParam {
	if in == nil {
		return nil
	}
	out := *in
	if in.Verbosity != nil {
		v := *in.Verbosity
		out.Verbosity = &v
	}
	if in.Format != nil {
		f := *in.Format
		if f.JSONSchemaParam != nil {
			j := *f.JSONSchemaParam
			j.Schema = maps.Clone(j.Schema)
			f.JSONSchemaParam = &j
		}
		out.Format = &f
	}
	return &out
}
