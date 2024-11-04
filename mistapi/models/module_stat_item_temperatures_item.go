package models

import (
    "encoding/json"
)

// ModuleStatItemTemperaturesItem represents a ModuleStatItemTemperaturesItem struct.
type ModuleStatItemTemperaturesItem struct {
    Celsius              *float64       `json:"celsius,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemTemperaturesItem.
// It customizes the JSON marshaling process for ModuleStatItemTemperaturesItem objects.
func (m ModuleStatItemTemperaturesItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemTemperaturesItem object to a map representation for JSON marshaling.
func (m ModuleStatItemTemperaturesItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Celsius != nil {
        structMap["celsius"] = m.Celsius
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemTemperaturesItem.
// It customizes the JSON unmarshaling process for ModuleStatItemTemperaturesItem objects.
func (m *ModuleStatItemTemperaturesItem) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemTemperaturesItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "celsius", "name", "status")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Celsius = temp.Celsius
    m.Name = temp.Name
    m.Status = temp.Status
    return nil
}

// tempModuleStatItemTemperaturesItem is a temporary struct used for validating the fields of ModuleStatItemTemperaturesItem.
type tempModuleStatItemTemperaturesItem  struct {
    Celsius *float64 `json:"celsius,omitempty"`
    Name    *string  `json:"name,omitempty"`
    Status  *string  `json:"status,omitempty"`
}
