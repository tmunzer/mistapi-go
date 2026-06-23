
# Response Mxedge Upgrade

Mist Edge upgrade details response

## Structure

`ResponseMxedgeUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | Upgrade channel used to select Mist Edge package versions<br><br>**Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.MxedgeUpgradeResponseCounts`](../../doc/models/mxedge-upgrade-response-counts.md) | Required | Counts of Mist Edge upgrades by current status |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Status` | `string` | Required | Current status of the Mist Edge upgrade<br><br>**Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | Rollout strategy used for the Mist Edge upgrade<br><br>**Constraints**: *Minimum Length*: `1` |
| `Versions` | `interface{}` | Required | Per-service target versions for this Mist Edge upgrade |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMxedgeUpgrade := models.ResponseMxedgeUpgrade{
        Channel:              "channel2",
        Counts:               models.MxedgeUpgradeResponseCounts{
            Failed:               166,
            Queued:               234,
            Success:              90,
            Upgrading:            112,
        },
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Status:               "status6",
        Strategy:             "strategy2",
        Versions:             interface{}("[key1, val1][key2, val2]"),
    }

}
```

