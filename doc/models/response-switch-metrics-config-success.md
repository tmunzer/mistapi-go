
# Response Switch Metrics Config Success

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSwitchMetricsConfigSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.ResponseSwitchMetricsConfigSuccessDetails`](../../doc/models/response-switch-metrics-config-success-details.md) | Optional | - |
| `Score` | `*int` | Optional | - |
| `TotalSwitchCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "details": {
    "config_success_count": 166,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "score": 170,
  "total_switch_count": 42,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

