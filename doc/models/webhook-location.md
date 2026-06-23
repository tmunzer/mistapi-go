
# Webhook Location

Sample of the `location` webhook payload.

## Structure

`WebhookLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationEvent`](../../doc/models/webhook-location-event.md) | Required | Location events included in this webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for generic location deliveries. enum: `location`<br><br>**Value**: `"location"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocation := models.WebhookLocation{
        Events:               []models.WebhookLocationEvent{
            models.WebhookLocationEvent{
                BatteryVoltage:         models.ToPointer(134),
                EddystoneUidInstance:   models.ToPointer("eddystone_uid_instance4"),
                EddystoneUidNamespace:  models.ToPointer("eddystone_uid_namespace6"),
                EddystoneUrlUrl:        models.ToPointer("eddystone_url_url4"),
                IbeaconMajor:           models.NewOptional(models.ToPointer(1234)),
                IbeaconMinor:           models.NewOptional(models.ToPointer(1234)),
                IbeaconUuid:            models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
                Id:                     uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                MapId:                  uuid.MustParse("00001148-0000-0000-0000-000000000000"),
                SiteId:                 uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                Timestamp:              float64(188.18),
                Type:                   "type0",
                X:                      float64(94.86),
                Y:                      float64(226.14),
            },
        },
        Topic:                "location",
    }

}
```

