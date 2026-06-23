
# Response Search

Paginated response for MSP organization search results

## Structure

`ResponseSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `int` | Required | Maximum number of MSP search results returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of MSP search results |
| `Page` | `int` | Required | Current page number returned for the MSP search results |
| `Results` | [`[]models.ResponseSearchItem`](../../doc/models/response-search-item.md) | Required | MSP search result entries returned by search<br><br>**Constraints**: *Unique Items Required* |
| `Total` | `int` | Required | Number of MSP search results matching the query across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSearch := models.ResponseSearch{
        Limit:                110,
        Next:                 models.ToPointer("next8"),
        Page:                 224,
        Results:              []models.ResponseSearchItem{
            models.ResponseSearchItem{
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                Text:                 "text4",
                Type:                 "type4",
            },
        },
        Total:                204,
    }

}
```

