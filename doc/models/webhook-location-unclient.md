
# Webhook Location Unclient

Sample of the `location-unclient` webhook payload.

## Structure

`WebhookLocationUnclient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationUnclientEvent`](../../doc/models/webhook-location-unclient-event.md) | Required | Unconnected client location events included in this webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for unconnected client location deliveries. enum: `location-unclient`<br><br>**Value**: `"location-unclient"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationUnclient := models.WebhookLocationUnclient{
        Events:               []models.WebhookLocationUnclientEvent{
            models.WebhookLocationUnclientEvent{
                Mac:                    models.ToPointer("5684dae9ac8b"),
                MapId:                  models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
                SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Timestamp:              models.ToPointer(float64(188.18)),
                Type:                   models.ToPointer("wifi"),
                X:                      models.ToPointer(float64(13.5)),
                Y:                      models.ToPointer(float64(3.2)),
            },
        },
        Topic:                "location-unclient",
    }

}
```

