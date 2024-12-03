
# Const Mxedge Model

*This model accepts additional fields of type interface{}.*

## Structure

`ConstMxedgeModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomPorts` | `*bool` | Optional | - |
| `Display` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Ports` | [`map[string]models.ConstMxedgeModelPort`](../../doc/models/const-mxedge-model-port.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "display": "X10",
  "model": "ME-X10",
  "custom_ports": false,
  "ports": {
    "key0": {
      "display": "display2",
      "speed": 14,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "display": "display2",
      "speed": 14,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

