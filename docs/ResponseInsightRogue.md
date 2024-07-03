# ResponseInsightRogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** | link to next set of results. If more results aren’t present, next is null. | [optional] 
**Results** | [**[]InsightRogueAp**](InsightRogueAp.md) |  | 
**Start** | **int32** |  | 

## Methods

### NewResponseInsightRogue

`func NewResponseInsightRogue(end int32, limit int32, results []InsightRogueAp, start int32, ) *ResponseInsightRogue`

NewResponseInsightRogue instantiates a new ResponseInsightRogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseInsightRogueWithDefaults

`func NewResponseInsightRogueWithDefaults() *ResponseInsightRogue`

NewResponseInsightRogueWithDefaults instantiates a new ResponseInsightRogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseInsightRogue) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseInsightRogue) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseInsightRogue) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseInsightRogue) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseInsightRogue) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseInsightRogue) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseInsightRogue) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseInsightRogue) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseInsightRogue) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseInsightRogue) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseInsightRogue) GetResults() []InsightRogueAp`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseInsightRogue) GetResultsOk() (*[]InsightRogueAp, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseInsightRogue) SetResults(v []InsightRogueAp)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseInsightRogue) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseInsightRogue) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseInsightRogue) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


