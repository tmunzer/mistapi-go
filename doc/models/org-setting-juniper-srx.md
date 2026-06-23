
# Org Setting Juniper Srx

Organization settings for Juniper SRX devices

## Structure

`OrgSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.JuniperSrxAutoUpgrade`](../../doc/models/juniper-srx-auto-upgrade.md) | Optional | SRX firmware auto-upgrade settings applied when a device is first onboarded |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingJuniperSrx := models.OrgSettingJuniperSrx{
        AutoUpgrade:          models.ToPointer(models.JuniperSrxAutoUpgrade{
            CustomVersions:       map[string]string{
                "key0": "custom_versions3",
                "key1": "custom_versions2",
            },
            Enabled:              models.ToPointer(false),
            Snapshot:             models.ToPointer(false),
            Version:              models.ToPointer("version2"),
        }),
    }

}
```

