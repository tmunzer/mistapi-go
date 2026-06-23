
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
    "github.com/google/uuid"
)

func main() {
    webhookClientInfo := models.WebhookClientInfo{
        Events:               []models.WebhookClientInfoEvent{
            models.WebhookClientInfoEvent{
                Hostname:             models.ToPointer("hostname6"),
                Ip:                   models.ToPointer("ip4"),
                Mac:                  models.ToPointer("mac4"),
                OrgId:                models.ToPointer(uuid.MustParse("00000dbc-0000-0000-0000-000000000000")),
                SiteId:               models.ToPointer(uuid.MustParse("0000245a-0000-0000-0000-000000000000")),
            },
            models.WebhookClientInfoEvent{
                Hostname:             models.ToPointer("hostname6"),
                Ip:                   models.ToPointer("ip4"),
                Mac:                  models.ToPointer("mac4"),
                OrgId:                models.ToPointer(uuid.MustParse("00000dbc-0000-0000-0000-000000000000")),
                SiteId:               models.ToPointer(uuid.MustParse("0000245a-0000-0000-0000-000000000000")),
            },
        },
        Topic:                models.ToPointer(models.WebhookClientInfoTopicEnum_CLIENTINFO),
    }

}
```

