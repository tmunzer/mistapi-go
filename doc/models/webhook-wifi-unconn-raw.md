
# Webhook Wifi Unconn Raw

Sample of the `wifi-unconn-raw` webhook payload.

This webhook topic provides raw data from packets emitted by unconnected devices.

Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client,
identified by their advertising MAC address (assets, discovered ble, connected wifi, unconnected wifi).  
The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry
packets that the client may emit.

Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client.
In order to get the location data for a client please see our location webhooks.
Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).

## Structure

`WebhookWifiUnconnRaw`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookWifiUnconnRawEvent`](../../doc/models/webhook-wifi-unconn-raw-event.md) | Required | Unconnected Wi-Fi raw packet events included in a webhook delivery<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | Webhook topic name for unconnected Wi-Fi raw packet deliveries. enum: `wifi-unconn-raw`<br><br>**Value**: `"wifi-unconn-raw"` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookWifiUnconnRaw := models.WebhookWifiUnconnRaw{
        Events:               []models.WebhookWifiUnconnRawEvent{
            models.WebhookWifiUnconnRawEvent{
                ApId:                 models.ToPointer("ap_id2"),
                ApLoc:                []float64{
                    float64(26.98),
                    float64(26.97),
                    float64(26.96),
                },
                ClientId:             models.ToPointer("client_id2"),
                ConnectedSite:        models.ToPointer(false),
                MapId:                models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Topic:                "wifi-unconn-raw",
    }

}
```

