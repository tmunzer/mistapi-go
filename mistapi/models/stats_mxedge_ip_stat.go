package models

import (
    "encoding/json"
)

// StatsMxedgeIpStat represents a StatsMxedgeIpStat struct.
// OOBM IP stats
type StatsMxedgeIpStat struct {
    Ip                   *string           `json:"ip,omitempty"`
    // Property key is the interface name. IPs for each net interface
    Ips                  map[string]string `json:"ips,omitempty"`
    // Property key is the interface name. MAC for each net interface
    Macs                 map[string]string `json:"macs,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeIpStat.
// It customizes the JSON marshaling process for StatsMxedgeIpStat objects.
func (s StatsMxedgeIpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeIpStat object to a map representation for JSON marshaling.
func (s StatsMxedgeIpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Ips != nil {
        structMap["ips"] = s.Ips
    }
    if s.Macs != nil {
        structMap["macs"] = s.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeIpStat.
// It customizes the JSON unmarshaling process for StatsMxedgeIpStat objects.
func (s *StatsMxedgeIpStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeIpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "ips", "macs")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Ip = temp.Ip
    s.Ips = temp.Ips
    s.Macs = temp.Macs
    return nil
}

// tempStatsMxedgeIpStat is a temporary struct used for validating the fields of StatsMxedgeIpStat.
type tempStatsMxedgeIpStat  struct {
    Ip   *string           `json:"ip,omitempty"`
    Ips  map[string]string `json:"ips,omitempty"`
    Macs map[string]string `json:"macs,omitempty"`
}
