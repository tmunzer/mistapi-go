
# Webhook Asset Raw Rssi

Sample of the `asset-raw-rssi` webhook payload.

This webhook topic provides raw data from packets emitted by named and filtered assets.

Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client,
identified by their advertising MAC address (assets, discovered ble, connected wifi, unconnected wifi).  
The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry
packets that the client may emit.

Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client.
In order to get the location data for a client please see our location webhooks.
Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).

## Structure

`WebhookAssetRawRssi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookAssetRawRssiEvent`](../../doc/models/webhook-asset-raw-rssi-event.md) | Required | Asset raw RSSI events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for asset raw RSSI deliveries. enum: `asset-raw-rssi`<br><br>**Value**: `"asset-raw-rssi"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookAssetRawRssi := models.WebhookAssetRawRssi{
        Events:               []models.WebhookAssetRawRssiEvent{
            models.WebhookAssetRawRssiEvent{
                ApLoc:                []float64{
                    float64(26.98),
                    float64(26.97),
                    float64(26.96),
                },
                Beam:                 models.ToPointer(9),
                DeviceId:             models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                IbeaconMajor:         models.NewOptional(models.ToPointer(1234)),
                IbeaconMinor:         models.NewOptional(models.ToPointer(1234)),
                IbeaconUuid:          models.NewOptional(models.ToPointer(uuid.MustParse("f3f17139-704a-f03a-2786-0400279e37c3"))),
                MapId:                models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Topic:                "asset-raw-rssi",
    }

}
```

