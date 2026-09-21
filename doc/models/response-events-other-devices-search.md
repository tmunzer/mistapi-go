
# Response Events Other Devices Search

Paginated response for other-device event search results

## Structure

`ResponseEventsOtherDevicesSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the other-device event search window |
| `Limit` | `*int` | Optional | Maximum number of other-device event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of other-device event records |
| `Results` | [`[]models.EventOtherdevice`](../../doc/models/event-otherdevice.md) | Optional | Other-device event records returned by a search response |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the other-device event search window |
| `Total` | `*int` | Optional | Number of other-device event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsOtherDevicesSearch := models.ResponseEventsOtherDevicesSearch{
        End:                  models.ToPointer(52),
        Limit:                models.ToPointer(118),
        Next:                 models.ToPointer("next2"),
        Results:              []models.EventOtherdevice{
            models.EventOtherdevice{
                DeviceMac:            models.ToPointer("device_mac0"),
                Mac:                  models.ToPointer("mac0"),
                Text:                 models.ToPointer("text4"),
            },
            models.EventOtherdevice{
                DeviceMac:            models.ToPointer("device_mac0"),
                Mac:                  models.ToPointer("mac0"),
                Text:                 models.ToPointer("text4"),
            },
            models.EventOtherdevice{
                DeviceMac:            models.ToPointer("device_mac0"),
                Mac:                  models.ToPointer("mac0"),
                Text:                 models.ToPointer("text4"),
            },
        },
        Start:                models.ToPointer(10),
    }

}
```

