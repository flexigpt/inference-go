package googleinteractionssdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"strings"

	"google.golang.org/genai/interactions/models/interactions"

	"github.com/flexigpt/inference-go/spec"
)

type googleInteractionsRequestBuild struct {
	Interaction       interactions.CreateModelInteraction
	ToolChoiceNameMap map[string]spec.ToolChoice
	WebSearchChoiceID string
	Warnings          []spec.Warning
}

func buildGoogleInteractionsCreateInteraction(
	req *spec.FetchCompletionRequest,
	stream bool,
) (googleInteractionsRequestBuild, error) {
	if req == nil {
		return googleInteractionsRequestBuild{}, errors.New(
			"google interactions: nil completion request",
		)
	}

	var build googleInteractionsRequestBuild
	steps, err := googleInteractionsInputSteps(req.Inputs, &build.Warnings)
	if err != nil {
		return build, err
	}
	if len(steps) == 0 {
		return build, errors.New(
			"google interactions: no supported input steps remain after conversion",
		)
	}

	build.Interaction = interactions.CreateModelInteraction{
		Model:      interactions.Model(req.ModelParam.Name),
		Input:      new(interactions.NewInteractionsInput(steps)),
		Store:      new(false),
		Background: new(false),
		Stream:     new(stream),
	}

	if systemPrompt := strings.TrimSpace(req.ModelParam.SystemPrompt); systemPrompt != "" {
		build.Interaction.SystemInstruction = new(systemPrompt)
	}

	config := &interactions.GenerationConfig{}
	configured := false

	if req.ModelParam.MaxOutputLength > 0 {
		config.MaxOutputTokens = new(req.ModelParam.MaxOutputLength)
		configured = true
	}
	if len(req.ModelParam.StopSequences) > 0 {
		config.StopSequences = append([]string(nil), req.ModelParam.StopSequences...)
		configured = true
	}
	if req.ModelParam.Temperature != nil {
		build.Warnings = appendGoogleInteractionsWarning(
			build.Warnings,
			"temperature_dropped_unsupported",
			"Interactions does not expose a supported normalized temperature control.",
		)
	}
	if req.ModelParam.AdditionalParametersRawJSON != nil &&
		strings.TrimSpace(*req.ModelParam.AdditionalParametersRawJSON) != "" {
		build.Warnings = appendGoogleInteractionsWarning(
			build.Warnings,
			"additionalParameters_dropped_unsupported",
			"additionalParametersRawJSON is not forwarded by the Interactions adapter.",
		)
	}

	thinkingConfigured, thinkingWarnings := applyGoogleInteractionsThinkingConfig(
		config,
		req.ModelParam.Reasoning,
	)
	if thinkingConfigured {
		configured = true
	}
	build.Warnings = append(build.Warnings, thinkingWarnings...)

	tools, nameMap, webSearchChoiceID, toolConfigured, toolWarnings, err := buildGoogleInteractionsTools(
		req.ToolChoices,
		req.ToolPolicy,
		config,
	)
	if err != nil {
		return build, err
	}
	build.ToolChoiceNameMap = nameMap
	build.WebSearchChoiceID = webSearchChoiceID
	build.Warnings = append(build.Warnings, toolWarnings...)

	if len(tools) > 0 {
		build.Interaction.Tools = tools
	}
	if toolConfigured {
		configured = true
	}
	if configured {
		build.Interaction.GenerationConfig = config
	}

	format, formatWarnings, err := googleInteractionsResponseFormat(
		req.ModelParam.OutputParam,
	)
	if err != nil {
		return build, err
	}
	build.Warnings = append(build.Warnings, formatWarnings...)
	build.Interaction.ResponseFormat = format

	return build, nil
}

func googleInteractionsResponseFormat(
	output *spec.OutputParam,
) (*interactions.CreateModelInteractionResponseFormat, []spec.Warning, error) {
	if output == nil || output.Format == nil {
		return nil, nil, nil
	}

	switch output.Format.Kind {
	case spec.OutputFormatKindText:
		// Text is the SDK and API default. Do not send response_format or
		// response_mime_type when the caller requests ordinary text output.
		return nil, nil, nil

	case spec.OutputFormatKindJSONSchema:
		jsonSchema := output.Format.JSONSchemaParam
		if jsonSchema == nil || len(jsonSchema.Schema) == 0 {
			return nil, nil, errors.New(
				"google interactions: JSON Schema output requires jsonSchemaParam.schema",
			)
		}

		format := interactions.NewCreateModelInteractionResponseFormat(
			interactions.NewResponseFormat(
				interactions.TextResponseFormat{
					MimeType: interactions.TextResponseFormatMimeTypeApplicationJSON.ToPointer(),
					Schema:   maps.Clone(jsonSchema.Schema),
				},
			),
		)

		var warnings []spec.Warning
		if jsonSchema.Name != "" ||
			jsonSchema.Description != "" ||
			jsonSchema.Strict {
			warnings = appendGoogleInteractionsWarning(
				warnings,
				"jsonSchema_metadata_not_forwarded",
				"Interactions receives the JSON Schema payload but not schema name, description, or strict metadata.",
			)
		}
		return &format, warnings, nil

	default:
		return nil, nil, fmt.Errorf(
			"google interactions: unsupported output format %q",
			output.Format.Kind,
		)
	}
}

func googleInteractionsInputSteps(
	inputs []spec.InputUnion,
	warnings *[]spec.Warning,
) ([]interactions.Step, error) {
	out := make([]interactions.Step, 0, len(inputs))

	for _, input := range inputs {
		switch input.Kind {
		case spec.InputKindInputMessage:
			if input.InputMessage == nil ||
				input.InputMessage.Role != spec.RoleUser {
				continue
			}

			content, err := googleInteractionsContents(
				input.InputMessage.Contents,
				warnings,
			)
			if err != nil {
				return nil, err
			}
			if len(content) > 0 {
				out = append(out, interactions.NewStep(
					interactions.UserInputStep{Content: content},
				))
			}

		case spec.InputKindOutputMessage:
			if input.OutputMessage == nil ||
				input.OutputMessage.Role != spec.RoleAssistant {
				continue
			}

			content, err := googleInteractionsContents(
				input.OutputMessage.Contents,
				warnings,
			)
			if err != nil {
				return nil, err
			}
			if len(content) > 0 {
				out = append(out, interactions.NewStep(
					interactions.ModelOutputStep{Content: content},
				))
			}

		case spec.InputKindReasoningMessage:
			reasoning := input.ReasoningMessage
			if reasoning == nil ||
				strings.TrimSpace(reasoning.Signature) == "" {
				continue
			}

			out = append(out, interactions.NewStep(
				interactions.ThoughtStep{
					Signature: googleInteractionsOptionalString(reasoning.Signature),
					Summary: googleInteractionsThoughtSummaryContents(
						reasoning.Summary,
					),
				},
			))

		case spec.InputKindFunctionToolCall, spec.InputKindCustomToolCall:
			call := input.FunctionToolCall
			if input.Kind == spec.InputKindCustomToolCall {
				call = input.CustomToolCall
			}
			if call != nil {
				step, err := googleInteractionsFunctionCallStep(call)
				if err != nil {
					return nil, err
				}
				if step != nil {
					out = append(out, *step)
				}
			}

		case spec.InputKindFunctionToolOutput, spec.InputKindCustomToolOutput:
			toolOutput := input.FunctionToolOutput
			if input.Kind == spec.InputKindCustomToolOutput {
				toolOutput = input.CustomToolOutput
			}
			if toolOutput != nil {
				step, err := googleInteractionsFunctionResultStep(toolOutput, warnings)
				if err != nil {
					return nil, err
				}
				if step != nil {
					out = append(out, *step)
				}
			}

		case spec.InputKindWebSearchToolCall:
			if input.WebSearchToolCall != nil {
				step, err := googleInteractionsWebSearchCallStep(
					input.WebSearchToolCall,
				)
				if err != nil {
					return nil, err
				}
				if step != nil {
					out = append(out, *step)
				}
			}

		case spec.InputKindWebSearchToolOutput:
			step, err := googleInteractionsWebSearchResultStep(
				input.WebSearchToolOutput,
				warnings,
			)
			if err != nil {
				return nil, err
			}
			if step != nil {
				out = append(out, *step)
			}
		}
	}

	return out, nil
}

func googleInteractionsContents(
	items []spec.InputOutputContentItemUnion,
	warnings *[]spec.Warning,
) ([]interactions.Content, error) {
	out := make([]interactions.Content, 0, len(items))

	for _, item := range items {
		switch item.Kind {
		case spec.ContentItemKindText:
			if item.TextItem == nil ||
				strings.TrimSpace(item.TextItem.Text) == "" {
				continue
			}
			out = append(out, interactions.NewContent(
				interactions.TextContent{Text: item.TextItem.Text},
			))

		case spec.ContentItemKindImage:
			content, ok, err := googleInteractionsImageContent(item.ImageItem)
			if err != nil {
				return nil, err
			}
			if ok {
				out = append(out, content)
			} else {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"image_dropped_empty",
					"An image content item without data or URI was dropped.",
				)
			}

		case spec.ContentItemKindFile:
			if item.FileItem != nil &&
				strings.TrimSpace(item.FileItem.AdditionalContext) != "" {
				out = append(out, interactions.NewContent(
					interactions.TextContent{
						Text: item.FileItem.AdditionalContext,
					},
				))
			}
			if item.FileItem != nil && item.FileItem.CitationConfig != nil {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"fileCitationConfig_dropped_unsupported",
					"File citation configuration is not mapped by the Interactions adapter.",
				)
			}

			content, ok, err := googleInteractionsDocumentContent(item.FileItem)
			if err != nil {
				return nil, err
			}
			if ok {
				out = append(out, content)
			} else {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"file_dropped_empty",
					"A file content item without data or URI was dropped.",
				)
			}
		default:
		}
	}

	return out, nil
}

func googleInteractionsImageContent(
	image *spec.ContentItemImage,
) (interactions.Content, bool, error) {
	if image == nil {
		return interactions.Content{}, false, nil
	}

	data := strings.TrimSpace(image.ImageData)
	uri := strings.TrimSpace(image.ImageURL)
	if data == "" && uri == "" {
		return interactions.Content{}, false, nil
	}

	mimeType := strings.ToLower(strings.TrimSpace(image.ImageMIME))
	if mimeType == "" {
		mimeType = spec.DefaultImageDataMIME
	}
	if !googleInteractionsImageMIMESupported(mimeType) {
		return interactions.Content{}, false, fmt.Errorf(
			"google interactions: unsupported image MIME type %q",
			mimeType,
		)
	}

	content := interactions.ImageContent{
		MimeType: interactions.ImageContentMimeType(mimeType).ToPointer(),
	}
	switch image.Detail {
	case spec.ImageDetailLow:
		content.Resolution = interactions.MediaResolutionLow.ToPointer()
	case spec.ImageDetailHigh:
		content.Resolution = interactions.MediaResolutionHigh.ToPointer()
	default:
	}

	if data != "" {
		content.Data = new(data)
		return interactions.NewContent(content), true, nil
	}
	if uri != "" {
		content.URI = new(uri)
		return interactions.NewContent(content), true, nil
	}

	return interactions.Content{}, false, nil
}

func googleInteractionsDocumentContent(
	file *spec.ContentItemFile,
) (interactions.Content, bool, error) {
	if file == nil {
		return interactions.Content{}, false, nil
	}

	data := strings.TrimSpace(file.FileData)
	uri := strings.TrimSpace(file.FileURL)
	if data == "" && uri == "" {
		return interactions.Content{}, false, nil
	}

	mimeType := strings.ToLower(strings.TrimSpace(file.FileMIME))
	if !googleInteractionsDocumentMIMESupported(mimeType) {
		return interactions.Content{}, false, fmt.Errorf(
			"google interactions: unsupported document MIME type %q",
			mimeType,
		)
	}

	content := interactions.DocumentContent{
		MimeType: interactions.DocumentContentMimeType(mimeType).ToPointer(),
	}
	if data != "" {
		content.Data = new(data)
		return interactions.NewContent(content), true, nil
	}
	if uri != "" {
		content.URI = new(uri)
		return interactions.NewContent(content), true, nil
	}

	return interactions.Content{}, false, nil
}

func googleInteractionsFunctionCallStep(
	call *spec.ToolCall,
) (*interactions.Step, error) {
	callID := strings.TrimSpace(call.CallID)
	if callID == "" {
		callID = strings.TrimSpace(call.ID)
	}
	name := strings.TrimSpace(call.Name)
	if callID == "" || name == "" {
		return nil, errors.New(
			"google interactions: function call requires callID and name",
		)
	}

	arguments := map[string]any{}
	if raw := strings.TrimSpace(call.Arguments); raw != "" {
		if err := json.Unmarshal([]byte(raw), &arguments); err != nil ||
			arguments == nil {
			return nil, fmt.Errorf(
				"google interactions: function call %q requires JSON object arguments",
				call.Name,
			)
		}
	}

	step := interactions.NewStep(interactions.FunctionCallStep{
		ID:        callID,
		Name:      name,
		Arguments: arguments,
	})
	return &step, nil
}

func googleInteractionsFunctionResultStep(
	output *spec.ToolOutput,
	warnings *[]spec.Warning,
) (*interactions.Step, error) {
	if output == nil {
		return nil, errors.New(
			"google interactions: nil function result",
		)
	}

	callID := strings.TrimSpace(output.CallID)
	if callID == "" {
		callID = strings.TrimSpace(output.ID)
	}
	if callID == "" {
		return nil, errors.New(
			"google interactions: function result requires callID",
		)
	}

	subcontents, err := googleInteractionsFunctionResultSubcontents(
		output.Contents,
		warnings,
	)
	if err != nil {
		return nil, err
	}

	var resultUnion interactions.FunctionResultStepResultUnion
	//nolint:gocritic // If else is ok.
	if len(subcontents) == 1 && subcontents[0].TextContent != nil {
		text := ""
		if subcontents[0].TextContent.Text != "" {
			text = subcontents[0].TextContent.Text
		}
		resultUnion = interactions.NewFunctionResultStepResultUnion(text)
	} else if len(subcontents) == 0 {
		resultUnion = interactions.NewFunctionResultStepResultUnion("")
	} else {
		resultUnion = interactions.NewFunctionResultStepResultUnion(subcontents)
	}

	step := interactions.NewStep(interactions.FunctionResultStep{
		CallID:  callID,
		Name:    googleInteractionsOptionalString(strings.TrimSpace(output.Name)),
		Result:  resultUnion,
		IsError: new(output.IsError),
	})
	return &step, nil
}

func googleInteractionsFunctionResultSubcontents(
	items []spec.ToolOutputItemUnion,
	warnings *[]spec.Warning,
) ([]interactions.FunctionResultSubcontent, error) {
	out := make([]interactions.FunctionResultSubcontent, 0, len(items))
	for _, item := range items {
		switch item.Kind {
		case spec.ContentItemKindText:
			if item.TextItem == nil ||
				strings.TrimSpace(item.TextItem.Text) == "" {
				continue
			}
			out = append(out, interactions.NewFunctionResultSubcontent(
				interactions.TextContent{Text: item.TextItem.Text},
			))
		case spec.ContentItemKindImage:
			sub, ok, err := googleInteractionsFunctionImageSubcontent(
				item.ImageItem,
			)
			if err != nil {
				return nil, err
			}
			if ok {
				out = append(out, sub)
			} else if warnings != nil {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"image_dropped_empty",
					"An image tool result without data or URI was dropped.",
				)
			}
		case spec.ContentItemKindFile:
			if item.FileItem != nil &&
				strings.TrimSpace(item.FileItem.AdditionalContext) != "" {
				out = append(out, interactions.NewFunctionResultSubcontent(
					interactions.TextContent{
						Text: item.FileItem.AdditionalContext,
					},
				))
			}
			if item.FileItem != nil && item.FileItem.CitationConfig != nil {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"fileCitationConfig_dropped_unsupported",
					"File citation configuration is not mapped by the Interactions adapter.",
				)
			}
			sub, ok, err := googleInteractionsFunctionDocumentSubcontent(
				item.FileItem,
			)
			if err != nil {
				return nil, err
			}
			if ok {
				out = append(out, sub)
			} else if warnings != nil {
				*warnings = appendGoogleInteractionsWarning(
					*warnings,
					"file_dropped_empty",
					"A file tool result without data or URI was dropped.",
				)
			}
		default:
			continue
		}
	}
	return out, nil
}

func googleInteractionsFunctionImageSubcontent(
	image *spec.ContentItemImage,
) (interactions.FunctionResultSubcontent, bool, error) {
	if image == nil {
		return interactions.FunctionResultSubcontent{}, false, nil
	}

	data := strings.TrimSpace(image.ImageData)
	uri := strings.TrimSpace(image.ImageURL)
	if data == "" && uri == "" {
		return interactions.FunctionResultSubcontent{}, false, nil
	}

	mimeType := strings.ToLower(strings.TrimSpace(image.ImageMIME))
	if mimeType == "" {
		mimeType = spec.DefaultImageDataMIME
	}
	if !googleInteractionsImageMIMESupported(mimeType) {
		return interactions.FunctionResultSubcontent{}, false, fmt.Errorf(
			"google interactions: unsupported image MIME type %q",
			mimeType,
		)
	}

	content := interactions.ImageContent{
		MimeType: interactions.ImageContentMimeType(mimeType).ToPointer(),
	}
	switch image.Detail {
	case spec.ImageDetailLow:
		content.Resolution = interactions.MediaResolutionLow.ToPointer()
	case spec.ImageDetailHigh:
		content.Resolution = interactions.MediaResolutionHigh.ToPointer()
	default:
	}
	if data != "" {
		content.Data = new(data)
		return interactions.NewFunctionResultSubcontent(content), true, nil
	}
	if uri != "" {
		content.URI = new(uri)
		return interactions.NewFunctionResultSubcontent(content), true, nil
	}
	return interactions.FunctionResultSubcontent{}, false, nil
}

func googleInteractionsWebSearchCallStep(
	call *spec.ToolCall,
) (*interactions.Step, error) {
	callID := strings.TrimSpace(call.CallID)
	if callID == "" {
		callID = strings.TrimSpace(call.ID)
	}
	if callID == "" {
		return nil, errors.New(
			"google interactions: Google Search call requires callID",
		)
	}

	queries := make([]string, 0, len(call.WebSearchToolCallItems))
	for _, item := range call.WebSearchToolCallItems {
		if item.Kind == spec.WebSearchToolCallKindSearch &&
			item.SearchItem != nil &&
			strings.TrimSpace(item.SearchItem.Query) != "" {
			queries = append(queries, item.SearchItem.Query)
		}
	}

	if len(queries) == 0 {
		//nolint:nilnil // No queries found.
		return nil, nil
	}

	step := interactions.NewStep(interactions.GoogleSearchCallStep{
		ID: callID,
		Arguments: interactions.GoogleSearchCallArguments{
			Queries: queries,
		},
		Signature: googleInteractionsOptionalString(call.Signature),
	})
	return &step, nil
}

func googleInteractionsWebSearchResultStep(
	output *spec.ToolOutput,
	warnings *[]spec.Warning,
) (*interactions.Step, error) {
	if output == nil {
		//nolint:nilnil // No output found.
		return nil, nil
	}

	callID := strings.TrimSpace(output.CallID)
	if callID == "" {
		callID = strings.TrimSpace(output.ID)
	}
	if callID == "" {
		//nolint:nilnil // No output found.
		return nil, nil
	}

	result := make(
		[]interactions.GoogleSearchResult,
		0,
		len(output.WebSearchToolOutputItems),
	)
	for _, item := range output.WebSearchToolOutputItems {
		if item.Kind != spec.WebSearchToolOutputKindSearch ||
			item.SearchItem == nil {
			continue
		}

		if item.SearchItem.URL != "" || item.SearchItem.Title != "" {
			*warnings = appendGoogleInteractionsWarning(
				*warnings,
				"webSearch_result_metadata_dropped",
				"Interactions Google Search history replays rendered search suggestions without URL or title metadata.",
			)
		}
		if text := strings.TrimSpace(item.SearchItem.RenderedContent); text != "" {
			result = append(result, interactions.GoogleSearchResult{
				SearchSuggestions: new(text),
			})
		}
	}

	step := interactions.NewStep(interactions.GoogleSearchResultStep{
		CallID:    callID,
		Result:    result,
		IsError:   new(output.IsError),
		Signature: googleInteractionsOptionalString(output.Signature),
	})
	return &step, nil
}

func googleInteractionsFunctionDocumentSubcontent(
	file *spec.ContentItemFile,
) (interactions.FunctionResultSubcontent, bool, error) {
	if file == nil {
		return interactions.FunctionResultSubcontent{}, false, nil
	}

	data := strings.TrimSpace(file.FileData)
	uri := strings.TrimSpace(file.FileURL)
	if data == "" && uri == "" {
		return interactions.FunctionResultSubcontent{}, false, nil
	}

	mimeType := strings.ToLower(strings.TrimSpace(file.FileMIME))
	if !googleInteractionsDocumentMIMESupported(mimeType) {
		return interactions.FunctionResultSubcontent{}, false, fmt.Errorf(
			"google interactions: unsupported document MIME type %q",
			mimeType,
		)
	}

	content := interactions.DocumentContent{
		MimeType: interactions.DocumentContentMimeType(mimeType).ToPointer(),
	}
	if data != "" {
		content.Data = new(data)
		return interactions.NewFunctionResultSubcontent(content), true, nil
	}
	if uri != "" {
		content.URI = new(uri)
		return interactions.NewFunctionResultSubcontent(content), true, nil
	}
	return interactions.FunctionResultSubcontent{}, false, nil
}

func googleInteractionsImageMIMESupported(mimeType string) bool {
	switch mimeType {
	case "image/png",
		"image/jpeg",
		"image/webp",
		"image/heic",
		"image/heif",
		"image/gif",
		"image/bmp",
		"image/tiff":
		return true
	default:
		return false
	}
}

func googleInteractionsDocumentMIMESupported(mimeType string) bool {
	switch mimeType {
	case "application/pdf",
		"text/csv",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"application/vnd.ms-powerpoint",
		"application/vnd.oasis.opendocument.presentation",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/msword",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.template",
		"application/vnd.oasis.opendocument.text",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-excel",
		"application/vnd.oasis.opendocument.spreadsheet":
		return true
	default:
		return false
	}
}
