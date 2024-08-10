package models

import (
    "encoding/json"
)

// SwitchPortUsage represents a SwitchPortUsage struct.
// Junos port usages
type SwitchPortUsage struct {
    // Only if `mode`==`trunk` whether to trunk all network/vlans
    AllNetworks                              *bool                                       `json:"all_networks,omitempty"`
    // Only if `mode`!=`dynamic` if DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state.
    // When it is not defined, it means using the system’s default setting which depends on whether the port is a access or trunk port.
    AllowDhcpd                               *bool                                       `json:"allow_dhcpd,omitempty"`
    // Only if `mode`!=`dynamic`
    AllowMultipleSupplicants                 *bool                                       `json:"allow_multiple_supplicants,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`==`dot1x` bypass auth for known clients if set to true when RADIUS server is down
    BypassAuthWhenServerDown                 *bool                                       `json:"bypass_auth_when_server_down,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`=`dot1x` bypass auth for all (including unknown clients) if set to true when RADIUS server is down
    BypassAuthWhenServerDownForUnkonwnClient *bool                                       `json:"bypass_auth_when_server_down_for_unkonwn_client,omitempty"`
    // Only if `mode`!=`dynamic`
    Description                              *string                                     `json:"description,omitempty"`
    // Only if `mode`!=`dynamic` if speed and duplex are specified, whether to disable autonegotiation
    DisableAutoneg                           *bool                                       `json:"disable_autoneg,omitempty"`
    // Only if `mode`!=`dynamic` whether the port is disabled
    Disabled                                 *bool                                       `json:"disabled,omitempty"`
    // Only if `mode`!=`dynamic` link connection mode. enum: `auto`, `full`, `half`
    Duplex                                   *SwitchPortUsageDuplexEnum                  `json:"duplex,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return
    DynamicVlanNetworks                      []string                                    `json:"dynamic_vlan_networks,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`==`dot1x` whether to enable MAC Auth
    EnableMacAuth                            *bool                                       `json:"enable_mac_auth,omitempty"`
    // Only if `mode`!=`dynamic`
    EnableQos                                *bool                                       `json:"enable_qos,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`==`dot1x` which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed)
    GuestNetwork                             Optional[string]                            `json:"guest_network"`
    // Only if `mode`!=`dynamic` inter_switch_link is used together with "isolation" under networks
    // NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together
    InterSwitchLink                          *bool                                       `json:"inter_switch_link,omitempty"`
    // Only if `mode`!=`dynamic` and `enable_mac_auth`==`true`
    MacAuthOnly                              *bool                                       `json:"mac_auth_only,omitempty"`
    // Only if `mode`!=`dynamic` and `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`
    MacAuthProtocol                          *SwitchPortUsageMacAuthProtocolEnum         `json:"mac_auth_protocol,omitempty"`
    // Only if `mode`!=`dynamic` max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform
    MacLimit                                 *int                                        `json:"mac_limit,omitempty"`
    // `mode`==`dynamic` must only be used with the port usage with the name `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk`
    Mode                                     *SwitchPortUsageModeEnum                    `json:"mode,omitempty"`
    // Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.
    Mtu                                      *int                                        `json:"mtu,omitempty"`
    // Only if `mode`==`trunk`, the list of network/vlans
    Networks                                 []string                                    `json:"networks,omitempty"`
    // Only if `mode`==`access` and `port_auth`!=`dot1x` whether the port should retain dynamically learned MAC addresses
    PersistMac                               *bool                                       `json:"persist_mac,omitempty"`
    // Only if `mode`!=`dynamic` whether PoE capabilities are disabled for a port
    PoeDisabled                              *bool                                       `json:"poe_disabled,omitempty"`
    // Only if `mode`!=`dynamic` if dot1x is desired, set to dot1x. enum: `dot1x`
    PortAuth                                 Optional[SwitchPortUsageDot1xEnum]          `json:"port_auth"`
    // Only if `mode`!=`dynamic` native network/vlan for untagged traffic
    PortNetwork                              *string                                     `json:"port_network,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range
    ReauthInterval                           *int                                        `json:"reauth_interval,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`==`dot1x` when radius server reject / fails
    RejectedNetwork                          Optional[string]                            `json:"rejected_network"`
    // Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)
    ResetDefaultWhen                         *SwitchPortUsageDynamicResetDefaultWhenEnum `json:"reset_default_when,omitempty"`
    // Only if `mode`==`dynamic`
    Rules                                    []SwitchPortUsageDynamicRule                `json:"rules,omitempty"`
    // Only if `mode`!=`dynamic` speed, default is auto to automatically negotiate speed
    Speed                                    *string                                     `json:"speed,omitempty"`
    // Switch storm control
    // Only if `mode`!=`dynamic`
    StormControl                             *SwitchPortUsageStormControl                `json:"storm_control,omitempty"`
    // Only if `mode`!=`dynamic` when enabled, the port is not expected to receive BPDU frames
    StpEdge                                  *bool                                       `json:"stp_edge,omitempty"`
    StpNoRootPort                            *bool                                       `json:"stp_no_root_port,omitempty"`
    StpP2p                                   *bool                                       `json:"stp_p2p,omitempty"`
    // Only if `mode`!=`dynamic` network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth
    VoipNetwork                              *string                                     `json:"voip_network,omitempty"`
    AdditionalProperties                     map[string]any                              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsage.
// It customizes the JSON marshaling process for SwitchPortUsage objects.
func (s SwitchPortUsage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsage object to a map representation for JSON marshaling.
func (s SwitchPortUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.BypassAuthWhenServerDownForUnkonwnClient != nil {
        structMap["bypass_auth_when_server_down_for_unkonwn_client"] = s.BypassAuthWhenServerDownForUnkonwnClient
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
    if s.InterSwitchLink != nil {
        structMap["inter_switch_link"] = s.InterSwitchLink
    }
    if s.MacAuthOnly != nil {
        structMap["mac_auth_only"] = s.MacAuthOnly
    }
    if s.MacAuthProtocol != nil {
        structMap["mac_auth_protocol"] = s.MacAuthProtocol
    }
    if s.MacLimit != nil {
        structMap["mac_limit"] = s.MacLimit
    }
    if s.Mode != nil {
        structMap["mode"] = s.Mode
    }
    if s.Mtu != nil {
        structMap["mtu"] = s.Mtu
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
        structMap["reauth_interval"] = s.ReauthInterval
    }
    if s.RejectedNetwork.IsValueSet() {
        if s.RejectedNetwork.Value() != nil {
            structMap["rejected_network"] = s.RejectedNetwork.Value()
        } else {
            structMap["rejected_network"] = nil
        }
    }
    if s.ResetDefaultWhen != nil {
        structMap["reset_default_when"] = s.ResetDefaultWhen
    }
    if s.Rules != nil {
        structMap["rules"] = s.Rules
    }
    if s.Speed != nil {
        structMap["speed"] = s.Speed
    }
    if s.StormControl != nil {
        structMap["storm_control"] = s.StormControl.toMap()
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
    if s.VoipNetwork != nil {
        structMap["voip_network"] = s.VoipNetwork
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "all_networks", "allow_dhcpd", "allow_multiple_supplicants", "bypass_auth_when_server_down", "bypass_auth_when_server_down_for_unkonwn_client", "description", "disable_autoneg", "disabled", "duplex", "dynamic_vlan_networks", "enable_mac_auth", "enable_qos", "guest_network", "inter_switch_link", "mac_auth_only", "mac_auth_protocol", "mac_limit", "mode", "mtu", "networks", "persist_mac", "poe_disabled", "port_auth", "port_network", "reauth_interval", "rejected_network", "reset_default_when", "rules", "speed", "storm_control", "stp_edge", "stp_no_root_port", "stp_p2p", "voip_network")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AllNetworks = temp.AllNetworks
    s.AllowDhcpd = temp.AllowDhcpd
    s.AllowMultipleSupplicants = temp.AllowMultipleSupplicants
    s.BypassAuthWhenServerDown = temp.BypassAuthWhenServerDown
    s.BypassAuthWhenServerDownForUnkonwnClient = temp.BypassAuthWhenServerDownForUnkonwnClient
    s.Description = temp.Description
    s.DisableAutoneg = temp.DisableAutoneg
    s.Disabled = temp.Disabled
    s.Duplex = temp.Duplex
    s.DynamicVlanNetworks = temp.DynamicVlanNetworks
    s.EnableMacAuth = temp.EnableMacAuth
    s.EnableQos = temp.EnableQos
    s.GuestNetwork = temp.GuestNetwork
    s.InterSwitchLink = temp.InterSwitchLink
    s.MacAuthOnly = temp.MacAuthOnly
    s.MacAuthProtocol = temp.MacAuthProtocol
    s.MacLimit = temp.MacLimit
    s.Mode = temp.Mode
    s.Mtu = temp.Mtu
    s.Networks = temp.Networks
    s.PersistMac = temp.PersistMac
    s.PoeDisabled = temp.PoeDisabled
    s.PortAuth = temp.PortAuth
    s.PortNetwork = temp.PortNetwork
    s.ReauthInterval = temp.ReauthInterval
    s.RejectedNetwork = temp.RejectedNetwork
    s.ResetDefaultWhen = temp.ResetDefaultWhen
    s.Rules = temp.Rules
    s.Speed = temp.Speed
    s.StormControl = temp.StormControl
    s.StpEdge = temp.StpEdge
    s.StpNoRootPort = temp.StpNoRootPort
    s.StpP2p = temp.StpP2p
    s.VoipNetwork = temp.VoipNetwork
    return nil
}

// tempSwitchPortUsage is a temporary struct used for validating the fields of SwitchPortUsage.
type tempSwitchPortUsage  struct {
    AllNetworks                              *bool                                       `json:"all_networks,omitempty"`
    AllowDhcpd                               *bool                                       `json:"allow_dhcpd,omitempty"`
    AllowMultipleSupplicants                 *bool                                       `json:"allow_multiple_supplicants,omitempty"`
    BypassAuthWhenServerDown                 *bool                                       `json:"bypass_auth_when_server_down,omitempty"`
    BypassAuthWhenServerDownForUnkonwnClient *bool                                       `json:"bypass_auth_when_server_down_for_unkonwn_client,omitempty"`
    Description                              *string                                     `json:"description,omitempty"`
    DisableAutoneg                           *bool                                       `json:"disable_autoneg,omitempty"`
    Disabled                                 *bool                                       `json:"disabled,omitempty"`
    Duplex                                   *SwitchPortUsageDuplexEnum                  `json:"duplex,omitempty"`
    DynamicVlanNetworks                      []string                                    `json:"dynamic_vlan_networks,omitempty"`
    EnableMacAuth                            *bool                                       `json:"enable_mac_auth,omitempty"`
    EnableQos                                *bool                                       `json:"enable_qos,omitempty"`
    GuestNetwork                             Optional[string]                            `json:"guest_network"`
    InterSwitchLink                          *bool                                       `json:"inter_switch_link,omitempty"`
    MacAuthOnly                              *bool                                       `json:"mac_auth_only,omitempty"`
    MacAuthProtocol                          *SwitchPortUsageMacAuthProtocolEnum         `json:"mac_auth_protocol,omitempty"`
    MacLimit                                 *int                                        `json:"mac_limit,omitempty"`
    Mode                                     *SwitchPortUsageModeEnum                    `json:"mode,omitempty"`
    Mtu                                      *int                                        `json:"mtu,omitempty"`
    Networks                                 []string                                    `json:"networks,omitempty"`
    PersistMac                               *bool                                       `json:"persist_mac,omitempty"`
    PoeDisabled                              *bool                                       `json:"poe_disabled,omitempty"`
    PortAuth                                 Optional[SwitchPortUsageDot1xEnum]          `json:"port_auth"`
    PortNetwork                              *string                                     `json:"port_network,omitempty"`
    ReauthInterval                           *int                                        `json:"reauth_interval,omitempty"`
    RejectedNetwork                          Optional[string]                            `json:"rejected_network"`
    ResetDefaultWhen                         *SwitchPortUsageDynamicResetDefaultWhenEnum `json:"reset_default_when,omitempty"`
    Rules                                    []SwitchPortUsageDynamicRule                `json:"rules,omitempty"`
    Speed                                    *string                                     `json:"speed,omitempty"`
    StormControl                             *SwitchPortUsageStormControl                `json:"storm_control,omitempty"`
    StpEdge                                  *bool                                       `json:"stp_edge,omitempty"`
    StpNoRootPort                            *bool                                       `json:"stp_no_root_port,omitempty"`
    StpP2p                                   *bool                                       `json:"stp_p2p,omitempty"`
    VoipNetwork                              *string                                     `json:"voip_network,omitempty"`
}
