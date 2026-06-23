
# Webhook Nac Events

Sample of the `nac-events` webhook payload.

## Structure

`WebhookNacEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.NacClientEvent`](../../doc/models/nac-client-event.md) | Optional | List of NAC authentication events |
| `Topic` | [`*models.WebhookNacEventsTopicEnum`](../../doc/models/webhook-nac-events-topic-enum.md) | Optional | Webhook topic name for NAC event deliveries. enum: `nac-events` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookNacEvents := models.WebhookNacEvents{
        Events:               []models.NacClientEvent{
            models.NacClientEvent{
                Ap:                   models.ToPointer("ap6"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_CERT),
                Bssid:                models.ToPointer("bssid4"),
                ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRED),
                DeviceMac:            models.ToPointer("device_mac4"),
            },
            models.NacClientEvent{
                Ap:                   models.ToPointer("ap6"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_CERT),
                Bssid:                models.ToPointer("bssid4"),
                ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRED),
                DeviceMac:            models.ToPointer("device_mac4"),
            },
            models.NacClientEvent{
                Ap:                   models.ToPointer("ap6"),
                AuthType:             models.ToPointer(models.NacAuthTypeEnum_CERT),
                Bssid:                models.ToPointer("bssid4"),
                ClientType:           models.ToPointer(models.NacAccessTypeEnum_WIRED),
                DeviceMac:            models.ToPointer("device_mac4"),
            },
        },
        Topic:                models.ToPointer(models.WebhookNacEventsTopicEnum_NACEVENTS),
    }

}
```

