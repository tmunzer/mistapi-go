
# Beacon

Beacon

*This model accepts additional fields of type interface{}.*

## Structure

`Beacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `EddystoneInstance` | `*string` | Optional | Eddystone-UID instance (6 bytes) in hexstring format |
| `EddystoneNamespace` | `*string` | Optional | Eddystone-UID namespace (10 bytes) in hexstring format |
| `EddystoneUrl` | `*string` | Optional | Eddystone-URL url |
| `ForSite` | `*bool` | Optional | - |
| `IbeaconMajor` | `*int` | Optional | Bluetooth tag major |
| `IbeaconMinor` | `*int` | Optional | Bluetooth tag minor |
| `IbeaconUuid` | `*uuid.UUID` | Optional | Bluetooth tag UUID |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | Optional, MAC of the beacon, currently used only to identify battery voltage |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Name / label of the device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Power` | `*int` | Optional | In dBm<br><br>**Default**: `-12`<br><br>**Constraints**: `>= -12`, `<= 100` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | [`*models.BeaconTypeEnum`](../../doc/models/beacon-type-enum.md) | Optional | enum: `eddystone-uid`, `eddystone-url`, `ibeacon`<br><br>**Default**: `"eddystone-uid"` |
| `X` | `*float64` | Optional | X in pixel |
| `Y` | `*float64` | Optional | Y in pixel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "power": -12,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "eddystone-uid",
  "created_time": 53.64,
  "eddystone_instance": "eddystone_instance6",
  "eddystone_namespace": "eddystone_namespace0",
  "eddystone_url": "eddystone_url8",
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

