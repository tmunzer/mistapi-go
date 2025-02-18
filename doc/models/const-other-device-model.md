
# Const Other Device Model

*This model accepts additional fields of type interface{}.*

## Structure

`ConstOtherDeviceModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `VendorModelId` | `*string` | Optional | - |
| `Display` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Vendor` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "_vendor_model_id": "_vendor_model_id2",
  "display": "display4",
  "model": "model0",
  "type": "type8",
  "vendor": "vendor8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

