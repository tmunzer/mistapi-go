
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
| `IbeaconMajor` | `*int` | Optional | - |
| `IbeaconMinor` | `*int` | Optional | - |
| `IbeaconUuid` | `*uuid.UUID` | Optional | - |
| `Id` | `uuid.UUID` | Required | unique id of the client (a client would have different id for different org) |
| `Mac` | `*string` | Optional | - |
| `MapId` | `uuid.UUID` | Required | map id |
| `MfgCompanyId` | `*int` | Optional | optional, BLE manufacturing company ID |
| `MfgData` | `*string` | Optional | optional, BLE manufacturing data in hex byte-string format (ie "112233AABBCC") |
| `Name` | `*string` | Optional | name of the client, may be empty |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `int` | Required | timestamp of the event, epoch |
| `Type` | `string` | Required | - |
| `WifiBeaconExtendedInfo` | [`[]models.WifiBeaconExtendedInfoItems`](../../doc/models/wifi-beacon-extended-info-items.md) | Optional | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload |
| `X` | `float64` | Required | x, in meter |
| `Y` | `float64` | Required | y, in meter |

## Example (as JSON)

```json
{
  "id": "00001b7a-0000-0000-0000-000000000000",
  "map_id": "000002b2-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 24,
  "type": "type4",
  "x": 132.2,
  "y": 7.48,
  "battery_voltage": 28,
  "eddystone_uid_instance": "eddystone_uid_instance8",
  "eddystone_uid_namespace": "eddystone_uid_namespace0",
  "eddystone_url_url": "eddystone_url_url0",
  "ibeacon_major": 72
}
```

