
# Response Events Rogue Search

Paginated response for rogue AP event search results

## Structure

`ResponseEventsRogueSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the rogue AP event search window |
| `Limit` | `int` | Required | Maximum number of rogue AP event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of rogue AP event records |
| `Results` | [`[]models.EventsRogue`](../../doc/models/events-rogue.md) | Required | Rogue AP event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the rogue AP event search window |
| `Total` | `int` | Required | Number of rogue AP event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsRogueSearch := models.ResponseEventsRogueSearch{
        End:                  150,
        Limit:                20,
        Next:                 models.ToPointer("next0"),
        Results:              []models.EventsRogue{
            models.EventsRogue{
                Ap:                   "ap8",
                Bssid:                "bssid0",
                Channel:              30,
                Rssi:                 68,
                Ssid:                 "ssid6",
                Timestamp:            float64(2.64),
            },
        },
        Start:                108,
        Total:                114,
    }

}
```

