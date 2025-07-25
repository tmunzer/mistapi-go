// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MxedgeTuntermIpConfig represents a MxedgeTuntermIpConfig struct.
// IPconfiguration of the Mist Tunnel interface
type MxedgeTuntermIpConfig struct {
    Gateway              string                 `json:"gateway"`
    Gateway6             *string                `json:"gateway6,omitempty"`
    // Untagged VLAN
    Ip                   string                 `json:"ip"`
    Ip6                  *string                `json:"ip6,omitempty"`
    Netmask              string                 `json:"netmask"`
    Netmask6             *string                `json:"netmask6,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermIpConfig) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermIpConfig[Gateway=%v, Gateway6=%v, Ip=%v, Ip6=%v, Netmask=%v, Netmask6=%v, AdditionalProperties=%v]",
    	m.Gateway, m.Gateway6, m.Ip, m.Ip6, m.Netmask, m.Netmask6, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIpConfig.
// It customizes the JSON marshaling process for MxedgeTuntermIpConfig objects.
func (m MxedgeTuntermIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIpConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["gateway"] = m.Gateway
    if m.Gateway6 != nil {
        structMap["gateway6"] = m.Gateway6
    }
    structMap["ip"] = m.Ip
    if m.Ip6 != nil {
        structMap["ip6"] = m.Ip6
    }
    structMap["netmask"] = m.Netmask
    if m.Netmask6 != nil {
        structMap["netmask6"] = m.Netmask6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermIpConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermIpConfig objects.
func (m *MxedgeTuntermIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Gateway = *temp.Gateway
    m.Gateway6 = temp.Gateway6
    m.Ip = *temp.Ip
    m.Ip6 = temp.Ip6
    m.Netmask = *temp.Netmask
    m.Netmask6 = temp.Netmask6
    return nil
}

// tempMxedgeTuntermIpConfig is a temporary struct used for validating the fields of MxedgeTuntermIpConfig.
type tempMxedgeTuntermIpConfig  struct {
    Gateway  *string `json:"gateway"`
    Gateway6 *string `json:"gateway6,omitempty"`
    Ip       *string `json:"ip"`
    Ip6      *string `json:"ip6,omitempty"`
    Netmask  *string `json:"netmask"`
    Netmask6 *string `json:"netmask6,omitempty"`
}

func (m *tempMxedgeTuntermIpConfig) validate() error {
    var errs []string
    if m.Gateway == nil {
        errs = append(errs, "required field `gateway` is missing for type `mxedge_tunterm_ip_config`")
    }
    if m.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `mxedge_tunterm_ip_config`")
    }
    if m.Netmask == nil {
        errs = append(errs, "required field `netmask` is missing for type `mxedge_tunterm_ip_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
