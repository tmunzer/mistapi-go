
# Response Rrm Channel Scores

Response containing RRM channel score records

## Structure

`ResponseRrmChannelScores`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.RrmChannelScore`](../../doc/models/rrm-channel-score.md) | Required | RRM channel score records returned for a band<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseRrmChannelScores := models.ResponseRrmChannelScores{
        Results:              []models.RrmChannelScore{
            models.RrmChannelScore{
                Channel:                  30,
                UtilScore:                float64(76.78),
                UtilScoreNoiseFloor:      float64(140.42),
                UtilScoreNonWifi:         float64(74.88),
                UtilScoreOther:           float64(108.4),
                UtilScoreRadar:           float64(115.1),
                UtilScoreUndecodableWifi: float64(176.2),
                UtilScoreUnknownWifi:     float64(27.24),
            },
        },
    }

}
```

