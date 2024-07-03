# MxedgeTuntermIgmpSnoopingQuerier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResponseTime** | Pointer to **int32** | querier’s query response interval, in tenths-of-seconds | [optional] 
**Mtu** | Pointer to **int32** | the MTU we use (needed when forming large IGMPv3 Reports) | [optional] 
**QueryInterval** | Pointer to **int32** | querier’s query interval, in seconds | [optional] 
**Robustness** | Pointer to **int32** | querier’s robustness | [optional] 
**Version** | Pointer to **int32** | querier’s maximum protocol version | [optional] 

## Methods

### NewMxedgeTuntermIgmpSnoopingQuerier

`func NewMxedgeTuntermIgmpSnoopingQuerier() *MxedgeTuntermIgmpSnoopingQuerier`

NewMxedgeTuntermIgmpSnoopingQuerier instantiates a new MxedgeTuntermIgmpSnoopingQuerier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeTuntermIgmpSnoopingQuerierWithDefaults

`func NewMxedgeTuntermIgmpSnoopingQuerierWithDefaults() *MxedgeTuntermIgmpSnoopingQuerier`

NewMxedgeTuntermIgmpSnoopingQuerierWithDefaults instantiates a new MxedgeTuntermIgmpSnoopingQuerier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResponseTime

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *MxedgeTuntermIgmpSnoopingQuerier) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *MxedgeTuntermIgmpSnoopingQuerier) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetMtu

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *MxedgeTuntermIgmpSnoopingQuerier) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *MxedgeTuntermIgmpSnoopingQuerier) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetQueryInterval

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetQueryInterval() int32`

GetQueryInterval returns the QueryInterval field if non-nil, zero value otherwise.

### GetQueryIntervalOk

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetQueryIntervalOk() (*int32, bool)`

GetQueryIntervalOk returns a tuple with the QueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInterval

`func (o *MxedgeTuntermIgmpSnoopingQuerier) SetQueryInterval(v int32)`

SetQueryInterval sets QueryInterval field to given value.

### HasQueryInterval

`func (o *MxedgeTuntermIgmpSnoopingQuerier) HasQueryInterval() bool`

HasQueryInterval returns a boolean if a field has been set.

### GetRobustness

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetRobustness() int32`

GetRobustness returns the Robustness field if non-nil, zero value otherwise.

### GetRobustnessOk

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetRobustnessOk() (*int32, bool)`

GetRobustnessOk returns a tuple with the Robustness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobustness

`func (o *MxedgeTuntermIgmpSnoopingQuerier) SetRobustness(v int32)`

SetRobustness sets Robustness field to given value.

### HasRobustness

`func (o *MxedgeTuntermIgmpSnoopingQuerier) HasRobustness() bool`

HasRobustness returns a boolean if a field has been set.

### GetVersion

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MxedgeTuntermIgmpSnoopingQuerier) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MxedgeTuntermIgmpSnoopingQuerier) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MxedgeTuntermIgmpSnoopingQuerier) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


