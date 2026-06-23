
# Webhook Occupancy Alerts Event

Site-level occupancy alert batch containing one or more zone alerts

## Structure

`WebhookOccupancyAlertsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertEvents` | [`[]models.WebhookOccupancyAlertsEventAlertEventsItems`](../../doc/models/webhook-occupancy-alerts-event-alert-events-items.md) | Optional, Read-only | List of occupancy alerts for non-compliance zones within the site detected around the same time<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ForSite` | `*bool` | Optional, Read-only | Whether this occupancy alert batch is scoped to a site rather than only to the organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `SiteName` | `string` | Required, Read-only | Site name associated with the occupancy alert batch |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookOccupancyAlertsEvent := models.WebhookOccupancyAlertsEvent{
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
        SiteName:             "site_name6",
    }

}
```

