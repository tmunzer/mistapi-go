# SleClassifierSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | [**SleClassifier**](SleClassifier.md) |  | 
**End** | **float32** |  | 
**Failures** | **[]map[string]interface{}** |  | 
**Impact** | [**SleClassifierSummaryImpact**](SleClassifierSummaryImpact.md) |  | 
**Metric** | **string** |  | 
**Start** | **float32** |  | 

## Methods

### NewSleClassifierSummary

`func NewSleClassifierSummary(classifier SleClassifier, end float32, failures []map[string]interface{}, impact SleClassifierSummaryImpact, metric string, start float32, ) *SleClassifierSummary`

NewSleClassifierSummary instantiates a new SleClassifierSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleClassifierSummaryWithDefaults

`func NewSleClassifierSummaryWithDefaults() *SleClassifierSummary`

NewSleClassifierSummaryWithDefaults instantiates a new SleClassifierSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *SleClassifierSummary) GetClassifier() SleClassifier`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleClassifierSummary) GetClassifierOk() (*SleClassifier, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleClassifierSummary) SetClassifier(v SleClassifier)`

SetClassifier sets Classifier field to given value.


### GetEnd

`func (o *SleClassifierSummary) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleClassifierSummary) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleClassifierSummary) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetFailures

`func (o *SleClassifierSummary) GetFailures() []map[string]interface{}`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *SleClassifierSummary) GetFailuresOk() (*[]map[string]interface{}, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *SleClassifierSummary) SetFailures(v []map[string]interface{})`

SetFailures sets Failures field to given value.


### GetImpact

`func (o *SleClassifierSummary) GetImpact() SleClassifierSummaryImpact`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *SleClassifierSummary) GetImpactOk() (*SleClassifierSummaryImpact, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *SleClassifierSummary) SetImpact(v SleClassifierSummaryImpact)`

SetImpact sets Impact field to given value.


### GetMetric

`func (o *SleClassifierSummary) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleClassifierSummary) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleClassifierSummary) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetStart

`func (o *SleClassifierSummary) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleClassifierSummary) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleClassifierSummary) SetStart(v float32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


