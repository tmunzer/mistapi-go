
# Org Setting Auto Device Naming Rule

Automatic device naming rule

## Structure

`OrgSettingAutoDeviceNamingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Expression` | `*string` | Optional | "[0:3]"            // "abcdef" -> "abc"  <br>"split(.)[1]"      // "a.b.c" -> "b"  <br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"' |
| `MatchDevice` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `Prefix` | `*string` | Optional | Text prepended to the generated device name |
| `Src` | [`*models.OrgSettingAutoDeviceNamingRuleSrcEnum`](../../doc/models/org-setting-auto-device-naming-rule-src-enum.md) | Optional | Device attribute used to generate the name. enum: `lldp_port_desc`, `mac` |
| `Suffix` | `*string` | Optional | Text appended to the generated device name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingAutoDeviceNamingRule := models.OrgSettingAutoDeviceNamingRule{
        Expression:           models.ToPointer("split(.)[1]"),
        MatchDevice:          models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
        Prefix:               models.ToPointer("prefix0"),
        Src:                  models.ToPointer(models.OrgSettingAutoDeviceNamingRuleSrcEnum_LLDPPORTDESC),
        Suffix:               models.ToPointer("suffix8"),
    }

}
```

