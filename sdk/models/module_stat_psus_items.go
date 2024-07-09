package models

import (
    "encoding/json"
)

// ModuleStatPsusItems represents a ModuleStatPsusItems struct.
type ModuleStatPsusItems struct {
    Name                 *string        `json:"name,omitempty"`
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatPsusItems.
// It customizes the JSON marshaling process for ModuleStatPsusItems objects.
func (m ModuleStatPsusItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatPsusItems object to a map representation for JSON marshaling.
func (m ModuleStatPsusItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatPsusItems.
// It customizes the JSON unmarshaling process for ModuleStatPsusItems objects.
func (m *ModuleStatPsusItems) UnmarshalJSON(input []byte) error {
    var temp moduleStatPsusItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "status")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Name = temp.Name
    m.Status = temp.Status
    return nil
}

// moduleStatPsusItems is a temporary struct used for validating the fields of ModuleStatPsusItems.
type moduleStatPsusItems  struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
