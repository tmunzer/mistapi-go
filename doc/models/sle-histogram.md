
# Sle Histogram

*This model accepts additional fields of type interface{}.*

## Structure

`SleHistogram`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`[]models.SleHistogramDataItem`](../../doc/models/sle-histogram-data-item.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `End` | `float64` | Required | - |
| `Metric` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Start` | `float64` | Required | - |
| `XLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `YLabel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "data": [
    {
      "range": [
        7.03,
        7.04
      ],
      "value": 95.62,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "end": 37.86,
  "metric": "metric0",
  "start": 249.92,
  "x_label": "x_label6",
  "y_label": "y_label6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

