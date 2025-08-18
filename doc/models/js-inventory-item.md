
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
| `WarrantyType` | [`*models.JsiWarrantyTypeEnum`](../../doc/models/jsi-warranty-type-enum.md) | Optional | Warranty type for Juniper Support Insight (JSI) devices. The warranty type<br>is used to determine the support level and duration of the warranty for the<br>device. enum:<br><br>* WTY00001: Standard Hardware Warranty<br>* WTY00002: Enhanced Hardware Warranty<br>* WTY00003: Dead On Arrival Warranty<br>* WTY00004: Limited Lifetime Warranty<br>* WTY00005: Software Warranty<br>* WTY00006: Limited Lifetime Warranty for WLA<br>* WTY00007: Warranty-JCPO EOL (DOA Not Included)<br>* WTY00008: MIST Enhanced Hardware Warranty<br>* WTY00009: MIST Standard Warranty<br>* WTY00099: Determine Lifetime warranty |
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

