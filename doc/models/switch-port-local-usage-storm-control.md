
# Switch Port Local Usage Storm Control

Switch storm control

## Structure

`SwitchPortLocalUsageStormControl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoBroadcast` | `*bool` | Optional | whether to disable storm control on broadcast traffic<br>**Default**: `false` |
| `NoMulticast` | `*bool` | Optional | whether to disable storm control on multicast traffic<br>**Default**: `false` |
| `NoRegisteredMulticast` | `*bool` | Optional | whether to disable storm control on registered multicast traffic<br>**Default**: `false` |
| `NoUnknownUnicast` | `*bool` | Optional | whether to disable storm control on unknown unicast traffic<br>**Default**: `false` |
| `Percentage` | `*int` | Optional | bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth<br>**Default**: `80`<br>**Constraints**: `>= 0`, `<= 100` |

## Example (as JSON)

```json
{
  "no_broadcast": false,
  "no_multicast": false,
  "no_registered_multicast": false,
  "no_unknown_unicast": false,
  "percentage": 80
}
```

