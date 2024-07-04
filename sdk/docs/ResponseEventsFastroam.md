# ResponseEventsFastroam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** | the link to query next set of results. value is null if no next page exists. | [optional] 
**Results** | [**[]EventFastroam**](EventFastroam.md) |  | 
**Start** | **int32** |  | 

## Methods

### NewResponseEventsFastroam

`func NewResponseEventsFastroam(end int32, limit int32, results []EventFastroam, start int32, ) *ResponseEventsFastroam`

NewResponseEventsFastroam instantiates a new ResponseEventsFastroam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEventsFastroamWithDefaults

`func NewResponseEventsFastroamWithDefaults() *ResponseEventsFastroam`

NewResponseEventsFastroamWithDefaults instantiates a new ResponseEventsFastroam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseEventsFastroam) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseEventsFastroam) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseEventsFastroam) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseEventsFastroam) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseEventsFastroam) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseEventsFastroam) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseEventsFastroam) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseEventsFastroam) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseEventsFastroam) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseEventsFastroam) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseEventsFastroam) GetResults() []EventFastroam`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseEventsFastroam) GetResultsOk() (*[]EventFastroam, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseEventsFastroam) SetResults(v []EventFastroam)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseEventsFastroam) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseEventsFastroam) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseEventsFastroam) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


