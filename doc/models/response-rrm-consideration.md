
# Response Rrm Consideration

Response containing current RRM considerations for an AP radio band

## Structure

`ResponseRrmConsideration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.RrmConsideration`](../../doc/models/rrm-consideration.md) | Required | RRM channel consideration records returned for an AP radio band<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseRrmConsideration := models.ResponseRrmConsideration{
        Results:              []models.RrmConsideration{
            models.RrmConsideration{
                Channel:              30,
                Noise:                float64(44.34),
                OtherRssi:            models.ToPointer(float64(33.36)),
                OtherSsid:            models.ToPointer("other_ssid6"),
                Rssi:                 models.ToPointer(float64(128.68)),
                UtilScore:            float64(76.78),
                UtilScoreNonWifi:     float64(74.88),
                UtilScoreOther:       float64(108.4),
            },
        },
    }

}
```

