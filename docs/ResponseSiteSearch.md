# ResponseSiteSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]ResponseSiteSearchItem**](ResponseSiteSearchItem.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseSiteSearch

`func NewResponseSiteSearch(end int32, limit int32, results []ResponseSiteSearchItem, start int32, total int32, ) *ResponseSiteSearch`

NewResponseSiteSearch instantiates a new ResponseSiteSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSiteSearchWithDefaults

`func NewResponseSiteSearchWithDefaults() *ResponseSiteSearch`

NewResponseSiteSearchWithDefaults instantiates a new ResponseSiteSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseSiteSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseSiteSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseSiteSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseSiteSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseSiteSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseSiteSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseSiteSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseSiteSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseSiteSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseSiteSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseSiteSearch) GetResults() []ResponseSiteSearchItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseSiteSearch) GetResultsOk() (*[]ResponseSiteSearchItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseSiteSearch) SetResults(v []ResponseSiteSearchItem)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseSiteSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseSiteSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseSiteSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseSiteSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseSiteSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseSiteSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


