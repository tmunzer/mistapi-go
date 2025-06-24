
# Sso

SSO

*This model accepts additional fields of type interface{}.*

## Structure

`Sso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `CustomLogoutUrl` | `*string` | Optional | If `idp_type`==`saml`, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process) |
| `DefaultRole` | `*string` | Optional | If `idp_type`==`saml`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `Domain` | `*string` | Optional | Random string generated during the SSO creation and used to generate the SAML URLs:<br><br>* ACS URL = `/api/v1/saml/{domain}/login` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/login`)<br>* Single Logout URL = `/api/v1/saml/{domain}/logout` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/logout`) |
| `GroupFilter` | `*string` | Optional | Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IdpCert` | `*string` | Optional | If `idp_type`==`saml`. IDP Cert (used to verify the signed response) |
| `IdpSignAlgo` | [`*models.SsoIdpSignAlgoEnum`](../../doc/models/sso-idp-sign-algo-enum.md) | Optional | Required if `idp_type`==`saml`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512` |
| `IdpSsoUrl` | `*string` | Optional | Required if `idp_type`==`saml`, IDP Single-Sign-On URL |
| `IdpType` | [`*models.SsoIdpTypeEnum`](../../doc/models/sso-idp-type-enum.md) | Optional | SSO IDP Type:<br><br>* For Admin SSO, enum: `saml`<br>* For NAC SSO, enum: `ldap`, `mxedge_proxy`, `oauth`, `openroaming`<br><br>**Default**: `"saml"` |
| `IgnoreUnmatchedRoles` | `*bool` | Optional | If `idp_type`==`saml`, ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role |
| `Issuer` | `*string` | Optional | If `idp_type`==`saml`. IDP issuer URL |
| `LdapBaseDn` | `*string` | Optional | Required if `idp_type`==`ldap`, whole domain or a specific organization unit (container) in Search base to specify where users and groups are found in the LDAP tree |
| `LdapBindDn` | `*string` | Optional | Required if `idp_type`==`ldap`, the account used to authenticate against the LDAP |
| `LdapBindPassword` | `*string` | Optional | Required if `idp_type`==`ldap`, the password used to authenticate against the LDAP |
| `LdapCacerts` | `[]string` | Optional | Required if `idp_type`==`ldap`, list of CA certificates to validate the LDAP certificate |
| `LdapClientCert` | `*string` | Optional | If `idp_type`==`ldap`, LDAPS Client certificate |
| `LdapClientKey` | `*string` | Optional | If `idp_type`==`ldap`, Key for the `ldap_client_cert` |
| `LdapGroupAttr` | `*string` | Optional | If `ldap_type`==`custom`<br><br>**Default**: `"memberOf"` |
| `LdapGroupDn` | `*string` | Optional | If `ldap_type`==`custom`<br><br>**Default**: `"base_dn"` |
| `LdapResolveGroups` | `*bool` | Optional | If `idp_type`==`ldap`, whether to recursively resolve LDAP groups<br><br>**Default**: `false` |
| `LdapServerHosts` | `[]string` | Optional | If `idp_type`==`ldap`, list of LDAP/LDAPS server IP Addresses or Hostnames |
| `LdapType` | [`*models.SsoLdapTypeEnum`](../../doc/models/sso-ldap-type-enum.md) | Optional | if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`<br><br>**Default**: `"azure"` |
| `LdapUserFilter` | `*string` | Optional | Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user |
| `MemberFilter` | `*string` | Optional | Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MxedgeProxy` | [`*models.SsoMxedgeProxy`](../../doc/models/sso-mxedge-proxy.md) | Optional | If `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster |
| `Name` | `string` | Required | Name |
| `NameidFormat` | [`*models.SsoNameidFormatEnum`](../../doc/models/sso-nameid-format-enum.md) | Optional | if `idp_type`==`saml`. enum: `email`, `unspecified`<br><br>**Default**: `"email"` |
| `OauthCcClientId` | `*string` | Optional | Required if `idp_type`==`oauth`, Client Credentials |
| `OauthCcClientSecret` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form "-----BEGIN RSA PRIVATE KEY--...." |
| `OauthDiscoveryUrl` | `*string` | Optional | If `idp_type`==`oauth` |
| `OauthPingIdentityRegion` | [`*models.OauthPingIdentityRegionEnum`](../../doc/models/oauth-ping-identity-region-enum.md) | Optional | enum: `us` (United States, default), `ca` (Canada), `eu` (Europe), `asia` (Asia), `au` (Australia)<br><br>**Default**: `"us"` |
| `OauthRopcClientId` | `*string` | Optional | If `idp_type`==`oauth`, ropc = Resource Owner Password Credentials |
| `OauthRopcClientSecret` | `*string` | Optional | If `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty |
| `OauthTenantId` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_tenant_id |
| `OauthType` | [`*models.SsoOauthTypeEnum`](../../doc/models/sso-oauth-type-enum.md) | Optional | if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`<br><br>**Default**: `"azure"` |
| `Openroaming` | [`*models.SsoOpenroaming`](../../doc/models/sso-openroaming.md) | Optional | if `idp_type`==`openroaming` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RoleAttrExtraction` | `*string` | Optional | If `idp_type`==`saml`, custom role attribute parsing scheme. Supported Role Parsing Schemes <table><tr><th>Name</th><th>Scheme</th></tr><tr><td>`cn`</td><td><ul><li>The expected role attribute format in SAML Assertion is "CN=cn,OU=ou1,OU=ou2,…"</li><li>CN (the key) is case-insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li></ul>E.g. if role attribute is "CN=cn,OU=ou1,OU=ou2" then parsed role value is "cn"</td></tr></table> |
| `RoleAttrFrom` | `*string` | Optional | If `idp_type`==`saml`, name of the attribute in SAML Assertion to extract role from<br><br>**Default**: `"Role"` |
| `ScimEnabled` | `*bool` | Optional | If `idp_type`==`oauth`, indicates if SCIM provisioning is enabled for the OAuth IDP<br><br>**Default**: `false` |
| `ScimSecretToken` | `*string` | Optional | If `idp_type`==`oauth`, scim_secret_token (auto-generated when not provided by caller and `scim_enabled`==`true`, empty string when `scim_enabled`==`false`) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "idp_type": "saml",
  "ldap_base_dn": "DC=abc,DC=com",
  "ldap_bind_dn": "CN=nas,CN=users,DC=abc,DC=com",
  "ldap_bind_password": "secret",
  "ldap_cacerts": [
    "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
    "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----"
  ],
  "ldap_client_cert": "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
  "ldap_client_key": "-----BEGIN PRI...",
  "ldap_group_attr": "memberOf",
  "ldap_group_dn": "base_dn",
  "ldap_resolve_groups": false,
  "ldap_server_hosts": [
    "hostname",
    "63.1.3.5"
  ],
  "ldap_type": "azure",
  "ldap_user_filter": "(mail=%s)",
  "member_filter": "(CN=%s)",
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "name": "name6",
  "nameid_format": "email",
  "oauth_cc_client_id": "e60da615-7def-4c5a-8196-43675f45e174",
  "oauth_cc_client_secret": "akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH",
  "oauth_ping_identity_region": "us",
  "oauth_ropc_client_id": "9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7",
  "oauth_ropc_client_secret": "blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI",
  "oauth_tenant_id": "dev-88336535",
  "oauth_type": "azure",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "role_attr_from": "Role",
  "scim_enabled": false,
  "scim_secret_token": "FBitbKPE1aecSloPGBuqqPxDUrFeZyZk",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 58.56,
  "custom_logout_url": "custom_logout_url0",
  "default_role": "default_role8",
  "domain": "domain2",
  "group_filter": "group_filter4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

