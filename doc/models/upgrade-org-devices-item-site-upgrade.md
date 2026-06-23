
# Upgrade Org Devices Item Site Upgrade

Site upgrade mapping returned for an organization upgrade

## Structure

`UpgradeOrgDevicesItemSiteUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `UpgradeId` | `*uuid.UUID` | Optional | Identifier of the site-level upgrade job |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeOrgDevicesItemSiteUpgrade := models.UpgradeOrgDevicesItemSiteUpgrade{
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        UpgradeId:            models.ToPointer(uuid.MustParse("ebbdbd0b-1bcf-4e55-8a6a-3416049a52b1")),
    }

}
```

