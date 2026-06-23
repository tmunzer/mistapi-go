
# Sle Summary Impact

AP and user impact counts for an SLE summary window

## Structure

`SleSummaryImpact`

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
    sleSummaryImpact := models.SleSummaryImpact{
        NumAps:               float64(199.82),
        NumUsers:             float64(55.26),
        TotalAps:             float64(38.54),
        TotalUsers:           float64(209.44),
    }

}
```

