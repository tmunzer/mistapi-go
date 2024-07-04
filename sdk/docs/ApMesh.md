# ApMesh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether mesh is enabled on this AP | [optional] [default to false]
**Group** | Pointer to **NullableInt32** | mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional | [optional] 
**Role** | Pointer to [**ApMeshRole**](ApMeshRole.md) |  | [optional] 

## Methods

### NewApMesh

`func NewApMesh() *ApMesh`

NewApMesh instantiates a new ApMesh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMeshWithDefaults

`func NewApMeshWithDefaults() *ApMesh`

NewApMeshWithDefaults instantiates a new ApMesh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApMesh) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApMesh) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApMesh) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApMesh) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroup

`func (o *ApMesh) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApMesh) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApMesh) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApMesh) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *ApMesh) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *ApMesh) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetRole

`func (o *ApMesh) GetRole() ApMeshRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApMesh) GetRoleOk() (*ApMeshRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApMesh) SetRole(v ApMeshRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApMesh) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


