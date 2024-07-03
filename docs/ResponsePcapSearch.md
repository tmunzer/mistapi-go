# ResponsePcapSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | **string** |  | 
**Results** | [**[]ResponsePcapSearchItem**](ResponsePcapSearchItem.md) |  | 
**Start** | **int32** |  | 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponsePcapSearch

`func NewResponsePcapSearch(end int32, limit int32, next string, results []ResponsePcapSearchItem, start int32, ) *ResponsePcapSearch`

NewResponsePcapSearch instantiates a new ResponsePcapSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePcapSearchWithDefaults

`func NewResponsePcapSearchWithDefaults() *ResponsePcapSearch`

NewResponsePcapSearchWithDefaults instantiates a new ResponsePcapSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponsePcapSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponsePcapSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponsePcapSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponsePcapSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponsePcapSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponsePcapSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponsePcapSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponsePcapSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponsePcapSearch) SetNext(v string)`

SetNext sets Next field to given value.


### GetResults

`func (o *ResponsePcapSearch) GetResults() []ResponsePcapSearchItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponsePcapSearch) GetResultsOk() (*[]ResponsePcapSearchItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponsePcapSearch) SetResults(v []ResponsePcapSearchItem)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponsePcapSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponsePcapSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponsePcapSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponsePcapSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponsePcapSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponsePcapSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponsePcapSearch) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


