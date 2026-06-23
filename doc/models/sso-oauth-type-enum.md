
# Sso Oauth Type Enum

if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`

## Enumeration

`SsoOauthTypeEnum`

## Fields

| Name |
|  --- |
| `azure` |
| `azure-gov` |
| `okta` |
| `ping_identity` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoOauthType := models.SsoOauthTypeEnum_AZURE

}
```

