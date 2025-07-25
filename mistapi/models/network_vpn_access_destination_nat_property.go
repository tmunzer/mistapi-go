// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// NetworkVpnAccessDestinationNatProperty represents a NetworkVpnAccessDestinationNatProperty struct.
type NetworkVpnAccessDestinationNatProperty struct {
    // The Destination NAT destination IP Address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}")
    InternalIp           *string                `json:"internal_ip,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Port                 *string                `json:"port,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NetworkVpnAccessDestinationNatProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NetworkVpnAccessDestinationNatProperty) String() string {
    return fmt.Sprintf(
    	"NetworkVpnAccessDestinationNatProperty[InternalIp=%v, Name=%v, Port=%v, AdditionalProperties=%v]",
    	n.InternalIp, n.Name, n.Port, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NetworkVpnAccessDestinationNatProperty.
// It customizes the JSON marshaling process for NetworkVpnAccessDestinationNatProperty objects.
func (n NetworkVpnAccessDestinationNatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "internal_ip", "name", "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkVpnAccessDestinationNatProperty object to a map representation for JSON marshaling.
func (n NetworkVpnAccessDestinationNatProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.InternalIp != nil {
        structMap["internal_ip"] = n.InternalIp
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.Port != nil {
        structMap["port"] = n.Port
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkVpnAccessDestinationNatProperty.
// It customizes the JSON unmarshaling process for NetworkVpnAccessDestinationNatProperty objects.
func (n *NetworkVpnAccessDestinationNatProperty) UnmarshalJSON(input []byte) error {
    var temp tempNetworkVpnAccessDestinationNatProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "internal_ip", "name", "port")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.InternalIp = temp.InternalIp
    n.Name = temp.Name
    n.Port = temp.Port
    return nil
}

// tempNetworkVpnAccessDestinationNatProperty is a temporary struct used for validating the fields of NetworkVpnAccessDestinationNatProperty.
type tempNetworkVpnAccessDestinationNatProperty  struct {
    InternalIp *string `json:"internal_ip,omitempty"`
    Name       *string `json:"name,omitempty"`
    Port       *string `json:"port,omitempty"`
}
