
# Rrm Channel Score

RRM utilization score for a channel

## Structure

`RrmChannelScore`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `int` | Required | RF channel number represented by this score record |
| `UtilScore` | `float64` | Required | Utilization score for the channel, 0-1, lower means cleaner RF |
| `UtilScoreNoiseFloor` | `float64` | Required | Score contribution from noise, 0-1, lower means cleaner RF |
| `UtilScoreNonWifi` | `float64` | Required | Score contribution from non-wifi utilization, 0-1, lower means cleaner RF |
| `UtilScoreOther` | `float64` | Required | Score contribution from RxOtherBss utilization (wifi packets destined for other radios), 0-1, lower means cleaner RF |
| `UtilScoreRadar` | `float64` | Required | Score contribution from radar detections, 0-1, lower means cleaner RF |
| `UtilScoreUndecodableWifi` | `float64` | Required | Score contribution from undecodable wifi utilization (wifi packets which can't be decoded), 0-1, lower means cleaner RF |
| `UtilScoreUnknownWifi` | `float64` | Required | Score contribution from unknown wifi utilization (wifi packets of unknown type), 0-1, lower means cleaner RF |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmChannelScore := models.RrmChannelScore{
        Channel:                  58,
        UtilScore:                float64(250.58),
        UtilScoreNoiseFloor:      float64(58.22),
        UtilScoreNonWifi:         float64(98.92),
        UtilScoreOther:           float64(26.2),
        UtilScoreRadar:           float64(32.9),
        UtilScoreUndecodableWifi: float64(94),
        UtilScoreUnknownWifi:     float64(201.04),
    }

}
```

