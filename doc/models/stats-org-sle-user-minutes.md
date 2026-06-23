
# Stats Org Sle User Minutes

User-minute totals for an organization SLE summary

## Structure

`StatsOrgSleUserMinutes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ok` | `float64` | Required | User minutes that met the SLE target |
| `Total` | `float64` | Required | Observed total user minutes for the SLE path |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsOrgSleUserMinutes := models.StatsOrgSleUserMinutes{
        Ok:                   float64(60.88),
        Total:                float64(196.58),
    }

}
```

