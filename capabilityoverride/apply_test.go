package capabilityoverride

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestApplyModelCapabilitiesOverrideTable(t *testing.T) {
	tests := []struct {
		name     string
		base     spec.ModelCapabilities
		override *ModelCapabilitiesOverride
		check    func(t *testing.T, got spec.ModelCapabilities)
	}{
		{
			name:     "nil override leaves base unchanged",
			base:     baseCapabilitiesForTest(),
			override: nil,
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()
				assertDeepEqual(t, got, baseCapabilitiesForTest())
			},
		},
		{
			name: "top-level slices and partial max-output param dialect override",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				ModalitiesIn: []spec.Modality{},
				ModalitiesOut: []spec.Modality{
					spec.ModalityFileOut,
				},
				ParamDialect: &ParamDialectOverride{
					MaxOutputTokensParamName: new(
						spec.MaxOutputTokensParamNameMaxTokens,
					),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				if got.ModalitiesIn == nil {
					t.Fatal("expected modalitiesIn to be a non-nil empty slice")
				}
				assertDeepEqual(t, got.ModalitiesIn, []spec.Modality{})
				assertDeepEqual(t, got.ModalitiesOut, []spec.Modality{
					spec.ModalityFileOut,
				})

				if got.ParamDialect == nil {
					t.Fatal("expected ParamDialect")
				}
				if got.ParamDialect.MaxOutputTokensParamName != spec.MaxOutputTokensParamNameMaxTokens {
					t.Fatalf("unexpected max output token param: %q", got.ParamDialect.MaxOutputTokensParamName)
				}
				if got.ParamDialect.ToolChoiceParamStyle != spec.ToolChoiceParamStyleAllowedTools {
					t.Fatalf(
						"expected existing tool choice style to remain, got %q",
						got.ParamDialect.ToolChoiceParamStyle,
					)
				}
			},
		},
		{
			name: "partial tool-choice param dialect override creates dialect",
			base: spec.ModelCapabilities{},
			override: &ModelCapabilitiesOverride{
				ParamDialect: &ParamDialectOverride{
					ToolChoiceParamStyle: new(
						spec.ToolChoiceParamStyleRequiredNamed,
					),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				if got.ParamDialect == nil {
					t.Fatal("expected ParamDialect")
				}
				if got.ParamDialect.MaxOutputTokensParamName != "" {
					t.Fatalf("expected zero max output token param, got %q", got.ParamDialect.MaxOutputTokensParamName)
				}
				if got.ParamDialect.ToolChoiceParamStyle != spec.ToolChoiceParamStyleRequiredNamed {
					t.Fatalf("unexpected tool choice style: %q", got.ParamDialect.ToolChoiceParamStyle)
				}
			},
		},
		{
			name: "reasoning override updates all supported fields and keeps unsupported base fields",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportsReasoningConfig: new(false),
					SupportedReasoningTypes: []spec.ReasoningType{},
					SupportedReasoningLevels: []spec.ReasoningLevel{
						spec.ReasoningLevelMinimal,
						spec.ReasoningLevelMax,
					},
					HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
						MinAllowed:      new(2),
						MaxAllowed:      new(8192),
						ZeroAllowed:     new(false),
						MinusOneAllowed: new(false),
					},
					SupportsSummaryStyle:             new(true),
					SupportsEncryptedReasoningInput:  new(false),
					TemperatureDisallowedWhenEnabled: new(true),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				rc := got.ReasoningCapabilities
				if rc == nil {
					t.Fatal("expected reasoning capabilities")
				}
				if rc.SupportsReasoningConfig {
					t.Fatal("expected supportsReasoningConfig=false")
				}
				if rc.SupportedReasoningTypes == nil {
					t.Fatal("expected supportedReasoningTypes to be non-nil empty slice")
				}
				assertDeepEqual(t, rc.SupportedReasoningTypes, []spec.ReasoningType{})
				assertDeepEqual(t, rc.SupportedReasoningLevels, []spec.ReasoningLevel{
					spec.ReasoningLevelMinimal,
					spec.ReasoningLevelMax,
				})
				if !rc.SupportsSummaryStyle {
					t.Fatal("expected supportsSummaryStyle=true")
				}
				if rc.SupportsEncryptedReasoningInput {
					t.Fatal("expected supportsEncryptedReasoningInput=false")
				}
				if !rc.TemperatureDisallowedWhenEnabled {
					t.Fatal("expected temperatureDisallowedWhenEnabled=true")
				}
				if rc.HybridTokenBudgetCapabilities == nil {
					t.Fatal("expected hybrid token budget capabilities")
				}
				if rc.HybridTokenBudgetCapabilities.MinAllowed != 2 {
					t.Fatalf("unexpected minAllowed: %d", rc.HybridTokenBudgetCapabilities.MinAllowed)
				}
				if rc.HybridTokenBudgetCapabilities.MaxAllowed != 8192 {
					t.Fatalf("unexpected maxAllowed: %d", rc.HybridTokenBudgetCapabilities.MaxAllowed)
				}
				if rc.HybridTokenBudgetCapabilities.ZeroAllowed {
					t.Fatal("expected zeroAllowed=false")
				}
				if rc.HybridTokenBudgetCapabilities.MinusOneAllowed {
					t.Fatal("expected minusOneAllowed=false")
				}
			},
		},
		{
			name: "reasoning override keeps base token budget when omitted",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportsSummaryStyle: new(true),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				rc := got.ReasoningCapabilities
				if rc.HybridTokenBudgetCapabilities == nil {
					t.Fatal("expected existing hybrid token budget capabilities to remain")
				}
				if rc.HybridTokenBudgetCapabilities.MinAllowed != 1 {
					t.Fatalf("unexpected hybrid token budget capabilities: %#v", rc.HybridTokenBudgetCapabilities)
				}
			},
		},
		{
			name: "stop sequence override updates false true and zero scalar values",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				StopSequenceCapabilities: &StopSequenceCapabilitiesOverride{
					IsSupported:             new(false),
					DisallowedWithReasoning: new(true),
					MaxSequences:            new(0),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				sc := got.StopSequenceCapabilities
				if sc == nil {
					t.Fatal("expected stop sequence capabilities")
				}
				if sc.IsSupported {
					t.Fatal("expected isSupported=false")
				}
				if !sc.DisallowedWithReasoning {
					t.Fatal("expected disallowedWithReasoning=true")
				}
				if sc.MaxSequences != 0 {
					t.Fatalf("expected maxSequences=0, got %d", sc.MaxSequences)
				}
			},
		},
		{
			name: "output override preserves empty supported formats and false-to-true scalar",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				OutputCapabilities: &OutputCapabilitiesOverride{
					SupportedOutputFormats: []spec.OutputFormatKind{},
					SupportsVerbosity:      new(true),
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				oc := got.OutputCapabilities
				if oc == nil {
					t.Fatal("expected output capabilities")
				}
				if oc.SupportedOutputFormats == nil {
					t.Fatal("expected supportedOutputFormats to be non-nil empty slice")
				}
				assertDeepEqual(t, oc.SupportedOutputFormats, []spec.OutputFormatKind{})
				if !oc.SupportsVerbosity {
					t.Fatal("expected supportsVerbosity=true")
				}
			},
		},
		{
			name: "tool override updates all supported fields including empty slices",
			base: baseCapabilitiesForTest(),
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolTypes: []spec.ToolType{},
					SupportedToolPolicyModes: []spec.ToolPolicyMode{
						spec.ToolPolicyModeNone,
						spec.ToolPolicyModeTool,
					},
					SupportsParallelToolCalls:        new(false),
					MaxForcedTools:                   new(0),
					SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{},
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				tc := got.ToolCapabilities
				if tc == nil {
					t.Fatal("expected tool capabilities")
				}
				if tc.SupportedToolTypes == nil {
					t.Fatal("expected supportedToolTypes to be non-nil empty slice")
				}
				assertDeepEqual(t, tc.SupportedToolTypes, []spec.ToolType{})
				assertDeepEqual(t, tc.SupportedToolPolicyModes, []spec.ToolPolicyMode{
					spec.ToolPolicyModeNone,
					spec.ToolPolicyModeTool,
				})
				if tc.SupportsParallelToolCalls {
					t.Fatal("expected supportsParallelToolCalls=false")
				}
				if tc.MaxForcedTools != 0 {
					t.Fatalf("expected maxForcedTools=0, got %d", tc.MaxForcedTools)
				}
				if tc.SupportedClientToolOutputFormats == nil {
					t.Fatal("expected supportedClientToolOutputFormats to be non-nil empty slice")
				}
				assertDeepEqual(t, tc.SupportedClientToolOutputFormats, []spec.ToolOutputFormatKind{})
			},
		},
		{
			name: "cache override creates all scopes and preserves empty slices",
			base: spec.ModelCapabilities{
				CacheCapabilities: &spec.CacheCapabilities{
					SupportsAutomaticCaching: true,
				},
			},
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					SupportsAutomaticCaching: new(false),
					TopLevel: &CacheControlCapabilitiesOverride{
						SupportsTTL:    new(false),
						SupportedKinds: []spec.CacheControlKind{},
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTL1h,
						},
						SupportsKey: new(true),
					},
					InputOutputContent: &CacheControlCapabilitiesOverride{
						SupportsTTL: new(true),
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKindEphemeral,
						},
						SupportedTTLs: []spec.CacheControlTTL{},
						SupportsKey:   new(false),
					},
					ReasoningContent: &CacheControlCapabilitiesOverride{
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKindEphemeral,
						},
					},
					ToolChoice: &CacheControlCapabilitiesOverride{
						SupportsKey: new(true),
					},
					ToolCall: &CacheControlCapabilitiesOverride{
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTLInMemory,
						},
					},
					ToolOutput: &CacheControlCapabilitiesOverride{
						SupportsTTL: new(true),
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKindEphemeral,
						},
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTL5m,
						},
						SupportsKey: new(true),
					},
				},
			},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				cc := got.CacheCapabilities
				if cc == nil {
					t.Fatal("expected cache capabilities")
				}
				if cc.SupportsAutomaticCaching {
					t.Fatal("expected supportsAutomaticCaching=false")
				}

				assertCacheControlCapabilities(t, cc.TopLevel, false,
					[]spec.CacheControlKind{},
					[]spec.CacheControlTTL{spec.CacheControlTTL1h},
					true,
				)
				assertCacheControlCapabilities(t, cc.InputOutputContent, true,
					[]spec.CacheControlKind{spec.CacheControlKindEphemeral},
					[]spec.CacheControlTTL{},
					false,
				)
				assertCacheControlCapabilities(t, cc.ReasoningContent, false,
					[]spec.CacheControlKind{spec.CacheControlKindEphemeral},
					nil,
					false,
				)
				assertCacheControlCapabilities(t, cc.ToolChoice, false, nil, nil, true)
				assertCacheControlCapabilities(t, cc.ToolCall, false, nil,
					[]spec.CacheControlTTL{spec.CacheControlTTLInMemory},
					false,
				)
				assertCacheControlCapabilities(t, cc.ToolOutput, true,
					[]spec.CacheControlKind{spec.CacheControlKindEphemeral},
					[]spec.CacheControlTTL{spec.CacheControlTTL5m},
					true,
				)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.base
			ApplyModelCapabilitiesOverrides(&got, tc.override)
			tc.check(t, got)
		})
	}
}

func TestApplyModelCapabilitiesOverridesAndDeriveTable(t *testing.T) {
	providerOverride := &ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityImageIn,
		},
		OutputCapabilities: &OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{
				spec.OutputFormatKindJSONSchema,
			},
			SupportsVerbosity: new(true),
		},
	}

	modelOverride := &ModelCapabilitiesOverride{
		ModalitiesIn: []spec.Modality{
			spec.ModalityFileIn,
		},
		OutputCapabilities: &OutputCapabilitiesOverride{
			SupportedOutputFormats: []spec.OutputFormatKind{},
		},
	}

	tests := []struct {
		name      string
		overrides []*ModelCapabilitiesOverride
		check     func(t *testing.T, got spec.ModelCapabilities)
	}{
		{
			name:      "no overrides returns cloned base",
			overrides: nil,
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()
				assertDeepEqual(t, got, baseCapabilitiesForTest())
			},
		},
		{
			name:      "nil override in list is ignored",
			overrides: []*ModelCapabilitiesOverride{nil},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()
				assertDeepEqual(t, got, baseCapabilitiesForTest())
			},
		},
		{
			name:      "later model override wins over provider override while omitted fields remain",
			overrides: []*ModelCapabilitiesOverride{providerOverride, modelOverride},
			check: func(t *testing.T, got spec.ModelCapabilities) {
				t.Helper()

				assertDeepEqual(t, got.ModalitiesIn, []spec.Modality{
					spec.ModalityFileIn,
				})

				if got.OutputCapabilities == nil {
					t.Fatal("expected output capabilities")
				}
				if got.OutputCapabilities.SupportedOutputFormats == nil {
					t.Fatal("expected model override to set non-nil empty output format slice")
				}
				assertDeepEqual(t, got.OutputCapabilities.SupportedOutputFormats, []spec.OutputFormatKind{})
				if !got.OutputCapabilities.SupportsVerbosity {
					t.Fatal("expected provider supportsVerbosity override to remain true")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base := baseCapabilitiesForTest()
			before := CloneModelCapabilities(base)

			derived := DeriveModelCapabilities(base, tc.overrides...)
			tc.check(t, derived)

			assertDeepEqual(t, base, before)

			applied := CloneModelCapabilities(base)
			ApplyModelCapabilitiesOverrides(&applied, tc.overrides...)
			assertDeepEqual(t, applied, derived)

			assertDeepEqual(t, base, before)
		})
	}
}

func TestApplyNilDestinationDoesNotPanicTable(t *testing.T) {
	tests := []struct {
		name string
		run  func()
	}{
		{
			name: "single override nil destination",
			run: func() {
				ApplyModelCapabilitiesOverrides(nil, fullCapabilitiesOverrideForTest())
			},
		},
		{
			name: "multiple overrides nil destination",
			run: func() {
				ApplyModelCapabilitiesOverrides(nil, fullCapabilitiesOverrideForTest())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assertDoesNotPanic(t, tc.run)
		})
	}
}

func assertDoesNotPanic(t *testing.T, run func()) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("expected no panic, got %v", r)
		}
	}()

	run()
}

func assertCacheControlCapabilities(
	t *testing.T,
	got *spec.CacheControlCapabilities,
	wantSupportsTTL bool,
	wantKinds []spec.CacheControlKind,
	wantTTLs []spec.CacheControlTTL,
	wantSupportsKey bool,
) {
	t.Helper()

	if got == nil {
		t.Fatal("expected cache control capabilities")
	}

	if got.SupportsTTL != wantSupportsTTL {
		t.Fatalf("unexpected supportsTTL: got %v want %v", got.SupportsTTL, wantSupportsTTL)
	}
	assertDeepEqual(t, got.SupportedKinds, wantKinds)
	assertDeepEqual(t, got.SupportedTTLs, wantTTLs)

	if got.SupportsKey != wantSupportsKey {
		t.Fatalf("unexpected supportsKey: got %v want %v", got.SupportsKey, wantSupportsKey)
	}
}

func TestEmptyCacheControlScopeOverrideIsNoOp(t *testing.T) {
	base := spec.ModelCapabilities{
		CacheCapabilities: &spec.CacheCapabilities{
			SupportsAutomaticCaching: true,
		},
	}

	got := CloneModelCapabilities(base)
	ApplyModelCapabilitiesOverrides(&got, &ModelCapabilitiesOverride{
		CacheCapabilities: &CacheCapabilitiesOverride{
			SupportsAutomaticCaching: new(false),
			TopLevel:                 &CacheControlCapabilitiesOverride{},
		},
	})

	if got.CacheCapabilities == nil {
		t.Fatal("expected cache capabilities")
	}
	if got.CacheCapabilities.SupportsAutomaticCaching {
		t.Fatal("expected supportsAutomaticCaching=false")
	}
	if got.CacheCapabilities.TopLevel != nil {
		t.Fatalf("expected empty topLevel override to be a no-op, got %#v", got.CacheCapabilities.TopLevel)
	}
}
