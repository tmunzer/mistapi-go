// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPortConfigIpConfig represents a GatewayPortConfigIpConfig struct.
// Junos IP Config
type GatewayPortConfigIpConfig struct {
    // Except for out-of_band interface (vme/em0/fxp0)
    Dns                  []string                `json:"dns,omitempty"`
    // Except for out-of_band interface (vme/em0/fxp0)
    DnsSuffix            []string                `json:"dns_suffix,omitempty"`
    // Except for out-of_band interface (vme/em0/fxp0). Interface Default Gateway IP Address (i.e. "192.168.1.1") or a Variable (i.e. "{{myvar}}")
    Gateway              *string                 `json:"gateway,omitempty"`
    // Except for out-of_band interface (vme/em0/fxp0). Interface Default Gateway IPv6 Address (i.e. "2001:db8::1") or a Variable (i.e. "{{myvar}}")
    Gateway6             *string                 `json:"gateway6,omitempty"`
    // Interface IP Address (i.e. "192.168.1.8") or a Variable (i.e. "{{myvar}}")
    Ip                   *string                 `json:"ip,omitempty"`
    // Interface IPv6 Address (i.e. "2001:db8::123") or a Variable (i.e. "{{myvar}}")
    Ipv6                 *string                 `json:"ipv6,omitempty"`
    // Used only if `subnet` is not specified in `networks`. Interface Netmask (i.e. "/24") or a Variable (i.e. "{{myvar}}")
    Netmask              *string                 `json:"netmask,omitempty"`
    // Used only if `subnet` is not specified in `networks`. Interface IPv6 Netmask (i.e. "/64") or a Variable (i.e. "{{myvar}}")
    Netmask6             *string                 `json:"netmask6,omitempty"`
    // Optional, the network to be used for mgmt
    Network              *string                 `json:"network,omitempty"`
    // If `type`==`pppoe`
    PoserPassword        *string                 `json:"poser_password,omitempty"`
    // if `type`==`pppoe`. enum: `chap`, `none`, `pap`
    PppoeAuth            *GatewayWanPpoeAuthEnum `json:"pppoe_auth,omitempty"`
    // If `type`==`pppoe`
    PppoeUsername        *string                 `json:"pppoe_username,omitempty"`
    // enum: `dhcp`, `pppoe`, `static`
    Type                 *GatewayWanTypeEnum     `json:"type,omitempty"`
    // enum: `autoconf`, `dhcp`, `static`
    Type6                *GatewayWanType6Enum    `json:"type6,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortConfigIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortConfigIpConfig) String() string {
    return fmt.Sprintf(
    	"GatewayPortConfigIpConfig[Dns=%v, DnsSuffix=%v, Gateway=%v, Gateway6=%v, Ip=%v, Ipv6=%v, Netmask=%v, Netmask6=%v, Network=%v, PoserPassword=%v, PppoeAuth=%v, PppoeUsername=%v, Type=%v, Type6=%v, AdditionalProperties=%v]",
    	g.Dns, g.DnsSuffix, g.Gateway, g.Gateway6, g.Ip, g.Ipv6, g.Netmask, g.Netmask6, g.Network, g.PoserPassword, g.PppoeAuth, g.PppoeUsername, g.Type, g.Type6, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortConfigIpConfig.
// It customizes the JSON marshaling process for GatewayPortConfigIpConfig objects.
func (g GatewayPortConfigIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "dns", "dns_suffix", "gateway", "gateway6", "ip", "ipv6", "netmask", "netmask6", "network", "poser_password", "pppoe_auth", "pppoe_username", "type", "type6"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortConfigIpConfig object to a map representation for JSON marshaling.
func (g GatewayPortConfigIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Dns != nil {
        structMap["dns"] = g.Dns
    }
    if g.DnsSuffix != nil {
        structMap["dns_suffix"] = g.DnsSuffix
    }
    if g.Gateway != nil {
        structMap["gateway"] = g.Gateway
    }
    if g.Gateway6 != nil {
        structMap["gateway6"] = g.Gateway6
    }
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.Ipv6 != nil {
        structMap["ipv6"] = g.Ipv6
    }
    if g.Netmask != nil {
        structMap["netmask"] = g.Netmask
    }
    if g.Netmask6 != nil {
        structMap["netmask6"] = g.Netmask6
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
    if g.Type6 != nil {
        structMap["type6"] = g.Type6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortConfigIpConfig.
// It customizes the JSON unmarshaling process for GatewayPortConfigIpConfig objects.
func (g *GatewayPortConfigIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortConfigIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "dns_suffix", "gateway", "gateway6", "ip", "ipv6", "netmask", "netmask6", "network", "poser_password", "pppoe_auth", "pppoe_username", "type", "type6")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Dns = temp.Dns
    g.DnsSuffix = temp.DnsSuffix
    g.Gateway = temp.Gateway
    g.Gateway6 = temp.Gateway6
    g.Ip = temp.Ip
    g.Ipv6 = temp.Ipv6
    g.Netmask = temp.Netmask
    g.Netmask6 = temp.Netmask6
    g.Network = temp.Network
    g.PoserPassword = temp.PoserPassword
    g.PppoeAuth = temp.PppoeAuth
    g.PppoeUsername = temp.PppoeUsername
    g.Type = temp.Type
    g.Type6 = temp.Type6
    return nil
}

// tempGatewayPortConfigIpConfig is a temporary struct used for validating the fields of GatewayPortConfigIpConfig.
type tempGatewayPortConfigIpConfig  struct {
    Dns           []string                `json:"dns,omitempty"`
    DnsSuffix     []string                `json:"dns_suffix,omitempty"`
    Gateway       *string                 `json:"gateway,omitempty"`
    Gateway6      *string                 `json:"gateway6,omitempty"`
    Ip            *string                 `json:"ip,omitempty"`
    Ipv6          *string                 `json:"ipv6,omitempty"`
    Netmask       *string                 `json:"netmask,omitempty"`
    Netmask6      *string                 `json:"netmask6,omitempty"`
    Network       *string                 `json:"network,omitempty"`
    PoserPassword *string                 `json:"poser_password,omitempty"`
    PppoeAuth     *GatewayWanPpoeAuthEnum `json:"pppoe_auth,omitempty"`
    PppoeUsername *string                 `json:"pppoe_username,omitempty"`
    Type          *GatewayWanTypeEnum     `json:"type,omitempty"`
    Type6         *GatewayWanType6Enum    `json:"type6,omitempty"`
}
