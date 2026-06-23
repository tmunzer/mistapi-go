
# Response Site Device Upgrade

Site device upgrade job details

## Structure

`ResponseSiteDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanaryPhases` | `[]int` | Optional | phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. |
| `CurrentPhase` | `*int` | Optional | Current canary or rrm phase in progress |
| `EnableP2p` | `*bool` | Optional | Whether to allow local AP-to-AP FW upgrade<br><br>**Default**: `false` |
| `Force` | `*bool` | Optional | Whether to force upgrade when requested version is same as running version |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `MaxFailurePercentage` | `*int` | Optional | Percentage of failures allowed |
| `MaxFailures` | `[]int` | Optional | If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `P2pClusterSize` | `*int` | Optional | size to split the devices for p2p<br><br>**Default**: `10` |
| `P2pParallelism` | `*int` | Optional | number of parallel p2p download batches to create<br><br>**Default**: `1` |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `StartTime` | `*int` | Optional | Firmware download start time in epoch |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued` |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)<br><br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | Version to upgrade to<br><br>**Constraints**: *Minimum Length*: `1` |
| `Targets` | [`*models.UpgradeDevicesTargets`](../../doc/models/upgrade-devices-targets.md) | Optional, Read-only | Read-only device target lists grouped by upgrade status |
| `UpgradePlan` | `map[string][]string` | Optional | If `strategy`!=`big_bang`, a dictionary of phase number to devices part of that phase |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSiteDeviceUpgrade := models.ResponseSiteDeviceUpgrade{
        CanaryPhases:         nil,
        CurrentPhase:         models.ToPointer(204),
        EnableP2p:            models.ToPointer(false),
        Force:                models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MaxFailurePercentage: models.ToPointer(82),
        P2pClusterSize:       models.ToPointer(10),
        P2pParallelism:       models.ToPointer(1),
        Strategy:             models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
    }

}
```

