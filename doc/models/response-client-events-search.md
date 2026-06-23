
# Response Client Events Search

Paginated client event search response

## Structure

`ResponseClientEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Search window end timestamp for client events, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of client event results requested |
| `Next` | `*string` | Optional | URL for the next page of client event results |
| `Results` | [`[]models.EventsClient`](../../doc/models/events-client.md) | Required | Client event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Search window start timestamp for client events, in epoch seconds |
| `Total` | `int` | Required | Number of client event records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClientEventsSearch := models.ResponseClientEventsSearch{
        End:                  222,
        Limit:                52,
        Next:                 models.ToPointer("next2"),
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
        Start:                180,
        Total:                214,
    }

}
```

