package anthropicsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/flexigpt/inference-go/internal/logutil"
	"github.com/flexigpt/inference-go/internal/sdkutil"
	"github.com/flexigpt/inference-go/spec"
)

var errEmptyInputPart = errors.New("empty input part")

type anthropicInputPart struct {
	Role                spec.RoleEnum
	Blocks              []anthropic.ContentBlockParamUnion
	UserToolResult      bool
	ClientToolUseIDs    []string
	ClientToolResultIDs []string
}

type anthropicTurnDraft struct {
	Role                spec.RoleEnum
	Blocks              []anthropic.ContentBlockParamUnion
	ToolResultBlocks    []anthropic.ContentBlockParamUnion
	ClientToolUseIDs    []string
	ClientToolResultIDs []string
}

// toAnthropicMessagesInput converts a sequence of generic InputUnion items into
// Anthropic MessageParam and system prompt blocks.
func toAnthropicMessagesInput(
	_ context.Context,
	systemPrompt string,
	inputs []spec.InputUnion,
) (msgs []anthropic.MessageParam, sysPrompts []anthropic.TextBlockParam, err error) {
	var out []anthropic.MessageParam
	var turns []anthropicTurnDraft
	var sysParts []string

	if s := strings.TrimSpace(systemPrompt); s != "" {
		sysParts = append(sysParts, s)
	}

	for _, in := range inputs {
		if sdkutil.IsInputUnionEmpty(in) {
			continue
		}
		part, partErr := inputUnionToAnthropicPart(in)
		if partErr != nil {
			if errors.Is(partErr, errEmptyInputPart) {
				continue
			}
			return nil, nil, partErr
		}
		if part == nil || len(part.Blocks) == 0 {
			continue
		}
		turns = appendAnthropicTurnPart(turns, part)
	}

	if err := validateAnthropicToolTurnOrdering(turns); err != nil {
		return nil, nil, err
	}

	out = make([]anthropic.MessageParam, 0, len(turns))
	for _, turn := range turns {
		blocks := emitBlocks(turn)
		if len(blocks) == 0 {
			continue
		}
		if turn.Role == spec.RoleUser {
			out = append(out, anthropic.NewUserMessage(blocks...))
		} else {
			out = append(out, anthropic.NewAssistantMessage(blocks...))
		}

	}

	// System prompt as a single text block.
	if len(sysParts) > 0 {
		sysStr := strings.Join(sysParts, "\n\n")
		sysPrompts = append(sysPrompts, anthropic.TextBlockParam{Text: sysStr})
	}

	return out, sysPrompts, nil
}

func emitBlocks(t anthropicTurnDraft) []anthropic.ContentBlockParamUnion {
	if t.Role == spec.RoleUser && len(t.ToolResultBlocks) > 0 {
		out := make([]anthropic.ContentBlockParamUnion, 0, len(t.ToolResultBlocks)+len(t.Blocks))
		out = append(out, t.ToolResultBlocks...)
		out = append(out, t.Blocks...)
		return out
	}
	return t.Blocks
}

func appendAnthropicTurnPart(
	turns []anthropicTurnDraft,
	part *anthropicInputPart,
) []anthropicTurnDraft {
	if part == nil || len(part.Blocks) == 0 {
		return turns
	}

	if len(turns) == 0 || turns[len(turns)-1].Role != part.Role {
		turns = append(turns, anthropicTurnDraft{Role: part.Role})
	}

	turn := &turns[len(turns)-1]
	if part.Role == spec.RoleUser && part.UserToolResult {
		turn.ToolResultBlocks = append(turn.ToolResultBlocks, part.Blocks...)
	} else {
		turn.Blocks = append(turn.Blocks, part.Blocks...)
	}
	turn.ClientToolUseIDs = append(turn.ClientToolUseIDs, part.ClientToolUseIDs...)
	turn.ClientToolResultIDs = append(turn.ClientToolResultIDs, part.ClientToolResultIDs...)

	return turns
}

func inputUnionToAnthropicPart(in spec.InputUnion) (*anthropicInputPart, error) {
	switch in.Kind {
	case spec.InputKindInputMessage:
		if in.InputMessage == nil || in.InputMessage.Role != spec.RoleUser {
			return nil, errEmptyInputPart
		}
		blocks := contentItemsToAnthropicContentBlocks(in.InputMessage.Contents)
		blocks = applyAnthropicContentBlockCacheControl(blocks, in.InputMessage.CacheControl)
		if len(blocks) == 0 {
			return nil, errEmptyInputPart
		}
		return &anthropicInputPart{
			Role:   spec.RoleUser,
			Blocks: blocks,
		}, nil

	case spec.InputKindOutputMessage:
		if in.OutputMessage == nil || in.OutputMessage.Role != spec.RoleAssistant {
			return nil, errEmptyInputPart
		}
		blocks := contentItemsToAnthropicContentBlocks(in.OutputMessage.Contents)
		blocks = applyAnthropicContentBlockCacheControl(blocks, in.OutputMessage.CacheControl)
		if len(blocks) == 0 {
			return nil, errEmptyInputPart
		}
		return &anthropicInputPart{
			Role:   spec.RoleAssistant,
			Blocks: blocks,
		}, nil

	case spec.InputKindReasoningMessage:
		if in.ReasoningMessage == nil {
			return nil, errEmptyInputPart
		}
		block := reasoningContentToAnthropicBlocks(in.ReasoningMessage)
		if block == nil {
			return nil, errEmptyInputPart
		}
		return &anthropicInputPart{
			Role:   spec.RoleAssistant,
			Blocks: []anthropic.ContentBlockParamUnion{*block},
		}, nil

	case spec.InputKindFunctionToolCall, spec.InputKindCustomToolCall, spec.InputKindWebSearchToolCall:
		var call *spec.ToolCall
		switch {
		case in.FunctionToolCall != nil:
			call = in.FunctionToolCall
		case in.CustomToolCall != nil:
			call = in.CustomToolCall
		case in.WebSearchToolCall != nil:
			call = in.WebSearchToolCall
		}
		if call == nil {
			return nil, errEmptyInputPart
		}
		if id := anthropicToolUseIDFromCall(call); id == "" {
			return nil, fmt.Errorf("anthropic: tool call %q is missing id/callID", call.Name)
		}

		block := toolCallToAnthropicToolUseBlock(call)
		if block == nil {
			return nil, errEmptyInputPart
		}

		part := &anthropicInputPart{
			Role:   spec.RoleAssistant,
			Blocks: []anthropic.ContentBlockParamUnion{*block},
		}
		if call.Type != spec.ToolTypeWebSearch {
			part.ClientToolUseIDs = []string{anthropicToolUseIDFromCall(call)}
		}
		return part, nil

	case spec.InputKindFunctionToolOutput, spec.InputKindCustomToolOutput, spec.InputKindWebSearchToolOutput:
		isWebSearchOutput := false
		var output *spec.ToolOutput
		switch {
		case in.FunctionToolOutput != nil:
			output = in.FunctionToolOutput
		case in.CustomToolOutput != nil:
			output = in.CustomToolOutput
		case in.WebSearchToolOutput != nil:
			output = in.WebSearchToolOutput
			isWebSearchOutput = true
		}
		if output == nil {
			return nil, errEmptyInputPart
		}
		if !isWebSearchOutput {
			if id := anthropicToolResultIDFromOutput(output); id == "" {
				return nil, errors.New("anthropic: client tool output is missing callID")
			}
		}

		block := toolOutputToAnthropicBlocks(output)
		if block == nil {
			return nil, errEmptyInputPart
		}

		if isWebSearchOutput {
			return &anthropicInputPart{
				Role:   spec.RoleAssistant,
				Blocks: []anthropic.ContentBlockParamUnion{*block},
			}, nil
		}

		return &anthropicInputPart{
			Role:                spec.RoleUser,
			Blocks:              []anthropic.ContentBlockParamUnion{*block},
			UserToolResult:      true,
			ClientToolResultIDs: []string{anthropicToolResultIDFromOutput(output)},
		}, nil

	default:
		return nil, errEmptyInputPart
	}
}

func validateAnthropicToolTurnOrdering(turns []anthropicTurnDraft) error {
	if len(turns) == 0 {
		return nil
	}

	var pendingOrdered []string
	pendingSet := map[string]struct{}{}

	clearPending := func() {
		pendingOrdered = nil
		pendingSet = map[string]struct{}{}
	}
	clearPending()

	for i, turn := range turns {
		if len(turn.ClientToolResultIDs) > 0 {
			if turn.Role != spec.RoleUser {
				return errors.New("anthropic: tool_result blocks must be sent in a user turn")
			}
			if len(pendingSet) == 0 {
				return fmt.Errorf(
					"anthropic: unexpected user tool_result turn at index %d without a preceding assistant client tool_use turn",
					i,
				)
			}

			seen := map[string]struct{}{}
			for _, id := range turn.ClientToolResultIDs {
				id = strings.TrimSpace(id)
				if id == "" {
					return fmt.Errorf("anthropic: tool_result in user turn at index %d is missing call_id", i)
				}
				if _, ok := pendingSet[id]; !ok {
					return fmt.Errorf(
						"anthropic: tool_result %q in user turn at index %d does not match the immediately preceding assistant client tool_use turn",
						id,
						i,
					)
				}
				if _, dup := seen[id]; dup {
					return fmt.Errorf("anthropic: duplicate tool_result %q in user turn at index %d", id, i)
				}
				seen[id] = struct{}{}
			}

			if len(seen) != len(pendingSet) {
				missing := make([]string, 0, len(pendingOrdered))
				for _, id := range pendingOrdered {
					if _, ok := seen[id]; !ok {
						missing = append(missing, id)
					}
				}
				return errors.New(
					"anthropic: assistant client tool_use turn must be answered by the immediately following user turn; missing tool_result for " + strings.Join(
						missing,
						", ",
					),
				)
			}

			clearPending()
		} else if len(pendingSet) > 0 {
			if turn.Role == spec.RoleUser {
				return errors.New(
					"anthropic: the user turn immediately following an assistant client tool_use turn must begin with matching tool_result blocks",
				)
			}
			return errors.New(
				"anthropic: assistant client tool_use turn must be immediately followed by a user tool_result turn",
			)
		}

		if len(turn.ClientToolUseIDs) > 0 {
			clearPending()
			for _, id := range turn.ClientToolUseIDs {
				id = strings.TrimSpace(id)
				if id == "" {
					return fmt.Errorf("anthropic: assistant tool_use turn at index %d contains an empty tool_use id", i)
				}
				if _, dup := pendingSet[id]; dup {
					return fmt.Errorf("anthropic: duplicate tool_use id %q in assistant turn at index %d", id, i)
				}
				pendingSet[id] = struct{}{}
				pendingOrdered = append(pendingOrdered, id)
			}
		}
	}

	if len(pendingSet) > 0 {
		return errors.New(
			"anthropic: input ends with an assistant client tool_use turn but no following user tool_result turn",
		)
	}
	return nil
}

// contentItemsToAnthropicContentBlocks converts generic content items into Anthropic
// content blocks (text/image/document).
func contentItemsToAnthropicContentBlocks(
	items []spec.InputOutputContentItemUnion,
) []anthropic.ContentBlockParamUnion {
	if len(items) == 0 {
		return nil
	}
	out := make([]anthropic.ContentBlockParamUnion, 0, len(items))

	for _, it := range items {
		switch it.Kind {
		case spec.ContentItemKindText:
			tb := contentItemTextToAnthropicTextBlockParam(it.TextItem)
			if tb != nil {
				out = append(out, anthropic.ContentBlockParamUnion{OfText: tb})
			}

		case spec.ContentItemKindImage:
			ib := contentItemImageToAnthropicImageBlockParam(it.ImageItem)
			if ib != nil {
				out = append(out, anthropic.ContentBlockParamUnion{OfImage: ib})
			}

		case spec.ContentItemKindFile:
			db := contentItemFileToAnthropicDocumentBlockParam(it.FileItem)
			if db != nil {
				out = append(out, anthropic.ContentBlockParamUnion{OfDocument: db})
			}

		case spec.ContentItemKindRefusal:
			// Anthropic does not have a dedicated "refusal" content block type.
			// Refusals are conveyed via stop_reason="refusal". We don't send
			// refusals back as input content.
			continue

		default:
			logutil.Debug("anthropic: unknown content item kind for message", "kind", it.Kind)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func reasoningContentToAnthropicBlocks(
	r *spec.ReasoningContent,
) *anthropic.ContentBlockParamUnion {
	if r == nil {
		return nil
	}

	if len(r.RedactedThinking) > 0 {
		// If redacted thinking is present it is redacted thinking block.
		data := strings.Join(r.RedactedThinking, " ")
		out := anthropic.NewRedactedThinkingBlock(data)
		return &out
	}

	if len(r.Thinking) > 0 && r.Signature != "" {
		data := strings.Join(r.Thinking, " ")
		out := anthropic.NewThinkingBlock(r.Signature, data)
		return &out
	}
	return nil
}

// toolCallToAnthropicToolUseBlock converts a ToolCall into an Anthropic tool_use block.
func toolCallToAnthropicToolUseBlock(
	toolCall *spec.ToolCall,
) *anthropic.ContentBlockParamUnion {
	toolUseID := anthropicToolUseIDFromCall(toolCall)
	if toolCall == nil || toolUseID == "" {
		return nil
	}
	switch toolCall.Type {
	case spec.ToolTypeFunction, spec.ToolTypeCustom:
		if strings.TrimSpace(toolCall.Name) == "" {
			return nil
		}
		args := strings.TrimSpace(toolCall.Arguments)
		if args == "" {
			args = "{}"
		}
		raw := json.RawMessage(args)

		out := anthropic.ContentBlockParamUnion{OfToolUse: &anthropic.ToolUseBlockParam{
			ID:    toolUseID,
			Name:  toolCall.Name,
			Input: raw,
		}}
		applyAnthropicToolUseCacheControl(&out, toolCall.CacheControl)
		return &out

	case spec.ToolTypeWebSearch:
		if len(toolCall.WebSearchToolCallItems) == 0 {
			return nil
		}
		// Anthropic has only 1 web search call item as of now.
		wcall := toolCall.WebSearchToolCallItems[0]
		if wcall.Kind != spec.WebSearchToolCallKindSearch || wcall.SearchItem == nil || wcall.SearchItem.Input == nil {
			// Only search supported.
			return nil
		}

		out := anthropic.ContentBlockParamUnion{OfServerToolUse: &anthropic.ServerToolUseBlockParam{
			ID:    toolUseID,
			Input: wcall.SearchItem.Input,
			Name:  anthropic.ServerToolUseBlockParamNameWebSearch,
		}}
		applyAnthropicToolUseCacheControl(&out, toolCall.CacheControl)
		return &out

	}
	return nil
}

func anthropicToolUseIDFromCall(toolCall *spec.ToolCall) string {
	if toolCall == nil {
		return ""
	}
	if id := strings.TrimSpace(toolCall.ID); id != "" {
		return id
	}
	return strings.TrimSpace(toolCall.CallID)
}

func anthropicToolResultIDFromOutput(toolOutput *spec.ToolOutput) string {
	if toolOutput == nil {
		return ""
	}
	return strings.TrimSpace(toolOutput.CallID)
}

func toolOutputToAnthropicBlocks(
	toolOutput *spec.ToolOutput,
) *anthropic.ContentBlockParamUnion {
	if toolOutput == nil || strings.TrimSpace(toolOutput.CallID) == "" {
		return nil
	}

	switch toolOutput.Type {
	case spec.ToolTypeFunction, spec.ToolTypeCustom:
		items := contentItemsToAnthropicToolResultBlocks(toolOutput.Contents)
		if len(items) == 0 {
			return nil
		}
		toolBlock := anthropic.ToolResultBlockParam{
			ToolUseID: toolOutput.CallID,
			Content:   items,
			IsError:   anthropic.Bool(toolOutput.IsError),
		}
		out := anthropic.ContentBlockParamUnion{OfToolResult: &toolBlock}
		applyAnthropicToolResultCacheControl(&out, toolOutput.CacheControl)
		return &out

	case spec.ToolTypeWebSearch:
		content := webSearchToolOutputItemsToAnthropicWebSearchContent(
			toolOutput.WebSearchToolOutputItems,
		)
		if content == nil {
			return nil
		}
		wsBlock := anthropic.WebSearchToolResultBlockParam{
			ToolUseID: toolOutput.CallID,
			Content:   *content,
			// Type omitted; zero value marshals as "web_search_tool_result".
		}
		out := anthropic.ContentBlockParamUnion{OfWebSearchToolResult: &wsBlock}
		applyAnthropicToolResultCacheControl(&out, toolOutput.CacheControl)
		return &out
	default:
		// Nothing to do more.
	}
	return nil
}

func contentItemsToAnthropicToolResultBlocks(
	items []spec.ToolOutputItemUnion,
) []anthropic.ToolResultBlockParamContentUnion {
	if len(items) == 0 {
		return nil
	}
	out := make([]anthropic.ToolResultBlockParamContentUnion, 0, len(items))

	for _, it := range items {
		switch it.Kind {
		case spec.ContentItemKindText:
			tb := contentItemTextToAnthropicTextBlockParam(it.TextItem)
			if tb != nil {
				out = append(out, anthropic.ToolResultBlockParamContentUnion{OfText: tb})
			}

		case spec.ContentItemKindImage:
			ib := contentItemImageToAnthropicImageBlockParam(it.ImageItem)
			if ib != nil {
				out = append(out, anthropic.ToolResultBlockParamContentUnion{OfImage: ib})
			}

		case spec.ContentItemKindFile:
			db := contentItemFileToAnthropicDocumentBlockParam(it.FileItem)
			if db != nil {
				out = append(out, anthropic.ToolResultBlockParamContentUnion{OfDocument: db})
			}
		case spec.ContentItemKindRefusal:
			// Invalid for this.
		default:
			logutil.Debug("anthropic: unknown content item kind for message", "kind", it.Kind)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func webSearchToolOutputItemsToAnthropicWebSearchContent(
	items []spec.WebSearchToolOutputItemUnion,
) *anthropic.WebSearchToolResultBlockParamContentUnion {
	if len(items) == 0 {
		return nil
	}

	// If there's an error item, treat the whole tool call as an error.
	for _, it := range items {
		if it.Kind == spec.WebSearchToolOutputKindError && it.ErrorItem != nil {
			errParam := anthropic.WebSearchToolRequestErrorParam{
				// Code is something like "invalid_tool_input", "unavailable", etc.
				ErrorCode: anthropic.WebSearchToolRequestErrorErrorCode(it.ErrorItem.Code),
				// Type is omitted; zero value marshals as "web_search_tool_result_error".
			}
			return &anthropic.WebSearchToolResultBlockParamContentUnion{
				OfRequestWebSearchToolResultError: &errParam,
			}
		}
	}

	// Otherwise, collect all search results.
	results := make([]anthropic.WebSearchResultBlockParam, 0, len(items))

	for _, it := range items {
		if it.Kind != spec.WebSearchToolOutputKindSearch || it.SearchItem == nil {
			continue
		}

		ws := it.SearchItem

		block := anthropic.WebSearchResultBlockParam{
			URL:              ws.URL,
			Title:            ws.Title,
			EncryptedContent: ws.EncryptedContent,
			// Type omitted; zero value marshals as "web_search_result".
		}

		// Optional page_age.
		if s := strings.TrimSpace(ws.PageAge); s != "" {
			block.PageAge = anthropic.String(s)
		}

		results = append(results, block)
	}

	if len(results) == 0 {
		return nil
	}

	return &anthropic.WebSearchToolResultBlockParamContentUnion{
		OfWebSearchToolResultBlockItem: results,
	}
}

func contentItemTextToAnthropicTextBlockParam(textItem *spec.ContentItemText) *anthropic.TextBlockParam {
	if textItem == nil {
		return nil
	}
	text := strings.TrimSpace(textItem.Text)
	if text == "" {
		return nil
	}
	tb := &anthropic.TextBlockParam{
		Text: text,
	}

	if anns := citationsToAnthropicTextCitations(textItem.Citations); len(anns) > 0 {
		tb.Citations = anns
	}
	return tb
}

// citationsToAnthropicTextCitations converts our generic URL citations into
// Anthropic TextCitationParamUnion (web_search_result_location).
func citationsToAnthropicTextCitations(
	citations []spec.Citation,
) []anthropic.TextCitationParamUnion {
	if len(citations) == 0 {
		return nil
	}
	out := make([]anthropic.TextCitationParamUnion, 0)
	for _, c := range citations {
		if c.URLCitation == nil {
			continue
		}
		out = append(out, anthropic.TextCitationParamUnion{
			OfWebSearchResultLocation: &anthropic.CitationWebSearchResultLocationParam{
				CitedText:      c.URLCitation.CitedText,
				EncryptedIndex: c.URLCitation.EncryptedIndex,
				Title:          anthropic.String(c.URLCitation.Title),
				URL:            c.URLCitation.URL,
			},
		})
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func contentItemImageToAnthropicImageBlockParam(imageItem *spec.ContentItemImage) *anthropic.ImageBlockParam {
	if imageItem == nil {
		return nil
	}

	if data := strings.TrimSpace(imageItem.ImageData); data != "" {
		mime := strings.TrimSpace(imageItem.ImageMIME)
		if mime == "" {
			mime = spec.DefaultImageDataMIME
		}
		return &anthropic.ImageBlockParam{
			Source: anthropic.ImageBlockParamSourceUnion{
				OfBase64: &anthropic.Base64ImageSourceParam{
					Data:      data,
					MediaType: anthropic.Base64ImageSourceMediaType(mime),
				},
			},
		}
	} else if u := strings.TrimSpace(imageItem.ImageURL); u != "" {
		return &anthropic.ImageBlockParam{
			Source: anthropic.ImageBlockParamSourceUnion{
				OfURL: &anthropic.URLImageSourceParam{
					URL: u,
				},
			},
		}
	}
	return nil
}

func contentItemFileToAnthropicDocumentBlockParam(fileItem *spec.ContentItemFile) *anthropic.DocumentBlockParam {
	if fileItem == nil {
		return nil
	}
	data := strings.TrimSpace(fileItem.FileData)
	url := strings.TrimSpace(fileItem.FileURL)
	mime := strings.TrimSpace(fileItem.FileMIME)
	// Map files to document blocks where possible.
	switch {
	case data != "" && strings.HasPrefix(mime, "application/pdf"):
		return &anthropic.DocumentBlockParam{
			Source: anthropic.DocumentBlockParamSourceUnion{
				OfBase64: &anthropic.Base64PDFSourceParam{
					Data: data,
				},
			},
		}

	case url != "" && strings.HasPrefix(mime, "application/pdf"):
		return &anthropic.DocumentBlockParam{
			Source: anthropic.DocumentBlockParamSourceUnion{
				OfURL: &anthropic.URLPDFSourceParam{
					URL: url,
				},
			},
		}

	case data != "" && strings.HasPrefix(mime, "text/"):
		// For plain text, Anthropic expects actual text, not base64. If you
		// want to support this fully, decode base64 here. For now we skip.
		logutil.Debug("anthropic: skipping non-pdf base64 file; plain-text decoding not implemented",
			"id", fileItem.ID, "name", fileItem.FileName, "mime", mime)
	default:
		// Other file types not supported as document blocks.
	}
	return nil
}
