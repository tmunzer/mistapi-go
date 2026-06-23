
# Stats Org

Organization statistics summary returned by the org stats endpoint

## Structure

`StatsOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `uuid.UUID` | Required | Organization-level alarm template identifier used as the default for sites |
| `AllowMist` | `bool` | Required | Whether Mist support access is allowed for this organization |
| `CreatedTime` | `float64` | Required, Read-only | When the object has been created, in epoch |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `float64` | Required, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `uuid.UUID` | Required, Read-only | Managed service provider identifier |
| `Name` | `string` | Required | Display name of the organization |
| `NumDevices` | `int` | Required | Total number of devices in the organization |
| `NumDevicesConnected` | `int` | Required | Number of organization devices currently connected to Mist |
| `NumDevicesDisconnected` | `int` | Required | Number of organization devices currently disconnected from Mist |
| `NumInventory` | `int` | Required | Number of devices in the organization's inventory |
| `NumSites` | `int` | Required | Number of sites in the organization |
| `OrggroupIds` | `[]uuid.UUID` | Required | Organization group identifiers associated with an organization stats record |
| `SessionExpiry` | `int64` | Required | Admin session lifetime for the organization, in minutes |
| `Sle` | [`[]models.StatsOrgSle`](../../doc/models/stats-org-sle.md) | Required | Service level expectation summaries for an organization<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsOrg := models.StatsOrg{
        AlarmtemplateId:        uuid.MustParse("000010fa-0000-0000-0000-000000000000"),
        AllowMist:              false,
        CreatedTime:            float64(63.06),
        Id:                     uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        ModifiedTime:           float64(15.9),
        MspId:                  uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909"),
        Name:                   "name6",
        NumDevices:             170,
        NumDevicesConnected:    230,
        NumDevicesDisconnected: 18,
        NumInventory:           136,
        NumSites:               4,
        OrggroupIds:            []uuid.UUID{
            uuid.MustParse("0000106e-0000-0000-0000-000000000000"),
            uuid.MustParse("0000106f-0000-0000-0000-000000000000"),
        },
        SessionExpiry:          int64(252),
        Sle:                    []models.StatsOrgSle{
            models.StatsOrgSle{
                Path:                 "path2",
                UserMinutes:          models.ToPointer(models.StatsOrgSleUserMinutes{
                    Ok:                   float64(13.84),
                    Total:                float64(12.38),
                }),
            },
        },
    }

}
```

