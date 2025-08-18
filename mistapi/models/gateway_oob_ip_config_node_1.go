// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayOobIpConfigNode1 represents a GatewayOobIpConfigNode1 struct.
// For HA Cluster, node1 can have different IP Config
type GatewayOobIpConfigNode1 struct {
	// If `type`==`static`
	Gateway *string `json:"gateway,omitempty"`
	Ip      *string `json:"ip,omitempty"`
	// Used only if `subnet` is not specified in `networks`
	Netmask *string `json:"netmask,omitempty"`
	// enum: `dhcp`, `static`
	Type *IpTypeEnum `json:"type,omitempty"`
	// If supported on the platform. If enabled, DNS will be using this routing-instance, too
	UseMgmtVrf *bool `json:"use_mgmt_vrf,omitempty"`
	// Whether to use `mgmt_junos` for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired
	UseMgmtVrfForHostOut *bool                  `json:"use_mgmt_vrf_for_host_out,omitempty"`
	VlanId               *string                `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayOobIpConfigNode1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayOobIpConfigNode1) String() string {
	return fmt.Sprintf(
		"GatewayOobIpConfigNode1[Gateway=%v, Ip=%v, Netmask=%v, Type=%v, UseMgmtVrf=%v, UseMgmtVrfForHostOut=%v, VlanId=%v, AdditionalProperties=%v]",
		g.Gateway, g.Ip, g.Netmask, g.Type, g.UseMgmtVrf, g.UseMgmtVrfForHostOut, g.VlanId, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayOobIpConfigNode1.
// It customizes the JSON marshaling process for GatewayOobIpConfigNode1 objects.
func (g GatewayOobIpConfigNode1) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"gateway", "ip", "netmask", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "vlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayOobIpConfigNode1 object to a map representation for JSON marshaling.
func (g GatewayOobIpConfigNode1) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayOobIpConfigNode1.
// It customizes the JSON unmarshaling process for GatewayOobIpConfigNode1 objects.
func (g *GatewayOobIpConfigNode1) UnmarshalJSON(input []byte) error {
	var temp tempGatewayOobIpConfigNode1
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateway", "ip", "netmask", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "vlan_id")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Gateway = temp.Gateway
	g.Ip = temp.Ip
	g.Netmask = temp.Netmask
	g.Type = temp.Type
	g.UseMgmtVrf = temp.UseMgmtVrf
	g.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
	g.VlanId = temp.VlanId
	return nil
}

// tempGatewayOobIpConfigNode1 is a temporary struct used for validating the fields of GatewayOobIpConfigNode1.
type tempGatewayOobIpConfigNode1 struct {
	Gateway              *string     `json:"gateway,omitempty"`
	Ip                   *string     `json:"ip,omitempty"`
	Netmask              *string     `json:"netmask,omitempty"`
	Type                 *IpTypeEnum `json:"type,omitempty"`
	UseMgmtVrf           *bool       `json:"use_mgmt_vrf,omitempty"`
	UseMgmtVrfForHostOut *bool       `json:"use_mgmt_vrf_for_host_out,omitempty"`
	VlanId               *string     `json:"vlan_id,omitempty"`
}
