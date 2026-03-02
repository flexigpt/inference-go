package spec

import (
	"context"
	"net/http"
)

// StreamContentKind enumerates the kinds of streaming events that can be delivered while a completion is in progress.
type StreamContentKind string

const (
	StreamContentKindText     StreamContentKind = "text"
	StreamContentKindThinking StreamContentKind = "thinking"
)

type StreamTextChunk struct {
	Text string `json:"text"`
}

type StreamThinkingChunk struct {
	Text string `json:"text"`
}

type StreamEvent struct {
	Kind StreamContentKind `json:"kind"`

	// Optional metadata to help consumers correlate events across models/providers.
	Provider      ProviderName `json:"provider,omitempty"`
	Model         ModelName    `json:"model,omitempty"`
	CompletionKey string       `json:"completionKey,omitempty"`

	// Exactly one of the below will be non-nil depending on Kind.
	Text     *StreamTextChunk     `json:"text,omitempty"`
	Thinking *StreamThinkingChunk `json:"thinking,omitempty"`
}

// StreamConfig controls low-level behavior of streaming delivery. All fields are optional; zero values mean "use
// library defaults".
type StreamConfig struct {
	// FlushIntervalMillis is the maximum delay between flushes of buffered stream data to the StreamHandler.
	FlushIntervalMillis int `json:"flushIntervalMillis,omitempty"`
	// FlushChunkSize is the approximate target size (in bytes/characters) for chunks passed to the StreamHandler.
	FlushChunkSize int `json:"flushChunkSize,omitempty"`
}

type StreamHandler func(event StreamEvent) error

// FetchCompletionOptions controls optional behaviors for FetchCompletion.
// A nil pointer is treated the same as &FetchCompletionOptions{}.
type FetchCompletionOptions struct {
	// CompletionKey is an opaque, runtime-only key that is passed in runtime operations
	// It can be used by the capability resolver to map arbitrary user/model identifiers to capability profiles.
	CompletionKey string `json:"-"`

	// StreamHandler, if non-nil, is invoked with incremental streaming events
	// when ModelParam.Stream is true. Returning a non-nil error will stop
	// streaming early and propagate that error back to the caller.
	StreamHandler StreamHandler `json:"-"`
	StreamConfig  *StreamConfig `json:"streamConfig,omitempty"`

	// CapabilityResolver, if non-nil, is used to resolve model capabilities for request validation/normalization.
	// Else, inbuilt SDK capabilities are enforced.
	CapabilityResolver ModelCapabilityResolver `json:"-"`
}

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FetchCompletionResponse struct {
	Outputs      []OutputUnion `json:"outputs,omitempty"`
	Usage        *Usage        `json:"usage,omitempty"`
	Error        *Error        `json:"error,omitempty"`
	Warnings     []Warning     `json:"warnings,omitempty"`
	DebugDetails any           `json:"debugDetails,omitempty"`
}

type FetchCompletionRequest struct {
	ModelParam ModelParam   `json:"modelParam"`
	Inputs     []InputUnion `json:"inputs"`

	// ToolPolicy - optional control on how (or whether) the model may use the provided ToolChoices.
	ToolPolicy  *ToolPolicy  `json:"toolPolicy,omitempty"`
	ToolChoices []ToolChoice `json:"toolChoices,omitempty"`
}

type CompletionSpanStart struct {
	Provider ProviderName
	Model    ModelName

	// Original request and options for this completion.
	// These may be nil and MUST be treated as read-only.
	Request *FetchCompletionRequest
	Options *FetchCompletionOptions
}

type CompletionSpanEnd struct {
	// Raw SDK response (e.g. *responses.Response for OpenAI). May be nil.
	ProviderResponse any

	// Error from the provider/stream if any.
	Err error

	// Normalized response object that the caller is about to return.
	// May be nil. MUST be treated as read-only.
	Response *FetchCompletionResponse
}

// CompletionSpan is the per-request handle. Only provider code sees this; external callers never construct it.
type CompletionSpan interface {
	// End is called exactly once, just before FetchCompletion returns.
	//
	// It can:
	//   - pull any per-request state from the context,
	//   - inspect raw/normalized responses and errors,
	//   - return arbitrary data to attach to Response.DebugDetails.
	//
	// Returning nil means "no debug details for this call".
	End(info *CompletionSpanEnd) any
}

// CompletionDebugger is the long-lived "client" object for a provider.
//
// Provider code (e.g., OpenAIResponsesAPI) owns one CompletionDebugger.
// Callers don't touch this directly.
type CompletionDebugger interface {
	// HTTPClient is called once when the provider initializes its SDK client.
	//
	// The debugger can:
	//   - wrap base (change Transport),
	//   - ignore base and create a new client,
	//   - or return nil to say "use provider's default client".
	HTTPClient(base *http.Client) *http.Client

	// StartSpan is called at the beginning of FetchCompletion.
	//
	// It can:
	//   - inspect the request to decide whether to debug,
	//   - attach per-request state to the context (e.g., via context keys),
	//   - return a span handle that will be ended when the call finishes.
	//
	// If span is nil, no debugging will be performed for this call.
	StartSpan(
		ctx context.Context,
		info *CompletionSpanStart,
	) (ctxWithSpan context.Context, span CompletionSpan)
}
