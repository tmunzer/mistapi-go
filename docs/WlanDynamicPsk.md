# WlanDynamicPsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPsk** | Pointer to **string** | default PSK to use if cloud WLC is not available, 8-63 characters | [optional] 
**DefaultVlanId** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**ForceLookup** | Pointer to **bool** | when 11r is enabled, we&#39;ll try to use the cached PMK, this can be disabled &#x60;false&#x60; means auto | [optional] [default to false]
**Source** | Pointer to [**DynamicPskSource**](DynamicPskSource.md) |  | [optional] [default to DYNAMICPSKSOURCE_RADIUS]
**VlanIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewWlanDynamicPsk

`func NewWlanDynamicPsk() *WlanDynamicPsk`

NewWlanDynamicPsk instantiates a new WlanDynamicPsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanDynamicPskWithDefaults

`func NewWlanDynamicPskWithDefaults() *WlanDynamicPsk`

NewWlanDynamicPskWithDefaults instantiates a new WlanDynamicPsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPsk

`func (o *WlanDynamicPsk) GetDefaultPsk() string`

GetDefaultPsk returns the DefaultPsk field if non-nil, zero value otherwise.

### GetDefaultPskOk

`func (o *WlanDynamicPsk) GetDefaultPskOk() (*string, bool)`

GetDefaultPskOk returns a tuple with the DefaultPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPsk

`func (o *WlanDynamicPsk) SetDefaultPsk(v string)`

SetDefaultPsk sets DefaultPsk field to given value.

### HasDefaultPsk

`func (o *WlanDynamicPsk) HasDefaultPsk() bool`

HasDefaultPsk returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *WlanDynamicPsk) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *WlanDynamicPsk) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *WlanDynamicPsk) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *WlanDynamicPsk) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### SetDefaultVlanIdNil

`func (o *WlanDynamicPsk) SetDefaultVlanIdNil(b bool)`

 SetDefaultVlanIdNil sets the value for DefaultVlanId to be an explicit nil

### UnsetDefaultVlanId
`func (o *WlanDynamicPsk) UnsetDefaultVlanId()`

UnsetDefaultVlanId ensures that no value is present for DefaultVlanId, not even an explicit nil
### GetEnabled

`func (o *WlanDynamicPsk) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanDynamicPsk) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanDynamicPsk) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanDynamicPsk) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForceLookup

`func (o *WlanDynamicPsk) GetForceLookup() bool`

GetForceLookup returns the ForceLookup field if non-nil, zero value otherwise.

### GetForceLookupOk

`func (o *WlanDynamicPsk) GetForceLookupOk() (*bool, bool)`

GetForceLookupOk returns a tuple with the ForceLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceLookup

`func (o *WlanDynamicPsk) SetForceLookup(v bool)`

SetForceLookup sets ForceLookup field to given value.

### HasForceLookup

`func (o *WlanDynamicPsk) HasForceLookup() bool`

HasForceLookup returns a boolean if a field has been set.

### GetSource

`func (o *WlanDynamicPsk) GetSource() DynamicPskSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WlanDynamicPsk) GetSourceOk() (*DynamicPskSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WlanDynamicPsk) SetSource(v DynamicPskSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *WlanDynamicPsk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVlanIds

`func (o *WlanDynamicPsk) GetVlanIds() []*int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *WlanDynamicPsk) GetVlanIdsOk() (*[]*int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *WlanDynamicPsk) SetVlanIds(v []*int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *WlanDynamicPsk) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


