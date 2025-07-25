// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// TuntermDhcpdConfig represents a TuntermDhcpdConfig struct.
// DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID
type TuntermDhcpdConfig struct {
    Enabled              *bool                                 `json:"enabled,omitempty"`
    Servers              []string                              `json:"servers,omitempty"`
    // enum: `relay`
    Type                 *TuntermDhcpdTypeEnum                 `json:"type,omitempty"`
    AdditionalProperties map[string]TuntermDhcpdConfigProperty `json:"_"`
}

// String implements the fmt.Stringer interface for TuntermDhcpdConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TuntermDhcpdConfig) String() string {
    return fmt.Sprintf(
    	"TuntermDhcpdConfig[Enabled=%v, Servers=%v, Type=%v, AdditionalProperties=%v]",
    	t.Enabled, t.Servers, t.Type, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TuntermDhcpdConfig.
// It customizes the JSON marshaling process for TuntermDhcpdConfig objects.
func (t TuntermDhcpdConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "enabled", "servers", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TuntermDhcpdConfig object to a map representation for JSON marshaling.
func (t TuntermDhcpdConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Enabled != nil {
        structMap["enabled"] = t.Enabled
    }
    if t.Servers != nil {
        structMap["servers"] = t.Servers
    }
    if t.Type != nil {
        structMap["type"] = t.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermDhcpdConfig.
// It customizes the JSON unmarshaling process for TuntermDhcpdConfig objects.
func (t *TuntermDhcpdConfig) UnmarshalJSON(input []byte) error {
    var temp tempTuntermDhcpdConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[TuntermDhcpdConfigProperty](input, "enabled", "servers", "type")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Enabled = temp.Enabled
    t.Servers = temp.Servers
    t.Type = temp.Type
    return nil
}

// tempTuntermDhcpdConfig is a temporary struct used for validating the fields of TuntermDhcpdConfig.
type tempTuntermDhcpdConfig  struct {
    Enabled *bool                 `json:"enabled,omitempty"`
    Servers []string              `json:"servers,omitempty"`
    Type    *TuntermDhcpdTypeEnum `json:"type,omitempty"`
}
