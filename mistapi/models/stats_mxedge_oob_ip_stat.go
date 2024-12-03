package models

import (
    "encoding/json"
)

// StatsMxedgeOobIpStat represents a StatsMxedgeOobIpStat struct.
type StatsMxedgeOobIpStat struct {
    Dns                  []string                  `json:"dns,omitempty"`
    Gateway              *string                   `json:"gateway,omitempty"`
    Gateway6             *string                   `json:"gateway6,omitempty"`
    Ip                   *string                   `json:"ip,omitempty"`
    Ip6                  *string                   `json:"ip6,omitempty"`
    Netmask              *string                   `json:"netmask,omitempty"`
    Netmask6             *string                   `json:"netmask6,omitempty"`
    // enum: `dhcp`, `disabled`, `static`
    Type                 *MxedgeMgmtOobIpTypeEnum  `json:"type,omitempty"`
    // enum: `autoconf`, `dhcp`, `disabled`, `static`
    Type8                *MxedgeMgmtOobIpType6Enum `json:"type8,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeOobIpStat.
// It customizes the JSON marshaling process for StatsMxedgeOobIpStat objects.
func (s StatsMxedgeOobIpStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "dns", "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6", "type", "type8"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeOobIpStat object to a map representation for JSON marshaling.
func (s StatsMxedgeOobIpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Dns != nil {
        structMap["dns"] = s.Dns
    }
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway
    }
    if s.Gateway6 != nil {
        structMap["gateway6"] = s.Gateway6
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Ip6 != nil {
        structMap["ip6"] = s.Ip6
    }
    if s.Netmask != nil {
        structMap["netmask"] = s.Netmask
    }
    if s.Netmask6 != nil {
        structMap["netmask6"] = s.Netmask6
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Type8 != nil {
        structMap["type8"] = s.Type8
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeOobIpStat.
// It customizes the JSON unmarshaling process for StatsMxedgeOobIpStat objects.
func (s *StatsMxedgeOobIpStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeOobIpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns", "gateway", "gateway6", "ip", "ip6", "netmask", "netmask6", "type", "type8")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Dns = temp.Dns
    s.Gateway = temp.Gateway
    s.Gateway6 = temp.Gateway6
    s.Ip = temp.Ip
    s.Ip6 = temp.Ip6
    s.Netmask = temp.Netmask
    s.Netmask6 = temp.Netmask6
    s.Type = temp.Type
    s.Type8 = temp.Type8
    return nil
}

// tempStatsMxedgeOobIpStat is a temporary struct used for validating the fields of StatsMxedgeOobIpStat.
type tempStatsMxedgeOobIpStat  struct {
    Dns      []string                  `json:"dns,omitempty"`
    Gateway  *string                   `json:"gateway,omitempty"`
    Gateway6 *string                   `json:"gateway6,omitempty"`
    Ip       *string                   `json:"ip,omitempty"`
    Ip6      *string                   `json:"ip6,omitempty"`
    Netmask  *string                   `json:"netmask,omitempty"`
    Netmask6 *string                   `json:"netmask6,omitempty"`
    Type     *MxedgeMgmtOobIpTypeEnum  `json:"type,omitempty"`
    Type8    *MxedgeMgmtOobIpType6Enum `json:"type8,omitempty"`
}
