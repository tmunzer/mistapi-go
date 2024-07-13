package models

import (
    "encoding/json"
)

// SnmpVacm represents a SnmpVacm struct.
type SnmpVacm struct {
    Access               []SnmpVacmAccessItem     `json:"access,omitempty"`
    SecurityToGroup      *SnmpVacmSecurityToGroup `json:"security_to_group,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacm.
// It customizes the JSON marshaling process for SnmpVacm objects.
func (s SnmpVacm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacm object to a map representation for JSON marshaling.
func (s SnmpVacm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Access != nil {
        structMap["access"] = s.Access
    }
    if s.SecurityToGroup != nil {
        structMap["security_to_group"] = s.SecurityToGroup.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpVacm.
// It customizes the JSON unmarshaling process for SnmpVacm objects.
func (s *SnmpVacm) UnmarshalJSON(input []byte) error {
    var temp snmpVacm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "access", "security_to_group")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Access = temp.Access
    s.SecurityToGroup = temp.SecurityToGroup
    return nil
}

// snmpVacm is a temporary struct used for validating the fields of SnmpVacm.
type snmpVacm  struct {
    Access          []SnmpVacmAccessItem     `json:"access,omitempty"`
    SecurityToGroup *SnmpVacmSecurityToGroup `json:"security_to_group,omitempty"`
}