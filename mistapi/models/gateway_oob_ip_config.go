// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayOobIpConfig represents a GatewayOobIpConfig struct.
// Out-of-band (vme/em0/fxp0) IP config
type GatewayOobIpConfig struct {
	// If `type`==`static`
	Gateway *string `json:"gateway,omitempty"`
	// If `type`==`static`
	Ip *string `json:"ip,omitempty"`
	// If `type`==`static`
	Netmask *string `json:"netmask,omitempty"`
	// For HA Cluster, node1 can have different IP Config
	Node1 *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
	// enum: `dhcp`, `static`
	Type *IpTypeEnum `json:"type,omitempty"`
	// If supported on the platform. If enabled, DNS will be using this routing-instance, too
	UseMgmtVrf *bool `json:"use_mgmt_vrf,omitempty"`
	// For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired
	UseMgmtVrfForHostOut *bool                  `json:"use_mgmt_vrf_for_host_out,omitempty"`
	VlanId               *string                `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayOobIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayOobIpConfig) String() string {
	return fmt.Sprintf(
		"GatewayOobIpConfig[Gateway=%v, Ip=%v, Netmask=%v, Node1=%v, Type=%v, UseMgmtVrf=%v, UseMgmtVrfForHostOut=%v, VlanId=%v, AdditionalProperties=%v]",
		g.Gateway, g.Ip, g.Netmask, g.Node1, g.Type, g.UseMgmtVrf, g.UseMgmtVrfForHostOut, g.VlanId, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayOobIpConfig.
// It customizes the JSON marshaling process for GatewayOobIpConfig objects.
func (g GatewayOobIpConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"gateway", "ip", "netmask", "node1", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "vlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayOobIpConfig object to a map representation for JSON marshaling.
func (g GatewayOobIpConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Gateway != nil {
		structMap["gateway"] = g.Gateway
	}
	if g.Ip != nil {
		structMap["ip"] = g.Ip
	}
	if g.Netmask != nil {
		structMap["netmask"] = g.Netmask
	}
	if g.Node1 != nil {
		structMap["node1"] = g.Node1.toMap()
	}
	if g.Type != nil {
		structMap["type"] = g.Type
	}
	if g.UseMgmtVrf != nil {
		structMap["use_mgmt_vrf"] = g.UseMgmtVrf
	}
	if g.UseMgmtVrfForHostOut != nil {
		structMap["use_mgmt_vrf_for_host_out"] = g.UseMgmtVrfForHostOut
	}
	if g.VlanId != nil {
		structMap["vlan_id"] = g.VlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayOobIpConfig.
// It customizes the JSON unmarshaling process for GatewayOobIpConfig objects.
func (g *GatewayOobIpConfig) UnmarshalJSON(input []byte) error {
	var temp tempGatewayOobIpConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateway", "ip", "netmask", "node1", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "vlan_id")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Gateway = temp.Gateway
	g.Ip = temp.Ip
	g.Netmask = temp.Netmask
	g.Node1 = temp.Node1
	g.Type = temp.Type
	g.UseMgmtVrf = temp.UseMgmtVrf
	g.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
	g.VlanId = temp.VlanId
	return nil
}

// tempGatewayOobIpConfig is a temporary struct used for validating the fields of GatewayOobIpConfig.
type tempGatewayOobIpConfig struct {
	Gateway              *string                  `json:"gateway,omitempty"`
	Ip                   *string                  `json:"ip,omitempty"`
	Netmask              *string                  `json:"netmask,omitempty"`
	Node1                *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
	Type                 *IpTypeEnum              `json:"type,omitempty"`
	UseMgmtVrf           *bool                    `json:"use_mgmt_vrf,omitempty"`
	UseMgmtVrfForHostOut *bool                    `json:"use_mgmt_vrf_for_host_out,omitempty"`
	VlanId               *string                  `json:"vlan_id,omitempty"`
}
