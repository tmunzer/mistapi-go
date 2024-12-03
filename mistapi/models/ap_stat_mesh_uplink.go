package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatMeshUplink represents a ApStatMeshUplink struct.
type ApStatMeshUplink struct {
    Band                 *string                `json:"band,omitempty"`
    Channel              *int                   `json:"channel,omitempty"`
    IdleTime             *int                   `json:"idle_time,omitempty"`
    LastSeen             *float64               `json:"last_seen,omitempty"`
    Proto                *string                `json:"proto,omitempty"`
    Rssi                 *int                   `json:"rssi,omitempty"`
    RxBps                *int                   `json:"rx_bps,omitempty"`
    RxBytes              *int                   `json:"rx_bytes,omitempty"`
    RxPackets            *int                   `json:"rx_packets,omitempty"`
    RxRate               *int                   `json:"rx_rate,omitempty"`
    RxRetries            *int                   `json:"rx_retries,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Snr                  *int                   `json:"snr,omitempty"`
    TxBps                *int                   `json:"tx_bps,omitempty"`
    TxBytes              *int                   `json:"tx_bytes,omitempty"`
    TxPackets            *int                   `json:"tx_packets,omitempty"`
    TxRate               *int                   `json:"tx_rate,omitempty"`
    TxRetries            *int                   `json:"tx_retries,omitempty"`
    UplinkApId           *uuid.UUID             `json:"uplink_ap_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatMeshUplink.
// It customizes the JSON marshaling process for ApStatMeshUplink objects.
func (a ApStatMeshUplink) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "band", "channel", "idle_time", "last_seen", "proto", "rssi", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "site_id", "snr", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries", "uplink_ap_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatMeshUplink object to a map representation for JSON marshaling.
func (a ApStatMeshUplink) toMap() map[string]any {
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
    if a.UplinkApId != nil {
        structMap["uplink_ap_id"] = a.UplinkApId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatMeshUplink.
// It customizes the JSON unmarshaling process for ApStatMeshUplink objects.
func (a *ApStatMeshUplink) UnmarshalJSON(input []byte) error {
    var temp tempApStatMeshUplink
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "channel", "idle_time", "last_seen", "proto", "rssi", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "site_id", "snr", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries", "uplink_ap_id")
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
    a.UplinkApId = temp.UplinkApId
    return nil
}

// tempApStatMeshUplink is a temporary struct used for validating the fields of ApStatMeshUplink.
type tempApStatMeshUplink  struct {
    Band       *string    `json:"band,omitempty"`
    Channel    *int       `json:"channel,omitempty"`
    IdleTime   *int       `json:"idle_time,omitempty"`
    LastSeen   *float64   `json:"last_seen,omitempty"`
    Proto      *string    `json:"proto,omitempty"`
    Rssi       *int       `json:"rssi,omitempty"`
    RxBps      *int       `json:"rx_bps,omitempty"`
    RxBytes    *int       `json:"rx_bytes,omitempty"`
    RxPackets  *int       `json:"rx_packets,omitempty"`
    RxRate     *int       `json:"rx_rate,omitempty"`
    RxRetries  *int       `json:"rx_retries,omitempty"`
    SiteId     *uuid.UUID `json:"site_id,omitempty"`
    Snr        *int       `json:"snr,omitempty"`
    TxBps      *int       `json:"tx_bps,omitempty"`
    TxBytes    *int       `json:"tx_bytes,omitempty"`
    TxPackets  *int       `json:"tx_packets,omitempty"`
    TxRate     *int       `json:"tx_rate,omitempty"`
    TxRetries  *int       `json:"tx_retries,omitempty"`
    UplinkApId *uuid.UUID `json:"uplink_ap_id,omitempty"`
}
