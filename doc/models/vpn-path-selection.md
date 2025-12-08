
# Vpn Path Selection

Only if `type`==`hub_spoke`

## Structure

`VpnPathSelection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Strategy` | [`*models.VpnPathSelectionStrategyEnum`](../../doc/models/vpn-path-selection-strategy-enum.md) | Optional | enum: `disabled`, `simple`, `manual`<br><br>**Default**: `"disabled"` |

## Example (as JSON)

```json
{
  "strategy": "disabled"
}
```

