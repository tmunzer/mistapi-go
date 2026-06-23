
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
                AlertEvents:          []models.WebhookOccupancyAlertsEventAlertEventsItems{
                    models.WebhookOccupancyAlertsEventAlertEventsItems{
                        CurrentOccupancy:     30,
                        MapId:                uuid.MustParse("00000fb2-0000-0000-0000-000000000000"),
                        OccupancyLimit:       78,
                        OrgId:                uuid.MustParse("00000f52-0000-0000-0000-000000000000"),
                        Timestamp:            float64(148.24),
                        Type:                 models.WebhookOccupancyAlertTypeEnum_COMPLIANCEOK,
                        ZoneId:               uuid.MustParse("000010f6-0000-0000-0000-000000000000"),
                        ZoneName:             "zone_name6",
                    },
                    models.WebhookOccupancyAlertsEventAlertEventsItems{
                        CurrentOccupancy:     30,
                        MapId:                uuid.MustParse("00000fb2-0000-0000-0000-000000000000"),
                        OccupancyLimit:       78,
                        OrgId:                uuid.MustParse("00000f52-0000-0000-0000-000000000000"),
                        Timestamp:            float64(148.24),
                        Type:                 models.WebhookOccupancyAlertTypeEnum_COMPLIANCEOK,
                        ZoneId:               uuid.MustParse("000010f6-0000-0000-0000-000000000000"),
                        ZoneName:             "zone_name6",
                    },
                    models.WebhookOccupancyAlertsEventAlertEventsItems{
                        CurrentOccupancy:     30,
                        MapId:                uuid.MustParse("00000fb2-0000-0000-0000-000000000000"),
                        OccupancyLimit:       78,
                        OrgId:                uuid.MustParse("00000f52-0000-0000-0000-000000000000"),
                        Timestamp:            float64(148.24),
                        Type:                 models.WebhookOccupancyAlertTypeEnum_COMPLIANCEOK,
                        ZoneId:               uuid.MustParse("000010f6-0000-0000-0000-000000000000"),
                        ZoneName:             "zone_name6",
                    },
                },
                ForSite:              models.ToPointer(false),
                SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
                SiteName:             "site_name2",
            },
        },
        Topic:                "occupancy-alerts",
    }

}
```

