
# Webhook Rssizone Event

RSSI zone transition event for a Wi-Fi, SDK, or asset client

## Structure

`WebhookRssizoneEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Client MAC address associated with the RSSI zone event |
| `MapId` | `uuid.UUID` | Required | Map associated with the RSSI zone event |
| `RssizoneId` | `uuid.UUID` | Required | RSSI zone identifier associated with the event |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Trigger` | [`models.WebhookZoneEventTriggerEnum`](../../doc/models/webhook-zone-event-trigger-enum.md) | Required | Zone transition direction, either enter or exit. enum: `enter`, `exit` |
| `Type` | [`models.WebhookZoneEventTypeEnum`](../../doc/models/webhook-zone-event-type-enum.md) | Required | Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookRssizoneEvent := models.WebhookRssizoneEvent{
        Mac:                  "mac8",
        MapId:                uuid.MustParse("000018a6-0000-0000-0000-000000000000"),
        RssizoneId:           uuid.MustParse("000007b4-0000-0000-0000-000000000000"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:            float64(74.68),
        Trigger:              models.WebhookZoneEventTriggerEnum_ENTER,
        Type:                 models.WebhookZoneEventTypeEnum_ASSET,
    }

}
```

