package models

import (
    "encoding/json"
)

// MxedgeStatsIpStat represents a MxedgeStatsIpStat struct.
// OOBM IP stats
type MxedgeStatsIpStat struct {
    Ip                   *string           `json:"ip,omitempty"`
    // Property key is the interface name. IPs for each net interface
    Ips                  map[string]string `json:"ips,omitempty"`
    // Property key is the interface name. MAC for each net interface
    Macs                 map[string]string `json:"macs,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsIpStat.
// It customizes the JSON marshaling process for MxedgeStatsIpStat objects.
func (m MxedgeStatsIpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsIpStat object to a map representation for JSON marshaling.
func (m MxedgeStatsIpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Ip != nil {
        structMap["ip"] = m.Ip
    }
    if m.Ips != nil {
        structMap["ips"] = m.Ips
    }
    if m.Macs != nil {
        structMap["macs"] = m.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsIpStat.
// It customizes the JSON unmarshaling process for MxedgeStatsIpStat objects.
func (m *MxedgeStatsIpStat) UnmarshalJSON(input []byte) error {
    var temp mxedgeStatsIpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "ips", "macs")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Ip = temp.Ip
    m.Ips = temp.Ips
    m.Macs = temp.Macs
    return nil
}

// mxedgeStatsIpStat is a temporary struct used for validating the fields of MxedgeStatsIpStat.
type mxedgeStatsIpStat  struct {
    Ip   *string           `json:"ip,omitempty"`
    Ips  map[string]string `json:"ips,omitempty"`
    Macs map[string]string `json:"macs,omitempty"`
}
