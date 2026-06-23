
# Org Setting Marvis

Organization settings for Marvis automation

## Structure

`OrgSettingMarvis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableProactiveMonitoring` | `*bool` | Optional | Disable proactive monitoring in Marvis. NOTE: support access must be enabled for the org (`allow_mist`=`true`) for proactive monitoring to function.<br><br>**Default**: `false` |
| `SelfDriving` | [`*models.MarvisSelfDriving`](../../doc/models/marvis-self-driving.md) | Optional | Self-driving network automation settings per domain |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingMarvis := models.OrgSettingMarvis{
        DisableProactiveMonitoring: models.ToPointer(false),
        SelfDriving:                models.ToPointer(models.MarvisSelfDriving{
            Wan:                  models.ToPointer(models.MarvisSelfDrivingDomain{
                Enabled:              models.ToPointer(false),
            }),
            Wired:                models.ToPointer(models.MarvisSelfDrivingDomain{
                Enabled:              models.ToPointer(false),
            }),
            Wireless:             models.ToPointer(models.MarvisSelfDrivingDomain{
                Enabled:              models.ToPointer(false),
            }),
        }),
    }

}
```

