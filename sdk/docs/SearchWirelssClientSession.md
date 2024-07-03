# SearchWirelssClientSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **float32** |  | 
**Limit** | **int32** |  | 
**Next** | Pointer to **string** |  | [optional] 
**Results** | [**[]WirelssClientSession**](WirelssClientSession.md) |  | 
**Start** | **float32** |  | 
**Total** | **int32** |  | 

## Methods

### NewSearchWirelssClientSession

`func NewSearchWirelssClientSession(end float32, limit int32, results []WirelssClientSession, start float32, total int32, ) *SearchWirelssClientSession`

NewSearchWirelssClientSession instantiates a new SearchWirelssClientSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWirelssClientSessionWithDefaults

`func NewSearchWirelssClientSessionWithDefaults() *SearchWirelssClientSession`

NewSearchWirelssClientSessionWithDefaults instantiates a new SearchWirelssClientSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *SearchWirelssClientSession) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SearchWirelssClientSession) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SearchWirelssClientSession) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *SearchWirelssClientSession) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchWirelssClientSession) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchWirelssClientSession) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *SearchWirelssClientSession) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchWirelssClientSession) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchWirelssClientSession) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *SearchWirelssClientSession) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResults

`func (o *SearchWirelssClientSession) GetResults() []WirelssClientSession`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchWirelssClientSession) GetResultsOk() (*[]WirelssClientSession, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchWirelssClientSession) SetResults(v []WirelssClientSession)`

SetResults sets Results field to given value.


### GetStart

`func (o *SearchWirelssClientSession) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SearchWirelssClientSession) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SearchWirelssClientSession) SetStart(v float32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *SearchWirelssClientSession) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchWirelssClientSession) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchWirelssClientSession) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


