package models

import (
    "encoding/json"
)

// SnmpVacmSecurityToGroupContentItem represents a SnmpVacmSecurityToGroupContentItem struct.
type SnmpVacmSecurityToGroupContentItem struct {
    // refer to group_name under access
    Group                *string        `json:"group,omitempty"`
    SecurityName         *string        `json:"security_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacmSecurityToGroupContentItem.
// It customizes the JSON marshaling process for SnmpVacmSecurityToGroupContentItem objects.
func (s SnmpVacmSecurityToGroupContentItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacmSecurityToGroupContentItem object to a map representation for JSON marshaling.
func (s SnmpVacmSecurityToGroupContentItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp snmpVacmSecurityToGroupContentItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "group", "security_name")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Group = temp.Group
    s.SecurityName = temp.SecurityName
    return nil
}

// snmpVacmSecurityToGroupContentItem is a temporary struct used for validating the fields of SnmpVacmSecurityToGroupContentItem.
type snmpVacmSecurityToGroupContentItem  struct {
    Group        *string `json:"group,omitempty"`
    SecurityName *string `json:"security_name,omitempty"`
}
