
# Psk Portal Sso

if `auth`==`sso`

## Structure

`PskPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedRoles` | `[]string` | Optional | // allowed roles for accessing psk portal, if none, any role is permitted |
| `IdpCert` | `*string` | Optional | - |
| `IdpSignAlgo` | [`*models.PskPortalSsoIdpSignAlgoEnum`](../../doc/models/psk-portal-sso-idp-sign-algo-enum.md) | Optional | Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`. enum: `sha1`, `sha256`, `sha384`, `sha512`<br>**Default**: `"sha256"` |
| `IdpSsoUrl` | `*string` | Optional | - |
| `Issuer` | `*string` | Optional | - |
| `NameidFormat` | `*string` | Optional | - |
| `RoleMapping` | `map[string]string` | Optional | Property key is the role name, property value is the SSO Attribute |
| `UseSsoRoleForPskRole` | `*bool` | Optional | if enabled, the `role` above will be ignored |

## Example (as JSON)

```json
{
  "idp_sign_algo": "sha256",
  "allowed_roles": [
    "allowed_roles7"
  ],
  "idp_cert": "idp_cert2",
  "idp_sso_url": "idp_sso_url8",
  "issuer": "issuer8"
}
```

