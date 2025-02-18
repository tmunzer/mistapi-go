
# Sle Classifier

*This model accepts additional fields of type interface{}.*

## Structure

`SleClassifier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Impact` | [`models.SleClassifierImpact`](../../doc/models/sle-classifier-impact.md) | Required | - |
| `Interval` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Samples` | [`*models.SleClassifierSamples`](../../doc/models/sle-classifier-samples.md) | Optional | - |
| `XLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
  "interval": 0.34,
  "name": "name8",
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
  "x_label": "x_label6",
  "y_label": "y_label6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

