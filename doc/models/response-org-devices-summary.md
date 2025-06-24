
# Response Org Devices Summary

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseOrgDevicesSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `*int` | Optional | - |
| `NumGateways` | `*int` | Optional | - |
| `NumMxedges` | `*int` | Optional | - |
| `NumSwitches` | `*int` | Optional | - |
| `NumUnassignedAps` | `*int` | Optional | - |
| `NumUnassignedGateways` | `*int` | Optional | - |
| `NumUnassignedSwitches` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_aps": 172,
  "num_gateways": 80,
  "num_mxedges": 40,
  "num_switches": 98,
  "num_unassigned_aps": 2,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

