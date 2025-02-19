package models

import (
    "encoding/json"
    "fmt"
)

// SwitchMistNac represents a SwitchMistNac struct.
// Enable mist_nac to use RadSec
type SwitchMistNac struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    Network              *string                `json:"network,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchMistNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchMistNac) String() string {
    return fmt.Sprintf(
    	"SwitchMistNac[Enabled=%v, Network=%v, AdditionalProperties=%v]",
    	s.Enabled, s.Network, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchMistNac.
// It customizes the JSON marshaling process for SwitchMistNac objects.
func (s SwitchMistNac) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "network"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchMistNac object to a map representation for JSON marshaling.
func (s SwitchMistNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSwitchMistNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "network")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.Network = temp.Network
    return nil
}

// tempSwitchMistNac is a temporary struct used for validating the fields of SwitchMistNac.
type tempSwitchMistNac  struct {
    Enabled *bool   `json:"enabled,omitempty"`
    Network *string `json:"network,omitempty"`
}
