
# Webhook Wifi Conn Raw Event Packet

RSSI observation for a connected Wi-Fi client packet

## Structure

`WebhookWifiConnRawEventPacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band on which the connected Wi-Fi packet was observed |
| `Rssi` | `*int` | Optional | Received signal strength of the connected Wi-Fi packet at the reporting AP, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookWifiConnRawEventPacket := models.WebhookWifiConnRawEventPacket{
        Band:                 models.ToPointer("band6"),
        Rssi:                 models.ToPointer(244),
    }

}
```

