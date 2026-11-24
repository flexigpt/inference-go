package googleinteractionssdk

import (
	"testing"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/spec"
)

func TestBuildGoogleInteractionsToolsPolicies(t *testing.T) {
	toolChoices := []spec.ToolChoice{
		{
			Type:      spec.ToolTypeFunction,
			ID:        "echo-tool",
			Name:      "echo_text",
			Arguments: map[string]any{"type": "object"},
		},
		{
			Type: spec.ToolTypeWebSearch,
			ID:   "search-tool",
			Name: spec.DefaultWebSearchToolName,
		},
	}

	tests := []struct {
		name            string
		choices         []spec.ToolChoice
		policy          *spec.ToolPolicy
		wantTools       int
		wantNameMap     int
		wantWebSearchID string
		wantConfigured  bool
		wantWarningCode string
		wantError       bool
	}{
		{
			name:            "default policy enables callable tools and Google Search",
			choices:         toolChoices,
			wantTools:       2,
			wantNameMap:     1,
			wantWebSearchID: "search-tool",
		},
		{
			name:            "auto policy enables callable tools and Google Search",
			choices:         toolChoices,
			policy:          &spec.ToolPolicy{Mode: spec.ToolPolicyModeAuto},
			wantTools:       2,
			wantNameMap:     1,
			wantWebSearchID: "search-tool",
		},
		{
			name:           "any policy forces callable tools and omits Google Search",
			choices:        toolChoices,
			policy:         &spec.ToolPolicy{Mode: spec.ToolPolicyModeAny},
			wantTools:      1,
			wantNameMap:    1,
			wantConfigured: true,
		},
		{
			name:    "tool policy forces exactly one callable tool",
			choices: toolChoices,
			policy: &spec.ToolPolicy{
				Mode: spec.ToolPolicyModeTool,
				AllowedTools: []spec.AllowedTool{
					{ToolChoiceID: "echo-tool"},
				},
			},
			wantTools:      1,
			wantNameMap:    1,
			wantConfigured: true,
		},
		{
			name:    "none policy omits every tool",
			choices: toolChoices,
			policy:  &spec.ToolPolicy{Mode: spec.ToolPolicyModeNone},
		},
		{
			name: "any policy rejects Google Search only",
			choices: []spec.ToolChoice{{
				Type: spec.ToolTypeWebSearch,
				ID:   "search-tool",
				Name: spec.DefaultWebSearchToolName,
			}},
			policy:    &spec.ToolPolicy{Mode: spec.ToolPolicyModeAny},
			wantError: true,
		},
		{
			name:    "disable parallel returns an explicit warning",
			choices: toolChoices,
			policy: &spec.ToolPolicy{
				Mode:            spec.ToolPolicyModeAuto,
				DisableParallel: true,
			},
			wantTools:       2,
			wantNameMap:     1,
			wantWebSearchID: "search-tool",
			wantWarningCode: "toolPolicy_disableParallel_unsupported",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &interactions.GenerationConfig{}
			gotTools, gotNameMap, gotWebSearchID, configured, warnings, err := buildGoogleInteractionsTools(
				tc.choices,
				tc.policy,
				config,
			)

			if tc.wantError {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("buildGoogleInteractionsTools: %v", err)
			}
			if got := len(gotTools); got != tc.wantTools {
				t.Fatalf("tool count got %d want %d", got, tc.wantTools)
			}
			if got := len(gotNameMap); got != tc.wantNameMap {
				t.Fatalf("name-map count got %d want %d", got, tc.wantNameMap)
			}
			if gotWebSearchID != tc.wantWebSearchID {
				t.Fatalf(
					"webSearchChoiceID got %q want %q",
					gotWebSearchID,
					tc.wantWebSearchID,
				)
			}
			if configured != tc.wantConfigured {
				t.Fatalf("configured got %v want %v", configured, tc.wantConfigured)
			}
			if tc.wantWarningCode != "" &&
				!googleInteractionsHasWarning(warnings, tc.wantWarningCode) {
				t.Fatalf("missing warning %q in %#v", tc.wantWarningCode, warnings)
			}
		})
	}
}
