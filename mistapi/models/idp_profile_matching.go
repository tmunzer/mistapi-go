// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// IdpProfileMatching represents a IdpProfileMatching struct.
type IdpProfileMatching struct {
    AttackName           []string                              `json:"attack_name,omitempty"`
    DstSubnet            []string                              `json:"dst_subnet,omitempty"`
    Severity             []IdpProfileMatchingSeverityValueEnum `json:"severity,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for IdpProfileMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IdpProfileMatching) String() string {
    return fmt.Sprintf(
    	"IdpProfileMatching[AttackName=%v, DstSubnet=%v, Severity=%v, AdditionalProperties=%v]",
    	i.AttackName, i.DstSubnet, i.Severity, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IdpProfileMatching.
// It customizes the JSON marshaling process for IdpProfileMatching objects.
func (i IdpProfileMatching) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "attack_name", "dst_subnet", "severity"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IdpProfileMatching object to a map representation for JSON marshaling.
func (i IdpProfileMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.AttackName != nil {
        structMap["attack_name"] = i.AttackName
    }
    if i.DstSubnet != nil {
        structMap["dst_subnet"] = i.DstSubnet
    }
    if i.Severity != nil {
        structMap["severity"] = i.Severity
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IdpProfileMatching.
// It customizes the JSON unmarshaling process for IdpProfileMatching objects.
func (i *IdpProfileMatching) UnmarshalJSON(input []byte) error {
    var temp tempIdpProfileMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "attack_name", "dst_subnet", "severity")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.AttackName = temp.AttackName
    i.DstSubnet = temp.DstSubnet
    i.Severity = temp.Severity
    return nil
}

// tempIdpProfileMatching is a temporary struct used for validating the fields of IdpProfileMatching.
type tempIdpProfileMatching  struct {
    AttackName []string                              `json:"attack_name,omitempty"`
    DstSubnet  []string                              `json:"dst_subnet,omitempty"`
    Severity   []IdpProfileMatchingSeverityValueEnum `json:"severity,omitempty"`
}
