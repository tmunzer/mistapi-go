// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// SwitchPortUsage represents a SwitchPortUsage struct.
// Junos port usages
type SwitchPortUsage struct {
	// Only if `mode`==`trunk`. Whether to trunk all network/vlans
	AllNetworks *bool `json:"all_networks,omitempty"`
	// Only applies when `mode`!=`dynamic`. Controls whether DHCP server traffic is allowed on ports using this configuration if DHCP snooping is enabled. This is a tri-state setting; `true`: ports become trusted ports allowing DHCP server traffic, `false`: ports become untrusted blocking DHCP server traffic, undefined: use system defaults (access ports default to untrusted, trunk ports default to trusted).
	AllowDhcpd *bool `json:"allow_dhcpd,omitempty"`
	// Only if `mode`!=`dynamic`
	AllowMultipleSupplicants *bool `json:"allow_multiple_supplicants,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Bypass auth for known clients if set to true when RADIUS server is down
	BypassAuthWhenServerDown *bool `json:"bypass_auth_when_server_down,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`=`dot1x`. Bypass auth for all (including unknown clients) if set to true when RADIUS server is down
	BypassAuthWhenServerDownForUnknownClient *bool `json:"bypass_auth_when_server_down_for_unknown_client,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Bypass auth for VOIP if set to true when RADIUS server is down
	BypassAuthWhenServerDownForVoip *bool `json:"bypass_auth_when_server_down_for_voip,omitempty"`
	// Only if `mode`!=`dynamic`. To be used together with `isolation` under networks. Signaling that this port connects to the networks isolated but wired clients belong to the same community can talk to each other
	CommunityVlanId *int `json:"community_vlan_id,omitempty"`
	// Only if `mode`!=`dynamic`
	Description *string `json:"description,omitempty"`
	// Only if `mode`!=`dynamic`. If speed and duplex are specified, whether to disable autonegotiation
	DisableAutoneg *bool `json:"disable_autoneg,omitempty"`
	// Only if `mode`!=`dynamic`. Whether the port is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Only if `mode`!=`dynamic`. Link connection mode. enum: `auto`, `full`, `half`
	Duplex *SwitchPortUsageDuplexEnum `json:"duplex,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return
	DynamicVlanNetworks []string `json:"dynamic_vlan_networks,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Whether to enable MAC Auth
	EnableMacAuth *bool `json:"enable_mac_auth,omitempty"`
	// Only if `mode`!=`dynamic`
	EnableQos *bool `json:"enable_qos,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed)
	GuestNetwork Optional[string] `json:"guest_network"`
	// Only if `mode`!=`dynamic`. `inter_isolation_network_link` is used together with `isolation` under networks, signaling that this port connects to isolated networks
	InterIsolationNetworkLink *bool `json:"inter_isolation_network_link,omitempty"`
	// Only if `mode`!=`dynamic`. `inter_switch_link` is used together with `isolation` under networks. NOTE: `inter_switch_link` works only between Juniper devices. This has to be applied to both ports connected together
	InterSwitchLink *bool `json:"inter_switch_link,omitempty"`
	// Only if `mode`!=`dynamic` and `enable_mac_auth`==`true`
	MacAuthOnly *bool `json:"mac_auth_only,omitempty"`
	// Only if `mode`!=`dynamic` + `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x.
	MacAuthPreferred *bool `json:"mac_auth_preferred,omitempty"`
	// Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`
	MacAuthProtocol *SwitchPortUsageMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
	// Only if `mode`!=`dynamic`, max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)
	MacLimit *SwitchPortUsageMacLimit `json:"mac_limit,omitempty"`
	// `mode`==`dynamic` must only be used if the port usage name is `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk`
	Mode *SwitchPortUsageModeEnum `json:"mode,omitempty"`
	// Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.
	Mtu *SwitchPortUsageMtu `json:"mtu,omitempty"`
	// Only if `mode`==`trunk`, the list of network/vlans
	Networks []string `json:"networks,omitempty"`
	// Only if `mode`==`access` and `port_auth`!=`dot1x`. Whether the port should retain dynamically learned MAC addresses
	PersistMac *bool `json:"persist_mac,omitempty"`
	// Only if `mode`!=`dynamic`. Whether PoE capabilities are disabled for a port
	PoeDisabled *bool `json:"poe_disabled,omitempty"`
	// PoE priority. enum: `low`, `high`
	PoePriority *PoePriorityEnum `json:"poe_priority,omitempty"`
	// Only if `mode`!=`dynamic`. If dot1x is desired, set to dot1x. enum: `dot1x`
	PortAuth Optional[SwitchPortUsageDot1xEnum] `json:"port_auth"`
	// Only if `mode`!=`dynamic`. Native network/vlan for untagged traffic
	PortNetwork *string `json:"port_network,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600)
	ReauthInterval *SwitchPortUsageReauthInterval `json:"reauth_interval,omitempty"`
	// Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)
	ResetDefaultWhen *SwitchPortUsageDynamicResetDefaultWhenEnum `json:"reset_default_when,omitempty"`
	// Only if `mode`==`dynamic`
	Rules []SwitchPortUsageDynamicRule `json:"rules,omitempty"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. Sets server fail fallback vlan
	ServerFailNetwork Optional[string] `json:"server_fail_network"`
	// Only if `mode`!=`dynamic` and `port_auth`==`dot1x`. When radius server reject / fails
	ServerRejectNetwork Optional[string] `json:"server_reject_network"`
	// Only if `mode`!=`dynamic`, Port speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
	Speed *SwitchPortUsageSpeedEnum `json:"speed,omitempty"`
	// Switch storm control. Only if `mode`!=`dynamic`
	StormControl *SwitchPortUsageStormControl `json:"storm_control,omitempty"`
	// Only if `mode`!=`dynamic` and `stp_required`==`false`. Drop bridge protocol data units (BPDUs ) that enter any interface or a specified interface
	StpDisable *bool `json:"stp_disable,omitempty"`
	// Only if `mode`!=`dynamic`. When enabled, the port is not expected to receive BPDU frames
	StpEdge *bool `json:"stp_edge,omitempty"`
	// Only if `mode`!=`dynamic`
	StpNoRootPort *bool `json:"stp_no_root_port,omitempty"`
	// Only if `mode`!=`dynamic`
	StpP2p *bool `json:"stp_p2p,omitempty"`
	// Only if `mode`!=`dynamic`. Whether to remain in block state if no BPDU is received
	StpRequired *bool `json:"stp_required,omitempty"`
	// Optional for Campus Fabric Core-Distribution ESI-LAG profile. Helper used by the UI to select this port profile as the ESI-Lag between Distribution and Access switches
	UiEvpntopoId *uuid.UUID `json:"ui_evpntopo_id,omitempty"`
	// If this is connected to a vstp network
	UseVstp *bool `json:"use_vstp,omitempty"`
	// Only if `mode`!=`dynamic`. Network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth
	VoipNetwork          Optional[string]       `json:"voip_network"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchPortUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchPortUsage) String() string {
	return fmt.Sprintf(
		"SwitchPortUsage[AllNetworks=%v, AllowDhcpd=%v, AllowMultipleSupplicants=%v, BypassAuthWhenServerDown=%v, BypassAuthWhenServerDownForUnknownClient=%v, BypassAuthWhenServerDownForVoip=%v, CommunityVlanId=%v, Description=%v, DisableAutoneg=%v, Disabled=%v, Duplex=%v, DynamicVlanNetworks=%v, EnableMacAuth=%v, EnableQos=%v, GuestNetwork=%v, InterIsolationNetworkLink=%v, InterSwitchLink=%v, MacAuthOnly=%v, MacAuthPreferred=%v, MacAuthProtocol=%v, MacLimit=%v, Mode=%v, Mtu=%v, Networks=%v, PersistMac=%v, PoeDisabled=%v, PoePriority=%v, PortAuth=%v, PortNetwork=%v, ReauthInterval=%v, ResetDefaultWhen=%v, Rules=%v, ServerFailNetwork=%v, ServerRejectNetwork=%v, Speed=%v, StormControl=%v, StpDisable=%v, StpEdge=%v, StpNoRootPort=%v, StpP2p=%v, StpRequired=%v, UiEvpntopoId=%v, UseVstp=%v, VoipNetwork=%v, AdditionalProperties=%v]",
		s.AllNetworks, s.AllowDhcpd, s.AllowMultipleSupplicants, s.BypassAuthWhenServerDown, s.BypassAuthWhenServerDownForUnknownClient, s.BypassAuthWhenServerDownForVoip, s.CommunityVlanId, s.Description, s.DisableAutoneg, s.Disabled, s.Duplex, s.DynamicVlanNetworks, s.EnableMacAuth, s.EnableQos, s.GuestNetwork, s.InterIsolationNetworkLink, s.InterSwitchLink, s.MacAuthOnly, s.MacAuthPreferred, s.MacAuthProtocol, s.MacLimit, s.Mode, s.Mtu, s.Networks, s.PersistMac, s.PoeDisabled, s.PoePriority, s.PortAuth, s.PortNetwork, s.ReauthInterval, s.ResetDefaultWhen, s.Rules, s.ServerFailNetwork, s.ServerRejectNetwork, s.Speed, s.StormControl, s.StpDisable, s.StpEdge, s.StpNoRootPort, s.StpP2p, s.StpRequired, s.UiEvpntopoId, s.UseVstp, s.VoipNetwork, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsage.
// It customizes the JSON marshaling process for SwitchPortUsage objects.
func (s SwitchPortUsage) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"all_networks", "allow_dhcpd", "allow_multiple_supplicants", "bypass_auth_when_server_down", "bypass_auth_when_server_down_for_unknown_client", "bypass_auth_when_server_down_for_voip", "community_vlan_id", "description", "disable_autoneg", "disabled", "duplex", "dynamic_vlan_networks", "enable_mac_auth", "enable_qos", "guest_network", "inter_isolation_network_link", "inter_switch_link", "mac_auth_only", "mac_auth_preferred", "mac_auth_protocol", "mac_limit", "mode", "mtu", "networks", "persist_mac", "poe_disabled", "poe_priority", "port_auth", "port_network", "reauth_interval", "reset_default_when", "rules", "server_fail_network", "server_reject_network", "speed", "storm_control", "stp_disable", "stp_edge", "stp_no_root_port", "stp_p2p", "stp_required", "ui_evpntopo_id", "use_vstp", "voip_network"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsage object to a map representation for JSON marshaling.
func (s SwitchPortUsage) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AllNetworks != nil {
		structMap["all_networks"] = s.AllNetworks
	}
	if s.AllowDhcpd != nil {
		structMap["allow_dhcpd"] = s.AllowDhcpd
	}
	if s.AllowMultipleSupplicants != nil {
		structMap["allow_multiple_supplicants"] = s.AllowMultipleSupplicants
	}
	if s.BypassAuthWhenServerDown != nil {
		structMap["bypass_auth_when_server_down"] = s.BypassAuthWhenServerDown
	}
	if s.BypassAuthWhenServerDownForUnknownClient != nil {
		structMap["bypass_auth_when_server_down_for_unknown_client"] = s.BypassAuthWhenServerDownForUnknownClient
	}
	if s.BypassAuthWhenServerDownForVoip != nil {
		structMap["bypass_auth_when_server_down_for_voip"] = s.BypassAuthWhenServerDownForVoip
	}
	if s.CommunityVlanId != nil {
		structMap["community_vlan_id"] = s.CommunityVlanId
	}
	if s.Description != nil {
		structMap["description"] = s.Description
	}
	if s.DisableAutoneg != nil {
		structMap["disable_autoneg"] = s.DisableAutoneg
	}
	if s.Disabled != nil {
		structMap["disabled"] = s.Disabled
	}
	if s.Duplex != nil {
		structMap["duplex"] = s.Duplex
	}
	if s.DynamicVlanNetworks != nil {
		structMap["dynamic_vlan_networks"] = s.DynamicVlanNetworks
	}
	if s.EnableMacAuth != nil {
		structMap["enable_mac_auth"] = s.EnableMacAuth
	}
	if s.EnableQos != nil {
		structMap["enable_qos"] = s.EnableQos
	}
	if s.GuestNetwork.IsValueSet() {
		if s.GuestNetwork.Value() != nil {
			structMap["guest_network"] = s.GuestNetwork.Value()
		} else {
			structMap["guest_network"] = nil
		}
	}
	if s.InterIsolationNetworkLink != nil {
		structMap["inter_isolation_network_link"] = s.InterIsolationNetworkLink
	}
	if s.InterSwitchLink != nil {
		structMap["inter_switch_link"] = s.InterSwitchLink
	}
	if s.MacAuthOnly != nil {
		structMap["mac_auth_only"] = s.MacAuthOnly
	}
	if s.MacAuthPreferred != nil {
		structMap["mac_auth_preferred"] = s.MacAuthPreferred
	}
	if s.MacAuthProtocol != nil {
		structMap["mac_auth_protocol"] = s.MacAuthProtocol
	}
	if s.MacLimit != nil {
		structMap["mac_limit"] = s.MacLimit.toMap()
	}
	if s.Mode != nil {
		structMap["mode"] = s.Mode
	}
	if s.Mtu != nil {
		structMap["mtu"] = s.Mtu.toMap()
	}
	if s.Networks != nil {
		structMap["networks"] = s.Networks
	}
	if s.PersistMac != nil {
		structMap["persist_mac"] = s.PersistMac
	}
	if s.PoeDisabled != nil {
		structMap["poe_disabled"] = s.PoeDisabled
	}
	if s.PoePriority != nil {
		structMap["poe_priority"] = s.PoePriority
	}
	if s.PortAuth.IsValueSet() {
		if s.PortAuth.Value() != nil {
			structMap["port_auth"] = s.PortAuth.Value()
		} else {
			structMap["port_auth"] = nil
		}
	}
	if s.PortNetwork != nil {
		structMap["port_network"] = s.PortNetwork
	}
	if s.ReauthInterval != nil {
		structMap["reauth_interval"] = s.ReauthInterval.toMap()
	}
	if s.ResetDefaultWhen != nil {
		structMap["reset_default_when"] = s.ResetDefaultWhen
	}
	if s.Rules != nil {
		structMap["rules"] = s.Rules
	}
	if s.ServerFailNetwork.IsValueSet() {
		if s.ServerFailNetwork.Value() != nil {
			structMap["server_fail_network"] = s.ServerFailNetwork.Value()
		} else {
			structMap["server_fail_network"] = nil
		}
	}
	if s.ServerRejectNetwork.IsValueSet() {
		if s.ServerRejectNetwork.Value() != nil {
			structMap["server_reject_network"] = s.ServerRejectNetwork.Value()
		} else {
			structMap["server_reject_network"] = nil
		}
	}
	if s.Speed != nil {
		structMap["speed"] = s.Speed
	}
	if s.StormControl != nil {
		structMap["storm_control"] = s.StormControl.toMap()
	}
	if s.StpDisable != nil {
		structMap["stp_disable"] = s.StpDisable
	}
	if s.StpEdge != nil {
		structMap["stp_edge"] = s.StpEdge
	}
	if s.StpNoRootPort != nil {
		structMap["stp_no_root_port"] = s.StpNoRootPort
	}
	if s.StpP2p != nil {
		structMap["stp_p2p"] = s.StpP2p
	}
	if s.StpRequired != nil {
		structMap["stp_required"] = s.StpRequired
	}
	if s.UiEvpntopoId != nil {
		structMap["ui_evpntopo_id"] = s.UiEvpntopoId
	}
	if s.UseVstp != nil {
		structMap["use_vstp"] = s.UseVstp
	}
	if s.VoipNetwork.IsValueSet() {
		if s.VoipNetwork.Value() != nil {
			structMap["voip_network"] = s.VoipNetwork.Value()
		} else {
			structMap["voip_network"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsage.
// It customizes the JSON unmarshaling process for SwitchPortUsage objects.
func (s *SwitchPortUsage) UnmarshalJSON(input []byte) error {
	var temp tempSwitchPortUsage
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "all_networks", "allow_dhcpd", "allow_multiple_supplicants", "bypass_auth_when_server_down", "bypass_auth_when_server_down_for_unknown_client", "bypass_auth_when_server_down_for_voip", "community_vlan_id", "description", "disable_autoneg", "disabled", "duplex", "dynamic_vlan_networks", "enable_mac_auth", "enable_qos", "guest_network", "inter_isolation_network_link", "inter_switch_link", "mac_auth_only", "mac_auth_preferred", "mac_auth_protocol", "mac_limit", "mode", "mtu", "networks", "persist_mac", "poe_disabled", "poe_priority", "port_auth", "port_network", "reauth_interval", "reset_default_when", "rules", "server_fail_network", "server_reject_network", "speed", "storm_control", "stp_disable", "stp_edge", "stp_no_root_port", "stp_p2p", "stp_required", "ui_evpntopo_id", "use_vstp", "voip_network")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AllNetworks = temp.AllNetworks
	s.AllowDhcpd = temp.AllowDhcpd
	s.AllowMultipleSupplicants = temp.AllowMultipleSupplicants
	s.BypassAuthWhenServerDown = temp.BypassAuthWhenServerDown
	s.BypassAuthWhenServerDownForUnknownClient = temp.BypassAuthWhenServerDownForUnknownClient
	s.BypassAuthWhenServerDownForVoip = temp.BypassAuthWhenServerDownForVoip
	s.CommunityVlanId = temp.CommunityVlanId
	s.Description = temp.Description
	s.DisableAutoneg = temp.DisableAutoneg
	s.Disabled = temp.Disabled
	s.Duplex = temp.Duplex
	s.DynamicVlanNetworks = temp.DynamicVlanNetworks
	s.EnableMacAuth = temp.EnableMacAuth
	s.EnableQos = temp.EnableQos
	s.GuestNetwork = temp.GuestNetwork
	s.InterIsolationNetworkLink = temp.InterIsolationNetworkLink
	s.InterSwitchLink = temp.InterSwitchLink
	s.MacAuthOnly = temp.MacAuthOnly
	s.MacAuthPreferred = temp.MacAuthPreferred
	s.MacAuthProtocol = temp.MacAuthProtocol
	s.MacLimit = temp.MacLimit
	s.Mode = temp.Mode
	s.Mtu = temp.Mtu
	s.Networks = temp.Networks
	s.PersistMac = temp.PersistMac
	s.PoeDisabled = temp.PoeDisabled
	s.PoePriority = temp.PoePriority
	s.PortAuth = temp.PortAuth
	s.PortNetwork = temp.PortNetwork
	s.ReauthInterval = temp.ReauthInterval
	s.ResetDefaultWhen = temp.ResetDefaultWhen
	s.Rules = temp.Rules
	s.ServerFailNetwork = temp.ServerFailNetwork
	s.ServerRejectNetwork = temp.ServerRejectNetwork
	s.Speed = temp.Speed
	s.StormControl = temp.StormControl
	s.StpDisable = temp.StpDisable
	s.StpEdge = temp.StpEdge
	s.StpNoRootPort = temp.StpNoRootPort
	s.StpP2p = temp.StpP2p
	s.StpRequired = temp.StpRequired
	s.UiEvpntopoId = temp.UiEvpntopoId
	s.UseVstp = temp.UseVstp
	s.VoipNetwork = temp.VoipNetwork
	return nil
}

// tempSwitchPortUsage is a temporary struct used for validating the fields of SwitchPortUsage.
type tempSwitchPortUsage struct {
	AllNetworks                              *bool                                       `json:"all_networks,omitempty"`
	AllowDhcpd                               *bool                                       `json:"allow_dhcpd,omitempty"`
	AllowMultipleSupplicants                 *bool                                       `json:"allow_multiple_supplicants,omitempty"`
	BypassAuthWhenServerDown                 *bool                                       `json:"bypass_auth_when_server_down,omitempty"`
	BypassAuthWhenServerDownForUnknownClient *bool                                       `json:"bypass_auth_when_server_down_for_unknown_client,omitempty"`
	BypassAuthWhenServerDownForVoip          *bool                                       `json:"bypass_auth_when_server_down_for_voip,omitempty"`
	CommunityVlanId                          *int                                        `json:"community_vlan_id,omitempty"`
	Description                              *string                                     `json:"description,omitempty"`
	DisableAutoneg                           *bool                                       `json:"disable_autoneg,omitempty"`
	Disabled                                 *bool                                       `json:"disabled,omitempty"`
	Duplex                                   *SwitchPortUsageDuplexEnum                  `json:"duplex,omitempty"`
	DynamicVlanNetworks                      []string                                    `json:"dynamic_vlan_networks,omitempty"`
	EnableMacAuth                            *bool                                       `json:"enable_mac_auth,omitempty"`
	EnableQos                                *bool                                       `json:"enable_qos,omitempty"`
	GuestNetwork                             Optional[string]                            `json:"guest_network"`
	InterIsolationNetworkLink                *bool                                       `json:"inter_isolation_network_link,omitempty"`
	InterSwitchLink                          *bool                                       `json:"inter_switch_link,omitempty"`
	MacAuthOnly                              *bool                                       `json:"mac_auth_only,omitempty"`
	MacAuthPreferred                         *bool                                       `json:"mac_auth_preferred,omitempty"`
	MacAuthProtocol                          *SwitchPortUsageMacAuthProtocolEnum         `json:"mac_auth_protocol,omitempty"`
	MacLimit                                 *SwitchPortUsageMacLimit                    `json:"mac_limit,omitempty"`
	Mode                                     *SwitchPortUsageModeEnum                    `json:"mode,omitempty"`
	Mtu                                      *SwitchPortUsageMtu                         `json:"mtu,omitempty"`
	Networks                                 []string                                    `json:"networks,omitempty"`
	PersistMac                               *bool                                       `json:"persist_mac,omitempty"`
	PoeDisabled                              *bool                                       `json:"poe_disabled,omitempty"`
	PoePriority                              *PoePriorityEnum                            `json:"poe_priority,omitempty"`
	PortAuth                                 Optional[SwitchPortUsageDot1xEnum]          `json:"port_auth"`
	PortNetwork                              *string                                     `json:"port_network,omitempty"`
	ReauthInterval                           *SwitchPortUsageReauthInterval              `json:"reauth_interval,omitempty"`
	ResetDefaultWhen                         *SwitchPortUsageDynamicResetDefaultWhenEnum `json:"reset_default_when,omitempty"`
	Rules                                    []SwitchPortUsageDynamicRule                `json:"rules,omitempty"`
	ServerFailNetwork                        Optional[string]                            `json:"server_fail_network"`
	ServerRejectNetwork                      Optional[string]                            `json:"server_reject_network"`
	Speed                                    *SwitchPortUsageSpeedEnum                   `json:"speed,omitempty"`
	StormControl                             *SwitchPortUsageStormControl                `json:"storm_control,omitempty"`
	StpDisable                               *bool                                       `json:"stp_disable,omitempty"`
	StpEdge                                  *bool                                       `json:"stp_edge,omitempty"`
	StpNoRootPort                            *bool                                       `json:"stp_no_root_port,omitempty"`
	StpP2p                                   *bool                                       `json:"stp_p2p,omitempty"`
	StpRequired                              *bool                                       `json:"stp_required,omitempty"`
	UiEvpntopoId                             *uuid.UUID                                  `json:"ui_evpntopo_id,omitempty"`
	UseVstp                                  *bool                                       `json:"use_vstp,omitempty"`
	VoipNetwork                              Optional[string]                            `json:"voip_network"`
}
