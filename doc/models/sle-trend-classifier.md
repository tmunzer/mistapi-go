
# Sle Trend Classifier

## Structure

`SleTrendClassifier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Samples` | [`*models.SleClassifierSamples`](../../doc/models/sle-classifier-samples.md) | Optional | - |
| `XLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "interval": 251.98,
  "name": "name0",
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
  "x_label": "x_label6",
  "y_label": "y_label8"
}
```

