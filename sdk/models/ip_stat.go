package models

import (
    "encoding/json"
)

// IpStat represents a IpStat struct.
type IpStat struct {
    Dns                  []string          `json:"dns,omitempty"`
    DnsSuffix            []string          `json:"dns_suffix,omitempty"`
    Gateway              *string           `json:"gateway,omitempty"`
    Gateway6             *string           `json:"gateway6,omitempty"`
    Ip                   *string           `json:"ip,omitempty"`
    Ip6                  *string           `json:"ip6,omitempty"`
    Ips                  map[string]string `json:"ips,omitempty"`
    Netmask              *string           `json:"netmask,omitempty"`
    Netmask6             *string           `json:"netmask6,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IpStat.
// It customizes the JSON marshaling process for IpStat objects.
func (i IpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IpStat object to a map representation for JSON marshaling.
func (i IpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Dns != nil {
        structMap["dns"] = i.Dns
    }
    if i.DnsSuffix != nil {
        structMap["dns_suffix"] = i.DnsSuffix
    }
    if i.Gateway != nil {
        structMap["gateway"] = i.Gateway
    }
    if i.Gateway6 != nil {
        structMap["gateway6"] = i.Gateway6
    }
    if i.Ip != nil {
        structMap["ip"] = i.Ip
    }
    if i.Ip6 != nil {
        structMap["ip6"] = i.Ip6
    }
    if i.Ips != nil {
        structMap["ips"] = i.Ips
    }
    if i.Netmask != nil {
        structMap["netmask"] = i.Netmask
    }
    if i.Netmask6 != nil {
        structMap["netmask6"] = i.Netmask6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IpStat.
// It customizes the JSON unmarshaling process for IpStat objects.
func (i *IpStat) UnmarshalJSON(input []byte) error {
    var temp ipStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dns", "dns_suffix", "gateway", "gateway6", "ip", "ip6", "ips", "netmask", "netmask6")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Dns = temp.Dns
    i.DnsSuffix = temp.DnsSuffix
    i.Gateway = temp.Gateway
    i.Gateway6 = temp.Gateway6
    i.Ip = temp.Ip
    i.Ip6 = temp.Ip6
    i.Ips = temp.Ips
    i.Netmask = temp.Netmask
    i.Netmask6 = temp.Netmask6
    return nil
}

// ipStat is a temporary struct used for validating the fields of IpStat.
type ipStat  struct {
    Dns       []string          `json:"dns,omitempty"`
    DnsSuffix []string          `json:"dns_suffix,omitempty"`
    Gateway   *string           `json:"gateway,omitempty"`
    Gateway6  *string           `json:"gateway6,omitempty"`
    Ip        *string           `json:"ip,omitempty"`
    Ip6       *string           `json:"ip6,omitempty"`
    Ips       map[string]string `json:"ips,omitempty"`
    Netmask   *string           `json:"netmask,omitempty"`
    Netmask6  *string           `json:"netmask6,omitempty"`
}
