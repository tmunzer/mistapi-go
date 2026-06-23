
# Site Setting Rtsa

Managed mobility and asset tracking settings

## Structure

`SiteSettingRtsa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppWaking` | `*bool` | Optional | Whether app wake-up support is enabled for managed mobility<br><br>**Default**: `false` |
| `DisableDeadReckoning` | `*bool` | Optional | Whether dead reckoning is disabled for managed mobility |
| `DisablePressureSensor` | `*bool` | Optional | Whether pressure sensor use is disabled for managed mobility<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether managed mobility features are enabled |
| `TrackAsset` | `*bool` | Optional | Whether BLE asset tracking is enabled for managed mobility<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingRtsa := models.SiteSettingRtsa{
        AppWaking:             models.ToPointer(false),
        DisableDeadReckoning:  models.ToPointer(false),
        DisablePressureSensor: models.ToPointer(false),
        Enabled:               models.ToPointer(false),
        TrackAsset:            models.ToPointer(false),
    }

}
```

