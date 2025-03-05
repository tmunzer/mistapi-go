package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ApPortConfig represents a ApPortConfig struct.
type ApPortConfig struct {
    Disabled             *bool                            `json:"disabled,omitempty"`
    // Optional dynamic vlan
    DynamicVlan          *ApPortConfigDynamicVlan         `json:"dynamic_vlan,omitempty"`
    EnableMacAuth        *bool                            `json:"enable_mac_auth,omitempty"`
    // enum:
    // * `all`: local breakout, All VLANs
    // * `limited`: local breakout, only the VLANs configured in `port_vlan_id` and `vlan_ids`
    // * `mxtunnel`: central breakout to an Org Mist Edge (requires `mxtunnel_id`)
    // * `site_mxedge`: central breakout to a Site Mist Edge (requires `mxtunnel_name`)
    // * `wxtunnel`': central breakout to an Org WxTunnel (requires `wxtunnel_id`)
    Forwarding           *ApPortConfigForwardingEnum      `json:"forwarding,omitempty"`
    // When `true`, we'll do dot1x then mac_auth. enable this to prefer mac_auth
    MacAuthPreferred     *bool                            `json:"mac_auth_preferred,omitempty"`
    // if `enable_mac_auth`==`true`, allows user to select an authentication protocol. enum: `eap-md5`, `eap-peap`, `pap`
    MacAuthProtocol      *ApPortConfigMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    MistNac              *WlanMistNac                     `json:"mist_nac,omitempty"`
    // If `forwarding`==`mxtunnel`, vlan_ids comes from mxtunnel
    MxTunnelId           *uuid.UUID                       `json:"mx_tunnel_id,omitempty"`
    // If `forwarding`==`site_mxedge`, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
    MxtunnelName         *string                          `json:"mxtunnel_name,omitempty"`
    // for Q-in-Q configuration
    OuterVlanId          *int                             `json:"outer_vlan_id,omitempty"`
    // When doing port auth. enum: `dot1x`, `none`
    PortAuth             *ApPortConfigPortAuthEnum        `json:"port_auth,omitempty"`
    // If `forwarding`==`limited`
    PortVlanId           *int                             `json:"port_vlan_id,omitempty"`
    // Junos Radius config
    RadiusConfig         *RadiusConfig                    `json:"radius_config,omitempty"`
    // RadSec settings
    Radsec               *Radsec                          `json:"radsec,omitempty"`
    // Optional to specify the vlan id for a tunnel if forwarding is for `wxtunnel`, `mxtunnel` or `site_mxedge`.
    // * if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.
    // * if forwarding == site_mxedge, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
    VlanId               *int                             `json:"vlan_id,omitempty"`
    // If `forwarding`==`limited`
    VlanIds              []int                            `json:"vlan_ids,omitempty"`
    // If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
    WxtunnelId           *uuid.UUID                       `json:"wxtunnel_id,omitempty"`
    // If `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
    WxtunnelRemoteId     *string                          `json:"wxtunnel_remote_id,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for ApPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApPortConfig) String() string {
    return fmt.Sprintf(
    	"ApPortConfig[Disabled=%v, DynamicVlan=%v, EnableMacAuth=%v, Forwarding=%v, MacAuthPreferred=%v, MacAuthProtocol=%v, MistNac=%v, MxTunnelId=%v, MxtunnelName=%v, OuterVlanId=%v, PortAuth=%v, PortVlanId=%v, RadiusConfig=%v, Radsec=%v, VlanId=%v, VlanIds=%v, WxtunnelId=%v, WxtunnelRemoteId=%v, AdditionalProperties=%v]",
    	a.Disabled, a.DynamicVlan, a.EnableMacAuth, a.Forwarding, a.MacAuthPreferred, a.MacAuthProtocol, a.MistNac, a.MxTunnelId, a.MxtunnelName, a.OuterVlanId, a.PortAuth, a.PortVlanId, a.RadiusConfig, a.Radsec, a.VlanId, a.VlanIds, a.WxtunnelId, a.WxtunnelRemoteId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApPortConfig.
// It customizes the JSON marshaling process for ApPortConfig objects.
func (a ApPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "disabled", "dynamic_vlan", "enable_mac_auth", "forwarding", "mac_auth_preferred", "mac_auth_protocol", "mist_nac", "mx_tunnel_id", "mxtunnel_name", "outer_vlan_id", "port_auth", "port_vlan_id", "radius_config", "radsec", "vlan_id", "vlan_ids", "wxtunnel_id", "wxtunnel_remote_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApPortConfig object to a map representation for JSON marshaling.
func (a ApPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Disabled != nil {
        structMap["disabled"] = a.Disabled
    }
    if a.DynamicVlan != nil {
        structMap["dynamic_vlan"] = a.DynamicVlan.toMap()
    }
    if a.EnableMacAuth != nil {
        structMap["enable_mac_auth"] = a.EnableMacAuth
    }
    if a.Forwarding != nil {
        structMap["forwarding"] = a.Forwarding
    }
    if a.MacAuthPreferred != nil {
        structMap["mac_auth_preferred"] = a.MacAuthPreferred
    }
    if a.MacAuthProtocol != nil {
        structMap["mac_auth_protocol"] = a.MacAuthProtocol
    }
    if a.MistNac != nil {
        structMap["mist_nac"] = a.MistNac.toMap()
    }
    if a.MxTunnelId != nil {
        structMap["mx_tunnel_id"] = a.MxTunnelId
    }
    if a.MxtunnelName != nil {
        structMap["mxtunnel_name"] = a.MxtunnelName
    }
    if a.OuterVlanId != nil {
        structMap["outer_vlan_id"] = a.OuterVlanId
    }
    if a.PortAuth != nil {
        structMap["port_auth"] = a.PortAuth
    }
    if a.PortVlanId != nil {
        structMap["port_vlan_id"] = a.PortVlanId
    }
    if a.RadiusConfig != nil {
        structMap["radius_config"] = a.RadiusConfig.toMap()
    }
    if a.Radsec != nil {
        structMap["radsec"] = a.Radsec.toMap()
    }
    if a.VlanId != nil {
        structMap["vlan_id"] = a.VlanId
    }
    if a.VlanIds != nil {
        structMap["vlan_ids"] = a.VlanIds
    }
    if a.WxtunnelId != nil {
        structMap["wxtunnel_id"] = a.WxtunnelId
    }
    if a.WxtunnelRemoteId != nil {
        structMap["wxtunnel_remote_id"] = a.WxtunnelRemoteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApPortConfig.
// It customizes the JSON unmarshaling process for ApPortConfig objects.
func (a *ApPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempApPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "dynamic_vlan", "enable_mac_auth", "forwarding", "mac_auth_preferred", "mac_auth_protocol", "mist_nac", "mx_tunnel_id", "mxtunnel_name", "outer_vlan_id", "port_auth", "port_vlan_id", "radius_config", "radsec", "vlan_id", "vlan_ids", "wxtunnel_id", "wxtunnel_remote_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Disabled = temp.Disabled
    a.DynamicVlan = temp.DynamicVlan
    a.EnableMacAuth = temp.EnableMacAuth
    a.Forwarding = temp.Forwarding
    a.MacAuthPreferred = temp.MacAuthPreferred
    a.MacAuthProtocol = temp.MacAuthProtocol
    a.MistNac = temp.MistNac
    a.MxTunnelId = temp.MxTunnelId
    a.MxtunnelName = temp.MxtunnelName
    a.OuterVlanId = temp.OuterVlanId
    a.PortAuth = temp.PortAuth
    a.PortVlanId = temp.PortVlanId
    a.RadiusConfig = temp.RadiusConfig
    a.Radsec = temp.Radsec
    a.VlanId = temp.VlanId
    a.VlanIds = temp.VlanIds
    a.WxtunnelId = temp.WxtunnelId
    a.WxtunnelRemoteId = temp.WxtunnelRemoteId
    return nil
}

// tempApPortConfig is a temporary struct used for validating the fields of ApPortConfig.
type tempApPortConfig  struct {
    Disabled         *bool                            `json:"disabled,omitempty"`
    DynamicVlan      *ApPortConfigDynamicVlan         `json:"dynamic_vlan,omitempty"`
    EnableMacAuth    *bool                            `json:"enable_mac_auth,omitempty"`
    Forwarding       *ApPortConfigForwardingEnum      `json:"forwarding,omitempty"`
    MacAuthPreferred *bool                            `json:"mac_auth_preferred,omitempty"`
    MacAuthProtocol  *ApPortConfigMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    MistNac          *WlanMistNac                     `json:"mist_nac,omitempty"`
    MxTunnelId       *uuid.UUID                       `json:"mx_tunnel_id,omitempty"`
    MxtunnelName     *string                          `json:"mxtunnel_name,omitempty"`
    OuterVlanId      *int                             `json:"outer_vlan_id,omitempty"`
    PortAuth         *ApPortConfigPortAuthEnum        `json:"port_auth,omitempty"`
    PortVlanId       *int                             `json:"port_vlan_id,omitempty"`
    RadiusConfig     *RadiusConfig                    `json:"radius_config,omitempty"`
    Radsec           *Radsec                          `json:"radsec,omitempty"`
    VlanId           *int                             `json:"vlan_id,omitempty"`
    VlanIds          []int                            `json:"vlan_ids,omitempty"`
    WxtunnelId       *uuid.UUID                       `json:"wxtunnel_id,omitempty"`
    WxtunnelRemoteId *string                          `json:"wxtunnel_remote_id,omitempty"`
}
