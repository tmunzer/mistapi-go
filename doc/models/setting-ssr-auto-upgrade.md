
# Setting Ssr Auto Upgrade

Automatic firmware upgrade settings applied when an SSR device is first onboarded

## Structure

`SettingSsrAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SSR model (e.g. "SSR130"). |
| `Enabled` | `*bool` | Optional | Whether SSR auto-upgrade is enabled for newly onboarded devices<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | Firmware version to deploy (e.g. 6.3.0-107.r1). Optional, used when custom_versions not specified |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    settingSsrAutoUpgrade := models.SettingSsrAutoUpgrade{
        Channel:              models.ToPointer(models.SsrUpgradeChannelEnum_STABLE),
        CustomVersions:       map[string]string{
            "key0": "custom_versions3",
            "key1": "custom_versions2",
        },
        Enabled:              models.ToPointer(false),
        Version:              models.ToPointer("6.3.0-107.r1"),
    }

}
```

