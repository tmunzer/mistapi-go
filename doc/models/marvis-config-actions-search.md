
# Marvis Config Actions Search

Paginated list of Marvis config actions

## Structure

`MarvisConfigActionsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Search window end timestamp, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of results requested |
| `Results` | [`[]models.MarvisConfigAction`](../../doc/models/marvis-config-action.md) | Optional | List of Marvis config actions |
| `Start` | `*int` | Optional | Search window start timestamp, in epoch seconds |
| `Total` | `*int` | Optional | Total number of matching results |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    marvisConfigActionsSearch := models.MarvisConfigActionsSearch{
        End:                  models.ToPointer(240),
        Limit:                models.ToPointer(186),
        Results:              []models.MarvisConfigAction{
            models.MarvisConfigAction{
                AdminId:              models.ToPointer(uuid.MustParse("00000014-0000-0000-0000-000000000000")),
                Id:                   models.ToPointer(uuid.MustParse("000023ba-0000-0000-0000-000000000000")),
                Mac:                  models.ToPointer("mac0"),
                Op:                   models.ToPointer("op0"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
            },
            models.MarvisConfigAction{
                AdminId:              models.ToPointer(uuid.MustParse("00000014-0000-0000-0000-000000000000")),
                Id:                   models.ToPointer(uuid.MustParse("000023ba-0000-0000-0000-000000000000")),
                Mac:                  models.ToPointer("mac0"),
                Op:                   models.ToPointer("op0"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
            },
        },
        Start:                models.ToPointer(198),
        Total:                models.ToPointer(24),
    }

}
```

