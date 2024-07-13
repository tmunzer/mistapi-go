package models

import (
    "encoding/json"
)

// GatewayPortConfigIpConfig represents a GatewayPortConfigIpConfig struct.
// Junos IP Config
type GatewayPortConfigIpConfig struct {
    // except for out-of_band interface (vme/em0/fxp0)
    Dns                  []string                `json:"dns,omitempty"`
    // except for out-of_band interface (vme/em0/fxp0)
    DnsSuffix            []string                `json:"dns_suffix,omitempty"`
    // except for out-of_band interface (vme/em0/fxp0)
    Gateway              *string                 `json:"gateway,omitempty"`
    Ip                   *string                 `json:"ip,omitempty"`
    // used only if `subnet` is not specified in `networks`
    Netmask              *string                 `json:"netmask,omitempty"`
    // optional, the network to be used for mgmt
    Network              *string                 `json:"network,omitempty"`
    // if `type`==`pppoe`
    PoserPassword        *string                 `json:"poser_password,omitempty"`
    // if `type`==`pppoe`
    PppoeAuth            *GatewayWanPpoeAuthEnum `json:"pppoe_auth,omitempty"`
    // if `type`==`pppoe`
    PppoeUsername        *string                 `json:"pppoe_username,omitempty"`
    Type                 *GatewayWanTypeEnum     `json:"type,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfigIpConfig.
// It customizes the JSON marshaling process for GatewayPortConfigIpConfig objects.
func (g GatewayPortConfigIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortConfigIpConfig object to a map representation for JSON marshaling.
func (g GatewayPortConfigIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Dns != nil {
        structMap["dns"] = g.Dns
    }
    if g.DnsSuffix != nil {
        structMap["dns_suffix"] = g.DnsSuffix
    }
    if g.Gateway != nil {
        structMap["gateway"] = g.Gateway
    }
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.Netmask != nil {
        structMap["netmask"] = g.Netmask
    }
    if g.Network != nil {
        structMap["network"] = g.Network
    }
    if g.PoserPassword != nil {
        structMap["poser_password"] = g.PoserPassword
    }
    if g.PppoeAuth != nil {
        structMap["pppoe_auth"] = g.PppoeAuth
    }
    if g.PppoeUsername != nil {
        structMap["pppoe_username"] = g.PppoeUsername
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortConfigIpConfig.
// It customizes the JSON unmarshaling process for GatewayPortConfigIpConfig objects.
func (g *GatewayPortConfigIpConfig) UnmarshalJSON(input []byte) error {
    var temp gatewayPortConfigIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dns", "dns_suffix", "gateway", "ip", "netmask", "network", "poser_password", "pppoe_auth", "pppoe_username", "type")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Dns = temp.Dns
    g.DnsSuffix = temp.DnsSuffix
    g.Gateway = temp.Gateway
    g.Ip = temp.Ip
    g.Netmask = temp.Netmask
    g.Network = temp.Network
    g.PoserPassword = temp.PoserPassword
    g.PppoeAuth = temp.PppoeAuth
    g.PppoeUsername = temp.PppoeUsername
    g.Type = temp.Type
    return nil
}

// gatewayPortConfigIpConfig is a temporary struct used for validating the fields of GatewayPortConfigIpConfig.
type gatewayPortConfigIpConfig  struct {
    Dns           []string                `json:"dns,omitempty"`
    DnsSuffix     []string                `json:"dns_suffix,omitempty"`
    Gateway       *string                 `json:"gateway,omitempty"`
    Ip            *string                 `json:"ip,omitempty"`
    Netmask       *string                 `json:"netmask,omitempty"`
    Network       *string                 `json:"network,omitempty"`
    PoserPassword *string                 `json:"poser_password,omitempty"`
    PppoeAuth     *GatewayWanPpoeAuthEnum `json:"pppoe_auth,omitempty"`
    PppoeUsername *string                 `json:"pppoe_username,omitempty"`
    Type          *GatewayWanTypeEnum     `json:"type,omitempty"`
}