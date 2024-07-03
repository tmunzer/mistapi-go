# ResponseDeviceSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]ResponseDeviceSearchResultsItems**](ResponseDeviceSearchResultsItems.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseDeviceSearch

`func NewResponseDeviceSearch(end int32, limit int32, results []ResponseDeviceSearchResultsItems, start int32, total int32, ) *ResponseDeviceSearch`

NewResponseDeviceSearch instantiates a new ResponseDeviceSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceSearchWithDefaults

`func NewResponseDeviceSearchWithDefaults() *ResponseDeviceSearch`

NewResponseDeviceSearchWithDefaults instantiates a new ResponseDeviceSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseDeviceSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseDeviceSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseDeviceSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseDeviceSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseDeviceSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseDeviceSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseDeviceSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseDeviceSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseDeviceSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseDeviceSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseDeviceSearch) GetResults() []ResponseDeviceSearchResultsItems`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseDeviceSearch) GetResultsOk() (*[]ResponseDeviceSearchResultsItems, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseDeviceSearch) SetResults(v []ResponseDeviceSearchResultsItems)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseDeviceSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseDeviceSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseDeviceSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseDeviceSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseDeviceSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseDeviceSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


