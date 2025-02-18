package models

import (
    "encoding/json"
    "fmt"
)

// ApIpConfig represents a ApIpConfig struct.
// IP AP settings
type ApIpConfig struct {
    // If `type`==`static`
    Dns                  []string               `json:"dns,omitempty"`
    // Required if `type`==`static`
    DnsSuffix            []string               `json:"dns_suffix,omitempty"`
    // Required if `type`==`static`
    Gateway              *string                `json:"gateway,omitempty"`
    Gateway6             *string                `json:"gateway6,omitempty"`
    // Required if `type`==`static`
    Ip                   *string                `json:"ip,omitempty"`
    Ip6                  *string                `json:"ip6,omitempty"`
    Mtu                  *int                   `json:"mtu,omitempty"`
    // Required if `type`==`static`
    Netmask              *string                `json:"netmask,omitempty"`
    Netmask6             *string                `json:"netmask6,omitempty"`
    // enum: `dhcp`, `static`
    Type                 *IpTypeEnum            `json:"type,omitempty"`
    // enum: `autoconf`, `dhcp`, `disabled`, `static`
    Type6                *IpType6Enum           `json:"type6,omitempty"`
    // Management VLAN id, default is 1 (untagged)
    VlanId               *int                   `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApIpConfig) String() string {
    return fmt.Sprintf(
    	"ApIpConfig[Dns=%v, DnsSuffix=%v, Gateway=%v, Gateway6=%v, Ip=%v, Ip6=%v, Mtu=%v, Netmask=%v, Netmask6=%v, Type=%v, Type6=%v, VlanId=%v, AdditionalProperties=%v]",
    	a.Dns, a.DnsSuffix, a.Gateway, a.Gateway6, a.Ip, a.Ip6, a.Mtu, a.Netmask, a.Netmask6, a.Type, a.Type6, a.VlanId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApIpConfig.
// It customizes the JSON marshaling process for ApIpConfig objects.
func (a ApIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "dns", "dns_suffix", "gateway", "gateway6", "ip", "ip6", "mtu", "netmask", "netmask6", "type", "type6", "vlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApIpConfig object to a map representation for JSON marshaling.
func (a ApIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Dns != nil {
        structMap["dns"] = a.Dns
    }
    if a.DnsSuffix != nil {
        structMap["dns_suffix"] = a.DnsSuffix
    }
    if a.Gateway != nil {
        structMap["gateway"] = a.Gateway
    }
    if a.Gateway6 != nil {
        structMap["gateway6"] = a.Gateway6
    }
    if a.Ip != nil {
        structMap["ip"] = a.Ip
    }
    if a.Ip6 != nil {
        structMap["ip6"] = a.Ip6
    }
    if a.Mtu != nil {
        structMap["mtu"] = a.Mtu
    }
    if a.Netmask != nil {
        structMap["netmask"] = a.Netmask
    }
    if a.Netmask6 != nil {
        structMap["netmask6"] = a.Netmask6
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Type6 != nil {
        structMap["type6"] = a.Type6
    }
    if a.VlanId != nil {
        structMap["vlan_id"] = a.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIpConfig.
// It customizes the JSON unmarshaling process for ApIpConfig objects.
func (a *ApIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempApIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "dns_suffix", "gateway", "gateway6", "ip", "ip6", "mtu", "netmask", "netmask6", "type", "type6", "vlan_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Dns = temp.Dns
    a.DnsSuffix = temp.DnsSuffix
    a.Gateway = temp.Gateway
    a.Gateway6 = temp.Gateway6
    a.Ip = temp.Ip
    a.Ip6 = temp.Ip6
    a.Mtu = temp.Mtu
    a.Netmask = temp.Netmask
    a.Netmask6 = temp.Netmask6
    a.Type = temp.Type
    a.Type6 = temp.Type6
    a.VlanId = temp.VlanId
    return nil
}

// tempApIpConfig is a temporary struct used for validating the fields of ApIpConfig.
type tempApIpConfig  struct {
    Dns       []string     `json:"dns,omitempty"`
    DnsSuffix []string     `json:"dns_suffix,omitempty"`
    Gateway   *string      `json:"gateway,omitempty"`
    Gateway6  *string      `json:"gateway6,omitempty"`
    Ip        *string      `json:"ip,omitempty"`
    Ip6       *string      `json:"ip6,omitempty"`
    Mtu       *int         `json:"mtu,omitempty"`
    Netmask   *string      `json:"netmask,omitempty"`
    Netmask6  *string      `json:"netmask6,omitempty"`
    Type      *IpTypeEnum  `json:"type,omitempty"`
    Type6     *IpType6Enum `json:"type6,omitempty"`
    VlanId    *int         `json:"vlan_id,omitempty"`
}
