
# Nac Portal Sso Role Matching

Mapping rule from an SSO role claim value to a NAC portal role

## Structure

`NacPortalSsoRoleMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Assigned` | `*string` | Optional | NAC portal role assigned when the SSO role value matches |
| `Match` | `*string` | Optional | SSO role value to match from the SAML assertion |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacPortalSsoRoleMatching := models.NacPortalSsoRoleMatching{
        Assigned:             models.ToPointer("user"),
        Match:                models.ToPointer("Student"),
    }

}
```

