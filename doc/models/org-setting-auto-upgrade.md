
# Org Setting Auto Upgrade

Organization-wide AP automatic firmware upgrade policy

## Structure

`OrgSettingAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Per-AP-model firmware versions or channels used for auto-upgrade. Property key is the AP model name (e.g. "AP41"), value is the firmware version or release channel (e.g. `stable`) |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether AP auto-upgrade is enabled. Note that Mist may auto-upgrade APs if the running version is no longer supported.<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` or HH:MM (24-hour format). Upgrade will happen within up to 1 hour from this time. |
| `Version` | [`*models.SiteAutoUpgradeVersionEnum`](../../doc/models/site-auto-upgrade-version-enum.md) | Optional | desired version. enum: `beta`, `custom`, `stable`<br><br>**Default**: `"stable"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingAutoUpgrade := models.OrgSettingAutoUpgrade{
        CustomVersions:       map[string]string{
            "AP21": "stable",
            "AP41": "0.1.5135",
            "AP61": "0.1.7215",
        },
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_SUN),
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
        Version:              models.ToPointer(models.SiteAutoUpgradeVersionEnum_BETA),
    }

}
```

