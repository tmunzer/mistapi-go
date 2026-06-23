
# Stats Zone

Zone statistics and occupancy counts for a site map zone

## Structure

`StatsZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsWait` | [`*models.StatsZoneAssetsWaits`](../../doc/models/stats-zone-assets-waits.md) | Optional | BLE asset wait time right now |
| `ClientsWait` | [`*models.StatsZoneClientsWaits`](../../doc/models/stats-zone-clients-waits.md) | Optional | Client wait time right now |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DiscoveredAssetsWait` | [`*models.StatsZoneDiscoveredAssetsWaits`](../../doc/models/stats-zone-discovered-assets-waits.md) | Optional | Discovered asset wait time right now |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `MapId` | `uuid.UUID` | Required | Map identifier for the zone |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the zone |
| `NumAssets` | `*int` | Optional | Number of BLE assets currently counted in the zone |
| `NumClients` | `*int` | Optional | Number of Wi-Fi clients currently counted in the zone, including connected and unconnected clients |
| `NumDiscoveredAssets` | `*int` | Optional | Number of discovered BLE assets currently counted in the zone |
| `NumSdkclients` | `*int` | Optional | Number of SDK clients currently counted in the zone |
| `NumUnconnectedClients` | `*int` | Optional | Number of unconnected Wi-Fi clients currently counted in the zone |
| `OccupancyLimit` | `*int` | Optional | Configured occupancy limit for the zone |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SdkclientsWait` | [`*models.StatsZoneSdkclientsWaits`](../../doc/models/stats-zone-sdkclients-waits.md) | Optional | SDK client wait time right now |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `UnconnectedClientsWait` | [`*models.StatsZoneUnconnectedClientsWaits`](../../doc/models/stats-zone-unconnected-clients-waits.md) | Optional | Unconnected Wi-Fi client wait time right now |
| `Vertices` | [`[]models.ZoneVertex`](../../doc/models/zone-vertex.md) | Optional | Vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area |
| `VerticesM` | [`[]models.ZoneVertexM`](../../doc/models/zone-vertex-m.md) | Optional | Vertices used to define a zone boundary, expressed in meters |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsZone := models.StatsZone{
        AssetsWait:             models.ToPointer(models.StatsZoneAssetsWaits{
            Avg:                  models.ToPointer(float64(77.3)),
            Max:                  models.ToPointer(float64(145.38)),
            Min:                  models.ToPointer(float64(184.04)),
            P95:                  models.ToPointer(float64(43.48)),
        }),
        ClientsWait:            models.ToPointer(models.StatsZoneClientsWaits{
            Avg:                  models.ToPointer(float64(182.48)),
            Max:                  models.ToPointer(float64(5.44)),
            Min:                  models.ToPointer(float64(78.86)),
            P95:                  models.ToPointer(float64(107.34)),
        }),
        CreatedTime:            models.ToPointer(float64(223.8)),
        DiscoveredAssetsWait:   models.ToPointer(models.StatsZoneDiscoveredAssetsWaits{
            Avg:                  models.ToPointer(float64(229.4)),
            Max:                  models.ToPointer(float64(41.48)),
            Min:                  models.ToPointer(float64(224.06)),
            P95:                  models.ToPointer(float64(195.58)),
        }),
        Id:                     uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MapId:                  uuid.MustParse("123449d4-d12f-4feb-b40f-5be0e2ae1234"),
        ModifiedTime:           models.ToPointer(float64(111.16)),
        Name:                   "Zone A",
        NumAssets:              models.ToPointer(0),
        NumClients:             models.ToPointer(80),
        NumDiscoveredAssets:    models.ToPointer(0),
        NumSdkclients:          models.ToPointer(10),
        NumUnconnectedClients:  models.ToPointer(80),
        OccupancyLimit:         models.ToPointer(4),
        OrgId:                  models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Vertices:               []models.ZoneVertex{
            models.ZoneVertex{
                X:                    float64(732),
                Y:                    float64(1821),
            },
            models.ZoneVertex{
                X:                    float64(732.5),
                Y:                    float64(1731),
            },
            models.ZoneVertex{
                X:                    float64(837.5),
                Y:                    float64(1731.5),
            },
            models.ZoneVertex{
                X:                    float64(839),
                Y:                    float64(1821),
            },
        },
        VerticesM:              []models.ZoneVertexM{
            models.ZoneVertexM{
                X:                    float64(24.1983341951072),
                Y:                    float64(60.198314985369144),
            },
            models.ZoneVertexM{
                X:                    float64(24.21486311190714),
                Y:                    float64(57.22310996138056),
            },
            models.ZoneVertexM{
                X:                    float64(27.685935639893827),
                Y:                    float64(57.23963887818049),
            },
            models.ZoneVertexM{
                X:                    float64(27.73552239029364),
                Y:                    float64(60.198314985369144),
            },
        },
    }

}
```

