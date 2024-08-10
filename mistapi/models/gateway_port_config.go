package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayPortConfig represents a GatewayPortConfig struct.
// Gateway port config
type GatewayPortConfig struct {
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

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfig.
// It customizes the JSON marshaling process for GatewayPortConfig objects.
func (g GatewayPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortConfig object to a map representation for JSON marshaling.
func (g GatewayPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Description != nil {
        structMap["description"] = g.Description
    }
    if g.DisableAutoneg != nil {
        structMap["disable_autoneg"] = g.DisableAutoneg
    }
    if g.Disabled != nil {
        structMap["disabled"] = g.Disabled
    }
    if g.DslType != nil {
        structMap["dsl_type"] = g.DslType
    }
    if g.DslVci != nil {
        structMap["dsl_vci"] = g.DslVci
    }
    if g.DslVpi != nil {
        structMap["dsl_vpi"] = g.DslVpi
    }
    if g.Duplex != nil {
        structMap["duplex"] = g.Duplex
    }
    if g.IpConfig != nil {
        structMap["ip_config"] = g.IpConfig.toMap()
    }
    if g.LteApn != nil {
        structMap["lte_apn"] = g.LteApn
    }
    if g.LteAuth != nil {
        structMap["lte_auth"] = g.LteAuth
    }
    if g.LteBackup != nil {
        structMap["lte_backup"] = g.LteBackup
    }
    if g.LtePassword != nil {
        structMap["lte_password"] = g.LtePassword
    }
    if g.LteUsername != nil {
        structMap["lte_username"] = g.LteUsername
    }
    if g.Mtu != nil {
        structMap["mtu"] = g.Mtu
    }
    if g.Name != nil {
        structMap["name"] = g.Name
    }
    if g.Networks != nil {
        structMap["networks"] = g.Networks
    }
    if g.OuterVlanId != nil {
        structMap["outer_vlan_id"] = g.OuterVlanId
    }
    if g.PoeDisabled != nil {
        structMap["poe_disabled"] = g.PoeDisabled
    }
    if g.PortNetwork != nil {
        structMap["port_network"] = g.PortNetwork
    }
    if g.PreserveDscp != nil {
        structMap["preserve_dscp"] = g.PreserveDscp
    }
    if g.Redundant != nil {
        structMap["redundant"] = g.Redundant
    }
    if g.RethIdx != nil {
        structMap["reth_idx"] = g.RethIdx
    }
    if g.RethNode != nil {
        structMap["reth_node"] = g.RethNode
    }
    if g.RethNodes != nil {
        structMap["reth_nodes"] = g.RethNodes
    }
    if g.Speed != nil {
        structMap["speed"] = g.Speed
    }
    if g.SsrNoVirtualMac != nil {
        structMap["ssr_no_virtual_mac"] = g.SsrNoVirtualMac
    }
    if g.SvrPortRange != nil {
        structMap["svr_port_range"] = g.SvrPortRange
    }
    if g.TrafficShaping != nil {
        structMap["traffic_shaping"] = g.TrafficShaping.toMap()
    }
    structMap["usage"] = g.Usage
    if g.VlanId != nil {
        structMap["vlan_id"] = g.VlanId
    }
    if g.VpnPaths != nil {
        structMap["vpn_paths"] = g.VpnPaths
    }
    if g.WanArpPolicer != nil {
        structMap["wan_arp_policer"] = g.WanArpPolicer
    }
    if g.WanExtIp != nil {
        structMap["wan_ext_ip"] = g.WanExtIp
    }
    if g.WanSourceNat != nil {
        structMap["wan_source_nat"] = g.WanSourceNat.toMap()
    }
    if g.WanType != nil {
        structMap["wan_type"] = g.WanType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortConfig.
// It customizes the JSON unmarshaling process for GatewayPortConfig objects.
func (g *GatewayPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortConfig
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
    
    g.AdditionalProperties = additionalProperties
    g.Description = temp.Description
    g.DisableAutoneg = temp.DisableAutoneg
    g.Disabled = temp.Disabled
    g.DslType = temp.DslType
    g.DslVci = temp.DslVci
    g.DslVpi = temp.DslVpi
    g.Duplex = temp.Duplex
    g.IpConfig = temp.IpConfig
    g.LteApn = temp.LteApn
    g.LteAuth = temp.LteAuth
    g.LteBackup = temp.LteBackup
    g.LtePassword = temp.LtePassword
    g.LteUsername = temp.LteUsername
    g.Mtu = temp.Mtu
    g.Name = temp.Name
    g.Networks = temp.Networks
    g.OuterVlanId = temp.OuterVlanId
    g.PoeDisabled = temp.PoeDisabled
    g.PortNetwork = temp.PortNetwork
    g.PreserveDscp = temp.PreserveDscp
    g.Redundant = temp.Redundant
    g.RethIdx = temp.RethIdx
    g.RethNode = temp.RethNode
    g.RethNodes = temp.RethNodes
    g.Speed = temp.Speed
    g.SsrNoVirtualMac = temp.SsrNoVirtualMac
    g.SvrPortRange = temp.SvrPortRange
    g.TrafficShaping = temp.TrafficShaping
    g.Usage = *temp.Usage
    g.VlanId = temp.VlanId
    g.VpnPaths = temp.VpnPaths
    g.WanArpPolicer = temp.WanArpPolicer
    g.WanExtIp = temp.WanExtIp
    g.WanSourceNat = temp.WanSourceNat
    g.WanType = temp.WanType
    return nil
}

// tempGatewayPortConfig is a temporary struct used for validating the fields of GatewayPortConfig.
type tempGatewayPortConfig  struct {
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

func (g *tempGatewayPortConfig) validate() error {
    var errs []string
    if g.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `gateway_port_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
