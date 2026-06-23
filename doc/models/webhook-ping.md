
# Webhook Ping

Sample of the `ping` webhook payload.\n\nThe `ping` webhook can be manually sent with the following API calls:\n- for a Site level webhook with the [Ping Site Webhook](../../doc/controllers/orgs-webhooks.md#ping-org-webhook) endpoint\n- for an Org level webhook with the [Ping Org Webhook](../../doc/controllers/orgs-webhooks.md#ping-org-webhook) endpoint

## Structure

`WebhookPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookPingEvent`](../../doc/models/webhook-ping-event.md) | Required | Ping events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for ping deliveries. enum: `ping`<br><br>**Value**: `"ping"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookPing := models.WebhookPing{
        Events:               []models.WebhookPingEvent{
            models.WebhookPingEvent{
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                Name:                 "name0",
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:            float64(188.18),
            },
        },
        Topic:                "ping",
    }

}
```

