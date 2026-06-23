// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsClearBpdu represents a UtilsClearBpdu struct.
// Request to clear detected BPDU errors on switch ports
type UtilsClearBpdu struct {
	// List of ports on which to clear the detected BPDU error, or `all` for all ports
	Ports                []string               `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsClearBpdu,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsClearBpdu) String() string {
	return fmt.Sprintf(
		"UtilsClearBpdu[Ports=%v, AdditionalProperties=%v]",
		u.Ports, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearBpdu.
// It customizes the JSON marshaling process for UtilsClearBpdu objects.
func (u UtilsClearBpdu) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"ports"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearBpdu object to a map representation for JSON marshaling.
func (u UtilsClearBpdu) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Ports != nil {
		structMap["ports"] = u.Ports
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearBpdu.
// It customizes the JSON unmarshaling process for UtilsClearBpdu objects.
func (u *UtilsClearBpdu) UnmarshalJSON(input []byte) error {
	var temp tempUtilsClearBpdu
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ports")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Ports = temp.Ports
	return nil
}

// tempUtilsClearBpdu is a temporary struct used for validating the fields of UtilsClearBpdu.
type tempUtilsClearBpdu struct {
	Ports []string `json:"ports,omitempty"`
}
