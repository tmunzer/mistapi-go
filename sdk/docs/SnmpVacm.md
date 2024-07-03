# SnmpVacm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**[]SnmpVacmAccessItem**](SnmpVacmAccessItem.md) |  | [optional] 
**SecurityToGroup** | Pointer to [**SnmpVacmSecurityToGroup**](SnmpVacmSecurityToGroup.md) |  | [optional] 

## Methods

### NewSnmpVacm

`func NewSnmpVacm() *SnmpVacm`

NewSnmpVacm instantiates a new SnmpVacm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpVacmWithDefaults

`func NewSnmpVacmWithDefaults() *SnmpVacm`

NewSnmpVacmWithDefaults instantiates a new SnmpVacm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SnmpVacm) GetAccess() []SnmpVacmAccessItem`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SnmpVacm) GetAccessOk() (*[]SnmpVacmAccessItem, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SnmpVacm) SetAccess(v []SnmpVacmAccessItem)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SnmpVacm) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetSecurityToGroup

`func (o *SnmpVacm) GetSecurityToGroup() SnmpVacmSecurityToGroup`

GetSecurityToGroup returns the SecurityToGroup field if non-nil, zero value otherwise.

### GetSecurityToGroupOk

`func (o *SnmpVacm) GetSecurityToGroupOk() (*SnmpVacmSecurityToGroup, bool)`

GetSecurityToGroupOk returns a tuple with the SecurityToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToGroup

`func (o *SnmpVacm) SetSecurityToGroup(v SnmpVacmSecurityToGroup)`

SetSecurityToGroup sets SecurityToGroup field to given value.

### HasSecurityToGroup

`func (o *SnmpVacm) HasSecurityToGroup() bool`

HasSecurityToGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


