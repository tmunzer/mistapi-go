
# Response Events Sky Atp Search

Paginated response for Sky ATP event search results

## Structure

`ResponseEventsSkyAtpSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the Sky ATP event search window |
| `Limit` | `int` | Required | Maximum number of Sky ATP event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of Sky ATP event records |
| `Results` | [`[]models.EventsSkyatp`](../../doc/models/events-skyatp.md) | Required | Sky ATP event records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the Sky ATP event search window |
| `Total` | `int` | Required | Number of Sky ATP event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseEventsSkyAtpSearch := models.ResponseEventsSkyAtpSearch{
        End:                  220,
        Limit:                50,
        Next:                 models.ToPointer("next0"),
        Results:              []models.EventsSkyatp{
            models.EventsSkyatp{
                DeviceMac:            "device_mac0",
                ForSite:              models.ToPointer(false),
                Ip:                   "ip0",
                Mac:                  "mac0",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                ThreatLevel:          152,
                Timestamp:            float64(2.64),
                Type:                 "type4",
            },
        },
        Start:                178,
        Total:                212,
    }

}
```

