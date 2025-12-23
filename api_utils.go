package inference

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"strings"

	"github.com/ppipada/inference-go/internal/debugclient"
	"github.com/ppipada/inference-go/internal/maputil"
	"github.com/ppipada/inference-go/spec"
)

// IsInputUnionEmpty reports whether an InputUnion has nothing "worth sending".
// Metadata-only messages (e.g. only IDs/roles with no actual content) are
// treated as empty. Adjust the helpers below if your semantics differ.
func IsInputUnionEmpty(in spec.InputUnion) bool {
	switch in.Kind {
	case spec.InputKindInputMessage:
		return isInputOutputContentEmpty(in.InputMessage)
	case spec.InputKindOutputMessage:
		return isInputOutputContentEmpty(in.OutputMessage)
	case spec.InputKindReasoningMessage:
		return isReasoningContentEmpty(in.ReasoningMessage)
	case spec.InputKindFunctionToolCall:
		return in.FunctionToolCall == nil
	case spec.InputKindFunctionToolOutput:
		return in.FunctionToolOutput == nil
	case spec.InputKindCustomToolCall:
		return in.CustomToolCall == nil
	case spec.InputKindCustomToolOutput:
		return in.CustomToolOutput == nil
	case spec.InputKindWebSearchToolCall:
		return in.WebSearchToolCall == nil
	case spec.InputKindWebSearchToolOutput:
		return in.WebSearchToolOutput == nil
	default:
		// Zero-value or unknown kind -> nothing to send.
		return true
	}
}

func isInputOutputContentEmpty(c *spec.InputOutputContent) bool {
	if c == nil {
		return true
	}
	if len(c.Contents) == 0 {
		// No content items at all -> nothing worth sending.
		return true
	}

	for _, it := range c.Contents {
		if !isContentItemEmpty(it) {
			return false
		}
	}
	return true
}

func isReasoningContentEmpty(r *spec.ReasoningContent) bool {
	if r == nil {
		return true
	}
	// Only reasoning text / encrypted content is considered "worth sending".
	return len(r.Summary) == 0 &&
		len(r.Thinking) == 0 &&
		len(r.RedactedThinking) == 0 &&
		len(r.EncryptedContent) == 0
}

func isContentItemEmpty(it spec.InputOutputContentItemUnion) bool {
	switch it.Kind {
	case spec.ContentItemKindText:
		if it.TextItem == nil {
			return true
		}
		return it.TextItem.Text == "" && len(it.TextItem.Citations) == 0

	case spec.ContentItemKindRefusal:
		if it.RefusalItem == nil {
			return true
		}
		return it.RefusalItem.Refusal == ""

	case spec.ContentItemKindImage:
		if it.ImageItem == nil {
			return true
		}
		img := it.ImageItem
		return img.ID == "" &&
			img.Detail == "" &&
			img.ImageName == "" &&
			img.ImageMIME == "" &&
			img.ImageURL == "" &&
			img.ImageData == ""

	case spec.ContentItemKindFile:
		if it.FileItem == nil {
			return true
		}
		f := it.FileItem
		return f.ID == "" &&
			f.FileName == "" &&
			f.FileURL == "" &&
			f.FileData == "" &&
			f.AdditionalContext == "" &&
			f.CitationConfig == nil

	default:
		// Unknown or zero-value kind.
		return true
	}
}

// attachDebugResp adds HTTP-debug information and error context—without panics.
//
// - ctx may or may not contain debug information.
// - respErr is the transport/SDK error (may be nil).
// - isNilResp tells whether the model returned an empty/invalid response.
// - rawModelJSON is an optional, provider-level JSON representation of the *final* model response (e.g. OpenAI
// responses `resp.RawJSON()` or `json.Marshal(fullResponse)` for other SDKs). If provided and the HTTP debug layer
// did not already set ResponseDetails.Data, we will sanitize and store this JSON there.
func attachDebugResp(
	ctx context.Context,
	completionResp *FetchCompletionResponse,
	respErr error,
	isNilResp bool,
	fullObj any,
) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("attach debug resp panic",
				"recover", r,
				"stack", string(debug.Stack()))
		}
	}()

	if completionResp == nil {
		return
	}

	debugDetails := map[string]any{
		"requestDetails":  map[string]any{},
		"responseDetails": map[string]any{},
		"errorDetails":    map[string]any{},
	}
	completionResp.DebugDetails = debugDetails

	debugResp, _ := debugclient.GetDebugHTTPResponse(ctx)

	// Always attach request/response debug info from the HTTP layer if available.
	if debugResp != nil {
		if debugResp.RequestDetails != nil {
			if d, err := maputil.StructWithJSONTagsToMap(debugResp.RequestDetails); err == nil {
				debugDetails["requestDetails"] = d
			}
		}
		if debugResp.ResponseDetails != nil {
			if d, err := maputil.StructWithJSONTagsToMap(debugResp.ResponseDetails); err == nil {
				debugDetails["responseDetails"] = d
			}
		}
	}

	// If the HTTP layer didn't populate ResponseDetails.Data (most common in
	// streaming/SSE cases), and we have a provider-level raw JSON for the final
	// model response, sanitize that and use it as the debug body.

	if fullObj != nil {
		// We got a object. Lets replace always.
		if m, err := maputil.StructWithJSONTagsToMap(fullObj); err == nil {
			if d, ok := debugDetails["responseDetails"].(map[string]any); ok {
				d["data"] = debugclient.ScrubAnyForDebug(m, true)
			}
		}
	}

	// Gather error-message fragments.
	var msgParts []string
	if debugResp != nil && debugResp.ErrorDetails != nil {
		if m := strings.TrimSpace(debugResp.ErrorDetails.Message); m != "" {
			msgParts = append(msgParts, m)
		}
	}
	if respErr != nil {
		msgParts = append(msgParts, respErr.Error())
	}
	if isNilResp {
		msgParts = append(msgParts, "got nil response from LLM api")
	}

	if len(msgParts) == 0 {
		// Nothing more to add; request/response details (if any) are already attached.

		return
	}

	// Prepare ErrorDetails without aliasing the debug struct pointer.
	if debugResp != nil && debugResp.ErrorDetails != nil {
		ed := *debugResp.ErrorDetails
		ed.Message = strings.Join(msgParts, "; ")

		if d, err := maputil.StructWithJSONTagsToMap(ed); err == nil {
			debugDetails["errorDetails"] = d
		}

	} else {
		if d, ok := debugDetails["errorDetails"].(map[string]any); ok {
			d["message"] = strings.Join(msgParts, "; ")
		}
	}
}

// toolName is a pair of an internal tool choice and the function name
// that will be sent to the API for that tool.
type toolName struct {
	Choice spec.ToolChoice
	Name   string
}

// buildToolChoiceNameMapping assigns short, human‑readable function names to tools.
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
func buildToolChoiceNameMapping(
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
			name = fmt.Sprintf("%s_%d", base, count)
		}

		toolNames = append(toolNames, toolName{
			Choice: ct,
			Name:   name,
		})
		toolNameMap[name] = ct
	}

	return toolNames, toolNameMap
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

func toolDescription(ct spec.ToolChoice) string {
	if desc := strings.TrimSpace(ct.Description); desc != "" {
		return desc
	}
	return ct.Name
}
