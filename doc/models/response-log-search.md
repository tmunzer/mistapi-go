
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
                AdminId:              models.NewOptional(models.ToPointer(uuid.MustParse("00000014-0000-0000-0000-000000000000"))),
                AdminName:            models.NewOptional(models.ToPointer("admin_name4")),
                After:                models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                Before:               models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000"))),
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Message:              "message6",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Timestamp:            float64(2.64),
            },
        },
        Start:                138,
        Total:                84,
    }

}
```

