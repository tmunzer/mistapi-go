package models

import (
    "encoding/json"
    "fmt"
)

// SnmpUsm represents a SnmpUsm struct.
type SnmpUsm struct {
    // enum: `local_engine`, `remote_engine`
    EngineType           *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    // Required only if `engine_type`==`remote_engine`
    RemoteEngineId       *string                `json:"remote_engine_id,omitempty"`
    Users                []SnmpUsmUser          `json:"users,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpUsm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpUsm) String() string {
    return fmt.Sprintf(
    	"SnmpUsm[EngineType=%v, RemoteEngineId=%v, Users=%v, AdditionalProperties=%v]",
    	s.EngineType, s.RemoteEngineId, s.Users, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpUsm.
// It customizes the JSON marshaling process for SnmpUsm objects.
func (s SnmpUsm) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "engine_type", "remote_engine_id", "users"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpUsm object to a map representation for JSON marshaling.
func (s SnmpUsm) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.EngineType != nil {
        structMap["engine_type"] = s.EngineType
    }
    if s.RemoteEngineId != nil {
        structMap["remote_engine_id"] = s.RemoteEngineId
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "engine_type", "remote_engine_id", "users")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.EngineType = temp.EngineType
    s.RemoteEngineId = temp.RemoteEngineId
    s.Users = temp.Users
    return nil
}

// tempSnmpUsm is a temporary struct used for validating the fields of SnmpUsm.
type tempSnmpUsm  struct {
    EngineType     *SnmpUsmEngineTypeEnum `json:"engine_type,omitempty"`
    RemoteEngineId *string                `json:"remote_engine_id,omitempty"`
    Users          []SnmpUsmUser          `json:"users,omitempty"`
}
