# ResponseStatsCalls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **float32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]CallStats**](CallStats.md) |  | [optional] 
**Start** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStatsCalls

`func NewResponseStatsCalls() *ResponseStatsCalls`

NewResponseStatsCalls instantiates a new ResponseStatsCalls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStatsCallsWithDefaults

`func NewResponseStatsCallsWithDefaults() *ResponseStatsCalls`

NewResponseStatsCallsWithDefaults instantiates a new ResponseStatsCalls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseStatsCalls) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseStatsCalls) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseStatsCalls) SetEnd(v float32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ResponseStatsCalls) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseStatsCalls) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseStatsCalls) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseStatsCalls) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseStatsCalls) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *ResponseStatsCalls) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResponseStatsCalls) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResponseStatsCalls) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResponseStatsCalls) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *ResponseStatsCalls) GetResults() []CallStats`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseStatsCalls) GetResultsOk() (*[]CallStats, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseStatsCalls) SetResults(v []CallStats)`

SetResults sets Results field to given value.

### HasResults

`func (o *ResponseStatsCalls) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStart

`func (o *ResponseStatsCalls) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseStatsCalls) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseStatsCalls) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ResponseStatsCalls) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotal

`func (o *ResponseStatsCalls) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseStatsCalls) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseStatsCalls) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponseStatsCalls) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


