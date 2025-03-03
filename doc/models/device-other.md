
# Device Other

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceOther`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceMac` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `State` | `*string` | Optional | - |
| `Vendor` | `*string` | Optional | - |
| `VendorApiId` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 33.6,
  "device_mac": "device_mac4",
  "mac": "mac4",
  "model": "model8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

