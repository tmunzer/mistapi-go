// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MxclusterNacedge represents a MxclusterNacedge struct.
// NAC Edge survivability settings for a Mist Edge cluster. Requires `mist_nac` to be enabled on the cluster
type MxclusterNacedge struct {
	// Cache TTL for last auth result in seconds
	AuthTtl *int `json:"auth_ttl,omitempty"`
	// Site UUIDs whose auth requests are cached by NAC Edges in the cluster
	CachingSiteIds []uuid.UUID `json:"caching_site_ids,omitempty"`
	// Default VLAN for all dot1x devices, if different from default_vlan
	DefaultDot1xVlan *string `json:"default_dot1x_vlan,omitempty"`
	// Default VLAN to assign for devices not in the cache
	DefaultVlan *string `json:"default_vlan,omitempty"`
	// Whether NAC Edge survivability is enabled for this cluster
	Enabled *bool `json:"enabled,omitempty"`
	// List of Mist NAC Edge hosts for AP survivability authentication
	NacEdgeHosts         []string               `json:"nac_edge_hosts,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterNacedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterNacedge) String() string {
	return fmt.Sprintf(
		"MxclusterNacedge[AuthTtl=%v, CachingSiteIds=%v, DefaultDot1xVlan=%v, DefaultVlan=%v, Enabled=%v, NacEdgeHosts=%v, AdditionalProperties=%v]",
		m.AuthTtl, m.CachingSiteIds, m.DefaultDot1xVlan, m.DefaultVlan, m.Enabled, m.NacEdgeHosts, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterNacedge.
// It customizes the JSON marshaling process for MxclusterNacedge objects.
func (m MxclusterNacedge) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"auth_ttl", "caching_site_ids", "default_dot1x_vlan", "default_vlan", "enabled", "nac_edge_hosts"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MxclusterNacedge object to a map representation for JSON marshaling.
func (m MxclusterNacedge) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AuthTtl != nil {
		structMap["auth_ttl"] = m.AuthTtl
	}
	if m.CachingSiteIds != nil {
		structMap["caching_site_ids"] = m.CachingSiteIds
	}
	if m.DefaultDot1xVlan != nil {
		structMap["default_dot1x_vlan"] = m.DefaultDot1xVlan
	}
	if m.DefaultVlan != nil {
		structMap["default_vlan"] = m.DefaultVlan
	}
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	if m.NacEdgeHosts != nil {
		structMap["nac_edge_hosts"] = m.NacEdgeHosts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterNacedge.
// It customizes the JSON unmarshaling process for MxclusterNacedge objects.
func (m *MxclusterNacedge) UnmarshalJSON(input []byte) error {
	var temp tempMxclusterNacedge
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_ttl", "caching_site_ids", "default_dot1x_vlan", "default_vlan", "enabled", "nac_edge_hosts")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AuthTtl = temp.AuthTtl
	m.CachingSiteIds = temp.CachingSiteIds
	m.DefaultDot1xVlan = temp.DefaultDot1xVlan
	m.DefaultVlan = temp.DefaultVlan
	m.Enabled = temp.Enabled
	m.NacEdgeHosts = temp.NacEdgeHosts
	return nil
}

// tempMxclusterNacedge is a temporary struct used for validating the fields of MxclusterNacedge.
type tempMxclusterNacedge struct {
	AuthTtl          *int        `json:"auth_ttl,omitempty"`
	CachingSiteIds   []uuid.UUID `json:"caching_site_ids,omitempty"`
	DefaultDot1xVlan *string     `json:"default_dot1x_vlan,omitempty"`
	DefaultVlan      *string     `json:"default_vlan,omitempty"`
	Enabled          *bool       `json:"enabled,omitempty"`
	NacEdgeHosts     []string    `json:"nac_edge_hosts,omitempty"`
}
