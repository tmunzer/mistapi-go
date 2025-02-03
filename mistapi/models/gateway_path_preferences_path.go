package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPathPreferencesPath represents a GatewayPathPreferencesPath struct.
type GatewayPathPreferencesPath struct {
    Cost                 *int                   `json:"cost,omitempty"`
    // For SSR Only. `true`, if this specific path is undesired
    Disabled             *bool                  `json:"disabled,omitempty"`
    // Only if `type`==`local`, if a different gateway is desired
    GatewayIp            *string                `json:"gateway_ip,omitempty"`
    // Only if `type`==`vpn`, if this vpn path can be used for internet
    InternetAccess       *bool                  `json:"internet_access,omitempty"`
    // Required when
    // * `type`==`vpn`: the name of the VPN Path to use
    // * `type`==`wan`: the name of the WAN interface to use'
    Name                 *string                `json:"name,omitempty"`
    // Required when `type`==`local`
    Networks             []string               `json:"networks,omitempty"`
    // If `type`==`local`, if destination IP is to be replaced
    TargetIps            []string               `json:"target_ips,omitempty"`
    // enum: `local`, `tunnel`, `vpn`, `wan`
    Type                 *GatewayPathTypeEnum   `json:"type,omitempty"`
    // Optional if `type`==`vpn`
    WanName              *string                `json:"wan_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPathPreferencesPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPathPreferencesPath) String() string {
    return fmt.Sprintf(
    	"GatewayPathPreferencesPath[Cost=%v, Disabled=%v, GatewayIp=%v, InternetAccess=%v, Name=%v, Networks=%v, TargetIps=%v, Type=%v, WanName=%v, AdditionalProperties=%v]",
    	g.Cost, g.Disabled, g.GatewayIp, g.InternetAccess, g.Name, g.Networks, g.TargetIps, g.Type, g.WanName, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPathPreferencesPath.
// It customizes the JSON marshaling process for GatewayPathPreferencesPath objects.
func (g GatewayPathPreferencesPath) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "cost", "disabled", "gateway_ip", "internet_access", "name", "networks", "target_ips", "type", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPathPreferencesPath object to a map representation for JSON marshaling.
func (g GatewayPathPreferencesPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
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
    var temp tempGatewayPathPreferencesPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cost", "disabled", "gateway_ip", "internet_access", "name", "networks", "target_ips", "type", "wan_name")
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

// tempGatewayPathPreferencesPath is a temporary struct used for validating the fields of GatewayPathPreferencesPath.
type tempGatewayPathPreferencesPath  struct {
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
