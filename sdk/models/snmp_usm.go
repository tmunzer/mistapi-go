package models

import (
    "encoding/json"
)

// SnmpUsm represents a SnmpUsm struct.
type SnmpUsm struct {
    // required only if `engine_type`==`remote_engine`
    EngineId             *string                `json:"engine-id,omitempty"`
    EngineType           *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    Users                []SnmpUsmpUser         `json:"users,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpUsm.
// It customizes the JSON marshaling process for SnmpUsm objects.
func (s SnmpUsm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpUsm object to a map representation for JSON marshaling.
func (s SnmpUsm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.EngineId != nil {
        structMap["engine-id"] = s.EngineId
    }
    if s.EngineType != nil {
        structMap["engine_type"] = s.EngineType
    }
    if s.Users != nil {
        structMap["users"] = s.Users
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpUsm.
// It customizes the JSON unmarshaling process for SnmpUsm objects.
func (s *SnmpUsm) UnmarshalJSON(input []byte) error {
    var temp snmpUsm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "engine-id", "engine_type", "users")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.EngineId = temp.EngineId
    s.EngineType = temp.EngineType
    s.Users = temp.Users
    return nil
}

// snmpUsm is a temporary struct used for validating the fields of SnmpUsm.
type snmpUsm  struct {
    EngineId   *string                `json:"engine-id,omitempty"`
    EngineType *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    Users      []SnmpUsmpUser         `json:"users,omitempty"`
}
