
# Response Device Metrics

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Results` | [`[]models.ResponseDeviceMetricsResultsItems`](../../doc/models/containers/response-device-metrics-results-items.md) | Required | - |
| `Rt` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 172,
  "interval": 82,
  "results": [
    "String0",
    "String1"
  ],
  "rt": [
    "rt4",
    "rt5"
  ],
  "start": 130,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

