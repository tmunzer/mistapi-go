# ResponseClientSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **float32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]ClientWireless**](ClientWireless.md) |  | 
**Start** | **float32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseClientSearch

`func NewResponseClientSearch(end float32, limit int32, results []ClientWireless, start float32, total int32, ) *ResponseClientSearch`

NewResponseClientSearch instantiates a new ResponseClientSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseClientSearchWithDefaults

`func NewResponseClientSearchWithDefaults() *ResponseClientSearch`

NewResponseClientSearchWithDefaults instantiates a new ResponseClientSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseClientSearch) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseClientSearch) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseClientSearch) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseClientSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseClientSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseClientSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseClientSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseClientSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseClientSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseClientSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseClientSearch) GetResults() []ClientWireless`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseClientSearch) GetResultsOk() (*[]ClientWireless, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseClientSearch) SetResults(v []ClientWireless)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseClientSearch) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseClientSearch) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseClientSearch) SetStart(v float32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseClientSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseClientSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseClientSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


