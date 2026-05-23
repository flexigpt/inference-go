package modelpreset

import (
	"github.com/flexigpt/inference-go/capabilityoverride"
	"github.com/flexigpt/inference-go/spec"
)

type ModelPresetID string

type ModelPreset struct {
	ID          ModelPresetID  `json:"id"`
	Name        spec.ModelName `json:"name"`
	DisplayName string         `json:"displayName"`

	// ModelParam is the default runtime request model configuration.
	// Callers should treat values returned by this package as immutable.
	ModelParam spec.ModelParam `json:"modelParam"`

	// CapabilitiesOverride is a runtime capability patch applied over provider/base SDK capabilities.
	// It is not the derived/effective capability profile.
	CapabilitiesOverride *capabilityoverride.ModelCapabilitiesOverride `json:"capabilitiesOverride,omitempty"`
}

type ProviderPreset struct {
	Name        spec.ProviderName    `json:"name"`
	DisplayName string               `json:"displayName"`
	SDKType     spec.ProviderSDKType `json:"sdkType"`

	Origin                   string            `json:"origin"`
	ChatCompletionPathPrefix string            `json:"chatCompletionPathPrefix"`
	APIKeyHeaderKey          string            `json:"apiKeyHeaderKey"`
	DefaultHeaders           map[string]string `json:"defaultHeaders,omitempty"`

	// CapabilitiesOverride is a provider-wide runtime capability patch.
	// Model preset overrides are applied after this.
	CapabilitiesOverride *capabilityoverride.ModelCapabilitiesOverride `json:"capabilitiesOverride,omitempty"`

	ModelPresets map[ModelPresetID]ModelPreset `json:"modelPresets"`
}

type Catalog struct {
	Providers map[spec.ProviderName]ProviderPreset `json:"providers"`
}
