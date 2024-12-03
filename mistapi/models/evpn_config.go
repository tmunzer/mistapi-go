package models

import (
    "encoding/json"
)

// EvpnConfig represents a EvpnConfig struct.
// EVPN Junos settings
type EvpnConfig struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    // enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`
    Role                 *EvpnConfigRoleEnum    `json:"role,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnConfig.
// It customizes the JSON marshaling process for EvpnConfig objects.
func (e EvpnConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "enabled", "role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnConfig object to a map representation for JSON marshaling.
func (e EvpnConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
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
    var temp tempEvpnConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "role")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Enabled = temp.Enabled
    e.Role = temp.Role
    return nil
}

// tempEvpnConfig is a temporary struct used for validating the fields of EvpnConfig.
type tempEvpnConfig  struct {
    Enabled *bool               `json:"enabled,omitempty"`
    Role    *EvpnConfigRoleEnum `json:"role,omitempty"`
}
