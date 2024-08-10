
# Sle Summary Sle

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

## Example (as JSON)

```json
{
  "interval": 247.38,
  "name": "name0",
  "samples": {
    "degraded": [
      183.49
    ],
    "total": [
      158.41
    ],
    "value": [
      26.92,
      26.93,
      26.94
    ]
  },
  "x_label": "x_label6",
  "y_label": "y_label8"
}
```

