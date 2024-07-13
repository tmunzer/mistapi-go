package models

import (
    "encoding/json"
)

// SynthetictestConfig represents a SynthetictestConfig struct.
type SynthetictestConfig struct {
    Disabled             *bool                     `json:"disabled,omitempty"`
    Vlans                []SynthetictestProperties `json:"vlans,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfig.
// It customizes the JSON marshaling process for SynthetictestConfig objects.
func (s SynthetictestConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfig object to a map representation for JSON marshaling.
func (s SynthetictestConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Disabled != nil {
        structMap["disabled"] = s.Disabled
    }
    if s.Vlans != nil {
        structMap["vlans"] = s.Vlans
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfig.
// It customizes the JSON unmarshaling process for SynthetictestConfig objects.
func (s *SynthetictestConfig) UnmarshalJSON(input []byte) error {
    var temp synthetictestConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disabled", "vlans")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Disabled = temp.Disabled
    s.Vlans = temp.Vlans
    return nil
}

// synthetictestConfig is a temporary struct used for validating the fields of SynthetictestConfig.
type synthetictestConfig  struct {
    Disabled *bool                     `json:"disabled,omitempty"`
    Vlans    []SynthetictestProperties `json:"vlans,omitempty"`
}
