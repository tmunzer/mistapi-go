package models

import (
    "encoding/json"
)

// VpnPathSelection represents a VpnPathSelection struct.
// Only if `type`==`hub_spoke`
type VpnPathSelection struct {
    // enum: `disabled`, `simple`, `manual`
    Strategy             *VpnPathSelectionStrategyEnum `json:"strategy,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VpnPathSelection.
// It customizes the JSON marshaling process for VpnPathSelection objects.
func (v VpnPathSelection) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPathSelection object to a map representation for JSON marshaling.
func (v VpnPathSelection) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Strategy != nil {
        structMap["strategy"] = v.Strategy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPathSelection.
// It customizes the JSON unmarshaling process for VpnPathSelection objects.
func (v *VpnPathSelection) UnmarshalJSON(input []byte) error {
    var temp tempVpnPathSelection
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "strategy")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Strategy = temp.Strategy
    return nil
}

// tempVpnPathSelection is a temporary struct used for validating the fields of VpnPathSelection.
type tempVpnPathSelection  struct {
    Strategy *VpnPathSelectionStrategyEnum `json:"strategy,omitempty"`
}
