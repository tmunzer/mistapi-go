package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatsBle represents a ApStatsBle struct.
type ApStatsBle struct {
    BeaconEnabled         Optional[bool]      `json:"beacon_enabled"`
    BeaconRate            Optional[int]       `json:"beacon_rate"`
    EddystoneUidEnabled   Optional[bool]      `json:"eddystone_uid_enabled"`
    EddystoneUidFreqMsec  Optional[int]       `json:"eddystone_uid_freq_msec"`
    EddystoneUidInstance  Optional[string]    `json:"eddystone_uid_instance"`
    EddystoneUidNamespace Optional[string]    `json:"eddystone_uid_namespace"`
    EddystoneUrlEnabled   Optional[bool]      `json:"eddystone_url_enabled"`
    // Frequency (msec) of data emmit by Eddystone-UID beacon
    EddystoneUrlFreqMsec  Optional[int]       `json:"eddystone_url_freq_msec"`
    EddystoneUrlUrl       Optional[string]    `json:"eddystone_url_url"`
    IbeaconEnabled        Optional[bool]      `json:"ibeacon_enabled"`
    IbeaconFreqMsec       Optional[int]       `json:"ibeacon_freq_msec"`
    IbeaconMajor          Optional[int]       `json:"ibeacon_major"`
    IbeaconMinor          Optional[int]       `json:"ibeacon_minor"`
    IbeaconUuid           Optional[uuid.UUID] `json:"ibeacon_uuid"`
    Major                 Optional[int]       `json:"major"`
    Minors                []int               `json:"minors,omitempty"`
    Power                 Optional[int]       `json:"power"`
    RxBytes               Optional[int]       `json:"rx_bytes"`
    RxPkts                Optional[int]       `json:"rx_pkts"`
    TxBytes               Optional[int64]     `json:"tx_bytes"`
    TxPkts                Optional[int]       `json:"tx_pkts"`
    // resets due to tx hung
    TxResets              Optional[int]       `json:"tx_resets"`
    Uuid                  Optional[uuid.UUID] `json:"uuid"`
    AdditionalProperties  map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsBle.
// It customizes the JSON marshaling process for ApStatsBle objects.
func (a ApStatsBle) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsBle object to a map representation for JSON marshaling.
func (a ApStatsBle) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.BeaconEnabled.IsValueSet() {
        if a.BeaconEnabled.Value() != nil {
            structMap["beacon_enabled"] = a.BeaconEnabled.Value()
        } else {
            structMap["beacon_enabled"] = nil
        }
    }
    if a.BeaconRate.IsValueSet() {
        if a.BeaconRate.Value() != nil {
            structMap["beacon_rate"] = a.BeaconRate.Value()
        } else {
            structMap["beacon_rate"] = nil
        }
    }
    if a.EddystoneUidEnabled.IsValueSet() {
        if a.EddystoneUidEnabled.Value() != nil {
            structMap["eddystone_uid_enabled"] = a.EddystoneUidEnabled.Value()
        } else {
            structMap["eddystone_uid_enabled"] = nil
        }
    }
    if a.EddystoneUidFreqMsec.IsValueSet() {
        if a.EddystoneUidFreqMsec.Value() != nil {
            structMap["eddystone_uid_freq_msec"] = a.EddystoneUidFreqMsec.Value()
        } else {
            structMap["eddystone_uid_freq_msec"] = nil
        }
    }
    if a.EddystoneUidInstance.IsValueSet() {
        if a.EddystoneUidInstance.Value() != nil {
            structMap["eddystone_uid_instance"] = a.EddystoneUidInstance.Value()
        } else {
            structMap["eddystone_uid_instance"] = nil
        }
    }
    if a.EddystoneUidNamespace.IsValueSet() {
        if a.EddystoneUidNamespace.Value() != nil {
            structMap["eddystone_uid_namespace"] = a.EddystoneUidNamespace.Value()
        } else {
            structMap["eddystone_uid_namespace"] = nil
        }
    }
    if a.EddystoneUrlEnabled.IsValueSet() {
        if a.EddystoneUrlEnabled.Value() != nil {
            structMap["eddystone_url_enabled"] = a.EddystoneUrlEnabled.Value()
        } else {
            structMap["eddystone_url_enabled"] = nil
        }
    }
    if a.EddystoneUrlFreqMsec.IsValueSet() {
        if a.EddystoneUrlFreqMsec.Value() != nil {
            structMap["eddystone_url_freq_msec"] = a.EddystoneUrlFreqMsec.Value()
        } else {
            structMap["eddystone_url_freq_msec"] = nil
        }
    }
    if a.EddystoneUrlUrl.IsValueSet() {
        if a.EddystoneUrlUrl.Value() != nil {
            structMap["eddystone_url_url"] = a.EddystoneUrlUrl.Value()
        } else {
            structMap["eddystone_url_url"] = nil
        }
    }
    if a.IbeaconEnabled.IsValueSet() {
        if a.IbeaconEnabled.Value() != nil {
            structMap["ibeacon_enabled"] = a.IbeaconEnabled.Value()
        } else {
            structMap["ibeacon_enabled"] = nil
        }
    }
    if a.IbeaconFreqMsec.IsValueSet() {
        if a.IbeaconFreqMsec.Value() != nil {
            structMap["ibeacon_freq_msec"] = a.IbeaconFreqMsec.Value()
        } else {
            structMap["ibeacon_freq_msec"] = nil
        }
    }
    if a.IbeaconMajor.IsValueSet() {
        if a.IbeaconMajor.Value() != nil {
            structMap["ibeacon_major"] = a.IbeaconMajor.Value()
        } else {
            structMap["ibeacon_major"] = nil
        }
    }
    if a.IbeaconMinor.IsValueSet() {
        if a.IbeaconMinor.Value() != nil {
            structMap["ibeacon_minor"] = a.IbeaconMinor.Value()
        } else {
            structMap["ibeacon_minor"] = nil
        }
    }
    if a.IbeaconUuid.IsValueSet() {
        if a.IbeaconUuid.Value() != nil {
            structMap["ibeacon_uuid"] = a.IbeaconUuid.Value()
        } else {
            structMap["ibeacon_uuid"] = nil
        }
    }
    if a.Major.IsValueSet() {
        if a.Major.Value() != nil {
            structMap["major"] = a.Major.Value()
        } else {
            structMap["major"] = nil
        }
    }
    if a.Minors != nil {
        structMap["minors"] = a.Minors
    }
    if a.Power.IsValueSet() {
        if a.Power.Value() != nil {
            structMap["power"] = a.Power.Value()
        } else {
            structMap["power"] = nil
        }
    }
    if a.RxBytes.IsValueSet() {
        if a.RxBytes.Value() != nil {
            structMap["rx_bytes"] = a.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if a.RxPkts.IsValueSet() {
        if a.RxPkts.Value() != nil {
            structMap["rx_pkts"] = a.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
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
    if a.TxResets.IsValueSet() {
        if a.TxResets.Value() != nil {
            structMap["tx_resets"] = a.TxResets.Value()
        } else {
            structMap["tx_resets"] = nil
        }
    }
    if a.Uuid.IsValueSet() {
        if a.Uuid.Value() != nil {
            structMap["uuid"] = a.Uuid.Value()
        } else {
            structMap["uuid"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsBle.
// It customizes the JSON unmarshaling process for ApStatsBle objects.
func (a *ApStatsBle) UnmarshalJSON(input []byte) error {
    var temp tempApStatsBle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "beacon_enabled", "beacon_rate", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_enabled", "ibeacon_freq_msec", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "major", "minors", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "tx_resets", "uuid")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.BeaconEnabled = temp.BeaconEnabled
    a.BeaconRate = temp.BeaconRate
    a.EddystoneUidEnabled = temp.EddystoneUidEnabled
    a.EddystoneUidFreqMsec = temp.EddystoneUidFreqMsec
    a.EddystoneUidInstance = temp.EddystoneUidInstance
    a.EddystoneUidNamespace = temp.EddystoneUidNamespace
    a.EddystoneUrlEnabled = temp.EddystoneUrlEnabled
    a.EddystoneUrlFreqMsec = temp.EddystoneUrlFreqMsec
    a.EddystoneUrlUrl = temp.EddystoneUrlUrl
    a.IbeaconEnabled = temp.IbeaconEnabled
    a.IbeaconFreqMsec = temp.IbeaconFreqMsec
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

// tempApStatsBle is a temporary struct used for validating the fields of ApStatsBle.
type tempApStatsBle  struct {
    BeaconEnabled         Optional[bool]      `json:"beacon_enabled"`
    BeaconRate            Optional[int]       `json:"beacon_rate"`
    EddystoneUidEnabled   Optional[bool]      `json:"eddystone_uid_enabled"`
    EddystoneUidFreqMsec  Optional[int]       `json:"eddystone_uid_freq_msec"`
    EddystoneUidInstance  Optional[string]    `json:"eddystone_uid_instance"`
    EddystoneUidNamespace Optional[string]    `json:"eddystone_uid_namespace"`
    EddystoneUrlEnabled   Optional[bool]      `json:"eddystone_url_enabled"`
    EddystoneUrlFreqMsec  Optional[int]       `json:"eddystone_url_freq_msec"`
    EddystoneUrlUrl       Optional[string]    `json:"eddystone_url_url"`
    IbeaconEnabled        Optional[bool]      `json:"ibeacon_enabled"`
    IbeaconFreqMsec       Optional[int]       `json:"ibeacon_freq_msec"`
    IbeaconMajor          Optional[int]       `json:"ibeacon_major"`
    IbeaconMinor          Optional[int]       `json:"ibeacon_minor"`
    IbeaconUuid           Optional[uuid.UUID] `json:"ibeacon_uuid"`
    Major                 Optional[int]       `json:"major"`
    Minors                []int               `json:"minors,omitempty"`
    Power                 Optional[int]       `json:"power"`
    RxBytes               Optional[int]       `json:"rx_bytes"`
    RxPkts                Optional[int]       `json:"rx_pkts"`
    TxBytes               Optional[int64]     `json:"tx_bytes"`
    TxPkts                Optional[int]       `json:"tx_pkts"`
    TxResets              Optional[int]       `json:"tx_resets"`
    Uuid                  Optional[uuid.UUID] `json:"uuid"`
}
