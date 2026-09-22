
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
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Message:              "",
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                Timestamp:            0.0,
            },
        },
        Topic:                "audits",
    }

}
```

