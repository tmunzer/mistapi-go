
# Response Client Sessions Search

Paginated wireless client session search response

## Structure

`ResponseClientSessionsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Search window end timestamp for wireless client sessions, in epoch seconds |
| `Limit` | `int` | Required | Maximum number of wireless client session results requested |
| `Next` | `*string` | Optional | URL for the next page of wireless client session results |
| `Results` | [`[]models.ResponseClientSessionsSearchItem`](../../doc/models/response-client-sessions-search-item.md) | Required | Wireless client session records returned by a search response |
| `Start` | `int` | Required | Search window start timestamp for wireless client sessions, in epoch seconds |
| `Total` | `int` | Required | Number of wireless client session records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseClientSessionsSearch := models.ResponseClientSessionsSearch{
        End:                  126,
        Limit:                212,
        Next:                 models.ToPointer("next4"),
        Results:              []models.ResponseClientSessionsSearchItem{
            models.ResponseClientSessionsSearchItem{
                Ap:                   "ap8",
                Band:                 "band8",
                ClientManufacture:    "client_manufacture0",
                Connect:              float64(206.78),
                Disconnect:           float64(129.48),
                Duration:             float64(104.42),
                Mac:                  "mac0",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Ssid:                 "ssid6",
                Tags:                 []string{
                    "tags1",
                    "tags2",
                },
                Timestamp:            float64(2.64),
                WlanId:               uuid.MustParse("00000742-0000-0000-0000-000000000000"),
            },
        },
        Start:                84,
        Total:                118,
    }

}
```

