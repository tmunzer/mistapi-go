package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// GatewayPortConfig represents a GatewayPortConfig struct.
// Gateway port config
type GatewayPortConfig struct {
    // if `aggregated`==`true`. To disable LCP support for the AE interface
    AeDisableLacp        *bool                          `json:"ae_disable_lacp,omitempty"`
    // if `aggregated`==`true`. Users could force to use the designated AE name (must be an integer between 0 and 127)
    AeIdx                Optional[string]               `json:"ae_idx"`
    // For SRX Only, if `aggregated`==`true`.Sets the state of the interface as UP when the peer has limited LACP capability. Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end. **Note:** Turning this on will enable force-up on one of the interfaces in the bundle only
    AeLacpForceUp        *bool                          `json:"ae_lacp_force_up,omitempty"`
    Aggregated           *bool                          `json:"aggregated,omitempty"`
    // if want to generate port up/down alarm, set it to true
    Critical             *bool                          `json:"critical,omitempty"`
    // Interface Description. Can be a variable (i.e. "{{myvar}}")
    Description          *string                        `json:"description,omitempty"`
    DisableAutoneg       *bool                          `json:"disable_autoneg,omitempty"`
    // port admin up (true) / down (false)
    Disabled             *bool                          `json:"disabled,omitempty"`
    // if `wan_type`==`dsl`. enum: `adsl`, `vdsl`
    DslType              *GatewayPortDslTypeEnum        `json:"dsl_type,omitempty"`
    // if `wan_type`==`dsl`
    // 16 bit int
    DslVci               *int                           `json:"dsl_vci,omitempty"`
    // if `wan_type`==`dsl`
    // 8 bit int
    DslVpi               *int                           `json:"dsl_vpi,omitempty"`
    // enum: `auto`, `full`, `half`
    Duplex               *GatewayPortDuplexEnum         `json:"duplex,omitempty"`
    // Junos IP Config
    IpConfig             *GatewayPortConfigIpConfig     `json:"ip_config,omitempty"`
    // if `wan_type`==`lte`
    LteApn               *string                        `json:"lte_apn,omitempty"`
    // if `wan_type`==`lte`. enum: `chap`, `none`, `pap`
    LteAuth              *GatewayPortLteAuthEnum        `json:"lte_auth,omitempty"`
    LteBackup            *bool                          `json:"lte_backup,omitempty"`
    // if `wan_type`==`lte`
    LtePassword          *string                        `json:"lte_password,omitempty"`
    // if `wan_type`==`lte`
    LteUsername          *string                        `json:"lte_username,omitempty"`
    Mtu                  *int                           `json:"mtu,omitempty"`
    // name that we'll use to derive config
    Name                 *string                        `json:"name,omitempty"`
    // if `usage`==`lan`, name of the [networks]($h/Orgs%20Networks/_overview) to attach to the interface
    Networks             []string                       `json:"networks,omitempty"`
    // for Q-in-Q
    OuterVlanId          *int                           `json:"outer_vlan_id,omitempty"`
    PoeDisabled          *bool                          `json:"poe_disabled,omitempty"`
    // Only for SRX and if `usage`==`lan`, the Untagged VLAN Network
    PortNetwork          *string                        `json:"port_network,omitempty"`
    // whether to preserve dscp when sending traffic over VPN (SSR-only)
    PreserveDscp         *bool                          `json:"preserve_dscp,omitempty"`
    // if HA mode
    Redundant            *bool                          `json:"redundant,omitempty"`
    // if HA mode
    RethIdx              *int                           `json:"reth_idx,omitempty"`
    // if HA mode
    RethNode             *string                        `json:"reth_node,omitempty"`
    // SSR only - supporting vlan-based redundancy (matching the size of `networks`)
    RethNodes            []string                       `json:"reth_nodes,omitempty"`
    Speed                *string                        `json:"speed,omitempty"`
    // when SSR is running as VM, this is required on certain hosting platforms
    SsrNoVirtualMac      *bool                          `json:"ssr_no_virtual_mac,omitempty"`
    // for SSR only
    SvrPortRange         *string                        `json:"svr_port_range,omitempty"`
    TrafficShaping       *GatewayTrafficShaping         `json:"traffic_shaping,omitempty"`
    // port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`
    Usage                GatewayPortUsageEnum           `json:"usage"`
    // if WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}")
    VlanId               *GatewayPortVlanIdWithVariable `json:"vlan_id,omitempty"`
    // Property key is the VPN name
    VpnPaths             map[string]GatewayPortVpnPath  `json:"vpn_paths,omitempty"`
    // Only when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`
    WanArpPolicer        *GatewayPortWanArpPolicerEnum  `json:"wan_arp_policer,omitempty"`
    // Only if `usage`==`wan`, optional. If spoke should reach this port by a different IP
    WanExtIp             *string                        `json:"wan_ext_ip,omitempty"`
    // Only if `usage`==`wan`. Property Key is the destianation CIDR (e.g "100.100.100.0/24")
    WanExtraRoutes       map[string]WanExtraRoutes      `json:"wan_extra_routes,omitempty"`
    // Only if `usage`==`wan`. If some networks are connected to this WAN port, it can be added here so policies can be defined
    WanNetworks          []string                       `json:"wan_networks,omitempty"`
    // Only if `usage`==`wan`
    WanProbeOverride     *GatewayWanProbeOverride       `json:"wan_probe_override,omitempty"`
    // Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip
    WanSourceNat         *GatewayPortWanSourceNat       `json:"wan_source_nat,omitempty"`
    // Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`
    WanType              *GatewayPortWanTypeEnum        `json:"wan_type,omitempty"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortConfig) String() string {
    return fmt.Sprintf(
    	"GatewayPortConfig[AeDisableLacp=%v, AeIdx=%v, AeLacpForceUp=%v, Aggregated=%v, Critical=%v, Description=%v, DisableAutoneg=%v, Disabled=%v, DslType=%v, DslVci=%v, DslVpi=%v, Duplex=%v, IpConfig=%v, LteApn=%v, LteAuth=%v, LteBackup=%v, LtePassword=%v, LteUsername=%v, Mtu=%v, Name=%v, Networks=%v, OuterVlanId=%v, PoeDisabled=%v, PortNetwork=%v, PreserveDscp=%v, Redundant=%v, RethIdx=%v, RethNode=%v, RethNodes=%v, Speed=%v, SsrNoVirtualMac=%v, SvrPortRange=%v, TrafficShaping=%v, Usage=%v, VlanId=%v, VpnPaths=%v, WanArpPolicer=%v, WanExtIp=%v, WanExtraRoutes=%v, WanNetworks=%v, WanProbeOverride=%v, WanSourceNat=%v, WanType=%v, AdditionalProperties=%v]",
    	g.AeDisableLacp, g.AeIdx, g.AeLacpForceUp, g.Aggregated, g.Critical, g.Description, g.DisableAutoneg, g.Disabled, g.DslType, g.DslVci, g.DslVpi, g.Duplex, g.IpConfig, g.LteApn, g.LteAuth, g.LteBackup, g.LtePassword, g.LteUsername, g.Mtu, g.Name, g.Networks, g.OuterVlanId, g.PoeDisabled, g.PortNetwork, g.PreserveDscp, g.Redundant, g.RethIdx, g.RethNode, g.RethNodes, g.Speed, g.SsrNoVirtualMac, g.SvrPortRange, g.TrafficShaping, g.Usage, g.VlanId, g.VpnPaths, g.WanArpPolicer, g.WanExtIp, g.WanExtraRoutes, g.WanNetworks, g.WanProbeOverride, g.WanSourceNat, g.WanType, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfig.
// It customizes the JSON marshaling process for GatewayPortConfig objects.
func (g GatewayPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "ae_disable_lacp", "ae_idx", "ae_lacp_force_up", "aggregated", "critical", "description", "disable_autoneg", "disabled", "dsl_type", "dsl_vci", "dsl_vpi", "duplex", "ip_config", "lte_apn", "lte_auth", "lte_backup", "lte_password", "lte_username", "mtu", "name", "networks", "outer_vlan_id", "poe_disabled", "port_network", "preserve_dscp", "redundant", "reth_idx", "reth_node", "reth_nodes", "speed", "ssr_no_virtual_mac", "svr_port_range", "traffic_shaping", "usage", "vlan_id", "vpn_paths", "wan_arp_policer", "wan_ext_ip", "wan_extra_routes", "wan_networks", "wan_probe_override", "wan_source_nat", "wan_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortConfig object to a map representation for JSON marshaling.
func (g GatewayPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.AeDisableLacp != nil {
        structMap["ae_disable_lacp"] = g.AeDisableLacp
    }
    if g.AeIdx.IsValueSet() {
        if g.AeIdx.Value() != nil {
            structMap["ae_idx"] = g.AeIdx.Value()
        } else {
            structMap["ae_idx"] = nil
        }
    }
    if g.AeLacpForceUp != nil {
        structMap["ae_lacp_force_up"] = g.AeLacpForceUp
    }
    if g.Aggregated != nil {
        structMap["aggregated"] = g.Aggregated
    }
    if g.Critical != nil {
        structMap["critical"] = g.Critical
    }
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
        structMap["vlan_id"] = g.VlanId.toMap()
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
    if g.WanExtraRoutes != nil {
        structMap["wan_extra_routes"] = g.WanExtraRoutes
    }
    if g.WanNetworks != nil {
        structMap["wan_networks"] = g.WanNetworks
    }
    if g.WanProbeOverride != nil {
        structMap["wan_probe_override"] = g.WanProbeOverride.toMap()
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ae_disable_lacp", "ae_idx", "ae_lacp_force_up", "aggregated", "critical", "description", "disable_autoneg", "disabled", "dsl_type", "dsl_vci", "dsl_vpi", "duplex", "ip_config", "lte_apn", "lte_auth", "lte_backup", "lte_password", "lte_username", "mtu", "name", "networks", "outer_vlan_id", "poe_disabled", "port_network", "preserve_dscp", "redundant", "reth_idx", "reth_node", "reth_nodes", "speed", "ssr_no_virtual_mac", "svr_port_range", "traffic_shaping", "usage", "vlan_id", "vpn_paths", "wan_arp_policer", "wan_ext_ip", "wan_extra_routes", "wan_networks", "wan_probe_override", "wan_source_nat", "wan_type")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.AeDisableLacp = temp.AeDisableLacp
    g.AeIdx = temp.AeIdx
    g.AeLacpForceUp = temp.AeLacpForceUp
    g.Aggregated = temp.Aggregated
    g.Critical = temp.Critical
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
    g.WanExtraRoutes = temp.WanExtraRoutes
    g.WanNetworks = temp.WanNetworks
    g.WanProbeOverride = temp.WanProbeOverride
    g.WanSourceNat = temp.WanSourceNat
    g.WanType = temp.WanType
    return nil
}

// tempGatewayPortConfig is a temporary struct used for validating the fields of GatewayPortConfig.
type tempGatewayPortConfig  struct {
    AeDisableLacp    *bool                          `json:"ae_disable_lacp,omitempty"`
    AeIdx            Optional[string]               `json:"ae_idx"`
    AeLacpForceUp    *bool                          `json:"ae_lacp_force_up,omitempty"`
    Aggregated       *bool                          `json:"aggregated,omitempty"`
    Critical         *bool                          `json:"critical,omitempty"`
    Description      *string                        `json:"description,omitempty"`
    DisableAutoneg   *bool                          `json:"disable_autoneg,omitempty"`
    Disabled         *bool                          `json:"disabled,omitempty"`
    DslType          *GatewayPortDslTypeEnum        `json:"dsl_type,omitempty"`
    DslVci           *int                           `json:"dsl_vci,omitempty"`
    DslVpi           *int                           `json:"dsl_vpi,omitempty"`
    Duplex           *GatewayPortDuplexEnum         `json:"duplex,omitempty"`
    IpConfig         *GatewayPortConfigIpConfig     `json:"ip_config,omitempty"`
    LteApn           *string                        `json:"lte_apn,omitempty"`
    LteAuth          *GatewayPortLteAuthEnum        `json:"lte_auth,omitempty"`
    LteBackup        *bool                          `json:"lte_backup,omitempty"`
    LtePassword      *string                        `json:"lte_password,omitempty"`
    LteUsername      *string                        `json:"lte_username,omitempty"`
    Mtu              *int                           `json:"mtu,omitempty"`
    Name             *string                        `json:"name,omitempty"`
    Networks         []string                       `json:"networks,omitempty"`
    OuterVlanId      *int                           `json:"outer_vlan_id,omitempty"`
    PoeDisabled      *bool                          `json:"poe_disabled,omitempty"`
    PortNetwork      *string                        `json:"port_network,omitempty"`
    PreserveDscp     *bool                          `json:"preserve_dscp,omitempty"`
    Redundant        *bool                          `json:"redundant,omitempty"`
    RethIdx          *int                           `json:"reth_idx,omitempty"`
    RethNode         *string                        `json:"reth_node,omitempty"`
    RethNodes        []string                       `json:"reth_nodes,omitempty"`
    Speed            *string                        `json:"speed,omitempty"`
    SsrNoVirtualMac  *bool                          `json:"ssr_no_virtual_mac,omitempty"`
    SvrPortRange     *string                        `json:"svr_port_range,omitempty"`
    TrafficShaping   *GatewayTrafficShaping         `json:"traffic_shaping,omitempty"`
    Usage            *GatewayPortUsageEnum          `json:"usage"`
    VlanId           *GatewayPortVlanIdWithVariable `json:"vlan_id,omitempty"`
    VpnPaths         map[string]GatewayPortVpnPath  `json:"vpn_paths,omitempty"`
    WanArpPolicer    *GatewayPortWanArpPolicerEnum  `json:"wan_arp_policer,omitempty"`
    WanExtIp         *string                        `json:"wan_ext_ip,omitempty"`
    WanExtraRoutes   map[string]WanExtraRoutes      `json:"wan_extra_routes,omitempty"`
    WanNetworks      []string                       `json:"wan_networks,omitempty"`
    WanProbeOverride *GatewayWanProbeOverride       `json:"wan_probe_override,omitempty"`
    WanSourceNat     *GatewayPortWanSourceNat       `json:"wan_source_nat,omitempty"`
    WanType          *GatewayPortWanTypeEnum        `json:"wan_type,omitempty"`
}

func (g *tempGatewayPortConfig) validate() error {
    var errs []string
    if g.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `gateway_port_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
