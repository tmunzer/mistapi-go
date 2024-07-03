# ReponseSyntheticTestSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]SyntheticTestInfo**](SyntheticTestInfo.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewReponseSyntheticTestSearch

`func NewReponseSyntheticTestSearch(end int32, limit int32, results []SyntheticTestInfo, start int32, total int32, ) *ReponseSyntheticTestSearch`

NewReponseSyntheticTestSearch instantiates a new ReponseSyntheticTestSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReponseSyntheticTestSearchWithDefaults

`func NewReponseSyntheticTestSearchWithDefaults() *ReponseSyntheticTestSearch`

NewReponseSyntheticTestSearchWithDefaults instantiates a new ReponseSyntheticTestSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ReponseSyntheticTestSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ReponseSyntheticTestSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ReponseSyntheticTestSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ReponseSyntheticTestSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReponseSyntheticTestSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReponseSyntheticTestSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *ReponseSyntheticTestSearch) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReponseSyntheticTestSearch) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReponseSyntheticTestSearch) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ReponseSyntheticTestSearch) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ReponseSyntheticTestSearch) GetResults() []SyntheticTestInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ReponseSyntheticTestSearch) GetResultsOk() (*[]SyntheticTestInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ReponseSyntheticTestSearch) SetResults(v []SyntheticTestInfo)`

SetResults sets Results field to given value.


### GetStart

`func (o *ReponseSyntheticTestSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ReponseSyntheticTestSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ReponseSyntheticTestSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ReponseSyntheticTestSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReponseSyntheticTestSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReponseSyntheticTestSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


