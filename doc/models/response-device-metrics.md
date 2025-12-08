
# Response Device Metrics

## Structure

`ResponseDeviceMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.ResponseDeviceMetricsResultsItems`](../../doc/models/containers/response-device-metrics-results-items.md) | Required | - |
| `Rt` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 172,
  "interval": 82,
  "limit": 2,
  "page": 112,
  "results": [
    "String0",
    "String1"
  ],
  "rt": [
    "rt4",
    "rt5"
  ],
  "start": 130
}
```

