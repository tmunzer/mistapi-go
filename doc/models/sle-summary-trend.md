
# Sle Summary Trend

## Structure

`SleSummaryTrend`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifiers` | [`[]models.SleTrendClassifier`](../../doc/models/sle-trend-classifier.md) | Required | **Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | - |
| `Sle` | [`models.SleSummarySle`](../../doc/models/sle-summary-sle.md) | Required | - |
| `Start` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "classifiers": [
    {
      "interval": 217.3,
      "name": "name2",
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
      "x_label": "x_label8",
      "y_label": "y_label0"
    }
  ],
  "end": 69.36,
  "sle": {
    "interval": 227.76,
    "name": "name8",
    "samples": {
      "degraded": [
        170.03
      ],
      "total": [
        144.95
      ],
      "value": [
        13.46,
        13.47,
        13.48
      ]
    },
    "x_label": "x_label4",
    "y_label": "y_label6"
  },
  "start": 25.42
}
```

