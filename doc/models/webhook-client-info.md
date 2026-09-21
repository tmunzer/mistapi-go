
# Webhook Client Info

Sample of the `client-info` webhook payload.

## Structure

`WebhookClientInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookClientInfoEvent`](../../doc/models/webhook-client-info-event.md) | Optional | Client information events included in a webhook delivery |
| `Topic` | [`*models.WebhookClientInfoTopicEnum`](../../doc/models/webhook-client-info-topic-enum.md) | Optional | Webhook topic name for client information deliveries. enum: `client-info` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookClientInfo := models.WebhookClientInfo{
        Events:               []models.WebhookClientInfoEvent{
            models.WebhookClientInfoEvent{
                Hostname:             models.ToPointer("hostname6"),
                Ip:                   models.ToPointer("ip4"),
                Mac:                  models.ToPointer("mac4"),
            },
            models.WebhookClientInfoEvent{
                Hostname:             models.ToPointer("hostname6"),
                Ip:                   models.ToPointer("ip4"),
                Mac:                  models.ToPointer("mac4"),
            },
        },
        Topic:                models.ToPointer(models.WebhookClientInfoTopicEnum_CLIENTINFO),
    }

}
```

