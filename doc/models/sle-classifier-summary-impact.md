
# Sle Classifier Summary Impact

AP and user impact counts for a classifier summary window

## Structure

`SleClassifierSummaryImpact`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `float64` | Required | Number of APs affected by degraded SLE experience |
| `NumUsers` | `float64` | Required | Number of users affected by degraded SLE experience |
| `TotalAps` | `float64` | Required | Total number of APs considered in the summary window |
| `TotalUsers` | `float64` | Required | Total number of users considered in the summary window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifierSummaryImpact := models.SleClassifierSummaryImpact{
        NumAps:               float64(231),
        NumUsers:             float64(24.08),
        TotalAps:             float64(7.36),
        TotalUsers:           float64(240.62),
    }

}
```

