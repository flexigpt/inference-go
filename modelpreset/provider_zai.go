package modelpreset

import "github.com/flexigpt/inference-go/spec"

const (
	ModelNameGLM47Flash30BA3BRepo  spec.ModelName = "zai-org/GLM-4.7-Flash"
	ModelNameGLM47Flash30BA3BLocal spec.ModelName = "glm-4.7-flash"

	ModelNameGLM52FireworksAI    spec.ModelName = "zai-org/GLM-5.2:fireworks-ai"
	ModelNameGLM51FireworksAI    spec.ModelName = "zai-org/GLM-5.1:fireworks-ai"
	ModelNameGLM51FP8FireworksAI spec.ModelName = "zai-org/GLM-5.1-FP8:fireworks-ai"

	ModelNameGLM52FP8ZAI spec.ModelName = "zai-org/GLM-5.2-FP8:zai-org"

	ModelNameGLM5Novita    spec.ModelName = "zai-org/GLM-5:novita"
	ModelNameGLM47Cerebras spec.ModelName = "zai-org/GLM-4.7:cerebras"

	ModelNameOpenRouterZAIGLM51 spec.ModelName = "z-ai/glm-5.1"
	ModelNameOpenRouterZAIGLM52 spec.ModelName = "z-ai/glm-5.2"
)

const (
	DisplayNameGLM47Flash30BA3B = "GLM-4.7 Flash 30B A3B"
	DisplayNameGLM47            = "GLM 4.7"
	DisplayNameGLM5             = "GLM 5"
	DisplayNameGLM51            = "GLM 5.1"
	DisplayNameGLM51FP8         = "GLM 5.1 FP8"
	DisplayNameGLM52            = "GLM 5.2"
	DisplayNameGLM52FP8         = "GLM 5.2 FP8"

	DisplayNameZAIGLM51    = "Z.AI GLM 5.1"
	DisplayNameZAIGLM52    = "Z.AI GLM 5.2"
	DisplayNameGLM52FP8ZAI = "GLM 5.2 FP8 (Z.AI)"

	DisplayNameGLM51FireworksAI    = "GLM 5.1 (Fireworks AI)"
	DisplayNameGLM51FP8FireworksAI = "GLM 5.1 FP8 (Fireworks AI)"
	DisplayNameGLM52FireworksAI    = "GLM 5.2 (Fireworks AI)"

	DisplayNameGLM47Cerebras = "GLM 4.7 (Cerebras)"
	DisplayNameGLM5Novita    = "GLM 5 (Novita)"
)

const (
	PresetGLM47Flash30BA3B ModelPresetID = "glm47flash30ba3b"
	PresetGLM47            ModelPresetID = "glm47"
	PresetGLM5             ModelPresetID = "glm5"
	PresetGLM51            ModelPresetID = "glm51"
	PresetGLM51FP8         ModelPresetID = "glm51fp8"
	PresetGLM52            ModelPresetID = "glm52"
	PresetGLM52FP8         ModelPresetID = "glm52fp8"

	PresetZAIGLM51    ModelPresetID = "zaiglm51"
	PresetZAIGLM52    ModelPresetID = "zaiglm52"
	PresetGLM52FP8ZAI ModelPresetID = "glm52fp8ZAI"

	PresetGLM52FireworksAI    ModelPresetID = "glm52FireworksAI"
	PresetGLM51FireworksAI    ModelPresetID = "glm51FireworksAI"
	PresetGLM51FP8FireworksAI ModelPresetID = "glm51fp8FireworksAI"

	PresetGLM5Novita    ModelPresetID = "glm5Novita"
	PresetGLM47Cerebras ModelPresetID = "glm47Cerebras"
)
