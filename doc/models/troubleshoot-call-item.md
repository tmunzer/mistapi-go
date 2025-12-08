
# Troubleshoot Call Item

## Structure

`TroubleshootCallItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApNumClients` | `*float64` | Optional | - |
| `ApRtt` | `*float64` | Optional | - |
| `AudioIn` | [`*models.CallTroubleshootData`](../../doc/models/call-troubleshoot-data.md) | Optional | - |
| `AudioOut` | [`*models.CallTroubleshootData`](../../doc/models/call-troubleshoot-data.md) | Optional | - |
| `ClientCpu` | `*float64` | Optional | - |
| `ClientNStreams` | `*float64` | Optional | - |
| `ClientRadioBand` | `*float64` | Optional | - |
| `ClientRssi` | `*float64` | Optional | - |
| `ClientRxBytes` | `*float64` | Optional | - |
| `ClientRxRates` | `*float64` | Optional | - |
| `ClientRxRetries` | `*float64` | Optional | - |
| `ClientTxBytes` | `*float64` | Optional | - |
| `ClientTxRates` | `*float64` | Optional | - |
| `ClientTxRetries` | `*float64` | Optional | - |
| `ClientVpnDistance` | `*float64` | Optional | - |
| `ClientWifiVersion` | `*float64` | Optional | - |
| `Expected` | `*float64` | Optional | - |
| `RadioApChange` | `*float64` | Optional | - |
| `RadioBandwidth` | `*float64` | Optional | - |
| `RadioChannel` | `*float64` | Optional | - |
| `RadioRxFailed` | `*float64` | Optional | - |
| `RadioTxPower` | `*float64` | Optional | - |
| `RadioUtil` | `*float64` | Optional | - |
| `RadioUtilInterference` | `*float64` | Optional | - |
| `SiteNumClients` | `*float64` | Optional | - |
| `SiteWanAvgDownloadMbps` | `*float64` | Optional | - |
| `SiteWanAvgUploadMbps` | `*float64` | Optional | - |
| `SiteWanDownloadMbps` | `*float64` | Optional | - |
| `SiteWanJitter` | `*float64` | Optional | - |
| `SiteWanRtt` | `*float64` | Optional | - |
| `SiteWanUploadMbps` | `*float64` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `VideoIn` | [`*models.CallTroubleshootData`](../../doc/models/call-troubleshoot-data.md) | Optional | - |
| `VideoOut` | [`*models.CallTroubleshootData`](../../doc/models/call-troubleshoot-data.md) | Optional | - |

## Example (as JSON)

```json
{
  "ap_num_clients": -0.08802365511655807,
  "ap_rtt": 0.09924473613500595,
  "client_cpu": 0.00834270566701889,
  "client_n_streams": 0.00734270566701889,
  "client_radio_band": 0.5841414928436279,
  "client_rssi": 0.7594696879386902,
  "client_rx_bytes": 2.365511655807E-05,
  "client_rx_rates": 0.02441493794322014,
  "client_rx_retries": -0.14325742423534393,
  "client_tx_bytes": 0.00102365511655807,
  "client_tx_rates": 0.22236637771129608,
  "client_tx_retries": 0.3308201730251312,
  "client_vpn_distance": -0.0001660545531194657,
  "client_wifi_version": 7.0566701889E-07,
  "expected": -2.8630001056670187,
  "radio_ap_change": 0.01850946433842182,
  "radio_bandwidth": -0.021175479516386986,
  "radio_channel": 0.11686426401138306,
  "radio_rx_failed": 1.1782013177871704,
  "radio_tx_power": 0.121039018034935,
  "radio_util": 0.2452986091375351,
  "radio_util_interference": 3.4367904663085938,
  "site_num_clients": 0.055026158690452576,
  "site_wan_avg_download_mbps": 3.0566701889E-07,
  "site_wan_avg_upload_mbps": 5.566701889E-08,
  "site_wan_download_mbps": 8.0566701889E-07,
  "site_wan_jitter": 1.00566701889E-06,
  "site_wan_rtt": 4.0566701889E-07,
  "site_wan_upload_mbps": 2.0566701889E-07,
  "audio_in": {
    "ap_num_clients": 152.32,
    "ap_rtt": 133.36,
    "client_cpu": 164.78,
    "client_n_streams": 206.36,
    "client_radio_band": 43.4
  },
  "audio_out": {
    "ap_num_clients": 71.16,
    "ap_rtt": 52.2,
    "client_cpu": 245.94,
    "client_n_streams": 125.2,
    "client_radio_band": 218.24
  }
}
```

