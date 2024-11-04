package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatMeshDownlink represents a ApStatMeshDownlink struct.
type ApStatMeshDownlink struct {
    Band                 *string        `json:"band,omitempty"`
    Channel              *int           `json:"channel,omitempty"`
    IdleTime             *int           `json:"idle_time,omitempty"`
    LastSeen             *float64       `json:"last_seen,omitempty"`
    Proto                *string        `json:"proto,omitempty"`
    Rssi                 *int           `json:"rssi,omitempty"`
    RxBps                *int           `json:"rx_bps,omitempty"`
    RxBytes              *int           `json:"rx_bytes,omitempty"`
    RxPackets            *int           `json:"rx_packets,omitempty"`
    RxRate               *int           `json:"rx_rate,omitempty"`
    RxRetries            *int           `json:"rx_retries,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Snr                  *int           `json:"snr,omitempty"`
    TxBps                *int           `json:"tx_bps,omitempty"`
    TxBytes              *int           `json:"tx_bytes,omitempty"`
    TxPackets            *int           `json:"tx_packets,omitempty"`
    TxRate               *int           `json:"tx_rate,omitempty"`
    TxRetries            *int           `json:"tx_retries,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatMeshDownlink.
// It customizes the JSON marshaling process for ApStatMeshDownlink objects.
func (a ApStatMeshDownlink) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatMeshDownlink object to a map representation for JSON marshaling.
func (a ApStatMeshDownlink) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Band != nil {
        structMap["band"] = a.Band
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.IdleTime != nil {
        structMap["idle_time"] = a.IdleTime
    }
    if a.LastSeen != nil {
        structMap["last_seen"] = a.LastSeen
    }
    if a.Proto != nil {
        structMap["proto"] = a.Proto
    }
    if a.Rssi != nil {
        structMap["rssi"] = a.Rssi
    }
    if a.RxBps != nil {
        structMap["rx_bps"] = a.RxBps
    }
    if a.RxBytes != nil {
        structMap["rx_bytes"] = a.RxBytes
    }
    if a.RxPackets != nil {
        structMap["rx_packets"] = a.RxPackets
    }
    if a.RxRate != nil {
        structMap["rx_rate"] = a.RxRate
    }
    if a.RxRetries != nil {
        structMap["rx_retries"] = a.RxRetries
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Snr != nil {
        structMap["snr"] = a.Snr
    }
    if a.TxBps != nil {
        structMap["tx_bps"] = a.TxBps
    }
    if a.TxBytes != nil {
        structMap["tx_bytes"] = a.TxBytes
    }
    if a.TxPackets != nil {
        structMap["tx_packets"] = a.TxPackets
    }
    if a.TxRate != nil {
        structMap["tx_rate"] = a.TxRate
    }
    if a.TxRetries != nil {
        structMap["tx_retries"] = a.TxRetries
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band", "channel", "idle_time", "last_seen", "proto", "rssi", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "site_id", "snr", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries")
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
    Band      *string    `json:"band,omitempty"`
    Channel   *int       `json:"channel,omitempty"`
    IdleTime  *int       `json:"idle_time,omitempty"`
    LastSeen  *float64   `json:"last_seen,omitempty"`
    Proto     *string    `json:"proto,omitempty"`
    Rssi      *int       `json:"rssi,omitempty"`
    RxBps     *int       `json:"rx_bps,omitempty"`
    RxBytes   *int       `json:"rx_bytes,omitempty"`
    RxPackets *int       `json:"rx_packets,omitempty"`
    RxRate    *int       `json:"rx_rate,omitempty"`
    RxRetries *int       `json:"rx_retries,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Snr       *int       `json:"snr,omitempty"`
    TxBps     *int       `json:"tx_bps,omitempty"`
    TxBytes   *int       `json:"tx_bytes,omitempty"`
    TxPackets *int       `json:"tx_packets,omitempty"`
    TxRate    *int       `json:"tx_rate,omitempty"`
    TxRetries *int       `json:"tx_retries,omitempty"`
}
