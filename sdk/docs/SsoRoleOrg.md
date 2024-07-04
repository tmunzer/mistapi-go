# SsoRoleOrg

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
**Privileges** | [**[]PrivilegeOrg**](PrivilegeOrg.md) |  | 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSsoRoleOrg

`func NewSsoRoleOrg(name string, privileges []PrivilegeOrg, ) *SsoRoleOrg`

NewSsoRoleOrg instantiates a new SsoRoleOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoRoleOrgWithDefaults

`func NewSsoRoleOrgWithDefaults() *SsoRoleOrg`

NewSsoRoleOrgWithDefaults instantiates a new SsoRoleOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *SsoRoleOrg) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SsoRoleOrg) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SsoRoleOrg) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SsoRoleOrg) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *SsoRoleOrg) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *SsoRoleOrg) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *SsoRoleOrg) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *SsoRoleOrg) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *SsoRoleOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsoRoleOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsoRoleOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SsoRoleOrg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *SsoRoleOrg) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *SsoRoleOrg) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *SsoRoleOrg) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *SsoRoleOrg) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *SsoRoleOrg) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *SsoRoleOrg) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *SsoRoleOrg) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *SsoRoleOrg) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetName

`func (o *SsoRoleOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsoRoleOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsoRoleOrg) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *SsoRoleOrg) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SsoRoleOrg) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SsoRoleOrg) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SsoRoleOrg) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPrivileges

`func (o *SsoRoleOrg) GetPrivileges() []PrivilegeOrg`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *SsoRoleOrg) GetPrivilegesOk() (*[]PrivilegeOrg, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *SsoRoleOrg) SetPrivileges(v []PrivilegeOrg)`

SetPrivileges sets Privileges field to given value.


### GetSiteId

`func (o *SsoRoleOrg) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SsoRoleOrg) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SsoRoleOrg) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SsoRoleOrg) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


