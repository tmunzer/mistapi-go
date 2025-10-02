// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// CallTroubleshootSummary represents a CallTroubleshootSummary struct.
type CallTroubleshootSummary struct {
	ApNumClients           *float64                     `json:"ap_num_clients,omitempty"`
	ApRtt                  *float64                     `json:"ap_rtt,omitempty"`
	AudioIn                *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
	AudioOut               *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
	ClientCpu              *float64                     `json:"client_cpu,omitempty"`
	ClientNStreams         *float64                     `json:"client_n_streams,omitempty"`
	ClientRadioBand        *float64                     `json:"client_radio_band,omitempty"`
	ClientRssi             *float64                     `json:"client_rssi,omitempty"`
	ClientRxBytes          *float64                     `json:"client_rx_bytes,omitempty"`
	ClientRxRates          *float64                     `json:"client_rx_rates,omitempty"`
	ClientRxRetries        *float64                     `json:"client_rx_retries,omitempty"`
	ClientTxBytes          *float64                     `json:"client_tx_bytes,omitempty"`
	ClientTxRates          *float64                     `json:"client_tx_rates,omitempty"`
	ClientTxRetries        *float64                     `json:"client_tx_retries,omitempty"`
	ClientVpnDistance      *float64                     `json:"client_vpn_distance,omitempty"`
	ClientWifiVersion      *float64                     `json:"client_wifi_version,omitempty"`
	Expected               *float64                     `json:"expected,omitempty"`
	RadioApChange          *float64                     `json:"radio_ap_change,omitempty"`
	RadioBandwidth         *float64                     `json:"radio_bandwidth,omitempty"`
	RadioChannel           *float64                     `json:"radio_channel,omitempty"`
	RadioRxFailed          *float64                     `json:"radio_rx_failed,omitempty"`
	RadioTxPower           *float64                     `json:"radio_tx_power,omitempty"`
	RadioUtil              *float64                     `json:"radio_util,omitempty"`
	RadioUtilInterference  *float64                     `json:"radio_util_interference,omitempty"`
	SiteNumClients         *float64                     `json:"site_num_clients,omitempty"`
	SiteWanAvgDownloadMbps *float64                     `json:"site_wan_avg_download_mbps,omitempty"`
	SiteWanAvgUploadMbps   *float64                     `json:"site_wan_avg_upload_mbps,omitempty"`
	SiteWanDownloadMbps    *float64                     `json:"site_wan_download_mbps,omitempty"`
	SiteWanJitter          *float64                     `json:"site_wan_jitter,omitempty"`
	SiteWanRtt             *float64                     `json:"site_wan_rtt,omitempty"`
	SiteWanUploadMbps      *float64                     `json:"site_wan_upload_mbps,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64                     `json:"timestamp,omitempty"`
	VideoIn              *CallTroubleshootSummaryData `json:"video_in,omitempty"`
	VideoOut             *CallTroubleshootSummaryData `json:"video_out,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for CallTroubleshootSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CallTroubleshootSummary) String() string {
	return fmt.Sprintf(
		"CallTroubleshootSummary[ApNumClients=%v, ApRtt=%v, AudioIn=%v, AudioOut=%v, ClientCpu=%v, ClientNStreams=%v, ClientRadioBand=%v, ClientRssi=%v, ClientRxBytes=%v, ClientRxRates=%v, ClientRxRetries=%v, ClientTxBytes=%v, ClientTxRates=%v, ClientTxRetries=%v, ClientVpnDistance=%v, ClientWifiVersion=%v, Expected=%v, RadioApChange=%v, RadioBandwidth=%v, RadioChannel=%v, RadioRxFailed=%v, RadioTxPower=%v, RadioUtil=%v, RadioUtilInterference=%v, SiteNumClients=%v, SiteWanAvgDownloadMbps=%v, SiteWanAvgUploadMbps=%v, SiteWanDownloadMbps=%v, SiteWanJitter=%v, SiteWanRtt=%v, SiteWanUploadMbps=%v, Timestamp=%v, VideoIn=%v, VideoOut=%v, AdditionalProperties=%v]",
		c.ApNumClients, c.ApRtt, c.AudioIn, c.AudioOut, c.ClientCpu, c.ClientNStreams, c.ClientRadioBand, c.ClientRssi, c.ClientRxBytes, c.ClientRxRates, c.ClientRxRetries, c.ClientTxBytes, c.ClientTxRates, c.ClientTxRetries, c.ClientVpnDistance, c.ClientWifiVersion, c.Expected, c.RadioApChange, c.RadioBandwidth, c.RadioChannel, c.RadioRxFailed, c.RadioTxPower, c.RadioUtil, c.RadioUtilInterference, c.SiteNumClients, c.SiteWanAvgDownloadMbps, c.SiteWanAvgUploadMbps, c.SiteWanDownloadMbps, c.SiteWanJitter, c.SiteWanRtt, c.SiteWanUploadMbps, c.Timestamp, c.VideoIn, c.VideoOut, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootSummary.
// It customizes the JSON marshaling process for CallTroubleshootSummary objects.
func (c CallTroubleshootSummary) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"ap_num_clients", "ap_rtt", "audio_in", "audio_out", "client_cpu", "client_n_streams", "client_radio_band", "client_rssi", "client_rx_bytes", "client_rx_rates", "client_rx_retries", "client_tx_bytes", "client_tx_rates", "client_tx_retries", "client_vpn_distance", "client_wifi_version", "expected", "radio_ap_change", "radio_bandwidth", "radio_channel", "radio_rx_failed", "radio_tx_power", "radio_util", "radio_util_interference", "site_num_clients", "site_wan_avg_download_mbps", "site_wan_avg_upload_mbps", "site_wan_download_mbps", "site_wan_jitter", "site_wan_rtt", "site_wan_upload_mbps", "timestamp", "video_in", "video_out"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootSummary object to a map representation for JSON marshaling.
func (c CallTroubleshootSummary) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.ApNumClients != nil {
		structMap["ap_num_clients"] = c.ApNumClients
	}
	if c.ApRtt != nil {
		structMap["ap_rtt"] = c.ApRtt
	}
	if c.AudioIn != nil {
		structMap["audio_in"] = c.AudioIn.toMap()
	}
	if c.AudioOut != nil {
		structMap["audio_out"] = c.AudioOut.toMap()
	}
	if c.ClientCpu != nil {
		structMap["client_cpu"] = c.ClientCpu
	}
	if c.ClientNStreams != nil {
		structMap["client_n_streams"] = c.ClientNStreams
	}
	if c.ClientRadioBand != nil {
		structMap["client_radio_band"] = c.ClientRadioBand
	}
	if c.ClientRssi != nil {
		structMap["client_rssi"] = c.ClientRssi
	}
	if c.ClientRxBytes != nil {
		structMap["client_rx_bytes"] = c.ClientRxBytes
	}
	if c.ClientRxRates != nil {
		structMap["client_rx_rates"] = c.ClientRxRates
	}
	if c.ClientRxRetries != nil {
		structMap["client_rx_retries"] = c.ClientRxRetries
	}
	if c.ClientTxBytes != nil {
		structMap["client_tx_bytes"] = c.ClientTxBytes
	}
	if c.ClientTxRates != nil {
		structMap["client_tx_rates"] = c.ClientTxRates
	}
	if c.ClientTxRetries != nil {
		structMap["client_tx_retries"] = c.ClientTxRetries
	}
	if c.ClientVpnDistance != nil {
		structMap["client_vpn_distance"] = c.ClientVpnDistance
	}
	if c.ClientWifiVersion != nil {
		structMap["client_wifi_version"] = c.ClientWifiVersion
	}
	if c.Expected != nil {
		structMap["expected"] = c.Expected
	}
	if c.RadioApChange != nil {
		structMap["radio_ap_change"] = c.RadioApChange
	}
	if c.RadioBandwidth != nil {
		structMap["radio_bandwidth"] = c.RadioBandwidth
	}
	if c.RadioChannel != nil {
		structMap["radio_channel"] = c.RadioChannel
	}
	if c.RadioRxFailed != nil {
		structMap["radio_rx_failed"] = c.RadioRxFailed
	}
	if c.RadioTxPower != nil {
		structMap["radio_tx_power"] = c.RadioTxPower
	}
	if c.RadioUtil != nil {
		structMap["radio_util"] = c.RadioUtil
	}
	if c.RadioUtilInterference != nil {
		structMap["radio_util_interference"] = c.RadioUtilInterference
	}
	if c.SiteNumClients != nil {
		structMap["site_num_clients"] = c.SiteNumClients
	}
	if c.SiteWanAvgDownloadMbps != nil {
		structMap["site_wan_avg_download_mbps"] = c.SiteWanAvgDownloadMbps
	}
	if c.SiteWanAvgUploadMbps != nil {
		structMap["site_wan_avg_upload_mbps"] = c.SiteWanAvgUploadMbps
	}
	if c.SiteWanDownloadMbps != nil {
		structMap["site_wan_download_mbps"] = c.SiteWanDownloadMbps
	}
	if c.SiteWanJitter != nil {
		structMap["site_wan_jitter"] = c.SiteWanJitter
	}
	if c.SiteWanRtt != nil {
		structMap["site_wan_rtt"] = c.SiteWanRtt
	}
	if c.SiteWanUploadMbps != nil {
		structMap["site_wan_upload_mbps"] = c.SiteWanUploadMbps
	}
	if c.Timestamp != nil {
		structMap["timestamp"] = c.Timestamp
	}
	if c.VideoIn != nil {
		structMap["video_in"] = c.VideoIn.toMap()
	}
	if c.VideoOut != nil {
		structMap["video_out"] = c.VideoOut.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshootSummary.
// It customizes the JSON unmarshaling process for CallTroubleshootSummary objects.
func (c *CallTroubleshootSummary) UnmarshalJSON(input []byte) error {
	var temp tempCallTroubleshootSummary
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_num_clients", "ap_rtt", "audio_in", "audio_out", "client_cpu", "client_n_streams", "client_radio_band", "client_rssi", "client_rx_bytes", "client_rx_rates", "client_rx_retries", "client_tx_bytes", "client_tx_rates", "client_tx_retries", "client_vpn_distance", "client_wifi_version", "expected", "radio_ap_change", "radio_bandwidth", "radio_channel", "radio_rx_failed", "radio_tx_power", "radio_util", "radio_util_interference", "site_num_clients", "site_wan_avg_download_mbps", "site_wan_avg_upload_mbps", "site_wan_download_mbps", "site_wan_jitter", "site_wan_rtt", "site_wan_upload_mbps", "timestamp", "video_in", "video_out")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.ApNumClients = temp.ApNumClients
	c.ApRtt = temp.ApRtt
	c.AudioIn = temp.AudioIn
	c.AudioOut = temp.AudioOut
	c.ClientCpu = temp.ClientCpu
	c.ClientNStreams = temp.ClientNStreams
	c.ClientRadioBand = temp.ClientRadioBand
	c.ClientRssi = temp.ClientRssi
	c.ClientRxBytes = temp.ClientRxBytes
	c.ClientRxRates = temp.ClientRxRates
	c.ClientRxRetries = temp.ClientRxRetries
	c.ClientTxBytes = temp.ClientTxBytes
	c.ClientTxRates = temp.ClientTxRates
	c.ClientTxRetries = temp.ClientTxRetries
	c.ClientVpnDistance = temp.ClientVpnDistance
	c.ClientWifiVersion = temp.ClientWifiVersion
	c.Expected = temp.Expected
	c.RadioApChange = temp.RadioApChange
	c.RadioBandwidth = temp.RadioBandwidth
	c.RadioChannel = temp.RadioChannel
	c.RadioRxFailed = temp.RadioRxFailed
	c.RadioTxPower = temp.RadioTxPower
	c.RadioUtil = temp.RadioUtil
	c.RadioUtilInterference = temp.RadioUtilInterference
	c.SiteNumClients = temp.SiteNumClients
	c.SiteWanAvgDownloadMbps = temp.SiteWanAvgDownloadMbps
	c.SiteWanAvgUploadMbps = temp.SiteWanAvgUploadMbps
	c.SiteWanDownloadMbps = temp.SiteWanDownloadMbps
	c.SiteWanJitter = temp.SiteWanJitter
	c.SiteWanRtt = temp.SiteWanRtt
	c.SiteWanUploadMbps = temp.SiteWanUploadMbps
	c.Timestamp = temp.Timestamp
	c.VideoIn = temp.VideoIn
	c.VideoOut = temp.VideoOut
	return nil
}

// tempCallTroubleshootSummary is a temporary struct used for validating the fields of CallTroubleshootSummary.
type tempCallTroubleshootSummary struct {
	ApNumClients           *float64                     `json:"ap_num_clients,omitempty"`
	ApRtt                  *float64                     `json:"ap_rtt,omitempty"`
	AudioIn                *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
	AudioOut               *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
	ClientCpu              *float64                     `json:"client_cpu,omitempty"`
	ClientNStreams         *float64                     `json:"client_n_streams,omitempty"`
	ClientRadioBand        *float64                     `json:"client_radio_band,omitempty"`
	ClientRssi             *float64                     `json:"client_rssi,omitempty"`
	ClientRxBytes          *float64                     `json:"client_rx_bytes,omitempty"`
	ClientRxRates          *float64                     `json:"client_rx_rates,omitempty"`
	ClientRxRetries        *float64                     `json:"client_rx_retries,omitempty"`
	ClientTxBytes          *float64                     `json:"client_tx_bytes,omitempty"`
	ClientTxRates          *float64                     `json:"client_tx_rates,omitempty"`
	ClientTxRetries        *float64                     `json:"client_tx_retries,omitempty"`
	ClientVpnDistance      *float64                     `json:"client_vpn_distance,omitempty"`
	ClientWifiVersion      *float64                     `json:"client_wifi_version,omitempty"`
	Expected               *float64                     `json:"expected,omitempty"`
	RadioApChange          *float64                     `json:"radio_ap_change,omitempty"`
	RadioBandwidth         *float64                     `json:"radio_bandwidth,omitempty"`
	RadioChannel           *float64                     `json:"radio_channel,omitempty"`
	RadioRxFailed          *float64                     `json:"radio_rx_failed,omitempty"`
	RadioTxPower           *float64                     `json:"radio_tx_power,omitempty"`
	RadioUtil              *float64                     `json:"radio_util,omitempty"`
	RadioUtilInterference  *float64                     `json:"radio_util_interference,omitempty"`
	SiteNumClients         *float64                     `json:"site_num_clients,omitempty"`
	SiteWanAvgDownloadMbps *float64                     `json:"site_wan_avg_download_mbps,omitempty"`
	SiteWanAvgUploadMbps   *float64                     `json:"site_wan_avg_upload_mbps,omitempty"`
	SiteWanDownloadMbps    *float64                     `json:"site_wan_download_mbps,omitempty"`
	SiteWanJitter          *float64                     `json:"site_wan_jitter,omitempty"`
	SiteWanRtt             *float64                     `json:"site_wan_rtt,omitempty"`
	SiteWanUploadMbps      *float64                     `json:"site_wan_upload_mbps,omitempty"`
	Timestamp              *float64                     `json:"timestamp,omitempty"`
	VideoIn                *CallTroubleshootSummaryData `json:"video_in,omitempty"`
	VideoOut               *CallTroubleshootSummaryData `json:"video_out,omitempty"`
}
