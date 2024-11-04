package models

import (
    "encoding/json"
)

// StatsMxedgeTuntermIpConfig represents a StatsMxedgeTuntermIpConfig struct.
type StatsMxedgeTuntermIpConfig struct {
    Gateway              *string        `json:"gateway,omitempty"`
    Ip                   *string        `json:"ip,omitempty"`
    Netmask              *string        `json:"netmask,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeTuntermIpConfig.
// It customizes the JSON marshaling process for StatsMxedgeTuntermIpConfig objects.
func (s StatsMxedgeTuntermIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeTuntermIpConfig object to a map representation for JSON marshaling.
func (s StatsMxedgeTuntermIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Netmask != nil {
        structMap["netmask"] = s.Netmask
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeTuntermIpConfig.
// It customizes the JSON unmarshaling process for StatsMxedgeTuntermIpConfig objects.
func (s *StatsMxedgeTuntermIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeTuntermIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway", "ip", "netmask")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Gateway = temp.Gateway
    s.Ip = temp.Ip
    s.Netmask = temp.Netmask
    return nil
}

// tempStatsMxedgeTuntermIpConfig is a temporary struct used for validating the fields of StatsMxedgeTuntermIpConfig.
type tempStatsMxedgeTuntermIpConfig  struct {
    Gateway *string `json:"gateway,omitempty"`
    Ip      *string `json:"ip,omitempty"`
    Netmask *string `json:"netmask,omitempty"`
}
