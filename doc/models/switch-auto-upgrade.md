
# Switch Auto Upgrade

Switch firmware auto-upgrade settings

## Structure

`SwitchAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Custom version to be used. The Property Key is the switch hardware and the property value is the firmware version |
| `Enabled` | `*bool` | Optional | Whether switch auto-upgrade is enabled |
| `Snapshot` | `*bool` | Optional | Whether to create a recovery snapshot during the upgrade process<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchAutoUpgrade := models.SwitchAutoUpgrade{
        CustomVersions:       map[string]string{
            "QFX5120-32C": "23.4R2-S2.1",
            "QFX5130-32CD": "23.4R2-S2.3",
        },
        Enabled:              models.ToPointer(false),
        Snapshot:             models.ToPointer(false),
    }

}
```

