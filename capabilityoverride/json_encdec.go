package capabilityoverride

import "encoding/json"

func (o ModelCapabilitiesOverride) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	if o.ModalitiesIn != nil {
		m["modalitiesIn"] = o.ModalitiesIn
	}
	if o.ModalitiesOut != nil {
		m["modalitiesOut"] = o.ModalitiesOut
	}
	if o.ReasoningCapabilities != nil {
		m["reasoningCapabilities"] = o.ReasoningCapabilities
	}
	if o.StopSequenceCapabilities != nil {
		m["stopSequenceCapabilities"] = o.StopSequenceCapabilities
	}
	if o.OutputCapabilities != nil {
		m["outputCapabilities"] = o.OutputCapabilities
	}
	if o.ToolCapabilities != nil {
		m["toolCapabilities"] = o.ToolCapabilities
	}
	if o.CacheCapabilities != nil {
		m["cacheCapabilities"] = o.CacheCapabilities
	}
	if o.ParamDialect != nil {
		m["paramDialect"] = o.ParamDialect
	}

	return json.Marshal(m)
}

func (o ReasoningCapabilitiesOverride) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	if o.SupportsReasoningConfig != nil {
		m["supportsReasoningConfig"] = *o.SupportsReasoningConfig
	}
	if o.SupportedReasoningTypes != nil {
		m["supportedReasoningTypes"] = o.SupportedReasoningTypes
	}
	if o.SupportedReasoningLevels != nil {
		m["supportedReasoningLevels"] = o.SupportedReasoningLevels
	}
	if o.HybridTokenBudgetCapabilities != nil {
		m["hybridTokenBudgetCapabilities"] = o.HybridTokenBudgetCapabilities
	}
	if o.SupportsSummaryStyle != nil {
		m["supportsSummaryStyle"] = *o.SupportsSummaryStyle
	}
	if o.SupportsEncryptedReasoningInput != nil {
		m["supportsEncryptedReasoningInput"] = *o.SupportsEncryptedReasoningInput
	}
	if o.TemperatureDisallowedWhenEnabled != nil {
		m["temperatureDisallowedWhenEnabled"] = *o.TemperatureDisallowedWhenEnabled
	}

	return json.Marshal(m)
}

func (o OutputCapabilitiesOverride) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	if o.SupportedOutputFormats != nil {
		m["supportedOutputFormats"] = o.SupportedOutputFormats
	}
	if o.SupportsVerbosity != nil {
		m["supportsVerbosity"] = *o.SupportsVerbosity
	}

	return json.Marshal(m)
}

func (o ToolCapabilitiesOverride) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	if o.SupportedToolTypes != nil {
		m["supportedToolTypes"] = o.SupportedToolTypes
	}
	if o.SupportedToolPolicyModes != nil {
		m["supportedToolPolicyModes"] = o.SupportedToolPolicyModes
	}
	if o.SupportsParallelToolCalls != nil {
		m["supportsParallelToolCalls"] = *o.SupportsParallelToolCalls
	}
	if o.MaxForcedTools != nil {
		m["maxForcedTools"] = *o.MaxForcedTools
	}
	if o.SupportedClientToolOutputFormats != nil {
		m["supportedClientToolOutputFormats"] = o.SupportedClientToolOutputFormats
	}

	return json.Marshal(m)
}

func (o CacheControlCapabilitiesOverride) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	if o.SupportsTTL != nil {
		m["supportsTTL"] = *o.SupportsTTL
	}
	if o.SupportedKinds != nil {
		m["supportedKinds"] = o.SupportedKinds
	}
	if o.SupportedTTLs != nil {
		m["supportedTTLs"] = o.SupportedTTLs
	}
	if o.SupportsKey != nil {
		m["supportsKey"] = *o.SupportsKey
	}

	return json.Marshal(m)
}
