
# Call Troubleshoot Summary Data

Media-direction call troubleshooting metrics in a summary response

## Structure

`CallTroubleshootSummaryData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApNumClients` | `*float64` | Optional | Media-direction metric value for AP client count |
| `ApRtt` | `*float64` | Optional | Media-direction metric value for AP round-trip time |
| `ClientCpu` | `*float64` | Optional | Media-direction metric value for client CPU load |
| `ClientNStreams` | `*float64` | Optional | Media-direction metric value for the number of client spatial streams |
| `ClientRadioBand` | `*float64` | Optional | Media-direction metric value for the client radio band |
| `ClientRssi` | `*float64` | Optional | Media-direction metric value for client RSSI |
| `ClientRxBytes` | `*float64` | Optional | Media-direction metric value for bytes received by the client |
| `ClientRxRates` | `*float64` | Optional | Media-direction metric value for client receive data rates |
| `ClientTxBytes` | `*float64` | Optional | Media-direction metric value for bytes transmitted by the client |
| `ClientTxRates` | `*float64` | Optional | Media-direction metric value for client transmit data rates |
| `ClientTxRetries` | `*float64` | Optional | Media-direction metric value for client transmit retries |
| `ClientVpnDistance` | `*float64` | Optional | Media-direction metric value for client VPN distance |
| `ClientWifiVersion` | `*float64` | Optional | Media-direction metric value for the client Wi-Fi version |
| `Expected` | `*float64` | Optional | Model baseline value expected for this media direction |
| `RadioBandwidth` | `*float64` | Optional | Media-direction metric value for radio channel bandwidth |
| `RadioChannel` | `*float64` | Optional | Media-direction metric value for the radio channel |
| `RadioTxPower` | `*float64` | Optional | Media-direction metric value for radio transmit power |
| `RadioUtil` | `*float64` | Optional | Media-direction metric value for radio utilization |
| `RadioUtilInterference` | `*float64` | Optional | Media-direction metric value for radio interference utilization |
| `SiteNumClients` | `*float64` | Optional | Media-direction metric value for the number of clients at the site |
| `SiteWanAvgDownloadMbps` | `*float64` | Optional | Media-direction metric value for site WAN average download throughput |
| `SiteWanAvgUploadMbps` | `*float64` | Optional | Media-direction metric value for site WAN average upload throughput |
| `SiteWanDownloadMbps` | `*float64` | Optional | Media-direction metric value for site WAN download throughput |
| `SiteWanJitter` | `*float64` | Optional | Media-direction metric value for site WAN jitter |
| `SiteWanRtt` | `*float64` | Optional | Media-direction metric value for site WAN round-trip time |
| `SiteWanUploadMbps` | `*float64` | Optional | Media-direction metric value for site WAN upload throughput |
| `WanAvgDownloadMbps` | `*float64` | Optional | Media-direction metric value for average WAN download throughput |
| `WanAvgUploadMbps` | `*float64` | Optional | Media-direction metric value for average WAN upload throughput |
| `WanJitter` | `*float64` | Optional | Media-direction metric value for WAN jitter |
| `WanMaxDownloadMbps` | `*float64` | Optional | Media-direction metric value for maximum WAN download throughput |
| `WanMaxUploadMbps` | `*float64` | Optional | Media-direction metric value for maximum WAN upload throughput |
| `WanRtt` | `*float64` | Optional | Media-direction metric value for WAN round-trip time |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    callTroubleshootSummaryData := models.CallTroubleshootSummaryData{
        ApNumClients:           models.ToPointer(float64(-0.6565111)),
        ApRtt:                  models.ToPointer(float64(0.16559607)),
        ClientCpu:              models.ToPointer(float64(3.7028809)),
        ClientNStreams:         models.ToPointer(float64(0.15803306)),
        ClientRadioBand:        models.ToPointer(float64(0.5576923)),
        ClientRssi:             models.ToPointer(float64(-1.0839354)),
        ClientRxBytes:          models.ToPointer(float64(2.2622051)),
        ClientRxRates:          models.ToPointer(float64(0.26726437)),
        ClientTxBytes:          models.ToPointer(float64(6.6164713)),
        ClientTxRates:          models.ToPointer(float64(0.62357205)),
        ClientTxRetries:        models.ToPointer(float64(1.702031)),
        ClientVpnDistance:      models.ToPointer(float64(1.6474955)),
        ClientWifiVersion:      models.ToPointer(float64(0.18267937)),
        Expected:               models.ToPointer(float64(30.941595)),
        RadioBandwidth:         models.ToPointer(float64(-0.06538621)),
        RadioChannel:           models.ToPointer(float64(-0.73391086)),
        RadioTxPower:           models.ToPointer(float64(0.10027129)),
        RadioUtil:              models.ToPointer(float64(12.770318)),
        RadioUtilInterference:  models.ToPointer(float64(-3.079999)),
        SiteNumClients:         models.ToPointer(float64(0.017364305)),
        SiteWanAvgDownloadMbps: models.ToPointer(float64(3.0566701889E-07)),
        SiteWanAvgUploadMbps:   models.ToPointer(float64(5.566701889E-08)),
        SiteWanDownloadMbps:    models.ToPointer(float64(8.0566701889E-07)),
        SiteWanJitter:          models.ToPointer(float64(0.7875519659784105)),
        SiteWanRtt:             models.ToPointer(float64(15.094849904378256)),
        SiteWanUploadMbps:      models.ToPointer(float64(2.0566701889E-07)),
        WanAvgDownloadMbps:     models.ToPointer(float64(1.4803165)),
        WanAvgUploadMbps:       models.ToPointer(float64(-0.038184267)),
        WanJitter:              models.ToPointer(float64(5.9680853)),
        WanMaxDownloadMbps:     models.ToPointer(float64(1.4803165)),
        WanMaxUploadMbps:       models.ToPointer(float64(-0.038184267)),
        WanRtt:                 models.ToPointer(float64(46.77899)),
    }

}
```

