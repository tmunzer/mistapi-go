
# Webhook Audits

Sample of the `audits` webhook payload.

## Structure

`WebhookAudits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.LogEvent`](../../doc/models/log-event.md) | Required | Audit log events returned by a log query<br><br>**Constraints**: *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for audit event deliveries. enum: `audits`<br><br>**Value**: `"audits"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookAudits := models.WebhookAudits{
        Events:               []models.LogEvent{
            models.LogEvent{
                AdminId:              models.NewOptional(models.ToPointer(uuid.MustParse("0000104e-0000-0000-0000-000000000000"))),
                AdminName:            models.NewOptional(models.ToPointer("admin_name8")),
                After:                models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                Before:               models.ToPointer(interface{}("[key1, val1][key2, val2]")),
                DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("0000254a-0000-0000-0000-000000000000"))),
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Message:              "message0",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Timestamp:            float64(188.18),
            },
        },
        Topic:                "audits",
    }

}
```

