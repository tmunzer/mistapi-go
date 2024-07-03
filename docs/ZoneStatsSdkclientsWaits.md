# ZoneStatsSdkclientsWaits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg** | Pointer to **float32** | average wait time in seconds | [optional] 
**Max** | Pointer to **float32** | longest wait time in seconds | [optional] 
**Min** | Pointer to **float32** | shortest wait time in seconds | [optional] 
**P95** | Pointer to **float32** | 95th percentile of all the wait time(s) | [optional] 

## Methods

### NewZoneStatsSdkclientsWaits

`func NewZoneStatsSdkclientsWaits() *ZoneStatsSdkclientsWaits`

NewZoneStatsSdkclientsWaits instantiates a new ZoneStatsSdkclientsWaits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStatsSdkclientsWaitsWithDefaults

`func NewZoneStatsSdkclientsWaitsWithDefaults() *ZoneStatsSdkclientsWaits`

NewZoneStatsSdkclientsWaitsWithDefaults instantiates a new ZoneStatsSdkclientsWaits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg

`func (o *ZoneStatsSdkclientsWaits) GetAvg() float32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *ZoneStatsSdkclientsWaits) GetAvgOk() (*float32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *ZoneStatsSdkclientsWaits) SetAvg(v float32)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *ZoneStatsSdkclientsWaits) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetMax

`func (o *ZoneStatsSdkclientsWaits) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ZoneStatsSdkclientsWaits) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ZoneStatsSdkclientsWaits) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ZoneStatsSdkclientsWaits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ZoneStatsSdkclientsWaits) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ZoneStatsSdkclientsWaits) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ZoneStatsSdkclientsWaits) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *ZoneStatsSdkclientsWaits) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetP95

`func (o *ZoneStatsSdkclientsWaits) GetP95() float32`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *ZoneStatsSdkclientsWaits) GetP95Ok() (*float32, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *ZoneStatsSdkclientsWaits) SetP95(v float32)`

SetP95 sets P95 field to given value.

### HasP95

`func (o *ZoneStatsSdkclientsWaits) HasP95() bool`

HasP95 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


