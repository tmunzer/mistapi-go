# PskPortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**PskPortalAuth**](PskPortalAuth.md) |  | [optional] [default to PSKPORTALAUTH_SSO]
**BgImageUrl** | Pointer to **string** |  | [optional] 
**CleanupPsk** | Pointer to **bool** | used to cleanup exited psk when portal delete or ssid changed | [optional] [default to false]
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**ExpireTime** | Pointer to **int32** | unit min | [optional] 
**ExpiryNotificationTime** | Pointer to **int32** | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire | [optional] 
**HidePsksCreatedByOtherAdmins** | Pointer to **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;admin&#x60; | [optional] [default to false]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**MaxUsage** | Pointer to **int32** | &#x60;max_usage&#x60;&#x3D;&#x3D;&#x60;0&#x60; means unlimited | [optional] [default to 0]
**ModifiedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**NotificationRenewUrl** | Pointer to **string** | optional, will include the link in the notification email the customer can either provide their own url or use the one generate from mist, or do a url shorterner against either | [optional] 
**NotifyExpiry** | Pointer to **bool** | If set to true, reminder notification will be sent when psk is about to expire | [optional] 
**NotifyOnCreateOrEdit** | Pointer to **bool** |  | [optional] [default to false]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PassphraseRules** | Pointer to [**PskPortalPassphraseRules**](PskPortalPassphraseRules.md) |  | [optional] 
**RequiredFields** | Pointer to **[]string** | what information to ask for (email is required by default) | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Ssid** | **string** | intended SSID | 
**Sso** | Pointer to [**PskPortalSso**](PskPortalSso.md) |  | [optional] 
**TemplateUrl** | Pointer to **string** | UI customization | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PskPortalType**](PskPortalType.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPskPortal

`func NewPskPortal(name string, ssid string, ) *PskPortal`

NewPskPortal instantiates a new PskPortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskPortalWithDefaults

`func NewPskPortalWithDefaults() *PskPortal`

NewPskPortalWithDefaults instantiates a new PskPortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *PskPortal) GetAuth() PskPortalAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *PskPortal) GetAuthOk() (*PskPortalAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *PskPortal) SetAuth(v PskPortalAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *PskPortal) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetBgImageUrl

`func (o *PskPortal) GetBgImageUrl() string`

GetBgImageUrl returns the BgImageUrl field if non-nil, zero value otherwise.

### GetBgImageUrlOk

`func (o *PskPortal) GetBgImageUrlOk() (*string, bool)`

GetBgImageUrlOk returns a tuple with the BgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImageUrl

`func (o *PskPortal) SetBgImageUrl(v string)`

SetBgImageUrl sets BgImageUrl field to given value.

### HasBgImageUrl

`func (o *PskPortal) HasBgImageUrl() bool`

HasBgImageUrl returns a boolean if a field has been set.

### GetCleanupPsk

`func (o *PskPortal) GetCleanupPsk() bool`

GetCleanupPsk returns the CleanupPsk field if non-nil, zero value otherwise.

### GetCleanupPskOk

`func (o *PskPortal) GetCleanupPskOk() (*bool, bool)`

GetCleanupPskOk returns a tuple with the CleanupPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPsk

`func (o *PskPortal) SetCleanupPsk(v bool)`

SetCleanupPsk sets CleanupPsk field to given value.

### HasCleanupPsk

`func (o *PskPortal) HasCleanupPsk() bool`

HasCleanupPsk returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PskPortal) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PskPortal) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PskPortal) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PskPortal) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *PskPortal) GetExpireTime() int32`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *PskPortal) GetExpireTimeOk() (*int32, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *PskPortal) SetExpireTime(v int32)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *PskPortal) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetExpiryNotificationTime

`func (o *PskPortal) GetExpiryNotificationTime() int32`

GetExpiryNotificationTime returns the ExpiryNotificationTime field if non-nil, zero value otherwise.

### GetExpiryNotificationTimeOk

`func (o *PskPortal) GetExpiryNotificationTimeOk() (*int32, bool)`

GetExpiryNotificationTimeOk returns a tuple with the ExpiryNotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryNotificationTime

`func (o *PskPortal) SetExpiryNotificationTime(v int32)`

SetExpiryNotificationTime sets ExpiryNotificationTime field to given value.

### HasExpiryNotificationTime

`func (o *PskPortal) HasExpiryNotificationTime() bool`

HasExpiryNotificationTime returns a boolean if a field has been set.

### GetHidePsksCreatedByOtherAdmins

`func (o *PskPortal) GetHidePsksCreatedByOtherAdmins() bool`

GetHidePsksCreatedByOtherAdmins returns the HidePsksCreatedByOtherAdmins field if non-nil, zero value otherwise.

### GetHidePsksCreatedByOtherAdminsOk

`func (o *PskPortal) GetHidePsksCreatedByOtherAdminsOk() (*bool, bool)`

GetHidePsksCreatedByOtherAdminsOk returns a tuple with the HidePsksCreatedByOtherAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePsksCreatedByOtherAdmins

`func (o *PskPortal) SetHidePsksCreatedByOtherAdmins(v bool)`

SetHidePsksCreatedByOtherAdmins sets HidePsksCreatedByOtherAdmins field to given value.

### HasHidePsksCreatedByOtherAdmins

`func (o *PskPortal) HasHidePsksCreatedByOtherAdmins() bool`

HasHidePsksCreatedByOtherAdmins returns a boolean if a field has been set.

### GetId

`func (o *PskPortal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PskPortal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PskPortal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PskPortal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxUsage

`func (o *PskPortal) GetMaxUsage() int32`

GetMaxUsage returns the MaxUsage field if non-nil, zero value otherwise.

### GetMaxUsageOk

`func (o *PskPortal) GetMaxUsageOk() (*int32, bool)`

GetMaxUsageOk returns a tuple with the MaxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsage

`func (o *PskPortal) SetMaxUsage(v int32)`

SetMaxUsage sets MaxUsage field to given value.

### HasMaxUsage

`func (o *PskPortal) HasMaxUsage() bool`

HasMaxUsage returns a boolean if a field has been set.

### GetModifiedTime

`func (o *PskPortal) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *PskPortal) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *PskPortal) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *PskPortal) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *PskPortal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PskPortal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PskPortal) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationRenewUrl

`func (o *PskPortal) GetNotificationRenewUrl() string`

GetNotificationRenewUrl returns the NotificationRenewUrl field if non-nil, zero value otherwise.

### GetNotificationRenewUrlOk

`func (o *PskPortal) GetNotificationRenewUrlOk() (*string, bool)`

GetNotificationRenewUrlOk returns a tuple with the NotificationRenewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRenewUrl

`func (o *PskPortal) SetNotificationRenewUrl(v string)`

SetNotificationRenewUrl sets NotificationRenewUrl field to given value.

### HasNotificationRenewUrl

`func (o *PskPortal) HasNotificationRenewUrl() bool`

HasNotificationRenewUrl returns a boolean if a field has been set.

### GetNotifyExpiry

`func (o *PskPortal) GetNotifyExpiry() bool`

GetNotifyExpiry returns the NotifyExpiry field if non-nil, zero value otherwise.

### GetNotifyExpiryOk

`func (o *PskPortal) GetNotifyExpiryOk() (*bool, bool)`

GetNotifyExpiryOk returns a tuple with the NotifyExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyExpiry

`func (o *PskPortal) SetNotifyExpiry(v bool)`

SetNotifyExpiry sets NotifyExpiry field to given value.

### HasNotifyExpiry

`func (o *PskPortal) HasNotifyExpiry() bool`

HasNotifyExpiry returns a boolean if a field has been set.

### GetNotifyOnCreateOrEdit

`func (o *PskPortal) GetNotifyOnCreateOrEdit() bool`

GetNotifyOnCreateOrEdit returns the NotifyOnCreateOrEdit field if non-nil, zero value otherwise.

### GetNotifyOnCreateOrEditOk

`func (o *PskPortal) GetNotifyOnCreateOrEditOk() (*bool, bool)`

GetNotifyOnCreateOrEditOk returns a tuple with the NotifyOnCreateOrEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnCreateOrEdit

`func (o *PskPortal) SetNotifyOnCreateOrEdit(v bool)`

SetNotifyOnCreateOrEdit sets NotifyOnCreateOrEdit field to given value.

### HasNotifyOnCreateOrEdit

`func (o *PskPortal) HasNotifyOnCreateOrEdit() bool`

HasNotifyOnCreateOrEdit returns a boolean if a field has been set.

### GetOrgId

`func (o *PskPortal) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PskPortal) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PskPortal) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PskPortal) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPassphraseRules

`func (o *PskPortal) GetPassphraseRules() PskPortalPassphraseRules`

GetPassphraseRules returns the PassphraseRules field if non-nil, zero value otherwise.

### GetPassphraseRulesOk

`func (o *PskPortal) GetPassphraseRulesOk() (*PskPortalPassphraseRules, bool)`

GetPassphraseRulesOk returns a tuple with the PassphraseRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseRules

`func (o *PskPortal) SetPassphraseRules(v PskPortalPassphraseRules)`

SetPassphraseRules sets PassphraseRules field to given value.

### HasPassphraseRules

`func (o *PskPortal) HasPassphraseRules() bool`

HasPassphraseRules returns a boolean if a field has been set.

### GetRequiredFields

`func (o *PskPortal) GetRequiredFields() []string`

GetRequiredFields returns the RequiredFields field if non-nil, zero value otherwise.

### GetRequiredFieldsOk

`func (o *PskPortal) GetRequiredFieldsOk() (*[]string, bool)`

GetRequiredFieldsOk returns a tuple with the RequiredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFields

`func (o *PskPortal) SetRequiredFields(v []string)`

SetRequiredFields sets RequiredFields field to given value.

### HasRequiredFields

`func (o *PskPortal) HasRequiredFields() bool`

HasRequiredFields returns a boolean if a field has been set.

### GetRole

`func (o *PskPortal) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PskPortal) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PskPortal) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PskPortal) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSsid

`func (o *PskPortal) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PskPortal) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PskPortal) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetSso

`func (o *PskPortal) GetSso() PskPortalSso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *PskPortal) GetSsoOk() (*PskPortalSso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *PskPortal) SetSso(v PskPortalSso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *PskPortal) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetTemplateUrl

`func (o *PskPortal) GetTemplateUrl() string`

GetTemplateUrl returns the TemplateUrl field if non-nil, zero value otherwise.

### GetTemplateUrlOk

`func (o *PskPortal) GetTemplateUrlOk() (*string, bool)`

GetTemplateUrlOk returns a tuple with the TemplateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrl

`func (o *PskPortal) SetTemplateUrl(v string)`

SetTemplateUrl sets TemplateUrl field to given value.

### HasTemplateUrl

`func (o *PskPortal) HasTemplateUrl() bool`

HasTemplateUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *PskPortal) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PskPortal) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PskPortal) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PskPortal) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetType

`func (o *PskPortal) GetType() PskPortalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PskPortal) GetTypeOk() (*PskPortalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PskPortal) SetType(v PskPortalType)`

SetType sets Type field to given value.

### HasType

`func (o *PskPortal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanId

`func (o *PskPortal) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PskPortal) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PskPortal) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *PskPortal) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


