package models

import (
    "encoding/json"
)

// VrfInstance represents a VrfInstance struct.
type VrfInstance struct {
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes          map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
    Networks             []string                 `json:"networks,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrfInstance.
// It customizes the JSON marshaling process for VrfInstance objects.
func (v VrfInstance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VrfInstance object to a map representation for JSON marshaling.
func (v VrfInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.ExtraRoutes != nil {
        structMap["extra_routes"] = v.ExtraRoutes
    }
    if v.Networks != nil {
        structMap["networks"] = v.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrfInstance.
// It customizes the JSON unmarshaling process for VrfInstance objects.
func (v *VrfInstance) UnmarshalJSON(input []byte) error {
    var temp vrfInstance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "extra_routes", "networks")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.ExtraRoutes = temp.ExtraRoutes
    v.Networks = temp.Networks
    return nil
}

// vrfInstance is a temporary struct used for validating the fields of VrfInstance.
type vrfInstance  struct {
    ExtraRoutes map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
    Networks    []string                 `json:"networks,omitempty"`
}
