
# Webhook Occupancy Alerts Event Alert Events Items

Occupancy compliance alert for a single zone

## Structure

`WebhookOccupancyAlertsEventAlertEventsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentOccupancy` | `int` | Required | Current number of occupants detected in the zone |
| `MapId` | `uuid.UUID` | Required | Map associated with the occupancy zone |
| `OccupancyLimit` | `int` | Required | Configured occupancy limit for the zone |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | [`models.WebhookOccupancyAlertTypeEnum`](../../doc/models/webhook-occupancy-alert-type-enum.md) | Required | Occupancy compliance state reported for the zone. enum: `COMPLIANCE-OK`, `COMPLIANCE-VIOLATION` |
| `ZoneId` | `uuid.UUID` | Required | Zone identifier associated with the occupancy alert |
| `ZoneName` | `string` | Required | Zone name associated with the occupancy alert |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookOccupancyAlertsEventAlertEventsItems := models.WebhookOccupancyAlertsEventAlertEventsItems{
        CurrentOccupancy:     70,
        MapId:                uuid.MustParse("0000145e-0000-0000-0000-000000000000"),
        OccupancyLimit:       178,
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Timestamp:            float64(236.28),
        Type:                 models.WebhookOccupancyAlertTypeEnum_COMPLIANCEOK,
        ZoneId:               uuid.MustParse("00000c4a-0000-0000-0000-000000000000"),
        ZoneName:             "zone_name0",
    }

}
```

