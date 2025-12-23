package maputil

import (
	"encoding/json"
	"errors"
	"fmt"
)

func StructWithJSONTagsToMap(data any) (map[string]any, error) {
	if data == nil {
		return nil, errors.New("input data cannot be nil")
	}
	// Marshal the struct to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}

	// Unmarshal the JSON into a map.
	var result map[string]any
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to map: %w", err)
	}

	return result, nil
}
