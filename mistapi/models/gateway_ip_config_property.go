package models

import (
    "encoding/json"
    "fmt"
)

// GatewayIpConfigProperty represents a GatewayIpConfigProperty struct.
type GatewayIpConfigProperty struct {
    Ip                   *string                `json:"ip,omitempty"`
    Netmask              *string                `json:"netmask,omitempty"`
    // Optional list of secondary IPs in CIDR format
    SecondaryIps         []string               `json:"secondary_ips,omitempty"`
    // enum: `dhcp`, `static`
    Type                 *IpTypeEnum            `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayIpConfigProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayIpConfigProperty) String() string {
    return fmt.Sprintf(
    	"GatewayIpConfigProperty[Ip=%v, Netmask=%v, SecondaryIps=%v, Type=%v, AdditionalProperties=%v]",
    	g.Ip, g.Netmask, g.SecondaryIps, g.Type, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayIpConfigProperty.
// It customizes the JSON marshaling process for GatewayIpConfigProperty objects.
func (g GatewayIpConfigProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "ip", "netmask", "secondary_ips", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayIpConfigProperty object to a map representation for JSON marshaling.
func (g GatewayIpConfigProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.Netmask != nil {
        structMap["netmask"] = g.Netmask
    }
    if g.SecondaryIps != nil {
        structMap["secondary_ips"] = g.SecondaryIps
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayIpConfigProperty.
// It customizes the JSON unmarshaling process for GatewayIpConfigProperty objects.
func (g *GatewayIpConfigProperty) UnmarshalJSON(input []byte) error {
    var temp tempGatewayIpConfigProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip", "netmask", "secondary_ips", "type")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Ip = temp.Ip
    g.Netmask = temp.Netmask
    g.SecondaryIps = temp.SecondaryIps
    g.Type = temp.Type
    return nil
}

// tempGatewayIpConfigProperty is a temporary struct used for validating the fields of GatewayIpConfigProperty.
type tempGatewayIpConfigProperty  struct {
    Ip           *string     `json:"ip,omitempty"`
    Netmask      *string     `json:"netmask,omitempty"`
    SecondaryIps []string    `json:"secondary_ips,omitempty"`
    Type         *IpTypeEnum `json:"type,omitempty"`
}
