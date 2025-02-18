package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermDhcpdConfig represents a MxedgeTuntermDhcpdConfig struct.
// Global and per-VLAN. Property key is the VLAN ID
type MxedgeTuntermDhcpdConfig struct {
    Enabled              *bool                                       `json:"enabled,omitempty"`
    // List of DHCP servers; required if `type`==`relay`
    Servers              []string                                    `json:"servers,omitempty"`
    // enum: `relay`
    Type                 *MxedgeTuntermDhcpdTypeEnum                 `json:"type,omitempty"`
    AdditionalProperties map[string]MxedgeTuntermDhcpdConfigProperty `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermDhcpdConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermDhcpdConfig) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermDhcpdConfig[Enabled=%v, Servers=%v, Type=%v, AdditionalProperties=%v]",
    	m.Enabled, m.Servers, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermDhcpdConfig.
// It customizes the JSON marshaling process for MxedgeTuntermDhcpdConfig objects.
func (m MxedgeTuntermDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "enabled", "servers", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermDhcpdConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[MxedgeTuntermDhcpdConfigProperty](input, "enabled", "servers", "type")
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
