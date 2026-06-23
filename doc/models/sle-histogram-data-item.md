
# Sle Histogram Data Item

Single bucket in an SLE histogram

## Structure

`SleHistogramDataItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Range` | `[]float64` | Optional | Lower and upper boundary values for an SLE histogram bucket |
| `Value` | `float64` | Required | Measured amount for this histogram bucket |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleHistogramDataItem := models.SleHistogramDataItem{
        Range:                []float64{
            float64(31.17),
        },
        Value:                float64(57.42),
    }

}
```

