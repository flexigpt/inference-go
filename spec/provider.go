package spec

import "time"

const (
	DefaultAuthorizationHeaderKey = "Authorization"
	DefaultAPITimeout             = 300 * time.Second

	DefaultFileDataMIME    = "application/octet-stream"
	DefaultImageDataMIME   = "image/png"
	DefaultApplicationJSON = "application/json"

	DefaultContentTypeHeaderKey = "content-type"
	DefaultAcceptHeaderKey      = "accept"
	DefaultContentTypeHeader    = DefaultApplicationJSON

	DefaultAnthropicOrigin                 = "https://api.anthropic.com"
	DefaultAnthropicChatCompletionPrefix   = "/v1/messages"
	DefaultAnthropicAuthorizationHeaderKey = "x-api-key"

	DefaultOpenAIOrigin                = "https://api.openai.com"
	DefaultOpenAIChatCompletionsPrefix = "/v1/chat/completions"
	DefaultOpenAIResponsesPrefix       = "/v1/responses"

	DefaultGoogleGenerateContentOrigin = "https://generativelanguage.googleapis.com"
	DefaultGoogleGenerateContentPrefix = "/"
	//nolint:gosec // APIKeyHeaderKey is the key string and not the key itself.
	DefaultGoogleGenerateContentAPIKeyHeaderKey = "x-goog-api-key"
)

var DefaultBaseHeaders = map[string]string{DefaultContentTypeHeaderKey: DefaultContentTypeHeader}

const (
	ProviderSDKTypeAnthropic             ProviderSDKType = "providerSDKTypeAnthropicMessages"
	ProviderSDKTypeOpenAIChatCompletions ProviderSDKType = "providerSDKTypeOpenAIChatCompletions"
	ProviderSDKTypeOpenAIResponses       ProviderSDKType = "providerSDKTypeOpenAIResponses"
	ProviderSDKTypeGoogleGenerateContent ProviderSDKType = "providerSDKTypeGoogleGenerateContent"
)

// ProviderParam represents information about a provider.
type ProviderParam struct {
	Name                     ProviderName      `json:"name"`
	SDKType                  ProviderSDKType   `json:"sdkType"`
	APIKey                   string            `json:"apiKey"`
	Origin                   string            `json:"origin"`
	ChatCompletionPathPrefix string            `json:"chatCompletionPathPrefix"`
	APIKeyHeaderKey          string            `json:"apiKeyHeaderKey"`
	DefaultHeaders           map[string]string `json:"defaultHeaders"`
}
