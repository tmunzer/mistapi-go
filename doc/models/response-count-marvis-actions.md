
# Response Count Marvis Actions

Distinct count response for Marvis action suggestions

## Structure

`ResponseCountMarvisActions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distinct` | `*string` | Optional | Field used to group Marvis action count results |
| `Limit` | `*int` | Optional | Maximum number of Marvis action count results requested |
| `Results` | [`[]models.ResponseCountMarvisActionsResult`](../../doc/models/response-count-marvis-actions-result.md) | Optional | Marvis action count results grouped by a distinct field |
| `Total` | `*int` | Optional | Number of Marvis action count result buckets returned |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseCountMarvisActions := models.ResponseCountMarvisActions{
        Distinct:             models.ToPointer("status"),
        Limit:                models.ToPointer(1000),
        Results:              []models.ResponseCountMarvisActionsResult{
            models.ResponseCountMarvisActionsResult{
                Count:                models.ToPointer(24),
                AdditionalProperties: map[string]string{
                    "status": "002e176a-0000-000-1111-002e208b20e1",
                },
            },
            models.ResponseCountMarvisActionsResult{
                Count:                models.ToPointer(12),
                AdditionalProperties: map[string]string{
                    "status": "2d3f176a-0000-000-2222-002e208f176a",
                },
            },
            models.ResponseCountMarvisActionsResult{
                Count:                models.ToPointer(15),
                AdditionalProperties: map[string]string{
                    "status": "08b2176a-0000-000-3333-002e208b2d3f",
                },
            },
        },
        Total:                models.ToPointer(3),
    }

}
```

