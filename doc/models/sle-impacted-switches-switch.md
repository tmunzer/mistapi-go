
# Sle Impacted Switches Switch

## Structure

`SleImpactedSwitchesSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `Interface` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Name` | `*string` | Optional | - |
| `SwitchMac` | `*string` | Optional | - |
| `SwitchModel` | `*string` | Optional | - |
| `SwitchVersion` | `*string` | Optional | - |
| `Total` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 156.2,
  "duration": 29.26,
  "interface": [
    "interface8",
    "interface9",
    "interface0"
  ],
  "name": "name0",
  "switch_mac": "switch_mac8"
}
```

