package capabilityoverride

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestCloneModelCapabilitiesTable(t *testing.T) {
	tests := []struct {
		name        string
		in          spec.ModelCapabilities
		want        spec.ModelCapabilities
		mutateInput func(in *spec.ModelCapabilities)
	}{
		{
			name: "zero value",
			in:   spec.ModelCapabilities{},
			want: spec.ModelCapabilities{},
		},
		{
			name:        "full capabilities deep clone",
			in:          baseCapabilitiesForTest(),
			want:        baseCapabilitiesForTest(),
			mutateInput: mutateModelCapabilitiesForCloneTest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CloneModelCapabilities(tc.in)
			assertDeepEqual(t, got, tc.want)

			if tc.mutateInput != nil {
				tc.mutateInput(&tc.in)
				assertDeepEqual(t, got, tc.want)
			}
		})
	}
}

func TestCloneModelCapabilitiesOverrideTable(t *testing.T) {
	fullIn := fullCapabilitiesOverrideForTest()
	fullWant := fullCapabilitiesOverrideForTest()

	partialMaxOnly := &ModelCapabilitiesOverride{
		ParamDialect: &ParamDialectOverride{
			MaxOutputTokensParamName: new(
				spec.MaxOutputTokensParamNameMaxTokens,
			),
		},
	}

	partialToolChoiceOnly := &ModelCapabilitiesOverride{
		ParamDialect: &ParamDialectOverride{
			ToolChoiceParamStyle: new(
				spec.ToolChoiceParamStyleRequiredNamed,
			),
		},
	}

	emptySlices := &ModelCapabilitiesOverride{
		ModalitiesIn:  []spec.Modality{},
		ModalitiesOut: []spec.Modality{},
		ReasoningCapabilities: &ReasoningCapabilitiesOverride{
			SupportedReasoningTypes:  []spec.ReasoningType{},
			SupportedReasoningLevels: []spec.ReasoningLevel{},
		},
		OutputCapabilities: &OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{},
		},
		ToolCapabilities: &ToolCapabilitiesOverride{
			SupportedToolTypes:               []spec.ToolType{},
			SupportedToolPolicyModes:         []spec.ToolPolicyMode{},
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{},
		},
		CacheCapabilities: &CacheCapabilitiesOverride{
			TopLevel: &CacheControlCapabilitiesOverride{
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
			},
		},
	}

	tests := []struct {
		name        string
		in          *ModelCapabilitiesOverride
		mutateInput func(in *ModelCapabilitiesOverride)
		check       func(t *testing.T, got *ModelCapabilitiesOverride)
	}{
		{
			name: "nil override",
			in:   nil,
			check: func(t *testing.T, got *ModelCapabilitiesOverride) {
				t.Helper()

				if got != nil {
					t.Fatalf("expected nil clone, got %#v", got)
				}
			},
		},
		{
			name:        "full override deep clone",
			in:          fullIn,
			mutateInput: mutateCapabilitiesOverrideForCloneTest,
			check: func(t *testing.T, got *ModelCapabilitiesOverride) {
				t.Helper()

				assertDeepEqual(t, got, fullWant)

				if got == fullIn {
					t.Fatal("expected different top-level pointer")
				}
				if got.ReasoningCapabilities == fullIn.ReasoningCapabilities {
					t.Fatal("expected different reasoning capabilities pointer")
				}
				if got.ReasoningCapabilities.HybridTokenBudgetCapabilities ==
					fullIn.ReasoningCapabilities.HybridTokenBudgetCapabilities {
					t.Fatal("expected different hybrid token budget capabilities pointer")
				}
				if got.ReasoningCapabilities.SupportsReasoningConfig ==
					fullIn.ReasoningCapabilities.SupportsReasoningConfig {
					t.Fatal("expected cloned reasoning bool pointer")
				}
				if got.StopSequenceCapabilities == fullIn.StopSequenceCapabilities {
					t.Fatal("expected different stop sequence capabilities pointer")
				}
				if got.OutputCapabilities == fullIn.OutputCapabilities {
					t.Fatal("expected different output capabilities pointer")
				}
				if got.ToolCapabilities == fullIn.ToolCapabilities {
					t.Fatal("expected different tool capabilities pointer")
				}
				if got.CacheCapabilities == fullIn.CacheCapabilities {
					t.Fatal("expected different cache capabilities pointer")
				}
				if got.CacheCapabilities.TopLevel == fullIn.CacheCapabilities.TopLevel {
					t.Fatal("expected different cache top-level pointer")
				}
				if got.ParamDialect == fullIn.ParamDialect {
					t.Fatal("expected different param dialect pointer")
				}
				if got.ParamDialect.MaxOutputTokensParamName ==
					fullIn.ParamDialect.MaxOutputTokensParamName {
					t.Fatal("expected cloned maxOutputTokensParamName pointer")
				}
			},
		},
		{
			name: "partial param dialect max only",
			in:   partialMaxOnly,
			check: func(t *testing.T, got *ModelCapabilitiesOverride) {
				t.Helper()

				if got == nil || got.ParamDialect == nil {
					t.Fatal("expected cloned param dialect")
				}
				if got.ParamDialect.MaxOutputTokensParamName == nil {
					t.Fatal("expected cloned maxOutputTokensParamName")
				}
				if *got.ParamDialect.MaxOutputTokensParamName !=
					spec.MaxOutputTokensParamNameMaxTokens {
					t.Fatalf("unexpected maxOutputTokensParamName: %q", *got.ParamDialect.MaxOutputTokensParamName)
				}
				if got.ParamDialect.ToolChoiceParamStyle != nil {
					t.Fatalf("expected nil toolChoiceParamStyle, got %#v", got.ParamDialect.ToolChoiceParamStyle)
				}
				if got.ParamDialect.MaxOutputTokensParamName ==
					partialMaxOnly.ParamDialect.MaxOutputTokensParamName {
					t.Fatal("expected cloned maxOutputTokensParamName pointer")
				}
			},
		},
		{
			name: "partial param dialect tool choice only",
			in:   partialToolChoiceOnly,
			check: func(t *testing.T, got *ModelCapabilitiesOverride) {
				t.Helper()

				if got == nil || got.ParamDialect == nil {
					t.Fatal("expected cloned param dialect")
				}
				if got.ParamDialect.ToolChoiceParamStyle == nil {
					t.Fatal("expected cloned toolChoiceParamStyle")
				}
				if *got.ParamDialect.ToolChoiceParamStyle !=
					spec.ToolChoiceParamStyleRequiredNamed {
					t.Fatalf("unexpected toolChoiceParamStyle: %q", *got.ParamDialect.ToolChoiceParamStyle)
				}
				if got.ParamDialect.MaxOutputTokensParamName != nil {
					t.Fatalf(
						"expected nil maxOutputTokensParamName, got %#v",
						got.ParamDialect.MaxOutputTokensParamName,
					)
				}
				if got.ParamDialect.ToolChoiceParamStyle ==
					partialToolChoiceOnly.ParamDialect.ToolChoiceParamStyle {
					t.Fatal("expected cloned toolChoiceParamStyle pointer")
				}
			},
		},
		{
			name: "empty slices remain non-nil empty slices",
			in:   emptySlices,
			check: func(t *testing.T, got *ModelCapabilitiesOverride) {
				t.Helper()

				if got.ModalitiesIn == nil || len(got.ModalitiesIn) != 0 {
					t.Fatalf("expected non-nil empty modalitiesIn, got %#v", got.ModalitiesIn)
				}
				if got.ModalitiesOut == nil || len(got.ModalitiesOut) != 0 {
					t.Fatalf("expected non-nil empty modalitiesOut, got %#v", got.ModalitiesOut)
				}
				if got.ReasoningCapabilities.SupportedReasoningTypes == nil ||
					len(got.ReasoningCapabilities.SupportedReasoningTypes) != 0 {
					t.Fatalf("expected non-nil empty supportedReasoningTypes, got %#v",
						got.ReasoningCapabilities.SupportedReasoningTypes)
				}
				if got.ReasoningCapabilities.SupportedReasoningLevels == nil ||
					len(got.ReasoningCapabilities.SupportedReasoningLevels) != 0 {
					t.Fatalf("expected non-nil empty supportedReasoningLevels, got %#v",
						got.ReasoningCapabilities.SupportedReasoningLevels)
				}
				if got.OutputCapabilities.SupportedOutputFormats == nil ||
					len(got.OutputCapabilities.SupportedOutputFormats) != 0 {
					t.Fatalf("expected non-nil empty supportedOutputFormats, got %#v",
						got.OutputCapabilities.SupportedOutputFormats)
				}
				if got.ToolCapabilities.SupportedToolTypes == nil ||
					len(got.ToolCapabilities.SupportedToolTypes) != 0 {
					t.Fatalf("expected non-nil empty supportedToolTypes, got %#v",
						got.ToolCapabilities.SupportedToolTypes)
				}
				if got.ToolCapabilities.SupportedToolPolicyModes == nil ||
					len(got.ToolCapabilities.SupportedToolPolicyModes) != 0 {
					t.Fatalf("expected non-nil empty supportedToolPolicyModes, got %#v",
						got.ToolCapabilities.SupportedToolPolicyModes)
				}
				if got.ToolCapabilities.SupportedClientToolOutputFormats == nil ||
					len(got.ToolCapabilities.SupportedClientToolOutputFormats) != 0 {
					t.Fatalf("expected non-nil empty supportedClientToolOutputFormats, got %#v",
						got.ToolCapabilities.SupportedClientToolOutputFormats)
				}
				if got.CacheCapabilities.TopLevel.SupportedKinds == nil ||
					len(got.CacheCapabilities.TopLevel.SupportedKinds) != 0 {
					t.Fatalf("expected non-nil empty supportedKinds, got %#v",
						got.CacheCapabilities.TopLevel.SupportedKinds)
				}
				if got.CacheCapabilities.TopLevel.SupportedTTLs == nil ||
					len(got.CacheCapabilities.TopLevel.SupportedTTLs) != 0 {
					t.Fatalf("expected non-nil empty supportedTTLs, got %#v",
						got.CacheCapabilities.TopLevel.SupportedTTLs)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CloneModelCapabilitiesOverride(tc.in)
			tc.check(t, got)

			if tc.mutateInput != nil {
				tc.mutateInput(tc.in)
				tc.check(t, got)
			}
		})
	}
}

func mutateModelCapabilitiesForCloneTest(in *spec.ModelCapabilities) {
	in.ModalitiesIn[0] = spec.ModalityFileIn
	in.ModalitiesOut[0] = spec.ModalityFileOut

	in.ReasoningCapabilities.SupportedReasoningTypes[0] = spec.ReasoningTypeSingleWithLevels
	in.ReasoningCapabilities.SupportedReasoningLevels[0] = spec.ReasoningLevelMax
	in.ReasoningCapabilities.HybridTokenBudgetCapabilities.MinAllowed = 99

	in.StopSequenceCapabilities.MaxSequences = 99

	in.OutputCapabilities.SupportedOutputFormats[0] = spec.OutputFormatKindJSONSchema

	in.ToolCapabilities.SupportedToolTypes[0] = spec.ToolTypeCustom
	in.ToolCapabilities.SupportedToolPolicyModes[0] = spec.ToolPolicyModeNone
	in.ToolCapabilities.SupportedClientToolOutputFormats[0] = spec.ToolOutputFormatKindContentItemList

	in.CacheCapabilities.TopLevel.SupportedTTLs[0] = spec.CacheControlTTL24h

	in.ParamDialect.ToolChoiceParamStyle = spec.ToolChoiceParamStyleRequiredNamed
}

func mutateCapabilitiesOverrideForCloneTest(in *ModelCapabilitiesOverride) {
	in.ModalitiesIn[0] = spec.ModalityAudioIn
	in.ModalitiesOut[0] = spec.ModalityFileOut

	*in.ReasoningCapabilities.SupportsReasoningConfig = false
	in.ReasoningCapabilities.SupportedReasoningTypes[0] = spec.ReasoningTypeSingleWithLevels
	in.ReasoningCapabilities.SupportedReasoningLevels[0] = spec.ReasoningLevelHigh
	*in.ReasoningCapabilities.HybridTokenBudgetCapabilities.MinAllowed = 99
	*in.ReasoningCapabilities.HybridTokenBudgetCapabilities.MaxAllowed = 100
	*in.ReasoningCapabilities.HybridTokenBudgetCapabilities.ZeroAllowed = false
	*in.ReasoningCapabilities.HybridTokenBudgetCapabilities.MinusOneAllowed = false
	*in.ReasoningCapabilities.SupportsSummaryStyle = false
	*in.ReasoningCapabilities.SupportsEncryptedReasoningInput = false
	*in.ReasoningCapabilities.TemperatureDisallowedWhenEnabled = false

	*in.StopSequenceCapabilities.IsSupported = false
	*in.StopSequenceCapabilities.DisallowedWithReasoning = false
	*in.StopSequenceCapabilities.MaxSequences = 99

	in.OutputCapabilities.SupportedOutputFormats[0] = spec.OutputFormatKindJSONSchema
	*in.OutputCapabilities.SupportsVerbosity = false

	in.ToolCapabilities.SupportedToolTypes[0] = spec.ToolTypeCustom
	in.ToolCapabilities.SupportedToolPolicyModes[0] = spec.ToolPolicyModeNone
	in.ToolCapabilities.SupportedClientToolOutputFormats[0] = spec.ToolOutputFormatKindContentItemList
	*in.ToolCapabilities.SupportsParallelToolCalls = false
	*in.ToolCapabilities.MaxForcedTools = 99

	*in.CacheCapabilities.SupportsAutomaticCaching = false
	*in.CacheCapabilities.TopLevel.SupportsTTL = false
	in.CacheCapabilities.TopLevel.SupportedKinds[0] = spec.CacheControlKind("mutated")
	in.CacheCapabilities.TopLevel.SupportedTTLs[0] = spec.CacheControlTTL24h
	*in.CacheCapabilities.TopLevel.SupportsKey = false

	*in.ParamDialect.MaxOutputTokensParamName = spec.MaxOutputTokensParamNameMaxCompletionTokens
	*in.ParamDialect.ToolChoiceParamStyle = spec.ToolChoiceParamStyleAllowedTools
}
