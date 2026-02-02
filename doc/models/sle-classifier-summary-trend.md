
# Sle Classifier Summary Trend

## Structure

`SleClassifierSummaryTrend`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | [`models.SleTrendClassifier`](../../doc/models/sle-trend-classifier.md) | Required | - |
| `End` | `float64` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "classifier": {
    "interval": 24.52,
    "name": "name4",
    "samples": {
      "degraded": [
        170.03
      ],
      "duration": [
        249.08,
        249.09,
        249.1
      ],
      "total": [
        144.95
      ]
    },
    "x_label": "x_label0",
    "y_label": "y_label2"
  },
  "end": 97.58,
  "metric": "metric8",
  "start": 53.64
}
```

