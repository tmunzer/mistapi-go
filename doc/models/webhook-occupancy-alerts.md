
# Webhook Occupancy Alerts

Sample of the `occupancy-alerts` webhook payload.

## Structure

`WebhookOccupancyAlerts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookOccupancyAlertsEvent`](../../doc/models/webhook-occupancy-alerts-event.md) | Required | Occupancy alert batches included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for occupancy alert deliveries. enum: `occupancy-alerts`<br><br>**Value**: `"occupancy-alerts"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookOccupancyAlerts := models.WebhookOccupancyAlerts{
        Events:               []models.WebhookOccupancyAlertsEvent{
            models.WebhookOccupancyAlertsEvent{
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SiteName:             "",
            },
        },
        Topic:                "occupancy-alerts",
    }

}
```

