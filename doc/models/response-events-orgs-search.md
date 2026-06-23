
# Response Events Orgs Search

Paginated response for organization event search results

## Structure

`ResponseEventsOrgsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp for the end of the organization event search window |
| `Limit` | `*int` | Optional | Maximum number of organization event records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of organization event records |
| `Results` | [`[]models.OrgEvent`](../../doc/models/org-event.md) | Optional | Organization event records returned by a search response |
| `Start` | `*int` | Optional | Epoch timestamp for the start of the organization event search window |
| `Total` | `*int` | Optional | Number of organization event records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseEventsOrgsSearch := models.ResponseEventsOrgsSearch{
        End:                  models.ToPointer(1688035193),
        Limit:                models.ToPointer(10),
        Next:                 models.ToPointer("next4"),
        Results:              []models.OrgEvent{
            models.OrgEvent{
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
                Text:                 models.ToPointer("text4"),
                Timestamp:            models.ToPointer(float64(2.64)),
                Type:                 models.ToPointer("type4"),
            },
        },
        Start:                models.ToPointer(1687948793),
    }

}
```

