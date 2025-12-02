
# Stats Asset

Asset statistics

## Structure

`StatsAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*float64` | Optional | Battery voltage, in mV |
| `Beam` | `*int` | Optional | - |
| `DeviceName` | `*string` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `EddystoneUidInstance` | `*string` | Optional | - |
| `EddystoneUidNamespace` | `*string` | Optional | - |
| `EddystoneUrlUrl` | `*string` | Optional | - |
| `IbeaconMajor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `string` | Required | Bluetooth MAC |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `Name` | `*string` | Optional | Name / label of the device |
| `Rssi` | `*int` | Optional | - |
| `Rssizones` | [`[]models.AssetRssiZone`](../../doc/models/asset-rssi-zone.md) | Optional | Only send this for individual asset stat |
| `Temperature` | `*float64` | Optional | - |
| `X` | `*float64` | Optional | X in pixel |
| `Y` | `*float64` | Optional | Y in pixel |
| `Zones` | [`[]models.AssetZone`](../../doc/models/asset-zone.md) | Optional | Only send this for individual asset stat |

## Example (as JSON)

```json
{
  "battery_voltage": 2970.0,
  "beam": 6,
  "device_name": "a",
  "duration": 120,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_url": "https://www.abc.com",
  "ibeacon_major": 1234,
  "ibeacon_minor": 1234,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "last_seen": 1470417522,
  "mac": "6fa474be7ae5",
  "map_id": "c45be59f-854d-4ef7-b782-dcd6309c84a9",
  "name": "6fa474be7ae5",
  "rssi": -60,
  "temperature": 23,
  "x": 280.19918140310193,
  "y": 420.2987721046529
}
```

