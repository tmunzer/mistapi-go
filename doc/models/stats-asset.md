
# Stats Asset

Asset statistics

## Structure

`StatsAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*float64` | Optional | battery voltage, in mV |
| `Beam` | `*int` | Optional | - |
| `DeviceName` | `*string` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `EddystoneUidInstance` | `*string` | Optional | - |
| `EddystoneUidNamespace` | `*string` | Optional | - |
| `EddystoneUrlUrl` | `*string` | Optional | - |
| `IbeaconMajor` | `*int` | Optional | - |
| `IbeaconMinor` | `*int` | Optional | - |
| `IbeaconUuid` | `*uuid.UUID` | Optional | - |
| `LastSeen` | `*float64` | Optional | last seen timestamp |
| `Mac` | `string` | Required | bluetooth MAC |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `Name` | `*string` | Optional | name / label of the device |
| `Rssi` | `*int` | Optional | - |
| `Rssizones` | [`[]models.AssetRssiZone`](../../doc/models/asset-rssi-zone.md) | Optional | only send this for individual asset stat |
| `Temperatur` | `*float64` | Optional | - |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |
| `Zones` | [`[]models.AssetZone`](../../doc/models/asset-zone.md) | Optional | only send this for individual asset stat |

## Example (as JSON)

```json
{
  "battery_voltage": 2970.0,
  "beam": 6,
  "device_name": "a",
  "duration": 120,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_url": "ttps://www.abc.com",
  "ibeacon_major": 12,
  "ibeacon_minor": 138,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "last_seen": 1428939600,
  "mac": "6fa474be7ae5",
  "map_id": "c45be59f-854d-4ef7-b782-dcd6309c84a9",
  "name": "6fa474be7ae5",
  "rssi": -60,
  "temperatur": 23,
  "x": 280.199181403102,
  "y": 420.298772104653
}
```

