
# Switch Port Usage Storm Control

Switch storm control. Only if `mode`!=`dynamic`

## Structure

`SwitchPortUsageStormControl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisablePort` | `*bool` | Optional | Whether to disable the port when storm control is triggered<br><br>**Default**: `false` |
| `NoBroadcast` | `*bool` | Optional | Whether to disable storm control on broadcast traffic<br><br>**Default**: `false` |
| `NoMulticast` | `*bool` | Optional | Whether to disable storm control on multicast traffic<br><br>**Default**: `false` |
| `NoRegisteredMulticast` | `*bool` | Optional | Whether to disable storm control on registered multicast traffic<br><br>**Default**: `false` |
| `NoUnknownUnicast` | `*bool` | Optional | Whether to disable storm control on unknown unicast traffic<br><br>**Default**: `false` |
| `Percentage` | `*int` | Optional | Bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth<br><br>**Default**: `80`<br><br>**Constraints**: `>= 0`, `<= 100` |

## Example (as JSON)

```json
{
  "disable_port": false,
  "no_broadcast": false,
  "no_multicast": false,
  "no_registered_multicast": false,
  "no_unknown_unicast": false,
  "percentage": 80
}
```

