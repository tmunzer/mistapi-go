
# Webhook Alarms

Sample of the `alarms` webhook payload.

**N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.

Events specific fields for other alarm event type can be found with API [List Alarm Definitions#]($e/Events%20Definitions/listAlarmDefinitions), under "fields" array of /alarm_defs response object.

## Structure

`WebhookAlarms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookAlarmEvent`](../../doc/models/webhook-alarm-event.md) | Required | Alarm events included in this webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for alarm event deliveries. enum: `alarms`<br><br>**Value**: `"alarms"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookAlarms := models.WebhookAlarms{
        Events:               []models.WebhookAlarmEvent{
            models.WebhookAlarmEvent{
                Aps:                  []string{
                    "aps1",
                    "aps2",
                },
                Bssids:               []string{
                    "bssids2",
                },
                Count:                models.ToPointer(152),
                EventId:              models.ToPointer(uuid.MustParse("000015dc-0000-0000-0000-000000000000")),
                ForSite:              models.ToPointer(false),
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
                OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(188.18),
                Type:                 "type0",
            },
        },
        Topic:                "alarms",
    }

}
```

