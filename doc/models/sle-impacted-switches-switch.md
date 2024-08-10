
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
  "degraded": 238.34,
  "duration": 111.4,
  "interface": [
    "interface0",
    "interface9",
    "interface8"
  ],
  "name": "name4",
  "switch_mac": "switch_mac2"
}
```

