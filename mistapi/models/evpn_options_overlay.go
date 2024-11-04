package models

import (
    "encoding/json"
)

// EvpnOptionsOverlay represents a EvpnOptionsOverlay struct.
type EvpnOptionsOverlay struct {
    // optional, these are defaults
    As                   *int           `json:"as,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptionsOverlay.
// It customizes the JSON marshaling process for EvpnOptionsOverlay objects.
func (e EvpnOptionsOverlay) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptionsOverlay object to a map representation for JSON marshaling.
func (e EvpnOptionsOverlay) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.As != nil {
        structMap["as"] = e.As
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptionsOverlay.
// It customizes the JSON unmarshaling process for EvpnOptionsOverlay objects.
func (e *EvpnOptionsOverlay) UnmarshalJSON(input []byte) error {
    var temp tempEvpnOptionsOverlay
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "as")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.As = temp.As
    return nil
}

// tempEvpnOptionsOverlay is a temporary struct used for validating the fields of EvpnOptionsOverlay.
type tempEvpnOptionsOverlay  struct {
    As *int `json:"as,omitempty"`
}
