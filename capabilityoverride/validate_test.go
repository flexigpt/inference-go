package capabilityoverride

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestValidateModelCapabilitiesOverrideTable(t *testing.T) {
	tests := []struct {
		name     string
		override *ModelCapabilitiesOverride
		wantErr  string
	}{
		{
			name:     "nil is valid",
			override: nil,
		},
		{
			name:     "empty override is valid",
			override: &ModelCapabilitiesOverride{},
		},
		{
			name:     "full valid override",
			override: fullCapabilitiesOverrideForTest(),
		},
		{
			name: "empty slice overrides are valid",
			override: &ModelCapabilitiesOverride{
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
			},
		},
		{
			name: "reasoning levels are valid without supported types in same patch",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningLevels: []spec.ReasoningLevel{
						spec.ReasoningLevelLow,
						spec.ReasoningLevelHigh,
					},
				},
			},
		},
		{
			name: "hybrid token budget is valid without supported types in same patch",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
						MinAllowed:      new(1),
						MaxAllowed:      new(4096),
						ZeroAllowed:     new(true),
						MinusOneAllowed: new(false),
					},
				},
			},
		},
		{
			name: "reasoning supported types can be narrowed independently",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningTypes: []spec.ReasoningType{spec.ReasoningTypeHybridWithTokens},
				},
			},
		},

		{
			name: "partial param dialect max only is valid",
			override: &ModelCapabilitiesOverride{
				ParamDialect: &ParamDialectOverride{
					MaxOutputTokensParamName: new(
						spec.MaxOutputTokensParamNameMaxTokens,
					),
				},
			},
		},
		{
			name: "partial param dialect tool style only is valid",
			override: &ModelCapabilitiesOverride{
				ParamDialect: &ParamDialectOverride{
					ToolChoiceParamStyle: new(
						spec.ToolChoiceParamStyleAllowedTools,
					),
				},
			},
		},
		{
			name: "modalitiesIn empty value",
			override: &ModelCapabilitiesOverride{
				ModalitiesIn: []spec.Modality{""},
			},
			wantErr: "modalitiesIn: [0] empty modality",
		},
		{
			name: "modalitiesOut unknown value",
			override: &ModelCapabilitiesOverride{
				ModalitiesOut: []spec.Modality{
					spec.Modality("bad"),
				},
			},
			wantErr: "modalitiesOut: [0] invalid output modality",
		},
		{
			name: "modalitiesIn duplicate",
			override: &ModelCapabilitiesOverride{
				ModalitiesIn: []spec.Modality{
					spec.ModalityTextIn,
					spec.ModalityTextIn,
				},
			},
			wantErr: "duplicate modality",
		},
		{
			name: "reasoning unknown type",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningTypes: []spec.ReasoningType{
						spec.ReasoningType("bad"),
					},
				},
			},
			wantErr: "supportedReasoningTypes[0] unknown type",
		},
		{
			name: "reasoning duplicate type",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningTypes: []spec.ReasoningType{
						spec.ReasoningTypeHybridWithTokens,
						spec.ReasoningTypeHybridWithTokens,
					},
				},
			},
			wantErr: "supportedReasoningTypes[1] duplicate",
		},
		{
			name: "reasoning unknown level",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningLevels: []spec.ReasoningLevel{
						spec.ReasoningLevel("bad"),
					},
				},
			},
			wantErr: "supportedReasoningLevels[0] unknown level",
		},
		{
			name: "reasoning duplicate level",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					SupportedReasoningLevels: []spec.ReasoningLevel{
						spec.ReasoningLevelHigh,
						spec.ReasoningLevelHigh,
					},
				},
			},
			wantErr: "supportedReasoningLevels[1] duplicate",
		},
		{
			name: "reasoning token budget negative min",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
						MinAllowed: new(-1),
					},
				},
			},
			wantErr: "minAllowed must be >= 0",
		},
		{
			name: "reasoning token budget negative max",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
						MaxAllowed: new(-1),
					},
				},
			},
			wantErr: "maxAllowed must be >= 0",
		},
		{
			name: "reasoning token budget max below min",
			override: &ModelCapabilitiesOverride{
				ReasoningCapabilities: &ReasoningCapabilitiesOverride{
					HybridTokenBudgetCapabilities: &ReasoningTokenBudgetCapabilitiesOverride{
						MinAllowed: new(1),
						MaxAllowed: new(0),
					},
				},
			},
			wantErr: "maxAllowed must be >= minAllowed",
		},

		{
			name: "stop sequence negative max",
			override: &ModelCapabilitiesOverride{
				StopSequenceCapabilities: &StopSequenceCapabilitiesOverride{
					MaxSequences: new(-1),
				},
			},
			wantErr: "maxSequences must be >= 0",
		},
		{
			name: "output unknown format",
			override: &ModelCapabilitiesOverride{
				OutputCapabilities: &OutputCapabilitiesOverride{
					SupportedOutputFormats: []spec.OutputFormatKind{
						spec.OutputFormatKind("bad"),
					},
				},
			},
			wantErr: "supportedOutputFormats[0] unknown kind",
		},
		{
			name: "output duplicate format",
			override: &ModelCapabilitiesOverride{
				OutputCapabilities: &OutputCapabilitiesOverride{
					SupportedOutputFormats: []spec.OutputFormatKind{
						spec.OutputFormatKindText,
						spec.OutputFormatKindText,
					},
				},
			},
			wantErr: "supportedOutputFormats[1] duplicate",
		},
		{
			name: "tool negative max forced tools",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					MaxForcedTools: new(-1),
				},
			},
			wantErr: "maxForcedTools must be >= 0",
		},
		{
			name: "tool unknown type",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolTypes: []spec.ToolType{
						spec.ToolType("bad"),
					},
				},
			},
			wantErr: "supportedToolTypes[0] unknown type",
		},
		{
			name: "tool duplicate type",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolTypes: []spec.ToolType{
						spec.ToolTypeFunction,
						spec.ToolTypeFunction,
					},
				},
			},
			wantErr: "supportedToolTypes[1] duplicate",
		},
		{
			name: "tool unknown policy mode",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolPolicyModes: []spec.ToolPolicyMode{
						spec.ToolPolicyMode("bad"),
					},
				},
			},
			wantErr: "supportedToolPolicyModes[0] unknown mode",
		},
		{
			name: "tool duplicate policy mode",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedToolPolicyModes: []spec.ToolPolicyMode{
						spec.ToolPolicyModeAuto,
						spec.ToolPolicyModeAuto,
					},
				},
			},
			wantErr: "supportedToolPolicyModes[1] duplicate",
		},
		{
			name: "tool unknown client output format",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
						spec.ToolOutputFormatKind("bad"),
					},
				},
			},
			wantErr: "supportedClientToolOutputFormats[0] unknown kind",
		},
		{
			name: "tool duplicate client output format",
			override: &ModelCapabilitiesOverride{
				ToolCapabilities: &ToolCapabilitiesOverride{
					SupportedClientToolOutputFormats: []spec.ToolOutputFormatKind{
						spec.ToolOutputFormatKindString,
						spec.ToolOutputFormatKindString,
					},
				},
			},
			wantErr: "supportedClientToolOutputFormats[1] duplicate",
		},
		{
			name: "cache topLevel unknown kind",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					TopLevel: &CacheControlCapabilitiesOverride{
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKind("bad"),
						},
					},
				},
			},
			wantErr: "topLevel: supportedKinds[0] unknown kind",
		},
		{
			name: "cache inputOutputContent duplicate kind",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					InputOutputContent: &CacheControlCapabilitiesOverride{
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKindEphemeral,
							spec.CacheControlKindEphemeral,
						},
					},
				},
			},
			wantErr: "inputOutputContent: supportedKinds[1] duplicate",
		},
		{
			name: "cache reasoningContent unknown ttl",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					ReasoningContent: &CacheControlCapabilitiesOverride{
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTL("bad"),
						},
					},
				},
			},
			wantErr: "reasoningContent: supportedTTLs[0] unknown TTL",
		},
		{
			name: "cache toolChoice duplicate ttl",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					ToolChoice: &CacheControlCapabilitiesOverride{
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTL5m,
							spec.CacheControlTTL5m,
						},
					},
				},
			},
			wantErr: "toolChoice: supportedTTLs[1] duplicate",
		},
		{
			name: "cache toolCall unknown ttl",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					ToolCall: &CacheControlCapabilitiesOverride{
						SupportedTTLs: []spec.CacheControlTTL{
							spec.CacheControlTTL("bad"),
						},
					},
				},
			},
			wantErr: "toolCall: supportedTTLs[0] unknown TTL",
		},
		{
			name: "cache toolOutput unknown kind",
			override: &ModelCapabilitiesOverride{
				CacheCapabilities: &CacheCapabilitiesOverride{
					ToolOutput: &CacheControlCapabilitiesOverride{
						SupportedKinds: []spec.CacheControlKind{
							spec.CacheControlKind("bad"),
						},
					},
				},
			},
			wantErr: "toolOutput: supportedKinds[0] unknown kind",
		},
		{
			name: "param dialect unknown max output token param",
			override: &ModelCapabilitiesOverride{
				ParamDialect: &ParamDialectOverride{
					MaxOutputTokensParamName: new(
						spec.MaxOutputTokensParamName("bad"),
					),
				},
			},
			wantErr: "maxOutputTokensParamName unknown value",
		},
		{
			name: "param dialect unknown tool choice style",
			override: &ModelCapabilitiesOverride{
				ParamDialect: &ParamDialectOverride{
					ToolChoiceParamStyle: new(
						spec.ToolChoiceParamStyle("bad"),
					),
				},
			},
			wantErr: "toolChoiceParamStyle unknown value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateModelCapabilitiesOverride(tc.override)
			if tc.wantErr == "" {
				assertNoErr(t, err)
				return
			}

			assertErrContains(t, err, tc.wantErr)
		})
	}
}
