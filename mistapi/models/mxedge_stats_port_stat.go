package models

import (
    "encoding/json"
)

// MxedgeStatsPortStat represents a MxedgeStatsPortStat struct.
type MxedgeStatsPortStat struct {
    FullDuplex           *bool          `json:"full_duplex,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    RxBytes              *float64       `json:"rx_bytes,omitempty"`
    RxErrors             *int           `json:"rx_errors,omitempty"`
    RxPkts               *int           `json:"rx_pkts,omitempty"`
    Speed                *int           `json:"speed,omitempty"`
    State                *string        `json:"state,omitempty"`
    TxBytes              *int           `json:"tx_bytes,omitempty"`
    TxErrors             *int           `json:"tx_errors,omitempty"`
    TxPkts               *int           `json:"tx_pkts,omitempty"`
    Up                   *bool          `json:"up,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsPortStat.
// It customizes the JSON marshaling process for MxedgeStatsPortStat objects.
func (m MxedgeStatsPortStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsPortStat object to a map representation for JSON marshaling.
func (m MxedgeStatsPortStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.FullDuplex != nil {
        structMap["full_duplex"] = m.FullDuplex
    }
    if m.Mac != nil {
        structMap["mac"] = m.Mac
    }
    if m.RxBytes != nil {
        structMap["rx_bytes"] = m.RxBytes
    }
    if m.RxErrors != nil {
        structMap["rx_errors"] = m.RxErrors
    }
    if m.RxPkts != nil {
        structMap["rx_pkts"] = m.RxPkts
    }
    if m.Speed != nil {
        structMap["speed"] = m.Speed
    }
    if m.State != nil {
        structMap["state"] = m.State
    }
    if m.TxBytes != nil {
        structMap["tx_bytes"] = m.TxBytes
    }
    if m.TxErrors != nil {
        structMap["tx_errors"] = m.TxErrors
    }
    if m.TxPkts != nil {
        structMap["tx_pkts"] = m.TxPkts
    }
    if m.Up != nil {
        structMap["up"] = m.Up
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsPortStat.
// It customizes the JSON unmarshaling process for MxedgeStatsPortStat objects.
func (m *MxedgeStatsPortStat) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeStatsPortStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "full_duplex", "mac", "rx_bytes", "rx_errors", "rx_pkts", "speed", "state", "tx_bytes", "tx_errors", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.FullDuplex = temp.FullDuplex
    m.Mac = temp.Mac
    m.RxBytes = temp.RxBytes
    m.RxErrors = temp.RxErrors
    m.RxPkts = temp.RxPkts
    m.Speed = temp.Speed
    m.State = temp.State
    m.TxBytes = temp.TxBytes
    m.TxErrors = temp.TxErrors
    m.TxPkts = temp.TxPkts
    m.Up = temp.Up
    return nil
}

// tempMxedgeStatsPortStat is a temporary struct used for validating the fields of MxedgeStatsPortStat.
type tempMxedgeStatsPortStat  struct {
    FullDuplex *bool    `json:"full_duplex,omitempty"`
    Mac        *string  `json:"mac,omitempty"`
    RxBytes    *float64 `json:"rx_bytes,omitempty"`
    RxErrors   *int     `json:"rx_errors,omitempty"`
    RxPkts     *int     `json:"rx_pkts,omitempty"`
    Speed      *int     `json:"speed,omitempty"`
    State      *string  `json:"state,omitempty"`
    TxBytes    *int     `json:"tx_bytes,omitempty"`
    TxErrors   *int     `json:"tx_errors,omitempty"`
    TxPkts     *int     `json:"tx_pkts,omitempty"`
    Up         *bool    `json:"up,omitempty"`
}
