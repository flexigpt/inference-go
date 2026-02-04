package sdkutil

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flexigpt/inference-go/spec"
)

// toolName is a pair of an internal tool choice and the function name
// that will be sent to the API for that tool.
type toolName struct {
	Choice spec.ToolChoice
	Name   string
}

// BuildToolChoiceNameMapping assigns short, human‑readable function names to tools.
//
// Rules:
//   - base name is the sanitized tool slug (lower‑cased, [a‑z0‑9_-] only)
//   - first tool with a given slug gets "<slug>"
//   - subsequent tools with the same slug get "<slug>_2", "<slug>_3", ...
//   - names are truncated to 64 characters (OpenAI function-tool limit)
//
// Returns:
//   - ordered: same cardinality/order as input tools, but with the
//     derived function name for each tool.
//   - nameToTool: map[functionName] => FetchCompletionToolChoice; used
//     to translate tool calls back to the original identity.
func BuildToolChoiceNameMapping(
	tools []spec.ToolChoice,
) (toolNames []toolName, toolNameMap map[string]spec.ToolChoice) {
	if len(tools) == 0 {
		return nil, nil
	}

	used := make(map[string]int, len(tools))
	toolNameMap = make(map[string]spec.ToolChoice, len(tools))
	toolNames = make([]toolName, 0, len(tools))

	for _, ct := range tools {
		// Base is the sanitized tool slug.
		base := sanitizeToolNameComponent(ct.Name)
		if base == "" {
			base = "tool"
		}

		// Enforce 64‑char limit.
		if len(base) > 64 {
			base = base[:64]
		}
		// Deduplicate: slug, slug_2, slug_3, ...
		count := used[base]
		var name string
		if count == 0 {
			name = base
			used[base] = 1
		} else {
			count++
			used[base] = count
			suffix := fmt.Sprintf("_%d", count)
			// Ensure final name <= 64.
			maxBaseLen := 64 - len(suffix)
			adj := base
			if maxBaseLen < 1 {
				// Extremely defensive; shouldn't happen with our suffix format.
				adj = "tool"
				maxBaseLen = 64 - len(suffix)
			}
			if len(adj) > maxBaseLen {
				adj = adj[:maxBaseLen]
			}
			name = adj + suffix
		}

		toolNames = append(toolNames, toolName{
			Choice: ct,
			Name:   name,
		})
		toolNameMap[name] = ct
	}

	return toolNames, toolNameMap
}

type ResolvedAllowedTool struct {
	Type spec.ToolType
	Name string
}

func ResolveAllowedTools(
	allowed []spec.AllowedTool,
	toolChoiceNameMap map[string]spec.ToolChoice,
) ([]ResolvedAllowedTool, error) {
	if len(toolChoiceNameMap) == 0 {
		return nil, errors.New("got empty toolChoiceNameMap")
	}
	if len(allowed) == 0 {
		return nil, errors.New("got empty allowed tool choices")
	}

	resolvedTools := make([]ResolvedAllowedTool, 0, len(allowed))
	for _, a := range allowed {
		n := strings.TrimSpace(a.ToolChoiceName)
		id := strings.TrimSpace(a.ToolChoiceID)
		if n == "" && id == "" {
			// No name or id in input, invalid choice.
			continue
		}
		// Find matching tool by ID and/or by name.
		// Note: toolChoiceNameMap key is the *provider/API* tool name; tc.Name is the original ToolChoice.Name.
		found := false
		for apiName, tc := range toolChoiceNameMap {
			if id != "" && tc.ID != id {
				continue
			}
			if n != "" {
				// Accept either original tool name or the derived provider name.
				if !strings.EqualFold(tc.Name, n) && !strings.EqualFold(apiName, n) {
					continue
				}
			}
			if tc.Type != spec.ToolTypeFunction && tc.Type != spec.ToolTypeCustom {
				// Skip unsupported tool types for "allowed tools" resolution.
				continue
			}
			resolvedTools = append(resolvedTools, ResolvedAllowedTool{
				Type: tc.Type,
				Name: apiName,
			})
			found = true
			break
		}
		if !found {
			// Keep scanning other allowed entries; final error if none resolved.
			continue
		}
	}
	if len(resolvedTools) > 0 {
		return resolvedTools, nil
	}
	return nil, errors.New("no eligible allowed tool found")
}

func sanitizeToolNameComponent(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '_' || r == '-':
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}
	out := strings.Trim(b.String(), "_-")
	return out
}

func ToolDescription(ct spec.ToolChoice) string {
	if desc := strings.TrimSpace(ct.Description); desc != "" {
		return desc
	}
	return ct.Name
}
