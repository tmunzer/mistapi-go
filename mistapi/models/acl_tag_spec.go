// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// AclTagSpec represents a AclTagSpec struct.
type AclTagSpec struct {
    // Matched dst port, "0" means any
    PortRange            *string                `json:"port_range,omitempty"`
    // `tcp` / `udp` / `icmp` / `icmp6` / `gre` / `any` / `:protocol_number`, `protocol_number` is between 1-254, default is `any` `protocol_number` is between 1-254
    Protocol             *string                `json:"protocol,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AclTagSpec,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AclTagSpec) String() string {
    return fmt.Sprintf(
    	"AclTagSpec[PortRange=%v, Protocol=%v, AdditionalProperties=%v]",
    	a.PortRange, a.Protocol, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AclTagSpec.
// It customizes the JSON marshaling process for AclTagSpec objects.
func (a AclTagSpec) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "port_range", "protocol"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AclTagSpec object to a map representation for JSON marshaling.
func (a AclTagSpec) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.PortRange != nil {
        structMap["port_range"] = a.PortRange
    }
    if a.Protocol != nil {
        structMap["protocol"] = a.Protocol
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AclTagSpec.
// It customizes the JSON unmarshaling process for AclTagSpec objects.
func (a *AclTagSpec) UnmarshalJSON(input []byte) error {
    var temp tempAclTagSpec
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port_range", "protocol")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.PortRange = temp.PortRange
    a.Protocol = temp.Protocol
    return nil
}

// tempAclTagSpec is a temporary struct used for validating the fields of AclTagSpec.
type tempAclTagSpec  struct {
    PortRange *string `json:"port_range,omitempty"`
    Protocol  *string `json:"protocol,omitempty"`
}
