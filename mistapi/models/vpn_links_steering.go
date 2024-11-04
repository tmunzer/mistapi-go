package models

import (
    "encoding/json"
)

// VpnLinksSteering represents a VpnLinksSteering struct.
type VpnLinksSteering struct {
    Preference           *int           `json:"preference,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VpnLinksSteering.
// It customizes the JSON marshaling process for VpnLinksSteering objects.
func (v VpnLinksSteering) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VpnLinksSteering object to a map representation for JSON marshaling.
func (v VpnLinksSteering) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Preference != nil {
        structMap["preference"] = v.Preference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnLinksSteering.
// It customizes the JSON unmarshaling process for VpnLinksSteering objects.
func (v *VpnLinksSteering) UnmarshalJSON(input []byte) error {
    var temp tempVpnLinksSteering
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "preference")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Preference = temp.Preference
    return nil
}

// tempVpnLinksSteering is a temporary struct used for validating the fields of VpnLinksSteering.
type tempVpnLinksSteering  struct {
    Preference *int `json:"preference,omitempty"`
}
