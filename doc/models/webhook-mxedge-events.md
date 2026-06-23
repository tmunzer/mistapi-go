
# Webhook Mxedge Events

Sample of the `mxedge-events` webhook payload.

## Structure

`WebhookMxedgeEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.MxedgeEvent`](../../doc/models/mxedge-event.md) | Required | Mist Edge event records<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for Mist Edge event deliveries. enum: `mxedge-events`<br><br>**Value**: `"mxedge-events"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookMxedgeEvents := models.WebhookMxedgeEvents{
        Events:               []models.MxedgeEvent{
            models.MxedgeEvent{
                AuditId:              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                Component:            models.NewOptional(models.ToPointer("PS1")),
                DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("0000254a-0000-0000-0000-000000000000"))),
                DeviceType:           models.ToPointer("device_type0"),
                FromVersion:          models.ToPointer("from_version2"),
                MxclusterId:          models.ToPointer("2815c917-58e7-472f-a190-bfd44fb58d05"),
                MxedgeId:             models.ToPointer("00000000-0000-0000-1000-020000dc585c"),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                Service:              models.ToPointer("tunterm"),
                Type:                 models.ToPointer("ME_SERVICE_STOPPED"),
            },
        },
        Topic:                "mxedge-events",
    }

}
```

