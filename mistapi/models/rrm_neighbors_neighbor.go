// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// RrmNeighborsNeighbor represents a RrmNeighborsNeighbor struct.
type RrmNeighborsNeighbor struct {
	Mac                  *string                `json:"mac,omitempty"`
	Rssi                 *int                   `json:"rssi,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RrmNeighborsNeighbor,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RrmNeighborsNeighbor) String() string {
	return fmt.Sprintf(
		"RrmNeighborsNeighbor[Mac=%v, Rssi=%v, AdditionalProperties=%v]",
		r.Mac, r.Rssi, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RrmNeighborsNeighbor.
// It customizes the JSON marshaling process for RrmNeighborsNeighbor objects.
func (r RrmNeighborsNeighbor) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"mac", "rssi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RrmNeighborsNeighbor object to a map representation for JSON marshaling.
func (r RrmNeighborsNeighbor) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "rssi")
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
