package models

import (
    "encoding/json"
)

// EvpnTopologySwitchConfigDhcpdConfig represents a EvpnTopologySwitchConfigDhcpdConfig struct.
type EvpnTopologySwitchConfigDhcpdConfig struct {
    // If DHCPD is enabled on the switch
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitchConfigDhcpdConfig.
// It customizes the JSON marshaling process for EvpnTopologySwitchConfigDhcpdConfig objects.
func (e EvpnTopologySwitchConfigDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitchConfigDhcpdConfig object to a map representation for JSON marshaling.
func (e EvpnTopologySwitchConfigDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Enabled != nil {
        structMap["enabled"] = e.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitchConfigDhcpdConfig.
// It customizes the JSON unmarshaling process for EvpnTopologySwitchConfigDhcpdConfig objects.
func (e *EvpnTopologySwitchConfigDhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitchConfigDhcpdConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Enabled = temp.Enabled
    return nil
}

// tempEvpnTopologySwitchConfigDhcpdConfig is a temporary struct used for validating the fields of EvpnTopologySwitchConfigDhcpdConfig.
type tempEvpnTopologySwitchConfigDhcpdConfig  struct {
    Enabled *bool `json:"enabled,omitempty"`
}
