package models

import (
    "encoding/json"
)

// SwitchDhcpdConfig represents a SwitchDhcpdConfig struct.
type SwitchDhcpdConfig struct {
    // if set to `true`, enable the DHCP server
    Enabled              *bool                                `json:"enabled,omitempty"`
    AdditionalProperties map[string]SwitchDhcpdConfigProperty `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchDhcpdConfig.
// It customizes the JSON marshaling process for SwitchDhcpdConfig objects.
func (s SwitchDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchDhcpdConfig object to a map representation for JSON marshaling.
func (s SwitchDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchDhcpdConfig.
// It customizes the JSON unmarshaling process for SwitchDhcpdConfig objects.
func (s *SwitchDhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempSwitchDhcpdConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[SwitchDhcpdConfigProperty](input, "enabled")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    return nil
}

// tempSwitchDhcpdConfig is a temporary struct used for validating the fields of SwitchDhcpdConfig.
type tempSwitchDhcpdConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
