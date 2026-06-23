
# Response Self Audit Logs

Paginated response for audit logs for the current admin across organizations

## Structure

`ResponseSelfAuditLogs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the self audit log search window |
| `Limit` | `int` | Required | Maximum number of self audit log entries returned in this page |
| `Page` | `int` | Required | Current page number returned for self audit log results |
| `Results` | [`[]models.AuditLog`](../../doc/models/audit-log.md) | Required | Audit log entries returned for the current admin<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the self audit log search window |
| `Total` | `int` | Required | Number of self audit log entries matching the filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSelfAuditLogs := models.ResponseSelfAuditLogs{
        End:                  216,
        Limit:                210,
        Page:                 68,
        Results:              []models.AuditLog{
            models.AuditLog{
                AdminId:              uuid.MustParse("456b7016-a916-a4b1-78dd-72b947c152b7"),
                AdminName:            "admin_name4",
                After:                models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                Before:               models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                ForSite:              models.ToPointer(false),
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                Message:              "message6",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(2.64),
            },
        },
        Start:                174,
        Total:                48,
    }

}
```

