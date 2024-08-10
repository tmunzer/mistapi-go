package models

import (
    "encoding/json"
)

// MxedgeStatsTuntermIpConfig represents a MxedgeStatsTuntermIpConfig struct.
type MxedgeStatsTuntermIpConfig struct {
    Gateway              *string        `json:"gateway,omitempty"`
    Ip                   *string        `json:"ip,omitempty"`
    Netmask              *string        `json:"netmask,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsTuntermIpConfig.
// It customizes the JSON marshaling process for MxedgeStatsTuntermIpConfig objects.
func (m MxedgeStatsTuntermIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsTuntermIpConfig object to a map representation for JSON marshaling.
func (m MxedgeStatsTuntermIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Gateway != nil {
        structMap["gateway"] = m.Gateway
    }
    if m.Ip != nil {
        structMap["ip"] = m.Ip
    }
    if m.Netmask != nil {
        structMap["netmask"] = m.Netmask
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsTuntermIpConfig.
// It customizes the JSON unmarshaling process for MxedgeStatsTuntermIpConfig objects.
func (m *MxedgeStatsTuntermIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeStatsTuntermIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway", "ip", "netmask")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Gateway = temp.Gateway
    m.Ip = temp.Ip
    m.Netmask = temp.Netmask
    return nil
}

// tempMxedgeStatsTuntermIpConfig is a temporary struct used for validating the fields of MxedgeStatsTuntermIpConfig.
type tempMxedgeStatsTuntermIpConfig  struct {
    Gateway *string `json:"gateway,omitempty"`
    Ip      *string `json:"ip,omitempty"`
    Netmask *string `json:"netmask,omitempty"`
}
