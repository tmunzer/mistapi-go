
# Sso Idp Type Enum

SSO IDP Type:

* For Admin SSO, enum: `saml`
* For NAC SSO, enum: `ldap`, `mxedge_proxy`, `oauth`, `openroaming`

## Enumeration

`SsoIdpTypeEnum`

## Fields

| Name |
|  --- |
| `ldap` |
| `mxedge_proxy` |
| `oauth` |
| `saml` |
| `openroaming` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoIdpType := models.SsoIdpTypeEnum_MXEDGEPROXY

}
```

