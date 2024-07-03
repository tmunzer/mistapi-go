# RepsonseCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distinct** | **string** |  | 
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Results** | [**[]CountResult**](CountResult.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewRepsonseCount

`func NewRepsonseCount(distinct string, end int32, limit int32, results []CountResult, start int32, total int32, ) *RepsonseCount`

NewRepsonseCount instantiates a new RepsonseCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepsonseCountWithDefaults

`func NewRepsonseCountWithDefaults() *RepsonseCount`

NewRepsonseCountWithDefaults instantiates a new RepsonseCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinct

`func (o *RepsonseCount) GetDistinct() string`

GetDistinct returns the Distinct field if non-nil, zero value otherwise.

### GetDistinctOk

`func (o *RepsonseCount) GetDistinctOk() (*string, bool)`

GetDistinctOk returns a tuple with the Distinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinct

`func (o *RepsonseCount) SetDistinct(v string)`

SetDistinct sets Distinct field to given value.


### GetEnd

`func (o *RepsonseCount) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *RepsonseCount) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *RepsonseCount) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *RepsonseCount) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RepsonseCount) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RepsonseCount) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *RepsonseCount) GetResults() []CountResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RepsonseCount) GetResultsOk() (*[]CountResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RepsonseCount) SetResults(v []CountResult)`

SetResults sets Results field to given value.


### GetStart

`func (o *RepsonseCount) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RepsonseCount) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RepsonseCount) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *RepsonseCount) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RepsonseCount) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RepsonseCount) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


