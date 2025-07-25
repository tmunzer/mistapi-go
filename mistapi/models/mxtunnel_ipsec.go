// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxtunnelIpsec represents a MxtunnelIpsec struct.
type MxtunnelIpsec struct {
    DnsServers           Optional[[]string]        `json:"dns_servers"`
    DnsSuffix            []string                  `json:"dns_suffix,omitempty"`
    Enabled              *bool                     `json:"enabled,omitempty"`
    ExtraRoutes          []MxtunnelIpsecExtraRoute `json:"extra_routes,omitempty"`
    SplitTunnel          *bool                     `json:"split_tunnel,omitempty"`
    UseMxedge            *bool                     `json:"use_mxedge,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for MxtunnelIpsec,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxtunnelIpsec) String() string {
    return fmt.Sprintf(
    	"MxtunnelIpsec[DnsServers=%v, DnsSuffix=%v, Enabled=%v, ExtraRoutes=%v, SplitTunnel=%v, UseMxedge=%v, AdditionalProperties=%v]",
    	m.DnsServers, m.DnsSuffix, m.Enabled, m.ExtraRoutes, m.SplitTunnel, m.UseMxedge, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxtunnelIpsec.
// It customizes the JSON marshaling process for MxtunnelIpsec objects.
func (m MxtunnelIpsec) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "dns_servers", "dns_suffix", "enabled", "extra_routes", "split_tunnel", "use_mxedge"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxtunnelIpsec object to a map representation for JSON marshaling.
func (m MxtunnelIpsec) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.DnsServers.IsValueSet() {
        if m.DnsServers.Value() != nil {
            structMap["dns_servers"] = m.DnsServers.Value()
        } else {
            structMap["dns_servers"] = nil
        }
    }
    if m.DnsSuffix != nil {
        structMap["dns_suffix"] = m.DnsSuffix
    }
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.ExtraRoutes != nil {
        structMap["extra_routes"] = m.ExtraRoutes
    }
    if m.SplitTunnel != nil {
        structMap["split_tunnel"] = m.SplitTunnel
    }
    if m.UseMxedge != nil {
        structMap["use_mxedge"] = m.UseMxedge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxtunnelIpsec.
// It customizes the JSON unmarshaling process for MxtunnelIpsec objects.
func (m *MxtunnelIpsec) UnmarshalJSON(input []byte) error {
    var temp tempMxtunnelIpsec
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns_servers", "dns_suffix", "enabled", "extra_routes", "split_tunnel", "use_mxedge")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.DnsServers = temp.DnsServers
    m.DnsSuffix = temp.DnsSuffix
    m.Enabled = temp.Enabled
    m.ExtraRoutes = temp.ExtraRoutes
    m.SplitTunnel = temp.SplitTunnel
    m.UseMxedge = temp.UseMxedge
    return nil
}

// tempMxtunnelIpsec is a temporary struct used for validating the fields of MxtunnelIpsec.
type tempMxtunnelIpsec  struct {
    DnsServers  Optional[[]string]        `json:"dns_servers"`
    DnsSuffix   []string                  `json:"dns_suffix,omitempty"`
    Enabled     *bool                     `json:"enabled,omitempty"`
    ExtraRoutes []MxtunnelIpsecExtraRoute `json:"extra_routes,omitempty"`
    SplitTunnel *bool                     `json:"split_tunnel,omitempty"`
    UseMxedge   *bool                     `json:"use_mxedge,omitempty"`
}
