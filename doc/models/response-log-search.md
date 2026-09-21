
# Response Log Search

Paginated response for audit log search results

## Structure

`ResponseLogSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the audit log search window |
| `Limit` | `int` | Required | Maximum number of audit log events returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of audit log events |
| `Results` | [`[]models.LogEvent`](../../doc/models/log-event.md) | Required | Audit log events returned by a log query<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the audit log search window |
| `Total` | `int` | Required | Number of audit log events matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseLogSearch := models.ResponseLogSearch{
        End:                  180,
        Limit:                246,
        Next:                 models.ToPointer("next2"),
        Results:              []models.LogEvent{
            models.LogEvent{
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Message:              "",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Timestamp:            0.0,
            },
        },
        Start:                138,
        Total:                84,
    }

}
```

