# Sdkinvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**ExpireTime** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** | name, will show up in mobile | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Quota** | Pointer to **int32** | number of time this invite can be used | [optional] 
**QuotaLimited** | Pointer to **bool** | whether quota limiting is enabled | [optional] [default to false]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSdkinvite

`func NewSdkinvite(name string, ) *Sdkinvite`

NewSdkinvite instantiates a new Sdkinvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkinviteWithDefaults

`func NewSdkinviteWithDefaults() *Sdkinvite`

NewSdkinviteWithDefaults instantiates a new Sdkinvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Sdkinvite) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Sdkinvite) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Sdkinvite) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Sdkinvite) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEnabled

`func (o *Sdkinvite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Sdkinvite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Sdkinvite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Sdkinvite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpireTime

`func (o *Sdkinvite) GetExpireTime() int32`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *Sdkinvite) GetExpireTimeOk() (*int32, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *Sdkinvite) SetExpireTime(v int32)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *Sdkinvite) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *Sdkinvite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sdkinvite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sdkinvite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sdkinvite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Sdkinvite) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Sdkinvite) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Sdkinvite) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Sdkinvite) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Sdkinvite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sdkinvite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sdkinvite) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *Sdkinvite) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Sdkinvite) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Sdkinvite) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Sdkinvite) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetQuota

`func (o *Sdkinvite) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Sdkinvite) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Sdkinvite) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Sdkinvite) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetQuotaLimited

`func (o *Sdkinvite) GetQuotaLimited() bool`

GetQuotaLimited returns the QuotaLimited field if non-nil, zero value otherwise.

### GetQuotaLimitedOk

`func (o *Sdkinvite) GetQuotaLimitedOk() (*bool, bool)`

GetQuotaLimitedOk returns a tuple with the QuotaLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimited

`func (o *Sdkinvite) SetQuotaLimited(v bool)`

SetQuotaLimited sets QuotaLimited field to given value.

### HasQuotaLimited

`func (o *Sdkinvite) HasQuotaLimited() bool`

HasQuotaLimited returns a boolean if a field has been set.

### GetSiteId

`func (o *Sdkinvite) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Sdkinvite) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Sdkinvite) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Sdkinvite) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


