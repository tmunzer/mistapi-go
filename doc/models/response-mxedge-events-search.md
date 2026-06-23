
# Response Mxedge Events Search

Search response for Mist Edge event records

## Structure

`ResponseMxedgeEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the Mist Edge event search window |
| `Limit` | `*int` | Optional | Maximum number of Mist Edge event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of Mist Edge event records |
| `Page` | `*int` | Optional | Returned page number for Mist Edge event records |
| `Results` | [`[]models.MxedgeEvent`](../../doc/models/mxedge-event.md) | Optional | Mist Edge event records<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the Mist Edge event search window |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMxedgeEventsSearch := models.ResponseMxedgeEventsSearch{
        End:                  models.ToPointer(1694708579),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next4"),
        Page:                 models.ToPointer(3),
        Results:              []models.MxedgeEvent{
            models.MxedgeEvent{
                AuditId:              models.ToPointer(uuid.MustParse("00000d00-0000-0000-0000-000000000000")),
                Component:            models.NewOptional(models.ToPointer("component4")),
                DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000"))),
                DeviceType:           models.ToPointer("device_type4"),
                FromVersion:          models.ToPointer("from_version8"),
            },
            models.MxedgeEvent{
                AuditId:              models.ToPointer(uuid.MustParse("00000d00-0000-0000-0000-000000000000")),
                Component:            models.NewOptional(models.ToPointer("component4")),
                DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000"))),
                DeviceType:           models.ToPointer("device_type4"),
                FromVersion:          models.ToPointer("from_version8"),
            },
        },
        Start:                models.ToPointer(1694622179),
    }

}
```

