package models

import (
    "encoding/json"
)

// ConfigSwitchRadius represents a ConfigSwitchRadius struct.
// by default, `radius_config` will be used. if a different one has to be used set `use_different_radius
type ConfigSwitchRadius struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    // Junos Radius config
    RadiusConfig         *RadiusConfig  `json:"radius_config,omitempty"`
    UseDifferentRadius   *string        `json:"use_different_radius,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConfigSwitchRadius.
// It customizes the JSON marshaling process for ConfigSwitchRadius objects.
func (c ConfigSwitchRadius) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigSwitchRadius object to a map representation for JSON marshaling.
func (c ConfigSwitchRadius) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Enabled != nil {
        structMap["enabled"] = c.Enabled
    }
    if c.RadiusConfig != nil {
        structMap["radius_config"] = c.RadiusConfig.toMap()
    }
    if c.UseDifferentRadius != nil {
        structMap["use_different_radius"] = c.UseDifferentRadius
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigSwitchRadius.
// It customizes the JSON unmarshaling process for ConfigSwitchRadius objects.
func (c *ConfigSwitchRadius) UnmarshalJSON(input []byte) error {
    var temp tempConfigSwitchRadius
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "radius_config", "use_different_radius")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Enabled = temp.Enabled
    c.RadiusConfig = temp.RadiusConfig
    c.UseDifferentRadius = temp.UseDifferentRadius
    return nil
}

// tempConfigSwitchRadius is a temporary struct used for validating the fields of ConfigSwitchRadius.
type tempConfigSwitchRadius  struct {
    Enabled            *bool         `json:"enabled,omitempty"`
    RadiusConfig       *RadiusConfig `json:"radius_config,omitempty"`
    UseDifferentRadius *string       `json:"use_different_radius,omitempty"`
}
