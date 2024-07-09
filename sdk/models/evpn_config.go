package models

import (
    "encoding/json"
)

// EvpnConfig represents a EvpnConfig struct.
// EVPN Junos settings
type EvpnConfig struct {
    Enabled              *bool               `json:"enabled,omitempty"`
    Role                 *EvpnConfigRoleEnum `json:"role,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnConfig.
// It customizes the JSON marshaling process for EvpnConfig objects.
func (e EvpnConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnConfig object to a map representation for JSON marshaling.
func (e EvpnConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Enabled != nil {
        structMap["enabled"] = e.Enabled
    }
    if e.Role != nil {
        structMap["role"] = e.Role
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnConfig.
// It customizes the JSON unmarshaling process for EvpnConfig objects.
func (e *EvpnConfig) UnmarshalJSON(input []byte) error {
    var temp evpnConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "role")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Enabled = temp.Enabled
    e.Role = temp.Role
    return nil
}

// evpnConfig is a temporary struct used for validating the fields of EvpnConfig.
type evpnConfig  struct {
    Enabled *bool               `json:"enabled,omitempty"`
    Role    *EvpnConfigRoleEnum `json:"role,omitempty"`
}
