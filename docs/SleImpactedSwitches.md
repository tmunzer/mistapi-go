# SleImpactedSwitches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Failure** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Switches** | Pointer to [**[]SleImpactedSwitchesSwitch**](SleImpactedSwitchesSwitch.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSleImpactedSwitches

`func NewSleImpactedSwitches() *SleImpactedSwitches`

NewSleImpactedSwitches instantiates a new SleImpactedSwitches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactedSwitchesWithDefaults

`func NewSleImpactedSwitchesWithDefaults() *SleImpactedSwitches`

NewSleImpactedSwitchesWithDefaults instantiates a new SleImpactedSwitches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *SleImpactedSwitches) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactedSwitches) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactedSwitches) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *SleImpactedSwitches) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetEnd

`func (o *SleImpactedSwitches) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactedSwitches) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactedSwitches) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SleImpactedSwitches) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFailure

`func (o *SleImpactedSwitches) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactedSwitches) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactedSwitches) SetFailure(v string)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *SleImpactedSwitches) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetLimit

`func (o *SleImpactedSwitches) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SleImpactedSwitches) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SleImpactedSwitches) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SleImpactedSwitches) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *SleImpactedSwitches) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactedSwitches) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactedSwitches) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SleImpactedSwitches) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPage

`func (o *SleImpactedSwitches) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SleImpactedSwitches) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SleImpactedSwitches) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SleImpactedSwitches) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStart

`func (o *SleImpactedSwitches) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactedSwitches) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactedSwitches) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SleImpactedSwitches) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSwitches

`func (o *SleImpactedSwitches) GetSwitches() []SleImpactedSwitchesSwitch`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *SleImpactedSwitches) GetSwitchesOk() (*[]SleImpactedSwitchesSwitch, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *SleImpactedSwitches) SetSwitches(v []SleImpactedSwitchesSwitch)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *SleImpactedSwitches) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetTotalCount

`func (o *SleImpactedSwitches) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SleImpactedSwitches) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SleImpactedSwitches) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SleImpactedSwitches) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


