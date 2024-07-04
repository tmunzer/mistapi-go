# ResponseEventsRrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** | the link to query next set of results. value is null if no next page exists. | [optional] 
**Results** | [**[]RrmEvent**](RrmEvent.md) |  | 
**Start** | **int32** |  | 

## Methods

### NewResponseEventsRrm

`func NewResponseEventsRrm(end int32, limit int32, results []RrmEvent, start int32, ) *ResponseEventsRrm`

NewResponseEventsRrm instantiates a new ResponseEventsRrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEventsRrmWithDefaults

`func NewResponseEventsRrmWithDefaults() *ResponseEventsRrm`

NewResponseEventsRrmWithDefaults instantiates a new ResponseEventsRrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseEventsRrm) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseEventsRrm) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseEventsRrm) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseEventsRrm) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseEventsRrm) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseEventsRrm) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseEventsRrm) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseEventsRrm) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseEventsRrm) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseEventsRrm) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseEventsRrm) GetResults() []RrmEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseEventsRrm) GetResultsOk() (*[]RrmEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseEventsRrm) SetResults(v []RrmEvent)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseEventsRrm) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseEventsRrm) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseEventsRrm) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

