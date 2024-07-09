package models

import (
    "encoding/json"
)

// ApRadioStats represents a ApRadioStats struct.
// radio stat
type ApRadioStats struct {
    // channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6
    Bandwidth              *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
    // current channel the radio is running on
    Channel                *int                `json:"channel,omitempty"`
    // Use dynamic chaining for downlink
    DynamicChainingEnalbed *bool               `json:"dynamic_chaining_enalbed,omitempty"`
    // radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af)
    Mac                    *string             `json:"mac,omitempty"`
    NumClients             *int                `json:"num_clients,omitempty"`
    // transmit power (in dBm)
    Power                  *int                `json:"power,omitempty"`
    RxBytes                *float64            `json:"rx_bytes,omitempty"`
    RxPkts                 *float64            `json:"rx_pkts,omitempty"`
    TxBytes                *float64            `json:"tx_bytes,omitempty"`
    TxPkts                 *float64            `json:"tx_pkts,omitempty"`
    // all utilization in percentage
    UtilAll                *int                `json:"util_all,omitempty"`
    // reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise
    UtilNonWifi            *int                `json:"util_non_wifi,omitempty"`
    // reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS
    UtilRxInBss            *int                `json:"util_rx_in_bss,omitempty"`
    // reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS
    UtilRxOtherBss         *int                `json:"util_rx_other_bss,omitempty"`
    // transmission utilization in percentage
    UtilTx                 *int                `json:"util_tx,omitempty"`
    // reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio
    UtilUndecodableWifi    *int                `json:"util_undecodable_wifi,omitempty"`
    // reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver
    UtilUnknownWifi        *int                `json:"util_unknown_wifi,omitempty"`
    AdditionalProperties   map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApRadioStats.
// It customizes the JSON marshaling process for ApRadioStats objects.
func (a ApRadioStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApRadioStats object to a map representation for JSON marshaling.
func (a ApRadioStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Bandwidth != nil {
        structMap["bandwidth"] = a.Bandwidth
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.DynamicChainingEnalbed != nil {
        structMap["dynamic_chaining_enalbed"] = a.DynamicChainingEnalbed
    }
    if a.Mac != nil {
        structMap["mac"] = a.Mac
    }
    if a.NumClients != nil {
        structMap["num_clients"] = a.NumClients
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
    if a.UtilAll != nil {
        structMap["util_all"] = a.UtilAll
    }
    if a.UtilNonWifi != nil {
        structMap["util_non_wifi"] = a.UtilNonWifi
    }
    if a.UtilRxInBss != nil {
        structMap["util_rx_in_bss"] = a.UtilRxInBss
    }
    if a.UtilRxOtherBss != nil {
        structMap["util_rx_other_bss"] = a.UtilRxOtherBss
    }
    if a.UtilTx != nil {
        structMap["util_tx"] = a.UtilTx
    }
    if a.UtilUndecodableWifi != nil {
        structMap["util_undecodable_wifi"] = a.UtilUndecodableWifi
    }
    if a.UtilUnknownWifi != nil {
        structMap["util_unknown_wifi"] = a.UtilUnknownWifi
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadioStats.
// It customizes the JSON unmarshaling process for ApRadioStats objects.
func (a *ApRadioStats) UnmarshalJSON(input []byte) error {
    var temp apRadioStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bandwidth", "channel", "dynamic_chaining_enalbed", "mac", "num_clients", "power", "rx_bytes", "rx_pkts", "tx_bytes", "tx_pkts", "util_all", "util_non_wifi", "util_rx_in_bss", "util_rx_other_bss", "util_tx", "util_undecodable_wifi", "util_unknown_wifi")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Bandwidth = temp.Bandwidth
    a.Channel = temp.Channel
    a.DynamicChainingEnalbed = temp.DynamicChainingEnalbed
    a.Mac = temp.Mac
    a.NumClients = temp.NumClients
    a.Power = temp.Power
    a.RxBytes = temp.RxBytes
    a.RxPkts = temp.RxPkts
    a.TxBytes = temp.TxBytes
    a.TxPkts = temp.TxPkts
    a.UtilAll = temp.UtilAll
    a.UtilNonWifi = temp.UtilNonWifi
    a.UtilRxInBss = temp.UtilRxInBss
    a.UtilRxOtherBss = temp.UtilRxOtherBss
    a.UtilTx = temp.UtilTx
    a.UtilUndecodableWifi = temp.UtilUndecodableWifi
    a.UtilUnknownWifi = temp.UtilUnknownWifi
    return nil
}

// apRadioStats is a temporary struct used for validating the fields of ApRadioStats.
type apRadioStats  struct {
    Bandwidth              *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
    Channel                *int                `json:"channel,omitempty"`
    DynamicChainingEnalbed *bool               `json:"dynamic_chaining_enalbed,omitempty"`
    Mac                    *string             `json:"mac,omitempty"`
    NumClients             *int                `json:"num_clients,omitempty"`
    Power                  *int                `json:"power,omitempty"`
    RxBytes                *float64            `json:"rx_bytes,omitempty"`
    RxPkts                 *float64            `json:"rx_pkts,omitempty"`
    TxBytes                *float64            `json:"tx_bytes,omitempty"`
    TxPkts                 *float64            `json:"tx_pkts,omitempty"`
    UtilAll                *int                `json:"util_all,omitempty"`
    UtilNonWifi            *int                `json:"util_non_wifi,omitempty"`
    UtilRxInBss            *int                `json:"util_rx_in_bss,omitempty"`
    UtilRxOtherBss         *int                `json:"util_rx_other_bss,omitempty"`
    UtilTx                 *int                `json:"util_tx,omitempty"`
    UtilUndecodableWifi    *int                `json:"util_undecodable_wifi,omitempty"`
    UtilUnknownWifi        *int                `json:"util_unknown_wifi,omitempty"`
}
