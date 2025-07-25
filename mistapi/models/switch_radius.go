// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchRadius represents a SwitchRadius struct.
// By default, `radius_config` will be used. if a different one has to be used set `use_different_radius
type SwitchRadius struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Junos Radius config
    RadiusConfig         *SwitchRadiusConfig    `json:"radius_config,omitempty"`
    UseDifferentRadius   *string                `json:"use_different_radius,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchRadius,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchRadius) String() string {
    return fmt.Sprintf(
    	"SwitchRadius[Enabled=%v, RadiusConfig=%v, UseDifferentRadius=%v, AdditionalProperties=%v]",
    	s.Enabled, s.RadiusConfig, s.UseDifferentRadius, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchRadius.
// It customizes the JSON marshaling process for SwitchRadius objects.
func (s SwitchRadius) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "radius_config", "use_different_radius"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchRadius object to a map representation for JSON marshaling.
func (s SwitchRadius) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.RadiusConfig != nil {
        structMap["radius_config"] = s.RadiusConfig.toMap()
    }
    if s.UseDifferentRadius != nil {
        structMap["use_different_radius"] = s.UseDifferentRadius
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchRadius.
// It customizes the JSON unmarshaling process for SwitchRadius objects.
func (s *SwitchRadius) UnmarshalJSON(input []byte) error {
    var temp tempSwitchRadius
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "radius_config", "use_different_radius")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.RadiusConfig = temp.RadiusConfig
    s.UseDifferentRadius = temp.UseDifferentRadius
    return nil
}

// tempSwitchRadius is a temporary struct used for validating the fields of SwitchRadius.
type tempSwitchRadius  struct {
    Enabled            *bool               `json:"enabled,omitempty"`
    RadiusConfig       *SwitchRadiusConfig `json:"radius_config,omitempty"`
    UseDifferentRadius *string             `json:"use_different_radius,omitempty"`
}
