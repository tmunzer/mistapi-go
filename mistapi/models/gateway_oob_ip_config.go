package models

import (
    "encoding/json"
)

// GatewayOobIpConfig represents a GatewayOobIpConfig struct.
// out-of-band (vme/em0/fxp0) IP config
type GatewayOobIpConfig struct {
    // if `type`==`static`
    Gateway              *string                  `json:"gateway,omitempty"`
    // if `type`==`static`
    Ip                   *string                  `json:"ip,omitempty"`
    // if `type`==`static`
    Netmask              *string                  `json:"netmask,omitempty"`
    // optional, the network to be used for mgmt
    Network              *string                  `json:"network,omitempty"`
    // for HA Cluster, node1 can have different IP Config
    Node1                *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
    Type                 *IpTypeEnum              `json:"type,omitempty"`
    // if supported on the platform. If enabled, DNS will be using this routing-instance, too
    UseMgmtVrf           *bool                    `json:"use_mgmt_vrf,omitempty"`
    // for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired,
    UseMgmtVrfForHostOut *bool                    `json:"use_mgmt_vrf_for_host_out,omitempty"`
    VlanId               *string                  `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayOobIpConfig.
// It customizes the JSON marshaling process for GatewayOobIpConfig objects.
func (g GatewayOobIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayOobIpConfig object to a map representation for JSON marshaling.
func (g GatewayOobIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Gateway != nil {
        structMap["gateway"] = g.Gateway
    }
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.Netmask != nil {
        structMap["netmask"] = g.Netmask
    }
    if g.Network != nil {
        structMap["network"] = g.Network
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
    var temp gatewayOobIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway", "ip", "netmask", "network", "node1", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "vlan_id")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Gateway = temp.Gateway
    g.Ip = temp.Ip
    g.Netmask = temp.Netmask
    g.Network = temp.Network
    g.Node1 = temp.Node1
    g.Type = temp.Type
    g.UseMgmtVrf = temp.UseMgmtVrf
    g.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
    g.VlanId = temp.VlanId
    return nil
}

// gatewayOobIpConfig is a temporary struct used for validating the fields of GatewayOobIpConfig.
type gatewayOobIpConfig  struct {
    Gateway              *string                  `json:"gateway,omitempty"`
    Ip                   *string                  `json:"ip,omitempty"`
    Netmask              *string                  `json:"netmask,omitempty"`
    Network              *string                  `json:"network,omitempty"`
    Node1                *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
    Type                 *IpTypeEnum              `json:"type,omitempty"`
    UseMgmtVrf           *bool                    `json:"use_mgmt_vrf,omitempty"`
    UseMgmtVrfForHostOut *bool                    `json:"use_mgmt_vrf_for_host_out,omitempty"`
    VlanId               *string                  `json:"vlan_id,omitempty"`
}
