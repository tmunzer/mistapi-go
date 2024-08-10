package models

import (
    "encoding/json"
)

// ModuleStatItemPoe represents a ModuleStatItemPoe struct.
type ModuleStatItemPoe struct {
    MaxPower             *float64       `json:"max_power,omitempty"`
    PowerDraw            *float64       `json:"power_draw,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemPoe.
// It customizes the JSON marshaling process for ModuleStatItemPoe objects.
func (m ModuleStatItemPoe) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemPoe object to a map representation for JSON marshaling.
func (m ModuleStatItemPoe) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.MaxPower != nil {
        structMap["max_power"] = m.MaxPower
    }
    if m.PowerDraw != nil {
        structMap["power_draw"] = m.PowerDraw
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemPoe.
// It customizes the JSON unmarshaling process for ModuleStatItemPoe objects.
func (m *ModuleStatItemPoe) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemPoe
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "max_power", "power_draw")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.MaxPower = temp.MaxPower
    m.PowerDraw = temp.PowerDraw
    return nil
}

// tempModuleStatItemPoe is a temporary struct used for validating the fields of ModuleStatItemPoe.
type tempModuleStatItemPoe  struct {
    MaxPower  *float64 `json:"max_power,omitempty"`
    PowerDraw *float64 `json:"power_draw,omitempty"`
}
