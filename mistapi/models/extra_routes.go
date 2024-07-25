package models

import (
    "encoding/json"
)

// ExtraRoutes represents a ExtraRoutes struct.
// Property key is the destination CIDR (e.g. "10.0.0.0/8")
type ExtraRoutes struct {
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ExtraRoutes.
// It customizes the JSON marshaling process for ExtraRoutes objects.
func (e ExtraRoutes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the ExtraRoutes object to a map representation for JSON marshaling.
func (e ExtraRoutes) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Via != nil {
        structMap["via"] = e.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRoutes.
// It customizes the JSON unmarshaling process for ExtraRoutes objects.
func (e *ExtraRoutes) UnmarshalJSON(input []byte) error {
    var temp extraRoutes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "via")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Via = temp.Via
    return nil
}

// extraRoutes is a temporary struct used for validating the fields of ExtraRoutes.
type extraRoutes  struct {
    Via *string `json:"via,omitempty"`
}
