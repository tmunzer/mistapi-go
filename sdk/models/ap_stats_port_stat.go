package models

import (
    "encoding/json"
)

// ApStatsPortStat represents a ApStatsPortStat struct.
type ApStatsPortStat struct {
    FullDuplex           *bool          `json:"full_duplex,omitempty"`
    RxBytes              *float64       `json:"rx_bytes,omitempty"`
    RxErrors             *float64       `json:"rx_errors,omitempty"`
    RxPkts               *float64       `json:"rx_pkts,omitempty"`
    Speed                *int           `json:"speed,omitempty"`
    TxBytes              *float64       `json:"tx_bytes,omitempty"`
    TxPkts               *float64       `json:"tx_pkts,omitempty"`
    Up                   *bool          `json:"up,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsPortStat.
// It customizes the JSON marshaling process for ApStatsPortStat objects.
func (a ApStatsPortStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsPortStat object to a map representation for JSON marshaling.
func (a ApStatsPortStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.FullDuplex != nil {
        structMap["full_duplex"] = a.FullDuplex
    }
    if a.RxBytes != nil {
        structMap["rx_bytes"] = a.RxBytes
    }
    if a.RxErrors != nil {
        structMap["rx_errors"] = a.RxErrors
    }
    if a.RxPkts != nil {
        structMap["rx_pkts"] = a.RxPkts
    }
    if a.Speed != nil {
        structMap["speed"] = a.Speed
    }
    if a.TxBytes != nil {
        structMap["tx_bytes"] = a.TxBytes
    }
    if a.TxPkts != nil {
        structMap["tx_pkts"] = a.TxPkts
    }
    if a.Up != nil {
        structMap["up"] = a.Up
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsPortStat.
// It customizes the JSON unmarshaling process for ApStatsPortStat objects.
func (a *ApStatsPortStat) UnmarshalJSON(input []byte) error {
    var temp apStatsPortStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "full_duplex", "rx_bytes", "rx_errors", "rx_pkts", "speed", "tx_bytes", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.FullDuplex = temp.FullDuplex
    a.RxBytes = temp.RxBytes
    a.RxErrors = temp.RxErrors
    a.RxPkts = temp.RxPkts
    a.Speed = temp.Speed
    a.TxBytes = temp.TxBytes
    a.TxPkts = temp.TxPkts
    a.Up = temp.Up
    return nil
}

// apStatsPortStat is a temporary struct used for validating the fields of ApStatsPortStat.
type apStatsPortStat  struct {
    FullDuplex *bool    `json:"full_duplex,omitempty"`
    RxBytes    *float64 `json:"rx_bytes,omitempty"`
    RxErrors   *float64 `json:"rx_errors,omitempty"`
    RxPkts     *float64 `json:"rx_pkts,omitempty"`
    Speed      *int     `json:"speed,omitempty"`
    TxBytes    *float64 `json:"tx_bytes,omitempty"`
    TxPkts     *float64 `json:"tx_pkts,omitempty"`
    Up         *bool    `json:"up,omitempty"`
}
