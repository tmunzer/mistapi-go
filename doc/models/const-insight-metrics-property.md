
# Const Insight Metrics Property

Definition of an insight metric returned by the constants API

## Structure

`ConstInsightMetricsProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ctype` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Description` | `*string` | Optional | Human-readable description of the insight metric |
| `Example` | [`*models.ConstInsightMetricsPropertyExampleAnyOf2`](../../doc/models/containers/const-insight-metrics-property-example-any-of-2.md) | Optional | Example values for an insight metric property, as an array or keyed object |
| `Intervals` | [`map[string]models.ConstInsightMetricsPropertyInterval`](../../doc/models/const-insight-metrics-property-interval.md) | Optional | Property key is the interval (e.g. 10m, 1h, ...) |
| `Keys` | `*interface{}` | Optional | Additional key metadata for an insight metric property |
| `Params` | [`map[string]models.ConstInsightMetricsPropertyParam`](../../doc/models/const-insight-metrics-property-param.md) | Optional | Property key is the parameter name |
| `ReportDurations` | [`map[string]models.ConstInsightMetricsPropertyReportDuration`](../../doc/models/const-insight-metrics-property-report-duration.md) | Optional | Property key is the duration (e.g. 1d, 1w, ...) |
| `ReportScopes` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Scopes` | [`[]models.ConstInsightMetricsPropertyScopeEnum`](../../doc/models/const-insight-metrics-property-scope-enum.md) | Optional | Entity scopes supported by an insight metric property |
| `SleBaselined` | `*bool` | Optional | Whether the insight metric uses an SLE baseline |
| `SleClassifiers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Type` | `*string` | Optional | Metric data type, such as timeseries |
| `Unit` | `*string` | Optional | Measurement unit for values returned by this metric |
| `Values` | `*interface{}` | Optional | Enumerated or structured value metadata for an insight metric property |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constInsightMetricsProperty := models.ConstInsightMetricsProperty{
        Ctype:                []string{
            "ctype4",
        },
        Description:          models.ToPointer("description8"),
        Example:              models.ToPointer(models.ConstInsightMetricsPropertyExampleAnyOf2Container.FromArrayOfConstInsightMetricsPropertyExample([]models.ConstInsightMetricsPropertyExample{
            models.ConstInsightMetricsPropertyExampleContainer.FromNumber(102),
            models.ConstInsightMetricsPropertyExampleContainer.FromNumber(103),
        })),
        Intervals:            map[string]models.ConstInsightMetricsPropertyInterval{
            "key0": models.ConstInsightMetricsPropertyInterval{
                Interval:             models.ToPointer(140),
                MaxAge:               models.ToPointer(96),
            },
            "key1": models.ConstInsightMetricsPropertyInterval{
                Interval:             models.ToPointer(140),
                MaxAge:               models.ToPointer(96),
            },
        },
        Keys:                 models.ToPointer(interface{}("[key1, val1][key2, val2]")),
    }

}
```

