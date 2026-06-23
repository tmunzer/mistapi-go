
# Site Setting Config Push Policy

Mist also uses some heuristic rules to prevent destructive configs from being pushed

## Structure

`SiteSettingConfigPushPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoPush` | `*bool` | Optional | Stop any new config from being pushed to the device<br><br>**Default**: `false` |
| `PushWindow` | [`*models.PushPolicyPushWindow`](../../doc/models/push-policy-push-window.md) | Optional | If enabled, new config will only be pushed to device within the specified time window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingConfigPushPolicy := models.SiteSettingConfigPushPolicy{
        NoPush:               models.ToPointer(false),
        PushWindow:           models.ToPointer(models.PushPolicyPushWindow{
            Enabled:              models.ToPointer(false),
            Hours:                models.ToPointer(models.Hours{
                Fri:                  models.ToPointer("fri2"),
                Mon:                  models.ToPointer("mon8"),
                Sat:                  models.ToPointer("sat0"),
                Sun:                  models.ToPointer("sun6"),
                Thu:                  models.ToPointer("thu6"),
            }),
        }),
    }

}
```

