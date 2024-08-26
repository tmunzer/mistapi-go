package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Sso represents a Sso struct.
// SSO
type Sso struct {
    CreatedTime           *float64             `json:"created_time,omitempty"`
    // optional, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process)
    CustomLogoutUrl       *string              `json:"custom_logout_url,omitempty"`
    // default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched
    DefaultRole           *string              `json:"default_role,omitempty"`
    Domain                *string              `json:"domain,omitempty"`
    Id                    *uuid.UUID           `json:"id,omitempty"`
    // if `idp_type`==`saml`. IDP Cert (used to verify the signed response)
    IdpCert               *string              `json:"idp_cert,omitempty"`
    // if `idp_type`==`saml`. Signing algorithm for SAML Assertion
    IdpSignAlgo           *string              `json:"idp_sign_algo,omitempty"`
    // IDP Single-Sign-On URL
    IdpSsoUrl             *string              `json:"idp_sso_url,omitempty"`
    // enum: `ldap`, `mxedge_proxy`, `oauth`, `saml`
    IdpType               *SsoIdpTypeEnum      `json:"idp_type,omitempty"`
    // ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role
    IgnoreUnmatchedRoles  *bool                `json:"ignore_unmatched_roles,omitempty"`
    // if `idp_type`==`saml`. IDP issuer URL
    Issuer                *string              `json:"issuer,omitempty"`
    // if `idp_type`==`ldap`
    LdapBaseDn            *string              `json:"ldap_base_dn,omitempty"`
    // if `idp_type`==`ldap`
    LdapBindDn            *string              `json:"ldap_bind_dn,omitempty"`
    // if `idp_type`==`ldap`
    LdapBindPassword      *string              `json:"ldap_bind_password,omitempty"`
    // if `idp_type`==`ldap`
    LdapCerts             []string             `json:"ldap_certs,omitempty"`
    // if `idp_type`==`ldap`
    LdapClientCert        *string              `json:"ldap_client_cert,omitempty"`
    // if `idp_type`==`ldap`
    LdapClientKey         *string              `json:"ldap_client_key,omitempty"`
    // Only if `ldap_type`==`custom`
    LdapGroupAttr         *string              `json:"ldap_group_attr,omitempty"`
    // Only if `ldap_type`==`custom`
    LdapGroupDn           *string              `json:"ldap_group_dn,omitempty"`
    // Only if `ldap_type`==`custom`
    LdapGroupFilter       *string              `json:"ldap_group_filter,omitempty"`
    // whether to recursively resolve LDAP groups
    LdapResolveGroups     *bool                `json:"ldap_resolve_groups,omitempty"`
    // if `idp_type`==`ldap`
    LdapServerHosts       []string             `json:"ldap_server_hosts,omitempty"`
    // if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`
    LdapType              *SsoLdapTypeEnum     `json:"ldap_type,omitempty"`
    // Only if `ldap_type`==`custom`
    LdapUserFilter        *string              `json:"ldap_user_filter,omitempty"`
    ModifiedTime          *float64             `json:"modified_time,omitempty"`
    MspId                 *uuid.UUID           `json:"msp_id,omitempty"`
    // if `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster
    MxedgeProxy           *SsoMxedgeProxy      `json:"mxedge_proxy,omitempty"`
    // name
    Name                  string               `json:"name"`
    // if `idp_type`==`saml`. enum: `email`, `unspecified`
    NameidFormat          *SsoNameidFormatEnum `json:"nameid_format,omitempty"`
    // if `oauth_type`==`okta`, Client Credentials
    OauthCcClientId       *string              `json:"oauth_cc_client_id,omitempty"`
    // if `oauth_type`==`okta` or `oauth_type`==`ping_identity`, oauth_cc_client_secret is RSA private key, of the form "-----BEGIN RSA PRIVATE KEY--...."
    OauthCcClientSecret   *string              `json:"oauth_cc_client_secret,omitempty"`
    // if `idp_type`==`oauth`
    OauthDiscoveryUrl     *string              `json:"oauth_discovery_url,omitempty"`
    // ropc = Resource Owner Password Credentials
    OauthRopcClientId     *string              `json:"oauth_ropc_client_id,omitempty"`
    // oauth_ropc_client_secret can be empty if oauth_type is azure or azure-gov
    OauthRopcClientSecret *string              `json:"oauth_ropc_client_secret,omitempty"`
    // if `oauth_type`==`okta` or `oauth_type`==`ping_identity`, oauth_tenant_id
    OauthTenantId         *string              `json:"oauth_tenant_id,omitempty"`
    // enum: `azure`, `azure-gov`, `okta`, `ping_identity`
    OauthType             *SsoOauthTypeEnum    `json:"oauth_type,omitempty"`
    OrgId                 *uuid.UUID           `json:"org_id,omitempty"`
    // optional, custom role attribute parsing scheme
    // Supported Role Parsing Schemes
    // <table><tr><th>Name</th><th>Scheme</th></tr><tr><td>cn</td><td><ul><li>The expected role attribute format in SAML Assertion is “CN=cn,OU=ou1,OU=ou2,…”</li><li>CN (the key) is case insensitive and exactly 1 CN is expected (or the entire entry will be ignored)</li><li>E.g. if role attribute is “CN=cn,OU=ou1,OU=ou2” then parsed role value is “cn”</li></ul></td></tr></table>
    RoleAttrExtraction    *string              `json:"role_attr_extraction,omitempty"`
    // name of the attribute in SAML Assertion to extract role from
    RoleAttrFrom          *string              `json:"role_attr_from,omitempty"`
    // indicates if SCIM provisioning is enabled for the OAuth IDP
    ScimEnabled           *bool                `json:"scim_enabled,omitempty"`
    // scim_secret_token (generated by caller, crypto-random) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP
    ScimSecretToken       *string              `json:"scim_secret_token,omitempty"`
    SiteId                *uuid.UUID           `json:"site_id,omitempty"`
    Type                  *string              `json:"type,omitempty"`
    AdditionalProperties  map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Sso.
// It customizes the JSON marshaling process for Sso objects.
func (s Sso) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Sso object to a map representation for JSON marshaling.
func (s Sso) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.LdapCerts != nil {
        structMap["ldap_certs"] = s.LdapCerts
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
    if s.LdapGroupFilter != nil {
        structMap["ldap_group_filter"] = s.LdapGroupFilter
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
    if s.Type != nil {
        structMap["type"] = s.Type
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "custom_logout_url", "default_role", "domain", "id", "idp_cert", "idp_sign_algo", "idp_sso_url", "idp_type", "ignore_unmatched_roles", "issuer", "ldap_base_dn", "ldap_bind_dn", "ldap_bind_password", "ldap_certs", "ldap_client_cert", "ldap_client_key", "ldap_group_attr", "ldap_group_dn", "ldap_group_filter", "ldap_resolve_groups", "ldap_server_hosts", "ldap_type", "ldap_user_filter", "modified_time", "msp_id", "mxedge_proxy", "name", "nameid_format", "oauth_cc_client_id", "oauth_cc_client_secret", "oauth_discovery_url", "oauth_ropc_client_id", "oauth_ropc_client_secret", "oauth_tenant_id", "oauth_type", "org_id", "role_attr_extraction", "role_attr_from", "scim_enabled", "scim_secret_token", "site_id", "type")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CreatedTime = temp.CreatedTime
    s.CustomLogoutUrl = temp.CustomLogoutUrl
    s.DefaultRole = temp.DefaultRole
    s.Domain = temp.Domain
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
    s.LdapCerts = temp.LdapCerts
    s.LdapClientCert = temp.LdapClientCert
    s.LdapClientKey = temp.LdapClientKey
    s.LdapGroupAttr = temp.LdapGroupAttr
    s.LdapGroupDn = temp.LdapGroupDn
    s.LdapGroupFilter = temp.LdapGroupFilter
    s.LdapResolveGroups = temp.LdapResolveGroups
    s.LdapServerHosts = temp.LdapServerHosts
    s.LdapType = temp.LdapType
    s.LdapUserFilter = temp.LdapUserFilter
    s.ModifiedTime = temp.ModifiedTime
    s.MspId = temp.MspId
    s.MxedgeProxy = temp.MxedgeProxy
    s.Name = *temp.Name
    s.NameidFormat = temp.NameidFormat
    s.OauthCcClientId = temp.OauthCcClientId
    s.OauthCcClientSecret = temp.OauthCcClientSecret
    s.OauthDiscoveryUrl = temp.OauthDiscoveryUrl
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
    s.Type = temp.Type
    return nil
}

// tempSso is a temporary struct used for validating the fields of Sso.
type tempSso  struct {
    CreatedTime           *float64             `json:"created_time,omitempty"`
    CustomLogoutUrl       *string              `json:"custom_logout_url,omitempty"`
    DefaultRole           *string              `json:"default_role,omitempty"`
    Domain                *string              `json:"domain,omitempty"`
    Id                    *uuid.UUID           `json:"id,omitempty"`
    IdpCert               *string              `json:"idp_cert,omitempty"`
    IdpSignAlgo           *string              `json:"idp_sign_algo,omitempty"`
    IdpSsoUrl             *string              `json:"idp_sso_url,omitempty"`
    IdpType               *SsoIdpTypeEnum      `json:"idp_type,omitempty"`
    IgnoreUnmatchedRoles  *bool                `json:"ignore_unmatched_roles,omitempty"`
    Issuer                *string              `json:"issuer,omitempty"`
    LdapBaseDn            *string              `json:"ldap_base_dn,omitempty"`
    LdapBindDn            *string              `json:"ldap_bind_dn,omitempty"`
    LdapBindPassword      *string              `json:"ldap_bind_password,omitempty"`
    LdapCerts             []string             `json:"ldap_certs,omitempty"`
    LdapClientCert        *string              `json:"ldap_client_cert,omitempty"`
    LdapClientKey         *string              `json:"ldap_client_key,omitempty"`
    LdapGroupAttr         *string              `json:"ldap_group_attr,omitempty"`
    LdapGroupDn           *string              `json:"ldap_group_dn,omitempty"`
    LdapGroupFilter       *string              `json:"ldap_group_filter,omitempty"`
    LdapResolveGroups     *bool                `json:"ldap_resolve_groups,omitempty"`
    LdapServerHosts       []string             `json:"ldap_server_hosts,omitempty"`
    LdapType              *SsoLdapTypeEnum     `json:"ldap_type,omitempty"`
    LdapUserFilter        *string              `json:"ldap_user_filter,omitempty"`
    ModifiedTime          *float64             `json:"modified_time,omitempty"`
    MspId                 *uuid.UUID           `json:"msp_id,omitempty"`
    MxedgeProxy           *SsoMxedgeProxy      `json:"mxedge_proxy,omitempty"`
    Name                  *string              `json:"name"`
    NameidFormat          *SsoNameidFormatEnum `json:"nameid_format,omitempty"`
    OauthCcClientId       *string              `json:"oauth_cc_client_id,omitempty"`
    OauthCcClientSecret   *string              `json:"oauth_cc_client_secret,omitempty"`
    OauthDiscoveryUrl     *string              `json:"oauth_discovery_url,omitempty"`
    OauthRopcClientId     *string              `json:"oauth_ropc_client_id,omitempty"`
    OauthRopcClientSecret *string              `json:"oauth_ropc_client_secret,omitempty"`
    OauthTenantId         *string              `json:"oauth_tenant_id,omitempty"`
    OauthType             *SsoOauthTypeEnum    `json:"oauth_type,omitempty"`
    OrgId                 *uuid.UUID           `json:"org_id,omitempty"`
    RoleAttrExtraction    *string              `json:"role_attr_extraction,omitempty"`
    RoleAttrFrom          *string              `json:"role_attr_from,omitempty"`
    ScimEnabled           *bool                `json:"scim_enabled,omitempty"`
    ScimSecretToken       *string              `json:"scim_secret_token,omitempty"`
    SiteId                *uuid.UUID           `json:"site_id,omitempty"`
    Type                  *string              `json:"type,omitempty"`
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
