package models

import (
    "encoding/json"
)

// SnmpConfigView represents a SnmpConfigView struct.
type SnmpConfigView struct {
    // if the root oid configured is included
    Include              *bool          `json:"include,omitempty"`
    Oid                  *string        `json:"oid,omitempty"`
    ViewName             *string        `json:"view_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigView.
// It customizes the JSON marshaling process for SnmpConfigView objects.
func (s SnmpConfigView) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigView object to a map representation for JSON marshaling.
func (s SnmpConfigView) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "include", "oid", "view_name")
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
