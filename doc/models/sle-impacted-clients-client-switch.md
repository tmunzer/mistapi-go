
# Sle Impacted Clients Client Switch

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedClientsClientSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SwitchMac` | `*string` | Optional | - |
| `SwitchName` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "interfaces": [
    "interfaces5",
    "interfaces6",
    "interfaces7"
  ],
  "switch_mac": "switch_mac2",
  "switch_name": "switch_name6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

