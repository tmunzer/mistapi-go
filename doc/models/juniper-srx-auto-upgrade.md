
# Juniper Srx Auto Upgrade

SRX firmware auto-upgrade settings applied when a device is first onboarded

## Structure

`JuniperSrxAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SRX Hardware model (e.g. "SRX4600") |
| `Enabled` | `*bool` | Optional | Whether SRX auto-upgrade is enabled for newly onboarded devices<br><br>**Default**: `false` |
| `Snapshot` | `*bool` | Optional | Whether to take a snapshot during the SRX upgrade process<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | Firmware version to deploy (e.g. 23.4R2-S5.5). Optional, used when custom_versions not specified |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    juniperSrxAutoUpgrade := models.JuniperSrxAutoUpgrade{
        CustomVersions:       map[string]string{
            "key0": "custom_versions5",
            "key1": "custom_versions6",
            "key2": "custom_versions7",
        },
        Enabled:              models.ToPointer(false),
        Snapshot:             models.ToPointer(false),
        Version:              models.ToPointer("23.4R2-S5.5"),
    }

}
```

