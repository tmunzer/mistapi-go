package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Sso represents a Sso struct.
// SSO
type Sso struct {
    // when the object has been created, in epoch
    CreatedTime             *float64                     `json:"created_time,omitempty"`
    // if `idp_type`==`saml`, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process)
    CustomLogoutUrl         *string                      `json:"custom_logout_url,omitempty"`
    // if `idp_type`==`saml`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched
    DefaultRole             *string                      `json:"default_role,omitempty"`
    // random string generated during the SSO creation and used to generate the SAML URLs:
    // * ACS URL = `/api/v1/saml/{domain}/login` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/login`)
    // * Single Logout URL = `/api/v1/saml/{domain}/logout` (e.g. `https://api.mist.com/api/v1/saml/s4t5vwv8/logout`)
    Domain                  *string                      `json:"domain,omitempty"`
    // Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group
    GroupFilter             *string                      `json:"group_filter,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                      *uuid.UUID                   `json:"id,omitempty"`
    // if `idp_type`==`saml`. IDP Cert (used to verify the signed response)
    IdpCert                 *string                      `json:"idp_cert,omitempty"`
    // Required if `idp_type`==`saml`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`
    IdpSignAlgo             *SsoIdpSignAlgoEnum          `json:"idp_sign_algo,omitempty"`
    // Required if `idp_type`==`saml`, IDP Single-Sign-On URL
    IdpSsoUrl               *string                      `json:"idp_sso_url,omitempty"`
    // * For Admin SSO, enum: `saml`
    // * For NAC SSO, enum: `ldap`, `mxedge_proxy`, `oauth`
    IdpType                 *SsoIdpTypeEnum              `json:"idp_type,omitempty"`
    // if `idp_type`==`saml`, ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role
    IgnoreUnmatchedRoles    *bool                        `json:"ignore_unmatched_roles,omitempty"`
    // if `idp_type`==`saml`. IDP issuer URL
    Issuer                  *string                      `json:"issuer,omitempty"`
    // Required if `idp_type`==`ldap`, whole domain or a specific organization unit (container) in Search base to specify where users and groups are found in the LDAP tree
    LdapBaseDn              *string                      `json:"ldap_base_dn,omitempty"`
    // Required if `idp_type`==`ldap`, the account used to authenticate against the LDAP
    LdapBindDn              *string                      `json:"ldap_bind_dn,omitempty"`
    // Required if `idp_type`==`ldap`, the password used to authenticate against the LDAP
    LdapBindPassword        *string                      `json:"ldap_bind_password,omitempty"`
    // Required if `idp_type`==`ldap`, list of CA certificates to validate the LDAP certificate
    LdapCacerts             []string                     `json:"ldap_cacerts,omitempty"`
    // if `idp_type`==`ldap`, LDAPS Client certificate
    LdapClientCert          *string                      `json:"ldap_client_cert,omitempty"`
    // if `idp_type`==`ldap`, Key for the `ldap_client_cert`
    LdapClientKey           *string                      `json:"ldap_client_key,omitempty"`
    // if `ldap_type`==`custom`
    LdapGroupAttr           *string                      `json:"ldap_group_attr,omitempty"`
    // if `ldap_type`==`custom`
    LdapGroupDn             *string                      `json:"ldap_group_dn,omitempty"`
    // if `idp_type`==`ldap`, whether to recursively resolve LDAP groups
    LdapResolveGroups       *bool                        `json:"ldap_resolve_groups,omitempty"`
    // if `idp_type`==`ldap`, list of LDAP/LDAPS server IP Addresses or Hostnames
    LdapServerHosts         []string                     `json:"ldap_server_hosts,omitempty"`
    // if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`
    LdapType                *SsoLdapTypeEnum             `json:"ldap_type,omitempty"`
    // Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user
    LdapUserFilter          *string                      `json:"ldap_user_filter,omitempty"`
    // Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member
    MemberFilter            *string                      `json:"member_filter,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime            *float64                     `json:"modified_time,omitempty"`
    MspId                   *uuid.UUID                   `json:"msp_id,omitempty"`
    // if `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster
    MxedgeProxy             *SsoMxedgeProxy              `json:"mxedge_proxy,omitempty"`
    // name
    Name                    string                       `json:"name"`
    // if `idp_type`==`saml`. enum: `email`, `unspecified`
    NameidFormat            *SsoNameidFormatEnum         `json:"nameid_format,omitempty"`
    // Required if `idp_type`==`oauth`, Client Credentials
    OauthCcClientId         *string                      `json:"oauth_cc_client_id,omitempty"`
    // Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form "-----BEGIN RSA PRIVATE KEY--...."
    OauthCcClientSecret     *string                      `json:"oauth_cc_client_secret,omitempty"`
    // if `idp_type`==`oauth`
    OauthDiscoveryUrl       *string                      `json:"oauth_discovery_url,omitempty"`
    // enum: `us` (United States, default), `ca` (Canada), `eu` (Europe), `asia` (Asia), `au` (Australia)
    OauthPingIdentityRegion *OauthPingIdentityRegionEnum `json:"oauth_ping_identity_region,omitempty"`
    // if `idp_type`==`oauth`, ropc = Resource Owner Password Credentials
    OauthRopcClientId       *string                      `json:"oauth_ropc_client_id,omitempty"`
    // if `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty
    OauthRopcClientSecret   *string                      `json:"oauth_ropc_client_secret,omitempty"`
    // Required if `idp_type`==`oauth`, oauth_tenant_id
    OauthTenantId           *string                      `json:"oauth_tenant_id,omitempty"`
    // if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`
    OauthType               *SsoOauthTypeEnum            `json:"oauth_type,omitempty"`
    OrgId                   *uuid.UUID                   `json:"org_id,omitempty"`
    // if `idp_type`==`saml`, custom role attribute parsing scheme
    // Supported Role Parsing Schemes
    // <table><tr><th>Name</th><th>Scheme</th></tr><tr><td>cn</td><td><ul><li>The expected role attribute format in SAML Assertion is “CN=cn,OU=ou1,OU=ou2,…”</li><li>CN (the key) is case insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li><li>E.g. if role attribute is “CN=cn,OU=ou1,OU=ou2” then parsed role value is “cn”</li></ul></td></tr></table>
    RoleAttrExtraction      *string                      `json:"role_attr_extraction,omitempty"`
    // if `idp_type`==`saml`, name of the attribute in SAML Assertion to extract role from
    RoleAttrFrom            *string                      `json:"role_attr_from,omitempty"`
    // if `idp_type`==`oauth`, indicates if SCIM provisioning is enabled for the OAuth IDP
    ScimEnabled             *bool                        `json:"scim_enabled,omitempty"`
    // if `idp_type`==`oauth`, scim_secret_token (auto-generated when not provided by caller and `scim_enabled`==`true`, empty string when `scim_enabled`==`false`) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP
    ScimSecretToken         *string                      `json:"scim_secret_token,omitempty"`
    SiteId                  *uuid.UUID                   `json:"site_id,omitempty"`
    AdditionalProperties    map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for Sso,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Sso) String() string {
    return fmt.Sprintf(
    	"Sso[CreatedTime=%v, CustomLogoutUrl=%v, DefaultRole=%v, Domain=%v, GroupFilter=%v, Id=%v, IdpCert=%v, IdpSignAlgo=%v, IdpSsoUrl=%v, IdpType=%v, IgnoreUnmatchedRoles=%v, Issuer=%v, LdapBaseDn=%v, LdapBindDn=%v, LdapBindPassword=%v, LdapCacerts=%v, LdapClientCert=%v, LdapClientKey=%v, LdapGroupAttr=%v, LdapGroupDn=%v, LdapResolveGroups=%v, LdapServerHosts=%v, LdapType=%v, LdapUserFilter=%v, MemberFilter=%v, ModifiedTime=%v, MspId=%v, MxedgeProxy=%v, Name=%v, NameidFormat=%v, OauthCcClientId=%v, OauthCcClientSecret=%v, OauthDiscoveryUrl=%v, OauthPingIdentityRegion=%v, OauthRopcClientId=%v, OauthRopcClientSecret=%v, OauthTenantId=%v, OauthType=%v, OrgId=%v, RoleAttrExtraction=%v, RoleAttrFrom=%v, ScimEnabled=%v, ScimSecretToken=%v, SiteId=%v, AdditionalProperties=%v]",
    	s.CreatedTime, s.CustomLogoutUrl, s.DefaultRole, s.Domain, s.GroupFilter, s.Id, s.IdpCert, s.IdpSignAlgo, s.IdpSsoUrl, s.IdpType, s.IgnoreUnmatchedRoles, s.Issuer, s.LdapBaseDn, s.LdapBindDn, s.LdapBindPassword, s.LdapCacerts, s.LdapClientCert, s.LdapClientKey, s.LdapGroupAttr, s.LdapGroupDn, s.LdapResolveGroups, s.LdapServerHosts, s.LdapType, s.LdapUserFilter, s.MemberFilter, s.ModifiedTime, s.MspId, s.MxedgeProxy, s.Name, s.NameidFormat, s.OauthCcClientId, s.OauthCcClientSecret, s.OauthDiscoveryUrl, s.OauthPingIdentityRegion, s.OauthRopcClientId, s.OauthRopcClientSecret, s.OauthTenantId, s.OauthType, s.OrgId, s.RoleAttrExtraction, s.RoleAttrFrom, s.ScimEnabled, s.ScimSecretToken, s.SiteId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Sso.
// It customizes the JSON marshaling process for Sso objects.
func (s Sso) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "created_time", "custom_logout_url", "default_role", "domain", "group_filter", "id", "idp_cert", "idp_sign_algo", "idp_sso_url", "idp_type", "ignore_unmatched_roles", "issuer", "ldap_base_dn", "ldap_bind_dn", "ldap_bind_password", "ldap_cacerts", "ldap_client_cert", "ldap_client_key", "ldap_group_attr", "ldap_group_dn", "ldap_resolve_groups", "ldap_server_hosts", "ldap_type", "ldap_user_filter", "member_filter", "modified_time", "msp_id", "mxedge_proxy", "name", "nameid_format", "oauth_cc_client_id", "oauth_cc_client_secret", "oauth_discovery_url", "oauth_ping_identity_region", "oauth_ropc_client_id", "oauth_ropc_client_secret", "oauth_tenant_id", "oauth_type", "org_id", "role_attr_extraction", "role_attr_from", "scim_enabled", "scim_secret_token", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Sso object to a map representation for JSON marshaling.
func (s Sso) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.CustomLogoutUrl != nil {
        structMap["custom_logout_url"] = s.CustomLogoutUrl
    }
    if s.DefaultRole != nil {
        structMap["default_role"] = s.DefaultRole
    }
    if s.Domain != nil {
        structMap["domain"] = s.Domain
    }
    if s.GroupFilter != nil {
        structMap["group_filter"] = s.GroupFilter
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.IdpCert != nil {
        structMap["idp_cert"] = s.IdpCert
    }
    if s.IdpSignAlgo != nil {
        structMap["idp_sign_algo"] = s.IdpSignAlgo
    }
    if s.IdpSsoUrl != nil {
        structMap["idp_sso_url"] = s.IdpSsoUrl
    }
    if s.IdpType != nil {
        structMap["idp_type"] = s.IdpType
    }
    if s.IgnoreUnmatchedRoles != nil {
        structMap["ignore_unmatched_roles"] = s.IgnoreUnmatchedRoles
    }
    if s.Issuer != nil {
        structMap["issuer"] = s.Issuer
    }
    if s.LdapBaseDn != nil {
        structMap["ldap_base_dn"] = s.LdapBaseDn
    }
    if s.LdapBindDn != nil {
        structMap["ldap_bind_dn"] = s.LdapBindDn
    }
    if s.LdapBindPassword != nil {
        structMap["ldap_bind_password"] = s.LdapBindPassword
    }
    if s.LdapCacerts != nil {
        structMap["ldap_cacerts"] = s.LdapCacerts
    }
    if s.LdapClientCert != nil {
        structMap["ldap_client_cert"] = s.LdapClientCert
    }
    if s.LdapClientKey != nil {
        structMap["ldap_client_key"] = s.LdapClientKey
    }
    if s.LdapGroupAttr != nil {
        structMap["ldap_group_attr"] = s.LdapGroupAttr
    }
    if s.LdapGroupDn != nil {
        structMap["ldap_group_dn"] = s.LdapGroupDn
    }
    if s.LdapResolveGroups != nil {
        structMap["ldap_resolve_groups"] = s.LdapResolveGroups
    }
    if s.LdapServerHosts != nil {
        structMap["ldap_server_hosts"] = s.LdapServerHosts
    }
    if s.LdapType != nil {
        structMap["ldap_type"] = s.LdapType
    }
    if s.LdapUserFilter != nil {
        structMap["ldap_user_filter"] = s.LdapUserFilter
    }
    if s.MemberFilter != nil {
        structMap["member_filter"] = s.MemberFilter
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.MspId != nil {
        structMap["msp_id"] = s.MspId
    }
    if s.MxedgeProxy != nil {
        structMap["mxedge_proxy"] = s.MxedgeProxy.toMap()
    }
    structMap["name"] = s.Name
    if s.NameidFormat != nil {
        structMap["nameid_format"] = s.NameidFormat
    }
    if s.OauthCcClientId != nil {
        structMap["oauth_cc_client_id"] = s.OauthCcClientId
    }
    if s.OauthCcClientSecret != nil {
        structMap["oauth_cc_client_secret"] = s.OauthCcClientSecret
    }
    if s.OauthDiscoveryUrl != nil {
        structMap["oauth_discovery_url"] = s.OauthDiscoveryUrl
    }
    if s.OauthPingIdentityRegion != nil {
        structMap["oauth_ping_identity_region"] = s.OauthPingIdentityRegion
    }
    if s.OauthRopcClientId != nil {
        structMap["oauth_ropc_client_id"] = s.OauthRopcClientId
    }
    if s.OauthRopcClientSecret != nil {
        structMap["oauth_ropc_client_secret"] = s.OauthRopcClientSecret
    }
    if s.OauthTenantId != nil {
        structMap["oauth_tenant_id"] = s.OauthTenantId
    }
    if s.OauthType != nil {
        structMap["oauth_type"] = s.OauthType
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.RoleAttrExtraction != nil {
        structMap["role_attr_extraction"] = s.RoleAttrExtraction
    }
    if s.RoleAttrFrom != nil {
        structMap["role_attr_from"] = s.RoleAttrFrom
    }
    if s.ScimEnabled != nil {
        structMap["scim_enabled"] = s.ScimEnabled
    }
    if s.ScimSecretToken != nil {
        structMap["scim_secret_token"] = s.ScimSecretToken
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Sso.
// It customizes the JSON unmarshaling process for Sso objects.
func (s *Sso) UnmarshalJSON(input []byte) error {
    var temp tempSso
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "custom_logout_url", "default_role", "domain", "group_filter", "id", "idp_cert", "idp_sign_algo", "idp_sso_url", "idp_type", "ignore_unmatched_roles", "issuer", "ldap_base_dn", "ldap_bind_dn", "ldap_bind_password", "ldap_cacerts", "ldap_client_cert", "ldap_client_key", "ldap_group_attr", "ldap_group_dn", "ldap_resolve_groups", "ldap_server_hosts", "ldap_type", "ldap_user_filter", "member_filter", "modified_time", "msp_id", "mxedge_proxy", "name", "nameid_format", "oauth_cc_client_id", "oauth_cc_client_secret", "oauth_discovery_url", "oauth_ping_identity_region", "oauth_ropc_client_id", "oauth_ropc_client_secret", "oauth_tenant_id", "oauth_type", "org_id", "role_attr_extraction", "role_attr_from", "scim_enabled", "scim_secret_token", "site_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CreatedTime = temp.CreatedTime
    s.CustomLogoutUrl = temp.CustomLogoutUrl
    s.DefaultRole = temp.DefaultRole
    s.Domain = temp.Domain
    s.GroupFilter = temp.GroupFilter
    s.Id = temp.Id
    s.IdpCert = temp.IdpCert
    s.IdpSignAlgo = temp.IdpSignAlgo
    s.IdpSsoUrl = temp.IdpSsoUrl
    s.IdpType = temp.IdpType
    s.IgnoreUnmatchedRoles = temp.IgnoreUnmatchedRoles
    s.Issuer = temp.Issuer
    s.LdapBaseDn = temp.LdapBaseDn
    s.LdapBindDn = temp.LdapBindDn
    s.LdapBindPassword = temp.LdapBindPassword
    s.LdapCacerts = temp.LdapCacerts
    s.LdapClientCert = temp.LdapClientCert
    s.LdapClientKey = temp.LdapClientKey
    s.LdapGroupAttr = temp.LdapGroupAttr
    s.LdapGroupDn = temp.LdapGroupDn
    s.LdapResolveGroups = temp.LdapResolveGroups
    s.LdapServerHosts = temp.LdapServerHosts
    s.LdapType = temp.LdapType
    s.LdapUserFilter = temp.LdapUserFilter
    s.MemberFilter = temp.MemberFilter
    s.ModifiedTime = temp.ModifiedTime
    s.MspId = temp.MspId
    s.MxedgeProxy = temp.MxedgeProxy
    s.Name = *temp.Name
    s.NameidFormat = temp.NameidFormat
    s.OauthCcClientId = temp.OauthCcClientId
    s.OauthCcClientSecret = temp.OauthCcClientSecret
    s.OauthDiscoveryUrl = temp.OauthDiscoveryUrl
    s.OauthPingIdentityRegion = temp.OauthPingIdentityRegion
    s.OauthRopcClientId = temp.OauthRopcClientId
    s.OauthRopcClientSecret = temp.OauthRopcClientSecret
    s.OauthTenantId = temp.OauthTenantId
    s.OauthType = temp.OauthType
    s.OrgId = temp.OrgId
    s.RoleAttrExtraction = temp.RoleAttrExtraction
    s.RoleAttrFrom = temp.RoleAttrFrom
    s.ScimEnabled = temp.ScimEnabled
    s.ScimSecretToken = temp.ScimSecretToken
    s.SiteId = temp.SiteId
    return nil
}

// tempSso is a temporary struct used for validating the fields of Sso.
type tempSso  struct {
    CreatedTime             *float64                     `json:"created_time,omitempty"`
    CustomLogoutUrl         *string                      `json:"custom_logout_url,omitempty"`
    DefaultRole             *string                      `json:"default_role,omitempty"`
    Domain                  *string                      `json:"domain,omitempty"`
    GroupFilter             *string                      `json:"group_filter,omitempty"`
    Id                      *uuid.UUID                   `json:"id,omitempty"`
    IdpCert                 *string                      `json:"idp_cert,omitempty"`
    IdpSignAlgo             *SsoIdpSignAlgoEnum          `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl               *string                      `json:"idp_sso_url,omitempty"`
    IdpType                 *SsoIdpTypeEnum              `json:"idp_type,omitempty"`
    IgnoreUnmatchedRoles    *bool                        `json:"ignore_unmatched_roles,omitempty"`
    Issuer                  *string                      `json:"issuer,omitempty"`
    LdapBaseDn              *string                      `json:"ldap_base_dn,omitempty"`
    LdapBindDn              *string                      `json:"ldap_bind_dn,omitempty"`
    LdapBindPassword        *string                      `json:"ldap_bind_password,omitempty"`
    LdapCacerts             []string                     `json:"ldap_cacerts,omitempty"`
    LdapClientCert          *string                      `json:"ldap_client_cert,omitempty"`
    LdapClientKey           *string                      `json:"ldap_client_key,omitempty"`
    LdapGroupAttr           *string                      `json:"ldap_group_attr,omitempty"`
    LdapGroupDn             *string                      `json:"ldap_group_dn,omitempty"`
    LdapResolveGroups       *bool                        `json:"ldap_resolve_groups,omitempty"`
    LdapServerHosts         []string                     `json:"ldap_server_hosts,omitempty"`
    LdapType                *SsoLdapTypeEnum             `json:"ldap_type,omitempty"`
    LdapUserFilter          *string                      `json:"ldap_user_filter,omitempty"`
    MemberFilter            *string                      `json:"member_filter,omitempty"`
    ModifiedTime            *float64                     `json:"modified_time,omitempty"`
    MspId                   *uuid.UUID                   `json:"msp_id,omitempty"`
    MxedgeProxy             *SsoMxedgeProxy              `json:"mxedge_proxy,omitempty"`
    Name                    *string                      `json:"name"`
    NameidFormat            *SsoNameidFormatEnum         `json:"nameid_format,omitempty"`
    OauthCcClientId         *string                      `json:"oauth_cc_client_id,omitempty"`
    OauthCcClientSecret     *string                      `json:"oauth_cc_client_secret,omitempty"`
    OauthDiscoveryUrl       *string                      `json:"oauth_discovery_url,omitempty"`
    OauthPingIdentityRegion *OauthPingIdentityRegionEnum `json:"oauth_ping_identity_region,omitempty"`
    OauthRopcClientId       *string                      `json:"oauth_ropc_client_id,omitempty"`
    OauthRopcClientSecret   *string                      `json:"oauth_ropc_client_secret,omitempty"`
    OauthTenantId           *string                      `json:"oauth_tenant_id,omitempty"`
    OauthType               *SsoOauthTypeEnum            `json:"oauth_type,omitempty"`
    OrgId                   *uuid.UUID                   `json:"org_id,omitempty"`
    RoleAttrExtraction      *string                      `json:"role_attr_extraction,omitempty"`
    RoleAttrFrom            *string                      `json:"role_attr_from,omitempty"`
    ScimEnabled             *bool                        `json:"scim_enabled,omitempty"`
    ScimSecretToken         *string                      `json:"scim_secret_token,omitempty"`
    SiteId                  *uuid.UUID                   `json:"site_id,omitempty"`
}

func (s *tempSso) validate() error {
    var errs []string
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sso`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
