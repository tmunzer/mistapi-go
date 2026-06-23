// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TuntermDhcpdConfigProperty represents a TuntermDhcpdConfigProperty struct.
// Per-VLAN DHCP relay configuration for tunnel termination
type TuntermDhcpdConfigProperty struct {
	// Whether DHCP relay is enabled for this tunneled VLAN
	Enabled *bool `json:"enabled,omitempty"`
	// Unique string values returned or accepted by this schema
	Servers []string `json:"servers,omitempty"`
	// DHCP handling mode for tunnel termination. enum: `relay`
	Type                 *TuntermDhcpdTypeEnum  `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TuntermDhcpdConfigProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TuntermDhcpdConfigProperty) String() string {
	return fmt.Sprintf(
		"TuntermDhcpdConfigProperty[Enabled=%v, Servers=%v, Type=%v, AdditionalProperties=%v]",
		t.Enabled, t.Servers, t.Type, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TuntermDhcpdConfigProperty.
// It customizes the JSON marshaling process for TuntermDhcpdConfigProperty objects.
func (t TuntermDhcpdConfigProperty) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"enabled", "servers", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TuntermDhcpdConfigProperty object to a map representation for JSON marshaling.
func (t TuntermDhcpdConfigProperty) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.Enabled != nil {
		structMap["enabled"] = t.Enabled
	}
	if t.Servers != nil {
		structMap["servers"] = t.Servers
	}
	if t.Type != nil {
		structMap["type"] = t.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermDhcpdConfigProperty.
// It customizes the JSON unmarshaling process for TuntermDhcpdConfigProperty objects.
func (t *TuntermDhcpdConfigProperty) UnmarshalJSON(input []byte) error {
	var temp tempTuntermDhcpdConfigProperty
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "servers", "type")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.Enabled = temp.Enabled
	t.Servers = temp.Servers
	t.Type = temp.Type
	return nil
}

// tempTuntermDhcpdConfigProperty is a temporary struct used for validating the fields of TuntermDhcpdConfigProperty.
type tempTuntermDhcpdConfigProperty struct {
	Enabled *bool                 `json:"enabled,omitempty"`
	Servers []string              `json:"servers,omitempty"`
	Type    *TuntermDhcpdTypeEnum `json:"type,omitempty"`
}
