# InsightMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Interval** | **int32** |  | 
**Results** | **[]map[string]interface{}** | results depends on the &#x60;metric&#x60; | 
**Start** | **int32** |  | 

## Methods

### NewInsightMetrics

`func NewInsightMetrics(end int32, interval int32, results []map[string]interface{}, start int32, ) *InsightMetrics`

NewInsightMetrics instantiates a new InsightMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightMetricsWithDefaults

`func NewInsightMetricsWithDefaults() *InsightMetrics`

NewInsightMetricsWithDefaults instantiates a new InsightMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *InsightMetrics) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InsightMetrics) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InsightMetrics) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetInterval

`func (o *InsightMetrics) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InsightMetrics) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InsightMetrics) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetResults

`func (o *InsightMetrics) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InsightMetrics) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InsightMetrics) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.


### GetStart

`func (o *InsightMetrics) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InsightMetrics) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InsightMetrics) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


