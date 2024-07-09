package models

import (
    "encoding/json"
)

// VrfExtraRoute represents a VrfExtraRoute struct.
type VrfExtraRoute struct {
    // Next-hop address
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrfExtraRoute.
// It customizes the JSON marshaling process for VrfExtraRoute objects.
func (v VrfExtraRoute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VrfExtraRoute object to a map representation for JSON marshaling.
func (v VrfExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Via != nil {
        structMap["via"] = v.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrfExtraRoute.
// It customizes the JSON unmarshaling process for VrfExtraRoute objects.
func (v *VrfExtraRoute) UnmarshalJSON(input []byte) error {
    var temp vrfExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "via")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Via = temp.Via
    return nil
}

// vrfExtraRoute is a temporary struct used for validating the fields of VrfExtraRoute.
type vrfExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}
