package models

import (
    "encoding/json"
)

// SwitchStatsIfStat represents a SwitchStatsIfStat struct.
type SwitchStatsIfStat struct {
    Ips                  []string       `json:"ips,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    RxBytes              *int           `json:"rx_bytes,omitempty"`
    RxPkts               *int           `json:"rx_pkts,omitempty"`
    TxBytes              *int           `json:"tx_bytes,omitempty"`
    TxPkts               *int           `json:"tx_pkts,omitempty"`
    Up                   *bool          `json:"up,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsIfStat.
// It customizes the JSON marshaling process for SwitchStatsIfStat objects.
func (s SwitchStatsIfStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsIfStat object to a map representation for JSON marshaling.
func (s SwitchStatsIfStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Ips != nil {
        structMap["ips"] = s.Ips
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
    }
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    }
    if s.RxPkts != nil {
        structMap["rx_pkts"] = s.RxPkts
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    }
    if s.TxPkts != nil {
        structMap["tx_pkts"] = s.TxPkts
    }
    if s.Up != nil {
        structMap["up"] = s.Up
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsIfStat.
// It customizes the JSON unmarshaling process for SwitchStatsIfStat objects.
func (s *SwitchStatsIfStat) UnmarshalJSON(input []byte) error {
    var temp switchStatsIfStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ips", "port_id", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Ips = temp.Ips
    s.PortId = temp.PortId
    s.RxBytes = temp.RxBytes
    s.RxPkts = temp.RxPkts
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Up = temp.Up
    return nil
}

// switchStatsIfStat is a temporary struct used for validating the fields of SwitchStatsIfStat.
type switchStatsIfStat  struct {
    Ips     []string `json:"ips,omitempty"`
    PortId  *string  `json:"port_id,omitempty"`
    RxBytes *int     `json:"rx_bytes,omitempty"`
    RxPkts  *int     `json:"rx_pkts,omitempty"`
    TxBytes *int     `json:"tx_bytes,omitempty"`
    TxPkts  *int     `json:"tx_pkts,omitempty"`
    Up      *bool    `json:"up,omitempty"`
}
