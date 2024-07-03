# ResponseAnomalySearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Page** | **int32** |  | 
**Results** | [**[]Anomaly**](Anomaly.md) |  | 
**Start** | **int32** |  | 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseAnomalySearch

`func NewResponseAnomalySearch(end int32, limit int32, page int32, results []Anomaly, start int32, ) *ResponseAnomalySearch`

NewResponseAnomalySearch instantiates a new ResponseAnomalySearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAnomalySearchWithDefaults

`func NewResponseAnomalySearchWithDefaults() *ResponseAnomalySearch`

NewResponseAnomalySearchWithDefaults instantiates a new ResponseAnomalySearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseAnomalySearch) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseAnomalySearch) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseAnomalySearch) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseAnomalySearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseAnomalySearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseAnomalySearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *ResponseAnomalySearch) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResponseAnomalySearch) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResponseAnomalySearch) SetPage(v int32)`

SetPage sets Page field to given value.


### GetResults

`func (o *ResponseAnomalySearch) GetResults() []Anomaly`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseAnomalySearch) GetResultsOk() (*[]Anomaly, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseAnomalySearch) SetResults(v []Anomaly)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseAnomalySearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseAnomalySearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseAnomalySearch) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseAnomalySearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseAnomalySearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseAnomalySearch) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponseAnomalySearch) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


