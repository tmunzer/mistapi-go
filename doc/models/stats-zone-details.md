
# Stats Zone Details

Detailed zone statistics and occupants for a site map zone

## Structure

`StatsZoneDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Assets` | `[]string` | Optional | List of ble assets currently in the zone and when they entered |
| `ClientWaits` | [`models.StatsZoneDetailsClientWaits`](../../doc/models/stats-zone-details-client-waits.md) | Required | Client wait time right now |
| `Clients` | `[]string` | Optional | List of clients currently in the zone and when they entered |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `MapId` | `uuid.UUID` | Required | Map identifier for the zone |
| `Name` | `string` | Required | Display name of the zone |
| `NumClients` | `int` | Required | Number of Wi-Fi clients currently counted in the zone |
| `NumSdkclients` | `int` | Required | Number of SDK clients currently counted in the zone |
| `Sdkclients` | `[]string` | Optional | List of SDK Clients currently in the zone and when they entered |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsZoneDetails := models.StatsZoneDetails{
        Assets:               []string{
            "assets2",
        },
        ClientWaits:          models.StatsZoneDetailsClientWaits{
            Avg:                  1200,
            Max:                  3610,
            Min:                  600,
            P95:                  2800,
        },
        Clients:              []string{
            "clients2",
            "clients3",
        },
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MapId:                uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc"),
        Name:                 "Board Room",
        NumClients:           80,
        NumSdkclients:        0,
        Sdkclients:           []string{
            "sdkclients6",
        },
    }

}
```

