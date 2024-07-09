package models

import (
    "encoding/json"
)

// ApIpConfig1 represents a ApIpConfig1 struct.
// IP AP settings
type ApIpConfig1 struct {
    // if `type`==`static`
    Dns                  []string       `json:"dns,omitempty"`
    // required if `type`==`static`
    DnsSuffix            []string       `json:"dns_suffix,omitempty"`
    // required if `type`==`static`
    Gateway              *string        `json:"gateway,omitempty"`
    Gateway6             *string        `json:"gateway6,omitempty"`
    // required if `type`==`static`
    Ip                   *string        `json:"ip,omitempty"`
    Ip6                  *string        `json:"ip6,omitempty"`
    Mtu                  *int           `json:"mtu,omitempty"`
    // required if `type`==`static`
    Netmask              *string        `json:"netmask,omitempty"`
    Netmask6             *string        `json:"netmask6,omitempty"`
    Type                 *IpTypeEnum    `json:"type,omitempty"`
    Type6                *IpType6Enum   `json:"type6,omitempty"`
    // management vlan id, default is 1 (untagged)
    VlanId               *int           `json:"vlan_id,omitempty"`
    // the network where this mgmt IP reside, this will be used as default network for outbound-ssh, dns, ntp, dns, tacplus, radius, syslog, snmp
    Network              *string        `json:"network,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApIpConfig1.
// It customizes the JSON marshaling process for ApIpConfig1 objects.
func (a ApIpConfig1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApIpConfig1 object to a map representation for JSON marshaling.
func (a ApIpConfig1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
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
    if a.Network != nil {
        structMap["network"] = a.Network
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIpConfig1.
// It customizes the JSON unmarshaling process for ApIpConfig1 objects.
func (a *ApIpConfig1) UnmarshalJSON(input []byte) error {
    var temp apIpConfig1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dns", "dns_suffix", "gateway", "gateway6", "ip", "ip6", "mtu", "netmask", "netmask6", "type", "type6", "vlan_id", "network")
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
    a.Network = temp.Network
    return nil
}

// apIpConfig1 is a temporary struct used for validating the fields of ApIpConfig1.
type apIpConfig1  struct {
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
    Network   *string      `json:"network,omitempty"`
}
