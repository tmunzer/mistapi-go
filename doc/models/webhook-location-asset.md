
# Webhook Location Asset

Sample of the `location_asset` webhook payload.

## Structure

`WebhookLocationAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationAssetEvent`](../../doc/models/webhook-location-asset-event.md) | Required | Asset location events included in this webhook delivery |
| `Topic` | `string` | Required, Constant | Webhook topic name for asset location deliveries. enum: `location-asset`<br><br>**Value**: `"location-asset"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookLocationAsset := models.WebhookLocationAsset{
        Events:               []models.WebhookLocationAssetEvent{
            models.WebhookLocationAssetEvent{
                BatteryVoltage:        models.ToPointer(3370),
                EddystoneUidInstance:  models.ToPointer("5c5b35000001"),
                EddystoneUidNamespace: models.ToPointer("2818e3868dec25629ede"),
                EddystoneUrlUrl:       models.ToPointer("https://www.abc.com"),
                IbeaconMajor:          models.NewOptional(models.ToPointer(1234)),
                IbeaconMinor:          models.NewOptional(models.ToPointer(1234)),
                IbeaconUuid:           models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
                Mac:                   models.ToPointer("7fc2936fd243"),
                MapId:                 models.ToPointer(uuid.MustParse("845a23bf-bed9-e43c-4c86-6fa474be7ae5")),
                MfgCompanyId:          models.ToPointer(935),
                MfgData:               models.ToPointer("648520a1020000"),
                SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                Type:                  models.ToPointer("asset"),
                X:                     models.ToPointer(float64(13.5)),
                Y:                     models.ToPointer(float64(3.2)),
            },
        },
        Topic:                "location-asset",
    }

}
```

