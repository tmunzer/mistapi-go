
# Stats Mxedge Service Stat

Runtime and package state for one Mist Edge service

## Structure

`StatsMxedgeServiceStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | External IP from ep-terminator’s point of view. valid only for service having its own cloud connection |
| `LastSeen` | `*float64` | Optional | Cloud Unix time, in seconds, when stats were last seen for this service, or the latest service last_seen time for whole-system records |
| `PackageState` | `*string` | Optional | Installation state of the Mist Edge service package |
| `PackageVersion` | `*string` | Optional | Installed version of the Mist Edge service package |
| `RunningState` | `*string` | Optional | Runtime state reported by the Mist Edge service |
| `Uptime` | `*int` | Optional | Elapsed running time reported by the Mist Edge service, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeServiceStat := models.StatsMxedgeServiceStat{
        ExtIp:                models.ToPointer("ext_ip8"),
        LastSeen:             models.ToPointer(float64(239.4)),
        PackageState:         models.ToPointer("package_state4"),
        PackageVersion:       models.ToPointer("package_version2"),
        RunningState:         models.ToPointer("running_state2"),
    }

}
```

