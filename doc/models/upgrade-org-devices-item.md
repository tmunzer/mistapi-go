
# Upgrade Org Devices Item

Organization device upgrade job returned by upgrade APIs

## Structure

`UpgradeOrgDevicesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `SiteUpgrades` | [`[]models.UpgradeOrgDevicesItemSiteUpgrade`](../../doc/models/upgrade-org-devices-item-site-upgrade.md) | Optional | Site-level upgrade jobs created for an organization upgrade |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeOrgDevicesItem := models.UpgradeOrgDevicesItem{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        SiteUpgrades:         []models.UpgradeOrgDevicesItemSiteUpgrade{
            models.UpgradeOrgDevicesItemSiteUpgrade{
                SiteId:               models.ToPointer(uuid.MustParse("00000ca4-0000-0000-0000-000000000000")),
                UpgradeId:            models.ToPointer(uuid.MustParse("000010fe-0000-0000-0000-000000000000")),
            },
            models.UpgradeOrgDevicesItemSiteUpgrade{
                SiteId:               models.ToPointer(uuid.MustParse("00000ca4-0000-0000-0000-000000000000")),
                UpgradeId:            models.ToPointer(uuid.MustParse("000010fe-0000-0000-0000-000000000000")),
            },
        },
    }

}
```

