
# Org Setting Installer

Organization-level permissions and grace period for installer workflows

## Structure

`OrgSettingInstaller`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowAllDevices` | `*bool` | Optional | Whether installers may work with all eligible devices |
| `AllowAllSites` | `*bool` | Optional | Whether installers may work with all sites |
| `ExtraSiteIds` | `[]uuid.UUID` | Optional | Additional site IDs that installers may access |
| `GracePeriod` | `*int` | Optional | Grace period, in days, during which installers can modify recent sites or devices |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSettingInstaller := models.OrgSettingInstaller{
        AllowAllDevices:      models.ToPointer(false),
        AllowAllSites:        models.ToPointer(false),
        ExtraSiteIds:         []uuid.UUID{
            uuid.MustParse("000010a0-0000-0000-0000-000000000000"),
        },
        GracePeriod:          models.ToPointer(38),
    }

}
```

