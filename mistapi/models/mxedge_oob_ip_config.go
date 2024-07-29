package models

import (
    "encoding/json"
)

// MxedgeOobIpConfig represents a MxedgeOobIpConfig struct.
// ip configuration of the Mist Edge out-of_band management interface
type MxedgeOobIpConfig struct {
    Autoconf6            *bool          `json:"autoconf6,omitempty"`
    Dhcp6                *bool          `json:"dhcp6,omitempty"`
    // IPv4 ignored if `type`!=`static`
    // IPv6 ignored if `type6`!=`static`
    Dns                  []string       `json:"dns,omitempty"`
    // if `type`=`static`
    Gateway              *string        `json:"gateway,omitempty"`
    Gateway6             *string        `json:"gateway6,omitempty"`
    // if `type`=`static`
    Ip                   *string        `json:"ip,omitempty"`
    Ip6                  *string        `json:"ip6,omitempty"`
    // if `type`=`static`
    Netmask              *string        `json:"netmask,omitempty"`
    Netmask6             *string        `json:"netmask6,omitempty"`
    // enum: `dhcp`, `static`
    Type                 *IpTypeEnum    `json:"type,omitempty"`
    // enum: `dhcp`, `static`
    Type6                *IpTypeEnum    `json:"type6,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeOobIpConfig.
// It customizes the JSON marshaling process for MxedgeOobIpConfig objects.
func (m MxedgeOobIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeOobIpConfig object to a map representation for JSON marshaling.
func (m MxedgeOobIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Autoconf6 != nil {
        structMap["autoconf6"] = m.Autoconf6
    }
    if m.Dhcp6 != nil {
        structMap["dhcp6"] = m.Dhcp6
    }
    if m.Dns != nil {
        structMap["dns"] = m.Dns
    }
    if m.Gateway != nil {
        structMap["gateway"] = m.Gateway
    }
    if m.Gateway6 != nil {
        structMap["gateway6"] = m.Gateway6
    }
    if m.Ip != nil {
        structMap["ip"] = m.Ip
    }
    if m.Ip6 != nil {
        structMap["ip6"] = m.Ip6
    }
    if m.Netmask != nil {
        structMap["netmask"] = m.Netmask
    }
    if m.Netmask6 != nil {
        structMap["netmask6"] = m.Netmask6
    }
    if m.Type != nil {
        structMap["type"] = m.Type
    }
    if m.Type6 != nil {
        structMap["type6"] = m.Type6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeOobIpConfig.
// It customizes the JSON unmarshaling process for MxedgeOobIpConfig objects.
func (m *MxedgeOobIpConfig) UnmarshalJSON(input []byte) error {
    var temp mxedgeOobIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "autoconf6", "dhcp6", "dns", "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6", "type", "type6")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Autoconf6 = temp.Autoconf6
    m.Dhcp6 = temp.Dhcp6
    m.Dns = temp.Dns
    m.Gateway = temp.Gateway
    m.Gateway6 = temp.Gateway6
    m.Ip = temp.Ip
    m.Ip6 = temp.Ip6
    m.Netmask = temp.Netmask
    m.Netmask6 = temp.Netmask6
    m.Type = temp.Type
    m.Type6 = temp.Type6
    return nil
}

// mxedgeOobIpConfig is a temporary struct used for validating the fields of MxedgeOobIpConfig.
type mxedgeOobIpConfig  struct {
    Autoconf6 *bool       `json:"autoconf6,omitempty"`
    Dhcp6     *bool       `json:"dhcp6,omitempty"`
    Dns       []string    `json:"dns,omitempty"`
    Gateway   *string     `json:"gateway,omitempty"`
    Gateway6  *string     `json:"gateway6,omitempty"`
    Ip        *string     `json:"ip,omitempty"`
    Ip6       *string     `json:"ip6,omitempty"`
    Netmask   *string     `json:"netmask,omitempty"`
    Netmask6  *string     `json:"netmask6,omitempty"`
    Type      *IpTypeEnum `json:"type,omitempty"`
    Type6     *IpTypeEnum `json:"type6,omitempty"`
}
