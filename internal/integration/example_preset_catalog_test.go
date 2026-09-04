package integration

import (
	"context"
	"fmt"

	"github.com/flexigpt/inference-go"
	"github.com/flexigpt/inference-go/modelpreset"
	"github.com/flexigpt/inference-go/spec"
)

func Example_presetProviderSetup() {
	ctx := context.Background()

	providerSet, err := inference.NewProviderSetAPI()
	if err != nil {
		fmt.Println("provider set error:", err)
		return
	}

	provider, err := modelpreset.Provider(modelpreset.ProviderOpenAIResponses)
	if err != nil {
		fmt.Println("provider lookup error:", err)
		return
	}

	model, err := modelpreset.Model(
		modelpreset.ProviderOpenAIResponses,
		modelpreset.PresetGPT56Luna,
	)
	if err != nil {
		fmt.Println("model lookup error:", err)
		return
	}

	if _, err := providerSet.AddProviderFromPreset(ctx, provider.Name, provider); err != nil {
		fmt.Println("provider registration error:", err)
		return
	}

	completionKey := string(model.ID)
	resolver, err := providerSet.NewPresetCapabilityResolver(
		ctx,
		provider.Name,
		provider,
		model,
		completionKey,
	)
	if err != nil {
		fmt.Println("resolver error:", err)
		return
	}

	capabilities, err := resolver.ResolveModelCapabilities(
		ctx,
		spec.ResolveModelCapabilitiesRequest{
			ProviderSDKType: provider.SDKType,
			ModelName:       model.Name,
			CompletionKey:   completionKey,
		},
	)
	if err != nil {
		fmt.Println("capability resolution error:", err)
		return
	}

	fmt.Printf(
		"provider=%s model=%s reasoning=%t\n",
		provider.Name,
		model.Name,
		capabilities.ReasoningCapabilities != nil &&
			capabilities.ReasoningCapabilities.SupportsReasoningConfig,
	)

	// Output:
	// provider=openairesponses model=gpt-5.6-luna reasoning=true
}
