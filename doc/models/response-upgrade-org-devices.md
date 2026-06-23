
# Response Upgrade Org Devices

Organization AP upgrade job details

## Structure

`ResponseUpgradeOrgDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableP2p` | `*bool` | Optional | Whether to allow local AP-to-AP FW upgrade |
| `Force` | `*bool` | Optional | Whether to force upgrade when requested version is same as running version |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Strategy` | [`*models.UpgradeDeviceStrategyEnum`](../../doc/models/upgrade-device-strategy-enum.md) | Optional | enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)<br><br>**Default**: `"big_bang"` |
| `TargetVersion` | `*string` | Optional | Firmware version targeted by the organization AP upgrade |
| `Upgrades` | [`[]models.UpgradeOrgDevicesUpgrade`](../../doc/models/upgrade-org-devices-upgrade.md) | Optional | Per-site upgrade entries created for an organization device upgrade |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseUpgradeOrgDevices := models.ResponseUpgradeOrgDevices{
        EnableP2p:            models.ToPointer(false),
        Force:                models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Strategy:             models.ToPointer(models.UpgradeDeviceStrategyEnum_BIGBANG),
        TargetVersion:        models.ToPointer("0.14.29411"),
    }

}
```

