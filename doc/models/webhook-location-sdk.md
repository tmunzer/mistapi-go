
# Webhook Location Sdk

Sample of the `location-sdk` webhook payload.

## Structure

`WebhookLocationSdk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationSdkEvent`](../../doc/models/webhook-location-sdk-event.md) | Required | SDK client location events included in this webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for SDK client location deliveries. enum: `location-sdk`<br><br>**Value**: `"location-sdk"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationSdk := models.WebhookLocationSdk{
        Events:               []models.WebhookLocationSdkEvent{
            models.WebhookLocationSdkEvent{
                Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                MapId:                models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
                Name:                 models.ToPointer("optional"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Timestamp:            models.ToPointer(float64(188.18)),
                Type:                 models.ToPointer("sdk"),
                X:                    models.ToPointer(float64(13.5)),
                Y:                    models.ToPointer(float64(3.2)),
            },
        },
        Topic:                "location-sdk",
    }

}
```

