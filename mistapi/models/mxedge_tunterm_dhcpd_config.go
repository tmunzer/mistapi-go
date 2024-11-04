package models

import (
    "encoding/json"
)

// MxedgeTuntermDhcpdConfig represents a MxedgeTuntermDhcpdConfig struct.
// global and per-VLAN. Property key is the VLAN ID
type MxedgeTuntermDhcpdConfig struct {
    Enabled              *bool                       `json:"enabled,omitempty"`
    // list of DHCP servers; required if `type`==`relay`
    Servers              []string                    `json:"servers,omitempty"`
    // enum: `relay`
    Type                 *MxedgeTuntermDhcpdTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermDhcpdConfig.
// It customizes the JSON marshaling process for MxedgeTuntermDhcpdConfig objects.
func (m MxedgeTuntermDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermDhcpdConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.Servers != nil {
        structMap["servers"] = m.Servers
    }
    if m.Type != nil {
        structMap["type"] = m.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermDhcpdConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermDhcpdConfig objects.
func (m *MxedgeTuntermDhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermDhcpdConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "servers", "type")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Enabled = temp.Enabled
    m.Servers = temp.Servers
    m.Type = temp.Type
    return nil
}

// tempMxedgeTuntermDhcpdConfig is a temporary struct used for validating the fields of MxedgeTuntermDhcpdConfig.
type tempMxedgeTuntermDhcpdConfig  struct {
    Enabled *bool                       `json:"enabled,omitempty"`
    Servers []string                    `json:"servers,omitempty"`
    Type    *MxedgeTuntermDhcpdTypeEnum `json:"type,omitempty"`
}
