
# Mxedge Upgrade Response Counts

Counts of Mist Edge upgrades by current status

## Structure

`MxedgeUpgradeResponseCounts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `int` | Required | Number of Mist Edge upgrades that failed |
| `Queued` | `int` | Required | Number of Mist Edge upgrades waiting to run |
| `Success` | `int` | Required | Number of Mist Edge upgrades completed successfully |
| `Upgrading` | `int` | Required | Number of Mist Edge upgrades currently in progress |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeUpgradeResponseCounts := models.MxedgeUpgradeResponseCounts{
        Failed:               142,
        Queued:               46,
        Success:              114,
        Upgrading:            136,
    }

}
```

