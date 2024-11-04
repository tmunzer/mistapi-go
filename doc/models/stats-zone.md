
# Stats Zone

Zone statistics

## Structure

`StatsZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsWaits` | [`*models.StatsZoneAssetsWaits`](../../doc/models/stats-zone-assets-waits.md) | Optional | ble asset wait time right now |
| `ClientsWaits` | [`*models.StatsZoneClientsWaits`](../../doc/models/stats-zone-clients-waits.md) | Optional | client wait time right now |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `MapId` | `uuid.UUID` | Required | map_id of the zone |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | name of the zone |
| `NumAssets` | `*int` | Optional | number of assets |
| `NumClients` | `*int` | Optional | number of wifi clients (unconnected + connected) |
| `NumSdkclients` | `*int` | Optional | number of sdk clients |
| `OccupancyLimit` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SdkclientsWaits` | [`*models.StatsZoneSdkclientsWaits`](../../doc/models/stats-zone-sdkclients-waits.md) | Optional | sdkclient wait time right now |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Vertices` | [`[]models.ZoneVertex`](../../doc/models/zone-vertex.md) | Optional | vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area |
| `VerticesM` | [`[]models.ZoneVertexM`](../../doc/models/zone-vertex-m.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "123449d4-d12f-4feb-b40f-5be0e2ae1234",
  "name": "Zone A",
  "num_assets": 0,
  "num_clients": 80,
  "num_sdkclients": 10,
  "occupancy_limit": 4,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vertices": [
    {
      "x": 732,
      "y": 1821
    },
    {
      "x": 732.5,
      "y": 1731
    },
    {
      "x": 837.5,
      "y": 1731.5
    },
    {
      "x": 839,
      "y": 1821
    }
  ],
  "vertices_m": [
    {
      "x": 24.1983341951072,
      "y": 60.198314985369144
    },
    {
      "x": 24.21486311190714,
      "y": 57.22310996138056
    },
    {
      "x": 27.685935639893827,
      "y": 57.23963887818049
    },
    {
      "x": 27.73552239029364,
      "y": 60.198314985369144
    }
  ],
  "assets_waits": {
    "avg": 81.08,
    "max": 149.16,
    "min": 75.74,
    "p95": 47.26
  },
  "clients_waits": {
    "avg": 73.04,
    "max": 141.12,
    "min": 188.3,
    "p95": 39.22
  },
  "created_time": 13.1,
  "modified_time": 65.86
}
```

