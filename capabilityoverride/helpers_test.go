package capabilityoverride

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func assertErrContains(t *testing.T, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", want)
	}

	if !strings.Contains(err.Error(), want) {
		t.Fatalf("expected error containing %q, got %q", want, err.Error())
	}
}

func assertDeepEqual(t *testing.T, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected value\n got: %#v\nwant: %#v", got, want)
	}
}

func mustJSONMap(t *testing.T, raw []byte) map[string]any {
	t.Helper()

	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatalf("failed to unmarshal JSON object %s: %v", string(raw), err)
	}

	return out
}

func assertJSONKeyAbsent(t *testing.T, m map[string]any, key string) {
	t.Helper()

	if _, ok := m[key]; ok {
		t.Fatalf("expected JSON key %q to be absent in %#v", key, m)
	}
}

func assertJSONArrayLen(t *testing.T, m map[string]any, key string, want int) {
	t.Helper()

	v, ok := m[key]
	if !ok {
		t.Fatalf("expected JSON key %q to be present in %#v", key, m)
	}

	got := valueAsSlice(t, v)
	if len(got) != want {
		t.Fatalf("expected JSON array %q length %d, got %d: %#v", key, want, len(got), got)
	}
}

func assertJSONBool(t *testing.T, m map[string]any, key string, want bool) {
	t.Helper()

	v, ok := m[key]
	if !ok {
		t.Fatalf("expected JSON key %q to be present in %#v", key, m)
	}

	got, ok := v.(bool)
	if !ok {
		t.Fatalf("expected JSON bool for key %q, got %T: %#v", key, v, v)
	}

	if got != want {
		t.Fatalf("expected JSON bool %q=%v, got %v", key, want, got)
	}
}

func assertJSONNumber(t *testing.T, m map[string]any, key string, want float64) {
	t.Helper()

	v, ok := m[key]
	if !ok {
		t.Fatalf("expected JSON key %q to be present in %#v", key, m)
	}

	got, ok := v.(float64)
	if !ok {
		t.Fatalf("expected JSON number for key %q, got %T: %#v", key, v, v)
	}

	if got != want {
		t.Fatalf("expected JSON number %q=%v, got %v", key, want, got)
	}
}

func baseCapabilitiesForTest() spec.ModelCapabilities {
	return spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &spec.ReasoningCapabilities{
			SupportsReasoningConfig: true,
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
			},
			HybridTokenBudgetCapabilities: &spec.ReasoningTokenBudgetCapabilities{
				MinAllowed:      1,
				MaxAllowed:      2048,
				ZeroAllowed:     true,
				MinusOneAllowed: true,
			},
			SupportsSummaryStyle:             false,
			SupportsEncryptedReasoningInput:  true,
			TemperatureDisallowedWhenEnabled: false,
		},
		StopSequenceCapabilities: &spec.StopSequenceCapabilities{
			IsSupported:             true,
			DisallowedWithReasoning: false,
			MaxSequences:            4,
		},
		OutputCapabilities: &spec.OutputCapabilities{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
			},
			SupportsVerbosity: false,
		},
		ToolCapabilities: &spec.ToolCapabilities{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
			},
			SupportsParallelToolCalls: true,
			MaxForcedTools:            1,
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
			},
		},
		CacheCapabilities: &spec.CacheCapabilities{
			SupportsAutomaticCaching: true,
			TopLevel:                 cacheControlCapabilitiesForTest(spec.CacheControlTTL5m),
			InputOutputContent:       cacheControlCapabilitiesForTest(spec.CacheControlTTL1h),
			ReasoningContent:         cacheControlCapabilitiesForTest(spec.CacheControlTTL24h),
			ToolChoice:               cacheControlCapabilitiesForTest(spec.CacheControlTTLInMemory),
			ToolCall:                 cacheControlCapabilitiesForTest(spec.CacheControlTTL5m),
			ToolOutput:               cacheControlCapabilitiesForTest(spec.CacheControlTTL1h),
		},
		ParamDialect: &spec.ParamDialect{
			MaxOutputTokensParamName: spec.MaxOutputTokensParamNameMaxCompletionTokens,
			ToolChoiceParamStyle:     spec.ToolChoiceParamStyleAllowedTools,
		},
	}
}

func cacheControlCapabilitiesForTest(
	ttl spec.CacheControlTTL,
) *spec.CacheControlCapabilities {
	return &spec.CacheControlCapabilities{
		SupportsTTL: true,
		SupportedKinds: []spec.CacheControlKind{
			spec.CacheControlKindEphemeral,
		},
		SupportedTTLs: []spec.CacheControlTTL{
			ttl,
		},
		SupportsKey: true,
	}
}

func fullCapabilitiesOverrideForTest() *ModelCapabilitiesOverride {
	return &ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityTextIn,
			spec.ModalityImageIn,
			spec.ModalityFileIn,
		},
		ModalitiesOut: []spec.Modality{
			spec.ModalityTextOut,
		},
		ReasoningCapabilities: &ReasoningCapabilitiesOverride{
			SupportsReasoningConfig: new(true),
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeHybridWithTokens,
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelNone,
				spec.ReasoningLevelMinimal,
				spec.ReasoningLevelLow,
				spec.ReasoningLevelMedium,
				spec.ReasoningLevelHigh,
				spec.ReasoningLevelXHigh,
				spec.ReasoningLevelMax,
			},
			HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
				MinAllowed:      new(1),
				MaxAllowed:      new(4096),
				ZeroAllowed:     new(true),
				MinusOneAllowed: new(true),
			},
			SupportsSummaryStyle:             new(true),
			SupportsEncryptedReasoningInput:  new(true),
			TemperatureDisallowedWhenEnabled: new(true),
		},
		StopSequenceCapabilities: &StopSequenceCapabilitiesOverride{
			IsSupported:             new(true),
			DisallowedWithReasoning: new(true),
			MaxSequences:            new(4),
		},
		OutputCapabilities: &OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindText,
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(true),
		},
		ToolCapabilities: &ToolCapabilitiesOverride{
			SupportedToolTypes: []spec.ToolType{
				spec.ToolTypeFunction,
				spec.ToolTypeCustom,
				spec.ToolTypeWebSearch,
			},
			SupportedToolPolicyModes: []spec.ToolPolicyMode{
				spec.ToolPolicyModeAuto,
				spec.ToolPolicyModeAny,
				spec.ToolPolicyModeTool,
				spec.ToolPolicyModeNone,
			},
			SupportsParallelToolCalls: new(true),
			MaxForcedTools:            new(2),
			SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
				spec.ToolOutputFormatKindString,
				spec.ToolOutputFormatKindContentItemList,
			},
		},
		CacheCapabilities: validCacheCapabilitiesOverrideForTest(),
		ParamDialect: &ParamDialectOverride{
			MaxOutputTokensParamName: new(
				spec.MaxOutputTokensParamNameMaxTokens,
			),
			ToolChoiceParamStyle: new(
				spec.ToolChoiceParamStyleRequiredNamed,
			),
		},
	}
}

func validCacheCapabilitiesOverrideForTest() *CacheCapabilitiesOverride {
	return &CacheCapabilitiesOverride{
		SupportsAutomaticCaching: new(true),
		TopLevel:                 validCacheControlCapabilitiesOverrideForTest(),
		InputOutputContent:       validCacheControlCapabilitiesOverrideForTest(),
		ReasoningContent:         validCacheControlCapabilitiesOverrideForTest(),
		ToolChoice:               validCacheControlCapabilitiesOverrideForTest(),
		ToolCall:                 validCacheControlCapabilitiesOverrideForTest(),
		ToolOutput:               validCacheControlCapabilitiesOverrideForTest(),
	}
}

func validCacheControlCapabilitiesOverrideForTest() *CacheControlCapabilitiesOverride {
	return &CacheControlCapabilitiesOverride{
		SupportsTTL: new(true),
		SupportedKinds: []spec.CacheControlKind{
			spec.CacheControlKindEphemeral,
		},
		SupportedTTLs: []spec.CacheControlTTL{
			spec.CacheControlTTL5m,
			spec.CacheControlTTL1h,
			spec.CacheControlTTL24h,
			spec.CacheControlTTLInMemory,
		},
		SupportsKey: new(true),
	}
}

func valueAsMap(t *testing.T, v any) map[string]any {
	t.Helper()

	out, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("expected JSON object, got %T: %#v", v, v)
	}

	return out
}

func valueAsSlice(t *testing.T, v any) []any {
	t.Helper()

	out, ok := v.([]any)
	if !ok {
		t.Fatalf("expected JSON array, got %T: %#v", v, v)
	}

	return out
}
