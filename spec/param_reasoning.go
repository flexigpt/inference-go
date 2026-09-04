package spec

type (
	ReasoningLevel string
	ReasoningType  string
)

const (
	ReasoningTypeHybridWithTokens ReasoningType = "hybridWithTokens"
	ReasoningTypeSingleWithLevels ReasoningType = "singleWithLevels"
)

const (
	ReasoningLevelNone    ReasoningLevel = "none"
	ReasoningLevelMinimal ReasoningLevel = "minimal"
	ReasoningLevelLow     ReasoningLevel = "low"
	ReasoningLevelMedium  ReasoningLevel = "medium"
	ReasoningLevelHigh    ReasoningLevel = "high"
	ReasoningLevelXHigh   ReasoningLevel = "xhigh"
	ReasoningLevelMax     ReasoningLevel = "max"
)

// ReasoningSummaryStyle - what kind of summary should be emitted for the reasoning performed by the model.
type ReasoningSummaryStyle string

const (
	ReasoningSummaryStyleOmitted  ReasoningSummaryStyle = "omitted"
	ReasoningSummaryStyleAuto     ReasoningSummaryStyle = "auto"
	ReasoningSummaryStyleConcise  ReasoningSummaryStyle = "concise"
	ReasoningSummaryStyleDetailed ReasoningSummaryStyle = "detailed"
)

// ReasoningContext controls which reasoning items are rendered back to the model on later turns.
// Only supported by OpenAI Responses SDK as of now.
type ReasoningContext string

const (
	ReasoningContextAuto        ReasoningContext = "auto"
	ReasoningContextCurrentTurn ReasoningContext = "current_turn"
	ReasoningContextAllTurns    ReasoningContext = "all_turns"
)

// ReasoningMode controls the reasoning execution mode for the request.
//
// When returned on a response, this is the effective execution mode.
// Only supported by OpenAI Responses SDK as of now.
type ReasoningMode string

const (
	ReasoningModeStandard ReasoningMode = "standard"
	ReasoningModePro      ReasoningMode = "pro"
)

type ReasoningParam struct {
	Type   ReasoningType  `json:"type"`
	Level  ReasoningLevel `json:"level"`
	Tokens int            `json:"tokens"`

	SummaryStyle *ReasoningSummaryStyle `json:"summaryStyle,omitempty"`
	Context      *ReasoningContext      `json:"context,omitempty"`
	Mode         *ReasoningMode         `json:"mode,omitempty"`
}
