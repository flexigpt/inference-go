package modelpreset

import (
	"slices"
	"strings"
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

// TestCatalogModelMembershipIsExhaustive is an explicit catalog manifest.
//
// It intentionally fails when a model is added, removed, or renamed without
// updating test coverage. This prevents catalog additions from being silently
// omitted from product-level review and integration coverage.
func TestCatalogModelMembershipIsExhaustive(t *testing.T) {
	expected := map[spec.ProviderName]string{
		ProviderAnthropic: "fable5 fable51 opus5 opus48 opus47 opus46 opus45 opus41 " +
			"sonnet5 sonnet46 sonnet45 sonnet4 haiku45",
		ProviderDeepSeek: "deepseekv4flash deepseekv4pro",
		ProviderGoogleGemini: "gemini38Flash gemini37Flash gemini36Flash gemini35Flash " +
			"gemini35FlashLite gemini3Flash gemini31Pro gemini31FlashLite gemini25Flash gemini25FlashLite",
		ProviderHuggingFace: "deepseekv4flashFireworksAI deepseekv4proFireworksAI " +
			"glm47Cerebras glm51FireworksAI glm51fp8FireworksAI glm52FireworksAI glm52fp8ZAI glm5Novita " +
			"gptoss120bFireworksAI gptoss20bFireworksAI " +
			"kimik2instruct0905Novita kimik2instructNovita kimik2thinkingFeatherlessAI " +
			"mimov25proDeepInfra mimov2flashFeatherlessAI " +
			"minimaxm25Novita minimaxm27FireworksAI " +
			"nemotron3superBF16FeatherlessAI nemotron3ultraBF16DeepInfra nemotron3ultraNVFP4FireworksAI " +
			"ornith1035bfp8DeepInfra " +
			"qwen3coder30ba3bFireworksAI qwen3codernextNovita step35flashFeatherlessAI",
		ProviderLocalAI: "gemma426ba4b gptoss20b qwen3635ba3b deepseekr18b qwen3vl30ba3b " +
			"ministral314b qwen3coder30ba3b glm47flash30ba3b phi4reasoning14b devstral224b",
		ProviderLMStudio: "gemma426ba4b gptoss20b qwen3635ba3b qwen3627b deepseekr18b " +
			"qwen3vl30b ministral314b qwen3coder30ba3b glm47flash30ba3b devstral224b",
		ProviderLlamaCPP: "llama4Behemoth llama4Maverick llama4Scout qwen3635ba3b",
		ProviderMeta:     "museSpark13 museSpark13Contributor museSpark12 museSpark12Contributor museSpark11",
		ProviderMiniMax: "minimaxm2 minimaxm21 minimaxm21Highspeed minimaxm25 minimaxm25Highspeed " +
			"minimaxm27 minimaxm27Highspeed minimaxM3",
		ProviderMistral:  "mistralMedium35 mistralSmall4 mistralLarge3 devstral2",
		ProviderMoonshot: "kimiK3 kimiK27Code kimiK27CodeHighspeed kimiK26",
		ProviderOllama: "gemma426b gemma4e4b gptoss20b qwen3635b qwen3627b deepseekr18b " +
			"qwen3vl30b ministral314b qwen3coder30b phi4reasoning14b",
		ProviderOpenAIChat: "gpt41 gpt41Mini gpt4o gpt4oMini",
		ProviderOpenAIResponses: "gpt6Astra gpt56sol gpt56terra gpt56luna gpt55 gpt54 " +
			"gpt54mini gpt54nano gpt53Codex gpt52 gpt52Codex gpt51 gpt51Codex gpt51CodexMax gpt5Mini",
		ProviderOpenRouter: "deepseekv4flash xiaomiMiMoV25 tencentHy3Preview minimaxM3 zaiglm52 " +
			"deepseekv4pro step37Flash nvidiaNemotron3UltraFree poolsideLagunaM1Free xiaomiMiMoV25Pro " +
			"nvidiaNemotron3SuperFree kimiK26 qwen37Max zaiglm51 kimiK27Code qwen37Plus minimaxm27 minimaxm25free",
		ProviderQwen: "qwen38Max qwen3824tA95B qwen3827B qwen37Max qwen37Max20260608 " +
			"qwen37Max20260520 qwen3Max qwen3Max20260123 qwen37Plus qwen37Plus20260526 " +
			"qwen37Flash qwen37Flash20260715 qwen36Plus qwen36Plus20260402 qwen36Flash " +
			"qwen36Flash20260416 qwen3635bA3B qwen3627B qwen35Plus qwen35Plus20260420 " +
			"qwen35Plus20260215 qwen35Flash qwen35Flash20260223 qwen35397bA17B " +
			"qwen35122bA10B qwen3527B qwen3535bA3B qwenPlus qwenFlash qwen3CoderPlus " +
			"qwen3CoderFlash qwenPlusCharacter qwenFlashCharacter",
		ProviderSGLang: "gemma426ba4b gptoss20b qwen3635ba3b qwen3vl30ba3b deepseekr18b " +
			"ministral314b qwen3coder30ba3b glm47flash30ba3b phi4reasoning14b devstral224b",
		ProviderVLLM: "gemma426ba4b gptoss20b qwen3635ba3b qwen3vl30ba3b deepseekr18b " +
			"ministral314b qwen3coder30ba3b glm47flash30ba3b phi4reasoning14b devstral224b",
		ProviderXAI:    "grok46 grok45 grok43 grokBuild01 grok42Reasoning grok42NonReasoning",
		ProviderXiaomi: "mimov25 mimov25pro",
		ProviderZAI: "glm53 glm52 glm51 glm5 glm5Turbo glm47 glm47FlashX glm47Flash " +
			"glm46 glm45 glm45X glm45Air glm45AirX glm45Flash glm5VTurbo glm46V glm46VFlash " +
			"glm46VFlashX glm45V",
		ProviderZAICodingPlan: "glm53 glm5Turbo glm47",
	}

	catalog := DefaultCatalog()
	if len(catalog.Providers) != len(expected) {
		t.Fatalf(
			"provider manifest count got %d want %d",
			len(catalog.Providers),
			len(expected),
		)
	}

	providerNames := make([]spec.ProviderName, 0, len(expected))
	for providerName := range expected {
		providerNames = append(providerNames, providerName)
	}
	slices.Sort(providerNames)

	for _, providerName := range providerNames {
		t.Run(string(providerName), func(t *testing.T) {
			provider, ok := catalog.Providers[providerName]
			if !ok {
				t.Fatalf("provider %q is missing from DefaultCatalog", providerName)
			}

			want := modelPresetIDsFromManifest(t, expected[providerName])
			got := make([]ModelPresetID, 0, len(provider.ModelPresets))
			for modelID := range provider.ModelPresets {
				got = append(got, modelID)
			}
			slices.Sort(got)

			if !slices.Equal(got, want) {
				t.Fatalf(
					"%s model membership mismatch\n got: %#v\nwant: %#v",
					providerName,
					got,
					want,
				)
			}
		})
	}

	for providerName := range catalog.Providers {
		if _, ok := expected[providerName]; !ok {
			t.Fatalf("catalog provider %q is missing from expected manifest", providerName)
		}
	}
}

func modelPresetIDsFromManifest(t *testing.T, manifest string) []ModelPresetID {
	t.Helper()

	rawIDs := strings.Fields(manifest)
	out := make([]ModelPresetID, 0, len(rawIDs))
	seen := make(map[ModelPresetID]struct{}, len(rawIDs))

	for _, rawID := range rawIDs {
		id := ModelPresetID(rawID)
		if _, exists := seen[id]; exists {
			t.Fatalf("duplicate model preset ID %q in test manifest", id)
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}

	slices.Sort(out)
	return out
}
