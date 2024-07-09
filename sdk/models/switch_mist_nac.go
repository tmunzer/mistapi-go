package models

import (
    "encoding/json"
)

// SwitchMistNac represents a SwitchMistNac struct.
// enable mist_nac to use radsec
type SwitchMistNac struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    Network              *string        `json:"network,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchMistNac.
// It customizes the JSON marshaling process for SwitchMistNac objects.
func (s SwitchMistNac) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMistNac object to a map representation for JSON marshaling.
func (s SwitchMistNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Network != nil {
        structMap["network"] = s.Network
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchMistNac.
// It customizes the JSON unmarshaling process for SwitchMistNac objects.
func (s *SwitchMistNac) UnmarshalJSON(input []byte) error {
    var temp switchMistNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "network")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Enabled = temp.Enabled
    s.Network = temp.Network
    return nil
}

// switchMistNac is a temporary struct used for validating the fields of SwitchMistNac.
type switchMistNac  struct {
    Enabled *bool   `json:"enabled,omitempty"`
    Network *string `json:"network,omitempty"`
}
