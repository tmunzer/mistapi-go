// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WanExtraRoutes6 represents a WanExtraRoutes6 struct.
// Additional IPv6 route for a WAN interface
type WanExtraRoutes6 struct {
	// IPv6 next-hop address for this WAN extra route
	Via                  *string                `json:"via,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WanExtraRoutes6,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WanExtraRoutes6) String() string {
	return fmt.Sprintf(
		"WanExtraRoutes6[Via=%v, AdditionalProperties=%v]",
		w.Via, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WanExtraRoutes6.
// It customizes the JSON marshaling process for WanExtraRoutes6 objects.
func (w WanExtraRoutes6) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"via"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WanExtraRoutes6 object to a map representation for JSON marshaling.
func (w WanExtraRoutes6) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Via != nil {
		structMap["via"] = w.Via
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WanExtraRoutes6.
// It customizes the JSON unmarshaling process for WanExtraRoutes6 objects.
func (w *WanExtraRoutes6) UnmarshalJSON(input []byte) error {
	var temp tempWanExtraRoutes6
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

// tempWanExtraRoutes6 is a temporary struct used for validating the fields of WanExtraRoutes6.
type tempWanExtraRoutes6 struct {
	Via *string `json:"via,omitempty"`
}
