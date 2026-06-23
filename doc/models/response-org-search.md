
# Response Org Search

Paginated response for MSP organization search results

## Structure

`ResponseOrgSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required, Read-only | Epoch timestamp, in seconds, for the end of the organization search window |
| `Limit` | `int` | Required, Read-only | Maximum number of organization records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of organization search results |
| `Results` | [`[]models.ResponseOrgSearchItem`](../../doc/models/response-org-search-item.md) | Required | Organization records returned by MSP organization search<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required, Read-only | Epoch timestamp, in seconds, for the start of the organization search window |
| `Total` | `int` | Required, Read-only | Number of organization records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseOrgSearch := models.ResponseOrgSearch{
        End:                  float64(20.8),
        Limit:                138,
        Next:                 models.ToPointer("next0"),
        Results:              []models.ResponseOrgSearchItem{
            models.ResponseOrgSearchItem{
                MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
                Name:                 models.ToPointer("name6"),
                NumAps:               models.ToPointer(140),
                NumGateways:          models.ToPointer(112),
                NumSites:             models.ToPointer(50),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
            },
        },
        Start:                float64(232.86),
        Total:                232,
    }

}
```

