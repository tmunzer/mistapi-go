package models

import (
    "encoding/json"
)

// VsInstanceProperty represents a VsInstanceProperty struct.
type VsInstanceProperty struct {
    Networks             []string       `json:"networks,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VsInstanceProperty.
// It customizes the JSON marshaling process for VsInstanceProperty objects.
func (v VsInstanceProperty) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VsInstanceProperty object to a map representation for JSON marshaling.
func (v VsInstanceProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Networks != nil {
        structMap["networks"] = v.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VsInstanceProperty.
// It customizes the JSON unmarshaling process for VsInstanceProperty objects.
func (v *VsInstanceProperty) UnmarshalJSON(input []byte) error {
    var temp tempVsInstanceProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "networks")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Networks = temp.Networks
    return nil
}

// tempVsInstanceProperty is a temporary struct used for validating the fields of VsInstanceProperty.
type tempVsInstanceProperty  struct {
    Networks []string `json:"networks,omitempty"`
}
