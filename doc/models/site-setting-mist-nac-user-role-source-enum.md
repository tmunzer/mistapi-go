
# Site Setting Mist Nac User Role Source Enum

Source of the Mist NAC user role sent to firewall gateways. enum: `idp_role`, `radius_group`, `none`

## Enumeration

`SiteSettingMistNacUserRoleSourceEnum`

## Fields

| Name |
|  --- |
| `idp_role` |
| `radius_group` |
| `none` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingMistNacUserRoleSource := models.SiteSettingMistNacUserRoleSourceEnum_IDPROLE

}
```

