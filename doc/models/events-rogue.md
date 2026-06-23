
# Events Rogue

Rogue wireless network event detected by a site AP

## Structure

`EventsRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | Access point MAC address that detected the rogue BSSID |
| `Bssid` | `string` | Required | Rogue BSSID observed by the AP |
| `Channel` | `int` | Required | Radio channel where the rogue BSSID was observed |
| `Rssi` | `int` | Required | Received signal strength of the rogue BSSID, in dBm |
| `Ssid` | `string` | Required | Wireless network SSID advertised by the rogue BSSID |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    eventsRogue := models.EventsRogue{
        Ap:                   "ap4",
        Bssid:                "bssid4",
        Channel:              244,
        Rssi:                 110,
        Ssid:                 "ssid2",
        Timestamp:            float64(148.98),
    }

}
```

