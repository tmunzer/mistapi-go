# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | whether webhook is enabled | [optional] [default to true]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Headers** | Pointer to **map[string]string** | if &#x60;type&#x60;&#x3D;&#x60;http-post&#x60;, additional custom HTTP headers to add the headers name and value must be string, total bytes of headers name and value must be less than 1000 | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | name of the webhook | [optional] 
**Oauth2ClientId** | Pointer to **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;client_credentials&#x60; | [optional] 
**Oauth2ClientSecret** | Pointer to **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;client_credentials&#x60; | [optional] 
**Oauth2GrantType** | Pointer to [**WebhookOauth2GrantType**](WebhookOauth2GrantType.md) |  | [optional] 
**Oauth2Password** | Pointer to **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;password&#x60; | [optional] 
**Oauth2Scopes** | Pointer to **[]string** | required when &#x60;type&#x60;&#x3D;&#x3D;&#x60;oauth2&#x60;, if provided, will be used in the token request | [optional] 
**Oauth2TokenUrl** | Pointer to **string** | required when &#x60;type&#x60;&#x3D;&#x3D;&#x60;oauth2&#x60; | [optional] 
**Oauth2Username** | Pointer to **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;password&#x60; | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Secret** | Pointer to **NullableString** | only if &#x60;type&#x60;&#x3D;&#x60;http-post&#x60; when &#x60;secret&#x60; is provided, two HTTP headers will be added:   * X-Mist-Signature-v2: HMAC_SHA256(secret, body)   * X-Mist-Signature: HMAC_SHA1(secret, body) | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SplunkToken** | Pointer to **NullableString** | required if &#x60;type&#x60;&#x3D;&#x60;splunk&#x60; If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.&#39; | [optional] 
**Topics** | Pointer to [**[]WebhookTopic**](WebhookTopic.md) | N.B. For org webhooks, only device_events/alarms/audits/client-join/client-sessions/nac-sessions/nac_events topics are supported. | [optional] 
**Type** | Pointer to [**WebhookType**](WebhookType.md) |  | [optional] [default to WEBHOOKTYPE_HTTP_POST]
**Url** | Pointer to **NullableString** |  | [optional] 
**VerifyCert** | Pointer to **bool** | when url uses HTTPS, whether to verify the certificate | [optional] [default to true]

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Webhook) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Webhook) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Webhook) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Webhook) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForSite

`func (o *Webhook) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Webhook) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Webhook) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Webhook) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeaders

`func (o *Webhook) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Webhook) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Webhook) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Webhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *Webhook) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *Webhook) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Webhook) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Webhook) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Webhook) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Webhook) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Webhook) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Webhook) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Webhook) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOauth2ClientId

`func (o *Webhook) GetOauth2ClientId() string`

GetOauth2ClientId returns the Oauth2ClientId field if non-nil, zero value otherwise.

### GetOauth2ClientIdOk

`func (o *Webhook) GetOauth2ClientIdOk() (*string, bool)`

GetOauth2ClientIdOk returns a tuple with the Oauth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientId

`func (o *Webhook) SetOauth2ClientId(v string)`

SetOauth2ClientId sets Oauth2ClientId field to given value.

### HasOauth2ClientId

`func (o *Webhook) HasOauth2ClientId() bool`

HasOauth2ClientId returns a boolean if a field has been set.

### GetOauth2ClientSecret

`func (o *Webhook) GetOauth2ClientSecret() string`

GetOauth2ClientSecret returns the Oauth2ClientSecret field if non-nil, zero value otherwise.

### GetOauth2ClientSecretOk

`func (o *Webhook) GetOauth2ClientSecretOk() (*string, bool)`

GetOauth2ClientSecretOk returns a tuple with the Oauth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientSecret

`func (o *Webhook) SetOauth2ClientSecret(v string)`

SetOauth2ClientSecret sets Oauth2ClientSecret field to given value.

### HasOauth2ClientSecret

`func (o *Webhook) HasOauth2ClientSecret() bool`

HasOauth2ClientSecret returns a boolean if a field has been set.

### GetOauth2GrantType

`func (o *Webhook) GetOauth2GrantType() WebhookOauth2GrantType`

GetOauth2GrantType returns the Oauth2GrantType field if non-nil, zero value otherwise.

### GetOauth2GrantTypeOk

`func (o *Webhook) GetOauth2GrantTypeOk() (*WebhookOauth2GrantType, bool)`

GetOauth2GrantTypeOk returns a tuple with the Oauth2GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2GrantType

`func (o *Webhook) SetOauth2GrantType(v WebhookOauth2GrantType)`

SetOauth2GrantType sets Oauth2GrantType field to given value.

### HasOauth2GrantType

`func (o *Webhook) HasOauth2GrantType() bool`

HasOauth2GrantType returns a boolean if a field has been set.

### GetOauth2Password

`func (o *Webhook) GetOauth2Password() string`

GetOauth2Password returns the Oauth2Password field if non-nil, zero value otherwise.

### GetOauth2PasswordOk

`func (o *Webhook) GetOauth2PasswordOk() (*string, bool)`

GetOauth2PasswordOk returns a tuple with the Oauth2Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Password

`func (o *Webhook) SetOauth2Password(v string)`

SetOauth2Password sets Oauth2Password field to given value.

### HasOauth2Password

`func (o *Webhook) HasOauth2Password() bool`

HasOauth2Password returns a boolean if a field has been set.

### GetOauth2Scopes

`func (o *Webhook) GetOauth2Scopes() []string`

GetOauth2Scopes returns the Oauth2Scopes field if non-nil, zero value otherwise.

### GetOauth2ScopesOk

`func (o *Webhook) GetOauth2ScopesOk() (*[]string, bool)`

GetOauth2ScopesOk returns a tuple with the Oauth2Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Scopes

`func (o *Webhook) SetOauth2Scopes(v []string)`

SetOauth2Scopes sets Oauth2Scopes field to given value.

### HasOauth2Scopes

`func (o *Webhook) HasOauth2Scopes() bool`

HasOauth2Scopes returns a boolean if a field has been set.

### GetOauth2TokenUrl

`func (o *Webhook) GetOauth2TokenUrl() string`

GetOauth2TokenUrl returns the Oauth2TokenUrl field if non-nil, zero value otherwise.

### GetOauth2TokenUrlOk

`func (o *Webhook) GetOauth2TokenUrlOk() (*string, bool)`

GetOauth2TokenUrlOk returns a tuple with the Oauth2TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2TokenUrl

`func (o *Webhook) SetOauth2TokenUrl(v string)`

SetOauth2TokenUrl sets Oauth2TokenUrl field to given value.

### HasOauth2TokenUrl

`func (o *Webhook) HasOauth2TokenUrl() bool`

HasOauth2TokenUrl returns a boolean if a field has been set.

### GetOauth2Username

`func (o *Webhook) GetOauth2Username() string`

GetOauth2Username returns the Oauth2Username field if non-nil, zero value otherwise.

### GetOauth2UsernameOk

`func (o *Webhook) GetOauth2UsernameOk() (*string, bool)`

GetOauth2UsernameOk returns a tuple with the Oauth2Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Username

`func (o *Webhook) SetOauth2Username(v string)`

SetOauth2Username sets Oauth2Username field to given value.

### HasOauth2Username

`func (o *Webhook) HasOauth2Username() bool`

HasOauth2Username returns a boolean if a field has been set.

### GetOrgId

`func (o *Webhook) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Webhook) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Webhook) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Webhook) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Webhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *Webhook) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *Webhook) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetSiteId

`func (o *Webhook) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Webhook) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Webhook) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Webhook) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSplunkToken

`func (o *Webhook) GetSplunkToken() string`

GetSplunkToken returns the SplunkToken field if non-nil, zero value otherwise.

### GetSplunkTokenOk

`func (o *Webhook) GetSplunkTokenOk() (*string, bool)`

GetSplunkTokenOk returns a tuple with the SplunkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkToken

`func (o *Webhook) SetSplunkToken(v string)`

SetSplunkToken sets SplunkToken field to given value.

### HasSplunkToken

`func (o *Webhook) HasSplunkToken() bool`

HasSplunkToken returns a boolean if a field has been set.

### SetSplunkTokenNil

`func (o *Webhook) SetSplunkTokenNil(b bool)`

 SetSplunkTokenNil sets the value for SplunkToken to be an explicit nil

### UnsetSplunkToken
`func (o *Webhook) UnsetSplunkToken()`

UnsetSplunkToken ensures that no value is present for SplunkToken, not even an explicit nil
### GetTopics

`func (o *Webhook) GetTopics() []WebhookTopic`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *Webhook) GetTopicsOk() (*[]WebhookTopic, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *Webhook) SetTopics(v []WebhookTopic)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *Webhook) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetType

`func (o *Webhook) GetType() WebhookType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Webhook) GetTypeOk() (*WebhookType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Webhook) SetType(v WebhookType)`

SetType sets Type field to given value.

### HasType

`func (o *Webhook) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Webhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Webhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVerifyCert

`func (o *Webhook) GetVerifyCert() bool`

GetVerifyCert returns the VerifyCert field if non-nil, zero value otherwise.

### GetVerifyCertOk

`func (o *Webhook) GetVerifyCertOk() (*bool, bool)`

GetVerifyCertOk returns a tuple with the VerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCert

`func (o *Webhook) SetVerifyCert(v bool)`

SetVerifyCert sets VerifyCert field to given value.

### HasVerifyCert

`func (o *Webhook) HasVerifyCert() bool`

HasVerifyCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


