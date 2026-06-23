
# Webhook Location Client

Sample of the `location-client` webhook payload.

## Structure

`WebhookLocationClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationClientEvent`](../../doc/models/webhook-location-client-event.md) | Required | Connected client location events included in this webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for connected client location deliveries. enum: `location-client`<br><br>**Value**: `"location-client"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationClient := models.WebhookLocationClient{
        Events:               []models.WebhookLocationClientEvent{
            models.WebhookLocationClientEvent{
                Mac:                    models.ToPointer("5684dae9ac8b"),
                MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
                SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Timestamp:              models.ToPointer(float64(188.18)),
                Type:                   models.ToPointer("wifi"),
                X:                      models.ToPointer(float64(13.5)),
                Y:                      models.ToPointer(float64(3.2)),
            },
        },
        Topic:                "location-client",
    }

}
```

