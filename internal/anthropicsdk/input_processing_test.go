package anthropicsdk

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestToAnthropicMessagesInput_SuccessCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		systemPrompt string
		inputs       []spec.InputUnion

		wantRoles []string
		wantTypes [][]string

		check func(t *testing.T, msgs []map[string]any, sysText string)
	}{
		{
			name:         "merges turns and puts tool_result before user text",
			systemPrompt: "You are helpful.",
			inputs: []spec.InputUnion{
				userText("What's the weather in SF?"),
				assistantText("Let me check."),
				functionToolCall("", "toolu_1", "get_weather", `{"city":"SF"}`),

				// Intentionally place plain user text before the tool result in raw inputs.
				// The adapter should still emit tool_result first within the user turn.
				userText("Please summarize briefly."),
				functionToolOutput("toolu_1", "get_weather", "72F and sunny"),
			},
			wantRoles: []string{"user", "assistant", "user"},
			wantTypes: [][]string{
				{"text"},
				{"text", "tool_use"},
				{"tool_result", "text"},
			},
			check: func(t *testing.T, msgs []map[string]any, sysText string) {
				t.Helper()

				if sysText != "You are helpful." {
					t.Fatalf("unexpected system prompt: got %q want %q", sysText, "You are helpful.")
				}

				assistantBlocks := contentBlocks(t, msgs[1])
				if got := stringField(t, assistantBlocks[1], "id"); got != "toolu_1" {
					t.Fatalf("assistant tool_use id: got %q want %q", got, "toolu_1")
				}

				userBlocks := contentBlocks(t, msgs[2])
				if got := stringField(t, userBlocks[0], "tool_use_id"); got != "toolu_1" {
					t.Fatalf("user tool_result tool_use_id: got %q want %q", got, "toolu_1")
				}
			},
		},
		{
			name: "multiple client tool uses answered in one user turn",
			inputs: []spec.InputUnion{
				userText("Do both steps."),
				functionToolCall("toolu_a", "toolu_a", "tool_a", `{}`),
				functionToolCall("toolu_b", "toolu_b", "tool_b", `{}`),
				functionToolOutput("toolu_a", "tool_a", "result a"),
				functionToolOutput("toolu_b", "tool_b", "result b"),
			},
			wantRoles: []string{"user", "assistant", "user"},
			wantTypes: [][]string{
				{"text"},
				{"tool_use", "tool_use"},
				{"tool_result", "tool_result"},
			},
			check: func(t *testing.T, msgs []map[string]any, _ string) {
				t.Helper()

				userBlocks := contentBlocks(t, msgs[2])
				if got := stringField(t, userBlocks[0], "tool_use_id"); got != "toolu_a" {
					t.Fatalf("first tool_result tool_use_id: got %q want %q", got, "toolu_a")
				}
				if got := stringField(t, userBlocks[1], "tool_use_id"); got != "toolu_b" {
					t.Fatalf("second tool_result tool_use_id: got %q want %q", got, "toolu_b")
				}
			},
		},
		{
			name: "web search server tool call and result stay on assistant side",
			inputs: []spec.InputUnion{
				userText("Search for Go release notes."),
				webSearchToolCall("wsu_1", "Go release notes", map[string]any{
					"query": "Go release notes",
				}),
				webSearchToolOutput("wsu_1", "https://go.dev/doc/devel/release", "Go release notes", "enc1"),
			},
			wantRoles: []string{"user", "assistant"},
			wantTypes: [][]string{
				{"text"},
				{"server_tool_use", "web_search_tool_result"},
			},
			check: func(t *testing.T, msgs []map[string]any, _ string) {
				t.Helper()

				assistantBlocks := contentBlocks(t, msgs[1])
				if got := stringField(t, assistantBlocks[0], "id"); got != "wsu_1" {
					t.Fatalf("server_tool_use id: got %q want %q", got, "wsu_1")
				}
				if got := stringField(t, assistantBlocks[1], "tool_use_id"); got != "wsu_1" {
					t.Fatalf("web_search_tool_result tool_use_id: got %q want %q", got, "wsu_1")
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			msgs, sysPrompts, err := toAnthropicMessagesInput(t.Context(), tc.systemPrompt, tc.inputs)
			if err != nil {
				t.Fatalf("toAnthropicMessagesInput returned error: %v", err)
			}

			gotSysText := ""
			switch len(sysPrompts) {
			case 0:
			case 1:
				gotSysText = sysPrompts[0].Text
			default:
				t.Fatalf("unexpected number of system prompts: got %d want <= 1", len(sysPrompts))
			}

			gotMsgs := decodeMessages(t, msgs)

			if got := messageRoles(t, gotMsgs); !reflect.DeepEqual(got, tc.wantRoles) {
				t.Fatalf("roles mismatch:\n got  %#v\n want %#v", got, tc.wantRoles)
			}

			if got := messageContentTypes(t, gotMsgs); !reflect.DeepEqual(got, tc.wantTypes) {
				t.Fatalf("content types mismatch:\n got  %#v\n want %#v", got, tc.wantTypes)
			}

			if tc.check != nil {
				tc.check(t, gotMsgs, gotSysText)
			}
		})
	}
}

func TestToAnthropicMessagesInput_ErrorCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		inputs  []spec.InputUnion
		wantErr string
	}{
		{
			name: "rejects orphan client tool_result",
			inputs: []spec.InputUnion{
				functionToolOutput("toolu_1", "get_weather", "72F and sunny"),
			},
			wantErr: "unexpected user tool_result turn",
		},
		{
			name: "rejects client tool use followed by plain user text without tool_result",
			inputs: []spec.InputUnion{
				userText("What's the weather?"),
				functionToolCall("toolu_1", "toolu_1", "get_weather", `{"city":"SF"}`),
				userText("Also make it brief."),
			},
			wantErr: "must begin with matching tool_result blocks",
		},
		{
			name: "rejects incomplete history ending with assistant client tool use",
			inputs: []spec.InputUnion{
				userText("What's the weather?"),
				functionToolCall("toolu_1", "toolu_1", "get_weather", `{"city":"SF"}`),
			},
			wantErr: "input ends with an assistant client tool_use turn",
		},
		{
			name: "rejects missing one of multiple tool results",
			inputs: []spec.InputUnion{
				userText("Do both."),
				functionToolCall("toolu_a", "toolu_a", "tool_a", `{}`),
				functionToolCall("toolu_b", "toolu_b", "tool_b", `{}`),
				functionToolOutput("toolu_a", "tool_a", "result a"),
			},
			wantErr: "missing tool_result for toolu_b",
		},
		{
			name: "rejects tool result that does not match immediately preceding assistant tool use turn",
			inputs: []spec.InputUnion{
				userText("Do one step."),
				functionToolCall("toolu_a", "toolu_a", "tool_a", `{}`),
				functionToolOutput("toolu_b", "tool_b", "result b"),
			},
			wantErr: `tool_result "toolu_b"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, _, err := toAnthropicMessagesInput(t.Context(), "", tc.inputs)
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("unexpected error:\n got  %q\n want substring %q", err.Error(), tc.wantErr)
			}
		})
	}
}

func userText(text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindInputMessage,
		InputMessage: &spec.InputOutputContent{
			Role: spec.RoleUser,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				},
			},
		},
	}
}

func assistantText(text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindOutputMessage,
		OutputMessage: &spec.InputOutputContent{
			Role: spec.RoleAssistant,
			Contents: []spec.InputOutputContentItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				},
			},
		},
	}
}

func functionToolCall(id, callID, name, args string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindFunctionToolCall,
		FunctionToolCall: &spec.ToolCall{
			Type:      spec.ToolTypeFunction,
			Role:      spec.RoleAssistant,
			ID:        id,
			CallID:    callID,
			Name:      name,
			Arguments: args,
		},
	}
}

func functionToolOutput(callID, name, text string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindFunctionToolOutput,
		FunctionToolOutput: &spec.ToolOutput{
			Type:   spec.ToolTypeFunction,
			Role:   spec.RoleTool,
			CallID: callID,
			Name:   name,
			Contents: []spec.ToolOutputItemUnion{
				{
					Kind:     spec.ContentItemKindText,
					TextItem: &spec.ContentItemText{Text: text},
				},
			},
		},
	}
}

func webSearchToolCall(id, query string, input map[string]any) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindWebSearchToolCall,
		WebSearchToolCall: &spec.ToolCall{
			Type:   spec.ToolTypeWebSearch,
			Role:   spec.RoleAssistant,
			ID:     id,
			CallID: id,
			Name:   spec.DefaultWebSearchToolName,
			WebSearchToolCallItems: []spec.WebSearchToolCallItemUnion{
				{
					Kind: spec.WebSearchToolCallKindSearch,
					SearchItem: &spec.WebSearchToolCallSearch{
						Query: query,
						Input: input,
					},
				},
			},
		},
	}
}

func webSearchToolOutput(callID, url, title, encryptedContent string) spec.InputUnion {
	return spec.InputUnion{
		Kind: spec.InputKindWebSearchToolOutput,
		WebSearchToolOutput: &spec.ToolOutput{
			Type:   spec.ToolTypeWebSearch,
			Role:   spec.RoleAssistant,
			CallID: callID,
			Name:   spec.DefaultWebSearchToolName,
			WebSearchToolOutputItems: []spec.WebSearchToolOutputItemUnion{
				{
					Kind: spec.WebSearchToolOutputKindSearch,
					SearchItem: &spec.WebSearchToolOutputSearch{
						URL:              url,
						Title:            title,
						EncryptedContent: encryptedContent,
					},
				},
			},
		},
	}
}

func decodeMessages(t *testing.T, v any) []map[string]any {
	t.Helper()

	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var out []map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("json.Unmarshal failed: %v\njson=%s", err, string(b))
	}
	return out
}

func messageRoles(t *testing.T, msgs []map[string]any) []string {
	t.Helper()

	out := make([]string, 0, len(msgs))
	for i, msg := range msgs {
		role, ok := msg["role"].(string)
		if !ok {
			t.Fatalf("message %d missing string role: %#v", i, msg["role"])
		}
		out = append(out, role)
	}
	return out
}

func messageContentTypes(t *testing.T, msgs []map[string]any) [][]string {
	t.Helper()

	out := make([][]string, 0, len(msgs))
	for i, msg := range msgs {
		blocks := contentBlocks(t, msg)
		types := make([]string, 0, len(blocks))
		for j, block := range blocks {
			typ, ok := block["type"].(string)
			if !ok {
				t.Fatalf("message %d block %d missing string type: %#v", i, j, block["type"])
			}
			types = append(types, typ)
		}
		out = append(out, types)
	}
	return out
}

func contentBlocks(t *testing.T, msg map[string]any) []map[string]any {
	t.Helper()

	raw, ok := msg["content"].([]any)
	if !ok {
		t.Fatalf("message content is not []any: %#v", msg["content"])
	}

	out := make([]map[string]any, 0, len(raw))
	for i, item := range raw {
		block, ok := item.(map[string]any)
		if !ok {
			t.Fatalf("content item %d is not map[string]any: %#v", i, item)
		}
		out = append(out, block)
	}
	return out
}

func stringField(t *testing.T, m map[string]any, key string) string {
	t.Helper()

	v, ok := m[key].(string)
	if !ok {
		t.Fatalf("field %q is not a string: %#v", key, m[key])
	}
	return v
}
