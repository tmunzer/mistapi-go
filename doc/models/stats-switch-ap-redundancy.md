
# Stats Switch Ap Redundancy

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSwitchApRedundancy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Modules` | [`map[string]models.StatsSwitchApRedundancyModule`](../../doc/models/stats-switch-ap-redundancy-module.md) | Optional | For a VC / stacked switches. |
| `NumAps` | `*int` | Optional | - |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_aps": 15,
  "num_aps_with_switch_redundancy": 8,
  "modules": {
    "key0": {
      "num_aps": 2,
      "num_aps_with_switch_redundancy": 254,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "num_aps": 2,
      "num_aps_with_switch_redundancy": 254,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "num_aps": 2,
      "num_aps_with_switch_redundancy": 254,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

