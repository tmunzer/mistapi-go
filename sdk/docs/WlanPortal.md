# WlanPortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmazonClientId** | Pointer to **NullableString** | amazon OAuth2 client id. This is optional. If not provided, it will use a default one. | [optional] [default to ""]
**AmazonClientSecret** | Pointer to **NullableString** | amazon OAuth2 client secret. If amazon_client_id was provided, provide a correspoinding value. Else leave blank. | [optional] [default to ""]
**AmazonEmailDomains** | Pointer to **[]string** | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. | [optional] [default to []]
**AmazonEnabled** | Pointer to **bool** | whether amazon is enabled as a login method | [optional] [default to false]
**AmazonExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**Auth** | Pointer to [**WlanPortalAuth**](WlanPortalAuth.md) |  | [optional] [default to WLANPORTALAUTH_NONE]
**AzureClientId** | Pointer to **NullableString** | Required if &#x60;azure_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;. Azure active directory app client id | [optional] [default to ""]
**AzureClientSecret** | Pointer to **NullableString** | Required if &#x60;azure_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;. Azure active directory app client secret | [optional] [default to ""]
**AzureEnabled** | Pointer to **bool** | whether Azure Active Directory is enabled as a login method | [optional] [default to false]
**AzureExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**AzureTenantId** | Pointer to **NullableString** | Required if &#x60;azure_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;. Azure active directory tenant id. | [optional] [default to ""]
**BroadnetPassword** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;broadnet&#x60; | [optional] [default to ""]
**BroadnetSid** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;broadnet&#x60; | [optional] [default to "MIST"]
**BroadnetUserId** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;broadnet&#x60; | [optional] [default to ""]
**BypassWhenCloudDown** | Pointer to **bool** | whether to bypass the guest portal when cloud not reachable (and apply the default policies) | [optional] [default to false]
**ClickatellApiKey** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;clickatell&#x60; | [optional] [default to ""]
**CrossSite** | Pointer to **bool** | whether to allow guest to roam between WLANs (with same &#x60;WLAN.ssid&#x60;, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming) | [optional] [default to false]
**EmailEnabled** | Pointer to **bool** | whether email (access code verification) is enabled as a login method | [optional] 
**Enabled** | Pointer to **bool** | whether guest portal is enabled | [optional] [default to false]
**Expire** | Pointer to **float64** | how long to remain authorized, in minutes | [optional] [default to 1440]
**ExternalPortalUrl** | Pointer to **string** | external portal URL (e.g. https://host/url) where we can append our query parameters to | [optional] [default to ""]
**FacebookClientId** | Pointer to **NullableString** | Required if &#x60;facebook_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;. Facebook OAuth2 app id. This is optional. If not provided, it will use a default one. | [optional] [default to ""]
**FacebookClientSecret** | Pointer to **NullableString** | Required if &#x60;facebook_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;. Facebook OAuth2 app secret. If facebook_client_id was provided, provide a correspoinding value. Else leave blank. | [optional] [default to ""]
**FacebookEmailDomains** | Pointer to **[]string** | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. | [optional] [default to []]
**FacebookEnabled** | Pointer to **bool** | whether facebook is enabled as a login method | [optional] [default to false]
**FacebookExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**Forward** | Pointer to **bool** | whether to forward the user to another URL after authorized | [optional] [default to false]
**ForwardUrl** | Pointer to **NullableString** | the URL to forward the user to | [optional] [default to ""]
**GoogleClientId** | Pointer to **NullableString** | Google OAuth2 app id. This is optional. If not provided, it will use a default one. | [optional] [default to ""]
**GoogleClientSecret** | Pointer to **NullableString** | Google OAuth2 app secret. If google_client_id was provided, provide a correspoinding value. Else leave blank. | [optional] [default to ""]
**GoogleEmailDomains** | Pointer to **[]string** | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. | [optional] [default to []]
**GoogleEnabled** | Pointer to **bool** | whether google is enabled as login method | [optional] [default to false]
**GoogleExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using google auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**GupshupPassword** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;gupshup&#x60; | [optional] [default to ""]
**GupshupUserid** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;gupshup&#x60; | [optional] [default to ""]
**MicrosoftClientId** | Pointer to **NullableString** | microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one. | [optional] [default to ""]
**MicrosoftClientSecret** | Pointer to **NullableString** | microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a correspoinding value. Else leave blank. | [optional] [default to ""]
**MicrosoftEmailDomains** | Pointer to **[]string** | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. | [optional] [default to []]
**MicrosoftEnabled** | Pointer to **bool** | whether microsoft 365 is enabled as a login method | [optional] [default to false]
**MicrosoftExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**PassphraseEnabled** | Pointer to **bool** | whether password is enabled | [optional] [default to false]
**PassphraseExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses &#x60;expire&#x60; | [optional] 
**Password** | Pointer to **NullableString** | passphrase | [optional] [default to ""]
**PortalAllowedHostnames** | Pointer to **string** | list of hostnames without http(s):// (matched by substring) | [optional] [default to ""]
**PortalAllowedSubnets** | Pointer to **string** | list of CIDRs | [optional] [default to ""]
**PortalApiSecret** | Pointer to **string** | api secret (auto-generated) that can be used to sign guest authorization requests | [optional] [default to ""]
**PortalDeniedHostnames** | Pointer to **string** | list of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames | [optional] [default to ""]
**PortalImage** | Pointer to **string** | Url of portal background image | [optional] [default to ""]
**PortalSsoUrl** | Pointer to **string** | for SAML, this is used as the ACS URL | [optional] [readonly] 
**PredefinedSponsorsEnabled** | Pointer to **bool** | whether to show list of sponsor emails mentioned in &#x60;sponsors&#x60; object as a dropdown. If both &#x60;sponsor_notify_all&#x60; and &#x60;predefined_sponsors_enabled&#x60; are false, behaviour is acc to &#x60;sponsor_email_domains&#x60; | [optional] [default to true]
**Privacy** | Pointer to **bool** |  | [optional] [default to true]
**PuzzelPassword** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;puzzel&#x60; | [optional] [default to ""]
**PuzzelServiceId** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;puzzel&#x60; | [optional] [default to ""]
**PuzzelUsername** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;puzzel&#x60; | [optional] [default to ""]
**SmsMessageFormat** | Pointer to **string** |  | [optional] [default to ""]
**SmsEnabled** | Pointer to **bool** | whether sms is enabled as a login method | [optional] [default to false]
**SmsExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**SmsProvider** | Pointer to [**WlanPortalSmsProvider**](WlanPortalSmsProvider.md) |  | [optional] [default to WLANPORTALSMSPROVIDER_MANUAL]
**SponsorAutoApprove** | Pointer to **bool** | whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled | [optional] [default to false]
**SponsorEmailDomains** | Pointer to **[]string** | list of domain allowed for sponsor email. Required if &#x60;sponsor_enabled&#x60; is &#x60;true&#x60; and &#x60;sponsors&#x60; is empty. | [optional] [default to []]
**SponsorEnabled** | Pointer to **bool** | whether sponsor is enabled | [optional] [default to false]
**SponsorExpire** | Pointer to **NullableFloat64** | interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire&#x60; | [optional] 
**SponsorLinkValidityDuration** | Pointer to **int32** | how long to remain valid sponsored guest request approve/deny link received in email, in minutes. | [optional] [default to 60]
**SponsorNotifyAll** | Pointer to **bool** | whether to notify all sponsors that are mentioned in &#x60;sponsors&#x60; object. Both &#x60;sponsor_notify_all&#x60; and &#x60;predefined_sponsors_enabled&#x60; should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order. | [optional] [default to false]
**SponsorStatusNotify** | Pointer to **bool** | if enabled, guest will get email about sponsor&#39;s action (approve/deny) | [optional] [default to false]
**Sponsors** | Pointer to **map[string]string** | object of allowed sponsors email with name. Required if &#x60;sponsor_enabled&#x60; is &#x60;true&#x60; and &#x60;sponsor_email_domains&#x60; is empty. Property key is the sponsor email, Property value is the sponsor name | [optional] [default to {}]
**SsoDefaultRole** | Pointer to **string** | default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched | [optional] [default to ""]
**SsoForcedRole** | Pointer to **string** |  | [optional] [default to ""]
**SsoIdpCert** | Pointer to **string** | IDP Cert (used to verify the signed response) | [optional] [default to ""]
**SsoIdpSignAlgo** | Pointer to **string** | signing algorithm for SAML Assertion | [optional] [default to ""]
**SsoIdpSsoUrl** | Pointer to **string** | IDP Single-Sign-On URL | [optional] [default to ""]
**SsoIssuer** | Pointer to **string** | IDP issuer URL | [optional] [default to ""]
**SsoNameidFormat** | Pointer to [**WlanPortalSsoNameidFormat**](WlanPortalSsoNameidFormat.md) |  | [optional] [default to WLANPORTALSSONAMEIDFORMAT_EMAIL]
**TelstraClientId** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;telstra&#x60;, Client ID provided by Telstra | [optional] [default to ""]
**TelstraClientSecret** | Pointer to **string** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;telstra&#x60;, Client secret provided by Telstra | [optional] [default to ""]
**Thumbnail** | Pointer to **string** | Url of portal background image thumbnail | [optional] [default to ""]
**TwilioAuthToken** | Pointer to **NullableString** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;twilio&#x60;, Auth token account with twilio account | [optional] [default to ""]
**TwilioPhoneNumber** | Pointer to **NullableString** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;twilio&#x60;, Twilio phone number associated with the account. See example for accepted format. | [optional] [default to ""]
**TwilioSid** | Pointer to **NullableString** | when &#x60;sms_provider&#x60;&#x3D;&#x3D;&#x60;twilio&#x60;, Account SID provided by Twilio | [optional] [default to ""]

## Methods

### NewWlanPortal

`func NewWlanPortal() *WlanPortal`

NewWlanPortal instantiates a new WlanPortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPortalWithDefaults

`func NewWlanPortalWithDefaults() *WlanPortal`

NewWlanPortalWithDefaults instantiates a new WlanPortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazonClientId

`func (o *WlanPortal) GetAmazonClientId() string`

GetAmazonClientId returns the AmazonClientId field if non-nil, zero value otherwise.

### GetAmazonClientIdOk

`func (o *WlanPortal) GetAmazonClientIdOk() (*string, bool)`

GetAmazonClientIdOk returns a tuple with the AmazonClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonClientId

`func (o *WlanPortal) SetAmazonClientId(v string)`

SetAmazonClientId sets AmazonClientId field to given value.

### HasAmazonClientId

`func (o *WlanPortal) HasAmazonClientId() bool`

HasAmazonClientId returns a boolean if a field has been set.

### SetAmazonClientIdNil

`func (o *WlanPortal) SetAmazonClientIdNil(b bool)`

 SetAmazonClientIdNil sets the value for AmazonClientId to be an explicit nil

### UnsetAmazonClientId
`func (o *WlanPortal) UnsetAmazonClientId()`

UnsetAmazonClientId ensures that no value is present for AmazonClientId, not even an explicit nil
### GetAmazonClientSecret

`func (o *WlanPortal) GetAmazonClientSecret() string`

GetAmazonClientSecret returns the AmazonClientSecret field if non-nil, zero value otherwise.

### GetAmazonClientSecretOk

`func (o *WlanPortal) GetAmazonClientSecretOk() (*string, bool)`

GetAmazonClientSecretOk returns a tuple with the AmazonClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonClientSecret

`func (o *WlanPortal) SetAmazonClientSecret(v string)`

SetAmazonClientSecret sets AmazonClientSecret field to given value.

### HasAmazonClientSecret

`func (o *WlanPortal) HasAmazonClientSecret() bool`

HasAmazonClientSecret returns a boolean if a field has been set.

### SetAmazonClientSecretNil

`func (o *WlanPortal) SetAmazonClientSecretNil(b bool)`

 SetAmazonClientSecretNil sets the value for AmazonClientSecret to be an explicit nil

### UnsetAmazonClientSecret
`func (o *WlanPortal) UnsetAmazonClientSecret()`

UnsetAmazonClientSecret ensures that no value is present for AmazonClientSecret, not even an explicit nil
### GetAmazonEmailDomains

`func (o *WlanPortal) GetAmazonEmailDomains() []string`

GetAmazonEmailDomains returns the AmazonEmailDomains field if non-nil, zero value otherwise.

### GetAmazonEmailDomainsOk

`func (o *WlanPortal) GetAmazonEmailDomainsOk() (*[]string, bool)`

GetAmazonEmailDomainsOk returns a tuple with the AmazonEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonEmailDomains

`func (o *WlanPortal) SetAmazonEmailDomains(v []string)`

SetAmazonEmailDomains sets AmazonEmailDomains field to given value.

### HasAmazonEmailDomains

`func (o *WlanPortal) HasAmazonEmailDomains() bool`

HasAmazonEmailDomains returns a boolean if a field has been set.

### GetAmazonEnabled

`func (o *WlanPortal) GetAmazonEnabled() bool`

GetAmazonEnabled returns the AmazonEnabled field if non-nil, zero value otherwise.

### GetAmazonEnabledOk

`func (o *WlanPortal) GetAmazonEnabledOk() (*bool, bool)`

GetAmazonEnabledOk returns a tuple with the AmazonEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonEnabled

`func (o *WlanPortal) SetAmazonEnabled(v bool)`

SetAmazonEnabled sets AmazonEnabled field to given value.

### HasAmazonEnabled

`func (o *WlanPortal) HasAmazonEnabled() bool`

HasAmazonEnabled returns a boolean if a field has been set.

### GetAmazonExpire

`func (o *WlanPortal) GetAmazonExpire() float64`

GetAmazonExpire returns the AmazonExpire field if non-nil, zero value otherwise.

### GetAmazonExpireOk

`func (o *WlanPortal) GetAmazonExpireOk() (*float64, bool)`

GetAmazonExpireOk returns a tuple with the AmazonExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonExpire

`func (o *WlanPortal) SetAmazonExpire(v float64)`

SetAmazonExpire sets AmazonExpire field to given value.

### HasAmazonExpire

`func (o *WlanPortal) HasAmazonExpire() bool`

HasAmazonExpire returns a boolean if a field has been set.

### SetAmazonExpireNil

`func (o *WlanPortal) SetAmazonExpireNil(b bool)`

 SetAmazonExpireNil sets the value for AmazonExpire to be an explicit nil

### UnsetAmazonExpire
`func (o *WlanPortal) UnsetAmazonExpire()`

UnsetAmazonExpire ensures that no value is present for AmazonExpire, not even an explicit nil
### GetAuth

`func (o *WlanPortal) GetAuth() WlanPortalAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *WlanPortal) GetAuthOk() (*WlanPortalAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *WlanPortal) SetAuth(v WlanPortalAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *WlanPortal) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAzureClientId

`func (o *WlanPortal) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *WlanPortal) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *WlanPortal) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *WlanPortal) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### SetAzureClientIdNil

`func (o *WlanPortal) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *WlanPortal) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *WlanPortal) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *WlanPortal) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *WlanPortal) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *WlanPortal) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### SetAzureClientSecretNil

`func (o *WlanPortal) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *WlanPortal) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureEnabled

`func (o *WlanPortal) GetAzureEnabled() bool`

GetAzureEnabled returns the AzureEnabled field if non-nil, zero value otherwise.

### GetAzureEnabledOk

`func (o *WlanPortal) GetAzureEnabledOk() (*bool, bool)`

GetAzureEnabledOk returns a tuple with the AzureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEnabled

`func (o *WlanPortal) SetAzureEnabled(v bool)`

SetAzureEnabled sets AzureEnabled field to given value.

### HasAzureEnabled

`func (o *WlanPortal) HasAzureEnabled() bool`

HasAzureEnabled returns a boolean if a field has been set.

### GetAzureExpire

`func (o *WlanPortal) GetAzureExpire() float64`

GetAzureExpire returns the AzureExpire field if non-nil, zero value otherwise.

### GetAzureExpireOk

`func (o *WlanPortal) GetAzureExpireOk() (*float64, bool)`

GetAzureExpireOk returns a tuple with the AzureExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureExpire

`func (o *WlanPortal) SetAzureExpire(v float64)`

SetAzureExpire sets AzureExpire field to given value.

### HasAzureExpire

`func (o *WlanPortal) HasAzureExpire() bool`

HasAzureExpire returns a boolean if a field has been set.

### SetAzureExpireNil

`func (o *WlanPortal) SetAzureExpireNil(b bool)`

 SetAzureExpireNil sets the value for AzureExpire to be an explicit nil

### UnsetAzureExpire
`func (o *WlanPortal) UnsetAzureExpire()`

UnsetAzureExpire ensures that no value is present for AzureExpire, not even an explicit nil
### GetAzureTenantId

`func (o *WlanPortal) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *WlanPortal) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *WlanPortal) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *WlanPortal) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *WlanPortal) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *WlanPortal) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetBroadnetPassword

`func (o *WlanPortal) GetBroadnetPassword() string`

GetBroadnetPassword returns the BroadnetPassword field if non-nil, zero value otherwise.

### GetBroadnetPasswordOk

`func (o *WlanPortal) GetBroadnetPasswordOk() (*string, bool)`

GetBroadnetPasswordOk returns a tuple with the BroadnetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadnetPassword

`func (o *WlanPortal) SetBroadnetPassword(v string)`

SetBroadnetPassword sets BroadnetPassword field to given value.

### HasBroadnetPassword

`func (o *WlanPortal) HasBroadnetPassword() bool`

HasBroadnetPassword returns a boolean if a field has been set.

### GetBroadnetSid

`func (o *WlanPortal) GetBroadnetSid() string`

GetBroadnetSid returns the BroadnetSid field if non-nil, zero value otherwise.

### GetBroadnetSidOk

`func (o *WlanPortal) GetBroadnetSidOk() (*string, bool)`

GetBroadnetSidOk returns a tuple with the BroadnetSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadnetSid

`func (o *WlanPortal) SetBroadnetSid(v string)`

SetBroadnetSid sets BroadnetSid field to given value.

### HasBroadnetSid

`func (o *WlanPortal) HasBroadnetSid() bool`

HasBroadnetSid returns a boolean if a field has been set.

### GetBroadnetUserId

`func (o *WlanPortal) GetBroadnetUserId() string`

GetBroadnetUserId returns the BroadnetUserId field if non-nil, zero value otherwise.

### GetBroadnetUserIdOk

`func (o *WlanPortal) GetBroadnetUserIdOk() (*string, bool)`

GetBroadnetUserIdOk returns a tuple with the BroadnetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadnetUserId

`func (o *WlanPortal) SetBroadnetUserId(v string)`

SetBroadnetUserId sets BroadnetUserId field to given value.

### HasBroadnetUserId

`func (o *WlanPortal) HasBroadnetUserId() bool`

HasBroadnetUserId returns a boolean if a field has been set.

### GetBypassWhenCloudDown

`func (o *WlanPortal) GetBypassWhenCloudDown() bool`

GetBypassWhenCloudDown returns the BypassWhenCloudDown field if non-nil, zero value otherwise.

### GetBypassWhenCloudDownOk

`func (o *WlanPortal) GetBypassWhenCloudDownOk() (*bool, bool)`

GetBypassWhenCloudDownOk returns a tuple with the BypassWhenCloudDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassWhenCloudDown

`func (o *WlanPortal) SetBypassWhenCloudDown(v bool)`

SetBypassWhenCloudDown sets BypassWhenCloudDown field to given value.

### HasBypassWhenCloudDown

`func (o *WlanPortal) HasBypassWhenCloudDown() bool`

HasBypassWhenCloudDown returns a boolean if a field has been set.

### GetClickatellApiKey

`func (o *WlanPortal) GetClickatellApiKey() string`

GetClickatellApiKey returns the ClickatellApiKey field if non-nil, zero value otherwise.

### GetClickatellApiKeyOk

`func (o *WlanPortal) GetClickatellApiKeyOk() (*string, bool)`

GetClickatellApiKeyOk returns a tuple with the ClickatellApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickatellApiKey

`func (o *WlanPortal) SetClickatellApiKey(v string)`

SetClickatellApiKey sets ClickatellApiKey field to given value.

### HasClickatellApiKey

`func (o *WlanPortal) HasClickatellApiKey() bool`

HasClickatellApiKey returns a boolean if a field has been set.

### GetCrossSite

`func (o *WlanPortal) GetCrossSite() bool`

GetCrossSite returns the CrossSite field if non-nil, zero value otherwise.

### GetCrossSiteOk

`func (o *WlanPortal) GetCrossSiteOk() (*bool, bool)`

GetCrossSiteOk returns a tuple with the CrossSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSite

`func (o *WlanPortal) SetCrossSite(v bool)`

SetCrossSite sets CrossSite field to given value.

### HasCrossSite

`func (o *WlanPortal) HasCrossSite() bool`

HasCrossSite returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *WlanPortal) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *WlanPortal) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *WlanPortal) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *WlanPortal) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanPortal) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanPortal) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanPortal) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanPortal) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpire

`func (o *WlanPortal) GetExpire() float64`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *WlanPortal) GetExpireOk() (*float64, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *WlanPortal) SetExpire(v float64)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *WlanPortal) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetExternalPortalUrl

`func (o *WlanPortal) GetExternalPortalUrl() string`

GetExternalPortalUrl returns the ExternalPortalUrl field if non-nil, zero value otherwise.

### GetExternalPortalUrlOk

`func (o *WlanPortal) GetExternalPortalUrlOk() (*string, bool)`

GetExternalPortalUrlOk returns a tuple with the ExternalPortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortalUrl

`func (o *WlanPortal) SetExternalPortalUrl(v string)`

SetExternalPortalUrl sets ExternalPortalUrl field to given value.

### HasExternalPortalUrl

`func (o *WlanPortal) HasExternalPortalUrl() bool`

HasExternalPortalUrl returns a boolean if a field has been set.

### GetFacebookClientId

`func (o *WlanPortal) GetFacebookClientId() string`

GetFacebookClientId returns the FacebookClientId field if non-nil, zero value otherwise.

### GetFacebookClientIdOk

`func (o *WlanPortal) GetFacebookClientIdOk() (*string, bool)`

GetFacebookClientIdOk returns a tuple with the FacebookClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookClientId

`func (o *WlanPortal) SetFacebookClientId(v string)`

SetFacebookClientId sets FacebookClientId field to given value.

### HasFacebookClientId

`func (o *WlanPortal) HasFacebookClientId() bool`

HasFacebookClientId returns a boolean if a field has been set.

### SetFacebookClientIdNil

`func (o *WlanPortal) SetFacebookClientIdNil(b bool)`

 SetFacebookClientIdNil sets the value for FacebookClientId to be an explicit nil

### UnsetFacebookClientId
`func (o *WlanPortal) UnsetFacebookClientId()`

UnsetFacebookClientId ensures that no value is present for FacebookClientId, not even an explicit nil
### GetFacebookClientSecret

`func (o *WlanPortal) GetFacebookClientSecret() string`

GetFacebookClientSecret returns the FacebookClientSecret field if non-nil, zero value otherwise.

### GetFacebookClientSecretOk

`func (o *WlanPortal) GetFacebookClientSecretOk() (*string, bool)`

GetFacebookClientSecretOk returns a tuple with the FacebookClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookClientSecret

`func (o *WlanPortal) SetFacebookClientSecret(v string)`

SetFacebookClientSecret sets FacebookClientSecret field to given value.

### HasFacebookClientSecret

`func (o *WlanPortal) HasFacebookClientSecret() bool`

HasFacebookClientSecret returns a boolean if a field has been set.

### SetFacebookClientSecretNil

`func (o *WlanPortal) SetFacebookClientSecretNil(b bool)`

 SetFacebookClientSecretNil sets the value for FacebookClientSecret to be an explicit nil

### UnsetFacebookClientSecret
`func (o *WlanPortal) UnsetFacebookClientSecret()`

UnsetFacebookClientSecret ensures that no value is present for FacebookClientSecret, not even an explicit nil
### GetFacebookEmailDomains

`func (o *WlanPortal) GetFacebookEmailDomains() []string`

GetFacebookEmailDomains returns the FacebookEmailDomains field if non-nil, zero value otherwise.

### GetFacebookEmailDomainsOk

`func (o *WlanPortal) GetFacebookEmailDomainsOk() (*[]string, bool)`

GetFacebookEmailDomainsOk returns a tuple with the FacebookEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookEmailDomains

`func (o *WlanPortal) SetFacebookEmailDomains(v []string)`

SetFacebookEmailDomains sets FacebookEmailDomains field to given value.

### HasFacebookEmailDomains

`func (o *WlanPortal) HasFacebookEmailDomains() bool`

HasFacebookEmailDomains returns a boolean if a field has been set.

### GetFacebookEnabled

`func (o *WlanPortal) GetFacebookEnabled() bool`

GetFacebookEnabled returns the FacebookEnabled field if non-nil, zero value otherwise.

### GetFacebookEnabledOk

`func (o *WlanPortal) GetFacebookEnabledOk() (*bool, bool)`

GetFacebookEnabledOk returns a tuple with the FacebookEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookEnabled

`func (o *WlanPortal) SetFacebookEnabled(v bool)`

SetFacebookEnabled sets FacebookEnabled field to given value.

### HasFacebookEnabled

`func (o *WlanPortal) HasFacebookEnabled() bool`

HasFacebookEnabled returns a boolean if a field has been set.

### GetFacebookExpire

`func (o *WlanPortal) GetFacebookExpire() float64`

GetFacebookExpire returns the FacebookExpire field if non-nil, zero value otherwise.

### GetFacebookExpireOk

`func (o *WlanPortal) GetFacebookExpireOk() (*float64, bool)`

GetFacebookExpireOk returns a tuple with the FacebookExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookExpire

`func (o *WlanPortal) SetFacebookExpire(v float64)`

SetFacebookExpire sets FacebookExpire field to given value.

### HasFacebookExpire

`func (o *WlanPortal) HasFacebookExpire() bool`

HasFacebookExpire returns a boolean if a field has been set.

### SetFacebookExpireNil

`func (o *WlanPortal) SetFacebookExpireNil(b bool)`

 SetFacebookExpireNil sets the value for FacebookExpire to be an explicit nil

### UnsetFacebookExpire
`func (o *WlanPortal) UnsetFacebookExpire()`

UnsetFacebookExpire ensures that no value is present for FacebookExpire, not even an explicit nil
### GetForward

`func (o *WlanPortal) GetForward() bool`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *WlanPortal) GetForwardOk() (*bool, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *WlanPortal) SetForward(v bool)`

SetForward sets Forward field to given value.

### HasForward

`func (o *WlanPortal) HasForward() bool`

HasForward returns a boolean if a field has been set.

### GetForwardUrl

`func (o *WlanPortal) GetForwardUrl() string`

GetForwardUrl returns the ForwardUrl field if non-nil, zero value otherwise.

### GetForwardUrlOk

`func (o *WlanPortal) GetForwardUrlOk() (*string, bool)`

GetForwardUrlOk returns a tuple with the ForwardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUrl

`func (o *WlanPortal) SetForwardUrl(v string)`

SetForwardUrl sets ForwardUrl field to given value.

### HasForwardUrl

`func (o *WlanPortal) HasForwardUrl() bool`

HasForwardUrl returns a boolean if a field has been set.

### SetForwardUrlNil

`func (o *WlanPortal) SetForwardUrlNil(b bool)`

 SetForwardUrlNil sets the value for ForwardUrl to be an explicit nil

### UnsetForwardUrl
`func (o *WlanPortal) UnsetForwardUrl()`

UnsetForwardUrl ensures that no value is present for ForwardUrl, not even an explicit nil
### GetGoogleClientId

`func (o *WlanPortal) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *WlanPortal) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *WlanPortal) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.

### HasGoogleClientId

`func (o *WlanPortal) HasGoogleClientId() bool`

HasGoogleClientId returns a boolean if a field has been set.

### SetGoogleClientIdNil

`func (o *WlanPortal) SetGoogleClientIdNil(b bool)`

 SetGoogleClientIdNil sets the value for GoogleClientId to be an explicit nil

### UnsetGoogleClientId
`func (o *WlanPortal) UnsetGoogleClientId()`

UnsetGoogleClientId ensures that no value is present for GoogleClientId, not even an explicit nil
### GetGoogleClientSecret

`func (o *WlanPortal) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *WlanPortal) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *WlanPortal) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.

### HasGoogleClientSecret

`func (o *WlanPortal) HasGoogleClientSecret() bool`

HasGoogleClientSecret returns a boolean if a field has been set.

### SetGoogleClientSecretNil

`func (o *WlanPortal) SetGoogleClientSecretNil(b bool)`

 SetGoogleClientSecretNil sets the value for GoogleClientSecret to be an explicit nil

### UnsetGoogleClientSecret
`func (o *WlanPortal) UnsetGoogleClientSecret()`

UnsetGoogleClientSecret ensures that no value is present for GoogleClientSecret, not even an explicit nil
### GetGoogleEmailDomains

`func (o *WlanPortal) GetGoogleEmailDomains() []string`

GetGoogleEmailDomains returns the GoogleEmailDomains field if non-nil, zero value otherwise.

### GetGoogleEmailDomainsOk

`func (o *WlanPortal) GetGoogleEmailDomainsOk() (*[]string, bool)`

GetGoogleEmailDomainsOk returns a tuple with the GoogleEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleEmailDomains

`func (o *WlanPortal) SetGoogleEmailDomains(v []string)`

SetGoogleEmailDomains sets GoogleEmailDomains field to given value.

### HasGoogleEmailDomains

`func (o *WlanPortal) HasGoogleEmailDomains() bool`

HasGoogleEmailDomains returns a boolean if a field has been set.

### GetGoogleEnabled

`func (o *WlanPortal) GetGoogleEnabled() bool`

GetGoogleEnabled returns the GoogleEnabled field if non-nil, zero value otherwise.

### GetGoogleEnabledOk

`func (o *WlanPortal) GetGoogleEnabledOk() (*bool, bool)`

GetGoogleEnabledOk returns a tuple with the GoogleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleEnabled

`func (o *WlanPortal) SetGoogleEnabled(v bool)`

SetGoogleEnabled sets GoogleEnabled field to given value.

### HasGoogleEnabled

`func (o *WlanPortal) HasGoogleEnabled() bool`

HasGoogleEnabled returns a boolean if a field has been set.

### GetGoogleExpire

`func (o *WlanPortal) GetGoogleExpire() float64`

GetGoogleExpire returns the GoogleExpire field if non-nil, zero value otherwise.

### GetGoogleExpireOk

`func (o *WlanPortal) GetGoogleExpireOk() (*float64, bool)`

GetGoogleExpireOk returns a tuple with the GoogleExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleExpire

`func (o *WlanPortal) SetGoogleExpire(v float64)`

SetGoogleExpire sets GoogleExpire field to given value.

### HasGoogleExpire

`func (o *WlanPortal) HasGoogleExpire() bool`

HasGoogleExpire returns a boolean if a field has been set.

### SetGoogleExpireNil

`func (o *WlanPortal) SetGoogleExpireNil(b bool)`

 SetGoogleExpireNil sets the value for GoogleExpire to be an explicit nil

### UnsetGoogleExpire
`func (o *WlanPortal) UnsetGoogleExpire()`

UnsetGoogleExpire ensures that no value is present for GoogleExpire, not even an explicit nil
### GetGupshupPassword

`func (o *WlanPortal) GetGupshupPassword() string`

GetGupshupPassword returns the GupshupPassword field if non-nil, zero value otherwise.

### GetGupshupPasswordOk

`func (o *WlanPortal) GetGupshupPasswordOk() (*string, bool)`

GetGupshupPasswordOk returns a tuple with the GupshupPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGupshupPassword

`func (o *WlanPortal) SetGupshupPassword(v string)`

SetGupshupPassword sets GupshupPassword field to given value.

### HasGupshupPassword

`func (o *WlanPortal) HasGupshupPassword() bool`

HasGupshupPassword returns a boolean if a field has been set.

### GetGupshupUserid

`func (o *WlanPortal) GetGupshupUserid() string`

GetGupshupUserid returns the GupshupUserid field if non-nil, zero value otherwise.

### GetGupshupUseridOk

`func (o *WlanPortal) GetGupshupUseridOk() (*string, bool)`

GetGupshupUseridOk returns a tuple with the GupshupUserid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGupshupUserid

`func (o *WlanPortal) SetGupshupUserid(v string)`

SetGupshupUserid sets GupshupUserid field to given value.

### HasGupshupUserid

`func (o *WlanPortal) HasGupshupUserid() bool`

HasGupshupUserid returns a boolean if a field has been set.

### GetMicrosoftClientId

`func (o *WlanPortal) GetMicrosoftClientId() string`

GetMicrosoftClientId returns the MicrosoftClientId field if non-nil, zero value otherwise.

### GetMicrosoftClientIdOk

`func (o *WlanPortal) GetMicrosoftClientIdOk() (*string, bool)`

GetMicrosoftClientIdOk returns a tuple with the MicrosoftClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftClientId

`func (o *WlanPortal) SetMicrosoftClientId(v string)`

SetMicrosoftClientId sets MicrosoftClientId field to given value.

### HasMicrosoftClientId

`func (o *WlanPortal) HasMicrosoftClientId() bool`

HasMicrosoftClientId returns a boolean if a field has been set.

### SetMicrosoftClientIdNil

`func (o *WlanPortal) SetMicrosoftClientIdNil(b bool)`

 SetMicrosoftClientIdNil sets the value for MicrosoftClientId to be an explicit nil

### UnsetMicrosoftClientId
`func (o *WlanPortal) UnsetMicrosoftClientId()`

UnsetMicrosoftClientId ensures that no value is present for MicrosoftClientId, not even an explicit nil
### GetMicrosoftClientSecret

`func (o *WlanPortal) GetMicrosoftClientSecret() string`

GetMicrosoftClientSecret returns the MicrosoftClientSecret field if non-nil, zero value otherwise.

### GetMicrosoftClientSecretOk

`func (o *WlanPortal) GetMicrosoftClientSecretOk() (*string, bool)`

GetMicrosoftClientSecretOk returns a tuple with the MicrosoftClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftClientSecret

`func (o *WlanPortal) SetMicrosoftClientSecret(v string)`

SetMicrosoftClientSecret sets MicrosoftClientSecret field to given value.

### HasMicrosoftClientSecret

`func (o *WlanPortal) HasMicrosoftClientSecret() bool`

HasMicrosoftClientSecret returns a boolean if a field has been set.

### SetMicrosoftClientSecretNil

`func (o *WlanPortal) SetMicrosoftClientSecretNil(b bool)`

 SetMicrosoftClientSecretNil sets the value for MicrosoftClientSecret to be an explicit nil

### UnsetMicrosoftClientSecret
`func (o *WlanPortal) UnsetMicrosoftClientSecret()`

UnsetMicrosoftClientSecret ensures that no value is present for MicrosoftClientSecret, not even an explicit nil
### GetMicrosoftEmailDomains

`func (o *WlanPortal) GetMicrosoftEmailDomains() []string`

GetMicrosoftEmailDomains returns the MicrosoftEmailDomains field if non-nil, zero value otherwise.

### GetMicrosoftEmailDomainsOk

`func (o *WlanPortal) GetMicrosoftEmailDomainsOk() (*[]string, bool)`

GetMicrosoftEmailDomainsOk returns a tuple with the MicrosoftEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftEmailDomains

`func (o *WlanPortal) SetMicrosoftEmailDomains(v []string)`

SetMicrosoftEmailDomains sets MicrosoftEmailDomains field to given value.

### HasMicrosoftEmailDomains

`func (o *WlanPortal) HasMicrosoftEmailDomains() bool`

HasMicrosoftEmailDomains returns a boolean if a field has been set.

### GetMicrosoftEnabled

`func (o *WlanPortal) GetMicrosoftEnabled() bool`

GetMicrosoftEnabled returns the MicrosoftEnabled field if non-nil, zero value otherwise.

### GetMicrosoftEnabledOk

`func (o *WlanPortal) GetMicrosoftEnabledOk() (*bool, bool)`

GetMicrosoftEnabledOk returns a tuple with the MicrosoftEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftEnabled

`func (o *WlanPortal) SetMicrosoftEnabled(v bool)`

SetMicrosoftEnabled sets MicrosoftEnabled field to given value.

### HasMicrosoftEnabled

`func (o *WlanPortal) HasMicrosoftEnabled() bool`

HasMicrosoftEnabled returns a boolean if a field has been set.

### GetMicrosoftExpire

`func (o *WlanPortal) GetMicrosoftExpire() float64`

GetMicrosoftExpire returns the MicrosoftExpire field if non-nil, zero value otherwise.

### GetMicrosoftExpireOk

`func (o *WlanPortal) GetMicrosoftExpireOk() (*float64, bool)`

GetMicrosoftExpireOk returns a tuple with the MicrosoftExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftExpire

`func (o *WlanPortal) SetMicrosoftExpire(v float64)`

SetMicrosoftExpire sets MicrosoftExpire field to given value.

### HasMicrosoftExpire

`func (o *WlanPortal) HasMicrosoftExpire() bool`

HasMicrosoftExpire returns a boolean if a field has been set.

### SetMicrosoftExpireNil

`func (o *WlanPortal) SetMicrosoftExpireNil(b bool)`

 SetMicrosoftExpireNil sets the value for MicrosoftExpire to be an explicit nil

### UnsetMicrosoftExpire
`func (o *WlanPortal) UnsetMicrosoftExpire()`

UnsetMicrosoftExpire ensures that no value is present for MicrosoftExpire, not even an explicit nil
### GetPassphraseEnabled

`func (o *WlanPortal) GetPassphraseEnabled() bool`

GetPassphraseEnabled returns the PassphraseEnabled field if non-nil, zero value otherwise.

### GetPassphraseEnabledOk

`func (o *WlanPortal) GetPassphraseEnabledOk() (*bool, bool)`

GetPassphraseEnabledOk returns a tuple with the PassphraseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseEnabled

`func (o *WlanPortal) SetPassphraseEnabled(v bool)`

SetPassphraseEnabled sets PassphraseEnabled field to given value.

### HasPassphraseEnabled

`func (o *WlanPortal) HasPassphraseEnabled() bool`

HasPassphraseEnabled returns a boolean if a field has been set.

### GetPassphraseExpire

`func (o *WlanPortal) GetPassphraseExpire() float64`

GetPassphraseExpire returns the PassphraseExpire field if non-nil, zero value otherwise.

### GetPassphraseExpireOk

`func (o *WlanPortal) GetPassphraseExpireOk() (*float64, bool)`

GetPassphraseExpireOk returns a tuple with the PassphraseExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseExpire

`func (o *WlanPortal) SetPassphraseExpire(v float64)`

SetPassphraseExpire sets PassphraseExpire field to given value.

### HasPassphraseExpire

`func (o *WlanPortal) HasPassphraseExpire() bool`

HasPassphraseExpire returns a boolean if a field has been set.

### SetPassphraseExpireNil

`func (o *WlanPortal) SetPassphraseExpireNil(b bool)`

 SetPassphraseExpireNil sets the value for PassphraseExpire to be an explicit nil

### UnsetPassphraseExpire
`func (o *WlanPortal) UnsetPassphraseExpire()`

UnsetPassphraseExpire ensures that no value is present for PassphraseExpire, not even an explicit nil
### GetPassword

`func (o *WlanPortal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WlanPortal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WlanPortal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WlanPortal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WlanPortal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WlanPortal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPortalAllowedHostnames

`func (o *WlanPortal) GetPortalAllowedHostnames() string`

GetPortalAllowedHostnames returns the PortalAllowedHostnames field if non-nil, zero value otherwise.

### GetPortalAllowedHostnamesOk

`func (o *WlanPortal) GetPortalAllowedHostnamesOk() (*string, bool)`

GetPortalAllowedHostnamesOk returns a tuple with the PortalAllowedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAllowedHostnames

`func (o *WlanPortal) SetPortalAllowedHostnames(v string)`

SetPortalAllowedHostnames sets PortalAllowedHostnames field to given value.

### HasPortalAllowedHostnames

`func (o *WlanPortal) HasPortalAllowedHostnames() bool`

HasPortalAllowedHostnames returns a boolean if a field has been set.

### GetPortalAllowedSubnets

`func (o *WlanPortal) GetPortalAllowedSubnets() string`

GetPortalAllowedSubnets returns the PortalAllowedSubnets field if non-nil, zero value otherwise.

### GetPortalAllowedSubnetsOk

`func (o *WlanPortal) GetPortalAllowedSubnetsOk() (*string, bool)`

GetPortalAllowedSubnetsOk returns a tuple with the PortalAllowedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalAllowedSubnets

`func (o *WlanPortal) SetPortalAllowedSubnets(v string)`

SetPortalAllowedSubnets sets PortalAllowedSubnets field to given value.

### HasPortalAllowedSubnets

`func (o *WlanPortal) HasPortalAllowedSubnets() bool`

HasPortalAllowedSubnets returns a boolean if a field has been set.

### GetPortalApiSecret

`func (o *WlanPortal) GetPortalApiSecret() string`

GetPortalApiSecret returns the PortalApiSecret field if non-nil, zero value otherwise.

### GetPortalApiSecretOk

`func (o *WlanPortal) GetPortalApiSecretOk() (*string, bool)`

GetPortalApiSecretOk returns a tuple with the PortalApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalApiSecret

`func (o *WlanPortal) SetPortalApiSecret(v string)`

SetPortalApiSecret sets PortalApiSecret field to given value.

### HasPortalApiSecret

`func (o *WlanPortal) HasPortalApiSecret() bool`

HasPortalApiSecret returns a boolean if a field has been set.

### GetPortalDeniedHostnames

`func (o *WlanPortal) GetPortalDeniedHostnames() string`

GetPortalDeniedHostnames returns the PortalDeniedHostnames field if non-nil, zero value otherwise.

### GetPortalDeniedHostnamesOk

`func (o *WlanPortal) GetPortalDeniedHostnamesOk() (*string, bool)`

GetPortalDeniedHostnamesOk returns a tuple with the PortalDeniedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalDeniedHostnames

`func (o *WlanPortal) SetPortalDeniedHostnames(v string)`

SetPortalDeniedHostnames sets PortalDeniedHostnames field to given value.

### HasPortalDeniedHostnames

`func (o *WlanPortal) HasPortalDeniedHostnames() bool`

HasPortalDeniedHostnames returns a boolean if a field has been set.

### GetPortalImage

`func (o *WlanPortal) GetPortalImage() string`

GetPortalImage returns the PortalImage field if non-nil, zero value otherwise.

### GetPortalImageOk

`func (o *WlanPortal) GetPortalImageOk() (*string, bool)`

GetPortalImageOk returns a tuple with the PortalImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalImage

`func (o *WlanPortal) SetPortalImage(v string)`

SetPortalImage sets PortalImage field to given value.

### HasPortalImage

`func (o *WlanPortal) HasPortalImage() bool`

HasPortalImage returns a boolean if a field has been set.

### GetPortalSsoUrl

`func (o *WlanPortal) GetPortalSsoUrl() string`

GetPortalSsoUrl returns the PortalSsoUrl field if non-nil, zero value otherwise.

### GetPortalSsoUrlOk

`func (o *WlanPortal) GetPortalSsoUrlOk() (*string, bool)`

GetPortalSsoUrlOk returns a tuple with the PortalSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalSsoUrl

`func (o *WlanPortal) SetPortalSsoUrl(v string)`

SetPortalSsoUrl sets PortalSsoUrl field to given value.

### HasPortalSsoUrl

`func (o *WlanPortal) HasPortalSsoUrl() bool`

HasPortalSsoUrl returns a boolean if a field has been set.

### GetPredefinedSponsorsEnabled

`func (o *WlanPortal) GetPredefinedSponsorsEnabled() bool`

GetPredefinedSponsorsEnabled returns the PredefinedSponsorsEnabled field if non-nil, zero value otherwise.

### GetPredefinedSponsorsEnabledOk

`func (o *WlanPortal) GetPredefinedSponsorsEnabledOk() (*bool, bool)`

GetPredefinedSponsorsEnabledOk returns a tuple with the PredefinedSponsorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedSponsorsEnabled

`func (o *WlanPortal) SetPredefinedSponsorsEnabled(v bool)`

SetPredefinedSponsorsEnabled sets PredefinedSponsorsEnabled field to given value.

### HasPredefinedSponsorsEnabled

`func (o *WlanPortal) HasPredefinedSponsorsEnabled() bool`

HasPredefinedSponsorsEnabled returns a boolean if a field has been set.

### GetPrivacy

`func (o *WlanPortal) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *WlanPortal) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *WlanPortal) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *WlanPortal) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPuzzelPassword

`func (o *WlanPortal) GetPuzzelPassword() string`

GetPuzzelPassword returns the PuzzelPassword field if non-nil, zero value otherwise.

### GetPuzzelPasswordOk

`func (o *WlanPortal) GetPuzzelPasswordOk() (*string, bool)`

GetPuzzelPasswordOk returns a tuple with the PuzzelPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzelPassword

`func (o *WlanPortal) SetPuzzelPassword(v string)`

SetPuzzelPassword sets PuzzelPassword field to given value.

### HasPuzzelPassword

`func (o *WlanPortal) HasPuzzelPassword() bool`

HasPuzzelPassword returns a boolean if a field has been set.

### GetPuzzelServiceId

`func (o *WlanPortal) GetPuzzelServiceId() string`

GetPuzzelServiceId returns the PuzzelServiceId field if non-nil, zero value otherwise.

### GetPuzzelServiceIdOk

`func (o *WlanPortal) GetPuzzelServiceIdOk() (*string, bool)`

GetPuzzelServiceIdOk returns a tuple with the PuzzelServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzelServiceId

`func (o *WlanPortal) SetPuzzelServiceId(v string)`

SetPuzzelServiceId sets PuzzelServiceId field to given value.

### HasPuzzelServiceId

`func (o *WlanPortal) HasPuzzelServiceId() bool`

HasPuzzelServiceId returns a boolean if a field has been set.

### GetPuzzelUsername

`func (o *WlanPortal) GetPuzzelUsername() string`

GetPuzzelUsername returns the PuzzelUsername field if non-nil, zero value otherwise.

### GetPuzzelUsernameOk

`func (o *WlanPortal) GetPuzzelUsernameOk() (*string, bool)`

GetPuzzelUsernameOk returns a tuple with the PuzzelUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzelUsername

`func (o *WlanPortal) SetPuzzelUsername(v string)`

SetPuzzelUsername sets PuzzelUsername field to given value.

### HasPuzzelUsername

`func (o *WlanPortal) HasPuzzelUsername() bool`

HasPuzzelUsername returns a boolean if a field has been set.

### GetSmsMessageFormat

`func (o *WlanPortal) GetSmsMessageFormat() string`

GetSmsMessageFormat returns the SmsMessageFormat field if non-nil, zero value otherwise.

### GetSmsMessageFormatOk

`func (o *WlanPortal) GetSmsMessageFormatOk() (*string, bool)`

GetSmsMessageFormatOk returns a tuple with the SmsMessageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMessageFormat

`func (o *WlanPortal) SetSmsMessageFormat(v string)`

SetSmsMessageFormat sets SmsMessageFormat field to given value.

### HasSmsMessageFormat

`func (o *WlanPortal) HasSmsMessageFormat() bool`

HasSmsMessageFormat returns a boolean if a field has been set.

### GetSmsEnabled

`func (o *WlanPortal) GetSmsEnabled() bool`

GetSmsEnabled returns the SmsEnabled field if non-nil, zero value otherwise.

### GetSmsEnabledOk

`func (o *WlanPortal) GetSmsEnabledOk() (*bool, bool)`

GetSmsEnabledOk returns a tuple with the SmsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsEnabled

`func (o *WlanPortal) SetSmsEnabled(v bool)`

SetSmsEnabled sets SmsEnabled field to given value.

### HasSmsEnabled

`func (o *WlanPortal) HasSmsEnabled() bool`

HasSmsEnabled returns a boolean if a field has been set.

### GetSmsExpire

`func (o *WlanPortal) GetSmsExpire() float64`

GetSmsExpire returns the SmsExpire field if non-nil, zero value otherwise.

### GetSmsExpireOk

`func (o *WlanPortal) GetSmsExpireOk() (*float64, bool)`

GetSmsExpireOk returns a tuple with the SmsExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsExpire

`func (o *WlanPortal) SetSmsExpire(v float64)`

SetSmsExpire sets SmsExpire field to given value.

### HasSmsExpire

`func (o *WlanPortal) HasSmsExpire() bool`

HasSmsExpire returns a boolean if a field has been set.

### SetSmsExpireNil

`func (o *WlanPortal) SetSmsExpireNil(b bool)`

 SetSmsExpireNil sets the value for SmsExpire to be an explicit nil

### UnsetSmsExpire
`func (o *WlanPortal) UnsetSmsExpire()`

UnsetSmsExpire ensures that no value is present for SmsExpire, not even an explicit nil
### GetSmsProvider

`func (o *WlanPortal) GetSmsProvider() WlanPortalSmsProvider`

GetSmsProvider returns the SmsProvider field if non-nil, zero value otherwise.

### GetSmsProviderOk

`func (o *WlanPortal) GetSmsProviderOk() (*WlanPortalSmsProvider, bool)`

GetSmsProviderOk returns a tuple with the SmsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsProvider

`func (o *WlanPortal) SetSmsProvider(v WlanPortalSmsProvider)`

SetSmsProvider sets SmsProvider field to given value.

### HasSmsProvider

`func (o *WlanPortal) HasSmsProvider() bool`

HasSmsProvider returns a boolean if a field has been set.

### GetSponsorAutoApprove

`func (o *WlanPortal) GetSponsorAutoApprove() bool`

GetSponsorAutoApprove returns the SponsorAutoApprove field if non-nil, zero value otherwise.

### GetSponsorAutoApproveOk

`func (o *WlanPortal) GetSponsorAutoApproveOk() (*bool, bool)`

GetSponsorAutoApproveOk returns a tuple with the SponsorAutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAutoApprove

`func (o *WlanPortal) SetSponsorAutoApprove(v bool)`

SetSponsorAutoApprove sets SponsorAutoApprove field to given value.

### HasSponsorAutoApprove

`func (o *WlanPortal) HasSponsorAutoApprove() bool`

HasSponsorAutoApprove returns a boolean if a field has been set.

### GetSponsorEmailDomains

`func (o *WlanPortal) GetSponsorEmailDomains() []string`

GetSponsorEmailDomains returns the SponsorEmailDomains field if non-nil, zero value otherwise.

### GetSponsorEmailDomainsOk

`func (o *WlanPortal) GetSponsorEmailDomainsOk() (*[]string, bool)`

GetSponsorEmailDomainsOk returns a tuple with the SponsorEmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmailDomains

`func (o *WlanPortal) SetSponsorEmailDomains(v []string)`

SetSponsorEmailDomains sets SponsorEmailDomains field to given value.

### HasSponsorEmailDomains

`func (o *WlanPortal) HasSponsorEmailDomains() bool`

HasSponsorEmailDomains returns a boolean if a field has been set.

### GetSponsorEnabled

`func (o *WlanPortal) GetSponsorEnabled() bool`

GetSponsorEnabled returns the SponsorEnabled field if non-nil, zero value otherwise.

### GetSponsorEnabledOk

`func (o *WlanPortal) GetSponsorEnabledOk() (*bool, bool)`

GetSponsorEnabledOk returns a tuple with the SponsorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEnabled

`func (o *WlanPortal) SetSponsorEnabled(v bool)`

SetSponsorEnabled sets SponsorEnabled field to given value.

### HasSponsorEnabled

`func (o *WlanPortal) HasSponsorEnabled() bool`

HasSponsorEnabled returns a boolean if a field has been set.

### GetSponsorExpire

`func (o *WlanPortal) GetSponsorExpire() float64`

GetSponsorExpire returns the SponsorExpire field if non-nil, zero value otherwise.

### GetSponsorExpireOk

`func (o *WlanPortal) GetSponsorExpireOk() (*float64, bool)`

GetSponsorExpireOk returns a tuple with the SponsorExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorExpire

`func (o *WlanPortal) SetSponsorExpire(v float64)`

SetSponsorExpire sets SponsorExpire field to given value.

### HasSponsorExpire

`func (o *WlanPortal) HasSponsorExpire() bool`

HasSponsorExpire returns a boolean if a field has been set.

### SetSponsorExpireNil

`func (o *WlanPortal) SetSponsorExpireNil(b bool)`

 SetSponsorExpireNil sets the value for SponsorExpire to be an explicit nil

### UnsetSponsorExpire
`func (o *WlanPortal) UnsetSponsorExpire()`

UnsetSponsorExpire ensures that no value is present for SponsorExpire, not even an explicit nil
### GetSponsorLinkValidityDuration

`func (o *WlanPortal) GetSponsorLinkValidityDuration() int32`

GetSponsorLinkValidityDuration returns the SponsorLinkValidityDuration field if non-nil, zero value otherwise.

### GetSponsorLinkValidityDurationOk

`func (o *WlanPortal) GetSponsorLinkValidityDurationOk() (*int32, bool)`

GetSponsorLinkValidityDurationOk returns a tuple with the SponsorLinkValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorLinkValidityDuration

`func (o *WlanPortal) SetSponsorLinkValidityDuration(v int32)`

SetSponsorLinkValidityDuration sets SponsorLinkValidityDuration field to given value.

### HasSponsorLinkValidityDuration

`func (o *WlanPortal) HasSponsorLinkValidityDuration() bool`

HasSponsorLinkValidityDuration returns a boolean if a field has been set.

### GetSponsorNotifyAll

`func (o *WlanPortal) GetSponsorNotifyAll() bool`

GetSponsorNotifyAll returns the SponsorNotifyAll field if non-nil, zero value otherwise.

### GetSponsorNotifyAllOk

`func (o *WlanPortal) GetSponsorNotifyAllOk() (*bool, bool)`

GetSponsorNotifyAllOk returns a tuple with the SponsorNotifyAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNotifyAll

`func (o *WlanPortal) SetSponsorNotifyAll(v bool)`

SetSponsorNotifyAll sets SponsorNotifyAll field to given value.

### HasSponsorNotifyAll

`func (o *WlanPortal) HasSponsorNotifyAll() bool`

HasSponsorNotifyAll returns a boolean if a field has been set.

### GetSponsorStatusNotify

`func (o *WlanPortal) GetSponsorStatusNotify() bool`

GetSponsorStatusNotify returns the SponsorStatusNotify field if non-nil, zero value otherwise.

### GetSponsorStatusNotifyOk

`func (o *WlanPortal) GetSponsorStatusNotifyOk() (*bool, bool)`

GetSponsorStatusNotifyOk returns a tuple with the SponsorStatusNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusNotify

`func (o *WlanPortal) SetSponsorStatusNotify(v bool)`

SetSponsorStatusNotify sets SponsorStatusNotify field to given value.

### HasSponsorStatusNotify

`func (o *WlanPortal) HasSponsorStatusNotify() bool`

HasSponsorStatusNotify returns a boolean if a field has been set.

### GetSponsors

`func (o *WlanPortal) GetSponsors() map[string]string`

GetSponsors returns the Sponsors field if non-nil, zero value otherwise.

### GetSponsorsOk

`func (o *WlanPortal) GetSponsorsOk() (*map[string]string, bool)`

GetSponsorsOk returns a tuple with the Sponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsors

`func (o *WlanPortal) SetSponsors(v map[string]string)`

SetSponsors sets Sponsors field to given value.

### HasSponsors

`func (o *WlanPortal) HasSponsors() bool`

HasSponsors returns a boolean if a field has been set.

### GetSsoDefaultRole

`func (o *WlanPortal) GetSsoDefaultRole() string`

GetSsoDefaultRole returns the SsoDefaultRole field if non-nil, zero value otherwise.

### GetSsoDefaultRoleOk

`func (o *WlanPortal) GetSsoDefaultRoleOk() (*string, bool)`

GetSsoDefaultRoleOk returns a tuple with the SsoDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDefaultRole

`func (o *WlanPortal) SetSsoDefaultRole(v string)`

SetSsoDefaultRole sets SsoDefaultRole field to given value.

### HasSsoDefaultRole

`func (o *WlanPortal) HasSsoDefaultRole() bool`

HasSsoDefaultRole returns a boolean if a field has been set.

### GetSsoForcedRole

`func (o *WlanPortal) GetSsoForcedRole() string`

GetSsoForcedRole returns the SsoForcedRole field if non-nil, zero value otherwise.

### GetSsoForcedRoleOk

`func (o *WlanPortal) GetSsoForcedRoleOk() (*string, bool)`

GetSsoForcedRoleOk returns a tuple with the SsoForcedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForcedRole

`func (o *WlanPortal) SetSsoForcedRole(v string)`

SetSsoForcedRole sets SsoForcedRole field to given value.

### HasSsoForcedRole

`func (o *WlanPortal) HasSsoForcedRole() bool`

HasSsoForcedRole returns a boolean if a field has been set.

### GetSsoIdpCert

`func (o *WlanPortal) GetSsoIdpCert() string`

GetSsoIdpCert returns the SsoIdpCert field if non-nil, zero value otherwise.

### GetSsoIdpCertOk

`func (o *WlanPortal) GetSsoIdpCertOk() (*string, bool)`

GetSsoIdpCertOk returns a tuple with the SsoIdpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdpCert

`func (o *WlanPortal) SetSsoIdpCert(v string)`

SetSsoIdpCert sets SsoIdpCert field to given value.

### HasSsoIdpCert

`func (o *WlanPortal) HasSsoIdpCert() bool`

HasSsoIdpCert returns a boolean if a field has been set.

### GetSsoIdpSignAlgo

`func (o *WlanPortal) GetSsoIdpSignAlgo() string`

GetSsoIdpSignAlgo returns the SsoIdpSignAlgo field if non-nil, zero value otherwise.

### GetSsoIdpSignAlgoOk

`func (o *WlanPortal) GetSsoIdpSignAlgoOk() (*string, bool)`

GetSsoIdpSignAlgoOk returns a tuple with the SsoIdpSignAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdpSignAlgo

`func (o *WlanPortal) SetSsoIdpSignAlgo(v string)`

SetSsoIdpSignAlgo sets SsoIdpSignAlgo field to given value.

### HasSsoIdpSignAlgo

`func (o *WlanPortal) HasSsoIdpSignAlgo() bool`

HasSsoIdpSignAlgo returns a boolean if a field has been set.

### GetSsoIdpSsoUrl

`func (o *WlanPortal) GetSsoIdpSsoUrl() string`

GetSsoIdpSsoUrl returns the SsoIdpSsoUrl field if non-nil, zero value otherwise.

### GetSsoIdpSsoUrlOk

`func (o *WlanPortal) GetSsoIdpSsoUrlOk() (*string, bool)`

GetSsoIdpSsoUrlOk returns a tuple with the SsoIdpSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdpSsoUrl

`func (o *WlanPortal) SetSsoIdpSsoUrl(v string)`

SetSsoIdpSsoUrl sets SsoIdpSsoUrl field to given value.

### HasSsoIdpSsoUrl

`func (o *WlanPortal) HasSsoIdpSsoUrl() bool`

HasSsoIdpSsoUrl returns a boolean if a field has been set.

### GetSsoIssuer

`func (o *WlanPortal) GetSsoIssuer() string`

GetSsoIssuer returns the SsoIssuer field if non-nil, zero value otherwise.

### GetSsoIssuerOk

`func (o *WlanPortal) GetSsoIssuerOk() (*string, bool)`

GetSsoIssuerOk returns a tuple with the SsoIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIssuer

`func (o *WlanPortal) SetSsoIssuer(v string)`

SetSsoIssuer sets SsoIssuer field to given value.

### HasSsoIssuer

`func (o *WlanPortal) HasSsoIssuer() bool`

HasSsoIssuer returns a boolean if a field has been set.

### GetSsoNameidFormat

`func (o *WlanPortal) GetSsoNameidFormat() WlanPortalSsoNameidFormat`

GetSsoNameidFormat returns the SsoNameidFormat field if non-nil, zero value otherwise.

### GetSsoNameidFormatOk

`func (o *WlanPortal) GetSsoNameidFormatOk() (*WlanPortalSsoNameidFormat, bool)`

GetSsoNameidFormatOk returns a tuple with the SsoNameidFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoNameidFormat

`func (o *WlanPortal) SetSsoNameidFormat(v WlanPortalSsoNameidFormat)`

SetSsoNameidFormat sets SsoNameidFormat field to given value.

### HasSsoNameidFormat

`func (o *WlanPortal) HasSsoNameidFormat() bool`

HasSsoNameidFormat returns a boolean if a field has been set.

### GetTelstraClientId

`func (o *WlanPortal) GetTelstraClientId() string`

GetTelstraClientId returns the TelstraClientId field if non-nil, zero value otherwise.

### GetTelstraClientIdOk

`func (o *WlanPortal) GetTelstraClientIdOk() (*string, bool)`

GetTelstraClientIdOk returns a tuple with the TelstraClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelstraClientId

`func (o *WlanPortal) SetTelstraClientId(v string)`

SetTelstraClientId sets TelstraClientId field to given value.

### HasTelstraClientId

`func (o *WlanPortal) HasTelstraClientId() bool`

HasTelstraClientId returns a boolean if a field has been set.

### GetTelstraClientSecret

`func (o *WlanPortal) GetTelstraClientSecret() string`

GetTelstraClientSecret returns the TelstraClientSecret field if non-nil, zero value otherwise.

### GetTelstraClientSecretOk

`func (o *WlanPortal) GetTelstraClientSecretOk() (*string, bool)`

GetTelstraClientSecretOk returns a tuple with the TelstraClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelstraClientSecret

`func (o *WlanPortal) SetTelstraClientSecret(v string)`

SetTelstraClientSecret sets TelstraClientSecret field to given value.

### HasTelstraClientSecret

`func (o *WlanPortal) HasTelstraClientSecret() bool`

HasTelstraClientSecret returns a boolean if a field has been set.

### GetThumbnail

`func (o *WlanPortal) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *WlanPortal) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *WlanPortal) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *WlanPortal) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetTwilioAuthToken

`func (o *WlanPortal) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *WlanPortal) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *WlanPortal) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *WlanPortal) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### SetTwilioAuthTokenNil

`func (o *WlanPortal) SetTwilioAuthTokenNil(b bool)`

 SetTwilioAuthTokenNil sets the value for TwilioAuthToken to be an explicit nil

### UnsetTwilioAuthToken
`func (o *WlanPortal) UnsetTwilioAuthToken()`

UnsetTwilioAuthToken ensures that no value is present for TwilioAuthToken, not even an explicit nil
### GetTwilioPhoneNumber

`func (o *WlanPortal) GetTwilioPhoneNumber() string`

GetTwilioPhoneNumber returns the TwilioPhoneNumber field if non-nil, zero value otherwise.

### GetTwilioPhoneNumberOk

`func (o *WlanPortal) GetTwilioPhoneNumberOk() (*string, bool)`

GetTwilioPhoneNumberOk returns a tuple with the TwilioPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioPhoneNumber

`func (o *WlanPortal) SetTwilioPhoneNumber(v string)`

SetTwilioPhoneNumber sets TwilioPhoneNumber field to given value.

### HasTwilioPhoneNumber

`func (o *WlanPortal) HasTwilioPhoneNumber() bool`

HasTwilioPhoneNumber returns a boolean if a field has been set.

### SetTwilioPhoneNumberNil

`func (o *WlanPortal) SetTwilioPhoneNumberNil(b bool)`

 SetTwilioPhoneNumberNil sets the value for TwilioPhoneNumber to be an explicit nil

### UnsetTwilioPhoneNumber
`func (o *WlanPortal) UnsetTwilioPhoneNumber()`

UnsetTwilioPhoneNumber ensures that no value is present for TwilioPhoneNumber, not even an explicit nil
### GetTwilioSid

`func (o *WlanPortal) GetTwilioSid() string`

GetTwilioSid returns the TwilioSid field if non-nil, zero value otherwise.

### GetTwilioSidOk

`func (o *WlanPortal) GetTwilioSidOk() (*string, bool)`

GetTwilioSidOk returns a tuple with the TwilioSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioSid

`func (o *WlanPortal) SetTwilioSid(v string)`

SetTwilioSid sets TwilioSid field to given value.

### HasTwilioSid

`func (o *WlanPortal) HasTwilioSid() bool`

HasTwilioSid returns a boolean if a field has been set.

### SetTwilioSidNil

`func (o *WlanPortal) SetTwilioSidNil(b bool)`

 SetTwilioSidNil sets the value for TwilioSid to be an explicit nil

### UnsetTwilioSid
`func (o *WlanPortal) UnsetTwilioSid()`

UnsetTwilioSid ensures that no value is present for TwilioSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


