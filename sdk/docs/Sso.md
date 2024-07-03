# Sso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**CustomLogoutUrl** | Pointer to **string** | optional, a URL we will redirect the user after user logout from Mist (for some IdP which supports a custom logout URL that is different from SP-initiated SLO process) | [optional] 
**DefaultRole** | Pointer to **string** | default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IdpCert** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;saml&#x60;. IDP Cert (used to verify the signed response) | [optional] 
**IdpSignAlgo** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;saml&#x60;. Signing algorithm for SAML Assertion | [optional] 
**IdpSsoUrl** | Pointer to **string** | IDP Single-Sign-On URL | [optional] 
**IdpType** | Pointer to [**SsoIdpType**](SsoIdpType.md) |  | [optional] [default to SSOIDPTYPE_SAML]
**IgnoreUnmatchedRoles** | Pointer to **bool** | ignore any unmatched roles provided in assertion. By default, an assertion is treated as invalid for any unmatched role | [optional] 
**Issuer** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;saml&#x60;. IDP issuer URL | [optional] 
**LdapBaseDn** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapBindDn** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapBindPassword** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapCerts** | Pointer to **[]string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapClientCert** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapClientKey** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapGroupAttr** | Pointer to **string** | Only if &#x60;ldap_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] 
**LdapGroupDn** | Pointer to **string** | Only if &#x60;ldap_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] 
**LdapGroupFilter** | Pointer to **string** | Only if &#x60;ldap_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] 
**LdapResolveGroups** | Pointer to **bool** | whether to recursively resolve LDAP groups | [optional] [default to false]
**LdapServerHosts** | Pointer to **[]string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;ldap&#x60; | [optional] 
**LdapType** | Pointer to [**SsoLdapType**](SsoLdapType.md) |  | [optional] [default to SSOLDAPTYPE_AZURE]
**LdapUserFilter** | Pointer to **string** | Only if &#x60;ldap_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**MxedgeProxy** | Pointer to [**SsoMxedgeProxy**](SsoMxedgeProxy.md) |  | [optional] 
**Name** | **string** | name | 
**NameidFormat** | Pointer to [**SsoNameidFormat**](SsoNameidFormat.md) |  | [optional] [default to SSONAMEIDFORMAT_EMAIL]
**OauthCcClientId** | Pointer to **string** | if &#x60;oauth_type&#x60;&#x3D;&#x3D;&#x60;okta&#x60;, Client Credentials | [optional] 
**OauthCcClientSecret** | Pointer to **string** | if &#x60;oauth_type&#x60;&#x3D;&#x3D;&#x60;okta&#x60;, oauth_cc_client_secret is RSA private key, of the form \&quot;-----BEGIN RSA PRIVATE KEY--....\&quot; | [optional] 
**OauthDiscoveryUrl** | Pointer to **string** | if &#x60;idp_type&#x60;&#x3D;&#x3D;&#x60;oauth&#x60; | [optional] 
**OauthRopcClientId** | Pointer to **string** | ropc &#x3D; Resource Owner Password Credentials | [optional] 
**OauthRopcClientSecret** | Pointer to **string** | oauth_ropc_client_secret can be empty if oauth_type is azure or azure-gov | [optional] 
**OauthTenantId** | Pointer to **string** | if &#x60;oauth_type&#x60;&#x3D;&#x3D;&#x60;okta&#x60;, oauth_tenant_id | [optional] 
**OauthType** | Pointer to [**SsoOauthType**](SsoOauthType.md) |  | [optional] [default to SSOOAUTHTYPE_AZURE]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RoleAttrExtraction** | Pointer to **string** | optional, custom role attribute parsing scheme  Supported Role Parsing Schemes &lt;table&gt;&lt;tr&gt;&lt;th&gt;Name&lt;/th&gt;&lt;th&gt;Scheme&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;cn&lt;/td&gt;&lt;td&gt;&lt;ul&gt;&lt;li&gt;The expected role attribute format in SAML Assertion is “CN&#x3D;cn,OU&#x3D;ou1,OU&#x3D;ou2,…”&lt;/li&gt;&lt;li&gt;CN (the key) is case insensitive and exactly 1 CN is expected (or the entire entry will be ignored)&lt;/li&gt;&lt;li&gt;E.g. if role attribute is “CN&#x3D;cn,OU&#x3D;ou1,OU&#x3D;ou2” then parsed role value is “cn”&lt;/li&gt;&lt;/ul&gt;&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | [optional] 
**RoleAttrFrom** | Pointer to **string** | name of the attribute in SAML Assertion to extract role from | [optional] [default to "role"]
**ScimEnabled** | Pointer to **bool** | indicates if SCIM provisioning is enabled for the OAuth IDP | [optional] [default to false]
**ScimSecretToken** | Pointer to **string** | scim_secret_token (generated by caller, crypto-random) is used as the Bearer token in the Authorization header of SCIM provisioning requests by the IDP | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSso

`func NewSso(name string, ) *Sso`

NewSso instantiates a new Sso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoWithDefaults

`func NewSsoWithDefaults() *Sso`

NewSsoWithDefaults instantiates a new Sso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Sso) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Sso) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Sso) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Sso) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomLogoutUrl

`func (o *Sso) GetCustomLogoutUrl() string`

GetCustomLogoutUrl returns the CustomLogoutUrl field if non-nil, zero value otherwise.

### GetCustomLogoutUrlOk

`func (o *Sso) GetCustomLogoutUrlOk() (*string, bool)`

GetCustomLogoutUrlOk returns a tuple with the CustomLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogoutUrl

`func (o *Sso) SetCustomLogoutUrl(v string)`

SetCustomLogoutUrl sets CustomLogoutUrl field to given value.

### HasCustomLogoutUrl

`func (o *Sso) HasCustomLogoutUrl() bool`

HasCustomLogoutUrl returns a boolean if a field has been set.

### GetDefaultRole

`func (o *Sso) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *Sso) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *Sso) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *Sso) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDomain

`func (o *Sso) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Sso) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Sso) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Sso) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *Sso) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sso) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sso) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sso) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpCert

`func (o *Sso) GetIdpCert() string`

GetIdpCert returns the IdpCert field if non-nil, zero value otherwise.

### GetIdpCertOk

`func (o *Sso) GetIdpCertOk() (*string, bool)`

GetIdpCertOk returns a tuple with the IdpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCert

`func (o *Sso) SetIdpCert(v string)`

SetIdpCert sets IdpCert field to given value.

### HasIdpCert

`func (o *Sso) HasIdpCert() bool`

HasIdpCert returns a boolean if a field has been set.

### GetIdpSignAlgo

`func (o *Sso) GetIdpSignAlgo() string`

GetIdpSignAlgo returns the IdpSignAlgo field if non-nil, zero value otherwise.

### GetIdpSignAlgoOk

`func (o *Sso) GetIdpSignAlgoOk() (*string, bool)`

GetIdpSignAlgoOk returns a tuple with the IdpSignAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSignAlgo

`func (o *Sso) SetIdpSignAlgo(v string)`

SetIdpSignAlgo sets IdpSignAlgo field to given value.

### HasIdpSignAlgo

`func (o *Sso) HasIdpSignAlgo() bool`

HasIdpSignAlgo returns a boolean if a field has been set.

### GetIdpSsoUrl

`func (o *Sso) GetIdpSsoUrl() string`

GetIdpSsoUrl returns the IdpSsoUrl field if non-nil, zero value otherwise.

### GetIdpSsoUrlOk

`func (o *Sso) GetIdpSsoUrlOk() (*string, bool)`

GetIdpSsoUrlOk returns a tuple with the IdpSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoUrl

`func (o *Sso) SetIdpSsoUrl(v string)`

SetIdpSsoUrl sets IdpSsoUrl field to given value.

### HasIdpSsoUrl

`func (o *Sso) HasIdpSsoUrl() bool`

HasIdpSsoUrl returns a boolean if a field has been set.

### GetIdpType

`func (o *Sso) GetIdpType() SsoIdpType`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *Sso) GetIdpTypeOk() (*SsoIdpType, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *Sso) SetIdpType(v SsoIdpType)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *Sso) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.

### GetIgnoreUnmatchedRoles

`func (o *Sso) GetIgnoreUnmatchedRoles() bool`

GetIgnoreUnmatchedRoles returns the IgnoreUnmatchedRoles field if non-nil, zero value otherwise.

### GetIgnoreUnmatchedRolesOk

`func (o *Sso) GetIgnoreUnmatchedRolesOk() (*bool, bool)`

GetIgnoreUnmatchedRolesOk returns a tuple with the IgnoreUnmatchedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUnmatchedRoles

`func (o *Sso) SetIgnoreUnmatchedRoles(v bool)`

SetIgnoreUnmatchedRoles sets IgnoreUnmatchedRoles field to given value.

### HasIgnoreUnmatchedRoles

`func (o *Sso) HasIgnoreUnmatchedRoles() bool`

HasIgnoreUnmatchedRoles returns a boolean if a field has been set.

### GetIssuer

`func (o *Sso) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Sso) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Sso) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Sso) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLdapBaseDn

`func (o *Sso) GetLdapBaseDn() string`

GetLdapBaseDn returns the LdapBaseDn field if non-nil, zero value otherwise.

### GetLdapBaseDnOk

`func (o *Sso) GetLdapBaseDnOk() (*string, bool)`

GetLdapBaseDnOk returns a tuple with the LdapBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDn

`func (o *Sso) SetLdapBaseDn(v string)`

SetLdapBaseDn sets LdapBaseDn field to given value.

### HasLdapBaseDn

`func (o *Sso) HasLdapBaseDn() bool`

HasLdapBaseDn returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *Sso) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *Sso) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *Sso) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *Sso) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *Sso) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *Sso) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *Sso) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *Sso) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCerts

`func (o *Sso) GetLdapCerts() []string`

GetLdapCerts returns the LdapCerts field if non-nil, zero value otherwise.

### GetLdapCertsOk

`func (o *Sso) GetLdapCertsOk() (*[]string, bool)`

GetLdapCertsOk returns a tuple with the LdapCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCerts

`func (o *Sso) SetLdapCerts(v []string)`

SetLdapCerts sets LdapCerts field to given value.

### HasLdapCerts

`func (o *Sso) HasLdapCerts() bool`

HasLdapCerts returns a boolean if a field has been set.

### GetLdapClientCert

`func (o *Sso) GetLdapClientCert() string`

GetLdapClientCert returns the LdapClientCert field if non-nil, zero value otherwise.

### GetLdapClientCertOk

`func (o *Sso) GetLdapClientCertOk() (*string, bool)`

GetLdapClientCertOk returns a tuple with the LdapClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapClientCert

`func (o *Sso) SetLdapClientCert(v string)`

SetLdapClientCert sets LdapClientCert field to given value.

### HasLdapClientCert

`func (o *Sso) HasLdapClientCert() bool`

HasLdapClientCert returns a boolean if a field has been set.

### GetLdapClientKey

`func (o *Sso) GetLdapClientKey() string`

GetLdapClientKey returns the LdapClientKey field if non-nil, zero value otherwise.

### GetLdapClientKeyOk

`func (o *Sso) GetLdapClientKeyOk() (*string, bool)`

GetLdapClientKeyOk returns a tuple with the LdapClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapClientKey

`func (o *Sso) SetLdapClientKey(v string)`

SetLdapClientKey sets LdapClientKey field to given value.

### HasLdapClientKey

`func (o *Sso) HasLdapClientKey() bool`

HasLdapClientKey returns a boolean if a field has been set.

### GetLdapGroupAttr

`func (o *Sso) GetLdapGroupAttr() string`

GetLdapGroupAttr returns the LdapGroupAttr field if non-nil, zero value otherwise.

### GetLdapGroupAttrOk

`func (o *Sso) GetLdapGroupAttrOk() (*string, bool)`

GetLdapGroupAttrOk returns a tuple with the LdapGroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttr

`func (o *Sso) SetLdapGroupAttr(v string)`

SetLdapGroupAttr sets LdapGroupAttr field to given value.

### HasLdapGroupAttr

`func (o *Sso) HasLdapGroupAttr() bool`

HasLdapGroupAttr returns a boolean if a field has been set.

### GetLdapGroupDn

`func (o *Sso) GetLdapGroupDn() string`

GetLdapGroupDn returns the LdapGroupDn field if non-nil, zero value otherwise.

### GetLdapGroupDnOk

`func (o *Sso) GetLdapGroupDnOk() (*string, bool)`

GetLdapGroupDnOk returns a tuple with the LdapGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupDn

`func (o *Sso) SetLdapGroupDn(v string)`

SetLdapGroupDn sets LdapGroupDn field to given value.

### HasLdapGroupDn

`func (o *Sso) HasLdapGroupDn() bool`

HasLdapGroupDn returns a boolean if a field has been set.

### GetLdapGroupFilter

`func (o *Sso) GetLdapGroupFilter() string`

GetLdapGroupFilter returns the LdapGroupFilter field if non-nil, zero value otherwise.

### GetLdapGroupFilterOk

`func (o *Sso) GetLdapGroupFilterOk() (*string, bool)`

GetLdapGroupFilterOk returns a tuple with the LdapGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupFilter

`func (o *Sso) SetLdapGroupFilter(v string)`

SetLdapGroupFilter sets LdapGroupFilter field to given value.

### HasLdapGroupFilter

`func (o *Sso) HasLdapGroupFilter() bool`

HasLdapGroupFilter returns a boolean if a field has been set.

### GetLdapResolveGroups

`func (o *Sso) GetLdapResolveGroups() bool`

GetLdapResolveGroups returns the LdapResolveGroups field if non-nil, zero value otherwise.

### GetLdapResolveGroupsOk

`func (o *Sso) GetLdapResolveGroupsOk() (*bool, bool)`

GetLdapResolveGroupsOk returns a tuple with the LdapResolveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapResolveGroups

`func (o *Sso) SetLdapResolveGroups(v bool)`

SetLdapResolveGroups sets LdapResolveGroups field to given value.

### HasLdapResolveGroups

`func (o *Sso) HasLdapResolveGroups() bool`

HasLdapResolveGroups returns a boolean if a field has been set.

### GetLdapServerHosts

`func (o *Sso) GetLdapServerHosts() []string`

GetLdapServerHosts returns the LdapServerHosts field if non-nil, zero value otherwise.

### GetLdapServerHostsOk

`func (o *Sso) GetLdapServerHostsOk() (*[]string, bool)`

GetLdapServerHostsOk returns a tuple with the LdapServerHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerHosts

`func (o *Sso) SetLdapServerHosts(v []string)`

SetLdapServerHosts sets LdapServerHosts field to given value.

### HasLdapServerHosts

`func (o *Sso) HasLdapServerHosts() bool`

HasLdapServerHosts returns a boolean if a field has been set.

### GetLdapType

`func (o *Sso) GetLdapType() SsoLdapType`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *Sso) GetLdapTypeOk() (*SsoLdapType, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *Sso) SetLdapType(v SsoLdapType)`

SetLdapType sets LdapType field to given value.

### HasLdapType

`func (o *Sso) HasLdapType() bool`

HasLdapType returns a boolean if a field has been set.

### GetLdapUserFilter

`func (o *Sso) GetLdapUserFilter() string`

GetLdapUserFilter returns the LdapUserFilter field if non-nil, zero value otherwise.

### GetLdapUserFilterOk

`func (o *Sso) GetLdapUserFilterOk() (*string, bool)`

GetLdapUserFilterOk returns a tuple with the LdapUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserFilter

`func (o *Sso) SetLdapUserFilter(v string)`

SetLdapUserFilter sets LdapUserFilter field to given value.

### HasLdapUserFilter

`func (o *Sso) HasLdapUserFilter() bool`

HasLdapUserFilter returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Sso) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Sso) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Sso) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Sso) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *Sso) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *Sso) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *Sso) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *Sso) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetMxedgeProxy

`func (o *Sso) GetMxedgeProxy() SsoMxedgeProxy`

GetMxedgeProxy returns the MxedgeProxy field if non-nil, zero value otherwise.

### GetMxedgeProxyOk

`func (o *Sso) GetMxedgeProxyOk() (*SsoMxedgeProxy, bool)`

GetMxedgeProxyOk returns a tuple with the MxedgeProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeProxy

`func (o *Sso) SetMxedgeProxy(v SsoMxedgeProxy)`

SetMxedgeProxy sets MxedgeProxy field to given value.

### HasMxedgeProxy

`func (o *Sso) HasMxedgeProxy() bool`

HasMxedgeProxy returns a boolean if a field has been set.

### GetName

`func (o *Sso) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sso) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sso) SetName(v string)`

SetName sets Name field to given value.


### GetNameidFormat

`func (o *Sso) GetNameidFormat() SsoNameidFormat`

GetNameidFormat returns the NameidFormat field if non-nil, zero value otherwise.

### GetNameidFormatOk

`func (o *Sso) GetNameidFormatOk() (*SsoNameidFormat, bool)`

GetNameidFormatOk returns a tuple with the NameidFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameidFormat

`func (o *Sso) SetNameidFormat(v SsoNameidFormat)`

SetNameidFormat sets NameidFormat field to given value.

### HasNameidFormat

`func (o *Sso) HasNameidFormat() bool`

HasNameidFormat returns a boolean if a field has been set.

### GetOauthCcClientId

`func (o *Sso) GetOauthCcClientId() string`

GetOauthCcClientId returns the OauthCcClientId field if non-nil, zero value otherwise.

### GetOauthCcClientIdOk

`func (o *Sso) GetOauthCcClientIdOk() (*string, bool)`

GetOauthCcClientIdOk returns a tuple with the OauthCcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthCcClientId

`func (o *Sso) SetOauthCcClientId(v string)`

SetOauthCcClientId sets OauthCcClientId field to given value.

### HasOauthCcClientId

`func (o *Sso) HasOauthCcClientId() bool`

HasOauthCcClientId returns a boolean if a field has been set.

### GetOauthCcClientSecret

`func (o *Sso) GetOauthCcClientSecret() string`

GetOauthCcClientSecret returns the OauthCcClientSecret field if non-nil, zero value otherwise.

### GetOauthCcClientSecretOk

`func (o *Sso) GetOauthCcClientSecretOk() (*string, bool)`

GetOauthCcClientSecretOk returns a tuple with the OauthCcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthCcClientSecret

`func (o *Sso) SetOauthCcClientSecret(v string)`

SetOauthCcClientSecret sets OauthCcClientSecret field to given value.

### HasOauthCcClientSecret

`func (o *Sso) HasOauthCcClientSecret() bool`

HasOauthCcClientSecret returns a boolean if a field has been set.

### GetOauthDiscoveryUrl

`func (o *Sso) GetOauthDiscoveryUrl() string`

GetOauthDiscoveryUrl returns the OauthDiscoveryUrl field if non-nil, zero value otherwise.

### GetOauthDiscoveryUrlOk

`func (o *Sso) GetOauthDiscoveryUrlOk() (*string, bool)`

GetOauthDiscoveryUrlOk returns a tuple with the OauthDiscoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthDiscoveryUrl

`func (o *Sso) SetOauthDiscoveryUrl(v string)`

SetOauthDiscoveryUrl sets OauthDiscoveryUrl field to given value.

### HasOauthDiscoveryUrl

`func (o *Sso) HasOauthDiscoveryUrl() bool`

HasOauthDiscoveryUrl returns a boolean if a field has been set.

### GetOauthRopcClientId

`func (o *Sso) GetOauthRopcClientId() string`

GetOauthRopcClientId returns the OauthRopcClientId field if non-nil, zero value otherwise.

### GetOauthRopcClientIdOk

`func (o *Sso) GetOauthRopcClientIdOk() (*string, bool)`

GetOauthRopcClientIdOk returns a tuple with the OauthRopcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRopcClientId

`func (o *Sso) SetOauthRopcClientId(v string)`

SetOauthRopcClientId sets OauthRopcClientId field to given value.

### HasOauthRopcClientId

`func (o *Sso) HasOauthRopcClientId() bool`

HasOauthRopcClientId returns a boolean if a field has been set.

### GetOauthRopcClientSecret

`func (o *Sso) GetOauthRopcClientSecret() string`

GetOauthRopcClientSecret returns the OauthRopcClientSecret field if non-nil, zero value otherwise.

### GetOauthRopcClientSecretOk

`func (o *Sso) GetOauthRopcClientSecretOk() (*string, bool)`

GetOauthRopcClientSecretOk returns a tuple with the OauthRopcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRopcClientSecret

`func (o *Sso) SetOauthRopcClientSecret(v string)`

SetOauthRopcClientSecret sets OauthRopcClientSecret field to given value.

### HasOauthRopcClientSecret

`func (o *Sso) HasOauthRopcClientSecret() bool`

HasOauthRopcClientSecret returns a boolean if a field has been set.

### GetOauthTenantId

`func (o *Sso) GetOauthTenantId() string`

GetOauthTenantId returns the OauthTenantId field if non-nil, zero value otherwise.

### GetOauthTenantIdOk

`func (o *Sso) GetOauthTenantIdOk() (*string, bool)`

GetOauthTenantIdOk returns a tuple with the OauthTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTenantId

`func (o *Sso) SetOauthTenantId(v string)`

SetOauthTenantId sets OauthTenantId field to given value.

### HasOauthTenantId

`func (o *Sso) HasOauthTenantId() bool`

HasOauthTenantId returns a boolean if a field has been set.

### GetOauthType

`func (o *Sso) GetOauthType() SsoOauthType`

GetOauthType returns the OauthType field if non-nil, zero value otherwise.

### GetOauthTypeOk

`func (o *Sso) GetOauthTypeOk() (*SsoOauthType, bool)`

GetOauthTypeOk returns a tuple with the OauthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthType

`func (o *Sso) SetOauthType(v SsoOauthType)`

SetOauthType sets OauthType field to given value.

### HasOauthType

`func (o *Sso) HasOauthType() bool`

HasOauthType returns a boolean if a field has been set.

### GetOrgId

`func (o *Sso) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Sso) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Sso) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Sso) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRoleAttrExtraction

`func (o *Sso) GetRoleAttrExtraction() string`

GetRoleAttrExtraction returns the RoleAttrExtraction field if non-nil, zero value otherwise.

### GetRoleAttrExtractionOk

`func (o *Sso) GetRoleAttrExtractionOk() (*string, bool)`

GetRoleAttrExtractionOk returns a tuple with the RoleAttrExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttrExtraction

`func (o *Sso) SetRoleAttrExtraction(v string)`

SetRoleAttrExtraction sets RoleAttrExtraction field to given value.

### HasRoleAttrExtraction

`func (o *Sso) HasRoleAttrExtraction() bool`

HasRoleAttrExtraction returns a boolean if a field has been set.

### GetRoleAttrFrom

`func (o *Sso) GetRoleAttrFrom() string`

GetRoleAttrFrom returns the RoleAttrFrom field if non-nil, zero value otherwise.

### GetRoleAttrFromOk

`func (o *Sso) GetRoleAttrFromOk() (*string, bool)`

GetRoleAttrFromOk returns a tuple with the RoleAttrFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttrFrom

`func (o *Sso) SetRoleAttrFrom(v string)`

SetRoleAttrFrom sets RoleAttrFrom field to given value.

### HasRoleAttrFrom

`func (o *Sso) HasRoleAttrFrom() bool`

HasRoleAttrFrom returns a boolean if a field has been set.

### GetScimEnabled

`func (o *Sso) GetScimEnabled() bool`

GetScimEnabled returns the ScimEnabled field if non-nil, zero value otherwise.

### GetScimEnabledOk

`func (o *Sso) GetScimEnabledOk() (*bool, bool)`

GetScimEnabledOk returns a tuple with the ScimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimEnabled

`func (o *Sso) SetScimEnabled(v bool)`

SetScimEnabled sets ScimEnabled field to given value.

### HasScimEnabled

`func (o *Sso) HasScimEnabled() bool`

HasScimEnabled returns a boolean if a field has been set.

### GetScimSecretToken

`func (o *Sso) GetScimSecretToken() string`

GetScimSecretToken returns the ScimSecretToken field if non-nil, zero value otherwise.

### GetScimSecretTokenOk

`func (o *Sso) GetScimSecretTokenOk() (*string, bool)`

GetScimSecretTokenOk returns a tuple with the ScimSecretToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimSecretToken

`func (o *Sso) SetScimSecretToken(v string)`

SetScimSecretToken sets ScimSecretToken field to given value.

### HasScimSecretToken

`func (o *Sso) HasScimSecretToken() bool`

HasScimSecretToken returns a boolean if a field has been set.

### GetSiteId

`func (o *Sso) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Sso) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Sso) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Sso) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetType

`func (o *Sso) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Sso) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Sso) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Sso) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


