# MxedgeTuntermIgmpSnoopingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Querier** | Pointer to [**MxedgeTuntermIgmpSnoopingQuerier**](MxedgeTuntermIgmpSnoopingQuerier.md) |  | [optional] 
**VlanIds** | Pointer to **[]int32** | the list of vlans on which tunterm performs IGMP snooping | [optional] 

## Methods

### NewMxedgeTuntermIgmpSnoopingConfig

`func NewMxedgeTuntermIgmpSnoopingConfig() *MxedgeTuntermIgmpSnoopingConfig`

NewMxedgeTuntermIgmpSnoopingConfig instantiates a new MxedgeTuntermIgmpSnoopingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeTuntermIgmpSnoopingConfigWithDefaults

`func NewMxedgeTuntermIgmpSnoopingConfigWithDefaults() *MxedgeTuntermIgmpSnoopingConfig`

NewMxedgeTuntermIgmpSnoopingConfigWithDefaults instantiates a new MxedgeTuntermIgmpSnoopingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxedgeTuntermIgmpSnoopingConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxedgeTuntermIgmpSnoopingConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetQuerier

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetQuerier() MxedgeTuntermIgmpSnoopingQuerier`

GetQuerier returns the Querier field if non-nil, zero value otherwise.

### GetQuerierOk

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetQuerierOk() (*MxedgeTuntermIgmpSnoopingQuerier, bool)`

GetQuerierOk returns a tuple with the Querier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerier

`func (o *MxedgeTuntermIgmpSnoopingConfig) SetQuerier(v MxedgeTuntermIgmpSnoopingQuerier)`

SetQuerier sets Querier field to given value.

### HasQuerier

`func (o *MxedgeTuntermIgmpSnoopingConfig) HasQuerier() bool`

HasQuerier returns a boolean if a field has been set.

### GetVlanIds

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *MxedgeTuntermIgmpSnoopingConfig) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *MxedgeTuntermIgmpSnoopingConfig) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *MxedgeTuntermIgmpSnoopingConfig) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


