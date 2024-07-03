# ZoneStatsClientsWaits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg** | Pointer to **float32** | average wait time in seconds | [optional] 
**Max** | Pointer to **float32** | longest wait time in seconds | [optional] 
**Min** | Pointer to **float32** | shortest wait time in seconds | [optional] 
**P95** | Pointer to **float32** | 95th percentile of all the wait time(s) | [optional] 

## Methods

### NewZoneStatsClientsWaits

`func NewZoneStatsClientsWaits() *ZoneStatsClientsWaits`

NewZoneStatsClientsWaits instantiates a new ZoneStatsClientsWaits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStatsClientsWaitsWithDefaults

`func NewZoneStatsClientsWaitsWithDefaults() *ZoneStatsClientsWaits`

NewZoneStatsClientsWaitsWithDefaults instantiates a new ZoneStatsClientsWaits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg

`func (o *ZoneStatsClientsWaits) GetAvg() float32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *ZoneStatsClientsWaits) GetAvgOk() (*float32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *ZoneStatsClientsWaits) SetAvg(v float32)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *ZoneStatsClientsWaits) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetMax

`func (o *ZoneStatsClientsWaits) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ZoneStatsClientsWaits) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ZoneStatsClientsWaits) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ZoneStatsClientsWaits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ZoneStatsClientsWaits) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ZoneStatsClientsWaits) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ZoneStatsClientsWaits) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *ZoneStatsClientsWaits) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetP95

`func (o *ZoneStatsClientsWaits) GetP95() float32`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *ZoneStatsClientsWaits) GetP95Ok() (*float32, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *ZoneStatsClientsWaits) SetP95(v float32)`

SetP95 sets P95 field to given value.

### HasP95

`func (o *ZoneStatsClientsWaits) HasP95() bool`

HasP95 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


