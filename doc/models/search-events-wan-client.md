
# Search Events Wan Client

Paginated response for WAN client event searches

## Structure

`SearchEventsWanClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound timestamp of the WAN client event search window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of WAN client event results returned by this page |
| `Next` | `*string` | Optional | URL for the next page of WAN client event results, when more results are available |
| `Results` | [`*models.EventsClientWan`](../../doc/models/events-client-wan.md) | Optional | WAN client event returned by WAN client event search APIs |
| `Start` | `*int` | Optional | Lower bound timestamp of the WAN client event search window, in epoch seconds |
| `Total` | `*int` | Optional | Count of WAN client event results matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    searchEventsWanClient := models.SearchEventsWanClient{
        End:                  models.ToPointer(12),
        Limit:                models.ToPointer(158),
        Next:                 models.ToPointer("next0"),
        Results:              models.ToPointer(models.EventsClientWan{
            When:                 models.ToPointer("When8"),
            EvType:               models.ToPointer("ev_type4"),
            Metadata:             models.ToPointer(interface{}("[key1, val1][key2, val2]")),
            OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
            RandomMac:            models.ToPointer(false),
        }),
        Start:                models.ToPointer(226),
    }

}
```

