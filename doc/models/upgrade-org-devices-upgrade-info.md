
# Upgrade Org Devices Upgrade Info

Site-level upgrade job details within an organization upgrade

## Structure

`UpgradeOrgDevicesUpgradeInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `StartTime` | `*int` | Optional | Epoch timestamp, in seconds, when the site-level upgrade starts |
| `Status` | [`*models.UpgradeDeviceStatusEnum`](../../doc/models/upgrade-device-status-enum.md) | Optional | status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued` |
| `Targets` | [`*models.UpgradeDevicesTargets`](../../doc/models/upgrade-devices-targets.md) | Optional, Read-only | Read-only device target lists grouped by upgrade status |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeOrgDevicesUpgradeInfo := models.UpgradeOrgDevicesUpgradeInfo{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        StartTime:            models.ToPointer(1717658765),
        Status:               models.ToPointer(models.UpgradeDeviceStatusEnum_CANCELLED),
        Targets:              models.ToPointer(models.UpgradeDevicesTargets{
            DownloadRequested:    []string{
                "download_requested6",
            },
            Downloaded:           []string{
                "downloaded0",
                "downloaded1",
                "downloaded2",
            },
            Downloading:          []string{
                "downloading6",
            },
            Failed:               []string{
                "failed6",
            },
            RebootInProgress:     []string{
                "reboot_in_progress3",
                "reboot_in_progress4",
            },
        }),
    }

}
```

