
# Events Client

Wireless client event returned by client event search APIs

## Structure

`EventsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | Access point MAC address associated with the wireless client event |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `Bssid` | `*string` | Optional | Wireless BSSID involved in the client event |
| `Channel` | `*int` | Optional | Radio channel used for the wireless client event |
| `KeyMgmt` | [`*models.ClientKeyMgmtEnum`](../../doc/models/client-key-mgmt-enum.md) | Optional | Key management protocol used for the latest authentication. enum: `WPA2-PSK`, `WPA2-PSK/CCMP`, `WPA2-PSK-FT`, `WPA2-PSK-SHA256`, `WPA3-EAP-SHA256`, `WPA3-EAP-SHA256/CCMP`, `WPA3-EAP-FT/GCMP256`, `WPA3-SAE-FT`, `WPA3-SAE-PSK` |
| `Proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Optional | enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n` |
| `Ssid` | `*string` | Optional | Wireless network SSID involved in the client event |
| `Text` | `*string` | Optional | Human-readable message for the client event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE |
| `TypeCode` | `*int` | Optional | Reason code for association or disassociation client events |
| `WlanId` | `*uuid.UUID` | Optional | Wireless LAN identifier associated with the client event |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    eventsClient := models.EventsClient{
        Ap:                   models.ToPointer("ap6"),
        Band:                 models.Dot11BandEnum_ENUM6DEDICATED,
        Bssid:                models.ToPointer("bssid8"),
        Channel:              models.ToPointer(156),
        KeyMgmt:              models.ToPointer(models.ClientKeyMgmtEnum_WPA2PSK),
        Proto:                models.ToPointer(models.Dot11ProtoEnum_N),
        Timestamp:            float64(67.74),
    }

}
```

