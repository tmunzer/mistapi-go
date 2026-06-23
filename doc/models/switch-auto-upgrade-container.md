
# Switch Auto Upgrade Container

Wrapper for switch firmware auto-upgrade settings

## Structure

`SwitchAutoUpgradeContainer`

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
    switchAutoUpgradeContainer := models.SwitchAutoUpgradeContainer{
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

