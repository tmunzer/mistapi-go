// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NetworkInternalAccess represents a NetworkInternalAccess struct.
type NetworkInternalAccess struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkInternalAccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkInternalAccess) String() string {
	return fmt.Sprintf(
		"NetworkInternalAccess[Enabled=%v, AdditionalProperties=%v]",
		n.Enabled, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkInternalAccess.
// It customizes the JSON marshaling process for NetworkInternalAccess objects.
func (n NetworkInternalAccess) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkInternalAccess object to a map representation for JSON marshaling.
func (n NetworkInternalAccess) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Enabled != nil {
		structMap["enabled"] = n.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkInternalAccess.
// It customizes the JSON unmarshaling process for NetworkInternalAccess objects.
func (n *NetworkInternalAccess) UnmarshalJSON(input []byte) error {
	var temp tempNetworkInternalAccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Enabled = temp.Enabled
	return nil
}

// tempNetworkInternalAccess is a temporary struct used for validating the fields of NetworkInternalAccess.
type tempNetworkInternalAccess struct {
	Enabled *bool `json:"enabled,omitempty"`
}
