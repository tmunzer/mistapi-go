// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NetworkVpnAccessStaticNatProperty represents a NetworkVpnAccessStaticNatProperty struct.
type NetworkVpnAccessStaticNatProperty struct {
	// The Static NAT destination IP Address. Must be an IP Address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}")
	InternalIp           *string                `json:"internal_ip,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkVpnAccessStaticNatProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkVpnAccessStaticNatProperty) String() string {
	return fmt.Sprintf(
		"NetworkVpnAccessStaticNatProperty[InternalIp=%v, Name=%v, AdditionalProperties=%v]",
		n.InternalIp, n.Name, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkVpnAccessStaticNatProperty.
// It customizes the JSON marshaling process for NetworkVpnAccessStaticNatProperty objects.
func (n NetworkVpnAccessStaticNatProperty) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"internal_ip", "name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkVpnAccessStaticNatProperty object to a map representation for JSON marshaling.
func (n NetworkVpnAccessStaticNatProperty) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.InternalIp != nil {
		structMap["internal_ip"] = n.InternalIp
	}
	if n.Name != nil {
		structMap["name"] = n.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkVpnAccessStaticNatProperty.
// It customizes the JSON unmarshaling process for NetworkVpnAccessStaticNatProperty objects.
func (n *NetworkVpnAccessStaticNatProperty) UnmarshalJSON(input []byte) error {
	var temp tempNetworkVpnAccessStaticNatProperty
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "internal_ip", "name")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.InternalIp = temp.InternalIp
	n.Name = temp.Name
	return nil
}

// tempNetworkVpnAccessStaticNatProperty is a temporary struct used for validating the fields of NetworkVpnAccessStaticNatProperty.
type tempNetworkVpnAccessStaticNatProperty struct {
	InternalIp *string `json:"internal_ip,omitempty"`
	Name       *string `json:"name,omitempty"`
}
