
# Vpn Path Selection

Only if `type`==`hub_spoke`

*This model accepts additional fields of type interface{}.*

## Structure

`VpnPathSelection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Strategy` | [`*models.VpnPathSelectionStrategyEnum`](../../doc/models/vpn-path-selection-strategy-enum.md) | Optional | enum: `disabled`, `simple`, `manual`<br>**Default**: `"disabled"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "strategy": "disabled",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

