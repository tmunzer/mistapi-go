package models

import (
    "encoding/json"
)

// ConstMxedgeModelPort represents a ConstMxedgeModelPort struct.
type ConstMxedgeModelPort struct {
    Display              *string        `json:"display,omitempty"`
    Speed                *int           `json:"speed,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstMxedgeModelPort.
// It customizes the JSON marshaling process for ConstMxedgeModelPort objects.
func (c ConstMxedgeModelPort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstMxedgeModelPort object to a map representation for JSON marshaling.
func (c ConstMxedgeModelPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Speed != nil {
        structMap["speed"] = c.Speed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstMxedgeModelPort.
// It customizes the JSON unmarshaling process for ConstMxedgeModelPort objects.
func (c *ConstMxedgeModelPort) UnmarshalJSON(input []byte) error {
    var temp constMxedgeModelPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "display", "speed")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Display = temp.Display
    c.Speed = temp.Speed
    return nil
}

// constMxedgeModelPort is a temporary struct used for validating the fields of ConstMxedgeModelPort.
type constMxedgeModelPort  struct {
    Display *string `json:"display,omitempty"`
    Speed   *int    `json:"speed,omitempty"`
}
