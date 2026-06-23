
# Response Search Var

Paginated response for organization variable search results

## Structure

`ResponseSearchVar`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the variable search window |
| `Limit` | `*int` | Optional | Maximum number of variable records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of variable search results |
| `Results` | [`[]models.ResponseSearchVarItem`](../../doc/models/response-search-var-item.md) | Optional | Organization variable records returned by search |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the variable search window |
| `Total` | `*int` | Optional | Number of variable records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSearchVar := models.ResponseSearchVar{
        End:                  models.ToPointer(218),
        Limit:                models.ToPointer(208),
        Next:                 models.ToPointer("next4"),
        Results:              []models.ResponseSearchVarItem{
            models.ResponseSearchVarItem{
                CreatedTime:          models.ToPointer(float64(73.76)),
                ModifiedTime:         models.ToPointer(float64(5.2)),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
                Src:                  models.ToPointer("src8"),
            },
            models.ResponseSearchVarItem{
                CreatedTime:          models.ToPointer(float64(73.76)),
                ModifiedTime:         models.ToPointer(float64(5.2)),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
                Src:                  models.ToPointer("src8"),
            },
            models.ResponseSearchVarItem{
                CreatedTime:          models.ToPointer(float64(73.76)),
                ModifiedTime:         models.ToPointer(float64(5.2)),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
                Src:                  models.ToPointer("src8"),
            },
        },
        Start:                models.ToPointer(176),
    }

}
```

