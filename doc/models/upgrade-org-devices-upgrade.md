
# Upgrade Org Devices Upgrade

Upgrade entry for a single site within an organization upgrade job

## Structure

`UpgradeOrgDevicesUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Upgrade` | [`*models.UpgradeOrgDevicesUpgradeInfo`](../../doc/models/upgrade-org-devices-upgrade-info.md) | Optional | Site-level upgrade job details within an organization upgrade |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeOrgDevicesUpgrade := models.UpgradeOrgDevicesUpgrade{
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Upgrade:              models.ToPointer(models.UpgradeOrgDevicesUpgradeInfo{
            Id:                   models.ToPointer(uuid.MustParse("000016ac-0000-0000-0000-000000000000")),
            StartTime:            models.ToPointer(228),
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
        }),
    }

}
```

