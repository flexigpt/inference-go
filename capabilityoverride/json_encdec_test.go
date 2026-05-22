package capabilityoverride

import (
	"encoding/json"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestMarshalJSONPreservesEmptySliceOverridesTable(t *testing.T) {
	tests := []struct {
		name  string
		value any
		check func(t *testing.T, m map[string]any)
	}{
		{
			name:  "empty model override omits nil fields",
			value: ModelCapabilitiesOverride{},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				if len(m) != 0 {
					t.Fatalf("expected empty JSON object, got %#v", m)
				}
			},
		},
		{
			name: "model top-level empty slices are emitted",
			value: ModelCapabilitiesOverride{
				ModalitiesIn:  []spec.Modality{},
				ModalitiesOut: []spec.Modality{},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				assertJSONArrayLen(t, m, "modalitiesIn", 0)
				assertJSONArrayLen(t, m, "modalitiesOut", 0)
			},
		},
		{
			name: "model nested empty slices and zero scalar pointers are emitted",
			value: ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportsReasoningConfig:  new(false),
					SupportedReasoningTypes:  []spec.ReasoningType{},
					SupportedReasoningLevels: []spec.ReasoningLevel{},
				},
				StopSequenceCapabilities: &StopSequenceCapabilitiesOverride{
					IsSupported:             new(false),
					DisallowedWithReasoning: new(false),
					MaxSequences:            new(0),
				},
				OutputCapabilities: &OutputCapabilitiesOverride{
					SupportedOutputFormats: []spec.OutputFormatKind{},
					SupportsVerbosity:      new(false),
				},
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolTypes:               []spec.ToolType{},
					SupportedToolPolicyModes:         []spec.ToolPolicyMode{},
					SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{},
					SupportsParallelToolCalls:        new(false),
					MaxForcedTools:                   new(0),
				},
				CacheCapabilities: &CacheCapabilitiesOverride{
					SupportsAutomaticCaching: new(false),
					TopLevel: &CacheControlCapabilitiesOverride{
						SupportsTTL:    new(false),
						SupportedKinds: []spec.CacheControlKind{},
						SupportedTTLs:  []spec.CacheControlTTL{},
						SupportsKey:    new(false),
					},
				},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				reasoning := valueAsMap(t, m["reasoningCapabilities"])
				assertJSONBool(t, reasoning, "supportsReasoningConfig", false)
				assertJSONArrayLen(t, reasoning, "supportedReasoningTypes", 0)
				assertJSONArrayLen(t, reasoning, "supportedReasoningLevels", 0)

				stop := valueAsMap(t, m["stopSequenceCapabilities"])
				assertJSONBool(t, stop, "isSupported", false)
				assertJSONBool(t, stop, "disallowedWithReasoning", false)
				assertJSONNumber(t, stop, "maxSequences", 0)

				output := valueAsMap(t, m["outputCapabilities"])
				assertJSONArrayLen(t, output, "supportedOutputFormats", 0)
				assertJSONBool(t, output, "supportsVerbosity", false)

				tool := valueAsMap(t, m["toolCapabilities"])
				assertJSONArrayLen(t, tool, "supportedToolTypes", 0)
				assertJSONArrayLen(t, tool, "supportedToolPolicyModes", 0)
				assertJSONArrayLen(t, tool, "supportedClientToolOutputFormats", 0)
				assertJSONBool(t, tool, "supportsParallelToolCalls", false)
				assertJSONNumber(t, tool, "maxForcedTools", 0)

				cache := valueAsMap(t, m["cacheCapabilities"])
				assertJSONBool(t, cache, "supportsAutomaticCaching", false)

				topLevel := valueAsMap(t, cache["topLevel"])
				assertJSONBool(t, topLevel, "supportsTTL", false)
				assertJSONArrayLen(t, topLevel, "supportedKinds", 0)
				assertJSONArrayLen(t, topLevel, "supportedTTLs", 0)
				assertJSONBool(t, topLevel, "supportsKey", false)
			},
		},
		{
			name: "direct reasoning override emits empty slices but omits nil sibling",
			value: ReasoningCapabilitiesOverride{
				SupportedReasoningTypes: []spec.ReasoningType{},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				assertJSONArrayLen(t, m, "supportedReasoningTypes", 0)
				assertJSONKeyAbsent(t, m, "supportedReasoningLevels")
				assertJSONKeyAbsent(t, m, "supportsReasoningConfig")
			},
		},
		{
			name: "direct output override emits empty slice",
			value: OutputCapabilitiesOverride{
				SupportedOutputFormats: []spec.OutputFormatKind{},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				assertJSONArrayLen(t, m, "supportedOutputFormats", 0)
				assertJSONKeyAbsent(t, m, "supportsVerbosity")
			},
		},
		{
			name: "direct tool override emits all empty slices",
			value: ToolCapabilitiesOverride{
				SupportedToolTypes:               []spec.ToolType{},
				SupportedToolPolicyModes:         []spec.ToolPolicyMode{},
				SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				assertJSONArrayLen(t, m, "supportedToolTypes", 0)
				assertJSONArrayLen(t, m, "supportedToolPolicyModes", 0)
				assertJSONArrayLen(t, m, "supportedClientToolOutputFormats", 0)
			},
		},
		{
			name: "direct cache control override emits empty slices",
			value: CacheControlCapabilitiesOverride{
				SupportedKinds: []spec.CacheControlKind{},
				SupportedTTLs:  []spec.CacheControlTTL{},
			},
			check: func(t *testing.T, m map[string]any) {
				t.Helper()

				assertJSONArrayLen(t, m, "supportedKinds", 0)
				assertJSONArrayLen(t, m, "supportedTTLs", 0)
				assertJSONKeyAbsent(t, m, "supportsTTL")
				assertJSONKeyAbsent(t, m, "supportsKey")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := json.Marshal(tc.value)
			assertNoErr(t, err)

			tc.check(t, mustJSONMap(t, raw))
		})
	}
}

func TestUnmarshalJSONPreservesNilVsEmptySlices(t *testing.T) {
	raw := []byte(`{
		"modalitiesIn": [],
		"reasoningCapabilities": {
			"supportsReasoningConfig": false,
			"supportedReasoningTypes": []
		},
		"outputCapabilities": {
			"supportedOutputFormats": []
		},
		"toolCapabilities": {
			"supportedToolTypes": [],
			"supportedClientToolOutputFormats": []
		},
		"cacheCapabilities": {
			"topLevel": {
				"supportedKinds": [],
				"supportedTTLs": []
			}
		}
	}`)

	var got ModelCapabilitiesOverride
	err := json.Unmarshal(raw, &got)
	assertNoErr(t, err)

	if got.ModalitiesIn == nil || len(got.ModalitiesIn) != 0 {
		t.Fatalf("expected non-nil empty modalitiesIn, got %#v", got.ModalitiesIn)
	}
	if got.ModalitiesOut != nil {
		t.Fatalf("expected nil modalitiesOut, got %#v", got.ModalitiesOut)
	}

	if got.ReasoningCapabilities == nil {
		t.Fatal("expected reasoning capabilities")
	}
	if got.ReasoningCapabilities.SupportedReasoningTypes == nil ||
		len(got.ReasoningCapabilities.SupportedReasoningTypes) != 0 {
		t.Fatalf("expected non-nil empty supportedReasoningTypes, got %#v",
			got.ReasoningCapabilities.SupportedReasoningTypes)
	}
	if got.ReasoningCapabilities.SupportedReasoningLevels != nil {
		t.Fatalf("expected nil supportedReasoningLevels, got %#v",
			got.ReasoningCapabilities.SupportedReasoningLevels)
	}
	if got.ReasoningCapabilities.SupportsReasoningConfig == nil ||
		*got.ReasoningCapabilities.SupportsReasoningConfig {
		t.Fatalf("expected supportsReasoningConfig=false pointer, got %#v",
			got.ReasoningCapabilities.SupportsReasoningConfig)
	}

	if got.OutputCapabilities == nil {
		t.Fatal("expected output capabilities")
	}
	if got.OutputCapabilities.SupportedOutputFormats == nil ||
		len(got.OutputCapabilities.SupportedOutputFormats) != 0 {
		t.Fatalf("expected non-nil empty supportedOutputFormats, got %#v",
			got.OutputCapabilities.SupportedOutputFormats)
	}

	if got.ToolCapabilities == nil {
		t.Fatal("expected tool capabilities")
	}
	if got.ToolCapabilities.SupportedToolTypes == nil ||
		len(got.ToolCapabilities.SupportedToolTypes) != 0 {
		t.Fatalf("expected non-nil empty supportedToolTypes, got %#v",
			got.ToolCapabilities.SupportedToolTypes)
	}
	if got.ToolCapabilities.SupportedToolPolicyModes != nil {
		t.Fatalf("expected nil supportedToolPolicyModes, got %#v",
			got.ToolCapabilities.SupportedToolPolicyModes)
	}
	if got.ToolCapabilities.SupportedClientToolOutputFormats == nil ||
		len(got.ToolCapabilities.SupportedClientToolOutputFormats) != 0 {
		t.Fatalf("expected non-nil empty supportedClientToolOutputFormats, got %#v",
			got.ToolCapabilities.SupportedClientToolOutputFormats)
	}

	if got.CacheCapabilities == nil || got.CacheCapabilities.TopLevel == nil {
		t.Fatalf("expected cache top-level capabilities, got %#v", got.CacheCapabilities)
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
	if got.CacheCapabilities.InputOutputContent != nil {
		t.Fatalf("expected nil inputOutputContent, got %#v", got.CacheCapabilities.InputOutputContent)
	}
}

func TestMarshalUnmarshalRoundTripPreservesEmptySlices(t *testing.T) {
	input := ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{},
		ReasoningCapabilities: &ReasoningCapabilitiesOverride{
			SupportedReasoningTypes: []spec.ReasoningType{},
		},
		OutputCapabilities: &OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{},
		},
		ToolCapabilities: &ToolCapabilitiesOverride{
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{},
		},
		CacheCapabilities: &CacheCapabilitiesOverride{
			TopLevel: &CacheControlCapabilitiesOverride{
				SupportedTTLs: []spec.CacheControlTTL{},
			},
		},
	}

	raw, err := json.Marshal(input)
	assertNoErr(t, err)

	var got ModelCapabilitiesOverride
	err = json.Unmarshal(raw, &got)
	assertNoErr(t, err)

	if got.ModalitiesIn == nil || len(got.ModalitiesIn) != 0 {
		t.Fatalf("expected non-nil empty modalitiesIn after roundtrip, got %#v", got.ModalitiesIn)
	}
	if got.ReasoningCapabilities.SupportedReasoningTypes == nil ||
		len(got.ReasoningCapabilities.SupportedReasoningTypes) != 0 {
		t.Fatalf("expected non-nil empty supportedReasoningTypes after roundtrip, got %#v",
			got.ReasoningCapabilities.SupportedReasoningTypes)
	}
	if got.OutputCapabilities.SupportedOutputFormats == nil ||
		len(got.OutputCapabilities.SupportedOutputFormats) != 0 {
		t.Fatalf("expected non-nil empty supportedOutputFormats after roundtrip, got %#v",
			got.OutputCapabilities.SupportedOutputFormats)
	}
	if got.ToolCapabilities.SupportedClientToolOutputFormats == nil ||
		len(got.ToolCapabilities.SupportedClientToolOutputFormats) != 0 {
		t.Fatalf("expected non-nil empty supportedClientToolOutputFormats after roundtrip, got %#v",
			got.ToolCapabilities.SupportedClientToolOutputFormats)
	}
	if got.CacheCapabilities.TopLevel.SupportedTTLs == nil ||
		len(got.CacheCapabilities.TopLevel.SupportedTTLs) != 0 {
		t.Fatalf("expected non-nil empty supportedTTLs after roundtrip, got %#v",
			got.CacheCapabilities.TopLevel.SupportedTTLs)
	}
}
