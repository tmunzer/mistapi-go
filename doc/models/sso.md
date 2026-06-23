
# Sso

Single sign-on identity provider configuration

*This model accepts additional fields of type interface{}.*

## Structure

`Sso`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `CustomLogoutUrl` | `*string` | Optional | If `idp_type`==`saml`, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process) |
| `DefaultRole` | `*string` | Optional | If `idp_type`==`saml`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `Domain` | `*string` | Optional, Read-only | Random string generated during the SSO creation and used to generate the SAML URLs:<br><br>* ACS URL = `/api/v1/saml/{domain}/login` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/login`)<br>* Single Logout URL = `/api/v1/saml/{domain}/logout` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/logout`) |
| `GroupFilter` | `*string` | Optional | Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
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
| `LdapGroupAttr` | `*string` | Optional | Group attribute used to resolve LDAP memberships. If `ldap_type`==`custom`<br><br>**Default**: `"memberOf"` |
| `LdapGroupDn` | `*string` | Optional | Group search base used for custom LDAP group lookup. If `ldap_type`==`custom`<br><br>**Default**: `"base_dn"` |
| `LdapResolveGroups` | `*bool` | Optional | If `idp_type`==`ldap`, whether to recursively resolve LDAP groups<br><br>**Default**: `false` |
| `LdapServerHosts` | `[]string` | Optional | If `idp_type`==`ldap`, list of LDAP/LDAPS server IP addresses or Hostnames |
| `LdapType` | [`*models.SsoLdapTypeEnum`](../../doc/models/sso-ldap-type-enum.md) | Optional | if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`<br><br>**Default**: `"azure"` |
| `LdapUserFilter` | `*string` | Optional | Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user |
| `MemberFilter` | `*string` | Optional | Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `MxedgeProxy` | [`*models.SsoMxedgeProxy`](../../doc/models/sso-mxedge-proxy.md) | Optional | Mist Edge proxy settings for NAC SSO. If `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster |
| `Name` | `string` | Required | Display name of the SSO configuration |
| `NameidFormat` | [`*models.SsoNameidFormatEnum`](../../doc/models/sso-nameid-format-enum.md) | Optional | if `idp_type`==`saml`. enum: `email`, `unspecified`<br><br>**Default**: `"email"` |
| `OauthCcClientId` | `*string` | Optional | Required if `idp_type`==`oauth`, Client Credentials |
| `OauthCcClientSecret` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form "-----BEGIN RSA PRIVATE KEY--...." |
| `OauthDiscoveryUrl` | `*string` | Optional | OAuth discovery document URL used when `idp_type`==`oauth` |
| `OauthPingIdentityRegion` | [`*models.OauthPingIdentityRegionEnum`](../../doc/models/oauth-ping-identity-region-enum.md) | Optional | enum: `us` (United States, default), `ca` (Canada), `eu` (Europe), `asia` (Asia), `au` (Australia)<br><br>**Default**: `"us"` |
| `OauthProviderDomain` | [`*models.OauthProviderDomainEnum`](../../doc/models/oauth-provider-domain-enum.md) | Optional | If `oauth_type`==`okta`, specifies the region-specific OAuth provider domain. enum: `okta.com`, `oktapreview.com`, `okta-emea.com`, `okta-gov.com`, `okta.mil`, `mtls.okta.com`<br><br>**Default**: `"okta.com"` |
| `OauthRopcClientId` | `*string` | Optional | If `idp_type`==`oauth`, ropc = Resource Owner Password Credentials |
| `OauthRopcClientSecret` | `*string` | Optional | If `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty |
| `OauthTenantId` | `*string` | Optional | Required if `idp_type`==`oauth`, oauth_tenant_id |
| `OauthType` | [`*models.SsoOauthTypeEnum`](../../doc/models/sso-oauth-type-enum.md) | Optional | if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`<br><br>**Default**: `"azure"` |
| `OpenroamingSsids` | `[]string` | Optional | Network SSID names enabled for OpenRoaming SSO |
| `OpenroamingWbaClientCert` | `*string` | Optional | Optional WBA-issued client certificate for OpenRoaming. If not provided, the default WBA-issued certificate for Juniper will be used. |
| `OpenroamingWbaClientKey` | `*string` | Optional | Optional WBA-issued client private key for OpenRoaming. If not provided, the default WBA-issued key for Juniper will be used. |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RoleAttrExtraction` | `*string` | Optional | If `idp_type`==`saml`, custom role attribute parsing scheme. Supported Role Parsing Schemes <table><tr><th>Name</th><th>Scheme</th></tr><tr><td>`cn`</td><td><ul><li>The expected role attribute format in SAML Assertion is "CN=cn,OU=ou1,OU=ou2,…"</li><li>CN (the key) is case-insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li></ul>E.g. if role attribute is "CN=cn,OU=ou1,OU=ou2" then parsed role value is "cn"</td></tr></table> |
| `RoleAttrFrom` | `*string` | Optional | If `idp_type`==`saml`, name of the attribute in SAML Assertion to extract role from<br><br>**Default**: `"Role"` |
| `ScimEnabled` | `*bool` | Optional | If `idp_type`==`oauth`, indicates if SCIM provisioning is enabled for the OAuth IDP<br><br>**Default**: `false` |
| `ScimSecretToken` | `*string` | Optional | If `idp_type`==`oauth`, scim_secret_token (auto-generated when not provided by caller and `scim_enabled`==`true`, empty string when `scim_enabled`==`false`) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sso := models.Sso{
        CreatedTime:              models.ToPointer(float64(58.56)),
        CustomLogoutUrl:          models.ToPointer("custom_logout_url0"),
        DefaultRole:              models.ToPointer("default_role8"),
        Domain:                   models.ToPointer("domain2"),
        GroupFilter:              models.ToPointer("group_filter4"),
        Id:                       models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IdpType:                  models.ToPointer(models.SsoIdpTypeEnum_SAML),
        LdapBaseDn:               models.ToPointer("DC=abc,DC=com"),
        LdapBindDn:               models.ToPointer("CN=nas,CN=users,DC=abc,DC=com"),
        LdapBindPassword:         models.ToPointer("secret"),
        LdapCacerts:              []string{
            "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
            "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
        },
        LdapClientCert:           models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
        LdapClientKey:            models.ToPointer("-----BEGIN PRI..."),
        LdapGroupAttr:            models.ToPointer("memberOf"),
        LdapGroupDn:              models.ToPointer("base_dn"),
        LdapResolveGroups:        models.ToPointer(false),
        LdapServerHosts:          []string{
            "hostname",
            "63.1.3.5",
        },
        LdapType:                 models.ToPointer(models.SsoLdapTypeEnum_AZURE),
        LdapUserFilter:           models.ToPointer("(mail=%s)"),
        MemberFilter:             models.ToPointer("(CN=%s)"),
        MspId:                    models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        Name:                     "name6",
        NameidFormat:             models.ToPointer(models.SsoNameidFormatEnum_EMAIL),
        OauthCcClientId:          models.ToPointer("e60da615-7def-4c5a-8196-43675f45e174"),
        OauthCcClientSecret:      models.ToPointer("akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH"),
        OauthPingIdentityRegion:  models.ToPointer(models.OauthPingIdentityRegionEnum_US),
        OauthProviderDomain:      models.ToPointer(models.OauthProviderDomainEnum_ENUMOKTACOM),
        OauthRopcClientId:        models.ToPointer("9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7"),
        OauthRopcClientSecret:    models.ToPointer("blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI"),
        OauthTenantId:            models.ToPointer("dev-88336535"),
        OauthType:                models.ToPointer(models.SsoOauthTypeEnum_AZURE),
        OpenroamingSsids:         []string{
            "ssid_name1",
            "ssid_name2",
        },
        OpenroamingWbaClientCert: models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
        OpenroamingWbaClientKey:  models.ToPointer("-----BEGIN PRIVATE KEY-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END PRIVATE KEY-----"),
        OrgId:                    models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RoleAttrFrom:             models.ToPointer("Role"),
        ScimEnabled:              models.ToPointer(false),
        ScimSecretToken:          models.ToPointer("FBitbKPE1aecSloPGBuqqPxDUrFeZyZk"),
        SiteId:                   models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties:     map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

