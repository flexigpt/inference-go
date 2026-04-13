package googlegeneratecontentsdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/genai"

	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var errEmptyInputPart = errors.New("empty input part")

// toGoogleGenerateContentContents converts a system prompt and a slice of spec
// InputUnion items into the []*genai.Content slice and optional system
// instruction *genai.Content expected by GenerateContent.
//
// Adjacent turns that share the same role are merged into a single Content so
// the conversation conforms to the expected user/model alternation pattern.
func toGoogleGenerateContentContents(
	_ context.Context,
	systemPrompt string,
	inputs []spec.InputUnion,
) (contents []*genai.Content, sysInstruction *genai.Content, err error) {
	if s := strings.TrimSpace(systemPrompt); s != "" {
		sysInstruction = &genai.Content{
			Parts: []*genai.Part{genai.NewPartFromText(s)},
		}
	}

	type draft struct {
		role  string
		parts []*genai.Part
	}

	var drafts []draft

	for _, in := range inputs {
		if sdkutil.IsInputUnionEmpty(in) {
			continue
		}

		role, parts, convErr := inputUnionToGenAIParts(in)
		if convErr != nil {
			if errors.Is(convErr, errEmptyInputPart) {
				continue
			}
			return nil, nil, convErr
		}
		if len(parts) == 0 {
			continue
		}

		// Merge into the last draft when the role matches.
		if len(drafts) > 0 && drafts[len(drafts)-1].role == role {
			drafts[len(drafts)-1].parts = append(drafts[len(drafts)-1].parts, parts...)
		} else {
			drafts = append(drafts, draft{role: role, parts: parts})
		}
	}

	if len(drafts) == 0 {
		return nil, sysInstruction, nil
	}

	out := make([]*genai.Content, 0, len(drafts))
	for _, d := range drafts {
		out = append(out, &genai.Content{Role: d.role, Parts: d.parts})
	}
	return out, sysInstruction, nil
}

// inputUnionToGenAIParts converts a single InputUnion to a (role, []*genai.Part) pair.
func inputUnionToGenAIParts(in spec.InputUnion) (role string, parts []*genai.Part, err error) {
	switch in.Kind {

	case spec.InputKindInputMessage:
		if in.InputMessage == nil || in.InputMessage.Role != spec.RoleUser {
			return "", nil, errEmptyInputPart
		}
		ps := contentItemsToGenAIParts(in.InputMessage.Contents)
		if len(ps) == 0 {
			return "", nil, errEmptyInputPart
		}
		return genai.RoleUser, ps, nil

	case spec.InputKindOutputMessage:
		if in.OutputMessage == nil || in.OutputMessage.Role != spec.RoleAssistant {
			return "", nil, errEmptyInputPart
		}
		ps := contentItemsToGenAIParts(in.OutputMessage.Contents)
		if len(ps) == 0 {
			return "", nil, errEmptyInputPart
		}
		return genai.RoleModel, ps, nil

	case spec.InputKindReasoningMessage:
		if in.ReasoningMessage == nil {
			return "", nil, errEmptyInputPart
		}
		p := reasoningContentToGenAIPart(in.ReasoningMessage)
		if p == nil {
			return "", nil, errEmptyInputPart
		}
		return genai.RoleModel, []*genai.Part{p}, nil

	case spec.InputKindFunctionToolCall, spec.InputKindCustomToolCall:
		var call *spec.ToolCall
		switch {
		case in.FunctionToolCall != nil:
			call = in.FunctionToolCall
		case in.CustomToolCall != nil:
			call = in.CustomToolCall
		}
		if call == nil {
			return "", nil, errEmptyInputPart
		}
		p, convErr := toolCallToGenAIFunctionCallPart(call)
		if convErr != nil {
			return "", nil, convErr
		}
		return genai.RoleModel, []*genai.Part{p}, nil

	case spec.InputKindFunctionToolOutput, spec.InputKindCustomToolOutput:
		var output *spec.ToolOutput
		switch {
		case in.FunctionToolOutput != nil:
			output = in.FunctionToolOutput
		case in.CustomToolOutput != nil:
			output = in.CustomToolOutput
		}
		if output == nil {
			return "", nil, errEmptyInputPart
		}
		if strings.TrimSpace(output.CallID) == "" {
			return "", nil, errors.New("googleGenerateContent: function tool output is missing callID")
		}
		if strings.TrimSpace(output.Name) == "" {
			return "", nil, errors.New("googleGenerateContent: function tool output is missing name")
		}
		p := toolOutputToGenAIFunctionResponsePart(output)
		if p == nil {
			return "", nil, errEmptyInputPart
		}
		return genai.RoleUser, []*genai.Part{p}, nil

	case spec.InputKindWebSearchToolCall, spec.InputKindWebSearchToolOutput:
		// Google grounding is handled transparently server-side; there are no
		// client-side tool_call / tool_result round-trips to echo back.
		return "", nil, errEmptyInputPart

	default:
		return "", nil, errEmptyInputPart
	}
}

// contentItemsToGenAIParts converts a slice of InputOutputContentItemUnion to
// genai.Part pointers, skipping unsupported types.
func contentItemsToGenAIParts(items []spec.InputOutputContentItemUnion) []*genai.Part {
	if len(items) == 0 {
		return nil
	}
	out := make([]*genai.Part, 0, len(items))
	for _, it := range items {
		switch it.Kind {
		case spec.ContentItemKindText:
			if it.TextItem != nil {
				if t := strings.TrimSpace(it.TextItem.Text); t != "" {
					out = append(out, genai.NewPartFromText(t))
				}
			}

		case spec.ContentItemKindImage:
			if p := contentItemImageToGenAIPart(it.ImageItem); p != nil {
				out = append(out, p)
			}

		case spec.ContentItemKindFile:
			if p := contentItemFileToGenAIPart(it.FileItem); p != nil {
				out = append(out, p)
			}

		case spec.ContentItemKindRefusal:
			// Refusals are model outputs; not a meaningful input representation.

		default:
			logutil.Debug("googleGenerateContent: unknown content item kind for message", "kind", it.Kind)
		}
	}
	return out
}

func contentItemImageToGenAIPart(imageItem *spec.ContentItemImage) *genai.Part {
	if imageItem == nil {
		return nil
	}

	mime := strings.TrimSpace(imageItem.ImageMIME)
	if mime == "" {
		mime = spec.DefaultImageDataMIME
	}

	if data := strings.TrimSpace(imageItem.ImageData); data != "" {
		raw, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			raw, err = base64.RawStdEncoding.DecodeString(data)
			if err != nil {
				logutil.Debug("googleGenerateContent: failed to decode base64 image data",
					"id", imageItem.ID, "err", err)
				return nil
			}
		}
		return genai.NewPartFromBytes(raw, mime)
	}

	if u := strings.TrimSpace(imageItem.ImageURL); u != "" {
		return genai.NewPartFromURI(u, mime)
	}
	return nil
}

func contentItemFileToGenAIPart(fileItem *spec.ContentItemFile) *genai.Part {
	if fileItem == nil {
		return nil
	}

	mime := strings.TrimSpace(fileItem.FileMIME)
	if mime == "" {
		mime = spec.DefaultFileDataMIME
	}

	if data := strings.TrimSpace(fileItem.FileData); data != "" {
		raw, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			raw, err = base64.RawStdEncoding.DecodeString(data)
			if err != nil {
				logutil.Debug("googleGenerateContent: failed to decode base64 file data",
					"id", fileItem.ID, "name", fileItem.FileName, "err", err)
				return nil
			}
		}
		return genai.NewPartFromBytes(raw, mime)
	}

	if u := strings.TrimSpace(fileItem.FileURL); u != "" {
		return genai.NewPartFromURI(u, mime)
	}
	return nil
}

// reasoningContentToGenAIPart converts a spec.ReasoningContent to a genai
// thought Part for history pass-back.
// Only Google-native reasoning (Thinking text + Signature) is converted;
// all other forms (EncryptedContent, RedactedThinking, Summary) are silently
// skipped since they originate from other providers.
func reasoningContentToGenAIPart(r *spec.ReasoningContent) *genai.Part {
	if r == nil {
		return nil
	}
	// Google-native: signed thought.
	if strings.TrimSpace(r.Signature) != "" {
		text := strings.TrimSpace(strings.Join(r.Thinking, " "))
		if text == "" {
			return nil
		}
		sig, ok := decodeThoughtSignature(r.Signature)
		if !ok {
			return nil
		}
		return &genai.Part{
			Text:             text,
			Thought:          true,
			ThoughtSignature: sig,
		}
	}
	return nil
}

// toolCallToGenAIFunctionCallPart converts a ToolCall to a genai FunctionCall
// Part for use in conversation history.
func toolCallToGenAIFunctionCallPart(call *spec.ToolCall) (*genai.Part, error) {
	if call == nil {
		return nil, errEmptyInputPart
	}

	name := strings.TrimSpace(call.Name)
	if name == "" {
		return nil, errors.New("googleGenerateContent: tool call is missing function name")
	}

	id := strings.TrimSpace(call.ID)
	if id == "" {
		id = strings.TrimSpace(call.CallID)
	}

	args := map[string]any{}
	if a := strings.TrimSpace(call.Arguments); a != "" && a != "{}" {
		if err := json.Unmarshal([]byte(a), &args); err != nil {
			return nil, fmt.Errorf("googleGenerateContent: tool call %q has invalid JSON arguments: %w", name, err)
		}
	}

	return &genai.Part{
		FunctionCall: &genai.FunctionCall{
			ID:   id,
			Name: name,
			Args: args,
		},
	}, nil
}

// toolOutputToGenAIFunctionResponsePart converts a ToolOutput to a genai
// FunctionResponse Part.
// Text contents are joined and placed under the "output" or "error" key as
// recommended by the Google GenAI API.
func toolOutputToGenAIFunctionResponsePart(output *spec.ToolOutput) *genai.Part {
	if output == nil {
		return nil
	}

	response := make(map[string]any)

	var textParts []string
	for _, c := range output.Contents {
		if c.Kind == spec.ContentItemKindText && c.TextItem != nil {
			if t := strings.TrimSpace(c.TextItem.Text); t != "" {
				textParts = append(textParts, t)
			}
		}
	}

	text := strings.Join(textParts, "\n")
	if output.IsError {
		if text != "" {
			response["error"] = text
		} else {
			response["error"] = "tool call failed"
		}
	} else if text != "" {
		response["output"] = text
	}

	return &genai.Part{
		FunctionResponse: &genai.FunctionResponse{
			ID:       strings.TrimSpace(output.CallID),
			Name:     strings.TrimSpace(output.Name),
			Response: response,
		},
	}
}
