
# Response Events Fastroam

Paginated response for fast roaming event results

## Structure

`ResponseEventsFastroam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the roaming event search window |
| `Limit` | `int` | Required | Maximum number of roaming event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of roaming event records; null when no next page exists |
| `Results` | [`[]models.EventFastroam`](../../doc/models/event-fastroam.md) | Required | Fast roaming event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the roaming event search window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseEventsFastroam := models.ResponseEventsFastroam{
        End:                  16,
        Limit:                154,
        Next:                 models.ToPointer("next8"),
        Results:              []models.EventFastroam{
            models.EventFastroam{
                ApMac:                "ap_mac8",
                ClientMac:            "client_mac4",
                Fromap:               "fromap6",
                Latency:              float64(250.14),
                Ssid:                 "ssid6",
                Subtype:              models.ToPointer("subtype8"),
                Timestamp:            float64(2.64),
                Type:                 models.ToPointer(models.EventFastroamTypeEnum_PINGPONG),
            },
        },
        Start:                230,
    }

}
```

