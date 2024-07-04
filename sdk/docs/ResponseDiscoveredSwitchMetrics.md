# ResponseDiscoveredSwitchMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **float32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]DiscoveredSwitchMetric**](DiscoveredSwitchMetric.md) |  | 
**Start** | **float32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseDiscoveredSwitchMetrics

`func NewResponseDiscoveredSwitchMetrics(end float32, limit int32, results []DiscoveredSwitchMetric, start float32, total int32, ) *ResponseDiscoveredSwitchMetrics`

NewResponseDiscoveredSwitchMetrics instantiates a new ResponseDiscoveredSwitchMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDiscoveredSwitchMetricsWithDefaults

`func NewResponseDiscoveredSwitchMetricsWithDefaults() *ResponseDiscoveredSwitchMetrics`

NewResponseDiscoveredSwitchMetricsWithDefaults instantiates a new ResponseDiscoveredSwitchMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseDiscoveredSwitchMetrics) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseDiscoveredSwitchMetrics) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseDiscoveredSwitchMetrics) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseDiscoveredSwitchMetrics) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseDiscoveredSwitchMetrics) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseDiscoveredSwitchMetrics) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseDiscoveredSwitchMetrics) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseDiscoveredSwitchMetrics) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseDiscoveredSwitchMetrics) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseDiscoveredSwitchMetrics) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseDiscoveredSwitchMetrics) GetResults() []DiscoveredSwitchMetric`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseDiscoveredSwitchMetrics) GetResultsOk() (*[]DiscoveredSwitchMetric, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseDiscoveredSwitchMetrics) SetResults(v []DiscoveredSwitchMetric)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseDiscoveredSwitchMetrics) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseDiscoveredSwitchMetrics) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseDiscoveredSwitchMetrics) SetStart(v float32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseDiscoveredSwitchMetrics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseDiscoveredSwitchMetrics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseDiscoveredSwitchMetrics) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


