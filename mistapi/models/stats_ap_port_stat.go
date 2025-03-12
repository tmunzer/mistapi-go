package models

import (
    "encoding/json"
    "fmt"
)

// StatsApPortStat represents a StatsApPortStat struct.
type StatsApPortStat struct {
    FullDuplex           Optional[bool]         `json:"full_duplex"`
    // Amount of traffic received since connection
    RxBytes              Optional[int64]        `json:"rx_bytes"`
    RxErrors             Optional[int]          `json:"rx_errors"`
    RxPeakBps            Optional[int]          `json:"rx_peak_bps"`
    // Amount of packets received since connection
    RxPkts               Optional[int64]        `json:"rx_pkts"`
    Speed                Optional[int]          `json:"speed"`
    // Amount of traffic sent since connection
    TxBytes              Optional[int64]        `json:"tx_bytes"`
    TxPeakBps            Optional[int]          `json:"tx_peak_bps"`
    // Amount of packets sent since connection
    TxPkts               Optional[int64]        `json:"tx_pkts"`
    Up                   Optional[bool]         `json:"up"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApPortStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApPortStat) String() string {
    return fmt.Sprintf(
    	"StatsApPortStat[FullDuplex=%v, RxBytes=%v, RxErrors=%v, RxPeakBps=%v, RxPkts=%v, Speed=%v, TxBytes=%v, TxPeakBps=%v, TxPkts=%v, Up=%v, AdditionalProperties=%v]",
    	s.FullDuplex, s.RxBytes, s.RxErrors, s.RxPeakBps, s.RxPkts, s.Speed, s.TxBytes, s.TxPeakBps, s.TxPkts, s.Up, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApPortStat.
// It customizes the JSON marshaling process for StatsApPortStat objects.
func (s StatsApPortStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "full_duplex", "rx_bytes", "rx_errors", "rx_peak_bps", "rx_pkts", "speed", "tx_bytes", "tx_peak_bps", "tx_pkts", "up"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApPortStat object to a map representation for JSON marshaling.
func (s StatsApPortStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.RxPeakBps.IsValueSet() {
        if s.RxPeakBps.Value() != nil {
            structMap["rx_peak_bps"] = s.RxPeakBps.Value()
        } else {
            structMap["rx_peak_bps"] = nil
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
    if s.TxPeakBps.IsValueSet() {
        if s.TxPeakBps.Value() != nil {
            structMap["tx_peak_bps"] = s.TxPeakBps.Value()
        } else {
            structMap["tx_peak_bps"] = nil
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "full_duplex", "rx_bytes", "rx_errors", "rx_peak_bps", "rx_pkts", "speed", "tx_bytes", "tx_peak_bps", "tx_pkts", "up")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.FullDuplex = temp.FullDuplex
    s.RxBytes = temp.RxBytes
    s.RxErrors = temp.RxErrors
    s.RxPeakBps = temp.RxPeakBps
    s.RxPkts = temp.RxPkts
    s.Speed = temp.Speed
    s.TxBytes = temp.TxBytes
    s.TxPeakBps = temp.TxPeakBps
    s.TxPkts = temp.TxPkts
    s.Up = temp.Up
    return nil
}

// tempStatsApPortStat is a temporary struct used for validating the fields of StatsApPortStat.
type tempStatsApPortStat  struct {
    FullDuplex Optional[bool]  `json:"full_duplex"`
    RxBytes    Optional[int64] `json:"rx_bytes"`
    RxErrors   Optional[int]   `json:"rx_errors"`
    RxPeakBps  Optional[int]   `json:"rx_peak_bps"`
    RxPkts     Optional[int64] `json:"rx_pkts"`
    Speed      Optional[int]   `json:"speed"`
    TxBytes    Optional[int64] `json:"tx_bytes"`
    TxPeakBps  Optional[int]   `json:"tx_peak_bps"`
    TxPkts     Optional[int64] `json:"tx_pkts"`
    Up         Optional[bool]  `json:"up"`
}
