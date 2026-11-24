package modelpreset

import (
	"errors"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

func TestGemini4UsesInteractionsProvider(t *testing.T) {
	provider, err := Provider(ProviderGoogleGeminiInteractions)
	if err != nil {
		t.Fatalf("Provider(%q): %v", ProviderGoogleGeminiInteractions, err)
	}

	if provider.SDKType != spec.ProviderSDKTypeGoogleInteractions {
		t.Fatalf(
			"sdk type got %q want %q",
			provider.SDKType,
			spec.ProviderSDKTypeGoogleInteractions,
		)
	}

	for _, id := range []ModelPresetID{
		PresetGemini4Neon,
		PresetGemini4Carbon,
		PresetGemini4Hydrogen,
		PresetGemini4Argon,
	} {
		if _, err := Model(ProviderGoogleGeminiInteractions, id); err != nil {
			t.Errorf("Interactions Model(%q): %v", id, err)
		}

		if _, err := Model(ProviderGoogleGemini, id); !errors.Is(err, ErrModelNotFound) {
			t.Errorf(
				"Generate Content Model(%q) error got %v want ErrModelNotFound",
				id,
				err,
			)
		}
	}
}

func TestLegacyGoogleGeminiProviderRetainsGenerateContentModels(t *testing.T) {
	provider, err := Provider(ProviderGoogleGemini)
	if err != nil {
		t.Fatal(err)
	}

	if provider.SDKType != spec.ProviderSDKTypeGoogleGenerateContent {
		t.Fatalf(
			"sdk type got %q want %q",
			provider.SDKType,
			spec.ProviderSDKTypeGoogleGenerateContent,
		)
	}
	if _, err := Model(ProviderGoogleGemini, PresetGemini25Flash); err != nil {
		t.Fatalf("Gemini 2.5 Flash must remain on Generate Content: %v", err)
	}
}
