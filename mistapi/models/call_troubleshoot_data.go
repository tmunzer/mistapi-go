package models

import (
    "encoding/json"
)

// CallTroubleshootData represents a CallTroubleshootData struct.
type CallTroubleshootData struct {
    ApNumClients          *float64       `json:"ap_num_clients,omitempty"`
    ApRtt                 *float64       `json:"ap_rtt,omitempty"`
    ClientCpu             *float64       `json:"client_cpu,omitempty"`
    ClientNStreams        *float64       `json:"client_n_streams,omitempty"`
    ClientRadioBand       *float64       `json:"client_radio_band,omitempty"`
    ClientRssi            *float64       `json:"client_rssi,omitempty"`
    ClientRxBytes         *float64       `json:"client_rx_bytes,omitempty"`
    ClientRxRates         *float64       `json:"client_rx_rates,omitempty"`
    ClientRxRetries       *float64       `json:"client_rx_retries,omitempty"`
    ClientTxBytes         *float64       `json:"client_tx_bytes,omitempty"`
    ClientTxRates         *float64       `json:"client_tx_rates,omitempty"`
    ClientTxRetries       *float64       `json:"client_tx_retries,omitempty"`
    ClientVpnDistaince    *float64       `json:"client_vpn_distaince,omitempty"`
    ClientWifiVersion     *float64       `json:"client_wifi_version,omitempty"`
    Expected              *float64       `json:"expected,omitempty"`
    RadioBandwidth        *float64       `json:"radio_bandwidth,omitempty"`
    RadioChannel          *float64       `json:"radio_channel,omitempty"`
    RadioTxPower          *float64       `json:"radio_tx_power,omitempty"`
    RadioUtil             *float64       `json:"radio_util,omitempty"`
    RadioUtilInterference *float64       `json:"radio_util_interference,omitempty"`
    SiteNumClients        *float64       `json:"site_num_clients,omitempty"`
    WanAvgDownloadMbps    *float64       `json:"wan_avg_download_mbps,omitempty"`
    WanAvgUploadMbps      *float64       `json:"wan_avg_upload_mbps,omitempty"`
    WanJitter             *float64       `json:"wan_jitter,omitempty"`
    WanMaxDownloadMbps    *float64       `json:"wan_max_download_mbps,omitempty"`
    WanMaxUploadMbps      *float64       `json:"wan_max_upload_mbps,omitempty"`
    WanRtt                *float64       `json:"wan_rtt,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootData.
// It customizes the JSON marshaling process for CallTroubleshootData objects.
func (c CallTroubleshootData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootData object to a map representation for JSON marshaling.
func (c CallTroubleshootData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApNumClients != nil {
        structMap["ap_num_clients"] = c.ApNumClients
    }
    if c.ApRtt != nil {
        structMap["ap_rtt"] = c.ApRtt
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
    if c.ClientVpnDistaince != nil {
        structMap["client_vpn_distaince"] = c.ClientVpnDistaince
    }
    if c.ClientWifiVersion != nil {
        structMap["client_wifi_version"] = c.ClientWifiVersion
    }
    if c.Expected != nil {
        structMap["expected"] = c.Expected
    }
    if c.RadioBandwidth != nil {
        structMap["radio_bandwidth"] = c.RadioBandwidth
    }
    if c.RadioChannel != nil {
        structMap["radio_channel"] = c.RadioChannel
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
    if c.WanAvgDownloadMbps != nil {
        structMap["wan_avg_download_mbps"] = c.WanAvgDownloadMbps
    }
    if c.WanAvgUploadMbps != nil {
        structMap["wan_avg_upload_mbps"] = c.WanAvgUploadMbps
    }
    if c.WanJitter != nil {
        structMap["wan_jitter"] = c.WanJitter
    }
    if c.WanMaxDownloadMbps != nil {
        structMap["wan_max_download_mbps"] = c.WanMaxDownloadMbps
    }
    if c.WanMaxUploadMbps != nil {
        structMap["wan_max_upload_mbps"] = c.WanMaxUploadMbps
    }
    if c.WanRtt != nil {
        structMap["wan_rtt"] = c.WanRtt
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshootData.
// It customizes the JSON unmarshaling process for CallTroubleshootData objects.
func (c *CallTroubleshootData) UnmarshalJSON(input []byte) error {
    var temp tempCallTroubleshootData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_num_clients", "ap_rtt", "client_cpu", "client_n_streams", "client_radio_band", "client_rssi", "client_rx_bytes", "client_rx_rates", "client_rx_retries", "client_tx_bytes", "client_tx_rates", "client_tx_retries", "client_vpn_distaince", "client_wifi_version", "expected", "radio_bandwidth", "radio_channel", "radio_tx_power", "radio_util", "radio_util_interference", "site_num_clients", "wan_avg_download_mbps", "wan_avg_upload_mbps", "wan_jitter", "wan_max_download_mbps", "wan_max_upload_mbps", "wan_rtt")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ApNumClients = temp.ApNumClients
    c.ApRtt = temp.ApRtt
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
    c.ClientVpnDistaince = temp.ClientVpnDistaince
    c.ClientWifiVersion = temp.ClientWifiVersion
    c.Expected = temp.Expected
    c.RadioBandwidth = temp.RadioBandwidth
    c.RadioChannel = temp.RadioChannel
    c.RadioTxPower = temp.RadioTxPower
    c.RadioUtil = temp.RadioUtil
    c.RadioUtilInterference = temp.RadioUtilInterference
    c.SiteNumClients = temp.SiteNumClients
    c.WanAvgDownloadMbps = temp.WanAvgDownloadMbps
    c.WanAvgUploadMbps = temp.WanAvgUploadMbps
    c.WanJitter = temp.WanJitter
    c.WanMaxDownloadMbps = temp.WanMaxDownloadMbps
    c.WanMaxUploadMbps = temp.WanMaxUploadMbps
    c.WanRtt = temp.WanRtt
    return nil
}

// tempCallTroubleshootData is a temporary struct used for validating the fields of CallTroubleshootData.
type tempCallTroubleshootData  struct {
    ApNumClients          *float64 `json:"ap_num_clients,omitempty"`
    ApRtt                 *float64 `json:"ap_rtt,omitempty"`
    ClientCpu             *float64 `json:"client_cpu,omitempty"`
    ClientNStreams        *float64 `json:"client_n_streams,omitempty"`
    ClientRadioBand       *float64 `json:"client_radio_band,omitempty"`
    ClientRssi            *float64 `json:"client_rssi,omitempty"`
    ClientRxBytes         *float64 `json:"client_rx_bytes,omitempty"`
    ClientRxRates         *float64 `json:"client_rx_rates,omitempty"`
    ClientRxRetries       *float64 `json:"client_rx_retries,omitempty"`
    ClientTxBytes         *float64 `json:"client_tx_bytes,omitempty"`
    ClientTxRates         *float64 `json:"client_tx_rates,omitempty"`
    ClientTxRetries       *float64 `json:"client_tx_retries,omitempty"`
    ClientVpnDistaince    *float64 `json:"client_vpn_distaince,omitempty"`
    ClientWifiVersion     *float64 `json:"client_wifi_version,omitempty"`
    Expected              *float64 `json:"expected,omitempty"`
    RadioBandwidth        *float64 `json:"radio_bandwidth,omitempty"`
    RadioChannel          *float64 `json:"radio_channel,omitempty"`
    RadioTxPower          *float64 `json:"radio_tx_power,omitempty"`
    RadioUtil             *float64 `json:"radio_util,omitempty"`
    RadioUtilInterference *float64 `json:"radio_util_interference,omitempty"`
    SiteNumClients        *float64 `json:"site_num_clients,omitempty"`
    WanAvgDownloadMbps    *float64 `json:"wan_avg_download_mbps,omitempty"`
    WanAvgUploadMbps      *float64 `json:"wan_avg_upload_mbps,omitempty"`
    WanJitter             *float64 `json:"wan_jitter,omitempty"`
    WanMaxDownloadMbps    *float64 `json:"wan_max_download_mbps,omitempty"`
    WanMaxUploadMbps      *float64 `json:"wan_max_upload_mbps,omitempty"`
    WanRtt                *float64 `json:"wan_rtt,omitempty"`
}
