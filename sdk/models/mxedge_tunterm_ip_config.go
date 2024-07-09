package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxedgeTuntermIpConfig represents a MxedgeTuntermIpConfig struct.
// ip configuration of the Mist Tunnel interface
type MxedgeTuntermIpConfig struct {
    Gateway              string         `json:"gateway"`
    Gateway6             *string        `json:"gateway6,omitempty"`
    // untagged VLAN
    Ip                   string         `json:"ip"`
    Ip6                  *string        `json:"ip6,omitempty"`
    Netmask              string         `json:"netmask"`
    Netmask6             *string        `json:"netmask6,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIpConfig.
// It customizes the JSON marshaling process for MxedgeTuntermIpConfig objects.
func (m MxedgeTuntermIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIpConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp mxedgeTuntermIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6")
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

// mxedgeTuntermIpConfig is a temporary struct used for validating the fields of MxedgeTuntermIpConfig.
type mxedgeTuntermIpConfig  struct {
    Gateway  *string `json:"gateway"`
    Gateway6 *string `json:"gateway6,omitempty"`
    Ip       *string `json:"ip"`
    Ip6      *string `json:"ip6,omitempty"`
    Netmask  *string `json:"netmask"`
    Netmask6 *string `json:"netmask6,omitempty"`
}

func (m *mxedgeTuntermIpConfig) validate() error {
    var errs []string
    if m.Gateway == nil {
        errs = append(errs, "required field `gateway` is missing for type `Mxedge_Tunterm_Ip_Config`")
    }
    if m.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `Mxedge_Tunterm_Ip_Config`")
    }
    if m.Netmask == nil {
        errs = append(errs, "required field `netmask` is missing for type `Mxedge_Tunterm_Ip_Config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
