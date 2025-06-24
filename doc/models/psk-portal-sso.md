
# Psk Portal Sso

If `auth`==`sso`

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedRoles` | `[]string` | Optional | // allowed roles for accessing psk portal, if none, any role is permitted |
| `IdpCert` | `*string` | Optional | - |
| `IdpSignAlgo` | [`*models.PskPortalSsoIdpSignAlgoEnum`](../../doc/models/psk-portal-sso-idp-sign-algo-enum.md) | Optional | Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`. enum: `sha1`, `sha256`, `sha384`, `sha512`<br><br>**Default**: `"sha256"` |
| `IdpSsoUrl` | `*string` | Optional | - |
| `Issuer` | `*string` | Optional | - |
| `NameidFormat` | `*string` | Optional | - |
| `RoleMapping` | `map[string]string` | Optional | Property key is the role name, property value is the SSO Attribute |
| `UseSsoRoleForPskRole` | `*bool` | Optional | If enabled, the `role` above will be ignored |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "idp_sign_algo": "sha256",
  "allowed_roles": [
    "allowed_roles7"
  ],
  "idp_cert": "idp_cert2",
  "idp_sso_url": "idp_sso_url8",
  "issuer": "issuer8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

