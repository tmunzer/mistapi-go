# ResponseDeviceMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Interval** | **int32** |  | 
**Results** | [**[]ResponseDeviceMetricsResultsItems**](ResponseDeviceMetricsResultsItems.md) |  | 
**Rt** | Pointer to **[]string** |  | [optional] 
**Start** | **int32** |  | 

## Methods

### NewResponseDeviceMetrics

`func NewResponseDeviceMetrics(end int32, interval int32, results []ResponseDeviceMetricsResultsItems, start int32, ) *ResponseDeviceMetrics`

NewResponseDeviceMetrics instantiates a new ResponseDeviceMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceMetricsWithDefaults

`func NewResponseDeviceMetricsWithDefaults() *ResponseDeviceMetrics`

NewResponseDeviceMetricsWithDefaults instantiates a new ResponseDeviceMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseDeviceMetrics) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseDeviceMetrics) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseDeviceMetrics) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetInterval

`func (o *ResponseDeviceMetrics) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ResponseDeviceMetrics) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ResponseDeviceMetrics) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetResults

`func (o *ResponseDeviceMetrics) GetResults() []ResponseDeviceMetricsResultsItems`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseDeviceMetrics) GetResultsOk() (*[]ResponseDeviceMetricsResultsItems, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseDeviceMetrics) SetResults(v []ResponseDeviceMetricsResultsItems)`

SetResults sets Results field to given value.


### GetRt

`func (o *ResponseDeviceMetrics) GetRt() []string`

GetRt returns the Rt field if non-nil, zero value otherwise.

### GetRtOk

`func (o *ResponseDeviceMetrics) GetRtOk() (*[]string, bool)`

GetRtOk returns a tuple with the Rt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRt

`func (o *ResponseDeviceMetrics) SetRt(v []string)`

SetRt sets Rt field to given value.

### HasRt

`func (o *ResponseDeviceMetrics) HasRt() bool`

HasRt returns a boolean if a field has been set.

### GetStart

`func (o *ResponseDeviceMetrics) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseDeviceMetrics) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseDeviceMetrics) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


