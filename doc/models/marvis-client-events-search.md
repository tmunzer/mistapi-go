
# Marvis Client Events Search

Paginated list of Marvis Client events

## Structure

`MarvisClientEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | Maximum number of results requested |
| `Results` | [`[]models.MarvisClientEvent`](../../doc/models/marvis-client-event.md) | Optional | List of Marvis Client events |
| `Total` | `*int` | Optional | Total number of matching results |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    marvisClientEventsSearch := models.MarvisClientEventsSearch{
        Limit:                models.ToPointer(2),
        Results:              []models.MarvisClientEvent{
            models.MarvisClientEvent{
                Band:                 models.ToPointer("band8"),
                Bssid:                models.ToPointer("bssid0"),
                Channel:              models.ToPointer(30),
                DeviceId:             models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000")),
                Hostname:             models.ToPointer("hostname8"),
            },
        },
        Total:                models.ToPointer(164),
    }

}
```

