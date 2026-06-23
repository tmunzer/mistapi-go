
# Webhook Wifi Unconn Raw Event

Raw packet observation for an unconnected Wi-Fi client

## Structure

`WebhookWifiUnconnRawEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApId` | `*string` | Optional | AP MAC address that reported the unconnected client packet |
| `ApLoc` | `[]float64` | Optional | optional, coordinates (if any) of reporting AP (updated once in 60s per client) |
| `ClientId` | `*string` | Optional | Client MAC address that emitted the unconnected Wi-Fi packet |
| `ConnectedSite` | `*bool` | Optional | Whether the observed client is connected to the reporting site |
| `MapId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Packets` | [`[]models.WebhookWifiUnconnRawEventPacket`](../../doc/models/webhook-wifi-unconn-raw-event-packet.md) | Optional | RSSI packet observations for an unconnected Wi-Fi raw event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookWifiUnconnRawEvent := models.WebhookWifiUnconnRawEvent{
        ApId:                 models.ToPointer("ap_id6"),
        ApLoc:                []float64{
            float64(53.08),
            float64(53.09),
            float64(53.1),
        },
        ClientId:             models.ToPointer("client_id6"),
        ConnectedSite:        models.ToPointer(false),
        MapId:                models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

