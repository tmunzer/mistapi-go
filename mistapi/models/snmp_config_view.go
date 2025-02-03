package models

import (
    "encoding/json"
    "fmt"
)

// SnmpConfigView represents a SnmpConfigView struct.
type SnmpConfigView struct {
    // If the root oid configured is included
    Include              *bool                  `json:"include,omitempty"`
    Oid                  *string                `json:"oid,omitempty"`
    ViewName             *string                `json:"view_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpConfigView,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpConfigView) String() string {
    return fmt.Sprintf(
    	"SnmpConfigView[Include=%v, Oid=%v, ViewName=%v, AdditionalProperties=%v]",
    	s.Include, s.Oid, s.ViewName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigView.
// It customizes the JSON marshaling process for SnmpConfigView objects.
func (s SnmpConfigView) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "include", "oid", "view_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigView object to a map representation for JSON marshaling.
func (s SnmpConfigView) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Include != nil {
        structMap["include"] = s.Include
    }
    if s.Oid != nil {
        structMap["oid"] = s.Oid
    }
    if s.ViewName != nil {
        structMap["view_name"] = s.ViewName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfigView.
// It customizes the JSON unmarshaling process for SnmpConfigView objects.
func (s *SnmpConfigView) UnmarshalJSON(input []byte) error {
    var temp tempSnmpConfigView
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "include", "oid", "view_name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Include = temp.Include
    s.Oid = temp.Oid
    s.ViewName = temp.ViewName
    return nil
}

// tempSnmpConfigView is a temporary struct used for validating the fields of SnmpConfigView.
type tempSnmpConfigView  struct {
    Include  *bool   `json:"include,omitempty"`
    Oid      *string `json:"oid,omitempty"`
    ViewName *string `json:"view_name,omitempty"`
}
