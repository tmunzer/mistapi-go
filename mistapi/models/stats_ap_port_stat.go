package models

import (
    "encoding/json"
)

// StatsApPortStat represents a StatsApPortStat struct.
type StatsApPortStat struct {
    FullDuplex           Optional[bool]    `json:"full_duplex"`
    RxBytes              Optional[float64] `json:"rx_bytes"`
    RxErrors             Optional[float64] `json:"rx_errors"`
    RxPkts               Optional[float64] `json:"rx_pkts"`
    Speed                Optional[int]     `json:"speed"`
    TxBytes              Optional[float64] `json:"tx_bytes"`
    TxPkts               Optional[float64] `json:"tx_pkts"`
    Up                   Optional[bool]    `json:"up"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApPortStat.
// It customizes the JSON marshaling process for StatsApPortStat objects.
func (s StatsApPortStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApPortStat object to a map representation for JSON marshaling.
func (s StatsApPortStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.FullDuplex.IsValueSet() {
        if s.FullDuplex.Value() != nil {
            structMap["full_duplex"] = s.FullDuplex.Value()
        } else {
            structMap["full_duplex"] = nil
        }
    }
    if s.RxBytes.IsValueSet() {
        if s.RxBytes.Value() != nil {
            structMap["rx_bytes"] = s.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if s.RxErrors.IsValueSet() {
        if s.RxErrors.Value() != nil {
            structMap["rx_errors"] = s.RxErrors.Value()
        } else {
            structMap["rx_errors"] = nil
        }
    }
    if s.RxPkts.IsValueSet() {
        if s.RxPkts.Value() != nil {
            structMap["rx_pkts"] = s.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if s.Speed.IsValueSet() {
        if s.Speed.Value() != nil {
            structMap["speed"] = s.Speed.Value()
        } else {
            structMap["speed"] = nil
        }
    }
    if s.TxBytes.IsValueSet() {
        if s.TxBytes.Value() != nil {
            structMap["tx_bytes"] = s.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if s.TxPkts.IsValueSet() {
        if s.TxPkts.Value() != nil {
            structMap["tx_pkts"] = s.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if s.Up.IsValueSet() {
        if s.Up.Value() != nil {
            structMap["up"] = s.Up.Value()
        } else {
            structMap["up"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApPortStat.
// It customizes the JSON unmarshaling process for StatsApPortStat objects.
func (s *StatsApPortStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsApPortStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "full_duplex", "rx_bytes", "rx_errors", "rx_pkts", "speed", "tx_bytes", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.FullDuplex = temp.FullDuplex
    s.RxBytes = temp.RxBytes
    s.RxErrors = temp.RxErrors
    s.RxPkts = temp.RxPkts
    s.Speed = temp.Speed
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Up = temp.Up
    return nil
}

// tempStatsApPortStat is a temporary struct used for validating the fields of StatsApPortStat.
type tempStatsApPortStat  struct {
    FullDuplex Optional[bool]    `json:"full_duplex"`
    RxBytes    Optional[float64] `json:"rx_bytes"`
    RxErrors   Optional[float64] `json:"rx_errors"`
    RxPkts     Optional[float64] `json:"rx_pkts"`
    Speed      Optional[int]     `json:"speed"`
    TxBytes    Optional[float64] `json:"tx_bytes"`
    TxPkts     Optional[float64] `json:"tx_pkts"`
    Up         Optional[bool]    `json:"up"`
}
