// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WanExtraRoutes represents a WanExtraRoutes struct.
type WanExtraRoutes struct {
	Via                  *string                `json:"via,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WanExtraRoutes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WanExtraRoutes) String() string {
	return fmt.Sprintf(
		"WanExtraRoutes[Via=%v, AdditionalProperties=%v]",
		w.Via, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WanExtraRoutes.
// It customizes the JSON marshaling process for WanExtraRoutes objects.
func (w WanExtraRoutes) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"via"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WanExtraRoutes object to a map representation for JSON marshaling.
func (w WanExtraRoutes) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Via != nil {
		structMap["via"] = w.Via
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WanExtraRoutes.
// It customizes the JSON unmarshaling process for WanExtraRoutes objects.
func (w *WanExtraRoutes) UnmarshalJSON(input []byte) error {
	var temp tempWanExtraRoutes
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Via = temp.Via
	return nil
}

// tempWanExtraRoutes is a temporary struct used for validating the fields of WanExtraRoutes.
type tempWanExtraRoutes struct {
	Via *string `json:"via,omitempty"`
}
