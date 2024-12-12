
# Stats Rssi Zone

Zone statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsRssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AssetsWaits` | [`*models.StatsZoneAssetsWaits`](../../doc/models/stats-zone-assets-waits.md) | Optional | ble asset wait time right now |
| `ClientsWaits` | [`*models.StatsZoneClientsWaits`](../../doc/models/stats-zone-clients-waits.md) | Optional | client wait time right now |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Devices` | [`[]models.StatsRssiZonesDevice`](../../doc/models/stats-rssi-zones-device.md) | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | name of the zone |
| `NumAssets` | `*int` | Optional | number of assets |
| `NumClients` | `*int` | Optional | number of wifi clients (unconnected + connected) |
| `NumSdkclients` | `*int` | Optional | number of sdk clients |
| `OccupancyLimit` | `*int` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SdkclientsWaits` | [`*models.StatsZoneSdkclientsWaits`](../../doc/models/stats-zone-sdkclients-waits.md) | Optional | sdkclient wait time right now |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "devices": [
    {
      "device_id": "00002576-0000-0000-0000-000000000000",
      "rssi": 170,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Zone A",
  "num_assets": 0,
  "num_clients": 80,
  "num_sdkclients": 10,
  "occupancy_limit": 4,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "assets_waits": {
    "avg": 81.08,
    "max": 149.16,
    "min": 75.74,
    "p95": 47.26,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "clients_waits": {
    "avg": 73.04,
    "max": 141.12,
    "min": 188.3,
    "p95": 39.22,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "created_time": 18.26,
  "modified_time": 60.7,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

