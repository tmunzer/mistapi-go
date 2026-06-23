
# Stats Cluster Config Redundancy Group Info Item

Redundancy group monitoring entry for a gateway cluster

## Structure

`StatsClusterConfigRedundancyGroupInfoItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | Redundancy group identifier reported by the gateway |
| `MonitoringFailure` | `*string` | Optional | Detected monitoring failure condition for the redundancy group |
| `Threshold` | `*int` | Optional | Configured monitoring threshold for the redundancy group |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsClusterConfigRedundancyGroupInfoItem := models.StatsClusterConfigRedundancyGroupInfoItem{
        Id:                   models.ToPointer(0),
        MonitoringFailure:    models.ToPointer("MonitoringFailure0"),
        Threshold:            models.ToPointer(194),
    }

}
```

