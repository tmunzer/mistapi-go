
# Wlan App Limit

Bandwidth limiting for apps (applies to up/down)

## Structure

`WlanAppLimit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `map[string]int` | Optional | Map from app key to bandwidth in kbps.<br>Property key is the app key, defined in Get Application List |
| `Enabled` | `*bool` | Optional | Whether application bandwidth limits are enabled for this WLAN<br><br>**Default**: `false` |
| `WxtagIds` | `map[string]int` | Optional | Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps. Property key is the `wxtag_id` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAppLimit := models.WlanAppLimit{
        Apps:                 map[string]int{
            "dropbox": 300,
            "netflix": 60,
        },
        Enabled:              models.ToPointer(false),
        WxtagIds:             map[string]int{
            "f99862d9-2726-931f-7559-3dfdf5d070d3": 30,
        },
    }

}
```

