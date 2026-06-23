
# Response Org System Events Search

Paginated response for organization system event search results

## Structure

`ResponseOrgSystemEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the system event search window |
| `Limit` | `int` | Required | Maximum number of system event records returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of system event search results |
| `Results` | [`[]models.OrgSystemEvent`](../../doc/models/org-system-event.md) | Required | Organization system event records returned by search<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the system event search window |
| `Total` | `int` | Required | Number of system event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseOrgSystemEventsSearch := models.ResponseOrgSystemEventsSearch{
        End:                  174,
        Limit:                252,
        Next:                 models.ToPointer("next6"),
        Results:              []models.OrgSystemEvent{
            models.OrgSystemEvent{
                ChangeCat:            models.ToPointer("admin_action"),
                Metadata:             models.ToPointer("metadata0"),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                Scope:                models.ToPointer("org"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Type:                 models.ToPointer("delete-wlan"),
            },
        },
        Start:                132,
        Total:                166,
    }

}
```

