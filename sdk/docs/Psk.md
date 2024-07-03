# Psk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSsoId** | Pointer to **string** | sso id for psk created from psk portal | [optional] [readonly] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Email** | Pointer to **string** | email to send psk expiring notifications to | [optional] 
**ExpireTime** | Pointer to **NullableInt32** | Expire time for this PSK key (epoch time in seconds). Default &#x60;null&#x60; (as no expiration) | [optional] 
**ExpiryNotificationTime** | Pointer to **int32** | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Mac** | Pointer to **string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;single&#x60;, the mac that this PSK ties to, empty if &#x60;auto-binding&#x60; | [optional] 
**Macs** | Pointer to **[]string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;macs&#x60;, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000 | [optional] 
**MaxUsage** | Pointer to **int32** | For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited) | [optional] [default to 0]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Note** | Pointer to **string** |  | [optional] 
**NotifyExpiry** | Pointer to **bool** | If set to true, reminder notification will be sent when psk is about to expire | [optional] [default to false]
**NotifyOnCreateOrEdit** | Pointer to **bool** | If set to true, notification will be sent when psk is created or edited | [optional] 
**OldPassphrase** | Pointer to **string** | previous passphrase of the PSK if it has been rotated | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Passphrase** | **string** | passphrase of the PSK (8-63 character or 64 in hex) | 
**Role** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssid** | **string** | SSID this PSK should be applicable to | 
**Usage** | Pointer to [**PskUsage**](PskUsage.md) |  | [optional] [default to PSKUSAGE_MULTI]
**VlanId** | Pointer to **int32** | VLAN for this PSK key | [optional] 

## Methods

### NewPsk

`func NewPsk(name string, passphrase string, ssid string, ) *Psk`

NewPsk instantiates a new Psk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskWithDefaults

`func NewPskWithDefaults() *Psk`

NewPskWithDefaults instantiates a new Psk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSsoId

`func (o *Psk) GetAdminSsoId() string`

GetAdminSsoId returns the AdminSsoId field if non-nil, zero value otherwise.

### GetAdminSsoIdOk

`func (o *Psk) GetAdminSsoIdOk() (*string, bool)`

GetAdminSsoIdOk returns a tuple with the AdminSsoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSsoId

`func (o *Psk) SetAdminSsoId(v string)`

SetAdminSsoId sets AdminSsoId field to given value.

### HasAdminSsoId

`func (o *Psk) HasAdminSsoId() bool`

HasAdminSsoId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Psk) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Psk) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Psk) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Psk) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEmail

`func (o *Psk) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Psk) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Psk) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Psk) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpireTime

`func (o *Psk) GetExpireTime() int32`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *Psk) GetExpireTimeOk() (*int32, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *Psk) SetExpireTime(v int32)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *Psk) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### SetExpireTimeNil

`func (o *Psk) SetExpireTimeNil(b bool)`

 SetExpireTimeNil sets the value for ExpireTime to be an explicit nil

### UnsetExpireTime
`func (o *Psk) UnsetExpireTime()`

UnsetExpireTime ensures that no value is present for ExpireTime, not even an explicit nil
### GetExpiryNotificationTime

`func (o *Psk) GetExpiryNotificationTime() int32`

GetExpiryNotificationTime returns the ExpiryNotificationTime field if non-nil, zero value otherwise.

### GetExpiryNotificationTimeOk

`func (o *Psk) GetExpiryNotificationTimeOk() (*int32, bool)`

GetExpiryNotificationTimeOk returns a tuple with the ExpiryNotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryNotificationTime

`func (o *Psk) SetExpiryNotificationTime(v int32)`

SetExpiryNotificationTime sets ExpiryNotificationTime field to given value.

### HasExpiryNotificationTime

`func (o *Psk) HasExpiryNotificationTime() bool`

HasExpiryNotificationTime returns a boolean if a field has been set.

### GetId

`func (o *Psk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Psk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Psk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Psk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *Psk) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Psk) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Psk) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Psk) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMacs

`func (o *Psk) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *Psk) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *Psk) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *Psk) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetMaxUsage

`func (o *Psk) GetMaxUsage() int32`

GetMaxUsage returns the MaxUsage field if non-nil, zero value otherwise.

### GetMaxUsageOk

`func (o *Psk) GetMaxUsageOk() (*int32, bool)`

GetMaxUsageOk returns a tuple with the MaxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsage

`func (o *Psk) SetMaxUsage(v int32)`

SetMaxUsage sets MaxUsage field to given value.

### HasMaxUsage

`func (o *Psk) HasMaxUsage() bool`

HasMaxUsage returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Psk) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Psk) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Psk) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Psk) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Psk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Psk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Psk) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *Psk) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Psk) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Psk) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Psk) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNotifyExpiry

`func (o *Psk) GetNotifyExpiry() bool`

GetNotifyExpiry returns the NotifyExpiry field if non-nil, zero value otherwise.

### GetNotifyExpiryOk

`func (o *Psk) GetNotifyExpiryOk() (*bool, bool)`

GetNotifyExpiryOk returns a tuple with the NotifyExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyExpiry

`func (o *Psk) SetNotifyExpiry(v bool)`

SetNotifyExpiry sets NotifyExpiry field to given value.

### HasNotifyExpiry

`func (o *Psk) HasNotifyExpiry() bool`

HasNotifyExpiry returns a boolean if a field has been set.

### GetNotifyOnCreateOrEdit

`func (o *Psk) GetNotifyOnCreateOrEdit() bool`

GetNotifyOnCreateOrEdit returns the NotifyOnCreateOrEdit field if non-nil, zero value otherwise.

### GetNotifyOnCreateOrEditOk

`func (o *Psk) GetNotifyOnCreateOrEditOk() (*bool, bool)`

GetNotifyOnCreateOrEditOk returns a tuple with the NotifyOnCreateOrEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnCreateOrEdit

`func (o *Psk) SetNotifyOnCreateOrEdit(v bool)`

SetNotifyOnCreateOrEdit sets NotifyOnCreateOrEdit field to given value.

### HasNotifyOnCreateOrEdit

`func (o *Psk) HasNotifyOnCreateOrEdit() bool`

HasNotifyOnCreateOrEdit returns a boolean if a field has been set.

### GetOldPassphrase

`func (o *Psk) GetOldPassphrase() string`

GetOldPassphrase returns the OldPassphrase field if non-nil, zero value otherwise.

### GetOldPassphraseOk

`func (o *Psk) GetOldPassphraseOk() (*string, bool)`

GetOldPassphraseOk returns a tuple with the OldPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassphrase

`func (o *Psk) SetOldPassphrase(v string)`

SetOldPassphrase sets OldPassphrase field to given value.

### HasOldPassphrase

`func (o *Psk) HasOldPassphrase() bool`

HasOldPassphrase returns a boolean if a field has been set.

### GetOrgId

`func (o *Psk) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Psk) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Psk) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Psk) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPassphrase

`func (o *Psk) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *Psk) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *Psk) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetRole

`func (o *Psk) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Psk) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Psk) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Psk) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSiteId

`func (o *Psk) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Psk) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Psk) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Psk) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsid

`func (o *Psk) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *Psk) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *Psk) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetUsage

`func (o *Psk) GetUsage() PskUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Psk) GetUsageOk() (*PskUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Psk) SetUsage(v PskUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Psk) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetVlanId

`func (o *Psk) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Psk) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Psk) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Psk) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


