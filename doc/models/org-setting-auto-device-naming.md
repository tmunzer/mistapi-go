
# Org Setting Auto Device Naming

Automatic device naming configuration for claimed devices

## Structure

`OrgSettingAutoDeviceNaming`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | Whether automatic device naming is enabled |
| `Rules` | [`models.Optional[[]models.OrgSettingAutoDeviceNamingRule]`](../../doc/models/org-setting-auto-device-naming-rule.md) | Optional | Automatic device naming rules, or null when automatic naming is not configured |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingAutoDeviceNaming := models.OrgSettingAutoDeviceNaming{
        Enable:               models.ToPointer(false),
        Rules:                models.NewOptional(models.ToPointer([]models.OrgSettingAutoDeviceNamingRule{
            models.OrgSettingAutoDeviceNamingRule{
                Expression:           models.ToPointer("expression4"),
                MatchDevice:          models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
                Prefix:               models.ToPointer("prefix6"),
                Src:                  models.ToPointer(models.OrgSettingAutoDeviceNamingRuleSrcEnum_LLDPPORTDESC),
                Suffix:               models.ToPointer("suffix2"),
            },
            models.OrgSettingAutoDeviceNamingRule{
                Expression:           models.ToPointer("expression4"),
                MatchDevice:          models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
                Prefix:               models.ToPointer("prefix6"),
                Src:                  models.ToPointer(models.OrgSettingAutoDeviceNamingRuleSrcEnum_LLDPPORTDESC),
                Suffix:               models.ToPointer("suffix2"),
            },
        })),
    }

}
```

