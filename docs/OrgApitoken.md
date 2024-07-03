# OrgApitoken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **NullableString** | email of the token creator / null if creator is deleted | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Key** | Pointer to **string** |  | [optional] [readonly] 
**LastUsed** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | name of the token | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Privileges** | Pointer to [**[]PrivilegeOrg**](PrivilegeOrg.md) | list of privileges the token has on the orgs/sites | [optional] 
**SrcIps** | Pointer to **[]string** | to restrict where the API can be called from | [optional] 

## Methods

### NewOrgApitoken

`func NewOrgApitoken() *OrgApitoken`

NewOrgApitoken instantiates a new OrgApitoken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgApitokenWithDefaults

`func NewOrgApitokenWithDefaults() *OrgApitoken`

NewOrgApitokenWithDefaults instantiates a new OrgApitoken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OrgApitoken) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrgApitoken) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrgApitoken) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrgApitoken) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OrgApitoken) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OrgApitoken) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedTime

`func (o *OrgApitoken) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OrgApitoken) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OrgApitoken) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *OrgApitoken) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *OrgApitoken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgApitoken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgApitoken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgApitoken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *OrgApitoken) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OrgApitoken) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OrgApitoken) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OrgApitoken) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *OrgApitoken) GetLastUsed() int32`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *OrgApitoken) GetLastUsedOk() (*int32, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *OrgApitoken) SetLastUsed(v int32)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *OrgApitoken) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *OrgApitoken) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *OrgApitoken) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetName

`func (o *OrgApitoken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgApitoken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgApitoken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgApitoken) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrgApitoken) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrgApitoken) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrgId

`func (o *OrgApitoken) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgApitoken) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgApitoken) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgApitoken) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPrivileges

`func (o *OrgApitoken) GetPrivileges() []PrivilegeOrg`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *OrgApitoken) GetPrivilegesOk() (*[]PrivilegeOrg, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *OrgApitoken) SetPrivileges(v []PrivilegeOrg)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *OrgApitoken) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSrcIps

`func (o *OrgApitoken) GetSrcIps() []string`

GetSrcIps returns the SrcIps field if non-nil, zero value otherwise.

### GetSrcIpsOk

`func (o *OrgApitoken) GetSrcIpsOk() (*[]string, bool)`

GetSrcIpsOk returns a tuple with the SrcIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIps

`func (o *OrgApitoken) SetSrcIps(v []string)`

SetSrcIps sets SrcIps field to given value.

### HasSrcIps

`func (o *OrgApitoken) HasSrcIps() bool`

HasSrcIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


