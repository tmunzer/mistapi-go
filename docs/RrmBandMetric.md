# RrmBandMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CochannelNeighbors** | **float32** | average number of co-channel neighbors | 
**Density** | **float32** | defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone) | 
**Interferences** | Pointer to [**map[string]RrmBandMetricInterference**](RrmBandMetricInterference.md) | Property key is the channel number | [optional] 
**Neighbors** | **float32** | average number of neighbors | 
**Noise** | **float32** | average noise in dBm | 

## Methods

### NewRrmBandMetric

`func NewRrmBandMetric(cochannelNeighbors float32, density float32, neighbors float32, noise float32, ) *RrmBandMetric`

NewRrmBandMetric instantiates a new RrmBandMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmBandMetricWithDefaults

`func NewRrmBandMetricWithDefaults() *RrmBandMetric`

NewRrmBandMetricWithDefaults instantiates a new RrmBandMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCochannelNeighbors

`func (o *RrmBandMetric) GetCochannelNeighbors() float32`

GetCochannelNeighbors returns the CochannelNeighbors field if non-nil, zero value otherwise.

### GetCochannelNeighborsOk

`func (o *RrmBandMetric) GetCochannelNeighborsOk() (*float32, bool)`

GetCochannelNeighborsOk returns a tuple with the CochannelNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCochannelNeighbors

`func (o *RrmBandMetric) SetCochannelNeighbors(v float32)`

SetCochannelNeighbors sets CochannelNeighbors field to given value.


### GetDensity

`func (o *RrmBandMetric) GetDensity() float32`

GetDensity returns the Density field if non-nil, zero value otherwise.

### GetDensityOk

`func (o *RrmBandMetric) GetDensityOk() (*float32, bool)`

GetDensityOk returns a tuple with the Density field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDensity

`func (o *RrmBandMetric) SetDensity(v float32)`

SetDensity sets Density field to given value.


### GetInterferences

`func (o *RrmBandMetric) GetInterferences() map[string]RrmBandMetricInterference`

GetInterferences returns the Interferences field if non-nil, zero value otherwise.

### GetInterferencesOk

`func (o *RrmBandMetric) GetInterferencesOk() (*map[string]RrmBandMetricInterference, bool)`

GetInterferencesOk returns a tuple with the Interferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterferences

`func (o *RrmBandMetric) SetInterferences(v map[string]RrmBandMetricInterference)`

SetInterferences sets Interferences field to given value.

### HasInterferences

`func (o *RrmBandMetric) HasInterferences() bool`

HasInterferences returns a boolean if a field has been set.

### GetNeighbors

`func (o *RrmBandMetric) GetNeighbors() float32`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *RrmBandMetric) GetNeighborsOk() (*float32, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *RrmBandMetric) SetNeighbors(v float32)`

SetNeighbors sets Neighbors field to given value.


### GetNoise

`func (o *RrmBandMetric) GetNoise() float32`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *RrmBandMetric) GetNoiseOk() (*float32, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *RrmBandMetric) SetNoise(v float32)`

SetNoise sets Noise field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


