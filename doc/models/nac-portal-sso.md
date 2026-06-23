
# Nac Portal Sso

SAML SSO configuration for a NAC portal

## Structure

`NacPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IdpCert` | `*string` | Optional | Identity provider certificate used to verify signed SAML responses |
| `IdpSignAlgo` | [`*models.NacPortalSsoIdpSignAlgoEnum`](../../doc/models/nac-portal-sso-idp-sign-algo-enum.md) | Optional | Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`.<br><br>**Default**: `"sha256"` |
| `IdpSsoUrl` | `*string` | Optional | Identity provider Single Sign-On URL for SAML authentication |
| `Issuer` | `*string` | Optional | Identity provider issuer URL for SAML authentication |
| `NameidFormat` | `*string` | Optional | SAML NameID format expected from the identity provider |
| `SsoRoleMatching` | [`[]models.NacPortalSsoRoleMatching`](../../doc/models/nac-portal-sso-role-matching.md) | Optional | List of SSO role mapping rules for the NAC portal |
| `UseSsoRoleForCert` | `*bool` | Optional | Whether to include the matched SSO role in the issued certificate subject for later policy matching |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacPortalSso := models.NacPortalSso{
        IdpCert:              models.ToPointer("-----BEGIN CERTIFICATE-----\\n<redacted>\\n-----END CERTIFICATE-----"),
        IdpSignAlgo:          models.ToPointer(models.NacPortalSsoIdpSignAlgoEnum_SHA256),
        IdpSsoUrl:            models.ToPointer("https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130"),
        Issuer:               models.ToPointer("https://app.onelogin.com/saml/metadata/138130"),
        NameidFormat:         models.ToPointer("email"),
    }

}
```

