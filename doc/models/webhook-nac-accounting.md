
# Webhook Nac Accounting

Sample of the `nac-accounting` webhook payload.

## Structure

`WebhookNacAccounting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacAccountingEvent`](../../doc/models/webhook-nac-accounting-event.md) | Optional | NAC accounting events included in a webhook delivery |
| `Topic` | [`*models.WebhookNacAccountingTopicEnum`](../../doc/models/webhook-nac-accounting-topic-enum.md) | Optional | Webhook topic name for NAC accounting deliveries. enum: `nac-accounting` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookNacAccounting := models.WebhookNacAccounting{
        Events:               []models.WebhookNacAccountingEvent{
            models.WebhookNacAccountingEvent{
                Ap:                   models.ToPointer("ap6"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_CERT),
                Bssid:                models.ToPointer("bssid4"),
                ClientIp:             models.ToPointer("client_ip0"),
                ClientType:           models.ToPointer("client_type6"),
            },
        },
        Topic:                models.ToPointer(models.WebhookNacAccountingTopicEnum_NACACCOUNTING),
    }

}
```

