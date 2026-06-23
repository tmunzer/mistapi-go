
# Webhook Wifi Conn Raw Event

Raw packet observation for a connected Wi-Fi client

## Structure

`WebhookWifiConnRawEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApId` | `*string` | Optional | AP MAC address that reported the connected client packet |
| `ApLoc` | `[]float64` | Optional | optional, coordinates (if any) of reporting AP (updated once in 60s per client) |
| `ClientId` | `*string` | Optional | Client MAC address that emitted the connected Wi-Fi packet |
| `ConnectedSite` | `*bool` | Optional | Whether the client is connected to the reporting site |
| `ExtendedInfoList` | [`[]models.WebhookWifiConnRawEventExtendedInfo`](../../doc/models/webhook-wifi-conn-raw-event-extended-info.md) | Optional | optional, list of specific telemetry packets emited by certain wifi tags (Eg. Centrak) |
| `MapId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Packets` | [`[]models.WebhookWifiConnRawEventPacket`](../../doc/models/webhook-wifi-conn-raw-event-packet.md) | Optional | RSSI packet observations for a connected Wi-Fi raw event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookWifiConnRawEvent := models.WebhookWifiConnRawEvent{
        ApId:                 models.ToPointer("ap_id0"),
        ApLoc:                []float64{
            float64(129.94),
            float64(129.95),
            float64(129.96),
        },
        ClientId:             models.ToPointer("client_id0"),
        ConnectedSite:        models.ToPointer(false),
        ExtendedInfoList:     []models.WebhookWifiConnRawEventExtendedInfo{
            models.WebhookWifiConnRawEventExtendedInfo{
                FrameCtrl:            models.ToPointer(248),
                Payload:              models.ToPointer("payload2"),
                SequenceCtrl:         models.ToPointer(42),
            },
            models.WebhookWifiConnRawEventExtendedInfo{
                FrameCtrl:            models.ToPointer(248),
                Payload:              models.ToPointer("payload2"),
                SequenceCtrl:         models.ToPointer(42),
            },
            models.WebhookWifiConnRawEventExtendedInfo{
                FrameCtrl:            models.ToPointer(248),
                Payload:              models.ToPointer("payload2"),
                SequenceCtrl:         models.ToPointer(42),
            },
        },
        MapId:                models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

