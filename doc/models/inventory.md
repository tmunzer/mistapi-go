
# Inventory

*This model accepts additional fields of type interface{}.*

## Structure

`Inventory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | Only if `type`==`switch` or `type`==`gateway`, whether the switch/gateway is adopted |
| `Connected` | `*bool` | Optional | Whether the device is connected |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[string]` | Optional | Deviceprofile id if assigned, null if not assigned |
| `Hostname` | `*string` | Optional | Hostname reported by the device |
| `HwRev` | `*string` | Optional | Device hardware revision number |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Jsi` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | Device MAC address |
| `Magic` | `*string` | Optional | Device claim code |
| `Model` | `*string` | Optional | Device model |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Device name if configured |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | Device serial |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sku` | `*string` | Optional | Device stock keeping unit |
| `Type` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br>**Default**: `"ap"` |
| `VcMac` | `*string` | Optional | If `type`==`switch` and device part of a Virtual Chassis, MAC Address of the Virtual Chassis. if `type`==`gateway` and device part of a Clust, MAC Address of the Cluster |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "ap",
  "adopted": false,
  "connected": false,
  "created_time": 98.34,
  "deviceprofile_id": "deviceprofile_id4",
  "hostname": "hostname0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

