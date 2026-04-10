package debugclient

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
