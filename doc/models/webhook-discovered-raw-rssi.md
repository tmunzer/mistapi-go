
# Webhook Discovered Raw Rssi

Sample of the `discovered-raw-rssi` webhook payload.

## Structure

`WebhookDiscoveredRawRssi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDiscoveredRawRssiEvent`](../../doc/models/webhook-discovered-raw-rssi-event.md) | Optional | Discovered BLE raw RSSI events included in a webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for discovered raw RSSI deliveries. enum: `discovered-raw-rssi`<br><br>**Value**: `"discovered-raw-rssi"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookDiscoveredRawRssi := models.WebhookDiscoveredRawRssi{
        Events:               []models.WebhookDiscoveredRawRssiEvent{
            models.WebhookDiscoveredRawRssiEvent{
                ApLoc:                []float64{
                    float64(26.98),
                    float64(26.97),
                    float64(26.96),
                },
                Beam:                 124,
                DeviceId:             uuid.MustParse("0000254a-0000-0000-0000-000000000000"),
                IbeaconMajor:         models.NewOptional(models.ToPointer(178)),
                IbeaconMinor:         models.NewOptional(models.ToPointer(40)),
                IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("0000149a-0000-0000-0000-000000000000"))),
                IsAsset:              models.ToPointer(false),
                Mac:                  "mac4",
                MapId:                uuid.MustParse("00001148-0000-0000-0000-000000000000"),
                OrgId:                uuid.MustParse("00000dbc-0000-0000-0000-000000000000"),
                Rssi:                 float64(58.22),
                SiteId:               uuid.MustParse("0000245a-0000-0000-0000-000000000000"),
            },
            models.WebhookDiscoveredRawRssiEvent{
                ApLoc:                []float64{
                    float64(26.98),
                    float64(26.97),
                    float64(26.96),
                },
                Beam:                 124,
                DeviceId:             uuid.MustParse("0000254a-0000-0000-0000-000000000000"),
                IbeaconMajor:         models.NewOptional(models.ToPointer(178)),
                IbeaconMinor:         models.NewOptional(models.ToPointer(40)),
                IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("0000149a-0000-0000-0000-000000000000"))),
                IsAsset:              models.ToPointer(false),
                Mac:                  "mac4",
                MapId:                uuid.MustParse("00001148-0000-0000-0000-000000000000"),
                OrgId:                uuid.MustParse("00000dbc-0000-0000-0000-000000000000"),
                Rssi:                 float64(58.22),
                SiteId:               uuid.MustParse("0000245a-0000-0000-0000-000000000000"),
            },
        },
        Topic:                "discovered-raw-rssi",
    }

}
```

