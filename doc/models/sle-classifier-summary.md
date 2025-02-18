
# Sle Classifier Summary

*This model accepts additional fields of type interface{}.*

## Structure

`SleClassifierSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | [`models.SleClassifier`](../../doc/models/sle-classifier.md) | Required | - |
| `End` | `float64` | Required | - |
| `Failures` | `[]interface{}` | Required | - |
| `Impact` | [`models.SleClassifierSummaryImpact`](../../doc/models/sle-classifier-summary-impact.md) | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": {
    "impact": {
      "num_aps": 250.52,
      "num_users": 4.56,
      "total_aps": 243.84,
      "total_users": 4.14,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "interval": 24.52,
    "name": "name4",
    "samples": {
      "degraded": [
        183.49
      ],
      "duration": [
        249.08,
        249.09,
        249.1
      ],
      "total": [
        158.41
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "x_label": "x_label0",
    "y_label": "y_label2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "end": 15.18,
  "failures": [
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
    "total_users": 4.14,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "metric": "metric8",
  "start": 227.24,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

