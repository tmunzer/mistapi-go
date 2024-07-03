# ResponseRrmNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** | the link to query next set of results. value is null if no next page exists. | [optional] 
**Results** | [**[]RrmNeighbors**](RrmNeighbors.md) |  | 
**Start** | **int32** |  | 

## Methods

### NewResponseRrmNeighbors

`func NewResponseRrmNeighbors(end int32, limit int32, results []RrmNeighbors, start int32, ) *ResponseRrmNeighbors`

NewResponseRrmNeighbors instantiates a new ResponseRrmNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseRrmNeighborsWithDefaults

`func NewResponseRrmNeighborsWithDefaults() *ResponseRrmNeighbors`

NewResponseRrmNeighborsWithDefaults instantiates a new ResponseRrmNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseRrmNeighbors) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseRrmNeighbors) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseRrmNeighbors) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseRrmNeighbors) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseRrmNeighbors) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseRrmNeighbors) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseRrmNeighbors) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseRrmNeighbors) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseRrmNeighbors) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseRrmNeighbors) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseRrmNeighbors) GetResults() []RrmNeighbors`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseRrmNeighbors) GetResultsOk() (*[]RrmNeighbors, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseRrmNeighbors) SetResults(v []RrmNeighbors)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseRrmNeighbors) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseRrmNeighbors) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseRrmNeighbors) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


