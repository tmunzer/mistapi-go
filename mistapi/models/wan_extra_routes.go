package models

import (
    "encoding/json"
)

// WanExtraRoutes represents a WanExtraRoutes struct.
type WanExtraRoutes struct {
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WanExtraRoutes.
// It customizes the JSON marshaling process for WanExtraRoutes objects.
func (w WanExtraRoutes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WanExtraRoutes object to a map representation for JSON marshaling.
func (w WanExtraRoutes) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Via != nil {
        structMap["via"] = w.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WanExtraRoutes.
// It customizes the JSON unmarshaling process for WanExtraRoutes objects.
func (w *WanExtraRoutes) UnmarshalJSON(input []byte) error {
    var temp tempWanExtraRoutes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "via")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Via = temp.Via
    return nil
}

// tempWanExtraRoutes is a temporary struct used for validating the fields of WanExtraRoutes.
type tempWanExtraRoutes  struct {
    Via *string `json:"via,omitempty"`
}
