
# Webhook Zone Event

Zone enter or exit event for a Wi-Fi client, SDK client, or asset

## Structure

`WebhookZoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetId` | `*uuid.UUID` | Optional | Only if `type`==`asset`. UUID of named asset |
| `Id` | `*uuid.UUID` | Optional | Only if `type`==`sdk`. UUID of the SDK Client |
| `Mac` | `*string` | Optional | Client or asset MAC address associated with the zone event |
| `MapId` | `uuid.UUID` | Required | Map associated with the zone event |
| `Name` | `*string` | Optional | Display name of the client or asset, when available |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Trigger` | [`models.WebhookZoneEventTriggerEnum`](../../doc/models/webhook-zone-event-trigger-enum.md) | Required | Zone transition direction, either enter or exit. enum: `enter`, `exit` |
| `Type` | [`models.WebhookZoneEventTypeEnum`](../../doc/models/webhook-zone-event-type-enum.md) | Required | Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi` |
| `ZoneId` | `uuid.UUID` | Required | Zone identifier associated with the event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookZoneEvent := models.WebhookZoneEvent{
        AssetId:              models.ToPointer(uuid.MustParse("00001fc4-0000-0000-0000-000000000000")),
        Id:                   models.ToPointer(uuid.MustParse("00001cea-0000-0000-0000-000000000000")),
        Mac:                  models.ToPointer("mac6"),
        MapId:                uuid.MustParse("00000142-0000-0000-0000-000000000000"),
        Name:                 models.ToPointer("name2"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:            float64(114.8),
        Trigger:              models.WebhookZoneEventTriggerEnum_ENTER,
        Type:                 models.WebhookZoneEventTypeEnum_SDK,
        ZoneId:               uuid.MustParse("00001f66-0000-0000-0000-000000000000"),
    }

}
```

