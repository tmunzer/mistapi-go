
# Site Setting Auto Upgrade Esl

Automatic AP ESL firmware upgrade policy. When both firmware and ESL auto-upgrade are enabled, ESL upgrade will be done only after firmware upgrade

## Structure

`SiteSettingAutoUpgradeEsl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowDowngrade` | `*bool` | Optional | If true, it will allow downgrade to a lower version<br><br>**Default**: `false` |
| `CustomVersions` | `map[string]string` | Optional | Custom versions for different models. Property key is the model name (e.g. "AP41") |
| `DayOfWeek` | [`*models.DayOfWeekEnum`](../../doc/models/day-of-week-enum.md) | Optional | enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed` |
| `Enabled` | `*bool` | Optional | Whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported)<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time |
| `Version` | `*string` | Optional | ESL firmware version used for auto-upgrade |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingAutoUpgradeEsl := models.SiteSettingAutoUpgradeEsl{
        AllowDowngrade:       models.ToPointer(false),
        CustomVersions:       map[string]string{
            "AP41": "2.4.6",
            "AP61": "2.5.0",
        },
        DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_ANY),
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
        Version:              models.ToPointer("2.5.0"),
    }

}
```

