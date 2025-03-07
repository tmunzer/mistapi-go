package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// JunosLocalPortConfig represents a JunosLocalPortConfig struct.
// Switch port config
type JunosLocalPortConfig struct {
    // Only if `mode`==`trunk` whether to trunk all network/vlans
    AllNetworks                              *bool                                    `json:"all_networks,omitempty"`
    // If DHCP snooping is enabled, whether DHCP server is allowed on the interfaces with. All the interfaces from port configs using this port usage are effected. Please notice that allow_dhcpd is a tri_state. When it is not defined, it means using the system's default setting which depends on whether the port is an access or trunk port.
    AllowDhcpd                               *bool                                    `json:"allow_dhcpd,omitempty"`
    AllowMultipleSupplicants                 *bool                                    `json:"allow_multiple_supplicants,omitempty"`
    // Only if `port_auth`==`dot1x` bypass auth for known clients if set to true when RADIUS server is down
    BypassAuthWhenServerDown                 *bool                                    `json:"bypass_auth_when_server_down,omitempty"`
    // Only if `port_auth`=`dot1x` bypass auth for all (including unknown clients) if set to true when RADIUS server is down
    BypassAuthWhenServerDownForUnknownClient *bool                                    `json:"bypass_auth_when_server_down_for_unknown_client,omitempty"`
    Description                              *string                                  `json:"description,omitempty"`
    // Only if `mode`!=`dynamic` if speed and duplex are specified, whether to disable autonegotiation
    DisableAutoneg                           *bool                                    `json:"disable_autoneg,omitempty"`
    // Whether the port is disabled
    Disabled                                 *bool                                    `json:"disabled,omitempty"`
    // link connection mode. enum: `auto`, `full`, `half`
    Duplex                                   *SwitchPortLocalUsageDuplexEnum          `json:"duplex,omitempty"`
    // Only if `port_auth`==`dot1x`, if dynamic vlan is used, specify the possible networks/vlans RADIUS can return
    DynamicVlanNetworks                      []string                                 `json:"dynamic_vlan_networks,omitempty"`
    // Only if `port_auth`==`dot1x` whether to enable MAC Auth
    EnableMacAuth                            *bool                                    `json:"enable_mac_auth,omitempty"`
    EnableQos                                *bool                                    `json:"enable_qos,omitempty"`
    // Only if `port_auth`==`dot1x` which network to put the device into if the device cannot do dot1x. default is null (i.e. not allowed)
    GuestNetwork                             Optional[string]                         `json:"guest_network"`
    // inter_switch_link is used together with "isolation" under networks. NOTE: inter_switch_link works only between Juniper device. This has to be applied to both ports connected together
    InterSwitchLink                          *bool                                    `json:"inter_switch_link,omitempty"`
    // Only if `enable_mac_auth`==`true`
    MacAuthOnly                              *bool                                    `json:"mac_auth_only,omitempty"`
    // Only if `enable_mac_auth`==`true` + `mac_auth_only`==`false`, dot1x will be given priority then mac_auth. Enable this to prefer mac_auth over dot1x.
    MacAuthPreferred                         *bool                                    `json:"mac_auth_preferred,omitempty"`
    // Only if `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`
    MacAuthProtocol                          *SwitchPortLocalUsageMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    // Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 or higher, with upper bound constrained by platform
    MacLimit                                 *int                                     `json:"mac_limit,omitempty"`
    // enum: `access`, `inet`, `trunk`
    Mode                                     *SwitchPortLocalUsageModeEnum            `json:"mode,omitempty"`
    // Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.
    Mtu                                      *int                                     `json:"mtu,omitempty"`
    // Only if `mode`==`trunk`, the list of network/vlans
    Networks                                 []string                                 `json:"networks,omitempty"`
    // Additional note for the port config override
    Note                                     *string                                  `json:"note,omitempty"`
    // Only if `mode`==`access` and `port_auth`!=`dot1x` whether the port should retain dynamically learned MAC addresses
    PersistMac                               *bool                                    `json:"persist_mac,omitempty"`
    // Whether PoE capabilities are disabled for a port
    PoeDisabled                              *bool                                    `json:"poe_disabled,omitempty"`
    // if dot1x is desired, set to dot1x. enum: `dot1x`
    PortAuth                                 Optional[SwitchPortLocalUsageDot1xEnum]  `json:"port_auth"`
    // Native network/vlan for untagged traffic
    PortNetwork                              *string                                  `json:"port_network,omitempty"`
    // Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600)
    ReauthInterval                           *SwitchPortUsageReauthInterval           `json:"reauth_interval,omitempty"`
    // Only if `port_auth`==`dot1x` sets server fail fallback vlan
    ServerFailNetwork                        Optional[string]                         `json:"server_fail_network"`
    // Only if `port_auth`==`dot1x` when radius server reject / fails
    ServerRejectNetwork                      Optional[string]                         `json:"server_reject_network"`
    // enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
    Speed                                    *JunosPortConfigSpeedEnum                `json:"speed,omitempty"`
    // Switch storm control
    StormControl                             *SwitchPortLocalUsageStormControl        `json:"storm_control,omitempty"`
    // When enabled, the port is not expected to receive BPDU frames
    StpEdge                                  *bool                                    `json:"stp_edge,omitempty"`
    StpNoRootPort                            *bool                                    `json:"stp_no_root_port,omitempty"`
    StpP2p                                   *bool                                    `json:"stp_p2p,omitempty"`
    // Port usage name.
    Usage                                    string                                   `json:"usage"`
    // If this is connected to a vstp network
    UseVstp                                  *bool                                    `json:"use_vstp,omitempty"`
    // Network/vlan for voip traffic, must also set port_network. to authenticate device, set port_auth
    VoipNetwork                              *string                                  `json:"voip_network,omitempty"`
    AdditionalProperties                     map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for JunosLocalPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JunosLocalPortConfig) String() string {
    return fmt.Sprintf(
    	"JunosLocalPortConfig[AllNetworks=%v, AllowDhcpd=%v, AllowMultipleSupplicants=%v, BypassAuthWhenServerDown=%v, BypassAuthWhenServerDownForUnknownClient=%v, Description=%v, DisableAutoneg=%v, Disabled=%v, Duplex=%v, DynamicVlanNetworks=%v, EnableMacAuth=%v, EnableQos=%v, GuestNetwork=%v, InterSwitchLink=%v, MacAuthOnly=%v, MacAuthPreferred=%v, MacAuthProtocol=%v, MacLimit=%v, Mode=%v, Mtu=%v, Networks=%v, Note=%v, PersistMac=%v, PoeDisabled=%v, PortAuth=%v, PortNetwork=%v, ReauthInterval=%v, ServerFailNetwork=%v, ServerRejectNetwork=%v, Speed=%v, StormControl=%v, StpEdge=%v, StpNoRootPort=%v, StpP2p=%v, Usage=%v, UseVstp=%v, VoipNetwork=%v, AdditionalProperties=%v]",
    	j.AllNetworks, j.AllowDhcpd, j.AllowMultipleSupplicants, j.BypassAuthWhenServerDown, j.BypassAuthWhenServerDownForUnknownClient, j.Description, j.DisableAutoneg, j.Disabled, j.Duplex, j.DynamicVlanNetworks, j.EnableMacAuth, j.EnableQos, j.GuestNetwork, j.InterSwitchLink, j.MacAuthOnly, j.MacAuthPreferred, j.MacAuthProtocol, j.MacLimit, j.Mode, j.Mtu, j.Networks, j.Note, j.PersistMac, j.PoeDisabled, j.PortAuth, j.PortNetwork, j.ReauthInterval, j.ServerFailNetwork, j.ServerRejectNetwork, j.Speed, j.StormControl, j.StpEdge, j.StpNoRootPort, j.StpP2p, j.Usage, j.UseVstp, j.VoipNetwork, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JunosLocalPortConfig.
// It customizes the JSON marshaling process for JunosLocalPortConfig objects.
func (j JunosLocalPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "all_networks", "allow_dhcpd", "allow_multiple_supplicants", "bypass_auth_when_server_down", "bypass_auth_when_server_down_for_unknown_client", "description", "disable_autoneg", "disabled", "duplex", "dynamic_vlan_networks", "enable_mac_auth", "enable_qos", "guest_network", "inter_switch_link", "mac_auth_only", "mac_auth_preferred", "mac_auth_protocol", "mac_limit", "mode", "mtu", "networks", "note", "persist_mac", "poe_disabled", "port_auth", "port_network", "reauth_interval", "server_fail_network", "server_reject_network", "speed", "storm_control", "stp_edge", "stp_no_root_port", "stp_p2p", "usage", "use_vstp", "voip_network"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JunosLocalPortConfig object to a map representation for JSON marshaling.
func (j JunosLocalPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.AllNetworks != nil {
        structMap["all_networks"] = j.AllNetworks
    }
    if j.AllowDhcpd != nil {
        structMap["allow_dhcpd"] = j.AllowDhcpd
    }
    if j.AllowMultipleSupplicants != nil {
        structMap["allow_multiple_supplicants"] = j.AllowMultipleSupplicants
    }
    if j.BypassAuthWhenServerDown != nil {
        structMap["bypass_auth_when_server_down"] = j.BypassAuthWhenServerDown
    }
    if j.BypassAuthWhenServerDownForUnknownClient != nil {
        structMap["bypass_auth_when_server_down_for_unknown_client"] = j.BypassAuthWhenServerDownForUnknownClient
    }
    if j.Description != nil {
        structMap["description"] = j.Description
    }
    if j.DisableAutoneg != nil {
        structMap["disable_autoneg"] = j.DisableAutoneg
    }
    if j.Disabled != nil {
        structMap["disabled"] = j.Disabled
    }
    if j.Duplex != nil {
        structMap["duplex"] = j.Duplex
    }
    if j.DynamicVlanNetworks != nil {
        structMap["dynamic_vlan_networks"] = j.DynamicVlanNetworks
    }
    if j.EnableMacAuth != nil {
        structMap["enable_mac_auth"] = j.EnableMacAuth
    }
    if j.EnableQos != nil {
        structMap["enable_qos"] = j.EnableQos
    }
    if j.GuestNetwork.IsValueSet() {
        if j.GuestNetwork.Value() != nil {
            structMap["guest_network"] = j.GuestNetwork.Value()
        } else {
            structMap["guest_network"] = nil
        }
    }
    if j.InterSwitchLink != nil {
        structMap["inter_switch_link"] = j.InterSwitchLink
    }
    if j.MacAuthOnly != nil {
        structMap["mac_auth_only"] = j.MacAuthOnly
    }
    if j.MacAuthPreferred != nil {
        structMap["mac_auth_preferred"] = j.MacAuthPreferred
    }
    if j.MacAuthProtocol != nil {
        structMap["mac_auth_protocol"] = j.MacAuthProtocol
    }
    if j.MacLimit != nil {
        structMap["mac_limit"] = j.MacLimit
    }
    if j.Mode != nil {
        structMap["mode"] = j.Mode
    }
    if j.Mtu != nil {
        structMap["mtu"] = j.Mtu
    }
    if j.Networks != nil {
        structMap["networks"] = j.Networks
    }
    if j.Note != nil {
        structMap["note"] = j.Note
    }
    if j.PersistMac != nil {
        structMap["persist_mac"] = j.PersistMac
    }
    if j.PoeDisabled != nil {
        structMap["poe_disabled"] = j.PoeDisabled
    }
    if j.PortAuth.IsValueSet() {
        if j.PortAuth.Value() != nil {
            structMap["port_auth"] = j.PortAuth.Value()
        } else {
            structMap["port_auth"] = nil
        }
    }
    if j.PortNetwork != nil {
        structMap["port_network"] = j.PortNetwork
    }
    if j.ReauthInterval != nil {
        structMap["reauth_interval"] = j.ReauthInterval.toMap()
    }
    if j.ServerFailNetwork.IsValueSet() {
        if j.ServerFailNetwork.Value() != nil {
            structMap["server_fail_network"] = j.ServerFailNetwork.Value()
        } else {
            structMap["server_fail_network"] = nil
        }
    }
    if j.ServerRejectNetwork.IsValueSet() {
        if j.ServerRejectNetwork.Value() != nil {
            structMap["server_reject_network"] = j.ServerRejectNetwork.Value()
        } else {
            structMap["server_reject_network"] = nil
        }
    }
    if j.Speed != nil {
        structMap["speed"] = j.Speed
    }
    if j.StormControl != nil {
        structMap["storm_control"] = j.StormControl.toMap()
    }
    if j.StpEdge != nil {
        structMap["stp_edge"] = j.StpEdge
    }
    if j.StpNoRootPort != nil {
        structMap["stp_no_root_port"] = j.StpNoRootPort
    }
    if j.StpP2p != nil {
        structMap["stp_p2p"] = j.StpP2p
    }
    structMap["usage"] = j.Usage
    if j.UseVstp != nil {
        structMap["use_vstp"] = j.UseVstp
    }
    if j.VoipNetwork != nil {
        structMap["voip_network"] = j.VoipNetwork
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JunosLocalPortConfig.
// It customizes the JSON unmarshaling process for JunosLocalPortConfig objects.
func (j *JunosLocalPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempJunosLocalPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "all_networks", "allow_dhcpd", "allow_multiple_supplicants", "bypass_auth_when_server_down", "bypass_auth_when_server_down_for_unknown_client", "description", "disable_autoneg", "disabled", "duplex", "dynamic_vlan_networks", "enable_mac_auth", "enable_qos", "guest_network", "inter_switch_link", "mac_auth_only", "mac_auth_preferred", "mac_auth_protocol", "mac_limit", "mode", "mtu", "networks", "note", "persist_mac", "poe_disabled", "port_auth", "port_network", "reauth_interval", "server_fail_network", "server_reject_network", "speed", "storm_control", "stp_edge", "stp_no_root_port", "stp_p2p", "usage", "use_vstp", "voip_network")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.AllNetworks = temp.AllNetworks
    j.AllowDhcpd = temp.AllowDhcpd
    j.AllowMultipleSupplicants = temp.AllowMultipleSupplicants
    j.BypassAuthWhenServerDown = temp.BypassAuthWhenServerDown
    j.BypassAuthWhenServerDownForUnknownClient = temp.BypassAuthWhenServerDownForUnknownClient
    j.Description = temp.Description
    j.DisableAutoneg = temp.DisableAutoneg
    j.Disabled = temp.Disabled
    j.Duplex = temp.Duplex
    j.DynamicVlanNetworks = temp.DynamicVlanNetworks
    j.EnableMacAuth = temp.EnableMacAuth
    j.EnableQos = temp.EnableQos
    j.GuestNetwork = temp.GuestNetwork
    j.InterSwitchLink = temp.InterSwitchLink
    j.MacAuthOnly = temp.MacAuthOnly
    j.MacAuthPreferred = temp.MacAuthPreferred
    j.MacAuthProtocol = temp.MacAuthProtocol
    j.MacLimit = temp.MacLimit
    j.Mode = temp.Mode
    j.Mtu = temp.Mtu
    j.Networks = temp.Networks
    j.Note = temp.Note
    j.PersistMac = temp.PersistMac
    j.PoeDisabled = temp.PoeDisabled
    j.PortAuth = temp.PortAuth
    j.PortNetwork = temp.PortNetwork
    j.ReauthInterval = temp.ReauthInterval
    j.ServerFailNetwork = temp.ServerFailNetwork
    j.ServerRejectNetwork = temp.ServerRejectNetwork
    j.Speed = temp.Speed
    j.StormControl = temp.StormControl
    j.StpEdge = temp.StpEdge
    j.StpNoRootPort = temp.StpNoRootPort
    j.StpP2p = temp.StpP2p
    j.Usage = *temp.Usage
    j.UseVstp = temp.UseVstp
    j.VoipNetwork = temp.VoipNetwork
    return nil
}

// tempJunosLocalPortConfig is a temporary struct used for validating the fields of JunosLocalPortConfig.
type tempJunosLocalPortConfig  struct {
    AllNetworks                              *bool                                    `json:"all_networks,omitempty"`
    AllowDhcpd                               *bool                                    `json:"allow_dhcpd,omitempty"`
    AllowMultipleSupplicants                 *bool                                    `json:"allow_multiple_supplicants,omitempty"`
    BypassAuthWhenServerDown                 *bool                                    `json:"bypass_auth_when_server_down,omitempty"`
    BypassAuthWhenServerDownForUnknownClient *bool                                    `json:"bypass_auth_when_server_down_for_unknown_client,omitempty"`
    Description                              *string                                  `json:"description,omitempty"`
    DisableAutoneg                           *bool                                    `json:"disable_autoneg,omitempty"`
    Disabled                                 *bool                                    `json:"disabled,omitempty"`
    Duplex                                   *SwitchPortLocalUsageDuplexEnum          `json:"duplex,omitempty"`
    DynamicVlanNetworks                      []string                                 `json:"dynamic_vlan_networks,omitempty"`
    EnableMacAuth                            *bool                                    `json:"enable_mac_auth,omitempty"`
    EnableQos                                *bool                                    `json:"enable_qos,omitempty"`
    GuestNetwork                             Optional[string]                         `json:"guest_network"`
    InterSwitchLink                          *bool                                    `json:"inter_switch_link,omitempty"`
    MacAuthOnly                              *bool                                    `json:"mac_auth_only,omitempty"`
    MacAuthPreferred                         *bool                                    `json:"mac_auth_preferred,omitempty"`
    MacAuthProtocol                          *SwitchPortLocalUsageMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    MacLimit                                 *int                                     `json:"mac_limit,omitempty"`
    Mode                                     *SwitchPortLocalUsageModeEnum            `json:"mode,omitempty"`
    Mtu                                      *int                                     `json:"mtu,omitempty"`
    Networks                                 []string                                 `json:"networks,omitempty"`
    Note                                     *string                                  `json:"note,omitempty"`
    PersistMac                               *bool                                    `json:"persist_mac,omitempty"`
    PoeDisabled                              *bool                                    `json:"poe_disabled,omitempty"`
    PortAuth                                 Optional[SwitchPortLocalUsageDot1xEnum]  `json:"port_auth"`
    PortNetwork                              *string                                  `json:"port_network,omitempty"`
    ReauthInterval                           *SwitchPortUsageReauthInterval           `json:"reauth_interval,omitempty"`
    ServerFailNetwork                        Optional[string]                         `json:"server_fail_network"`
    ServerRejectNetwork                      Optional[string]                         `json:"server_reject_network"`
    Speed                                    *JunosPortConfigSpeedEnum                `json:"speed,omitempty"`
    StormControl                             *SwitchPortLocalUsageStormControl        `json:"storm_control,omitempty"`
    StpEdge                                  *bool                                    `json:"stp_edge,omitempty"`
    StpNoRootPort                            *bool                                    `json:"stp_no_root_port,omitempty"`
    StpP2p                                   *bool                                    `json:"stp_p2p,omitempty"`
    Usage                                    *string                                  `json:"usage"`
    UseVstp                                  *bool                                    `json:"use_vstp,omitempty"`
    VoipNetwork                              *string                                  `json:"voip_network,omitempty"`
}

func (j *tempJunosLocalPortConfig) validate() error {
    var errs []string
    if j.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `junos_local_port_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
