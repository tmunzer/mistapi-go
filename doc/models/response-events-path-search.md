
# Response Events Path Search

Paginated response for service path event search results

## Structure

`ResponseEventsPathSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the service path event search window |
| `Limit` | `*int` | Optional | Maximum number of service path event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of service path event records |
| `Results` | [`[]models.ServicePathEvent`](../../doc/models/service-path-event.md) | Optional | Service path event records returned by a search response |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the service path event search window |
| `Total` | `*int` | Optional | Number of service path event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseEventsPathSearch := models.ResponseEventsPathSearch{
        End:                  models.ToPointer(1697096379),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next6"),
        Results:              []models.ServicePathEvent{
            models.ServicePathEvent{
                Mac:                  models.ToPointer("mac0"),
                Model:                models.ToPointer("model4"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                Policy:               models.ToPointer("policy8"),
                PortId:               models.ToPointer("port_id6"),
            },
            models.ServicePathEvent{
                Mac:                  models.ToPointer("mac0"),
                Model:                models.ToPointer("model4"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                Policy:               models.ToPointer("policy8"),
                PortId:               models.ToPointer("port_id6"),
            },
        },
        Start:                models.ToPointer(1697009979),
        Total:                models.ToPointer(2),
    }

}
```

