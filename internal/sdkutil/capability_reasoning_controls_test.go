package sdkutil

import (
	"slices"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestNormalizeRequestForSDKReasoningContextAndMode(t *testing.T) {
	tests := []struct {
		name            string
		sdkType         spec.ProviderSDKType
		supportsContext bool
		supportsMode    bool
		wantContext     bool
		wantMode        bool
		wantWarnings    []string
	}{
		{
			name:            "OpenAI Responses preserves supported context and mode",
			sdkType:         spec.ProviderSDKTypeOpenAIResponses,
			supportsContext: true,
			supportsMode:    true,
			wantContext:     true,
			wantMode:        true,
		},
		{
			name:            "Anthropic drops unsupported context and mode",
			sdkType:         spec.ProviderSDKTypeAnthropic,
			supportsContext: false,
			supportsMode:    false,
			wantContext:     false,
			wantMode:        false,
			wantWarnings: []string{
				"reasoning_context_dropped_unsupported",
				"reasoning_mode_dropped_unsupported",
			},
		},
		{
			name:            "Google drops unsupported context and mode",
			sdkType:         spec.ProviderSDKTypeGoogleGenerateContent,
			supportsContext: false,
			supportsMode:    false,
			wantContext:     false,
			wantMode:        false,
			wantWarnings: []string{
				"reasoning_context_dropped_unsupported",
				"reasoning_mode_dropped_unsupported",
			},
		},
		{
			name:            "context remains when only mode is unsupported",
			sdkType:         spec.ProviderSDKTypeOpenAIResponses,
			supportsContext: true,
			supportsMode:    false,
			wantContext:     true,
			wantMode:        false,
			wantWarnings: []string{
				"reasoning_mode_dropped_unsupported",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			contextValue := spec.ReasoningContextAllTurns
			modeValue := spec.ReasoningModePro

			request := &spec.FetchCompletionRequest{
				ModelParam: spec.ModelParam{
					Name: "test-model",
					Reasoning: &spec.ReasoningParam{
						Type:    spec.ReasoningTypeSingleWithLevels,
						Level:   spec.ReasoningLevelLow,
						Context: &contextValue,
						Mode:    &modeValue,
					},
				},
				Inputs: []spec.InputUnion{{
					Kind: spec.InputKindInputMessage,
					InputMessage: &spec.InputOutputContent{
						Role: spec.RoleUser,
						Contents: []spec.InputOutputContentItemUnion{{
							Kind:     spec.ContentItemKindText,
							TextItem: &spec.ContentItemText{Text: "test"},
						}},
					},
				}},
			}

			caps := reasoningControlsTestCapabilities(
				tc.supportsContext,
				tc.supportsMode,
			)
			normalized, _, warnings, err := NormalizeRequestForSDK(
				t.Context(),
				request,
				nil,
				tc.sdkType,
				caps,
			)
			if err != nil {
				t.Fatalf("NormalizeRequestForSDK: %v", err)
			}
			if normalized.ModelParam.Reasoning == nil {
				t.Fatal("reasoning unexpectedly dropped")
			}

			gotReasoning := normalized.ModelParam.Reasoning
			if (gotReasoning.Context != nil) != tc.wantContext {
				t.Fatalf(
					"context present=%v want=%v",
					gotReasoning.Context != nil,
					tc.wantContext,
				)
			}
			if (gotReasoning.Mode != nil) != tc.wantMode {
				t.Fatalf(
					"mode present=%v want=%v",
					gotReasoning.Mode != nil,
					tc.wantMode,
				)
			}

			gotWarningCodes := make([]string, 0, len(warnings))
			for _, warning := range warnings {
				gotWarningCodes = append(gotWarningCodes, warning.Code)
			}
			if !slices.Equal(gotWarningCodes, tc.wantWarnings) {
				t.Fatalf(
					"warning codes got %#v want %#v",
					gotWarningCodes,
					tc.wantWarnings,
				)
			}

			if request.ModelParam.Reasoning.Context == nil ||
				request.ModelParam.Reasoning.Mode == nil {
				t.Fatal("normalization mutated the original request")
			}
		})
	}
}

func reasoningControlsTestCapabilities(
	supportsContext bool,
	supportsMode bool,
) spec.ModelCapabilities {
	return spec.ModelCapabilities{
		ModalitiesIn: []spec.Modality{spec.ModalityTextIn},
		ReasoningCapabilities: &spec.ReasoningCapabilities{
			SupportsReasoningConfig: true,
			SupportedReasoningTypes: []spec.ReasoningType{
				spec.ReasoningTypeSingleWithLevels,
			},
			SupportedReasoningLevels: []spec.ReasoningLevel{
				spec.ReasoningLevelLow,
			},
			SupportsSummaryStyle:     true,
			SupportsReasoningContext: supportsContext,
			SupportsReasoningMode:    supportsMode,
		},
	}
}
