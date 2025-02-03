package models

import (
    "encoding/json"
    "fmt"
)

// JunosIpConfig represents a JunosIpConfig struct.
// Junos IP Config
type JunosIpConfig struct {
    Dns                  []string               `json:"dns,omitempty"`
    DnsSuffix            []string               `json:"dns_suffix,omitempty"`
    Gateway              *string                `json:"gateway,omitempty"`
    Ip                   *string                `json:"ip,omitempty"`
    // Used only if `subnet` is not specified in `networks`
    Netmask              *string                `json:"netmask,omitempty"`
    // Network where this mgmt IP reside, this will be used as default network for outbound-ssh, dns, ntp, dns, tacplus, radius, syslog, snmp
    Network              *string                `json:"network,omitempty"`
    // enum: `dhcp`, `static`
    Type                 *IpTypeEnum            `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JunosIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JunosIpConfig) String() string {
    return fmt.Sprintf(
    	"JunosIpConfig[Dns=%v, DnsSuffix=%v, Gateway=%v, Ip=%v, Netmask=%v, Network=%v, Type=%v, AdditionalProperties=%v]",
    	j.Dns, j.DnsSuffix, j.Gateway, j.Ip, j.Netmask, j.Network, j.Type, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JunosIpConfig.
// It customizes the JSON marshaling process for JunosIpConfig objects.
func (j JunosIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "dns", "dns_suffix", "gateway", "ip", "netmask", "network", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JunosIpConfig object to a map representation for JSON marshaling.
func (j JunosIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.Dns != nil {
        structMap["dns"] = j.Dns
    }
    if j.DnsSuffix != nil {
        structMap["dns_suffix"] = j.DnsSuffix
    }
    if j.Gateway != nil {
        structMap["gateway"] = j.Gateway
    }
    if j.Ip != nil {
        structMap["ip"] = j.Ip
    }
    if j.Netmask != nil {
        structMap["netmask"] = j.Netmask
    }
    if j.Network != nil {
        structMap["network"] = j.Network
    }
    if j.Type != nil {
        structMap["type"] = j.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JunosIpConfig.
// It customizes the JSON unmarshaling process for JunosIpConfig objects.
func (j *JunosIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempJunosIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "dns_suffix", "gateway", "ip", "netmask", "network", "type")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.Dns = temp.Dns
    j.DnsSuffix = temp.DnsSuffix
    j.Gateway = temp.Gateway
    j.Ip = temp.Ip
    j.Netmask = temp.Netmask
    j.Network = temp.Network
    j.Type = temp.Type
    return nil
}

// tempJunosIpConfig is a temporary struct used for validating the fields of JunosIpConfig.
type tempJunosIpConfig  struct {
    Dns       []string    `json:"dns,omitempty"`
    DnsSuffix []string    `json:"dns_suffix,omitempty"`
    Gateway   *string     `json:"gateway,omitempty"`
    Ip        *string     `json:"ip,omitempty"`
    Netmask   *string     `json:"netmask,omitempty"`
    Network   *string     `json:"network,omitempty"`
    Type      *IpTypeEnum `json:"type,omitempty"`
}
