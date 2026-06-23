
# Marvis Client Event

A Marvis Client event record

## Structure

`MarvisClientEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Wi-Fi band at the time of the event |
| `Bssid` | `*string` | Optional | BSSID the client roamed to (for roam events) |
| `Channel` | `*int` | Optional | Channel the client roamed to (for roam events) |
| `DeviceId` | `*uuid.UUID` | Optional | UUID of the device the Marvis Client is installed on |
| `Hostname` | `*string` | Optional | Device hostname |
| `Location` | [`*models.StatsMarvisClientLocation`](../../doc/models/stats-marvis-client-location.md) | Optional | Last known location fix for a Marvis Client device |
| `NeighborApReport` | [`[]models.MarvisClientEventNeighborAp`](../../doc/models/marvis-client-event-neighbor-ap.md) | Optional | List of neighboring APs observed at the time of the event |
| `OrgId` | `*uuid.UUID` | Optional | Organization UUID |
| `Percent` | `*int` | Optional | Battery level percentage at the time of the event (for battery events) |
| `PreBssid` | `*string` | Optional | BSSID the client roamed from (for roam events) |
| `PreChannel` | `*int` | Optional | Channel the client roamed from (for roam events) |
| `PreRssi` | `*int` | Optional | RSSI before the roam event, in dBm |
| `Rssi` | `*int` | Optional | Wi-Fi RSSI at the time of the event, in dBm |
| `Ssid` | `*string` | Optional | SSID the client was connected to |
| `Timestamp` | `*int` | Optional | Event timestamp, in epoch seconds |
| `Type` | `*string` | Optional | Event type |
| `WifiIp` | `*string` | Optional | Device Wi-Fi IP address at the time of the event |
| `WifiMac` | `*string` | Optional | Device Wi-Fi MAC address |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    marvisClientEvent := models.MarvisClientEvent{
        Band:                 models.ToPointer("band6"),
        Bssid:                models.ToPointer("bssid8"),
        Channel:              models.ToPointer(204),
        DeviceId:             models.ToPointer(uuid.MustParse("00001432-0000-0000-0000-000000000000")),
        Hostname:             models.ToPointer("hostname0"),
    }

}
```

