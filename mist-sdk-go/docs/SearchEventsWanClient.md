# SearchEventsWanClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**EventsClientWan**](EventsClientWan.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchEventsWanClient

`func NewSearchEventsWanClient() *SearchEventsWanClient`

NewSearchEventsWanClient instantiates a new SearchEventsWanClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEventsWanClientWithDefaults

`func NewSearchEventsWanClientWithDefaults() *SearchEventsWanClient`

NewSearchEventsWanClientWithDefaults instantiates a new SearchEventsWanClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *SearchEventsWanClient) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SearchEventsWanClient) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SearchEventsWanClient) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SearchEventsWanClient) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLimit

`func (o *SearchEventsWanClient) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchEventsWanClient) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchEventsWanClient) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchEventsWanClient) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetResults

`func (o *SearchEventsWanClient) GetResults() EventsClientWan`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchEventsWanClient) GetResultsOk() (*EventsClientWan, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchEventsWanClient) SetResults(v EventsClientWan)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchEventsWanClient) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStart

`func (o *SearchEventsWanClient) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SearchEventsWanClient) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SearchEventsWanClient) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SearchEventsWanClient) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotal

`func (o *SearchEventsWanClient) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchEventsWanClient) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchEventsWanClient) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchEventsWanClient) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


