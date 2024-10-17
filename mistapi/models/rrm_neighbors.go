package models

import (
	"encoding/json"
)

// RrmNeighbors represents a RrmNeighbors struct.
type RrmNeighbors struct {
	Mac                  *string                `json:"mac,omitempty"`
	Neighbors            []RrmNeighborsNeighbor `json:"neighbors,omitempty"`
	AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmNeighbors.
// It customizes the JSON marshaling process for RrmNeighbors objects.
func (r RrmNeighbors) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RrmNeighbors object to a map representation for JSON marshaling.
func (r RrmNeighbors) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Mac != nil {
		structMap["mac"] = r.Mac
	}
	if r.Neighbors != nil {
		structMap["neighbors"] = r.Neighbors
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmNeighbors.
// It customizes the JSON unmarshaling process for RrmNeighbors objects.
func (r *RrmNeighbors) UnmarshalJSON(input []byte) error {
	var temp tempRrmNeighbors
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "neighbors")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Mac = temp.Mac
	r.Neighbors = temp.Neighbors
	return nil
}

// tempRrmNeighbors is a temporary struct used for validating the fields of RrmNeighbors.
type tempRrmNeighbors struct {
	Mac       *string                `json:"mac,omitempty"`
	Neighbors []RrmNeighborsNeighbor `json:"neighbors,omitempty"`
}
