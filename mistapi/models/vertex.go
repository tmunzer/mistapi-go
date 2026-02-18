// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// Vertex represents a Vertex struct.
type Vertex struct {
	// X coordinate
	X *float64 `json:"X,omitempty"`
	// Y coordinate
	Y                    *float64               `json:"Y,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Vertex,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v Vertex) String() string {
	return fmt.Sprintf(
		"Vertex[X=%v, Y=%v, AdditionalProperties=%v]",
		v.X, v.Y, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Vertex.
// It customizes the JSON marshaling process for Vertex objects.
func (v Vertex) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"X", "Y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the Vertex object to a map representation for JSON marshaling.
func (v Vertex) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	if v.X != nil {
		structMap["X"] = v.X
	}
	if v.Y != nil {
		structMap["Y"] = v.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Vertex.
// It customizes the JSON unmarshaling process for Vertex objects.
func (v *Vertex) UnmarshalJSON(input []byte) error {
	var temp tempVertex
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "X", "Y")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.X = temp.X
	v.Y = temp.Y
	return nil
}

// tempVertex is a temporary struct used for validating the fields of Vertex.
type tempVertex struct {
	X *float64 `json:"X,omitempty"`
	Y *float64 `json:"Y,omitempty"`
}
