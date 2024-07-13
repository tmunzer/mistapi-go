package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApPortConfig represents a ApPortConfig struct.
type ApPortConfig struct {
    Disabled             *bool                            `json:"disabled,omitempty"`
    // optional dynamic vlan
    DynamicVlan          *ApPortConfigDynamicVlan         `json:"dynamic_vlan,omitempty"`
    EnableMacAuth        *bool                            `json:"enable_mac_auth,omitempty"`
    Forwarding           *ApPortConfigForwardingEnum      `json:"forwarding,omitempty"`
    // if `enable_mac_auth`==`true`, allows user to select an authentication protocol
    MacAuthProtocol      *ApPortConfigMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    MistNac              *WlanMistNac                     `json:"mist_nac,omitempty"`
    // if `forwarding`==`mxtunnel`, vlan_ids comes from mxtunnel
    MxTunnelId           *uuid.UUID                       `json:"mx_tunnel_id,omitempty"`
    // if `forwarding`==`site_mxedge`, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
    MxtunnelName         *string                          `json:"mxtunnel_name,omitempty"`
    // When doing port auth
    PortAuth             *ApPortConfigPortAuthEnum        `json:"port_auth,omitempty"`
    // if `forwrding`==`limited`
    PortVlanId           *int                             `json:"port_vlan_id,omitempty"`
    // Junos Radius config
    RadiusConfig         *RadiusConfig                    `json:"radius_config,omitempty"`
    // Radsec settings
    Radsec               *Radsec                          `json:"radsec,omitempty"`
    // optional to specify the vlan id for a tunnel if forwarding is for `wxtunnel`, `mxtunnel` or `site_mxedge`.
    // * if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.
    // * if forwarding == site_mxedge, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
    VlanId               *int                             `json:"vlan_id,omitempty"`
    // if `forwrding`==`limited`
    VlandIds             []int                            `json:"vland_ids,omitempty"`
    // if `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
    WxtunnelId           *uuid.UUID                       `json:"wxtunnel_id,omitempty"`
    // if `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
    WxtunnelRemoteId     *string                          `json:"wxtunnel_remote_id,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApPortConfig.
// It customizes the JSON marshaling process for ApPortConfig objects.
func (a ApPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApPortConfig object to a map representation for JSON marshaling.
func (a ApPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
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
    if a.VlandIds != nil {
        structMap["vland_ids"] = a.VlandIds
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
    var temp apPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disabled", "dynamic_vlan", "enable_mac_auth", "forwarding", "mac_auth_protocol", "mist_nac", "mx_tunnel_id", "mxtunnel_name", "port_auth", "port_vlan_id", "radius_config", "radsec", "vlan_id", "vland_ids", "wxtunnel_id", "wxtunnel_remote_id")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Disabled = temp.Disabled
    a.DynamicVlan = temp.DynamicVlan
    a.EnableMacAuth = temp.EnableMacAuth
    a.Forwarding = temp.Forwarding
    a.MacAuthProtocol = temp.MacAuthProtocol
    a.MistNac = temp.MistNac
    a.MxTunnelId = temp.MxTunnelId
    a.MxtunnelName = temp.MxtunnelName
    a.PortAuth = temp.PortAuth
    a.PortVlanId = temp.PortVlanId
    a.RadiusConfig = temp.RadiusConfig
    a.Radsec = temp.Radsec
    a.VlanId = temp.VlanId
    a.VlandIds = temp.VlandIds
    a.WxtunnelId = temp.WxtunnelId
    a.WxtunnelRemoteId = temp.WxtunnelRemoteId
    return nil
}

// apPortConfig is a temporary struct used for validating the fields of ApPortConfig.
type apPortConfig  struct {
    Disabled         *bool                            `json:"disabled,omitempty"`
    DynamicVlan      *ApPortConfigDynamicVlan         `json:"dynamic_vlan,omitempty"`
    EnableMacAuth    *bool                            `json:"enable_mac_auth,omitempty"`
    Forwarding       *ApPortConfigForwardingEnum      `json:"forwarding,omitempty"`
    MacAuthProtocol  *ApPortConfigMacAuthProtocolEnum `json:"mac_auth_protocol,omitempty"`
    MistNac          *WlanMistNac                     `json:"mist_nac,omitempty"`
    MxTunnelId       *uuid.UUID                       `json:"mx_tunnel_id,omitempty"`
    MxtunnelName     *string                          `json:"mxtunnel_name,omitempty"`
    PortAuth         *ApPortConfigPortAuthEnum        `json:"port_auth,omitempty"`
    PortVlanId       *int                             `json:"port_vlan_id,omitempty"`
    RadiusConfig     *RadiusConfig                    `json:"radius_config,omitempty"`
    Radsec           *Radsec                          `json:"radsec,omitempty"`
    VlanId           *int                             `json:"vlan_id,omitempty"`
    VlandIds         []int                            `json:"vland_ids,omitempty"`
    WxtunnelId       *uuid.UUID                       `json:"wxtunnel_id,omitempty"`
    WxtunnelRemoteId *string                          `json:"wxtunnel_remote_id,omitempty"`
}
