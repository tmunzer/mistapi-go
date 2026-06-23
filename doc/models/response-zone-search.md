
# Response Zone Search

Paginated response for site zone visit search results

## Structure

`ResponseZoneSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | Epoch timestamp, in seconds, for the end of the zone visit search window |
| `Limit` | `*int` | Optional | Maximum number of zone visit records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of zone visit search results |
| `Results` | [`[]models.ResponseZoneSearchItem`](../../doc/models/response-zone-search-item.md) | Optional | Zone visit records returned by search |
| `Start` | `*float64` | Optional | Epoch timestamp, in seconds, for the start of the zone visit search window |
| `Total` | `*int` | Optional | Number of zone visit records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseZoneSearch := models.ResponseZoneSearch{
        End:                  models.ToPointer(float64(1541705289.769911)),
        Limit:                models.ToPointer(1),
        Next:                 models.ToPointer("/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/zones/visits/search?limit=2&end=1541705247.000&scope_id=85fbba9e-4e12-11e6-9188-0242ac110007&user_type=asset&start=1541618889.77"),
        Results:              []models.ResponseZoneSearchItem{
            models.ResponseZoneSearchItem{
                Enter:                models.ToPointer(226),
                Scope:                models.ToPointer("scope6"),
                Timestamp:            models.ToPointer(float64(2.64)),
                User:                 models.ToPointer("user6"),
            },
        },
        Start:                models.ToPointer(float64(1541618889.769886)),
        Total:                models.ToPointer(5892),
    }

}
```

