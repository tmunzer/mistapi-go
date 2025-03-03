
# Js Inventory Item

*This model accepts additional fields of type interface{}.*

## Structure

`JsInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EolTime` | `*int` | Optional | End of life time |
| `EosTime` | `*int` | Optional | End of support time |
| `Model` | `*string` | Optional | Model of the install base inventory |
| `Serial` | `*string` | Optional | Serial Number of the inventory |
| `Sku` | `*string` | Optional | Serviceable device stock |
| `Type` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `WarrantyType` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "eol_time": 58,
  "eos_time": 104,
  "model": "model8",
  "serial": "serial0",
  "sku": "sku6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

