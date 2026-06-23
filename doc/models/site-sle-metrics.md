
# Site Sle Metrics

SLE metrics available for a site

## Structure

`SiteSleMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Supported` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSleMetrics := models.SiteSleMetrics{
        Enabled:              []string{
            "enabled3",
            "enabled4",
            "enabled5",
        },
        Supported:            []string{
            "supported8",
            "supported9",
            "supported0",
        },
    }

}
```

