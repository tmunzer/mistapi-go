
# Sle Histogram

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

## Example (as JSON)

```json
{
  "data": [
    {
      "range": [
        7.03,
        7.04
      ],
      "value": 95.62
    }
  ],
  "end": 54.28,
  "metric": "metric2",
  "start": 10.34,
  "x_label": "x_label6",
  "y_label": "y_label8"
}
```

