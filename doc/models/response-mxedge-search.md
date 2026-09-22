
# Response Mxedge Search

Search response for Mist Edge records

## Structure

`ResponseMxedgeSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the Mist Edge search window |
| `Limit` | `*int` | Optional | Maximum number of Mist Edge records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of Mist Edge records |
| `Results` | [`[]models.SearchMxedge`](../../doc/models/search-mxedge.md) | Optional | Mist Edge search records returned by a search response |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the Mist Edge search window |
| `Total` | `*int` | Optional | Number of Mist Edge records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseMxedgeSearch := models.ResponseMxedgeSearch{
        End:                  models.ToPointer(1694708579),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next4"),
        Results:              []models.SearchMxedge{
            models.SearchMxedge{
                Distro:               models.ToPointer("distro8"),
                LastSeen:             models.ToPointer(float64(165.16)),
                Model:                models.ToPointer("model4"),
            },
        },
        Start:                models.ToPointer(1694622179),
        Total:                models.ToPointer(2),
    }

}
```

