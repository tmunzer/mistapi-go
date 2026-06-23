
# Stats Ap Auto Upgrade

Auto-upgrade status for an AP

## Structure

`StatsApAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lastcheck` | `models.Optional[int64]` | Optional, Read-only | Time when the AP last checked for auto-upgrade, in epoch seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApAutoUpgrade := models.StatsApAutoUpgrade{
        Lastcheck:            models.NewOptional(models.ToPointer(int64(1720594762))),
    }

}
```

