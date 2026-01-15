package sdkutil

import "github.com/flexigpt/inference-go/spec"

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
