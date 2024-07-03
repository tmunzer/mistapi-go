# ResponseSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Page** | **int32** |  | 
**Results** | [**[]ResponseSearchItem**](ResponseSearchItem.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseSearch

`func NewResponseSearch(limit int32, page int32, results []ResponseSearchItem, total int32, ) *ResponseSearch`

NewResponseSearch instantiates a new ResponseSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSearchWithDefaults

`func NewResponseSearchWithDefaults() *ResponseSearch`

NewResponseSearchWithDefaults instantiates a new ResponseSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ResponseSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponseSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPage

`func (o *ResponseSearch) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResponseSearch) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResponseSearch) SetPage(v int32)`

SetPage sets Page field to given value.


### GetResults

`func (o *ResponseSearch) GetResults() []ResponseSearchItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseSearch) GetResultsOk() (*[]ResponseSearchItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseSearch) SetResults(v []ResponseSearchItem)`

SetResults sets Results field to given value.


### GetTotal

`func (o *ResponseSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


