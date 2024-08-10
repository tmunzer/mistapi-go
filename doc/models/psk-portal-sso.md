
# Psk Portal Sso

if `auth`==`sso`

## Structure

`PskPortalSso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedRoles` | `[]string` | Optional | // allowed roles for accessing psk portal, if none, any role is permitted |
| `IdpCert` | `*string` | Optional | - |
| `IdpSignAlgo` | `*string` | Optional | - |
| `IdpSsoUrl` | `*string` | Optional | - |
| `Issuer` | `*string` | Optional | - |
| `NameidFormat` | `*string` | Optional | - |
| `RoleMapping` | `map[string]string` | Optional | Property key is the role name, property value is the SSO Attribute |
| `UseSsoRoleForPskRole` | `*bool` | Optional | if enabled, the `role` above will be ignored |

## Example (as JSON)

```json
{
  "allowed_roles": [
    "allowed_roles7"
  ],
  "idp_cert": "idp_cert2",
  "idp_sign_algo": "idp_sign_algo4",
  "idp_sso_url": "idp_sso_url8",
  "issuer": "issuer8"
}
```

