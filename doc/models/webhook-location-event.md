
# Webhook Location Event

## Structure

`WebhookLocationEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryVoltage` | `*int` | Optional | - |
| `EddystoneUidInstance` | `*string` | Optional | - |
| `EddystoneUidNamespace` | `*string` | Optional | - |
| `EddystoneUrlUrl` | `*string` | Optional | - |
| `IbeaconMajor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconMinor` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `IbeaconUuid` | `models.Optional[uuid.UUID]` | Optional | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | - |
| `MapId` | `uuid.UUID` | Required | Map id |
| `MfgCompanyId` | `*int` | Optional | Optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | Optional, BLE manufacturing data in hex byte-string format (ie "112233AABBCC") |
| `Name` | `*string` | Optional | Name of the client, may be empty |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Type` | `string` | Required | - |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | Optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `float64` | Required | x, in meter |
| `Y` | `float64` | Required | y, in meter |

## Example (as JSON)

```json
{
  "ibeacon_major": 1234,
  "ibeacon_minor": 1234,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "00001f3e-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 8.44,
  "type": "type4",
  "x": 84.88,
  "y": 46.4,
  "battery_voltage": 80,
  "eddystone_uid_instance": "eddystone_uid_instance0",
  "eddystone_uid_namespace": "eddystone_uid_namespace8",
  "eddystone_url_url": "eddystone_url_url8"
}
```

