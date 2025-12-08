
# Vc Port

*This model accepts additional fields of type interface{}.*

## Structure

`VcPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mode` | [`*models.VcPortModeEnum`](../../doc/models/vc-port-mode-enum.md) | Optional | enum: `network`, `vcp-higig`, `vcp-hgoe` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mode": "vcp-higig",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

