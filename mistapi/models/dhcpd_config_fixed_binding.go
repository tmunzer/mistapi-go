// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// DhcpdConfigFixedBinding represents a DhcpdConfigFixedBinding struct.
type DhcpdConfigFixedBinding struct {
	Ip                   *string                `json:"ip,omitempty"`
	Ip6                  *string                `json:"ip6,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DhcpdConfigFixedBinding,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DhcpdConfigFixedBinding) String() string {
	return fmt.Sprintf(
		"DhcpdConfigFixedBinding[Ip=%v, Ip6=%v, Name=%v, AdditionalProperties=%v]",
		d.Ip, d.Ip6, d.Name, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigFixedBinding.
// It customizes the JSON marshaling process for DhcpdConfigFixedBinding objects.
func (d DhcpdConfigFixedBinding) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"ip", "ip6", "name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigFixedBinding object to a map representation for JSON marshaling.
func (d DhcpdConfigFixedBinding) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Ip != nil {
		structMap["ip"] = d.Ip
	}
	if d.Ip6 != nil {
		structMap["ip6"] = d.Ip6
	}
	if d.Name != nil {
		structMap["name"] = d.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigFixedBinding.
// It customizes the JSON unmarshaling process for DhcpdConfigFixedBinding objects.
func (d *DhcpdConfigFixedBinding) UnmarshalJSON(input []byte) error {
	var temp tempDhcpdConfigFixedBinding
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip", "ip6", "name")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Ip = temp.Ip
	d.Ip6 = temp.Ip6
	d.Name = temp.Name
	return nil
}

// tempDhcpdConfigFixedBinding is a temporary struct used for validating the fields of DhcpdConfigFixedBinding.
type tempDhcpdConfigFixedBinding struct {
	Ip   *string `json:"ip,omitempty"`
	Ip6  *string `json:"ip6,omitempty"`
	Name *string `json:"name,omitempty"`
}
