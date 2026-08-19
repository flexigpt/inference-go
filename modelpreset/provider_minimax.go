package modelpreset

import "github.com/flexigpt/inference-go/spec"

const (
	ModelNameOpenRouterMiniMaxM3      spec.ModelName = "minimax/minimax-m3"
	ModelNameOpenRouterMiniMaxM27     spec.ModelName = "minimax/minimax-m2.7"
	ModelNameOpenRouterMiniMaxM25Free spec.ModelName = "minimax/minimax-m2.5:free"

	ModelNameMiniMaxM27FireworksAI spec.ModelName = "MiniMaxAI/MiniMax-M2.7:fireworks-ai"
	ModelNameMiniMaxM25Novita      spec.ModelName = "MiniMaxAI/MiniMax-M2.5:novita"
)

const (
	DisplayNameMiniMaxM25     = "MiniMax M2.5"
	DisplayNameMiniMaxM25Free = "MiniMax M2.5 Free"
	DisplayNameMiniMaxM27     = "MiniMax M2.7"
	DisplayNameMiniMaxM3      = "MiniMax M3"

	DisplayNameMiniMaxM25Novita      = "MiniMax M2.5 (Novita)"
	DisplayNameMiniMaxM27FireworksAI = "MiniMax M2.7 (Fireworks AI)"
)

const (
	PresetMiniMaxM25     ModelPresetID = "minimaxm25"
	PresetMiniMaxM25Free ModelPresetID = "minimaxm25free"
	PresetMiniMaxM27     ModelPresetID = "minimaxm27"
	PresetMiniMaxM3      ModelPresetID = "minimaxM3"

	PresetMiniMaxM27FireworksAI ModelPresetID = "minimaxm27FireworksAI"
	PresetMiniMaxM25Novita      ModelPresetID = "minimaxm25Novita"
)
