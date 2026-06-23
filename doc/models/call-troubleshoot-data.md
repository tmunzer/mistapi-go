
# Call Troubleshoot Data

Per-media call troubleshooting metric values

## Structure

`CallTroubleshootData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApNumClients` | `*float64` | Optional | Troubleshooting metric value for AP client count |
| `ApRtt` | `*float64` | Optional | Troubleshooting metric value for AP round-trip time |
| `ClientCpu` | `*float64` | Optional | Troubleshooting metric value for client CPU load |
| `ClientNStreams` | `*float64` | Optional | Troubleshooting metric value for the number of client spatial streams |
| `ClientRadioBand` | `*float64` | Optional | Troubleshooting metric value for the client radio band |
| `ClientRssi` | `*float64` | Optional | Troubleshooting metric value for client RSSI |
| `ClientRxBytes` | `*float64` | Optional | Troubleshooting metric value for bytes received by the client |
| `ClientRxRates` | `*float64` | Optional | Troubleshooting metric value for client receive data rates |
| `ClientRxRetries` | `*float64` | Optional | Troubleshooting metric value for client receive retries |
| `ClientTxBytes` | `*float64` | Optional | Troubleshooting metric value for bytes transmitted by the client |
| `ClientTxRates` | `*float64` | Optional | Troubleshooting metric value for client transmit data rates |
| `ClientTxRetries` | `*float64` | Optional | Troubleshooting metric value for client transmit retries |
| `ClientVpnDistance` | `*float64` | Optional | Troubleshooting metric value for client VPN distance |
| `ClientWifiVersion` | `*float64` | Optional | Troubleshooting metric value for the client Wi-Fi version |
| `Expected` | `*float64` | Optional | Model baseline value expected for this media direction |
| `RadioBandwidth` | `*float64` | Optional | Troubleshooting metric value for radio channel bandwidth |
| `RadioChannel` | `*float64` | Optional | Troubleshooting metric value for the radio channel |
| `RadioTxPower` | `*float64` | Optional | Troubleshooting metric value for radio transmit power |
| `RadioUtil` | `*float64` | Optional | Troubleshooting metric value for radio utilization |
| `RadioUtilInterference` | `*float64` | Optional | Troubleshooting metric value for radio interference utilization |
| `SiteNumClients` | `*float64` | Optional | Troubleshooting metric value for the number of clients at the site |
| `WanAvgDownloadMbps` | `*float64` | Optional | Troubleshooting metric value for average WAN download throughput |
| `WanAvgUploadMbps` | `*float64` | Optional | Troubleshooting metric value for average WAN upload throughput |
| `WanJitter` | `*float64` | Optional | Troubleshooting metric value for WAN jitter |
| `WanMaxDownloadMbps` | `*float64` | Optional | Troubleshooting metric value for maximum WAN download throughput |
| `WanMaxUploadMbps` | `*float64` | Optional | Troubleshooting metric value for maximum WAN upload throughput |
| `WanRtt` | `*float64` | Optional | Troubleshooting metric value for WAN round-trip time |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    callTroubleshootData := models.CallTroubleshootData{
        ApNumClients:          models.ToPointer(float64(-0.6565111)),
        ApRtt:                 models.ToPointer(float64(0.16559607)),
        ClientCpu:             models.ToPointer(float64(3.7028809)),
        ClientNStreams:        models.ToPointer(float64(0.15803306)),
        ClientRadioBand:       models.ToPointer(float64(0.5576923)),
        ClientRssi:            models.ToPointer(float64(-1.0839354)),
        ClientRxBytes:         models.ToPointer(float64(2.2622051)),
        ClientRxRates:         models.ToPointer(float64(0.62357205)),
        ClientRxRetries:       models.ToPointer(float64(0.26726437)),
        ClientTxBytes:         models.ToPointer(float64(0.15803306)),
        ClientTxRates:         models.ToPointer(float64(0.62357205)),
        ClientTxRetries:       models.ToPointer(float64(0.77553505)),
        ClientVpnDistance:     models.ToPointer(float64(1.6474955)),
        ClientWifiVersion:     models.ToPointer(float64(0.18267937)),
        Expected:              models.ToPointer(float64(30.941595)),
        RadioBandwidth:        models.ToPointer(float64(-0.06538621)),
        RadioChannel:          models.ToPointer(float64(-0.73391086)),
        RadioTxPower:          models.ToPointer(float64(0.10027129)),
        RadioUtil:             models.ToPointer(float64(12.770318)),
        RadioUtilInterference: models.ToPointer(float64(-3.079999)),
        SiteNumClients:        models.ToPointer(float64(0.017364305)),
        WanAvgDownloadMbps:    models.ToPointer(float64(1.4803165)),
        WanAvgUploadMbps:      models.ToPointer(float64(-0.038184267)),
        WanJitter:             models.ToPointer(float64(5.9680853)),
        WanMaxDownloadMbps:    models.ToPointer(float64(1.4803165)),
        WanMaxUploadMbps:      models.ToPointer(float64(-0.038184267)),
        WanRtt:                models.ToPointer(float64(46.77899)),
    }

}
```

