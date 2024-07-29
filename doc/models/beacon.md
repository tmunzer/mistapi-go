
# Beacon

Beacon

## Structure

`Beacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `EddystoneInstance` | `*string` | Optional | Eddystone-UID instance (6 bytes) in hexstring format |
| `EddystoneNamespace` | `*string` | Optional | Eddystone-UID namespace (10 bytes) in hexstring format |
| `EddystoneUrl` | `*string` | Optional | Eddystone-URL url |
| `ForSite` | `*bool` | Optional | - |
| `IbeaconMajor` | `*int` | Optional | bluetooth tag major |
| `IbeaconMinor` | `*int` | Optional | bluetooth tag minor |
| `IbeaconUuid` | `*uuid.UUID` | Optional | bluetooth tag UUID |
| `Id` | `*uuid.UUID` | Optional | - |
| `Mac` | `*string` | Optional | optiona, MAC of the beacon, currently used only to identify battery voltage |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | name / label of the device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Power` | `*int` | Optional | in dBm<br>**Default**: `-12`<br>**Constraints**: `>= -12`, `<= 100` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | [`*models.BeaconTypeEnum`](../../doc/models/beacon-type-enum.md) | Optional | enum: `eddystone-uid`, `eddystone-url`, `ibeacon`<br>**Default**: `"eddystone-uid"` |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "power": -12,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "eddystone-uid",
  "created_time": 161.28,
  "eddystone_instance": "eddystone_instance0",
  "eddystone_namespace": "eddystone_namespace4",
  "eddystone_url": "eddystone_url4",
  "for_site": false
}
```

