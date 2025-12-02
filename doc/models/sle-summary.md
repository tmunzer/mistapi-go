
# Sle Summary

## Structure

`SleSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifiers` | [`[]models.SleClassifier`](../../doc/models/sle-classifier.md) | Required | **Constraints**: *Unique Items Required* |
| `End` | `float64` | Required | - |
| `Events` | `[]interface{}` | Required | - |
| `Impact` | [`models.SleSummaryImpact`](../../doc/models/sle-summary-impact.md) | Required | - |
| `Sle` | [`models.SleSummarySle`](../../doc/models/sle-summary-sle.md) | Required | - |
| `Start` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "classifiers": [
    {
      "impact": {
        "num_aps": 250.52,
        "num_users": 4.56,
        "total_aps": 243.84,
        "total_users": 4.14
      },
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
  "end": 255.54,
  "events": [
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    },
    {
      "key1": "val1",
      "key2": "val2"
    }
  ],
  "impact": {
    "num_aps": 250.52,
    "num_users": 4.56,
    "total_aps": 243.84,
    "total_users": 4.14
  },
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
  "start": 211.6
}
```

