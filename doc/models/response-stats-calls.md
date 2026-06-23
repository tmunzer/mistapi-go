
# Response Stats Calls

Paginated response for site call statistics search results

## Structure

`ResponseStatsCalls`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | Epoch timestamp, in seconds, for the end of the call statistics search window |
| `Limit` | `*int` | Optional | Maximum number of call statistics records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of call statistics results |
| `Results` | [`[]models.StatsCall`](../../doc/models/stats-call.md) | Optional | Collaboration call statistics returned by search |
| `Start` | `*float64` | Optional | Epoch timestamp, in seconds, for the start of the call statistics search window |
| `Total` | `*int` | Optional | Number of call statistics records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseStatsCalls := models.ResponseStatsCalls{
        End:                  models.ToPointer(float64(80.26)),
        Limit:                models.ToPointer(80),
        Next:                 models.ToPointer("next6"),
        Results:              []models.StatsCall{
            models.StatsCall{
                App:                  models.ToPointer("app6"),
                AudioQuality:         models.ToPointer(66),
                EndTime:              models.ToPointer(186),
                Mac:                  models.ToPointer("mac0"),
                MeetingId:            models.ToPointer("meeting_id2"),
            },
            models.StatsCall{
                App:                  models.ToPointer("app6"),
                AudioQuality:         models.ToPointer(66),
                EndTime:              models.ToPointer(186),
                Mac:                  models.ToPointer("mac0"),
                MeetingId:            models.ToPointer("meeting_id2"),
            },
            models.StatsCall{
                App:                  models.ToPointer("app6"),
                AudioQuality:         models.ToPointer(66),
                EndTime:              models.ToPointer(186),
                Mac:                  models.ToPointer("mac0"),
                MeetingId:            models.ToPointer("meeting_id2"),
            },
        },
        Start:                models.ToPointer(float64(36.32)),
    }

}
```

