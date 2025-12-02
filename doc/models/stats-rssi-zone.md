
# Stats Rssi Zone

Zone statistics

## Structure

`StatsRssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsWait` | [`*models.StatsZoneAssetsWaits`](../../doc/models/stats-zone-assets-waits.md) | Optional | BLE asset wait time right now |
| `ClientsWait` | [`*models.StatsZoneClientsWaits`](../../doc/models/stats-zone-clients-waits.md) | Optional | Client wait time right now |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Devices` | [`[]models.StatsRssiZonesDevice`](../../doc/models/stats-rssi-zones-device.md) | Required | - |
| `DiscoveredAssetsWait` | [`*models.StatsZoneDiscoveredAssetsWaits`](../../doc/models/stats-zone-discovered-assets-waits.md) | Optional | Discovered asset wait time right now |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Name of the zone |
| `NumAssets` | `*int` | Optional | Number of assets |
| `NumClients` | `*int` | Optional | Number of Wi-Fi clients (unconnected + connected) |
| `NumDiscoveredAssets` | `*int` | Optional | Number of discoveredassets |
| `NumSdkclients` | `*int` | Optional | Number of sdk clients |
| `NumUnconnectedClients` | `*int` | Optional | Number of unconnected Wi-Fi clients |
| `OccupancyLimit` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SdkclientsWait` | [`*models.StatsZoneSdkclientsWaits`](../../doc/models/stats-zone-sdkclients-waits.md) | Optional | SDK client wait time right now |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UnconnectedClientsWait` | [`*models.StatsZoneUnconnectedClientsWaits`](../../doc/models/stats-zone-unconnected-clients-waits.md) | Optional | Unconnected Wi-Fi client wait time right now |

## Example (as JSON)

```json
{
  "devices": [
    {
      "device_id": "00002576-0000-0000-0000-000000000000",
      "rssi": 170
    }
  ],
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Zone A",
  "num_assets": 0,
  "num_clients": 80,
  "num_discovered_assets": 0,
  "num_sdkclients": 10,
  "num_unconnected_clients": 80,
  "occupancy_limit": 4,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "assets_wait": {
    "avg": 77.3,
    "max": 145.38,
    "min": 184.04,
    "p95": 43.48
  },
  "clients_wait": {
    "avg": 182.48,
    "max": 5.44,
    "min": 78.86,
    "p95": 107.34
  },
  "created_time": 18.26,
  "discovered_assets_wait": {
    "avg": 229.4,
    "max": 41.48,
    "min": 224.06,
    "p95": 195.58
  },
  "modified_time": 60.7
}
```

