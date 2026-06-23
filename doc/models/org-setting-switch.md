
# Org Setting Switch

Configuration defaults for switches in this organization

## Structure

`OrgSettingSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SwitchAutoUpgrade`](../../doc/models/switch-auto-upgrade.md) | Optional | Switch firmware auto-upgrade settings |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingSwitch := models.OrgSettingSwitch{
        AutoUpgrade:          models.ToPointer(models.SwitchAutoUpgrade{
            CustomVersions:       map[string]string{
                "key0": "custom_versions3",
                "key1": "custom_versions2",
            },
            Enabled:              models.ToPointer(false),
            Snapshot:             models.ToPointer(false),
        }),
    }

}
```

