
# Psk Portal Sso

Single sign-on settings used when `auth`==`sso`

## Structure

`PskPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedRoles` | `[]string` | Optional | SSO roles allowed to access the PSK portal; if empty, any role is permitted |
| `IdpCert` | `*string` | Optional | Identity provider signing certificate in PEM format |
| `IdpSignAlgo` | [`*models.PskPortalSsoIdpSignAlgoEnum`](../../doc/models/psk-portal-sso-idp-sign-algo-enum.md) | Optional | Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`. enum: `sha1`, `sha256`, `sha384`, `sha512`<br><br>**Default**: `"sha256"` |
| `IdpSsoUrl` | `*string` | Optional | Identity provider SSO URL for SAML login |
| `Issuer` | `*string` | Optional | SAML issuer value for the identity provider |
| `NameidFormat` | `*string` | Optional | SAML NameID format used for the portal login |
| `RoleMapping` | `map[string]string` | Optional | Maps PSK portal role names to SSO attribute values |
| `UseSsoRoleForPskRole` | `*bool` | Optional | Whether to use SSO role mapping for the PSK role and ignore the portal-level `role` field |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pskPortalSso := models.PskPortalSso{
        AllowedRoles:         []string{
            "allowed_roles5",
            "allowed_roles6",
            "allowed_roles7",
        },
        IdpCert:              models.ToPointer("idp_cert0"),
        IdpSignAlgo:          models.ToPointer(models.PskPortalSsoIdpSignAlgoEnum_SHA256),
        IdpSsoUrl:            models.ToPointer("idp_sso_url6"),
        Issuer:               models.ToPointer("issuer6"),
    }

}
```

