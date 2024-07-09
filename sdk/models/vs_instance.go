package models

import (
    "encoding/json"
)

// VsInstance represents a VsInstance struct.
type VsInstance struct {
    Networks             []string       `json:"networks,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VsInstance.
// It customizes the JSON marshaling process for VsInstance objects.
func (v VsInstance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VsInstance object to a map representation for JSON marshaling.
func (v VsInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Networks != nil {
        structMap["networks"] = v.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VsInstance.
// It customizes the JSON unmarshaling process for VsInstance objects.
func (v *VsInstance) UnmarshalJSON(input []byte) error {
    var temp vsInstance
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

// vsInstance is a temporary struct used for validating the fields of VsInstance.
type vsInstance  struct {
    Networks []string `json:"networks,omitempty"`
}
