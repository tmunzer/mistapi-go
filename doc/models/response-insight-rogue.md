
# Response Insight Rogue

Paginated response for rogue or neighbor AP insight results

## Structure

`ResponseInsightRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the rogue AP insight window |
| `Limit` | `int` | Required | Maximum number of rogue or neighbor AP observations returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of rogue or neighbor AP observations; null when no next page exists |
| `Results` | [`[]models.InsightRogueAp`](../../doc/models/insight-rogue-ap.md) | Required | Rogue or neighbor AP insight observations returned by a query<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the rogue AP insight window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseInsightRogue := models.ResponseInsightRogue{
        End:                  176,
        Limit:                250,
        Next:                 models.ToPointer("next6"),
        Results:              []models.InsightRogueAp{
            models.InsightRogueAp{
                ApMac:                "ap_mac8",
                AvgRssi:              float64(170.84),
                Bssid:                "bssid0",
                Channel:              "channel8",
                DeltaX:               models.ToPointer(float64(25.74)),
                DeltaY:               models.ToPointer(float64(53.12)),
                NumAps:               140,
                SeenOnLan:            models.ToPointer(false),
                Ssid:                 models.ToPointer("ssid6"),
                TimesHeard:           models.ToPointer(110),
            },
        },
        Start:                134,
    }

}
```

