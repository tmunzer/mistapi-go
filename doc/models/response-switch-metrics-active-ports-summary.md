
# Response Switch Metrics Active Ports Summary

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSwitchMetricsActivePortsSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.SwitchMetricsActivePortsSummaryDetails`](../../doc/models/switch-metrics-active-ports-summary-details.md) | Optional | - |
| `Score` | `*int` | Optional | - |
| `TotalSwitchCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "active_port_count": 136,
    "total_port_count": 42,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 222,
  "total_switch_count": 10,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

