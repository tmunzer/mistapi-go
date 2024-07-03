# ApPortConfigDynamicVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVlanId** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vlans** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApPortConfigDynamicVlan

`func NewApPortConfigDynamicVlan() *ApPortConfigDynamicVlan`

NewApPortConfigDynamicVlan instantiates a new ApPortConfigDynamicVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPortConfigDynamicVlanWithDefaults

`func NewApPortConfigDynamicVlanWithDefaults() *ApPortConfigDynamicVlan`

NewApPortConfigDynamicVlanWithDefaults instantiates a new ApPortConfigDynamicVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVlanId

`func (o *ApPortConfigDynamicVlan) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *ApPortConfigDynamicVlan) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *ApPortConfigDynamicVlan) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *ApPortConfigDynamicVlan) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetEnabled

`func (o *ApPortConfigDynamicVlan) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApPortConfigDynamicVlan) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApPortConfigDynamicVlan) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApPortConfigDynamicVlan) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ApPortConfigDynamicVlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApPortConfigDynamicVlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApPortConfigDynamicVlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApPortConfigDynamicVlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlans

`func (o *ApPortConfigDynamicVlan) GetVlans() map[string]string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *ApPortConfigDynamicVlan) GetVlansOk() (*map[string]string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *ApPortConfigDynamicVlan) SetVlans(v map[string]string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *ApPortConfigDynamicVlan) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


