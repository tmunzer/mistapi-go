// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SnmpVacmSecurityToGroupContentItem represents a SnmpVacmSecurityToGroupContentItem struct.
type SnmpVacmSecurityToGroupContentItem struct {
    // Refer to group_name under access
    Group                *string                `json:"group,omitempty"`
    SecurityName         *string                `json:"security_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpVacmSecurityToGroupContentItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpVacmSecurityToGroupContentItem) String() string {
    return fmt.Sprintf(
    	"SnmpVacmSecurityToGroupContentItem[Group=%v, SecurityName=%v, AdditionalProperties=%v]",
    	s.Group, s.SecurityName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacmSecurityToGroupContentItem.
// It customizes the JSON marshaling process for SnmpVacmSecurityToGroupContentItem objects.
func (s SnmpVacmSecurityToGroupContentItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "group", "security_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacmSecurityToGroupContentItem object to a map representation for JSON marshaling.
func (s SnmpVacmSecurityToGroupContentItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Group != nil {
        structMap["group"] = s.Group
    }
    if s.SecurityName != nil {
        structMap["security_name"] = s.SecurityName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpVacmSecurityToGroupContentItem.
// It customizes the JSON unmarshaling process for SnmpVacmSecurityToGroupContentItem objects.
func (s *SnmpVacmSecurityToGroupContentItem) UnmarshalJSON(input []byte) error {
    var temp tempSnmpVacmSecurityToGroupContentItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "group", "security_name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Group = temp.Group
    s.SecurityName = temp.SecurityName
    return nil
}

// tempSnmpVacmSecurityToGroupContentItem is a temporary struct used for validating the fields of SnmpVacmSecurityToGroupContentItem.
type tempSnmpVacmSecurityToGroupContentItem  struct {
    Group        *string `json:"group,omitempty"`
    SecurityName *string `json:"security_name,omitempty"`
}
