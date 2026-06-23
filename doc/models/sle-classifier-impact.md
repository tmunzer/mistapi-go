
# Sle Classifier Impact

AP and user impact counts for an SLE classifier

## Structure

`SleClassifierImpact`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumAps` | `float64` | Required | Number of APs affected by degraded SLE experience |
| `NumUsers` | `float64` | Required | Number of users affected by degraded SLE experience |
| `TotalAps` | `float64` | Required | Total number of APs considered for the classifier |
| `TotalUsers` | `float64` | Required | Total number of users considered for the classifier |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleClassifierImpact := models.SleClassifierImpact{
        NumAps:               float64(165.58),
        NumUsers:             float64(89.5),
        TotalAps:             float64(72.78),
        TotalUsers:           float64(175.2),
    }

}
```

