
# Webhook Location Centrak

Sample of the `location-centrak` webhook payload.

## Structure

`WebhookLocationCentrak`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationCentrakEvent`](../../doc/models/webhook-location-centrak-event.md) | Required | CenTrak location events included in this webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for CenTrak location deliveries. enum: `location-centrak`<br><br>**Value**: `"location-centrak"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationCentrak := models.WebhookLocationCentrak{
        Events:               []models.WebhookLocationCentrakEvent{
            models.WebhookLocationCentrakEvent{
                Mac:                    models.ToPointer("mac4"),
                MapId:                  models.ToPointer("map_id4"),
                MfgCompanyId:           models.ToPointer(234),
                MfgData:                models.ToPointer("mfg_data2"),
                SiteId:                 models.ToPointer(uuid.MustParse("0000245a-0000-0000-0000-000000000000")),
            },
        },
        Topic:                "location-centrak",
    }

}
```

