
# Webhook Wifi Unconn Raw Event Packet

RSSI observation for an unconnected Wi-Fi client packet

## Structure

`WebhookWifiUnconnRawEventPacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band on which the unconnected Wi-Fi packet was observed |
| `Rssi` | `*int` | Optional | Received signal strength of the unconnected Wi-Fi packet at the reporting AP, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookWifiUnconnRawEventPacket := models.WebhookWifiUnconnRawEventPacket{
        Band:                 models.ToPointer("band4"),
        Rssi:                 models.ToPointer(182),
    }

}
```

