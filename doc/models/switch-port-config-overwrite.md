
# Switch Port Config Overwrite

Switch port config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchPortConfigOverwrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | - |
| `Disabled` | `*bool` | Optional | Whether the port is disabled<br><br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortUsageDuplexOverwriteEnum`](../../doc/models/switch-port-usage-duplex-overwrite-enum.md) | Optional | Link connection mode. enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `MacLimit` | [`*models.SwitchPortUsageMacLimitOverwrite`](../../doc/models/containers/switch-port-usage-mac-limit-overwrite.md) | Optional | Max number of mac addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform) |
| `PoeDisabled` | `*bool` | Optional | Whether PoE capabilities are disabled for a port<br><br>**Default**: `false` |
| `PortNetwork` | `*string` | Optional | Native network/vlan for untagged traffic |
| `Speed` | [`*models.SwitchPortUsageSpeedOverwriteEnum`](../../doc/models/switch-port-usage-speed-overwrite-enum.md) | Optional | Port Speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "duplex": "auto",
  "poe_disabled": false,
  "speed": "auto",
  "description": "description8",
  "mac_limit": 246,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

