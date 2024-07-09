
# Response Device Metrics

## Structure

`ResponseDeviceMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Results` | [`[]models.ResponseDeviceMetricsResults`](../../doc/models/containers/response-device-metrics-results.md) | Required | This is Array of a container for one-of cases. |
| `Rt` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 78,
  "interval": 244,
  "results": [
    "String0"
  ],
  "rt": [
    "rt6"
  ],
  "start": 36
}
```

