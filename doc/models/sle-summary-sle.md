
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
  "x_label": "x_label6",
  "y_label": "y_label8"
}
```

