
# Switch Port Local Usage Storm Control

Switch storm control

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchPortLocalUsageStormControl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoBroadcast` | `*bool` | Optional | Whether to disable storm control on broadcast traffic<br>**Default**: `false` |
| `NoMulticast` | `*bool` | Optional | Whether to disable storm control on multicast traffic<br>**Default**: `false` |
| `NoRegisteredMulticast` | `*bool` | Optional | Whether to disable storm control on registered multicast traffic<br>**Default**: `false` |
| `NoUnknownUnicast` | `*bool` | Optional | Whether to disable storm control on unknown unicast traffic<br>**Default**: `false` |
| `Percentage` | `*int` | Optional | Bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth<br>**Default**: `80`<br>**Constraints**: `>= 0`, `<= 100` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "no_broadcast": false,
  "no_multicast": false,
  "no_registered_multicast": false,
  "no_unknown_unicast": false,
  "percentage": 80,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

