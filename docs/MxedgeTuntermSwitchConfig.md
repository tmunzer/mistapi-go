# MxedgeTuntermSwitchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortVlanId** | Pointer to **int32** |  | [optional] 
**VlanIds** | Pointer to [**[]MxedgeTuntermSwitchConfigVlanId**](MxedgeTuntermSwitchConfigVlanId.md) |  | [optional] 

## Methods

### NewMxedgeTuntermSwitchConfig

`func NewMxedgeTuntermSwitchConfig() *MxedgeTuntermSwitchConfig`

NewMxedgeTuntermSwitchConfig instantiates a new MxedgeTuntermSwitchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeTuntermSwitchConfigWithDefaults

`func NewMxedgeTuntermSwitchConfigWithDefaults() *MxedgeTuntermSwitchConfig`

NewMxedgeTuntermSwitchConfigWithDefaults instantiates a new MxedgeTuntermSwitchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortVlanId

`func (o *MxedgeTuntermSwitchConfig) GetPortVlanId() int32`

GetPortVlanId returns the PortVlanId field if non-nil, zero value otherwise.

### GetPortVlanIdOk

`func (o *MxedgeTuntermSwitchConfig) GetPortVlanIdOk() (*int32, bool)`

GetPortVlanIdOk returns a tuple with the PortVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanId

`func (o *MxedgeTuntermSwitchConfig) SetPortVlanId(v int32)`

SetPortVlanId sets PortVlanId field to given value.

### HasPortVlanId

`func (o *MxedgeTuntermSwitchConfig) HasPortVlanId() bool`

HasPortVlanId returns a boolean if a field has been set.

### GetVlanIds

`func (o *MxedgeTuntermSwitchConfig) GetVlanIds() []MxedgeTuntermSwitchConfigVlanId`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *MxedgeTuntermSwitchConfig) GetVlanIdsOk() (*[]MxedgeTuntermSwitchConfigVlanId, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *MxedgeTuntermSwitchConfig) SetVlanIds(v []MxedgeTuntermSwitchConfigVlanId)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *MxedgeTuntermSwitchConfig) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


