package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PortConfig1 represents a PortConfig1 struct.
// Property key is the interface(s) name (e.g. "eth1,eth2")
type PortConfig1 struct {
    Description          *string                       `json:"description,omitempty"`
    DisableAutoneg       *bool                         `json:"disable_autoneg,omitempty"`
    // port admin up (true) / down (false)
    Disabled             *bool                         `json:"disabled,omitempty"`
    // if `wan_type`==`lte`. enum: `adsl`, `vdsl`
    DslType              *GatewayPortDslTypeEnum       `json:"dsl_type,omitempty"`
    // if `wan_type`==`dsl`
    // 16 bit int
    DslVci               *int                          `json:"dsl_vci,omitempty"`
    // if `wan_type`==`dsl`
    // 8 bit int
    DslVpi               *int                          `json:"dsl_vpi,omitempty"`
    // enum: `auto`, `full`, `half`
    Duplex               *GatewayPortDuplexEnum        `json:"duplex,omitempty"`
    // Junos IP Config
    IpConfig             *GatewayPortConfigIpConfig    `json:"ip_config,omitempty"`
    // if `wan_type`==`lte`
    LteApn               *string                       `json:"lte_apn,omitempty"`
    // if `wan_type`==`lte`. enum: `chap`, `none`, `pap`
    LteAuth              *GatewayPortLteAuthEnum       `json:"lte_auth,omitempty"`
    LteBackup            *bool                         `json:"lte_backup,omitempty"`
    // if `wan_type`==`lte`
    LtePassword          *string                       `json:"lte_password,omitempty"`
    // if `wan_type`==`lte`
    LteUsername          *string                       `json:"lte_username,omitempty"`
    Mtu                  *int                          `json:"mtu,omitempty"`
    // name that we'll use to derive config
    Name                 *string                       `json:"name,omitempty"`
    // if `usage`==`lan`
    Networks             []string                      `json:"networks,omitempty"`
    // for Q-in-Q
    OuterVlanId          *int                          `json:"outer_vlan_id,omitempty"`
    PoeDisabled          *bool                         `json:"poe_disabled,omitempty"`
    // if `usage`==`lan`
    PortNetwork          *string                       `json:"port_network,omitempty"`
    // whether to preserve dscp when sending traffic over VPN (SSR-only)
    PreserveDscp         *bool                         `json:"preserve_dscp,omitempty"`
    // if HA mode
    Redundant            *bool                         `json:"redundant,omitempty"`
    // if HA mode
    RethIdx              *int                          `json:"reth_idx,omitempty"`
    // if HA mode
    RethNode             *string                       `json:"reth_node,omitempty"`
    // SSR only - supporting vlan-based redundancy (matching the size of `networks`)
    RethNodes            []string                      `json:"reth_nodes,omitempty"`
    Speed                *string                       `json:"speed,omitempty"`
    // when SSR is running as VM, this is required on certain hosting platforms
    SsrNoVirtualMac      *bool                         `json:"ssr_no_virtual_mac,omitempty"`
    // for SSR only
    SvrPortRange         *string                       `json:"svr_port_range,omitempty"`
    TrafficShaping       *GatewayTrafficShaping        `json:"traffic_shaping,omitempty"`
    // port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`
    Usage                GatewayPortUsageEnum          `json:"usage"`
    // if WAN interface is on a VLAN
    VlanId               *int                          `json:"vlan_id,omitempty"`
    VpnPaths             map[string]GatewayPortVpnPath `json:"vpn_paths,omitempty"`
    // when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`
    WanArpPolicer        *GatewayPortWanArpPolicerEnum `json:"wan_arp_policer,omitempty"`
    // optional, if spoke should reach this port by a different IP
    WanExtIp             *string                       `json:"wan_ext_ip,omitempty"`
    // optional, by default, source-NAT is performed on all WAN Ports using the interface-ip
    WanSourceNat         *GatewayPortWanSourceNat      `json:"wan_source_nat,omitempty"`
    // if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`
    WanType              *GatewayPortWanTypeEnum       `json:"wan_type,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PortConfig1.
// It customizes the JSON marshaling process for PortConfig1 objects.
func (p PortConfig1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PortConfig1 object to a map representation for JSON marshaling.
func (p PortConfig1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Description != nil {
        structMap["description"] = p.Description
    }
    if p.DisableAutoneg != nil {
        structMap["disable_autoneg"] = p.DisableAutoneg
    }
    if p.Disabled != nil {
        structMap["disabled"] = p.Disabled
    }
    if p.DslType != nil {
        structMap["dsl_type"] = p.DslType
    }
    if p.DslVci != nil {
        structMap["dsl_vci"] = p.DslVci
    }
    if p.DslVpi != nil {
        structMap["dsl_vpi"] = p.DslVpi
    }
    if p.Duplex != nil {
        structMap["duplex"] = p.Duplex
    }
    if p.IpConfig != nil {
        structMap["ip_config"] = p.IpConfig.toMap()
    }
    if p.LteApn != nil {
        structMap["lte_apn"] = p.LteApn
    }
    if p.LteAuth != nil {
        structMap["lte_auth"] = p.LteAuth
    }
    if p.LteBackup != nil {
        structMap["lte_backup"] = p.LteBackup
    }
    if p.LtePassword != nil {
        structMap["lte_password"] = p.LtePassword
    }
    if p.LteUsername != nil {
        structMap["lte_username"] = p.LteUsername
    }
    if p.Mtu != nil {
        structMap["mtu"] = p.Mtu
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Networks != nil {
        structMap["networks"] = p.Networks
    }
    if p.OuterVlanId != nil {
        structMap["outer_vlan_id"] = p.OuterVlanId
    }
    if p.PoeDisabled != nil {
        structMap["poe_disabled"] = p.PoeDisabled
    }
    if p.PortNetwork != nil {
        structMap["port_network"] = p.PortNetwork
    }
    if p.PreserveDscp != nil {
        structMap["preserve_dscp"] = p.PreserveDscp
    }
    if p.Redundant != nil {
        structMap["redundant"] = p.Redundant
    }
    if p.RethIdx != nil {
        structMap["reth_idx"] = p.RethIdx
    }
    if p.RethNode != nil {
        structMap["reth_node"] = p.RethNode
    }
    if p.RethNodes != nil {
        structMap["reth_nodes"] = p.RethNodes
    }
    if p.Speed != nil {
        structMap["speed"] = p.Speed
    }
    if p.SsrNoVirtualMac != nil {
        structMap["ssr_no_virtual_mac"] = p.SsrNoVirtualMac
    }
    if p.SvrPortRange != nil {
        structMap["svr_port_range"] = p.SvrPortRange
    }
    if p.TrafficShaping != nil {
        structMap["traffic_shaping"] = p.TrafficShaping.toMap()
    }
    structMap["usage"] = p.Usage
    if p.VlanId != nil {
        structMap["vlan_id"] = p.VlanId
    }
    if p.VpnPaths != nil {
        structMap["vpn_paths"] = p.VpnPaths
    }
    if p.WanArpPolicer != nil {
        structMap["wan_arp_policer"] = p.WanArpPolicer
    }
    if p.WanExtIp != nil {
        structMap["wan_ext_ip"] = p.WanExtIp
    }
    if p.WanSourceNat != nil {
        structMap["wan_source_nat"] = p.WanSourceNat.toMap()
    }
    if p.WanType != nil {
        structMap["wan_type"] = p.WanType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PortConfig1.
// It customizes the JSON unmarshaling process for PortConfig1 objects.
func (p *PortConfig1) UnmarshalJSON(input []byte) error {
    var temp tempPortConfig1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "description", "disable_autoneg", "disabled", "dsl_type", "dsl_vci", "dsl_vpi", "duplex", "ip_config", "lte_apn", "lte_auth", "lte_backup", "lte_password", "lte_username", "mtu", "name", "networks", "outer_vlan_id", "poe_disabled", "port_network", "preserve_dscp", "redundant", "reth_idx", "reth_node", "reth_nodes", "speed", "ssr_no_virtual_mac", "svr_port_range", "traffic_shaping", "usage", "vlan_id", "vpn_paths", "wan_arp_policer", "wan_ext_ip", "wan_source_nat", "wan_type")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Description = temp.Description
    p.DisableAutoneg = temp.DisableAutoneg
    p.Disabled = temp.Disabled
    p.DslType = temp.DslType
    p.DslVci = temp.DslVci
    p.DslVpi = temp.DslVpi
    p.Duplex = temp.Duplex
    p.IpConfig = temp.IpConfig
    p.LteApn = temp.LteApn
    p.LteAuth = temp.LteAuth
    p.LteBackup = temp.LteBackup
    p.LtePassword = temp.LtePassword
    p.LteUsername = temp.LteUsername
    p.Mtu = temp.Mtu
    p.Name = temp.Name
    p.Networks = temp.Networks
    p.OuterVlanId = temp.OuterVlanId
    p.PoeDisabled = temp.PoeDisabled
    p.PortNetwork = temp.PortNetwork
    p.PreserveDscp = temp.PreserveDscp
    p.Redundant = temp.Redundant
    p.RethIdx = temp.RethIdx
    p.RethNode = temp.RethNode
    p.RethNodes = temp.RethNodes
    p.Speed = temp.Speed
    p.SsrNoVirtualMac = temp.SsrNoVirtualMac
    p.SvrPortRange = temp.SvrPortRange
    p.TrafficShaping = temp.TrafficShaping
    p.Usage = *temp.Usage
    p.VlanId = temp.VlanId
    p.VpnPaths = temp.VpnPaths
    p.WanArpPolicer = temp.WanArpPolicer
    p.WanExtIp = temp.WanExtIp
    p.WanSourceNat = temp.WanSourceNat
    p.WanType = temp.WanType
    return nil
}

// tempPortConfig1 is a temporary struct used for validating the fields of PortConfig1.
type tempPortConfig1  struct {
    Description     *string                       `json:"description,omitempty"`
    DisableAutoneg  *bool                         `json:"disable_autoneg,omitempty"`
    Disabled        *bool                         `json:"disabled,omitempty"`
    DslType         *GatewayPortDslTypeEnum       `json:"dsl_type,omitempty"`
    DslVci          *int                          `json:"dsl_vci,omitempty"`
    DslVpi          *int                          `json:"dsl_vpi,omitempty"`
    Duplex          *GatewayPortDuplexEnum        `json:"duplex,omitempty"`
    IpConfig        *GatewayPortConfigIpConfig    `json:"ip_config,omitempty"`
    LteApn          *string                       `json:"lte_apn,omitempty"`
    LteAuth         *GatewayPortLteAuthEnum       `json:"lte_auth,omitempty"`
    LteBackup       *bool                         `json:"lte_backup,omitempty"`
    LtePassword     *string                       `json:"lte_password,omitempty"`
    LteUsername     *string                       `json:"lte_username,omitempty"`
    Mtu             *int                          `json:"mtu,omitempty"`
    Name            *string                       `json:"name,omitempty"`
    Networks        []string                      `json:"networks,omitempty"`
    OuterVlanId     *int                          `json:"outer_vlan_id,omitempty"`
    PoeDisabled     *bool                         `json:"poe_disabled,omitempty"`
    PortNetwork     *string                       `json:"port_network,omitempty"`
    PreserveDscp    *bool                         `json:"preserve_dscp,omitempty"`
    Redundant       *bool                         `json:"redundant,omitempty"`
    RethIdx         *int                          `json:"reth_idx,omitempty"`
    RethNode        *string                       `json:"reth_node,omitempty"`
    RethNodes       []string                      `json:"reth_nodes,omitempty"`
    Speed           *string                       `json:"speed,omitempty"`
    SsrNoVirtualMac *bool                         `json:"ssr_no_virtual_mac,omitempty"`
    SvrPortRange    *string                       `json:"svr_port_range,omitempty"`
    TrafficShaping  *GatewayTrafficShaping        `json:"traffic_shaping,omitempty"`
    Usage           *GatewayPortUsageEnum         `json:"usage"`
    VlanId          *int                          `json:"vlan_id,omitempty"`
    VpnPaths        map[string]GatewayPortVpnPath `json:"vpn_paths,omitempty"`
    WanArpPolicer   *GatewayPortWanArpPolicerEnum `json:"wan_arp_policer,omitempty"`
    WanExtIp        *string                       `json:"wan_ext_ip,omitempty"`
    WanSourceNat    *GatewayPortWanSourceNat      `json:"wan_source_nat,omitempty"`
    WanType         *GatewayPortWanTypeEnum       `json:"wan_type,omitempty"`
}

func (p *tempPortConfig1) validate() error {
    var errs []string
    if p.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `PortConfig1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
