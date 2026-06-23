
# Response Events Rrm

Paginated response for RRM event results

## Structure

`ResponseEventsRrm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the RRM event search window |
| `Limit` | `int` | Required | Maximum number of RRM event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of RRM event records; null when no next page exists |
| `Results` | [`[]models.RrmEvent`](../../doc/models/rrm-event.md) | Required | RRM event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the RRM event search window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsRrm := models.ResponseEventsRrm{
        End:                  202,
        Limit:                32,
        Next:                 models.ToPointer("next4"),
        Results:              []models.RrmEvent{
            models.RrmEvent{
                Ap:                   "5c5b350e0001",
                Band:                 models.Dot11BandEnum_ENUM5,
                Bandwidth:            models.Dot11BandwidthEnum_ENUM20,
                Channel:              30,
                Event:                models.RrmEventTypeEnum_RADARDETECTED,
                Power:                226,
                PreBandwidth:         models.RrmEventPreBandwidthEnum_ENUM80,
                PreChannel:           164,
                PrePower:             float64(92.2),
                PreUsage:             "pre_usage6",
                Timestamp:            float64(2.64),
                Usage:                "usage4",
            },
        },
        Start:                160,
    }

}
```

