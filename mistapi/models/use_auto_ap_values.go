package models

import (
    "encoding/json"
)

// UseAutoApValues represents a UseAutoApValues struct.
type UseAutoApValues struct {
    // If accept is true, accepts placement for devices in list otherwise. If false, reject for devices in list.
    Accept               *bool                   `json:"accept,omitempty"`
    // The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`
    For                  *UseAutoApValuesForEnum `json:"for,omitempty"`
    // A list of macs to accept/reject. If a list is not provided the API will accept/reject for the full map.
    Macs                 []string                `json:"macs,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UseAutoApValues.
// It customizes the JSON marshaling process for UseAutoApValues objects.
func (u UseAutoApValues) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UseAutoApValues object to a map representation for JSON marshaling.
func (u UseAutoApValues) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Accept != nil {
        structMap["accept"] = u.Accept
    }
    if u.For != nil {
        structMap["for"] = u.For
    }
    if u.Macs != nil {
        structMap["macs"] = u.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UseAutoApValues.
// It customizes the JSON unmarshaling process for UseAutoApValues objects.
func (u *UseAutoApValues) UnmarshalJSON(input []byte) error {
    var temp tempUseAutoApValues
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accept", "for", "macs")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Accept = temp.Accept
    u.For = temp.For
    u.Macs = temp.Macs
    return nil
}

// tempUseAutoApValues is a temporary struct used for validating the fields of UseAutoApValues.
type tempUseAutoApValues  struct {
    Accept *bool                   `json:"accept,omitempty"`
    For    *UseAutoApValuesForEnum `json:"for,omitempty"`
    Macs   []string                `json:"macs,omitempty"`
}
