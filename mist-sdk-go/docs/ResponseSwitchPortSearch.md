# ResponseSwitchPortSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Results** | [**[]SwitchPortStats**](SwitchPortStats.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseSwitchPortSearch

`func NewResponseSwitchPortSearch(end int32, limit int32, results []SwitchPortStats, start int32, total int32, ) *ResponseSwitchPortSearch`

NewResponseSwitchPortSearch instantiates a new ResponseSwitchPortSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSwitchPortSearchWithDefaults

`func NewResponseSwitchPortSearchWithDefaults() *ResponseSwitchPortSearch`

NewResponseSwitchPortSearchWithDefaults instantiates a new ResponseSwitchPortSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseSwitchPortSearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseSwitchPortSearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseSwitchPortSearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseSwitchPortSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseSwitchPortSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseSwitchPortSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *ResponseSwitchPortSearch) GetResults() []SwitchPortStats`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseSwitchPortSearch) GetResultsOk() (*[]SwitchPortStats, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseSwitchPortSearch) SetResults(v []SwitchPortStats)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseSwitchPortSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseSwitchPortSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseSwitchPortSearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseSwitchPortSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseSwitchPortSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseSwitchPortSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


