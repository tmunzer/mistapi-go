
# Const Device Ap Extios

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceApExtios`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultDir` | [`*models.ConstDeviceApExtiosDefaultDirEnum`](../../doc/models/const-device-ap-extios-default-dir-enum.md) | Optional | enum: `IN`, `OUT` |
| `Input` | `*bool` | Optional | - |
| `Output` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "default_dir": "IN",
  "input": false,
  "output": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

