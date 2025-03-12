package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ApStatMeshDownlink represents a ApStatMeshDownlink struct.
type ApStatMeshDownlink struct {
    Band                 *string                `json:"band,omitempty"`
    Channel              *int                   `json:"channel,omitempty"`
    IdleTime             *int                   `json:"idle_time,omitempty"`
    // Last seen timestamp
    LastSeen             Optional[float64]      `json:"last_seen"`
    Proto                *string                `json:"proto,omitempty"`
    Rssi                 *int                   `json:"rssi,omitempty"`
    // Rate of receiving traffic, bits/seconds, last known
    RxBps                Optional[int64]        `json:"rx_bps"`
    // Amount of traffic received since connection
    RxBytes              Optional[int64]        `json:"rx_bytes"`
    // Amount of packets received since connection
    RxPackets            Optional[int64]        `json:"rx_packets"`
    // RX Rate, Mbps
    RxRate               Optional[float64]      `json:"rx_rate"`
    // Amount of rx retries
    RxRetries            Optional[int]          `json:"rx_retries"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Snr                  *int                   `json:"snr,omitempty"`
    // Rate of transmitting traffic, bits/seconds, last known
    TxBps                Optional[int64]        `json:"tx_bps"`
    // Amount of traffic sent since connection
    TxBytes              Optional[int64]        `json:"tx_bytes"`
    // Amount of packets sent since connection
    TxPackets            Optional[int64]        `json:"tx_packets"`
    // TX Rate, Mbps
    TxRate               Optional[float64]      `json:"tx_rate"`
    // Amount of tx retries
    TxRetries            Optional[int]          `json:"tx_retries"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApStatMeshDownlink,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApStatMeshDownlink) String() string {
    return fmt.Sprintf(
    	"ApStatMeshDownlink[Band=%v, Channel=%v, IdleTime=%v, LastSeen=%v, Proto=%v, Rssi=%v, RxBps=%v, RxBytes=%v, RxPackets=%v, RxRate=%v, RxRetries=%v, SiteId=%v, Snr=%v, TxBps=%v, TxBytes=%v, TxPackets=%v, TxRate=%v, TxRetries=%v, AdditionalProperties=%v]",
    	a.Band, a.Channel, a.IdleTime, a.LastSeen, a.Proto, a.Rssi, a.RxBps, a.RxBytes, a.RxPackets, a.RxRate, a.RxRetries, a.SiteId, a.Snr, a.TxBps, a.TxBytes, a.TxPackets, a.TxRate, a.TxRetries, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApStatMeshDownlink.
// It customizes the JSON marshaling process for ApStatMeshDownlink objects.
func (a ApStatMeshDownlink) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "band", "channel", "idle_time", "last_seen", "proto", "rssi", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "site_id", "snr", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatMeshDownlink object to a map representation for JSON marshaling.
func (a ApStatMeshDownlink) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Band != nil {
        structMap["band"] = a.Band
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.IdleTime != nil {
        structMap["idle_time"] = a.IdleTime
    }
    if a.LastSeen.IsValueSet() {
        if a.LastSeen.Value() != nil {
            structMap["last_seen"] = a.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if a.Proto != nil {
        structMap["proto"] = a.Proto
    }
    if a.Rssi != nil {
        structMap["rssi"] = a.Rssi
    }
    if a.RxBps.IsValueSet() {
        if a.RxBps.Value() != nil {
            structMap["rx_bps"] = a.RxBps.Value()
        } else {
            structMap["rx_bps"] = nil
        }
    }
    if a.RxBytes.IsValueSet() {
        if a.RxBytes.Value() != nil {
            structMap["rx_bytes"] = a.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if a.RxPackets.IsValueSet() {
        if a.RxPackets.Value() != nil {
            structMap["rx_packets"] = a.RxPackets.Value()
        } else {
            structMap["rx_packets"] = nil
        }
    }
    if a.RxRate.IsValueSet() {
        if a.RxRate.Value() != nil {
            structMap["rx_rate"] = a.RxRate.Value()
        } else {
            structMap["rx_rate"] = nil
        }
    }
    if a.RxRetries.IsValueSet() {
        if a.RxRetries.Value() != nil {
            structMap["rx_retries"] = a.RxRetries.Value()
        } else {
            structMap["rx_retries"] = nil
        }
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Snr != nil {
        structMap["snr"] = a.Snr
    }
    if a.TxBps.IsValueSet() {
        if a.TxBps.Value() != nil {
            structMap["tx_bps"] = a.TxBps.Value()
        } else {
            structMap["tx_bps"] = nil
        }
    }
    if a.TxBytes.IsValueSet() {
        if a.TxBytes.Value() != nil {
            structMap["tx_bytes"] = a.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if a.TxPackets.IsValueSet() {
        if a.TxPackets.Value() != nil {
            structMap["tx_packets"] = a.TxPackets.Value()
        } else {
            structMap["tx_packets"] = nil
        }
    }
    if a.TxRate.IsValueSet() {
        if a.TxRate.Value() != nil {
            structMap["tx_rate"] = a.TxRate.Value()
        } else {
            structMap["tx_rate"] = nil
        }
    }
    if a.TxRetries.IsValueSet() {
        if a.TxRetries.Value() != nil {
            structMap["tx_retries"] = a.TxRetries.Value()
        } else {
            structMap["tx_retries"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatMeshDownlink.
// It customizes the JSON unmarshaling process for ApStatMeshDownlink objects.
func (a *ApStatMeshDownlink) UnmarshalJSON(input []byte) error {
    var temp tempApStatMeshDownlink
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channel", "idle_time", "last_seen", "proto", "rssi", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "site_id", "snr", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Band = temp.Band
    a.Channel = temp.Channel
    a.IdleTime = temp.IdleTime
    a.LastSeen = temp.LastSeen
    a.Proto = temp.Proto
    a.Rssi = temp.Rssi
    a.RxBps = temp.RxBps
    a.RxBytes = temp.RxBytes
    a.RxPackets = temp.RxPackets
    a.RxRate = temp.RxRate
    a.RxRetries = temp.RxRetries
    a.SiteId = temp.SiteId
    a.Snr = temp.Snr
    a.TxBps = temp.TxBps
    a.TxBytes = temp.TxBytes
    a.TxPackets = temp.TxPackets
    a.TxRate = temp.TxRate
    a.TxRetries = temp.TxRetries
    return nil
}

// tempApStatMeshDownlink is a temporary struct used for validating the fields of ApStatMeshDownlink.
type tempApStatMeshDownlink  struct {
    Band      *string           `json:"band,omitempty"`
    Channel   *int              `json:"channel,omitempty"`
    IdleTime  *int              `json:"idle_time,omitempty"`
    LastSeen  Optional[float64] `json:"last_seen"`
    Proto     *string           `json:"proto,omitempty"`
    Rssi      *int              `json:"rssi,omitempty"`
    RxBps     Optional[int64]   `json:"rx_bps"`
    RxBytes   Optional[int64]   `json:"rx_bytes"`
    RxPackets Optional[int64]   `json:"rx_packets"`
    RxRate    Optional[float64] `json:"rx_rate"`
    RxRetries Optional[int]     `json:"rx_retries"`
    SiteId    *uuid.UUID        `json:"site_id,omitempty"`
    Snr       *int              `json:"snr,omitempty"`
    TxBps     Optional[int64]   `json:"tx_bps"`
    TxBytes   Optional[int64]   `json:"tx_bytes"`
    TxPackets Optional[int64]   `json:"tx_packets"`
    TxRate    Optional[float64] `json:"tx_rate"`
    TxRetries Optional[int]     `json:"tx_retries"`
}
