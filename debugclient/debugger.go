package debugclient

import (
	"context"
	"net/http"
	"strings"

	"github.com/ppipada/inference-go/internal/sdkutil"
	"github.com/ppipada/inference-go/spec"
)

// DebugConfig controls how HTTP debug information is captured and redacted.
//
// The zero value corresponds to the default behavior:
//
//   - debugging enabled
//   - request/response bodies captured
//   - content (LLM text, large/base64 blobs) stripped/scrubbed
//   - no slog logging
type DebugConfig struct {
	// Disable turns off all debugging when true.
	Disable bool `json:"disable,omitempty"`

	// DisableRequestBody prevents capturing request bodies when true.
	DisableRequestBody bool `json:"disableRequestBody,omitempty"`

	// DisableResponseBody prevents capturing response bodies when true.
	DisableResponseBody bool `json:"disableResponseBody,omitempty"`

	// DisableContentStripping, if true, leaves LLM text content and large/base64
	// payloads untouched. When false (default), scrubbers remove user/assistant
	// text and large/base64 blobs while preserving other metadata.
	DisableContentStripping bool `json:"disableContentStripping,omitempty"`

	// LogToSlog logs HTTP request/response details at debug level when true.
	LogToSlog bool `json:"logToSlog,omitempty"`
}

// HTTPCompletionDebugger implements spec.CompletionDebugger using the HTTP
// instrumentation in this package. It produces an opaque debug payload
// (HTTPDebugState) suitable for attachment to FetchCompletionResponse.DebugDetails.
type HTTPCompletionDebugger struct {
	config DebugConfig
}

// NewHTTPCompletionDebugger constructs a CompletionDebugger that instruments
// HTTP traffic using an internal RoundTripper and produces a scrubbed debug
// blob from HTTP-level data and the raw provider response.
//
// Config may be nil; in that case DebugConfig{} (defaults) is used.
func NewHTTPCompletionDebugger(config *DebugConfig) spec.CompletionDebugger {
	var c DebugConfig
	if config != nil {
		c = *config
	}
	return &HTTPCompletionDebugger{config: c}
}

// HTTPClient implements spec.CompletionDebugger.HTTPClient.
func (d *HTTPCompletionDebugger) HTTPClient(base *http.Client) *http.Client {
	if d.config.Disable {
		return base
	}

	if base == nil {
		base = &http.Client{Transport: http.DefaultTransport}
	}
	rt := base.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}

	clone := *base
	clone.Transport = &logTransport{
		base: rt,
		cfg:  d.config,
	}
	return &clone
}

type httpSpan struct {
	cfg  DebugConfig
	ctx  context.Context
	info *spec.CompletionSpanStart
}

// StartSpan implements spec.CompletionDebugger.StartSpan.
func (d *HTTPCompletionDebugger) StartSpan(
	ctx context.Context,
	info *spec.CompletionSpanStart,
) (context.Context, spec.CompletionSpan) {
	if d.config.Disable {
		return ctx, nil
	}

	ctx = withHTTPDebugState(ctx)
	span := &httpSpan{
		cfg:  d.config,
		ctx:  ctx,
		info: info,
	}
	return ctx, span
}

// APIRequestDetails describes a single HTTP request captured by the debugger.
type APIRequestDetails struct {
	URL         *string        `json:"url,omitempty"`
	Method      *string        `json:"method,omitempty"`
	Headers     map[string]any `json:"headers,omitempty"`
	Params      map[string]any `json:"params,omitempty"`
	Data        any            `json:"data,omitempty"`
	Timeout     *int           `json:"timeout,omitempty"`
	CurlCommand *string        `json:"curlCommand,omitempty"`
}

// APIResponseDetails describes a single HTTP response captured by the debugger.
type APIResponseDetails struct {
	Data    any            `json:"data,omitempty"`
	Status  int            `json:"status"`
	Headers map[string]any `json:"headers,omitempty"`
}

// APIErrorDetails summarizes an HTTP-level error.
type APIErrorDetails struct {
	Message         string              `json:"message"`
	RequestDetails  *APIRequestDetails  `json:"requestDetails,omitempty"`
	ResponseDetails *APIResponseDetails `json:"responseDetails,omitempty"`
}

// HTTPDebugState is the full HTTP- and provider-level debug payload returned by HTTPCompletionDebugger.
type HTTPDebugState struct {
	RequestDetails  *APIRequestDetails  `json:"requestDetails,omitempty"`
	ResponseDetails *APIResponseDetails `json:"responseDetails,omitempty"`
	ErrorDetails    *APIErrorDetails    `json:"errorDetails,omitempty"`

	// ProviderResponse holds a scrubbed form of the raw provider SDK response
	// (e.g. *responses.Response for OpenAI), if available.
	ProviderResponse any `json:"providerResponse,omitempty"`
}

// End implements spec.CompletionSpan.End.
func (s *httpSpan) End(end *spec.CompletionSpanEnd) any {
	defer sdkutil.Recover("debugclient.httpSpan.End panic")

	if s.cfg.Disable {
		return nil
	}

	state, _ := httpDebugStateFromContext(s.ctx)
	if state == nil {
		state = &HTTPDebugState{}
	}

	// Attach scrubbed provider response, if any.
	if end.ProviderResponse != nil {
		if m, err := structToMap(end.ProviderResponse); err == nil {
			strip := !s.cfg.DisableContentStripping
			state.ProviderResponse = scrubAnyForDebug(m, strip)
		}
	}

	// Compose error message from HTTP-level error + provider error.
	var msgParts []string
	if state.ErrorDetails != nil {
		if msg := strings.TrimSpace(state.ErrorDetails.Message); msg != "" {
			msgParts = append(msgParts, msg)
		}
	}
	if end.Err != nil {
		msgParts = append(msgParts, end.Err.Error())
	}

	if len(msgParts) > 0 {
		if state.ErrorDetails == nil {
			state.ErrorDetails = &APIErrorDetails{}
		}
		state.ErrorDetails.Message = strings.Join(msgParts, "; ")
	}

	return state
}
