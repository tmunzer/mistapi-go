
# Webhook Location Event

*This model accepts additional fields of type interface{}.*

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
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "00001f3e-0000-0000-0000-000000000000",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 76,
  "type": "type4",
  "x": 84.88,
  "y": 46.4,
  "battery_voltage": 80,
  "eddystone_uid_instance": "eddystone_uid_instance0",
  "eddystone_uid_namespace": "eddystone_uid_namespace8",
  "eddystone_url_url": "eddystone_url_url8",
  "ibeacon_major": 124,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

