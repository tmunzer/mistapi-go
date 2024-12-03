package models

import (
    "encoding/json"
)

// StatsMxedgePortStat represents a StatsMxedgePortStat struct.
type StatsMxedgePortStat struct {
    FullDuplex           *bool                  `json:"full_duplex,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    RxBytes              *float64               `json:"rx_bytes,omitempty"`
    RxErrors             *int                   `json:"rx_errors,omitempty"`
    RxPkts               *int                   `json:"rx_pkts,omitempty"`
    Speed                *int                   `json:"speed,omitempty"`
    State                *string                `json:"state,omitempty"`
    TxBytes              *int                   `json:"tx_bytes,omitempty"`
    TxErrors             *int                   `json:"tx_errors,omitempty"`
    TxPkts               *int                   `json:"tx_pkts,omitempty"`
    Up                   *bool                  `json:"up,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgePortStat.
// It customizes the JSON marshaling process for StatsMxedgePortStat objects.
func (s StatsMxedgePortStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "full_duplex", "mac", "rx_bytes", "rx_errors", "rx_pkts", "speed", "state", "tx_bytes", "tx_errors", "tx_pkts", "up"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgePortStat object to a map representation for JSON marshaling.
func (s StatsMxedgePortStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.FullDuplex != nil {
        structMap["full_duplex"] = s.FullDuplex
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    }
    if s.RxErrors != nil {
        structMap["rx_errors"] = s.RxErrors
    }
    if s.RxPkts != nil {
        structMap["rx_pkts"] = s.RxPkts
    }
    if s.Speed != nil {
        structMap["speed"] = s.Speed
    }
    if s.State != nil {
        structMap["state"] = s.State
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    }
    if s.TxErrors != nil {
        structMap["tx_errors"] = s.TxErrors
    }
    if s.TxPkts != nil {
        structMap["tx_pkts"] = s.TxPkts
    }
    if s.Up != nil {
        structMap["up"] = s.Up
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgePortStat.
// It customizes the JSON unmarshaling process for StatsMxedgePortStat objects.
func (s *StatsMxedgePortStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgePortStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "full_duplex", "mac", "rx_bytes", "rx_errors", "rx_pkts", "speed", "state", "tx_bytes", "tx_errors", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.FullDuplex = temp.FullDuplex
    s.Mac = temp.Mac
    s.RxBytes = temp.RxBytes
    s.RxErrors = temp.RxErrors
    s.RxPkts = temp.RxPkts
    s.Speed = temp.Speed
    s.State = temp.State
    s.TxBytes = temp.TxBytes
    s.TxErrors = temp.TxErrors
    s.TxPkts = temp.TxPkts
    s.Up = temp.Up
    return nil
}

// tempStatsMxedgePortStat is a temporary struct used for validating the fields of StatsMxedgePortStat.
type tempStatsMxedgePortStat  struct {
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
