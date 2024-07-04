# ResponsePortStatsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]SwitchPortStats**](SwitchPortStats.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponsePortStatsSearch

`func NewResponsePortStatsSearch(end int32, limit int32, results []SwitchPortStats, start int32, total int32, ) *ResponsePortStatsSearch`

NewResponsePortStatsSearch instantiates a new ResponsePortStatsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePortStatsSearchWithDefaults

`func NewResponsePortStatsSearchWithDefaults() *ResponsePortStatsSearch`

NewResponsePortStatsSearchWithDefaults instantiates a new ResponsePortStatsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponsePortStatsSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponsePortStatsSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponsePortStatsSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponsePortStatsSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponsePortStatsSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponsePortStatsSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ResponsePortStatsSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponsePortStatsSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponsePortStatsSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponsePortStatsSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponsePortStatsSearch) GetResults() []SwitchPortStats`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponsePortStatsSearch) GetResultsOk() (*[]SwitchPortStats, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponsePortStatsSearch) SetResults(v []SwitchPortStats)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponsePortStatsSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponsePortStatsSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponsePortStatsSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponsePortStatsSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponsePortStatsSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponsePortStatsSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


