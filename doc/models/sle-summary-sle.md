
# Sle Summary Sle

*This model accepts additional fields of type interface{}.*

## Structure

`SleSummarySle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `float64` | Required | - |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Samples` | [`models.SleSummarySleSamples`](../../doc/models/sle-summary-sle-samples.md) | Required | - |
| `XLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "interval": 247.38,
  "name": "name0",
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
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "x_label": "x_label6",
  "y_label": "y_label8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

