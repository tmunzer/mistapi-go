
# Sle Impacted Interfaces Interface

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedInterfacesInterface`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | - |
| `Duration` | `*float64` | Optional | - |
| `InterfaceName` | `*string` | Optional | - |
| `SwitchMac` | `*string` | Optional | - |
| `SwitchName` | `*string` | Optional | - |
| `Total` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "degraded": 225.26,
  "duration": 98.32,
  "interface_name": "interface_name2",
  "switch_mac": "switch_mac4",
  "switch_name": "switch_name8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

