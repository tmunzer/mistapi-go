package models

import (
    "encoding/json"
)

// ApStatsPortStat represents a ApStatsPortStat struct.
type ApStatsPortStat struct {
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
    if a.FullDuplex.IsValueSet() {
        if a.FullDuplex.Value() != nil {
            structMap["full_duplex"] = a.FullDuplex.Value()
        } else {
            structMap["full_duplex"] = nil
        }
    }
    if a.RxBytes.IsValueSet() {
        if a.RxBytes.Value() != nil {
            structMap["rx_bytes"] = a.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if a.RxErrors.IsValueSet() {
        if a.RxErrors.Value() != nil {
            structMap["rx_errors"] = a.RxErrors.Value()
        } else {
            structMap["rx_errors"] = nil
        }
    }
    if a.RxPkts.IsValueSet() {
        if a.RxPkts.Value() != nil {
            structMap["rx_pkts"] = a.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if a.Speed.IsValueSet() {
        if a.Speed.Value() != nil {
            structMap["speed"] = a.Speed.Value()
        } else {
            structMap["speed"] = nil
        }
    }
    if a.TxBytes.IsValueSet() {
        if a.TxBytes.Value() != nil {
            structMap["tx_bytes"] = a.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if a.TxPkts.IsValueSet() {
        if a.TxPkts.Value() != nil {
            structMap["tx_pkts"] = a.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if a.Up.IsValueSet() {
        if a.Up.Value() != nil {
            structMap["up"] = a.Up.Value()
        } else {
            structMap["up"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsPortStat.
// It customizes the JSON unmarshaling process for ApStatsPortStat objects.
func (a *ApStatsPortStat) UnmarshalJSON(input []byte) error {
    var temp tempApStatsPortStat
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

// tempApStatsPortStat is a temporary struct used for validating the fields of ApStatsPortStat.
type tempApStatsPortStat  struct {
    FullDuplex Optional[bool]    `json:"full_duplex"`
    RxBytes    Optional[float64] `json:"rx_bytes"`
    RxErrors   Optional[float64] `json:"rx_errors"`
    RxPkts     Optional[float64] `json:"rx_pkts"`
    Speed      Optional[int]     `json:"speed"`
    TxBytes    Optional[float64] `json:"tx_bytes"`
    TxPkts     Optional[float64] `json:"tx_pkts"`
    Up         Optional[bool]    `json:"up"`
}
