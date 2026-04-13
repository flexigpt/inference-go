package spec

import "time"

const (
	DefaultAuthorizationHeaderKey = "Authorization"
	DefaultAPITimeout             = 300 * time.Second

	DefaultAnthropicOrigin                 = "https://api.anthropic.com"
	DefaultAnthropicChatCompletionPrefix   = "/v1/messages"
	DefaultAnthropicAuthorizationHeaderKey = "x-api-key"

	DefaultOpenAIOrigin                = "https://api.openai.com"
	DefaultOpenAIChatCompletionsPrefix = "/v1/chat/completions"

	DefaultGoogleGenerateContentOrigin = "https://generativelanguage.googleapis.com"

	DefaultFileDataMIME  = "application/octet-stream"
	DefaultImageDataMIME = "image/png"
)

var OpenAIChatCompletionsDefaultHeaders = map[string]string{"content-type": "application/json"}

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
