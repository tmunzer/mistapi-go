
# Insight Metrics

## Structure

`InsightMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Interval` | `int` | Required | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.InsightMetricsResultsItem`](../../doc/models/containers/insight-metrics-results-item.md) | Required | Results depends on the `metric` - some return numbers (e.g. bytes, ap-count), others return objects<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 16,
  "interval": 182,
  "limit": 154,
  "results": [
    234.97
  ],
  "start": 230
}
```

