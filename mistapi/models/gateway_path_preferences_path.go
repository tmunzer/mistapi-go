package models

import (
    "encoding/json"
)

// GatewayPathPreferencesPath represents a GatewayPathPreferencesPath struct.
type GatewayPathPreferencesPath struct {
    Cost                 *int                 `json:"cost,omitempty"`
    // For SSR Only. `true`, if this specific path is undesired
    Disabled             *bool                `json:"disabled,omitempty"`
    // only if `type`==`local`, if a different gateway is desired
    GatewayIp            *string              `json:"gateway_ip,omitempty"`
    // only if `type`==`vpn`, if this vpn path can be used for internet
    InternetAccess       *bool                `json:"internet_access,omitempty"`
    // required when 
    //   * `type`==`vpn`: the name of the VPN Path to use 
    //   * `type`==`wan`: the name of the WAN interface to use'
    Name                 *string              `json:"name,omitempty"`
    // required when `type`==`local`
    Networks             []string             `json:"networks,omitempty"`
    // if `type`==`local`, if destination IP is to be replaced
    TargetIps            []string             `json:"target_ips,omitempty"`
    // enum: `local`, `tunnel`, `vpn`, `wan`
    Type                 *GatewayPathTypeEnum `json:"type,omitempty"`
    // required when`type`==`tunnel`, optional if `type`==`vpn` wan
    WanName              *string              `json:"wan_name,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPathPreferencesPath.
// It customizes the JSON marshaling process for GatewayPathPreferencesPath objects.
func (g GatewayPathPreferencesPath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPathPreferencesPath object to a map representation for JSON marshaling.
func (g GatewayPathPreferencesPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Cost != nil {
        structMap["cost"] = g.Cost
    }
    if g.Disabled != nil {
        structMap["disabled"] = g.Disabled
    }
    if g.GatewayIp != nil {
        structMap["gateway_ip"] = g.GatewayIp
    }
    if g.InternetAccess != nil {
        structMap["internet_access"] = g.InternetAccess
    }
    if g.Name != nil {
        structMap["name"] = g.Name
    }
    if g.Networks != nil {
        structMap["networks"] = g.Networks
    }
    if g.TargetIps != nil {
        structMap["target_ips"] = g.TargetIps
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    if g.WanName != nil {
        structMap["wan_name"] = g.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPathPreferencesPath.
// It customizes the JSON unmarshaling process for GatewayPathPreferencesPath objects.
func (g *GatewayPathPreferencesPath) UnmarshalJSON(input []byte) error {
    var temp gatewayPathPreferencesPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cost", "disabled", "gateway_ip", "internet_access", "name", "networks", "target_ips", "type", "wan_name")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Cost = temp.Cost
    g.Disabled = temp.Disabled
    g.GatewayIp = temp.GatewayIp
    g.InternetAccess = temp.InternetAccess
    g.Name = temp.Name
    g.Networks = temp.Networks
    g.TargetIps = temp.TargetIps
    g.Type = temp.Type
    g.WanName = temp.WanName
    return nil
}

// gatewayPathPreferencesPath is a temporary struct used for validating the fields of GatewayPathPreferencesPath.
type gatewayPathPreferencesPath  struct {
    Cost           *int                 `json:"cost,omitempty"`
    Disabled       *bool                `json:"disabled,omitempty"`
    GatewayIp      *string              `json:"gateway_ip,omitempty"`
    InternetAccess *bool                `json:"internet_access,omitempty"`
    Name           *string              `json:"name,omitempty"`
    Networks       []string             `json:"networks,omitempty"`
    TargetIps      []string             `json:"target_ips,omitempty"`
    Type           *GatewayPathTypeEnum `json:"type,omitempty"`
    WanName        *string              `json:"wan_name,omitempty"`
}
