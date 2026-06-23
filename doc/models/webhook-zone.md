
# Webhook Zone

Sample of the `zone` webhook payload.

## Structure

`WebhookZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookZoneEvent`](../../doc/models/webhook-zone-event.md) | Required | Zone transition events included in this webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for zone transition deliveries. enum: `zone`<br><br>**Value**: `"zone"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookZone := models.WebhookZone{
        Events:               []models.WebhookZoneEvent{
            models.WebhookZoneEvent{
                AssetId:              models.ToPointer(uuid.MustParse("00001e56-0000-0000-0000-000000000000")),
                Id:                   models.ToPointer(uuid.MustParse("00000ce4-0000-0000-0000-000000000000")),
                Mac:                  models.ToPointer("mac4"),
                MapId:                uuid.MustParse("00001148-0000-0000-0000-000000000000"),
                Name:                 models.ToPointer("name0"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(188.18),
                Trigger:              models.WebhookZoneEventTriggerEnum_ENTER,
                Type:                 models.WebhookZoneEventTypeEnum_SDK,
                ZoneId:               uuid.MustParse("00000f60-0000-0000-0000-000000000000"),
            },
        },
        Topic:                "zone",
    }

}
```

