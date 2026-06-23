
# Response Site Device Upgrades Item

Site device upgrade job summary

## Structure

`ResponseSiteDeviceUpgradesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Counts` | [`*models.UpgradeSiteDevicesCounts`](../../doc/models/upgrade-site-devices-counts.md) | Optional, Read-only | Device counts grouped by status for a site upgrade job |
| `CurrentPhase` | `*int` | Optional | Current canary or rrm phase in progress |
| `EnableP2p` | `*bool` | Optional | Whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | Whether to force upgrade when requested version is same as running version |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `MaxFailurePercentage` | `*int` | Optional | Percentage of failures allowed |
| `MaxFailures` | `[]int` | Optional | If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used |
| `RebootAt` | `*int` | Optional | reboot start time in epoch |
| `StartTime` | `*int` | Optional | Firmware download start time in epoch |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued` |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)<br><br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | Version to upgrade to<br><br>**Constraints**: *Minimum Length*: `1` |
| `UpgradePlan` | `*interface{}` | Optional | a dictionary of rrm phase number to devices part of that phase |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSiteDeviceUpgradesItem := models.ResponseSiteDeviceUpgradesItem{
        Counts:               models.ToPointer(models.UpgradeSiteDevicesCounts{
            DownloadRequested:    models.ToPointer(138),
            Downloaded:           models.ToPointer(70),
            Failed:               models.ToPointer(166),
            RebootInProgress:     models.ToPointer(88),
            Rebooted:             models.ToPointer(76),
        }),
        CurrentPhase:         models.ToPointer(150),
        EnableP2p:            models.ToPointer(false),
        Force:                models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MaxFailurePercentage: models.ToPointer(28),
        Strategy:             models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
    }

}
```

