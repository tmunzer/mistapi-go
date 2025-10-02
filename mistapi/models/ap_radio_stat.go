// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApRadioStat represents a ApRadioStat struct.
// Radio stat
type ApRadioStat struct {
	// channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
	Bandwidth *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
	// Current channel the radio is running on
	Channel Optional[int] `json:"channel"`
	// Use dynamic chaining for downlink
	DynamicChainingEnabled Optional[bool] `json:"dynamic_chaining_enabled"`
	// Radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af)
	Mac        Optional[string] `json:"mac"`
	NoiseFloor Optional[int]    `json:"noise_floor"`
	NumClients Optional[int]    `json:"num_clients"`
	// How many WLANs are applied to the radio
	NumWlans *int `json:"num_wlans,omitempty"`
	// Transmit power (in dBm)
	Power Optional[int] `json:"power"`
	// Amount of traffic received since connection
	RxBytes Optional[int64] `json:"rx_bytes"`
	// Amount of packets received since connection
	RxPkts Optional[int64] `json:"rx_pkts"`
	// Amount of traffic sent since connection
	TxBytes Optional[int64] `json:"tx_bytes"`
	// Amount of packets sent since connection
	TxPkts Optional[int64]  `json:"tx_pkts"`
	Usage  Optional[string] `json:"usage"`
	// All utilization in percentage
	UtilAll Optional[int] `json:"util_all"`
	// Reception of "No Packets" utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise
	UtilNonWifi Optional[int] `json:"util_non_wifi"`
	// Reception of "In BSS" utilization in percentage, only frames that are received from AP/STAs within the BSS
	UtilRxInBss Optional[int] `json:"util_rx_in_bss"`
	// Reception of "Other BSS" utilization in percentage, all frames received from AP/STAs that are outside the BSS
	UtilRxOtherBss Optional[int] `json:"util_rx_other_bss"`
	// Transmission utilization in percentage
	UtilTx Optional[int] `json:"util_tx"`
	// Reception of "UnDecodable Wifi" utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio
	UtilUndecodableWifi Optional[int] `json:"util_undecodable_wifi"`
	// Reception of "No Category" utilization in percentage, all 802.11 frames that are corrupted at the receiver
	UtilUnknownWifi      Optional[int]          `json:"util_unknown_wifi"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApRadioStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApRadioStat) String() string {
	return fmt.Sprintf(
		"ApRadioStat[Bandwidth=%v, Channel=%v, DynamicChainingEnabled=%v, Mac=%v, NoiseFloor=%v, NumClients=%v, NumWlans=%v, Power=%v, RxBytes=%v, RxPkts=%v, TxBytes=%v, TxPkts=%v, Usage=%v, UtilAll=%v, UtilNonWifi=%v, UtilRxInBss=%v, UtilRxOtherBss=%v, UtilTx=%v, UtilUndecodableWifi=%v, UtilUnknownWifi=%v, AdditionalProperties=%v]",
		a.Bandwidth, a.Channel, a.DynamicChainingEnabled, a.Mac, a.NoiseFloor, a.NumClients, a.NumWlans, a.Power, a.RxBytes, a.RxPkts, a.TxBytes, a.TxPkts, a.Usage, a.UtilAll, a.UtilNonWifi, a.UtilRxInBss, a.UtilRxOtherBss, a.UtilTx, a.UtilUndecodableWifi, a.UtilUnknownWifi, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApRadioStat.
// It customizes the JSON marshaling process for ApRadioStat objects.
func (a ApRadioStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"bandwidth", "channel", "dynamic_chaining_enabled", "mac", "noise_floor", "num_clients", "num_wlans", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "usage", "util_all", "util_non_wifi", "util_rx_in_bss", "util_rx_other_bss", "util_tx", "util_undecodable_wifi", "util_unknown_wifi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApRadioStat object to a map representation for JSON marshaling.
func (a ApRadioStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Bandwidth != nil {
		structMap["bandwidth"] = a.Bandwidth
	}
	if a.Channel.IsValueSet() {
		if a.Channel.Value() != nil {
			structMap["channel"] = a.Channel.Value()
		} else {
			structMap["channel"] = nil
		}
	}
	if a.DynamicChainingEnabled.IsValueSet() {
		if a.DynamicChainingEnabled.Value() != nil {
			structMap["dynamic_chaining_enabled"] = a.DynamicChainingEnabled.Value()
		} else {
			structMap["dynamic_chaining_enabled"] = nil
		}
	}
	if a.Mac.IsValueSet() {
		if a.Mac.Value() != nil {
			structMap["mac"] = a.Mac.Value()
		} else {
			structMap["mac"] = nil
		}
	}
	if a.NoiseFloor.IsValueSet() {
		if a.NoiseFloor.Value() != nil {
			structMap["noise_floor"] = a.NoiseFloor.Value()
		} else {
			structMap["noise_floor"] = nil
		}
	}
	if a.NumClients.IsValueSet() {
		if a.NumClients.Value() != nil {
			structMap["num_clients"] = a.NumClients.Value()
		} else {
			structMap["num_clients"] = nil
		}
	}
	if a.NumWlans != nil {
		structMap["num_wlans"] = a.NumWlans
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
	if a.Usage.IsValueSet() {
		if a.Usage.Value() != nil {
			structMap["usage"] = a.Usage.Value()
		} else {
			structMap["usage"] = nil
		}
	}
	if a.UtilAll.IsValueSet() {
		if a.UtilAll.Value() != nil {
			structMap["util_all"] = a.UtilAll.Value()
		} else {
			structMap["util_all"] = nil
		}
	}
	if a.UtilNonWifi.IsValueSet() {
		if a.UtilNonWifi.Value() != nil {
			structMap["util_non_wifi"] = a.UtilNonWifi.Value()
		} else {
			structMap["util_non_wifi"] = nil
		}
	}
	if a.UtilRxInBss.IsValueSet() {
		if a.UtilRxInBss.Value() != nil {
			structMap["util_rx_in_bss"] = a.UtilRxInBss.Value()
		} else {
			structMap["util_rx_in_bss"] = nil
		}
	}
	if a.UtilRxOtherBss.IsValueSet() {
		if a.UtilRxOtherBss.Value() != nil {
			structMap["util_rx_other_bss"] = a.UtilRxOtherBss.Value()
		} else {
			structMap["util_rx_other_bss"] = nil
		}
	}
	if a.UtilTx.IsValueSet() {
		if a.UtilTx.Value() != nil {
			structMap["util_tx"] = a.UtilTx.Value()
		} else {
			structMap["util_tx"] = nil
		}
	}
	if a.UtilUndecodableWifi.IsValueSet() {
		if a.UtilUndecodableWifi.Value() != nil {
			structMap["util_undecodable_wifi"] = a.UtilUndecodableWifi.Value()
		} else {
			structMap["util_undecodable_wifi"] = nil
		}
	}
	if a.UtilUnknownWifi.IsValueSet() {
		if a.UtilUnknownWifi.Value() != nil {
			structMap["util_unknown_wifi"] = a.UtilUnknownWifi.Value()
		} else {
			structMap["util_unknown_wifi"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadioStat.
// It customizes the JSON unmarshaling process for ApRadioStat objects.
func (a *ApRadioStat) UnmarshalJSON(input []byte) error {
	var temp tempApRadioStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bandwidth", "channel", "dynamic_chaining_enabled", "mac", "noise_floor", "num_clients", "num_wlans", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "usage", "util_all", "util_non_wifi", "util_rx_in_bss", "util_rx_other_bss", "util_tx", "util_undecodable_wifi", "util_unknown_wifi")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Bandwidth = temp.Bandwidth
	a.Channel = temp.Channel
	a.DynamicChainingEnabled = temp.DynamicChainingEnabled
	a.Mac = temp.Mac
	a.NoiseFloor = temp.NoiseFloor
	a.NumClients = temp.NumClients
	a.NumWlans = temp.NumWlans
	a.Power = temp.Power
	a.RxBytes = temp.RxBytes
	a.RxPkts = temp.RxPkts
	a.TxBytes = temp.TxBytes
	a.TxPkts = temp.TxPkts
	a.Usage = temp.Usage
	a.UtilAll = temp.UtilAll
	a.UtilNonWifi = temp.UtilNonWifi
	a.UtilRxInBss = temp.UtilRxInBss
	a.UtilRxOtherBss = temp.UtilRxOtherBss
	a.UtilTx = temp.UtilTx
	a.UtilUndecodableWifi = temp.UtilUndecodableWifi
	a.UtilUnknownWifi = temp.UtilUnknownWifi
	return nil
}

// tempApRadioStat is a temporary struct used for validating the fields of ApRadioStat.
type tempApRadioStat struct {
	Bandwidth              *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
	Channel                Optional[int]       `json:"channel"`
	DynamicChainingEnabled Optional[bool]      `json:"dynamic_chaining_enabled"`
	Mac                    Optional[string]    `json:"mac"`
	NoiseFloor             Optional[int]       `json:"noise_floor"`
	NumClients             Optional[int]       `json:"num_clients"`
	NumWlans               *int                `json:"num_wlans,omitempty"`
	Power                  Optional[int]       `json:"power"`
	RxBytes                Optional[int64]     `json:"rx_bytes"`
	RxPkts                 Optional[int64]     `json:"rx_pkts"`
	TxBytes                Optional[int64]     `json:"tx_bytes"`
	TxPkts                 Optional[int64]     `json:"tx_pkts"`
	Usage                  Optional[string]    `json:"usage"`
	UtilAll                Optional[int]       `json:"util_all"`
	UtilNonWifi            Optional[int]       `json:"util_non_wifi"`
	UtilRxInBss            Optional[int]       `json:"util_rx_in_bss"`
	UtilRxOtherBss         Optional[int]       `json:"util_rx_other_bss"`
	UtilTx                 Optional[int]       `json:"util_tx"`
	UtilUndecodableWifi    Optional[int]       `json:"util_undecodable_wifi"`
	UtilUnknownWifi        Optional[int]       `json:"util_unknown_wifi"`
}
