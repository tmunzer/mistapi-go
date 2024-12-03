
# Sle Impacted Chassis Chassis Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedChassisChassisItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chassis` | `*string` | Optional | - |
| `Degraded` | `*float64` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `Role` | `*string` | Optional | - |
| `SwitchMac` | `*string` | Optional | - |
| `SwitchName` | `*string` | Optional | - |
| `Total` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "chassis": "chassis8",
  "degraded": 159.04,
  "duration": 32.1,
  "role": "role8",
  "switch_mac": "switch_mac2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

