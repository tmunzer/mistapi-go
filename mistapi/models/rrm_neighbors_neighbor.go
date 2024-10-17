package models

import (
	"encoding/json"
)

// RrmNeighborsNeighbor represents a RrmNeighborsNeighbor struct.
type RrmNeighborsNeighbor struct {
	Mac                  *string        `json:"mac,omitempty"`
	Rssi                 *int           `json:"rssi,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmNeighborsNeighbor.
// It customizes the JSON marshaling process for RrmNeighborsNeighbor objects.
func (r RrmNeighborsNeighbor) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RrmNeighborsNeighbor object to a map representation for JSON marshaling.
func (r RrmNeighborsNeighbor) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Mac != nil {
		structMap["mac"] = r.Mac
	}
	if r.Rssi != nil {
		structMap["rssi"] = r.Rssi
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmNeighborsNeighbor.
// It customizes the JSON unmarshaling process for RrmNeighborsNeighbor objects.
func (r *RrmNeighborsNeighbor) UnmarshalJSON(input []byte) error {
	var temp tempRrmNeighborsNeighbor
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "rssi")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Mac = temp.Mac
	r.Rssi = temp.Rssi
	return nil
}

// tempRrmNeighborsNeighbor is a temporary struct used for validating the fields of RrmNeighborsNeighbor.
type tempRrmNeighborsNeighbor struct {
	Mac  *string `json:"mac,omitempty"`
	Rssi *int    `json:"rssi,omitempty"`
}
