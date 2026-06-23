
# Sdkstats Wireless Client

Detailed statistics for an individual SDK client

## Structure

`SdkstatsWirelessClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | Map identifier for the SDK client's location, if known |
| `Name` | `*string` | Optional | Display name provided for the SDK client |
| `NetworkConnection` | [`*models.StatsSdkclientNetworkConnection`](../../doc/models/stats-sdkclient-network-connection.md) | Optional | Current network connection details reported for an SDK client |
| `Uuid` | `uuid.UUID` | Required | Application UUID for the SDK client |
| `Vbeacons` | [`[]models.SdkstatsWirelessClientVbeacon`](../../doc/models/sdkstats-wireless-client-vbeacon.md) | Optional | Virtual beacon associations for the SDK client, including when each began<br><br>**Constraints**: *Unique Items Required* |
| `X` | `*float64` | Optional | Horizontal map coordinate of the SDK client location, in pixels, if known |
| `Y` | `*float64` | Optional | Vertical map coordinate of the SDK client location, in pixels, if known |
| `Zones` | [`[]models.SdkstatsWirelessClientZone`](../../doc/models/sdkstats-wireless-client-zone.md) | Optional | Zone memberships for the SDK client, including when each membership began<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sdkstatsWirelessClient := models.SdkstatsWirelessClient{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        MapId:                models.NewOptional(models.ToPointer(uuid.MustParse("00001d10-0000-0000-0000-000000000000"))),
        Name:                 models.ToPointer("name4"),
        NetworkConnection:    models.ToPointer(models.StatsSdkclientNetworkConnection{
            Mac:                  "mac2",
            Rssi:                 float64(235.8),
            SignalLevel:          float64(47.82),
            Type:                 "type2",
        }),
        Uuid:                 uuid.MustParse("0000203a-0000-0000-0000-000000000000"),
        Vbeacons:             []models.SdkstatsWirelessClientVbeacon{
            models.SdkstatsWirelessClientVbeacon{
                Id:                   uuid.MustParse("00001cc0-0000-0000-0000-000000000000"),
                Since:                float64(115.2),
            },
            models.SdkstatsWirelessClientVbeacon{
                Id:                   uuid.MustParse("00001cc0-0000-0000-0000-000000000000"),
                Since:                float64(115.2),
            },
            models.SdkstatsWirelessClientVbeacon{
                Id:                   uuid.MustParse("00001cc0-0000-0000-0000-000000000000"),
                Since:                float64(115.2),
            },
        },
    }

}
```

