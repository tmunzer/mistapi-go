
# Org Setting Auto Assignment Rule

Automatic assignment rule used by org settings

## Structure

`OrgSettingAutoAssignmentRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateNewSiteIfNeeded` | `*bool` | Optional | If `src`==`geoip`. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this<br><br>**Default**: `false` |
| `Expression` | `models.Optional[string]` | Optional | If `src`==`name`, `src`==`lldp_system_name`, `src`==`dns_suffix`  <br>"[0:3]"            // "abcdef" -> "abc"  <br>"split(.)[1]"      // "a.b.c" -> "b"  <br>"split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"' |
| `GatewaytemplateId` | `*string` | Optional | If `src`==`geoip` and `create_new_site_if_needed`==`true`. If a gateway template is desired for this newly created site |
| `MatchCountry` | `*string` | Optional | If `src`==`geoip`, country or region value that must match the device location |
| `MatchDeviceType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `MatchModel` | `*string` | Optional | Optional additional device model filter for this assignment rule |
| `Model` | `*string` | Optional | If `src`==`model`, device model value to match |
| `Prefix` | `models.Optional[string]` | Optional | If `src`==`name`, prefix that must be present in the device name |
| `Src` | [`models.OrgSettingAutoSiteAssignmentSrcEnum`](../../doc/models/org-setting-auto-site-assignment-src-enum.md) | Required | enum: `ext_ip`, `dns_suffix`, `geoip`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet` |
| `Subnet` | `*string` | Optional | If `src`==`subnet` or `ext_ip`==`ext_ip` |
| `Suffix` | `models.Optional[string]` | Optional | If `src`==`name`, suffix that must be present in the device name |
| `Value` | `*string` | Optional | If<br><br>* `src`==`ext_ip`, `src`==`subnet` or `src`==`model`, the site name<br>* `src`==`geoip`: site name for the device to be assigned to (\"city\" / \"city+country\" / ...)" |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingAutoAssignmentRule := models.OrgSettingAutoAssignmentRule{
        CreateNewSiteIfNeeded: models.ToPointer(false),
        Expression:            models.NewOptional(models.ToPointer("split(.)[1]")),
        GatewaytemplateId:     models.ToPointer("gatewaytemplate_id0"),
        MatchCountry:          models.ToPointer("match_country8"),
        MatchDeviceType:       models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
        Prefix:                models.NewOptional(models.ToPointer("XX-")),
        Src:                   models.OrgSettingAutoSiteAssignmentSrcEnum_GEOIP,
        Suffix:                models.NewOptional(models.ToPointer("-YY")),
    }

}
```

