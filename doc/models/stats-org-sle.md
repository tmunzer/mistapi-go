
# Stats Org Sle

Organization SLE summary for a service path

## Structure

`StatsOrgSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Path` | `string` | Required | SLE path represented by this organization summary |
| `UserMinutes` | [`*models.StatsOrgSleUserMinutes`](../../doc/models/stats-org-sle-user-minutes.md) | Optional | User-minute totals for an organization SLE summary |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsOrgSle := models.StatsOrgSle{
        Path:                 "path4",
        UserMinutes:          models.ToPointer(models.StatsOrgSleUserMinutes{
            Ok:                   float64(13.84),
            Total:                float64(12.38),
        }),
    }

}
```

