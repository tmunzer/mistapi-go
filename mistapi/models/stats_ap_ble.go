// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// StatsApBle represents a StatsApBle struct.
type StatsApBle struct {
	BeaconEnabled         Optional[bool]   `json:"beacon_enabled"`
	BeaconRate            Optional[int]    `json:"beacon_rate"`
	EddystoneUidEnabled   Optional[bool]   `json:"eddystone_uid_enabled"`
	EddystoneUidFreqMsec  Optional[int]    `json:"eddystone_uid_freq_msec"`
	EddystoneUidInstance  Optional[string] `json:"eddystone_uid_instance"`
	EddystoneUidNamespace Optional[string] `json:"eddystone_uid_namespace"`
	EddystoneUrlEnabled   Optional[bool]   `json:"eddystone_url_enabled"`
	// Frequency (msec) of data emit by Eddystone-UID beacon
	EddystoneUrlFreqMsec Optional[int]       `json:"eddystone_url_freq_msec"`
	EddystoneUrlUrl      Optional[string]    `json:"eddystone_url_url"`
	IbeaconEnabled       Optional[bool]      `json:"ibeacon_enabled"`
	IbeaconFreqMsec      Optional[int]       `json:"ibeacon_freq_msec"`
	IbeaconMajor         Optional[int]       `json:"ibeacon_major"`
	IbeaconMinor         Optional[int]       `json:"ibeacon_minor"`
	IbeaconUuid          Optional[uuid.UUID] `json:"ibeacon_uuid"`
	Major                Optional[int]       `json:"major"`
	Minors               []int               `json:"minors,omitempty"`
	Power                Optional[int]       `json:"power"`
	// Amount of traffic received since connection
	RxBytes Optional[int64] `json:"rx_bytes"`
	// Amount of packets received since connection
	RxPkts Optional[int64] `json:"rx_pkts"`
	// Amount of traffic sent since connection
	TxBytes Optional[int64] `json:"tx_bytes"`
	// Amount of packets sent since connection
	TxPkts Optional[int64] `json:"tx_pkts"`
	// Resets due to tx hung
	TxResets             Optional[int]          `json:"tx_resets"`
	Uuid                 Optional[uuid.UUID]    `json:"uuid"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApBle,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApBle) String() string {
	return fmt.Sprintf(
		"StatsApBle[BeaconEnabled=%v, BeaconRate=%v, EddystoneUidEnabled=%v, EddystoneUidFreqMsec=%v, EddystoneUidInstance=%v, EddystoneUidNamespace=%v, EddystoneUrlEnabled=%v, EddystoneUrlFreqMsec=%v, EddystoneUrlUrl=%v, IbeaconEnabled=%v, IbeaconFreqMsec=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, Major=%v, Minors=%v, Power=%v, RxBytes=%v, RxPkts=%v, TxBytes=%v, TxPkts=%v, TxResets=%v, Uuid=%v, AdditionalProperties=%v]",
		s.BeaconEnabled, s.BeaconRate, s.EddystoneUidEnabled, s.EddystoneUidFreqMsec, s.EddystoneUidInstance, s.EddystoneUidNamespace, s.EddystoneUrlEnabled, s.EddystoneUrlFreqMsec, s.EddystoneUrlUrl, s.IbeaconEnabled, s.IbeaconFreqMsec, s.IbeaconMajor, s.IbeaconMinor, s.IbeaconUuid, s.Major, s.Minors, s.Power, s.RxBytes, s.RxPkts, s.TxBytes, s.TxPkts, s.TxResets, s.Uuid, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApBle.
// It customizes the JSON marshaling process for StatsApBle objects.
func (s StatsApBle) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"beacon_enabled", "beacon_rate", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_enabled", "ibeacon_freq_msec", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "major", "minors", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "tx_resets", "uuid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApBle object to a map representation for JSON marshaling.
func (s StatsApBle) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.BeaconEnabled.IsValueSet() {
		if s.BeaconEnabled.Value() != nil {
			structMap["beacon_enabled"] = s.BeaconEnabled.Value()
		} else {
			structMap["beacon_enabled"] = nil
		}
	}
	if s.BeaconRate.IsValueSet() {
		if s.BeaconRate.Value() != nil {
			structMap["beacon_rate"] = s.BeaconRate.Value()
		} else {
			structMap["beacon_rate"] = nil
		}
	}
	if s.EddystoneUidEnabled.IsValueSet() {
		if s.EddystoneUidEnabled.Value() != nil {
			structMap["eddystone_uid_enabled"] = s.EddystoneUidEnabled.Value()
		} else {
			structMap["eddystone_uid_enabled"] = nil
		}
	}
	if s.EddystoneUidFreqMsec.IsValueSet() {
		if s.EddystoneUidFreqMsec.Value() != nil {
			structMap["eddystone_uid_freq_msec"] = s.EddystoneUidFreqMsec.Value()
		} else {
			structMap["eddystone_uid_freq_msec"] = nil
		}
	}
	if s.EddystoneUidInstance.IsValueSet() {
		if s.EddystoneUidInstance.Value() != nil {
			structMap["eddystone_uid_instance"] = s.EddystoneUidInstance.Value()
		} else {
			structMap["eddystone_uid_instance"] = nil
		}
	}
	if s.EddystoneUidNamespace.IsValueSet() {
		if s.EddystoneUidNamespace.Value() != nil {
			structMap["eddystone_uid_namespace"] = s.EddystoneUidNamespace.Value()
		} else {
			structMap["eddystone_uid_namespace"] = nil
		}
	}
	if s.EddystoneUrlEnabled.IsValueSet() {
		if s.EddystoneUrlEnabled.Value() != nil {
			structMap["eddystone_url_enabled"] = s.EddystoneUrlEnabled.Value()
		} else {
			structMap["eddystone_url_enabled"] = nil
		}
	}
	if s.EddystoneUrlFreqMsec.IsValueSet() {
		if s.EddystoneUrlFreqMsec.Value() != nil {
			structMap["eddystone_url_freq_msec"] = s.EddystoneUrlFreqMsec.Value()
		} else {
			structMap["eddystone_url_freq_msec"] = nil
		}
	}
	if s.EddystoneUrlUrl.IsValueSet() {
		if s.EddystoneUrlUrl.Value() != nil {
			structMap["eddystone_url_url"] = s.EddystoneUrlUrl.Value()
		} else {
			structMap["eddystone_url_url"] = nil
		}
	}
	if s.IbeaconEnabled.IsValueSet() {
		if s.IbeaconEnabled.Value() != nil {
			structMap["ibeacon_enabled"] = s.IbeaconEnabled.Value()
		} else {
			structMap["ibeacon_enabled"] = nil
		}
	}
	if s.IbeaconFreqMsec.IsValueSet() {
		if s.IbeaconFreqMsec.Value() != nil {
			structMap["ibeacon_freq_msec"] = s.IbeaconFreqMsec.Value()
		} else {
			structMap["ibeacon_freq_msec"] = nil
		}
	}
	if s.IbeaconMajor.IsValueSet() {
		if s.IbeaconMajor.Value() != nil {
			structMap["ibeacon_major"] = s.IbeaconMajor.Value()
		} else {
			structMap["ibeacon_major"] = nil
		}
	}
	if s.IbeaconMinor.IsValueSet() {
		if s.IbeaconMinor.Value() != nil {
			structMap["ibeacon_minor"] = s.IbeaconMinor.Value()
		} else {
			structMap["ibeacon_minor"] = nil
		}
	}
	if s.IbeaconUuid.IsValueSet() {
		if s.IbeaconUuid.Value() != nil {
			structMap["ibeacon_uuid"] = s.IbeaconUuid.Value()
		} else {
			structMap["ibeacon_uuid"] = nil
		}
	}
	if s.Major.IsValueSet() {
		if s.Major.Value() != nil {
			structMap["major"] = s.Major.Value()
		} else {
			structMap["major"] = nil
		}
	}
	if s.Minors != nil {
		structMap["minors"] = s.Minors
	}
	if s.Power.IsValueSet() {
		if s.Power.Value() != nil {
			structMap["power"] = s.Power.Value()
		} else {
			structMap["power"] = nil
		}
	}
	if s.RxBytes.IsValueSet() {
		if s.RxBytes.Value() != nil {
			structMap["rx_bytes"] = s.RxBytes.Value()
		} else {
			structMap["rx_bytes"] = nil
		}
	}
	if s.RxPkts.IsValueSet() {
		if s.RxPkts.Value() != nil {
			structMap["rx_pkts"] = s.RxPkts.Value()
		} else {
			structMap["rx_pkts"] = nil
		}
	}
	if s.TxBytes.IsValueSet() {
		if s.TxBytes.Value() != nil {
			structMap["tx_bytes"] = s.TxBytes.Value()
		} else {
			structMap["tx_bytes"] = nil
		}
	}
	if s.TxPkts.IsValueSet() {
		if s.TxPkts.Value() != nil {
			structMap["tx_pkts"] = s.TxPkts.Value()
		} else {
			structMap["tx_pkts"] = nil
		}
	}
	if s.TxResets.IsValueSet() {
		if s.TxResets.Value() != nil {
			structMap["tx_resets"] = s.TxResets.Value()
		} else {
			structMap["tx_resets"] = nil
		}
	}
	if s.Uuid.IsValueSet() {
		if s.Uuid.Value() != nil {
			structMap["uuid"] = s.Uuid.Value()
		} else {
			structMap["uuid"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApBle.
// It customizes the JSON unmarshaling process for StatsApBle objects.
func (s *StatsApBle) UnmarshalJSON(input []byte) error {
	var temp tempStatsApBle
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "beacon_enabled", "beacon_rate", "eddystone_uid_enabled", "eddystone_uid_freq_msec", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_enabled", "eddystone_url_freq_msec", "eddystone_url_url", "ibeacon_enabled", "ibeacon_freq_msec", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "major", "minors", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "tx_resets", "uuid")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.BeaconEnabled = temp.BeaconEnabled
	s.BeaconRate = temp.BeaconRate
	s.EddystoneUidEnabled = temp.EddystoneUidEnabled
	s.EddystoneUidFreqMsec = temp.EddystoneUidFreqMsec
	s.EddystoneUidInstance = temp.EddystoneUidInstance
	s.EddystoneUidNamespace = temp.EddystoneUidNamespace
	s.EddystoneUrlEnabled = temp.EddystoneUrlEnabled
	s.EddystoneUrlFreqMsec = temp.EddystoneUrlFreqMsec
	s.EddystoneUrlUrl = temp.EddystoneUrlUrl
	s.IbeaconEnabled = temp.IbeaconEnabled
	s.IbeaconFreqMsec = temp.IbeaconFreqMsec
	s.IbeaconMajor = temp.IbeaconMajor
	s.IbeaconMinor = temp.IbeaconMinor
	s.IbeaconUuid = temp.IbeaconUuid
	s.Major = temp.Major
	s.Minors = temp.Minors
	s.Power = temp.Power
	s.RxBytes = temp.RxBytes
	s.RxPkts = temp.RxPkts
	s.TxBytes = temp.TxBytes
	s.TxPkts = temp.TxPkts
	s.TxResets = temp.TxResets
	s.Uuid = temp.Uuid
	return nil
}

// tempStatsApBle is a temporary struct used for validating the fields of StatsApBle.
type tempStatsApBle struct {
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
	RxBytes               Optional[int64]     `json:"rx_bytes"`
	RxPkts                Optional[int64]     `json:"rx_pkts"`
	TxBytes               Optional[int64]     `json:"tx_bytes"`
	TxPkts                Optional[int64]     `json:"tx_pkts"`
	TxResets              Optional[int]       `json:"tx_resets"`
	Uuid                  Optional[uuid.UUID] `json:"uuid"`
}
