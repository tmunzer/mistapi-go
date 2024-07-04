# SleImpactedInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Failure** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]SleImpactedInterfacesInterface**](SleImpactedInterfacesInterface.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSleImpactedInterfaces

`func NewSleImpactedInterfaces() *SleImpactedInterfaces`

NewSleImpactedInterfaces instantiates a new SleImpactedInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactedInterfacesWithDefaults

`func NewSleImpactedInterfacesWithDefaults() *SleImpactedInterfaces`

NewSleImpactedInterfacesWithDefaults instantiates a new SleImpactedInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *SleImpactedInterfaces) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactedInterfaces) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactedInterfaces) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *SleImpactedInterfaces) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetEnd

`func (o *SleImpactedInterfaces) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactedInterfaces) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactedInterfaces) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SleImpactedInterfaces) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailure

`func (o *SleImpactedInterfaces) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactedInterfaces) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactedInterfaces) SetFailure(v string)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *SleImpactedInterfaces) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetInterfaces

`func (o *SleImpactedInterfaces) GetInterfaces() []SleImpactedInterfacesInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *SleImpactedInterfaces) GetInterfacesOk() (*[]SleImpactedInterfacesInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *SleImpactedInterfaces) SetInterfaces(v []SleImpactedInterfacesInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *SleImpactedInterfaces) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLimit

`func (o *SleImpactedInterfaces) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SleImpactedInterfaces) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SleImpactedInterfaces) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SleImpactedInterfaces) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *SleImpactedInterfaces) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactedInterfaces) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactedInterfaces) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SleImpactedInterfaces) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPage

`func (o *SleImpactedInterfaces) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SleImpactedInterfaces) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SleImpactedInterfaces) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SleImpactedInterfaces) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStart

`func (o *SleImpactedInterfaces) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactedInterfaces) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactedInterfaces) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SleImpactedInterfaces) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalCount

`func (o *SleImpactedInterfaces) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SleImpactedInterfaces) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SleImpactedInterfaces) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SleImpactedInterfaces) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


