
# Response Events Search

Paginated response for wireless client event search results

## Structure

`ResponseEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the wireless client event search window |
| `Limit` | `int` | Required | Maximum number of wireless client event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of wireless client event records |
| `Results` | [`[]models.EventsClient`](../../doc/models/events-client.md) | Required | Wireless client event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the wireless client event search window |
| `Total` | `int` | Required | Number of wireless client event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsSearch := models.ResponseEventsSearch{
        End:                  144,
        Limit:                230,
        Next:                 models.ToPointer("next0"),
        Results:              []models.EventsClient{
            models.EventsClient{
                Ap:                   models.ToPointer("ap8"),
                Band:                 models.Dot11BandEnum_ENUM5,
                Bssid:                models.ToPointer("bssid0"),
                Channel:              models.ToPointer(30),
                KeyMgmt:              models.ToPointer(models.ClientKeyMgmtEnum_WPA2PSK),
                Proto:                models.ToPointer(models.Dot11ProtoEnum_B),
                Timestamp:            float64(2.64),
            },
        },
        Start:                102,
        Total:                136,
    }

}
```

