
# Wlan Bonjour Service Properties

Bonjour service discovery settings for one advertised service

## Structure

`WlanBonjourServiceProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableLocal` | `*bool` | Optional | Whether to prevent wireless clients to discover bonjour devices on the same WLAN<br><br>**Default**: `false` |
| `RadiusGroups` | `[]string` | Optional | Optional, if the service is further restricted for certain RADIUS groups |
| `Scope` | [`*models.WlanBonjourServicePropertiesScopeEnum`](../../doc/models/wlan-bonjour-service-properties-scope-enum.md) | Optional | how bonjour services should be discovered for the same WLAN. enum: `same_ap`, `same_map`, `same_site`<br><br>**Default**: `"same_site"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanBonjourServiceProperties := models.WlanBonjourServiceProperties{
        DisableLocal:         models.ToPointer(false),
        RadiusGroups:         []string{
            "radius_groups9",
            "radius_groups8",
            "radius_groups7",
        },
        Scope:                models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum_SAMESITE),
    }

}
```

