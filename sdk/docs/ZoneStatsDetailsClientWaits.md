# ZoneStatsDetailsClientWaits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg** | **int32** | average wait time in seconds | 
**Max** | **int32** | longest wait time in seconds | 
**Min** | **int32** | shortest wait time in seconds | 
**P95** | **int32** | 95th percentile of all the wait time(s) | 

## Methods

### NewZoneStatsDetailsClientWaits

`func NewZoneStatsDetailsClientWaits(avg int32, max int32, min int32, p95 int32, ) *ZoneStatsDetailsClientWaits`

NewZoneStatsDetailsClientWaits instantiates a new ZoneStatsDetailsClientWaits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStatsDetailsClientWaitsWithDefaults

`func NewZoneStatsDetailsClientWaitsWithDefaults() *ZoneStatsDetailsClientWaits`

NewZoneStatsDetailsClientWaitsWithDefaults instantiates a new ZoneStatsDetailsClientWaits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg

`func (o *ZoneStatsDetailsClientWaits) GetAvg() int32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *ZoneStatsDetailsClientWaits) GetAvgOk() (*int32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *ZoneStatsDetailsClientWaits) SetAvg(v int32)`

SetAvg sets Avg field to given value.


### GetMax

`func (o *ZoneStatsDetailsClientWaits) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ZoneStatsDetailsClientWaits) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ZoneStatsDetailsClientWaits) SetMax(v int32)`

SetMax sets Max field to given value.


### GetMin

`func (o *ZoneStatsDetailsClientWaits) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ZoneStatsDetailsClientWaits) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ZoneStatsDetailsClientWaits) SetMin(v int32)`

SetMin sets Min field to given value.


### GetP95

`func (o *ZoneStatsDetailsClientWaits) GetP95() int32`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *ZoneStatsDetailsClientWaits) GetP95Ok() (*int32, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *ZoneStatsDetailsClientWaits) SetP95(v int32)`

SetP95 sets P95 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


