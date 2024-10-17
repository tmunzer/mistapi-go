package models

import (
	"encoding/json"
)

// HaClusterDelete represents a HaClusterDelete struct.
type HaClusterDelete struct {
	// node0 mac address
	Mac                  *string        `json:"mac,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for HaClusterDelete.
// It customizes the JSON marshaling process for HaClusterDelete objects.
func (h HaClusterDelete) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(h.toMap())
}

// toMap converts the HaClusterDelete object to a map representation for JSON marshaling.
func (h HaClusterDelete) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, h.AdditionalProperties)
	if h.Mac != nil {
		structMap["mac"] = h.Mac
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterDelete.
// It customizes the JSON unmarshaling process for HaClusterDelete objects.
func (h *HaClusterDelete) UnmarshalJSON(input []byte) error {
	var temp tempHaClusterDelete
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac")
	if err != nil {
		return err
	}

	h.AdditionalProperties = additionalProperties
	h.Mac = temp.Mac
	return nil
}

// tempHaClusterDelete is a temporary struct used for validating the fields of HaClusterDelete.
type tempHaClusterDelete struct {
	Mac *string `json:"mac,omitempty"`
}
