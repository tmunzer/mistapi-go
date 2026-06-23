// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// GatewayPortConfig represents a GatewayPortConfig struct.
// Gateway port configuration for LAN, WAN, tunnel, and HA interfaces
type GatewayPortConfig struct {
	// If `aggregated`==`true`. To disable LCP support for the AE interface
	AeDisableLacp *bool `json:"ae_disable_lacp,omitempty"`
	// If `aggregated`==`true`. Users could force to use the designated AE name (must be an integer between 0 and 127)
	AeIdx Optional[string] `json:"ae_idx"`
	// For SRX only, if `aggregated`==`true`.Sets the state of the interface as UP when the peer has limited LACP capability. Use case: When a device connected to this AE port is ZTPing for the first time, it will not have LACP configured on the other end. **Note:** Turning this on will enable force-up on one of the interfaces in the bundle only
	AeLacpForceUp *bool `json:"ae_lacp_force_up,omitempty"`
	// Whether the port participates in an aggregated Ethernet interface
	Aggregated *bool `json:"aggregated,omitempty"`
	// To generate port up/down alarm, set it to true
	Critical *bool `json:"critical,omitempty"`
	// Interface Description. Can be a variable (i.e. "{{myvar}}")
	Description *string `json:"description,omitempty"`
	// Whether Ethernet autonegotiation is disabled on the port
	DisableAutoneg *bool `json:"disable_autoneg,omitempty"`
	// Port admin up (true) / down (false)
	Disabled *bool `json:"disabled,omitempty"`
	// if `wan_type`==`dsl`. enum: `adsl`, `vdsl`
	DslType *GatewayPortDslTypeEnum `json:"dsl_type,omitempty"`
	// If `wan_type`==`dsl`, 16 bit int
	DslVci *int `json:"dsl_vci,omitempty"`
	// If `wan_type`==`dsl`, 8 bit int
	DslVpi *int `json:"dsl_vpi,omitempty"`
	// enum: `auto`, `full`, `half`
	Duplex *GatewayPortDuplexEnum `json:"duplex,omitempty"`
	// Junos IP configuration for a gateway port interface
	IpConfig *GatewayPortConfigIpConfig `json:"ip_config,omitempty"`
	// If `wan_type`==`lte`. APN used by the LTE uplink
	LteApn *string `json:"lte_apn,omitempty"`
	// if `wan_type`==`lte`. enum: `chap`, `none`, `pap`
	LteAuth *GatewayPortLteAuthEnum `json:"lte_auth,omitempty"`
	// Whether the LTE uplink is used as a backup WAN connection
	LteBackup *bool `json:"lte_backup,omitempty"`
	// If `wan_type`==`lte`. Password used for LTE uplink authentication
	LtePassword *string `json:"lte_password,omitempty"`
	// If `wan_type`==`lte`. Username used for LTE uplink authentication
	LteUsername *string `json:"lte_username,omitempty"`
	// Layer 3 MTU configured on the port
	Mtu *int `json:"mtu,omitempty"`
	// Interface name used to derive device configuration
	Name *string `json:"name,omitempty"`
	// If `usage`==`lan`, name of the [networks]($h/Orgs%20Networks/_overview) to attach to the interface
	Networks []string `json:"networks,omitempty"`
	// For Q-in-Q. Outer VLAN ID used for QinQ encapsulation
	OuterVlanId *int `json:"outer_vlan_id,omitempty"`
	// Whether PoE output is disabled on the port
	PoeDisabled *bool `json:"poe_disabled,omitempty"`
	// Whether Perpetual PoE capabilities are enabled for a port
	PoeKeepStateWhenReboot *bool `json:"poe_keep_state_when_reboot,omitempty"`
	// Only for SRX and if `usage`==`lan`, the name of the Network to be used as the Untagged VLAN
	PortNetwork *string `json:"port_network,omitempty"`
	// Whether to preserve dscp when sending traffic over VPN (SSR-only)
	PreserveDscp *bool `json:"preserve_dscp,omitempty"`
	// If HA mode. Whether the port participates in the redundant Ethernet configuration
	Redundant *bool `json:"redundant,omitempty"`
	// If HA mode, SRX Only - support redundancy-group. 1-128 for physical SRX, 1-64 for virtual SRX
	RedundantGroup *int `json:"redundant_group,omitempty"`
	// For SRX only and if HA Mode. `-1` means it will be managed by the device. Use `>= 0` values to manage it manually. Ensure no conflicting values are assigned across all ports.
	RethIdx *GatewayPortConfigRethIdx `json:"reth_idx,omitempty"`
	// If HA mode. Node associated with the redundant Ethernet interface
	RethNode *string `json:"reth_node,omitempty"`
	// SSR only - supporting vlan-based redundancy (matching the size of `networks`)
	RethNodes []string `json:"reth_nodes,omitempty"`
	// Link speed configured on the port
	Speed *string `json:"speed,omitempty"`
	// When SSR is running as VM, this is required on certain hosting platforms
	SsrNoVirtualMac *bool `json:"ssr_no_virtual_mac,omitempty"`
	// For SSR only. Port range configured on the interface
	SvrPortRange *string `json:"svr_port_range,omitempty"`
	// Traffic shaping settings for a gateway interface or VPN path
	TrafficShaping *GatewayTrafficShaping `json:"traffic_shaping,omitempty"`
	// port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`
	Usage GatewayPortUsageEnum `json:"usage"`
	// If WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}")
	VlanId *GatewayPortVlanIdWithVariable `json:"vlan_id,omitempty"`
	// Property key is the VPN name
	VpnPaths map[string]GatewayPortVpnPath `json:"vpn_paths,omitempty"`
	// Only when `wan_type`==`broadband`. enum: `default`, `max`, `recommended`
	WanArpPolicer *GatewayPortWanArpPolicerEnum `json:"wan_arp_policer,omitempty"`
	// Only if `usage`==`wan`, optional. If spoke should reach this port by a different IP
	WanExtIp *string `json:"wan_ext_ip,omitempty"`
	// Only if `usage`==`wan`, optional. If spoke should reach this port by a different IPv6
	WanExtIp6 *string `json:"wan_ext_ip6,omitempty"`
	// Only if `usage`==`wan`. Property Key is the destination CIDR (e.g. "100.100.100.0/24")
	WanExtraRoutes map[string]WanExtraRoutes `json:"wan_extra_routes,omitempty"`
	// Only if `usage`==`wan`. Property Key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
	WanExtraRoutes6 map[string]WanExtraRoutes6 `json:"wan_extra_routes6,omitempty"`
	// Only if `usage`==`wan`. If some networks are connected to this WAN port, it can be added here so policies can be defined
	WanNetworks []string `json:"wan_networks,omitempty"`
	// Only if `usage`==`wan`. WAN health probe override for this gateway port
	WanProbeOverride *GatewayWanProbeOverride `json:"wan_probe_override,omitempty"`
	// Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip
	WanSourceNat *GatewayPortWanSourceNat `json:"wan_source_nat,omitempty"`
	// Controls whether Marvis/scheduler can run speedtest on this port. enum: `auto`, `enabled`, `disabled`
	WanSpeedtestMode *GatewayPortConfigWanSpeedtestModeEnum `json:"wan_speedtest_mode,omitempty"`
	// Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`
	WanType              *GatewayPortWanTypeEnum `json:"wan_type,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortConfig) String() string {
	return fmt.Sprintf(
		"GatewayPortConfig[AeDisableLacp=%v, AeIdx=%v, AeLacpForceUp=%v, Aggregated=%v, Critical=%v, Description=%v, DisableAutoneg=%v, Disabled=%v, DslType=%v, DslVci=%v, DslVpi=%v, Duplex=%v, IpConfig=%v, LteApn=%v, LteAuth=%v, LteBackup=%v, LtePassword=%v, LteUsername=%v, Mtu=%v, Name=%v, Networks=%v, OuterVlanId=%v, PoeDisabled=%v, PoeKeepStateWhenReboot=%v, PortNetwork=%v, PreserveDscp=%v, Redundant=%v, RedundantGroup=%v, RethIdx=%v, RethNode=%v, RethNodes=%v, Speed=%v, SsrNoVirtualMac=%v, SvrPortRange=%v, TrafficShaping=%v, Usage=%v, VlanId=%v, VpnPaths=%v, WanArpPolicer=%v, WanExtIp=%v, WanExtIp6=%v, WanExtraRoutes=%v, WanExtraRoutes6=%v, WanNetworks=%v, WanProbeOverride=%v, WanSourceNat=%v, WanSpeedtestMode=%v, WanType=%v, AdditionalProperties=%v]",
		g.AeDisableLacp, g.AeIdx, g.AeLacpForceUp, g.Aggregated, g.Critical, g.Description, g.DisableAutoneg, g.Disabled, g.DslType, g.DslVci, g.DslVpi, g.Duplex, g.IpConfig, g.LteApn, g.LteAuth, g.LteBackup, g.LtePassword, g.LteUsername, g.Mtu, g.Name, g.Networks, g.OuterVlanId, g.PoeDisabled, g.PoeKeepStateWhenReboot, g.PortNetwork, g.PreserveDscp, g.Redundant, g.RedundantGroup, g.RethIdx, g.RethNode, g.RethNodes, g.Speed, g.SsrNoVirtualMac, g.SvrPortRange, g.TrafficShaping, g.Usage, g.VlanId, g.VpnPaths, g.WanArpPolicer, g.WanExtIp, g.WanExtIp6, g.WanExtraRoutes, g.WanExtraRoutes6, g.WanNetworks, g.WanProbeOverride, g.WanSourceNat, g.WanSpeedtestMode, g.WanType, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfig.
// It customizes the JSON marshaling process for GatewayPortConfig objects.
func (g GatewayPortConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"ae_disable_lacp", "ae_idx", "ae_lacp_force_up", "aggregated", "critical", "description", "disable_autoneg", "disabled", "dsl_type", "dsl_vci", "dsl_vpi", "duplex", "ip_config", "lte_apn", "lte_auth", "lte_backup", "lte_password", "lte_username", "mtu", "name", "networks", "outer_vlan_id", "poe_disabled", "poe_keep_state_when_reboot", "port_network", "preserve_dscp", "redundant", "redundant_group", "reth_idx", "reth_node", "reth_nodes", "speed", "ssr_no_virtual_mac", "svr_port_range", "traffic_shaping", "usage", "vlan_id", "vpn_paths", "wan_arp_policer", "wan_ext_ip", "wan_ext_ip6", "wan_extra_routes", "wan_extra_routes6", "wan_networks", "wan_probe_override", "wan_source_nat", "wan_speedtest_mode", "wan_type"); err != nil {
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
	if g.PoeKeepStateWhenReboot != nil {
		structMap["poe_keep_state_when_reboot"] = g.PoeKeepStateWhenReboot
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
	if g.RedundantGroup != nil {
		structMap["redundant_group"] = g.RedundantGroup
	}
	if g.RethIdx != nil {
		structMap["reth_idx"] = g.RethIdx.toMap()
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
	if g.WanExtIp6 != nil {
		structMap["wan_ext_ip6"] = g.WanExtIp6
	}
	if g.WanExtraRoutes != nil {
		structMap["wan_extra_routes"] = g.WanExtraRoutes
	}
	if g.WanExtraRoutes6 != nil {
		structMap["wan_extra_routes6"] = g.WanExtraRoutes6
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
	if g.WanSpeedtestMode != nil {
		structMap["wan_speedtest_mode"] = g.WanSpeedtestMode
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ae_disable_lacp", "ae_idx", "ae_lacp_force_up", "aggregated", "critical", "description", "disable_autoneg", "disabled", "dsl_type", "dsl_vci", "dsl_vpi", "duplex", "ip_config", "lte_apn", "lte_auth", "lte_backup", "lte_password", "lte_username", "mtu", "name", "networks", "outer_vlan_id", "poe_disabled", "poe_keep_state_when_reboot", "port_network", "preserve_dscp", "redundant", "redundant_group", "reth_idx", "reth_node", "reth_nodes", "speed", "ssr_no_virtual_mac", "svr_port_range", "traffic_shaping", "usage", "vlan_id", "vpn_paths", "wan_arp_policer", "wan_ext_ip", "wan_ext_ip6", "wan_extra_routes", "wan_extra_routes6", "wan_networks", "wan_probe_override", "wan_source_nat", "wan_speedtest_mode", "wan_type")
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
	g.PoeKeepStateWhenReboot = temp.PoeKeepStateWhenReboot
	g.PortNetwork = temp.PortNetwork
	g.PreserveDscp = temp.PreserveDscp
	g.Redundant = temp.Redundant
	g.RedundantGroup = temp.RedundantGroup
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
	g.WanExtIp6 = temp.WanExtIp6
	g.WanExtraRoutes = temp.WanExtraRoutes
	g.WanExtraRoutes6 = temp.WanExtraRoutes6
	g.WanNetworks = temp.WanNetworks
	g.WanProbeOverride = temp.WanProbeOverride
	g.WanSourceNat = temp.WanSourceNat
	g.WanSpeedtestMode = temp.WanSpeedtestMode
	g.WanType = temp.WanType
	return nil
}

// tempGatewayPortConfig is a temporary struct used for validating the fields of GatewayPortConfig.
type tempGatewayPortConfig struct {
	AeDisableLacp          *bool                                  `json:"ae_disable_lacp,omitempty"`
	AeIdx                  Optional[string]                       `json:"ae_idx"`
	AeLacpForceUp          *bool                                  `json:"ae_lacp_force_up,omitempty"`
	Aggregated             *bool                                  `json:"aggregated,omitempty"`
	Critical               *bool                                  `json:"critical,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	DisableAutoneg         *bool                                  `json:"disable_autoneg,omitempty"`
	Disabled               *bool                                  `json:"disabled,omitempty"`
	DslType                *GatewayPortDslTypeEnum                `json:"dsl_type,omitempty"`
	DslVci                 *int                                   `json:"dsl_vci,omitempty"`
	DslVpi                 *int                                   `json:"dsl_vpi,omitempty"`
	Duplex                 *GatewayPortDuplexEnum                 `json:"duplex,omitempty"`
	IpConfig               *GatewayPortConfigIpConfig             `json:"ip_config,omitempty"`
	LteApn                 *string                                `json:"lte_apn,omitempty"`
	LteAuth                *GatewayPortLteAuthEnum                `json:"lte_auth,omitempty"`
	LteBackup              *bool                                  `json:"lte_backup,omitempty"`
	LtePassword            *string                                `json:"lte_password,omitempty"`
	LteUsername            *string                                `json:"lte_username,omitempty"`
	Mtu                    *int                                   `json:"mtu,omitempty"`
	Name                   *string                                `json:"name,omitempty"`
	Networks               []string                               `json:"networks,omitempty"`
	OuterVlanId            *int                                   `json:"outer_vlan_id,omitempty"`
	PoeDisabled            *bool                                  `json:"poe_disabled,omitempty"`
	PoeKeepStateWhenReboot *bool                                  `json:"poe_keep_state_when_reboot,omitempty"`
	PortNetwork            *string                                `json:"port_network,omitempty"`
	PreserveDscp           *bool                                  `json:"preserve_dscp,omitempty"`
	Redundant              *bool                                  `json:"redundant,omitempty"`
	RedundantGroup         *int                                   `json:"redundant_group,omitempty"`
	RethIdx                *GatewayPortConfigRethIdx              `json:"reth_idx,omitempty"`
	RethNode               *string                                `json:"reth_node,omitempty"`
	RethNodes              []string                               `json:"reth_nodes,omitempty"`
	Speed                  *string                                `json:"speed,omitempty"`
	SsrNoVirtualMac        *bool                                  `json:"ssr_no_virtual_mac,omitempty"`
	SvrPortRange           *string                                `json:"svr_port_range,omitempty"`
	TrafficShaping         *GatewayTrafficShaping                 `json:"traffic_shaping,omitempty"`
	Usage                  *GatewayPortUsageEnum                  `json:"usage"`
	VlanId                 *GatewayPortVlanIdWithVariable         `json:"vlan_id,omitempty"`
	VpnPaths               map[string]GatewayPortVpnPath          `json:"vpn_paths,omitempty"`
	WanArpPolicer          *GatewayPortWanArpPolicerEnum          `json:"wan_arp_policer,omitempty"`
	WanExtIp               *string                                `json:"wan_ext_ip,omitempty"`
	WanExtIp6              *string                                `json:"wan_ext_ip6,omitempty"`
	WanExtraRoutes         map[string]WanExtraRoutes              `json:"wan_extra_routes,omitempty"`
	WanExtraRoutes6        map[string]WanExtraRoutes6             `json:"wan_extra_routes6,omitempty"`
	WanNetworks            []string                               `json:"wan_networks,omitempty"`
	WanProbeOverride       *GatewayWanProbeOverride               `json:"wan_probe_override,omitempty"`
	WanSourceNat           *GatewayPortWanSourceNat               `json:"wan_source_nat,omitempty"`
	WanSpeedtestMode       *GatewayPortConfigWanSpeedtestModeEnum `json:"wan_speedtest_mode,omitempty"`
	WanType                *GatewayPortWanTypeEnum                `json:"wan_type,omitempty"`
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
