package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatsBleStat represents a ApStatsBleStat struct.
type ApStatsBleStat struct {
    BeaconRate            *int           `json:"beacon_rate,omitempty"`
    EddystoneUidEnabled   *bool          `json:"eddystone_uid_enabled,omitempty"`
    EddystoneUidFreqMsec  *int           `json:"eddystone_uid_freq_msec,omitempty"`
    EddystoneUidInstance  *string        `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace *string        `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlEnabled   *bool          `json:"eddystone_url_enabled,omitempty"`
    // Frequency (msec) of data emmit by Eddystone-UID beacon
    EddystoneUrlFreqMsec  *int           `json:"eddystone_url_freq_msec,omitempty"`
    EddystoneUrlUrl       *string        `json:"eddystone_url_url,omitempty"`
    IbeaconEnabled        *bool          `json:"ibeacon_enabled,omitempty"`
    IbeaconMajor          *int           `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int           `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID     `json:"ibeacon_uuid,omitempty"`
    Major                 *int           `json:"major,omitempty"`
    Minors                []int          `json:"minors,omitempty"`
    Power                 *int           `json:"power,omitempty"`
    RxBytes               *int           `json:"rx_bytes,omitempty"`
    RxPkts                *int           `json:"rx_pkts,omitempty"`
    TxBytes               *int64         `json:"tx_bytes,omitempty"`
    TxPkts                *int           `json:"tx_pkts,omitempty"`
    // resets due to tx hung
    TxResets              *int           `json:"tx_resets,omitempty"`
    Uuid                  *uuid.UUID     `json:"uuid,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsBleStat.
// It customizes the JSON marshaling process for ApStatsBleStat objects.
func (a ApStatsBleStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsBleStat object to a map representation for JSON marshaling.
func (a ApStatsBleStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.BeaconRate != nil {
        structMap["beacon_rate"] = a.BeaconRate
    }
    if a.EddystoneUidEnabled != nil {
        structMap["eddystone_uid_enabled"] = a.EddystoneUidEnabled
    }
    if a.EddystoneUidFreqMsec != nil {
        structMap["eddystone_uid_freq_msec"] = a.EddystoneUidFreqMsec
    }
    if a.EddystoneUidInstance != nil {
        structMap["eddystone_uid_instance"] = a.EddystoneUidInstance
    }
    if a.EddystoneUidNamespace != nil {
        structMap["eddystone_uid_namespace"] = a.EddystoneUidNamespace
    }
    if a.EddystoneUrlEnabled != nil {
        structMap["eddystone_url_enabled"] = a.EddystoneUrlEnabled
    }
    if a.EddystoneUrlFreqMsec != nil {
        structMap["eddystone_url_freq_msec"] = a.EddystoneUrlFreqMsec
    }
    if a.EddystoneUrlUrl != nil {
        structMap["eddystone_url_url"] = a.EddystoneUrlUrl
    }
    if a.IbeaconEnabled != nil {
        structMap["ibeacon_enabled"] = a.IbeaconEnabled
    }
    if a.IbeaconMajor != nil {
        structMap["ibeacon_major"] = a.IbeaconMajor
    }
    if a.IbeaconMinor != nil {
        structMap["ibeacon_minor"] = a.IbeaconMinor
    }
    if a.IbeaconUuid != nil {
        structMap["ibeacon_uuid"] = a.IbeaconUuid
    }
    if a.Major != nil {
        structMap["major"] = a.Major
    }
    if a.Minors != nil {
        structMap["minors"] = a.Minors
    }
    if a.Power != nil {
        structMap["power"] = a.Power
    }
    if a.RxBytes != nil {
        structMap["rx_bytes"] = a.RxBytes
    }
    if a.RxPkts != nil {
        structMap["rx_pkts"] = a.RxPkts
    }
    if a.TxBytes != nil {
        structMap["tx_bytes"] = a.TxBytes
    }
    if a.TxPkts != nil {
        structMap["tx_pkts"] = a.TxPkts
    }
    if a.TxResets != nil {
        structMap["tx_resets"] = a.TxResets
    }
    if a.Uuid != nil {
        structMap["uuid"] = a.Uuid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsBleStat.
// It customizes the JSON unmarshaling process for ApStatsBleStat objects.
func (a *ApStatsBleStat) UnmarshalJSON(input []byte) error {
    var temp apStatsBleStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "beacon_rate", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_enabled", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "major", "minors", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "tx_resets", "uuid")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.BeaconRate = temp.BeaconRate
    a.EddystoneUidEnabled = temp.EddystoneUidEnabled
    a.EddystoneUidFreqMsec = temp.EddystoneUidFreqMsec
    a.EddystoneUidInstance = temp.EddystoneUidInstance
    a.EddystoneUidNamespace = temp.EddystoneUidNamespace
    a.EddystoneUrlEnabled = temp.EddystoneUrlEnabled
    a.EddystoneUrlFreqMsec = temp.EddystoneUrlFreqMsec
    a.EddystoneUrlUrl = temp.EddystoneUrlUrl
    a.IbeaconEnabled = temp.IbeaconEnabled
    a.IbeaconMajor = temp.IbeaconMajor
    a.IbeaconMinor = temp.IbeaconMinor
    a.IbeaconUuid = temp.IbeaconUuid
    a.Major = temp.Major
    a.Minors = temp.Minors
    a.Power = temp.Power
    a.RxBytes = temp.RxBytes
    a.RxPkts = temp.RxPkts
    a.TxBytes = temp.TxBytes
    a.TxPkts = temp.TxPkts
    a.TxResets = temp.TxResets
    a.Uuid = temp.Uuid
    return nil
}

// apStatsBleStat is a temporary struct used for validating the fields of ApStatsBleStat.
type apStatsBleStat  struct {
    BeaconRate            *int       `json:"beacon_rate,omitempty"`
    EddystoneUidEnabled   *bool      `json:"eddystone_uid_enabled,omitempty"`
    EddystoneUidFreqMsec  *int       `json:"eddystone_uid_freq_msec,omitempty"`
    EddystoneUidInstance  *string    `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace *string    `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlEnabled   *bool      `json:"eddystone_url_enabled,omitempty"`
    EddystoneUrlFreqMsec  *int       `json:"eddystone_url_freq_msec,omitempty"`
    EddystoneUrlUrl       *string    `json:"eddystone_url_url,omitempty"`
    IbeaconEnabled        *bool      `json:"ibeacon_enabled,omitempty"`
    IbeaconMajor          *int       `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int       `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID `json:"ibeacon_uuid,omitempty"`
    Major                 *int       `json:"major,omitempty"`
    Minors                []int      `json:"minors,omitempty"`
    Power                 *int       `json:"power,omitempty"`
    RxBytes               *int       `json:"rx_bytes,omitempty"`
    RxPkts                *int       `json:"rx_pkts,omitempty"`
    TxBytes               *int64     `json:"tx_bytes,omitempty"`
    TxPkts                *int       `json:"tx_pkts,omitempty"`
    TxResets              *int       `json:"tx_resets,omitempty"`
    Uuid                  *uuid.UUID `json:"uuid,omitempty"`
}
