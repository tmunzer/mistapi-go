
# Call Troubleshoot Summary

Summary row for call troubleshooting metrics

## Structure

`CallTroubleshootSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApNumClients` | `*float64` | Optional | Summary troubleshooting metric value for AP client count |
| `ApRtt` | `*float64` | Optional | Summary troubleshooting metric value for AP round-trip time |
| `AudioIn` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | Media-direction call troubleshooting metrics in a summary response |
| `AudioOut` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | Media-direction call troubleshooting metrics in a summary response |
| `ClientCpu` | `*float64` | Optional | Summary troubleshooting metric value for client CPU load |
| `ClientNStreams` | `*float64` | Optional | Summary troubleshooting metric value for the number of client spatial streams |
| `ClientRadioBand` | `*float64` | Optional | Summary troubleshooting metric value for the client radio band |
| `ClientRssi` | `*float64` | Optional | Summary troubleshooting metric value for client RSSI |
| `ClientRxBytes` | `*float64` | Optional | Summary troubleshooting metric value for bytes received by the client |
| `ClientRxRates` | `*float64` | Optional | Summary troubleshooting metric value for client receive data rates |
| `ClientRxRetries` | `*float64` | Optional | Summary troubleshooting metric value for client receive retries |
| `ClientTxBytes` | `*float64` | Optional | Summary troubleshooting metric value for bytes transmitted by the client |
| `ClientTxRates` | `*float64` | Optional | Summary troubleshooting metric value for client transmit data rates |
| `ClientTxRetries` | `*float64` | Optional | Summary troubleshooting metric value for client transmit retries |
| `ClientVpnDistance` | `*float64` | Optional | Summary troubleshooting metric value for client VPN distance |
| `ClientWifiVersion` | `*float64` | Optional | Summary troubleshooting metric value for the client Wi-Fi version |
| `Expected` | `*float64` | Optional | Model baseline value expected for this call sample |
| `RadioApChange` | `*float64` | Optional | Summary troubleshooting metric value for AP changes on the radio path |
| `RadioBandwidth` | `*float64` | Optional | Summary troubleshooting metric value for radio channel bandwidth |
| `RadioChannel` | `*float64` | Optional | Summary troubleshooting metric value for the radio channel |
| `RadioRxFailed` | `*float64` | Optional | Summary troubleshooting metric value for failed radio receive attempts |
| `RadioTxPower` | `*float64` | Optional | Summary troubleshooting metric value for radio transmit power |
| `RadioUtil` | `*float64` | Optional | Summary troubleshooting metric value for radio utilization |
| `RadioUtilInterference` | `*float64` | Optional | Summary troubleshooting metric value for radio interference utilization |
| `SiteNumClients` | `*float64` | Optional | Summary troubleshooting metric value for the number of clients at the site |
| `SiteWanAvgDownloadMbps` | `*float64` | Optional | Summary troubleshooting metric value for site WAN average download throughput |
| `SiteWanAvgUploadMbps` | `*float64` | Optional | Summary troubleshooting metric value for site WAN average upload throughput |
| `SiteWanDownloadMbps` | `*float64` | Optional | Summary troubleshooting metric value for site WAN download throughput |
| `SiteWanJitter` | `*float64` | Optional | Summary troubleshooting metric value for site WAN jitter |
| `SiteWanRtt` | `*float64` | Optional | Summary troubleshooting metric value for site WAN round-trip time |
| `SiteWanUploadMbps` | `*float64` | Optional | Summary troubleshooting metric value for site WAN upload throughput |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `VideoIn` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | Media-direction call troubleshooting metrics in a summary response |
| `VideoOut` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | Media-direction call troubleshooting metrics in a summary response |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    callTroubleshootSummary := models.CallTroubleshootSummary{
        ApNumClients:           models.ToPointer(float64(-0.08802365511655807)),
        ApRtt:                  models.ToPointer(float64(0.09924473613500595)),
        AudioIn:                models.ToPointer(models.CallTroubleshootSummaryData{
            ApNumClients:           models.ToPointer(float64(152.32)),
            ApRtt:                  models.ToPointer(float64(133.36)),
            ClientCpu:              models.ToPointer(float64(164.78)),
            ClientNStreams:         models.ToPointer(float64(206.36)),
            ClientRadioBand:        models.ToPointer(float64(43.4)),
        }),
        AudioOut:               models.ToPointer(models.CallTroubleshootSummaryData{
            ApNumClients:           models.ToPointer(float64(71.16)),
            ApRtt:                  models.ToPointer(float64(52.2)),
            ClientCpu:              models.ToPointer(float64(245.94)),
            ClientNStreams:         models.ToPointer(float64(125.2)),
            ClientRadioBand:        models.ToPointer(float64(218.24)),
        }),
        ClientCpu:              models.ToPointer(float64(0.00834270566701889)),
        ClientNStreams:         models.ToPointer(float64(0.00734270566701889)),
        ClientRadioBand:        models.ToPointer(float64(0.5841414928436279)),
        ClientRssi:             models.ToPointer(float64(0.7594696879386902)),
        ClientRxBytes:          models.ToPointer(float64(2.365511655807E-05)),
        ClientRxRates:          models.ToPointer(float64(0.02441493794322014)),
        ClientRxRetries:        models.ToPointer(float64(-0.14325742423534393)),
        ClientTxBytes:          models.ToPointer(float64(0.00102365511655807)),
        ClientTxRates:          models.ToPointer(float64(0.22236637771129608)),
        ClientTxRetries:        models.ToPointer(float64(0.3308201730251312)),
        ClientVpnDistance:      models.ToPointer(float64(-0.0001660545531194657)),
        ClientWifiVersion:      models.ToPointer(float64(7.0566701889E-07)),
        Expected:               models.ToPointer(float64(-2.8630001056670187)),
        RadioApChange:          models.ToPointer(float64(0.01850946433842182)),
        RadioBandwidth:         models.ToPointer(float64(-0.021175479516386986)),
        RadioChannel:           models.ToPointer(float64(0.11686426401138306)),
        RadioRxFailed:          models.ToPointer(float64(1.1782013177871704)),
        RadioTxPower:           models.ToPointer(float64(0.121039018034935)),
        RadioUtil:              models.ToPointer(float64(0.2452986091375351)),
        RadioUtilInterference:  models.ToPointer(float64(3.4367904663085938)),
        SiteNumClients:         models.ToPointer(float64(0.055026158690452576)),
        SiteWanAvgDownloadMbps: models.ToPointer(float64(3.0566701889E-07)),
        SiteWanAvgUploadMbps:   models.ToPointer(float64(5.566701889E-08)),
        SiteWanDownloadMbps:    models.ToPointer(float64(8.0566701889E-07)),
        SiteWanJitter:          models.ToPointer(float64(0.7875519659784105)),
        SiteWanRtt:             models.ToPointer(float64(15.094849904378256)),
        SiteWanUploadMbps:      models.ToPointer(float64(2.0566701889E-07)),
    }

}
```

