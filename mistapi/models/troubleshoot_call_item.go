// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TroubleshootCallItem represents a TroubleshootCallItem struct.
type TroubleshootCallItem struct {
	ApNumClients           *float64              `json:"ap_num_clients,omitempty"`
	ApRtt                  *float64              `json:"ap_rtt,omitempty"`
	AudioIn                *CallTroubleshootData `json:"audio_in,omitempty"`
	AudioOut               *CallTroubleshootData `json:"audio_out,omitempty"`
	ClientCpu              *float64              `json:"client_cpu,omitempty"`
	ClientNStreams         *float64              `json:"client_n_streams,omitempty"`
	ClientRadioBand        *float64              `json:"client_radio_band,omitempty"`
	ClientRssi             *float64              `json:"client_rssi,omitempty"`
	ClientRxBytes          *float64              `json:"client_rx_bytes,omitempty"`
	ClientRxRates          *float64              `json:"client_rx_rates,omitempty"`
	ClientRxRetries        *float64              `json:"client_rx_retries,omitempty"`
	ClientTxBytes          *float64              `json:"client_tx_bytes,omitempty"`
	ClientTxRates          *float64              `json:"client_tx_rates,omitempty"`
	ClientTxRetries        *float64              `json:"client_tx_retries,omitempty"`
	ClientVpnDistance      *float64              `json:"client_vpn_distance,omitempty"`
	ClientWifiVersion      *float64              `json:"client_wifi_version,omitempty"`
	Expected               *float64              `json:"expected,omitempty"`
	RadioApChange          *float64              `json:"radio_ap_change,omitempty"`
	RadioBandwidth         *float64              `json:"radio_bandwidth,omitempty"`
	RadioChannel           *float64              `json:"radio_channel,omitempty"`
	RadioRxFailed          *float64              `json:"radio_rx_failed,omitempty"`
	RadioTxPower           *float64              `json:"radio_tx_power,omitempty"`
	RadioUtil              *float64              `json:"radio_util,omitempty"`
	RadioUtilInterference  *float64              `json:"radio_util_interference,omitempty"`
	SiteNumClients         *float64              `json:"site_num_clients,omitempty"`
	SiteWanAvgDownloadMbps *float64              `json:"site_wan_avg_download_mbps,omitempty"`
	SiteWanAvgUploadMbps   *float64              `json:"site_wan_avg_upload_mbps,omitempty"`
	SiteWanDownloadMbps    *float64              `json:"site_wan_download_mbps,omitempty"`
	SiteWanJitter          *float64              `json:"site_wan_jitter,omitempty"`
	SiteWanRtt             *float64              `json:"site_wan_rtt,omitempty"`
	SiteWanUploadMbps      *float64              `json:"site_wan_upload_mbps,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	VideoIn              *CallTroubleshootData  `json:"video_in,omitempty"`
	VideoOut             *CallTroubleshootData  `json:"video_out,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TroubleshootCallItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TroubleshootCallItem) String() string {
	return fmt.Sprintf(
		"TroubleshootCallItem[ApNumClients=%v, ApRtt=%v, AudioIn=%v, AudioOut=%v, ClientCpu=%v, ClientNStreams=%v, ClientRadioBand=%v, ClientRssi=%v, ClientRxBytes=%v, ClientRxRates=%v, ClientRxRetries=%v, ClientTxBytes=%v, ClientTxRates=%v, ClientTxRetries=%v, ClientVpnDistance=%v, ClientWifiVersion=%v, Expected=%v, RadioApChange=%v, RadioBandwidth=%v, RadioChannel=%v, RadioRxFailed=%v, RadioTxPower=%v, RadioUtil=%v, RadioUtilInterference=%v, SiteNumClients=%v, SiteWanAvgDownloadMbps=%v, SiteWanAvgUploadMbps=%v, SiteWanDownloadMbps=%v, SiteWanJitter=%v, SiteWanRtt=%v, SiteWanUploadMbps=%v, Timestamp=%v, VideoIn=%v, VideoOut=%v, AdditionalProperties=%v]",
		t.ApNumClients, t.ApRtt, t.AudioIn, t.AudioOut, t.ClientCpu, t.ClientNStreams, t.ClientRadioBand, t.ClientRssi, t.ClientRxBytes, t.ClientRxRates, t.ClientRxRetries, t.ClientTxBytes, t.ClientTxRates, t.ClientTxRetries, t.ClientVpnDistance, t.ClientWifiVersion, t.Expected, t.RadioApChange, t.RadioBandwidth, t.RadioChannel, t.RadioRxFailed, t.RadioTxPower, t.RadioUtil, t.RadioUtilInterference, t.SiteNumClients, t.SiteWanAvgDownloadMbps, t.SiteWanAvgUploadMbps, t.SiteWanDownloadMbps, t.SiteWanJitter, t.SiteWanRtt, t.SiteWanUploadMbps, t.Timestamp, t.VideoIn, t.VideoOut, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TroubleshootCallItem.
// It customizes the JSON marshaling process for TroubleshootCallItem objects.
func (t TroubleshootCallItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"ap_num_clients", "ap_rtt", "audio_in", "audio_out", "client_cpu", "client_n_streams", "client_radio_band", "client_rssi", "client_rx_bytes", "client_rx_rates", "client_rx_retries", "client_tx_bytes", "client_tx_rates", "client_tx_retries", "client_vpn_distance", "client_wifi_version", "expected", "radio_ap_change", "radio_bandwidth", "radio_channel", "radio_rx_failed", "radio_tx_power", "radio_util", "radio_util_interference", "site_num_clients", "site_wan_avg_download_mbps", "site_wan_avg_upload_mbps", "site_wan_download_mbps", "site_wan_jitter", "site_wan_rtt", "site_wan_upload_mbps", "timestamp", "video_in", "video_out"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TroubleshootCallItem object to a map representation for JSON marshaling.
func (t TroubleshootCallItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.ApNumClients != nil {
		structMap["ap_num_clients"] = t.ApNumClients
	}
	if t.ApRtt != nil {
		structMap["ap_rtt"] = t.ApRtt
	}
	if t.AudioIn != nil {
		structMap["audio_in"] = t.AudioIn.toMap()
	}
	if t.AudioOut != nil {
		structMap["audio_out"] = t.AudioOut.toMap()
	}
	if t.ClientCpu != nil {
		structMap["client_cpu"] = t.ClientCpu
	}
	if t.ClientNStreams != nil {
		structMap["client_n_streams"] = t.ClientNStreams
	}
	if t.ClientRadioBand != nil {
		structMap["client_radio_band"] = t.ClientRadioBand
	}
	if t.ClientRssi != nil {
		structMap["client_rssi"] = t.ClientRssi
	}
	if t.ClientRxBytes != nil {
		structMap["client_rx_bytes"] = t.ClientRxBytes
	}
	if t.ClientRxRates != nil {
		structMap["client_rx_rates"] = t.ClientRxRates
	}
	if t.ClientRxRetries != nil {
		structMap["client_rx_retries"] = t.ClientRxRetries
	}
	if t.ClientTxBytes != nil {
		structMap["client_tx_bytes"] = t.ClientTxBytes
	}
	if t.ClientTxRates != nil {
		structMap["client_tx_rates"] = t.ClientTxRates
	}
	if t.ClientTxRetries != nil {
		structMap["client_tx_retries"] = t.ClientTxRetries
	}
	if t.ClientVpnDistance != nil {
		structMap["client_vpn_distance"] = t.ClientVpnDistance
	}
	if t.ClientWifiVersion != nil {
		structMap["client_wifi_version"] = t.ClientWifiVersion
	}
	if t.Expected != nil {
		structMap["expected"] = t.Expected
	}
	if t.RadioApChange != nil {
		structMap["radio_ap_change"] = t.RadioApChange
	}
	if t.RadioBandwidth != nil {
		structMap["radio_bandwidth"] = t.RadioBandwidth
	}
	if t.RadioChannel != nil {
		structMap["radio_channel"] = t.RadioChannel
	}
	if t.RadioRxFailed != nil {
		structMap["radio_rx_failed"] = t.RadioRxFailed
	}
	if t.RadioTxPower != nil {
		structMap["radio_tx_power"] = t.RadioTxPower
	}
	if t.RadioUtil != nil {
		structMap["radio_util"] = t.RadioUtil
	}
	if t.RadioUtilInterference != nil {
		structMap["radio_util_interference"] = t.RadioUtilInterference
	}
	if t.SiteNumClients != nil {
		structMap["site_num_clients"] = t.SiteNumClients
	}
	if t.SiteWanAvgDownloadMbps != nil {
		structMap["site_wan_avg_download_mbps"] = t.SiteWanAvgDownloadMbps
	}
	if t.SiteWanAvgUploadMbps != nil {
		structMap["site_wan_avg_upload_mbps"] = t.SiteWanAvgUploadMbps
	}
	if t.SiteWanDownloadMbps != nil {
		structMap["site_wan_download_mbps"] = t.SiteWanDownloadMbps
	}
	if t.SiteWanJitter != nil {
		structMap["site_wan_jitter"] = t.SiteWanJitter
	}
	if t.SiteWanRtt != nil {
		structMap["site_wan_rtt"] = t.SiteWanRtt
	}
	if t.SiteWanUploadMbps != nil {
		structMap["site_wan_upload_mbps"] = t.SiteWanUploadMbps
	}
	if t.Timestamp != nil {
		structMap["timestamp"] = t.Timestamp
	}
	if t.VideoIn != nil {
		structMap["video_in"] = t.VideoIn.toMap()
	}
	if t.VideoOut != nil {
		structMap["video_out"] = t.VideoOut.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TroubleshootCallItem.
// It customizes the JSON unmarshaling process for TroubleshootCallItem objects.
func (t *TroubleshootCallItem) UnmarshalJSON(input []byte) error {
	var temp tempTroubleshootCallItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_num_clients", "ap_rtt", "audio_in", "audio_out", "client_cpu", "client_n_streams", "client_radio_band", "client_rssi", "client_rx_bytes", "client_rx_rates", "client_rx_retries", "client_tx_bytes", "client_tx_rates", "client_tx_retries", "client_vpn_distance", "client_wifi_version", "expected", "radio_ap_change", "radio_bandwidth", "radio_channel", "radio_rx_failed", "radio_tx_power", "radio_util", "radio_util_interference", "site_num_clients", "site_wan_avg_download_mbps", "site_wan_avg_upload_mbps", "site_wan_download_mbps", "site_wan_jitter", "site_wan_rtt", "site_wan_upload_mbps", "timestamp", "video_in", "video_out")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.ApNumClients = temp.ApNumClients
	t.ApRtt = temp.ApRtt
	t.AudioIn = temp.AudioIn
	t.AudioOut = temp.AudioOut
	t.ClientCpu = temp.ClientCpu
	t.ClientNStreams = temp.ClientNStreams
	t.ClientRadioBand = temp.ClientRadioBand
	t.ClientRssi = temp.ClientRssi
	t.ClientRxBytes = temp.ClientRxBytes
	t.ClientRxRates = temp.ClientRxRates
	t.ClientRxRetries = temp.ClientRxRetries
	t.ClientTxBytes = temp.ClientTxBytes
	t.ClientTxRates = temp.ClientTxRates
	t.ClientTxRetries = temp.ClientTxRetries
	t.ClientVpnDistance = temp.ClientVpnDistance
	t.ClientWifiVersion = temp.ClientWifiVersion
	t.Expected = temp.Expected
	t.RadioApChange = temp.RadioApChange
	t.RadioBandwidth = temp.RadioBandwidth
	t.RadioChannel = temp.RadioChannel
	t.RadioRxFailed = temp.RadioRxFailed
	t.RadioTxPower = temp.RadioTxPower
	t.RadioUtil = temp.RadioUtil
	t.RadioUtilInterference = temp.RadioUtilInterference
	t.SiteNumClients = temp.SiteNumClients
	t.SiteWanAvgDownloadMbps = temp.SiteWanAvgDownloadMbps
	t.SiteWanAvgUploadMbps = temp.SiteWanAvgUploadMbps
	t.SiteWanDownloadMbps = temp.SiteWanDownloadMbps
	t.SiteWanJitter = temp.SiteWanJitter
	t.SiteWanRtt = temp.SiteWanRtt
	t.SiteWanUploadMbps = temp.SiteWanUploadMbps
	t.Timestamp = temp.Timestamp
	t.VideoIn = temp.VideoIn
	t.VideoOut = temp.VideoOut
	return nil
}

// tempTroubleshootCallItem is a temporary struct used for validating the fields of TroubleshootCallItem.
type tempTroubleshootCallItem struct {
	ApNumClients           *float64              `json:"ap_num_clients,omitempty"`
	ApRtt                  *float64              `json:"ap_rtt,omitempty"`
	AudioIn                *CallTroubleshootData `json:"audio_in,omitempty"`
	AudioOut               *CallTroubleshootData `json:"audio_out,omitempty"`
	ClientCpu              *float64              `json:"client_cpu,omitempty"`
	ClientNStreams         *float64              `json:"client_n_streams,omitempty"`
	ClientRadioBand        *float64              `json:"client_radio_band,omitempty"`
	ClientRssi             *float64              `json:"client_rssi,omitempty"`
	ClientRxBytes          *float64              `json:"client_rx_bytes,omitempty"`
	ClientRxRates          *float64              `json:"client_rx_rates,omitempty"`
	ClientRxRetries        *float64              `json:"client_rx_retries,omitempty"`
	ClientTxBytes          *float64              `json:"client_tx_bytes,omitempty"`
	ClientTxRates          *float64              `json:"client_tx_rates,omitempty"`
	ClientTxRetries        *float64              `json:"client_tx_retries,omitempty"`
	ClientVpnDistance      *float64              `json:"client_vpn_distance,omitempty"`
	ClientWifiVersion      *float64              `json:"client_wifi_version,omitempty"`
	Expected               *float64              `json:"expected,omitempty"`
	RadioApChange          *float64              `json:"radio_ap_change,omitempty"`
	RadioBandwidth         *float64              `json:"radio_bandwidth,omitempty"`
	RadioChannel           *float64              `json:"radio_channel,omitempty"`
	RadioRxFailed          *float64              `json:"radio_rx_failed,omitempty"`
	RadioTxPower           *float64              `json:"radio_tx_power,omitempty"`
	RadioUtil              *float64              `json:"radio_util,omitempty"`
	RadioUtilInterference  *float64              `json:"radio_util_interference,omitempty"`
	SiteNumClients         *float64              `json:"site_num_clients,omitempty"`
	SiteWanAvgDownloadMbps *float64              `json:"site_wan_avg_download_mbps,omitempty"`
	SiteWanAvgUploadMbps   *float64              `json:"site_wan_avg_upload_mbps,omitempty"`
	SiteWanDownloadMbps    *float64              `json:"site_wan_download_mbps,omitempty"`
	SiteWanJitter          *float64              `json:"site_wan_jitter,omitempty"`
	SiteWanRtt             *float64              `json:"site_wan_rtt,omitempty"`
	SiteWanUploadMbps      *float64              `json:"site_wan_upload_mbps,omitempty"`
	Timestamp              *float64              `json:"timestamp,omitempty"`
	VideoIn                *CallTroubleshootData `json:"video_in,omitempty"`
	VideoOut               *CallTroubleshootData `json:"video_out,omitempty"`
}
