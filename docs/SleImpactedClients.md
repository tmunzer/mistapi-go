# SleImpactedClients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**Clients** | Pointer to [**[]SleImpactedClientsClient**](SleImpactedClientsClient.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Failure** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSleImpactedClients

`func NewSleImpactedClients() *SleImpactedClients`

NewSleImpactedClients instantiates a new SleImpactedClients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactedClientsWithDefaults

`func NewSleImpactedClientsWithDefaults() *SleImpactedClients`

NewSleImpactedClientsWithDefaults instantiates a new SleImpactedClients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *SleImpactedClients) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactedClients) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactedClients) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *SleImpactedClients) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetClients

`func (o *SleImpactedClients) GetClients() []SleImpactedClientsClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *SleImpactedClients) GetClientsOk() (*[]SleImpactedClientsClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *SleImpactedClients) SetClients(v []SleImpactedClientsClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *SleImpactedClients) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetEnd

`func (o *SleImpactedClients) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactedClients) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactedClients) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SleImpactedClients) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailure

`func (o *SleImpactedClients) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactedClients) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactedClients) SetFailure(v string)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *SleImpactedClients) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetLimit

`func (o *SleImpactedClients) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SleImpactedClients) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SleImpactedClients) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SleImpactedClients) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *SleImpactedClients) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactedClients) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactedClients) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SleImpactedClients) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPage

`func (o *SleImpactedClients) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SleImpactedClients) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SleImpactedClients) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SleImpactedClients) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStart

`func (o *SleImpactedClients) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactedClients) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactedClients) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SleImpactedClients) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalCount

`func (o *SleImpactedClients) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SleImpactedClients) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SleImpactedClients) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SleImpactedClients) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


