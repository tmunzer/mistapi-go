# SleClassifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impact** | [**SleClassifierImpact**](SleClassifierImpact.md) |  | 
**Interval** | **float32** |  | 
**Name** | **string** |  | 
**Samples** | Pointer to [**SleClassifierSamples**](SleClassifierSamples.md) |  | [optional] 
**XLabel** | **string** |  | 
**YLabel** | **string** |  | 

## Methods

### NewSleClassifier

`func NewSleClassifier(impact SleClassifierImpact, interval float32, name string, xLabel string, yLabel string, ) *SleClassifier`

NewSleClassifier instantiates a new SleClassifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleClassifierWithDefaults

`func NewSleClassifierWithDefaults() *SleClassifier`

NewSleClassifierWithDefaults instantiates a new SleClassifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpact

`func (o *SleClassifier) GetImpact() SleClassifierImpact`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *SleClassifier) GetImpactOk() (*SleClassifierImpact, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *SleClassifier) SetImpact(v SleClassifierImpact)`

SetImpact sets Impact field to given value.


### GetInterval

`func (o *SleClassifier) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SleClassifier) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SleClassifier) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetName

`func (o *SleClassifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SleClassifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SleClassifier) SetName(v string)`

SetName sets Name field to given value.


### GetSamples

`func (o *SleClassifier) GetSamples() SleClassifierSamples`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *SleClassifier) GetSamplesOk() (*SleClassifierSamples, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *SleClassifier) SetSamples(v SleClassifierSamples)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *SleClassifier) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetXLabel

`func (o *SleClassifier) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *SleClassifier) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *SleClassifier) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.


### GetYLabel

`func (o *SleClassifier) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *SleClassifier) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *SleClassifier) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


