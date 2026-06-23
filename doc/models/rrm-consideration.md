
# Rrm Consideration

RRM consideration metrics for one candidate channel

## Structure

`RrmConsideration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `int` | Required | Candidate RF channel evaluated by RRM |
| `Noise` | `float64` | Required | Observed noise floor for the candidate channel, in dBm |
| `OtherRssi` | `*float64` | Optional | Avg RSSI heard from other APs (that does NOT belongs to the same site) |
| `OtherSsid` | `*string` | Optional | SSID from other AP that we heard from with the max RSSI |
| `Rssi` | `*float64` | Optional | Avg RSSI heard from APs (that belongs to the same site) |
| `UtilScore` | `float64` | Required | utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreNonWifi` | `float64` | Required | non-Wi-Fi utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreOther` | `float64` | Required | other utilization score, 0-1, lower means less utilization (cleaner RF) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmConsideration := models.RrmConsideration{
        Channel:              104,
        Noise:                float64(47.08),
        OtherRssi:            models.ToPointer(float64(124.78)),
        OtherSsid:            models.ToPointer("other_ssid8"),
        Rssi:                 models.ToPointer(float64(220.1)),
        UtilScore:            float64(168.2),
        UtilScoreNonWifi:     float64(239.46),
        UtilScoreOther:       float64(199.82),
    }

}
```

