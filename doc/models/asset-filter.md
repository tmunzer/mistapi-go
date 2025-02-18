
# Asset Filter

Asset Filter

*This model accepts additional fields of type interface{}.*

## Structure

`AssetFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | - |
| `Beam` | `*int` | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Disasbled` | `*bool` | Optional | Whether the asset filter is disabled |
| `EddystoneUidNamespace` | `*string` | Optional | Eddystone uid namespace used to filter assets |
| `EddystoneUrl` | `*string` | Optional | Eddystone url used to filter assets |
| `ForSite` | `*bool` | Optional | - |
| `IbeaconMajor` | `*int` | Optional | ibeacon major value used to filter assets |
| `IbeaconUuid` | `*uuid.UUID` | Optional | ibeacon uuid used to filter assets |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `MfgCompanyId` | `*int` | Optional | BLE manufacturing-specific company-id used to filter assets |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `ServiceUuid` | `*uuid.UUID` | Optional | BLE service data uuid used to filter assets |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "eddystone_uid_namespace": "2818e3868dec25629ede",
  "eddystone_url": "https://www.abc.com",
  "ibeacon_major": 13,
  "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "mfg_company_id": 935,
  "name": "Visitor Tags",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "service_uuid": "0000fe6a-0000-1000-8000-0030459b3cfb",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap_mac": "ap_mac0",
  "beam": 124,
  "created_time": 126.18,
  "disasbled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

