package models

import (
    "encoding/json"
)

// ConstMxedgeModel represents a ConstMxedgeModel struct.
type ConstMxedgeModel struct {
    CustomPorts          *bool                           `json:"custom_ports,omitempty"`
    Display              *string                         `json:"display,omitempty"`
    Model                *string                         `json:"model,omitempty"`
    Ports                map[string]ConstMxedgeModelPort `json:"ports,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstMxedgeModel.
// It customizes the JSON marshaling process for ConstMxedgeModel objects.
func (c ConstMxedgeModel) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstMxedgeModel object to a map representation for JSON marshaling.
func (c ConstMxedgeModel) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.CustomPorts != nil {
        structMap["custom_ports"] = c.CustomPorts
    }
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.Ports != nil {
        structMap["ports"] = c.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstMxedgeModel.
// It customizes the JSON unmarshaling process for ConstMxedgeModel objects.
func (c *ConstMxedgeModel) UnmarshalJSON(input []byte) error {
    var temp constMxedgeModel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "custom_ports", "display", "model", "ports")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.CustomPorts = temp.CustomPorts
    c.Display = temp.Display
    c.Model = temp.Model
    c.Ports = temp.Ports
    return nil
}

// constMxedgeModel is a temporary struct used for validating the fields of ConstMxedgeModel.
type constMxedgeModel  struct {
    CustomPorts *bool                           `json:"custom_ports,omitempty"`
    Display     *string                         `json:"display,omitempty"`
    Model       *string                         `json:"model,omitempty"`
    Ports       map[string]ConstMxedgeModelPort `json:"ports,omitempty"`
}
