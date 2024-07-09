
# Ap Redundancy

## Structure

`ApRedundancy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Modules` | [`map[string]models.ApRedundancyModule`](../../doc/models/ap-redundancy-module.md) | Optional | Property key is the node id |
| `NumAps` | `*int` | Optional | - |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "num_aps": 15,
  "num_aps_with_switch_redundancy": 8,
  "modules": {
    "key0": {
      "num_aps": 2,
      "num_aps_with_switch_redundancy": 254
    },
    "key1": {
      "num_aps": 2,
      "num_aps_with_switch_redundancy": 254
    }
  }
}
```

