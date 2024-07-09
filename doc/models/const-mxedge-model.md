
# Const Mxedge Model

## Structure

`ConstMxedgeModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomPorts` | `*bool` | Optional | - |
| `Display` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Ports` | [`map[string]models.ConstMxedgeModelPort`](../../doc/models/const-mxedge-model-port.md) | Optional | - |

## Example (as JSON)

```json
{
  "display": "X10",
  "model": "ME-X10",
  "custom_ports": false,
  "ports": {
    "key0": {
      "display": "display2",
      "speed": 14
    },
    "key1": {
      "display": "display2",
      "speed": 14
    }
  }
}
```

