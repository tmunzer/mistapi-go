// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SwitchOobIpConfig represents a SwitchOobIpConfig struct.
// Switch OOB IP Config:
// - If HA configuration: key parameter will be nodeX (eg: node1)
// - If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`
type SwitchOobIpConfig struct {
    Gateway              *string                `json:"gateway,omitempty"`
    Ip                   *string                `json:"ip,omitempty"`
    // Used only if `subnet` is not specified in `networks`
    Netmask              *string                `json:"netmask,omitempty"`
    // Optional, the network to be used for mgmt
    Network              *string                `json:"network,omitempty"`
    // enum: `dhcp`, `static`
    Type                 *IpTypeEnum            `json:"type,omitempty"`
    // If supported on the platform. If enabled, DNS will be using this routing-instance, too
    UseMgmtVrf           *bool                  `json:"use_mgmt_vrf,omitempty"`
    // For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired
    UseMgmtVrfForHostOut *bool                  `json:"use_mgmt_vrf_for_host_out,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchOobIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchOobIpConfig) String() string {
    return fmt.Sprintf(
    	"SwitchOobIpConfig[Gateway=%v, Ip=%v, Netmask=%v, Network=%v, Type=%v, UseMgmtVrf=%v, UseMgmtVrfForHostOut=%v, AdditionalProperties=%v]",
    	s.Gateway, s.Ip, s.Netmask, s.Network, s.Type, s.UseMgmtVrf, s.UseMgmtVrfForHostOut, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchOobIpConfig.
// It customizes the JSON marshaling process for SwitchOobIpConfig objects.
func (s SwitchOobIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "gateway", "ip", "netmask", "network", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchOobIpConfig object to a map representation for JSON marshaling.
func (s SwitchOobIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Netmask != nil {
        structMap["netmask"] = s.Netmask
    }
    if s.Network != nil {
        structMap["network"] = s.Network
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.UseMgmtVrf != nil {
        structMap["use_mgmt_vrf"] = s.UseMgmtVrf
    }
    if s.UseMgmtVrfForHostOut != nil {
        structMap["use_mgmt_vrf_for_host_out"] = s.UseMgmtVrfForHostOut
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchOobIpConfig.
// It customizes the JSON unmarshaling process for SwitchOobIpConfig objects.
func (s *SwitchOobIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempSwitchOobIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateway", "ip", "netmask", "network", "type", "use_mgmt_vrf", "use_mgmt_vrf_for_host_out")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Gateway = temp.Gateway
    s.Ip = temp.Ip
    s.Netmask = temp.Netmask
    s.Network = temp.Network
    s.Type = temp.Type
    s.UseMgmtVrf = temp.UseMgmtVrf
    s.UseMgmtVrfForHostOut = temp.UseMgmtVrfForHostOut
    return nil
}

// tempSwitchOobIpConfig is a temporary struct used for validating the fields of SwitchOobIpConfig.
type tempSwitchOobIpConfig  struct {
    Gateway              *string     `json:"gateway,omitempty"`
    Ip                   *string     `json:"ip,omitempty"`
    Netmask              *string     `json:"netmask,omitempty"`
    Network              *string     `json:"network,omitempty"`
    Type                 *IpTypeEnum `json:"type,omitempty"`
    UseMgmtVrf           *bool       `json:"use_mgmt_vrf,omitempty"`
    UseMgmtVrfForHostOut *bool       `json:"use_mgmt_vrf_for_host_out,omitempty"`
}
