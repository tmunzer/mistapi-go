
# Webhook Rssizone

Sample of the `rssizone` webhook payload.

## Structure

`WebhookRssizone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookRssizoneEvent`](../../doc/models/webhook-rssizone-event.md) | Required | RSSI zone transition events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for RSSI zone event deliveries. enum: `rssizone`<br><br>**Value**: `"rssizone"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookRssizone := models.WebhookRssizone{
        Events:               []models.WebhookRssizoneEvent{
            models.WebhookRssizoneEvent{
                Mac:                  "mac4",
                MapId:                uuid.MustParse("00001148-0000-0000-0000-000000000000"),
                RssizoneId:           uuid.MustParse("00000056-0000-0000-0000-000000000000"),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(188.18),
                Trigger:              models.WebhookZoneEventTriggerEnum_ENTER,
                Type:                 models.WebhookZoneEventTypeEnum_SDK,
            },
        },
        Topic:                "rssizone",
    }

}
```

