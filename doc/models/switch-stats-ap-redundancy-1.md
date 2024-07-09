
# Switch Stats Ap Redundancy 1

## Structure

`SwitchStatsApRedundancy1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Modules` | [`*models.ApRedundancyModules`](../../doc/models/ap-redundancy-modules.md) | Optional | for a VC / stacked switches. |
| `NumAps` | `*int` | Optional | - |
| `NumApsWithSwitchRedundancy` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "num_aps": 15,
  "num_aps_with_switch_redundancy": 8,
  "modules": {
    "num_aps": 2,
    "num_aps_with_switch_redundancy": 254
  }
}
```

