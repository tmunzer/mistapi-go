
# Response Switch Metrics Active Ports Summary

## Structure

`ResponseSwitchMetricsActivePortsSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.SwitchMetricsActivePortsSummaryDetails`](../../doc/models/switch-metrics-active-ports-summary-details.md) | Optional | - |
| `Score` | `*int` | Optional | - |
| `TotalSwitchCount` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "active_port_count": 136,
    "total_port_count": 42
  },
  "score": 222,
  "total_switch_count": 10
}
```

