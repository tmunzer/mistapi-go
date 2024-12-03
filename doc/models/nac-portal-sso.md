
# Nac Portal Sso

*This model accepts additional fields of type interface{}.*

## Structure

`NacPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IdpCert` | `*string` | Optional | - |
| `IdpSignAlgo` | [`*models.NacPortalSsoIdpSignAlgoEnum`](../../doc/models/nac-portal-sso-idp-sign-algo-enum.md) | Optional | Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`.<br>**Default**: `"sha256"` |
| `IdpSsoUrl` | `*string` | Optional | - |
| `Issuer` | `*string` | Optional | - |
| `NameidFormat` | `*string` | Optional | - |
| `SsoRoleMatching` | [`[]models.NacPortalSsoRoleMatching`](../../doc/models/nac-portal-sso-role-matching.md) | Optional | - |
| `UseSsoRoleForCert` | `*bool` | Optional | if it's desired to inject a role into Cert's Subject (so it can be used later on in policy) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "idp_cert": "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
  "idp_sign_algo": "sha256",
  "idp_sso_url": "https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130",
  "issuer": "https://app.onelogin.com/saml/metadata/138130",
  "nameid_format": "email",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

