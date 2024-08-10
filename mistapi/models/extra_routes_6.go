package models

import (
    "encoding/json"
)

// ExtraRoutes6 represents a ExtraRoutes6 struct.
// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
type ExtraRoutes6 struct {
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ExtraRoutes6.
// It customizes the JSON marshaling process for ExtraRoutes6 objects.
func (e ExtraRoutes6) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the ExtraRoutes6 object to a map representation for JSON marshaling.
func (e ExtraRoutes6) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Via != nil {
        structMap["via"] = e.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRoutes6.
// It customizes the JSON unmarshaling process for ExtraRoutes6 objects.
func (e *ExtraRoutes6) UnmarshalJSON(input []byte) error {
    var temp tempExtraRoutes6
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

// tempExtraRoutes6 is a temporary struct used for validating the fields of ExtraRoutes6.
type tempExtraRoutes6  struct {
    Via *string `json:"via,omitempty"`
}
