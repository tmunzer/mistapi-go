# SsoRoleMsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Privileges** | [**[]PrivilegeMsp**](PrivilegeMsp.md) |  | 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSsoRoleMsp

`func NewSsoRoleMsp(name string, privileges []PrivilegeMsp, ) *SsoRoleMsp`

NewSsoRoleMsp instantiates a new SsoRoleMsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoRoleMspWithDefaults

`func NewSsoRoleMspWithDefaults() *SsoRoleMsp`

NewSsoRoleMspWithDefaults instantiates a new SsoRoleMsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *SsoRoleMsp) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SsoRoleMsp) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SsoRoleMsp) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SsoRoleMsp) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *SsoRoleMsp) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *SsoRoleMsp) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *SsoRoleMsp) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *SsoRoleMsp) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *SsoRoleMsp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsoRoleMsp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsoRoleMsp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SsoRoleMsp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *SsoRoleMsp) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *SsoRoleMsp) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *SsoRoleMsp) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *SsoRoleMsp) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *SsoRoleMsp) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *SsoRoleMsp) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *SsoRoleMsp) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *SsoRoleMsp) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetName

`func (o *SsoRoleMsp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsoRoleMsp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsoRoleMsp) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *SsoRoleMsp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SsoRoleMsp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SsoRoleMsp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SsoRoleMsp) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPrivileges

`func (o *SsoRoleMsp) GetPrivileges() []PrivilegeMsp`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *SsoRoleMsp) GetPrivilegesOk() (*[]PrivilegeMsp, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *SsoRoleMsp) SetPrivileges(v []PrivilegeMsp)`

SetPrivileges sets Privileges field to given value.


### GetSiteId

`func (o *SsoRoleMsp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SsoRoleMsp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SsoRoleMsp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SsoRoleMsp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


