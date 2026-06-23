
# Sle Threshold

Threshold metadata and configured value for an SLE metric

*This model accepts additional fields of type interface{}.*

## Structure

`SleThreshold`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*float64` | Optional, Read-only | Built-in threshold value for this SLE metric |
| `Direction` | `*string` | Optional, Read-only | Comparison direction used for SLE threshold evaluation<br><br>**Constraints**: *Minimum Length*: `1` |
| `Maximum` | `*float64` | Optional | Highest allowed threshold value for this SLE metric |
| `Metric` | `*string` | Optional, Read-only | SLE metric name associated with this threshold<br><br>**Constraints**: *Minimum Length*: `1` |
| `Minimum` | `*float64` | Optional | Lowest allowed threshold value for this SLE metric |
| `Threshold` | `*string` | Optional, Read-only | Configured threshold value for this SLE metric<br><br>**Constraints**: *Minimum Length*: `1` |
| `Units` | `*string` | Optional, Read-only | Measurement units for this SLE threshold<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleThreshold := models.SleThreshold{
        Default:              models.ToPointer(float64(89.22)),
        Direction:            models.ToPointer("direction8"),
        Maximum:              models.ToPointer(float64(224.96)),
        Metric:               models.ToPointer("metric6"),
        Minimum:              models.ToPointer(float64(196.84)),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

