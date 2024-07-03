# SleImpactSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | [**[]SleImpactSummaryApItem**](SleImpactSummaryApItem.md) |  | 
**Band** | [**[]SleImpactSummaryBandItem**](SleImpactSummaryBandItem.md) |  | 
**Classifier** | **string** |  | 
**DeviceOs** | [**[]SleImpactSummaryDeviceOsItem**](SleImpactSummaryDeviceOsItem.md) |  | 
**DeviceType** | [**[]SleImpactSummaryDeviceTypeItem**](SleImpactSummaryDeviceTypeItem.md) |  | 
**End** | **float32** |  | 
**Failure** | **string** |  | 
**Metric** | **string** |  | 
**Start** | **float32** |  | 
**Wlan** | [**[]SleImpactSummaryWlanItem**](SleImpactSummaryWlanItem.md) |  | 

## Methods

### NewSleImpactSummary

`func NewSleImpactSummary(ap []SleImpactSummaryApItem, band []SleImpactSummaryBandItem, classifier string, deviceOs []SleImpactSummaryDeviceOsItem, deviceType []SleImpactSummaryDeviceTypeItem, end float32, failure string, metric string, start float32, wlan []SleImpactSummaryWlanItem, ) *SleImpactSummary`

NewSleImpactSummary instantiates a new SleImpactSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleImpactSummaryWithDefaults

`func NewSleImpactSummaryWithDefaults() *SleImpactSummary`

NewSleImpactSummaryWithDefaults instantiates a new SleImpactSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *SleImpactSummary) GetAp() []SleImpactSummaryApItem`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *SleImpactSummary) GetApOk() (*[]SleImpactSummaryApItem, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *SleImpactSummary) SetAp(v []SleImpactSummaryApItem)`

SetAp sets Ap field to given value.


### GetBand

`func (o *SleImpactSummary) GetBand() []SleImpactSummaryBandItem`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *SleImpactSummary) GetBandOk() (*[]SleImpactSummaryBandItem, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *SleImpactSummary) SetBand(v []SleImpactSummaryBandItem)`

SetBand sets Band field to given value.


### GetClassifier

`func (o *SleImpactSummary) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *SleImpactSummary) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *SleImpactSummary) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.


### GetDeviceOs

`func (o *SleImpactSummary) GetDeviceOs() []SleImpactSummaryDeviceOsItem`

GetDeviceOs returns the DeviceOs field if non-nil, zero value otherwise.

### GetDeviceOsOk

`func (o *SleImpactSummary) GetDeviceOsOk() (*[]SleImpactSummaryDeviceOsItem, bool)`

GetDeviceOsOk returns a tuple with the DeviceOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOs

`func (o *SleImpactSummary) SetDeviceOs(v []SleImpactSummaryDeviceOsItem)`

SetDeviceOs sets DeviceOs field to given value.


### GetDeviceType

`func (o *SleImpactSummary) GetDeviceType() []SleImpactSummaryDeviceTypeItem`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SleImpactSummary) GetDeviceTypeOk() (*[]SleImpactSummaryDeviceTypeItem, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SleImpactSummary) SetDeviceType(v []SleImpactSummaryDeviceTypeItem)`

SetDeviceType sets DeviceType field to given value.


### GetEnd

`func (o *SleImpactSummary) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleImpactSummary) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleImpactSummary) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetFailure

`func (o *SleImpactSummary) GetFailure() string`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *SleImpactSummary) GetFailureOk() (*string, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *SleImpactSummary) SetFailure(v string)`

SetFailure sets Failure field to given value.


### GetMetric

`func (o *SleImpactSummary) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleImpactSummary) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleImpactSummary) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetStart

`func (o *SleImpactSummary) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleImpactSummary) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleImpactSummary) SetStart(v float32)`

SetStart sets Start field to given value.


### GetWlan

`func (o *SleImpactSummary) GetWlan() []SleImpactSummaryWlanItem`

GetWlan returns the Wlan field if non-nil, zero value otherwise.

### GetWlanOk

`func (o *SleImpactSummary) GetWlanOk() (*[]SleImpactSummaryWlanItem, bool)`

GetWlanOk returns a tuple with the Wlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlan

`func (o *SleImpactSummary) SetWlan(v []SleImpactSummaryWlanItem)`

SetWlan sets Wlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


