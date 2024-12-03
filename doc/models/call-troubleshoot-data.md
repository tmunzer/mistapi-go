
# Call Troubleshoot Data

*This model accepts additional fields of type interface{}.*

## Structure

`CallTroubleshootData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApNumClients` | `*float64` | Optional | - |
| `ApRtt` | `*float64` | Optional | - |
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
| `ClientVpnDistaince` | `*float64` | Optional | - |
| `ClientWifiVersion` | `*float64` | Optional | - |
| `Expected` | `*float64` | Optional | - |
| `RadioBandwidth` | `*float64` | Optional | - |
| `RadioChannel` | `*float64` | Optional | - |
| `RadioTxPower` | `*float64` | Optional | - |
| `RadioUtil` | `*float64` | Optional | - |
| `RadioUtilInterference` | `*float64` | Optional | - |
| `SiteNumClients` | `*float64` | Optional | - |
| `WanAvgDownloadMbps` | `*float64` | Optional | - |
| `WanAvgUploadMbps` | `*float64` | Optional | - |
| `WanJitter` | `*float64` | Optional | - |
| `WanMaxDownloadMbps` | `*float64` | Optional | - |
| `WanMaxUploadMbps` | `*float64` | Optional | - |
| `WanRtt` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_num_clients": -0.6565111,
  "ap_rtt": 0.16559607,
  "client_cpu": 3.7028809,
  "client_n_streams": 0.15803306,
  "client_radio_band": 0.5576923,
  "client_rssi": -1.0839354,
  "client_rx_bytes": 2.2622051,
  "client_rx_rates": 0.62357205,
  "client_rx_retries": 0.26726437,
  "client_tx_bytes": 0.15803306,
  "client_tx_rates": 0.62357205,
  "client_tx_retries": 0.77553505,
  "client_vpn_distaince": 1.6474955,
  "client_wifi_version": 0.18267937,
  "expected": 30.941595,
  "radio_bandwidth": -0.06538621,
  "radio_channel": -0.73391086,
  "radio_tx_power": 0.10027129,
  "radio_util": 12.770318,
  "radio_util_interference": -3.079999,
  "site_num_clients": 0.017364305,
  "wan_avg_download_mbps": 1.4803165,
  "wan_avg_upload_mbps": -0.038184267,
  "wan_jitter": 5.9680853,
  "wan_max_download_mbps": 1.4803165,
  "wan_max_upload_mbps": -0.038184267,
  "wan_rtt": 46.77899,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

