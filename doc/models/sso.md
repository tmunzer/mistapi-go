
# Sso

SSO

## Structure

`Sso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `CustomLogoutUrl` | `*string` | Optional | if `idp_type`==`saml`, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process) |
| `DefaultRole` | `*string` | Optional | if `idp_type`==`saml`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `Domain` | `*string` | Optional | if `idp_type`==`ldap` or `idp_type`==`oauth`, list of domain names to match to use this IDP profile |
| `Id` | `*uuid.UUID` | Optional | - |
| `IdpCert` | `*string` | Optional | if `idp_type`==`saml`. IDP Cert (used to verify the signed response) |
| `IdpSignAlgo` | [`*models.SsoIdpSignAlgoEnum`](../../doc/models/sso-idp-sign-algo-enum.md) | Optional | Required if `idp_type`==`saml`, Signing algorithm for SAML Assertion. enum `sha1`, `sha256`, `sha384`, `sha512` |
| `IdpSsoUrl` | `*string` | Optional | Required if `idp_type`==`saml`, IDP Single-Sign-On URL |
| `IdpType` | [`models.SsoIdpTypeEnum`](../../doc/models/sso-idp-type-enum.md) | Required | * For Admin SSO, enum: `saml`<br>* For NAC SSO, enum: `ldap`, `mxedge_proxy`, `oauth`<br>**Default**: `"saml"` |
| `IgnoreUnmatchedRoles` | `*bool` | Optional | if `idp_type`==`saml`, ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role |
| `Issuer` | `*string` | Optional | if `idp_type`==`saml`. IDP issuer URL |
| `LdapBaseDn` | `*string` | Optional | Required if `idp_type`==`ldap` |
| `LdapBindDn` | `*string` | Optional | Required if `idp_type`==`ldap` |
| `LdapBindPassword` | `*string` | Optional | Required if `idp_type`==`ldap` |
| `LdapCerts` | `[]string` | Optional | if `idp_type`==`ldap` |
| `LdapClientCert` | `*string` | Optional | if `idp_type`==`ldap` |
| `LdapClientKey` | `*string` | Optional | if `idp_type`==`ldap` |
| `LdapGroupAttr` | `*string` | Optional | if `ldap_type`==`custom` |
| `LdapGroupDn` | `*string` | Optional | if `ldap_type`==`custom` |
| `LdapGroupFilter` | `*string` | Optional | Required if `ldap_type`==`custom` |
| `LdapMemberFilter` | `*string` | Optional | Required if `ldap_type`==`custom` |
| `LdapResolveGroups` | `*bool` | Optional | if `idp_type`==`ldap`, whether to recursively resolve LDAP groups<br>**Default**: `false` |
| `LdapServerHosts` | `[]string` | Optional | Required if `idp_type`==`ldap` |
| `LdapType` | [`*models.SsoLdapTypeEnum`](../../doc/models/sso-ldap-type-enum.md) | Optional | if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`<br>**Default**: `"azure"` |
| `LdapUserFilter` | `*string` | Optional | Required if `ldap_type`==`custom` |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MxedgeProxy` | [`*models.SsoMxedgeProxy`](../../doc/models/sso-mxedge-proxy.md) | Optional | if `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster |
| `Name` | `string` | Required | name |
| `NameidFormat` | [`*models.SsoNameidFormatEnum`](../../doc/models/sso-nameid-format-enum.md) | Optional | if `idp_type`==`saml`. enum: `email`, `unspecified`<br>**Default**: `"email"` |
| `OauthCcClientId` | `*string` | Optional | Required if `idp_type`==`oauth`, Client Credentials |
| `OauthCcClientSecret` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form "-----BEGIN RSA PRIVATE KEY--...." |
| `OauthDiscoveryUrl` | `*string` | Optional | if `idp_type`==`oauth` |
| `OauthRopcClientId` | `*string` | Optional | if `idp_type`==`oauth`, ropc = Resource Owner Password Credentials |
| `OauthRopcClientSecret` | `*string` | Optional | if `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty |
| `OauthTenantId` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_tenant_id |
| `OauthType` | [`*models.SsoOauthTypeEnum`](../../doc/models/sso-oauth-type-enum.md) | Optional | if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`<br>**Default**: `"azure"` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RoleAttrExtraction` | `*string` | Optional | if `idp_type`==`saml`, custom role attribute parsing scheme<br><br>Supported Role Parsing Schemes<br><br><table><tr><th>Name</th><th>Scheme</th></tr><tr><td>cn</td><td><ul><li>The expected role attribute format in SAML Assertion is “CN=cn,OU=ou1,OU=ou2,…”</li><li>CN (the key) is case insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li><li>E.g. if role attribute is “CN=cn,OU=ou1,OU=ou2” then parsed role value is “cn”</li></ul></td></tr></table><br> |
| `RoleAttrFrom` | `*string` | Optional | if `idp_type`==`saml`, name of the attribute in SAML Assertion to extract role from<br>**Default**: `"Role"` |
| `ScimEnabled` | `*bool` | Optional | if `idp_type`==`oauth`, indicates if SCIM provisioning is enabled for the OAuth IDP<br>**Default**: `false` |
| `ScimSecretToken` | `*string` | Optional | if `idp_type`==`oauth`, scim_secret_token (generated by caller, crypto-random) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "idp_type": "saml",
  "ldap_base_dn": "DC=mycorp,DC=org",
  "ldap_bind_dn": "admin@mycorp.org",
  "ldap_bind_password": "secret",
  "ldap_member_filter": "(CN=%s)",
  "ldap_resolve_groups": false,
  "ldap_type": "azure",
  "name": "name6",
  "nameid_format": "email",
  "oauth_cc_client_id": "e60da615-7def-4c5a-8196-43675f45e174",
  "oauth_cc_client_secret": "akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH",
  "oauth_ropc_client_id": "9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7",
  "oauth_ropc_client_secret": "blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI",
  "oauth_tenant_id": "dev-88336535",
  "oauth_type": "azure",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "role_attr_from": "Role",
  "scim_enabled": false,
  "scim_secret_token": "secret token",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 58.56,
  "custom_logout_url": "custom_logout_url0",
  "default_role": "default_role8",
  "domain": "domain2",
  "id": "00001dca-0000-0000-0000-000000000000"
}
```

