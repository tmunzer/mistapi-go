# SleHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SleHistogramDataItem**](SleHistogramDataItem.md) |  | 
**End** | **float32** |  | 
**Metric** | **string** |  | 
**Start** | **float32** |  | 
**XLabel** | **string** |  | 
**YLabel** | **string** |  | 

## Methods

### NewSleHistogram

`func NewSleHistogram(data []SleHistogramDataItem, end float32, metric string, start float32, xLabel string, yLabel string, ) *SleHistogram`

NewSleHistogram instantiates a new SleHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleHistogramWithDefaults

`func NewSleHistogramWithDefaults() *SleHistogram`

NewSleHistogramWithDefaults instantiates a new SleHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SleHistogram) GetData() []SleHistogramDataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SleHistogram) GetDataOk() (*[]SleHistogramDataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SleHistogram) SetData(v []SleHistogramDataItem)`

SetData sets Data field to given value.


### GetEnd

`func (o *SleHistogram) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SleHistogram) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SleHistogram) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetMetric

`func (o *SleHistogram) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SleHistogram) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SleHistogram) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetStart

`func (o *SleHistogram) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SleHistogram) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SleHistogram) SetStart(v float32)`

SetStart sets Start field to given value.


### GetXLabel

`func (o *SleHistogram) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *SleHistogram) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *SleHistogram) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.


### GetYLabel

`func (o *SleHistogram) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *SleHistogram) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *SleHistogram) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


