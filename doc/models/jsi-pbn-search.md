
# Jsi Pbn Search

Juniper Security Intelligence PBN search response with result metadata

## Structure

`JsiPbnSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Upper bound timestamp for the PBN search window |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | Pagination URL for the next page of PBN advisories |
| `Results` | [`[]models.JsiPbnItem`](../../doc/models/jsi-pbn-item.md) | Optional | List of PBN advisories |
| `Start` | `*int` | Optional | Lower bound timestamp for the PBN search window |
| `Total` | `*int` | Optional | Count of PBN advisories matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiPbnSearch := models.JsiPbnSearch{
        End:                  models.ToPointer(30),
        Limit:                models.ToPointer(140),
        Next:                 models.ToPointer("next4"),
        Results:              []models.JsiPbnItem{
            models.JsiPbnItem{
                BugType:              models.ToPointer("bug_type0"),
                CustomerRisk:         models.ToPointer("customer_risk4"),
                FixedIn:              models.ToPointer("fixed_in8"),
                Id:                   models.ToPointer("id6"),
                IntroducedIn:         models.ToPointer("introduced_in2"),
            },
            models.JsiPbnItem{
                BugType:              models.ToPointer("bug_type0"),
                CustomerRisk:         models.ToPointer("customer_risk4"),
                FixedIn:              models.ToPointer("fixed_in8"),
                Id:                   models.ToPointer("id6"),
                IntroducedIn:         models.ToPointer("introduced_in2"),
            },
        },
        Start:                models.ToPointer(244),
    }

}
```

