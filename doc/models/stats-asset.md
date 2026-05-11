
# Stats Asset

Asset statistics

## Structure

`StatsAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ttl` | `*int` | Optional | Time-to-live in seconds; how long this asset data is valid in cache |
| `BatteryVoltage` | `*float64` | Optional | Battery voltage, in mV |
| `Beam` | `*int` | Optional | - |
| `By` | `*string` | Optional | Source type |
| `DeviceId` | `*uuid.UUID` | Optional | Device ID of the loudest AP |
| `DeviceName` | `*string` | Optional | - |
| `Duration` | `*int` | Optional | - |
| `EddystoneUidInstance` | `*string` | Optional | - |
| `EddystoneUidNamespace` | `*string` | Optional | - |
| `EddystoneUrlUrl` | `*string` | Optional | - |
| `IbeaconMajor` | `models.Optional[int]` | Optional | Major number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | Minor number for iBeacon<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `string` | Required | Bluetooth MAC |
| `Manufacture` | `*string` | Optional | Manufacturer name resolved from company ID |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `MfgCompanyId` | `*int` | Optional | BLE manufacturer company ID from advertisement |
| `MfgData` | `*string` | Optional | Manufacturer-specific data (hex encoded) |
| `Name` | `*string` | Optional | Name / label of the device |
| `Rssi` | `*int` | Optional | Signal strength (RSSI) of the loudest AP in dBm |
| `Rssizones` | [`[]models.AssetRssiZone`](../../doc/models/asset-rssi-zone.md) | Optional | Only send this for individual asset stat |
| `ServicePackets` | [`[]models.StatsAssetServicePacket`](../../doc/models/stats-asset-service-packet.md) | Optional | List of all service data advertisements (maximum length of 10)<br><br>**Constraints**: *Maximum Items*: `10` |
| `Temperature` | `*float64` | Optional | - |
| `X` | `*float64` | Optional | X in pixel |
| `Y` | `*float64` | Optional | Y in pixel |
| `Zones` | [`[]models.AssetZone`](../../doc/models/asset-zone.md) | Optional | Only send this for individual asset stat |

## Example (as JSON)

```json
{
  "battery_voltage": 2970.0,
  "beam": 6,
  "by": "asset",
  "device_id": "00000000-0000-0000-1000-5c5b35000001",
  "device_name": "a",
  "duration": 120,
  "eddystone_uid_instance": "5c5b35000001",
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url_url": "https://www.abc.com",
  "ibeacon_major": 1234,
  "ibeacon_minor": 1234,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 1470417522,
  "mac": "6fa474be7ae5",
  "map_id": "c45be59f-854d-4ef7-b782-dcd6309c84a9",
  "mfg_company_id": 935,
  "mfg_data": "648520a1020000",
  "name": "6fa474be7ae5",
  "rssi": -60,
  "temperature": 23,
  "x": 280.19918140310193,
  "y": 420.2987721046529,
  "_ttl": 190
}
```

