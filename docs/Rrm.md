# Rrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band24** | [**map[string]RrmBand**](RrmBand.md) | proposal on band 2.4G, key is ap_id, value is the proposal | 
**Band24Metric** | [**RrmBandMetric**](RrmBandMetric.md) |  | 
**Band5** | [**map[string]RrmBand**](RrmBand.md) | proposal on band 5G, key is ap_id, value is the proposal | 
**Band5Metric** | [**RrmBandMetric**](RrmBandMetric.md) |  | 
**Band6** | Pointer to [**map[string]RrmBand**](RrmBand.md) | proposal on band 6G, key is ap_id, value is the proposal | [optional] 
**Band6Metric** | Pointer to [**RrmBandMetric**](RrmBandMetric.md) |  | [optional] 
**Rftemplate** | [**RfTemplate**](RfTemplate.md) |  | 
**RftemplateId** | **string** |  | 
**RftemplateName** | **string** |  | 
**Status** | [**RrmStatus**](RrmStatus.md) |  | 
**Timestamp** | **float32** | time where the status was updated | 

## Methods

### NewRrm

`func NewRrm(band24 map[string]RrmBand, band24Metric RrmBandMetric, band5 map[string]RrmBand, band5Metric RrmBandMetric, rftemplate RfTemplate, rftemplateId string, rftemplateName string, status RrmStatus, timestamp float32, ) *Rrm`

NewRrm instantiates a new Rrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmWithDefaults

`func NewRrmWithDefaults() *Rrm`

NewRrmWithDefaults instantiates a new Rrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand24

`func (o *Rrm) GetBand24() map[string]RrmBand`

GetBand24 returns the Band24 field if non-nil, zero value otherwise.

### GetBand24Ok

`func (o *Rrm) GetBand24Ok() (*map[string]RrmBand, bool)`

GetBand24Ok returns a tuple with the Band24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24

`func (o *Rrm) SetBand24(v map[string]RrmBand)`

SetBand24 sets Band24 field to given value.


### GetBand24Metric

`func (o *Rrm) GetBand24Metric() RrmBandMetric`

GetBand24Metric returns the Band24Metric field if non-nil, zero value otherwise.

### GetBand24MetricOk

`func (o *Rrm) GetBand24MetricOk() (*RrmBandMetric, bool)`

GetBand24MetricOk returns a tuple with the Band24Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Metric

`func (o *Rrm) SetBand24Metric(v RrmBandMetric)`

SetBand24Metric sets Band24Metric field to given value.


### GetBand5

`func (o *Rrm) GetBand5() map[string]RrmBand`

GetBand5 returns the Band5 field if non-nil, zero value otherwise.

### GetBand5Ok

`func (o *Rrm) GetBand5Ok() (*map[string]RrmBand, bool)`

GetBand5Ok returns a tuple with the Band5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5

`func (o *Rrm) SetBand5(v map[string]RrmBand)`

SetBand5 sets Band5 field to given value.


### GetBand5Metric

`func (o *Rrm) GetBand5Metric() RrmBandMetric`

GetBand5Metric returns the Band5Metric field if non-nil, zero value otherwise.

### GetBand5MetricOk

`func (o *Rrm) GetBand5MetricOk() (*RrmBandMetric, bool)`

GetBand5MetricOk returns a tuple with the Band5Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Metric

`func (o *Rrm) SetBand5Metric(v RrmBandMetric)`

SetBand5Metric sets Band5Metric field to given value.


### GetBand6

`func (o *Rrm) GetBand6() map[string]RrmBand`

GetBand6 returns the Band6 field if non-nil, zero value otherwise.

### GetBand6Ok

`func (o *Rrm) GetBand6Ok() (*map[string]RrmBand, bool)`

GetBand6Ok returns a tuple with the Band6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6

`func (o *Rrm) SetBand6(v map[string]RrmBand)`

SetBand6 sets Band6 field to given value.

### HasBand6

`func (o *Rrm) HasBand6() bool`

HasBand6 returns a boolean if a field has been set.

### GetBand6Metric

`func (o *Rrm) GetBand6Metric() RrmBandMetric`

GetBand6Metric returns the Band6Metric field if non-nil, zero value otherwise.

### GetBand6MetricOk

`func (o *Rrm) GetBand6MetricOk() (*RrmBandMetric, bool)`

GetBand6MetricOk returns a tuple with the Band6Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Metric

`func (o *Rrm) SetBand6Metric(v RrmBandMetric)`

SetBand6Metric sets Band6Metric field to given value.

### HasBand6Metric

`func (o *Rrm) HasBand6Metric() bool`

HasBand6Metric returns a boolean if a field has been set.

### GetRftemplate

`func (o *Rrm) GetRftemplate() RfTemplate`

GetRftemplate returns the Rftemplate field if non-nil, zero value otherwise.

### GetRftemplateOk

`func (o *Rrm) GetRftemplateOk() (*RfTemplate, bool)`

GetRftemplateOk returns a tuple with the Rftemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplate

`func (o *Rrm) SetRftemplate(v RfTemplate)`

SetRftemplate sets Rftemplate field to given value.


### GetRftemplateId

`func (o *Rrm) GetRftemplateId() string`

GetRftemplateId returns the RftemplateId field if non-nil, zero value otherwise.

### GetRftemplateIdOk

`func (o *Rrm) GetRftemplateIdOk() (*string, bool)`

GetRftemplateIdOk returns a tuple with the RftemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplateId

`func (o *Rrm) SetRftemplateId(v string)`

SetRftemplateId sets RftemplateId field to given value.


### GetRftemplateName

`func (o *Rrm) GetRftemplateName() string`

GetRftemplateName returns the RftemplateName field if non-nil, zero value otherwise.

### GetRftemplateNameOk

`func (o *Rrm) GetRftemplateNameOk() (*string, bool)`

GetRftemplateNameOk returns a tuple with the RftemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRftemplateName

`func (o *Rrm) SetRftemplateName(v string)`

SetRftemplateName sets RftemplateName field to given value.


### GetStatus

`func (o *Rrm) GetStatus() RrmStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rrm) GetStatusOk() (*RrmStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rrm) SetStatus(v RrmStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Rrm) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Rrm) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Rrm) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


