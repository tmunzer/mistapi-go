# RrmBand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] 
**Channel** | Pointer to **int32** | proposed channel | [optional] 
**CurrBandwidht** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] 
**CurrChannel** | Pointer to **int32** | current channel | [optional] 
**CurrPower** | Pointer to **int32** | current tx power | [optional] 
**CurrUsage** | Pointer to **string** | current radio band | [optional] 
**Power** | Pointer to **int32** | proposed tx power | [optional] 
**Usage** | Pointer to **string** | proposed radio band | [optional] 

## Methods

### NewRrmBand

`func NewRrmBand() *RrmBand`

NewRrmBand instantiates a new RrmBand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmBandWithDefaults

`func NewRrmBandWithDefaults() *RrmBand`

NewRrmBandWithDefaults instantiates a new RrmBand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *RrmBand) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *RrmBand) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *RrmBand) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *RrmBand) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *RrmBand) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RrmBand) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RrmBand) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *RrmBand) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCurrBandwidht

`func (o *RrmBand) GetCurrBandwidht() Dot11Bandwidth`

GetCurrBandwidht returns the CurrBandwidht field if non-nil, zero value otherwise.

### GetCurrBandwidhtOk

`func (o *RrmBand) GetCurrBandwidhtOk() (*Dot11Bandwidth, bool)`

GetCurrBandwidhtOk returns a tuple with the CurrBandwidht field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrBandwidht

`func (o *RrmBand) SetCurrBandwidht(v Dot11Bandwidth)`

SetCurrBandwidht sets CurrBandwidht field to given value.

### HasCurrBandwidht

`func (o *RrmBand) HasCurrBandwidht() bool`

HasCurrBandwidht returns a boolean if a field has been set.

### GetCurrChannel

`func (o *RrmBand) GetCurrChannel() int32`

GetCurrChannel returns the CurrChannel field if non-nil, zero value otherwise.

### GetCurrChannelOk

`func (o *RrmBand) GetCurrChannelOk() (*int32, bool)`

GetCurrChannelOk returns a tuple with the CurrChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrChannel

`func (o *RrmBand) SetCurrChannel(v int32)`

SetCurrChannel sets CurrChannel field to given value.

### HasCurrChannel

`func (o *RrmBand) HasCurrChannel() bool`

HasCurrChannel returns a boolean if a field has been set.

### GetCurrPower

`func (o *RrmBand) GetCurrPower() int32`

GetCurrPower returns the CurrPower field if non-nil, zero value otherwise.

### GetCurrPowerOk

`func (o *RrmBand) GetCurrPowerOk() (*int32, bool)`

GetCurrPowerOk returns a tuple with the CurrPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrPower

`func (o *RrmBand) SetCurrPower(v int32)`

SetCurrPower sets CurrPower field to given value.

### HasCurrPower

`func (o *RrmBand) HasCurrPower() bool`

HasCurrPower returns a boolean if a field has been set.

### GetCurrUsage

`func (o *RrmBand) GetCurrUsage() string`

GetCurrUsage returns the CurrUsage field if non-nil, zero value otherwise.

### GetCurrUsageOk

`func (o *RrmBand) GetCurrUsageOk() (*string, bool)`

GetCurrUsageOk returns a tuple with the CurrUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrUsage

`func (o *RrmBand) SetCurrUsage(v string)`

SetCurrUsage sets CurrUsage field to given value.

### HasCurrUsage

`func (o *RrmBand) HasCurrUsage() bool`

HasCurrUsage returns a boolean if a field has been set.

### GetPower

`func (o *RrmBand) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *RrmBand) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *RrmBand) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *RrmBand) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetUsage

`func (o *RrmBand) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RrmBand) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RrmBand) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RrmBand) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


