package models

import (
    "encoding/json"
)

// SwitchOobIpConfig1 represents a SwitchOobIpConfig1 struct.
// - If HA configuration: key parameter will be nodeX (eg: node1)
// - If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`
type SwitchOobIpConfig1 struct {
    // if `type`==`static`
    Gateway              *string                  `json:"gateway,omitempty"`
    // if `type`==`static`
    Ip                   *string                  `json:"ip,omitempty"`
    // used only if `subnet` is not specified in `networks`
    Netmask              *string                  `json:"netmask,omitempty"`
    // optional, the network to be used for mgmt
    Network              *string                  `json:"network,omitempty"`
    Type                 *IpTypeEnum              `json:"type,omitempty"`
    // f supported on the platform. If enabled, DNS will be using this routing-instance, too
    UseMgmtVrf           *bool                    `json:"use_mgmt_vrf,omitempty"`
    // for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired,
    UseMgmtVrfForHostOut *bool                    `json:"use_mgmt_vrf_for_host_out,omitempty"`
    // for HA Cluster, node1 can have different IP Config
    Node1                *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
    VlanId               *string                  `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchOobIpConfig1.
// It customizes the JSON marshaling process for SwitchOobIpConfig1 objects.
func (s SwitchOobIpConfig1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchOobIpConfig1 object to a map representation for JSON marshaling.
func (s SwitchOobIpConfig1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Netmask != nil {
        structMap["netmask"] = s.Netmask
    }
    if s.Network != nil {
        structMap["network"] = s.Network
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.UseMgmtVrf != nil {
        structMap["use_mgmt_vrf"] = s.UseMgmtVrf
    }
    if s.UseMgmtVrfForHostOut != nil {
        structMap["use_mgmt_vrf_for_host_out"] = s.UseMgmtVrfForHostOut
    }
    if s.Node1 != nil {
        structMap["node1"] = s.Node1.toMap()
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchOobIpConfig1.
// It customizes the JSON unmarshaling process for SwitchOobIpConfig1 objects.
func (s *SwitchOobIpConfig1) UnmarshalJSON(input []byte) error {
    var temp switchOobIpConfig1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway", "ip", "netmask", "network", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out", "node1", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Gateway = temp.Gateway
    s.Ip = temp.Ip
    s.Netmask = temp.Netmask
    s.Network = temp.Network
    s.Type = temp.Type
    s.UseMgmtVrf = temp.UseMgmtVrf
    s.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
    s.Node1 = temp.Node1
    s.VlanId = temp.VlanId
    return nil
}

// switchOobIpConfig1 is a temporary struct used for validating the fields of SwitchOobIpConfig1.
type switchOobIpConfig1  struct {
    Gateway              *string                  `json:"gateway,omitempty"`
    Ip                   *string                  `json:"ip,omitempty"`
    Netmask              *string                  `json:"netmask,omitempty"`
    Network              *string                  `json:"network,omitempty"`
    Type                 *IpTypeEnum              `json:"type,omitempty"`
    UseMgmtVrf           *bool                    `json:"use_mgmt_vrf,omitempty"`
    UseMgmtVrfForHostOut *bool                    `json:"use_mgmt_vrf_for_host_out,omitempty"`
    Node1                *GatewayOobIpConfigNode1 `json:"node1,omitempty"`
    VlanId               *string                  `json:"vlan_id,omitempty"`
}
