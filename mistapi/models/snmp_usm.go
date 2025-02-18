package models

import (
    "encoding/json"
    "fmt"
)

// SnmpUsm represents a SnmpUsm struct.
type SnmpUsm struct {
    // Required only if `engine_type`==`remote_engine`
    EngineId             *string                `json:"engine-id,omitempty"`
    // enum: `local_engine`, `remote_engine`
    EngineType           *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    Users                []SnmpUsmpUser         `json:"users,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpUsm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpUsm) String() string {
    return fmt.Sprintf(
    	"SnmpUsm[EngineId=%v, EngineType=%v, Users=%v, AdditionalProperties=%v]",
    	s.EngineId, s.EngineType, s.Users, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpUsm.
// It customizes the JSON marshaling process for SnmpUsm objects.
func (s SnmpUsm) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "engine-id", "engine_type", "users"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpUsm object to a map representation for JSON marshaling.
func (s SnmpUsm) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSnmpUsm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "engine-id", "engine_type", "users")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.EngineId = temp.EngineId
    s.EngineType = temp.EngineType
    s.Users = temp.Users
    return nil
}

// tempSnmpUsm is a temporary struct used for validating the fields of SnmpUsm.
type tempSnmpUsm  struct {
    EngineId   *string                `json:"engine-id,omitempty"`
    EngineType *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    Users      []SnmpUsmpUser         `json:"users,omitempty"`
}
