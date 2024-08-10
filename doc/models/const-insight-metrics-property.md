
# Const Insight Metrics Property

## Structure

`ConstInsightMetricsProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ctype` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Description` | `*string` | Optional | - |
| `Example` | [`[]models.ConstInsightMetricsPropertyExample`](../../doc/models/containers/const-insight-metrics-property-example.md) | Optional | - |
| `Intervals` | [`map[string]models.ConstInsightMetricsPropertyInterval`](../../doc/models/const-insight-metrics-property-interval.md) | Optional | Property key is the interval (e.g. 10m, 1h, ...) |
| `Keys` | `*interface{}` | Optional | - |
| `Params` | [`map[string]models.ConstInsightMetricsPropertyParam`](../../doc/models/const-insight-metrics-property-param.md) | Optional | Property key is the parameter name |
| `ReportDuration` | [`map[string]models.ConstInsightMetricsPropertyReportDuration`](../../doc/models/const-insight-metrics-property-report-duration.md) | Optional | Property key is the duration (e.g. 1d, 1w, ...) |
| `ReportScopes` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Scopes` | [`[]models.ConstInsightMetricsPropertyScopeEnum`](../../doc/models/const-insight-metrics-property-scope-enum.md) | Optional | - |
| `SleBaselined` | `*bool` | Optional | - |
| `SleClassifiers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Type` | `*string` | Optional | - |
| `Unit` | `*string` | Optional | - |
| `Values` | `*interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ctype": [
    "ctype0",
    "ctype9",
    "ctype8"
  ],
  "description": "description8",
  "example": [
    60
  ],
  "intervals": {
    "key0": {
      "interval": 140,
      "max_age": 96
    },
    "key1": {
      "interval": 140,
      "max_age": 96
    },
    "key2": {
      "interval": 140,
      "max_age": 96
    }
  },
  "keys": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

