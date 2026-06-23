
# Stats Rssi Zone

RSSI-based zone statistics for a site

## Structure

`StatsRssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsWait` | [`*models.StatsZoneAssetsWaits`](../../doc/models/stats-zone-assets-waits.md) | Optional | BLE asset wait time right now |
| `ClientsWait` | [`*models.StatsZoneClientsWaits`](../../doc/models/stats-zone-clients-waits.md) | Optional | Client wait time right now |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Devices` | [`[]models.StatsRssiZonesDevice`](../../doc/models/stats-rssi-zones-device.md) | Required | AP devices and RSSI thresholds that define an RSSI zone |
| `DiscoveredAssetsWait` | [`*models.StatsZoneDiscoveredAssetsWaits`](../../doc/models/stats-zone-discovered-assets-waits.md) | Optional | Discovered asset wait time right now |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the RSSI zone |
| `NumAssets` | `*int` | Optional | Number of BLE assets currently counted in the RSSI zone |
| `NumClients` | `*int` | Optional | Number of Wi-Fi clients (unconnected + connected) |
| `NumDiscoveredAssets` | `*int` | Optional | Number of discovered BLE assets currently counted in the RSSI zone |
| `NumSdkclients` | `*int` | Optional | Number of SDK clients currently counted in the RSSI zone |
| `NumUnconnectedClients` | `*int` | Optional | Number of unconnected Wi-Fi clients |
| `OccupancyLimit` | `*int` | Optional | Configured occupancy limit for the RSSI zone |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SdkclientsWait` | [`*models.StatsZoneSdkclientsWaits`](../../doc/models/stats-zone-sdkclients-waits.md) | Optional | SDK client wait time right now |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `UnconnectedClientsWait` | [`*models.StatsZoneUnconnectedClientsWaits`](../../doc/models/stats-zone-unconnected-clients-waits.md) | Optional | Unconnected Wi-Fi client wait time right now |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsRssiZone := models.StatsRssiZone{
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
        CreatedTime:            models.ToPointer(float64(3.88)),
        Devices:                []models.StatsRssiZonesDevice{
            models.StatsRssiZonesDevice{
                DeviceId:             models.ToPointer(uuid.MustParse("00002576-0000-0000-0000-000000000000")),
                Rssi:                 models.ToPointer(170),
            },
        },
        DiscoveredAssetsWait:   models.ToPointer(models.StatsZoneDiscoveredAssetsWaits{
            Avg:                  models.ToPointer(float64(229.4)),
            Max:                  models.ToPointer(float64(41.48)),
            Min:                  models.ToPointer(float64(224.06)),
            P95:                  models.ToPointer(float64(195.58)),
        }),
        Id:                     uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        ModifiedTime:           models.ToPointer(float64(75.08)),
        Name:                   "Zone A",
        NumAssets:              models.ToPointer(0),
        NumClients:             models.ToPointer(80),
        NumDiscoveredAssets:    models.ToPointer(0),
        NumSdkclients:          models.ToPointer(10),
        NumUnconnectedClients:  models.ToPointer(80),
        OccupancyLimit:         models.ToPointer(4),
        OrgId:                  models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

