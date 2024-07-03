# RrmEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApId** | **string** |  | 
**Band** | [**Dot11Band**](Dot11Band.md) |  | 
**Bandwidth** | [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | 
**Channel** | **int32** | channel for the band from rrm | 
**Event** | [**RrmEventType**](RrmEventType.md) |  | 
**Power** | **int32** | tx power of the radio | 
**PreBandwidth** | [**RrmEventPreBandwidth**](RrmEventPreBandwidth.md) |  | 
**PreChannel** | **int32** | (previously) channel for the band, 0 means no previously available | 
**PrePower** | **float32** | (previously) tx power of the radio, 0 means no previously available | 
**PreUsage** | **string** |  | 
**Timestamp** | **float32** | timestamp of the event | 
**Usage** | **string** |  | 

## Methods

### NewRrmEvent

`func NewRrmEvent(apId string, band Dot11Band, bandwidth Dot11Bandwidth, channel int32, event RrmEventType, power int32, preBandwidth RrmEventPreBandwidth, preChannel int32, prePower float32, preUsage string, timestamp float32, usage string, ) *RrmEvent`

NewRrmEvent instantiates a new RrmEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmEventWithDefaults

`func NewRrmEventWithDefaults() *RrmEvent`

NewRrmEventWithDefaults instantiates a new RrmEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApId

`func (o *RrmEvent) GetApId() string`

GetApId returns the ApId field if non-nil, zero value otherwise.

### GetApIdOk

`func (o *RrmEvent) GetApIdOk() (*string, bool)`

GetApIdOk returns a tuple with the ApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApId

`func (o *RrmEvent) SetApId(v string)`

SetApId sets ApId field to given value.


### GetBand

`func (o *RrmEvent) GetBand() Dot11Band`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *RrmEvent) GetBandOk() (*Dot11Band, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *RrmEvent) SetBand(v Dot11Band)`

SetBand sets Band field to given value.


### GetBandwidth

`func (o *RrmEvent) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *RrmEvent) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *RrmEvent) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.


### GetChannel

`func (o *RrmEvent) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RrmEvent) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RrmEvent) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetEvent

`func (o *RrmEvent) GetEvent() RrmEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RrmEvent) GetEventOk() (*RrmEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RrmEvent) SetEvent(v RrmEventType)`

SetEvent sets Event field to given value.


### GetPower

`func (o *RrmEvent) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *RrmEvent) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *RrmEvent) SetPower(v int32)`

SetPower sets Power field to given value.


### GetPreBandwidth

`func (o *RrmEvent) GetPreBandwidth() RrmEventPreBandwidth`

GetPreBandwidth returns the PreBandwidth field if non-nil, zero value otherwise.

### GetPreBandwidthOk

`func (o *RrmEvent) GetPreBandwidthOk() (*RrmEventPreBandwidth, bool)`

GetPreBandwidthOk returns a tuple with the PreBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBandwidth

`func (o *RrmEvent) SetPreBandwidth(v RrmEventPreBandwidth)`

SetPreBandwidth sets PreBandwidth field to given value.


### GetPreChannel

`func (o *RrmEvent) GetPreChannel() int32`

GetPreChannel returns the PreChannel field if non-nil, zero value otherwise.

### GetPreChannelOk

`func (o *RrmEvent) GetPreChannelOk() (*int32, bool)`

GetPreChannelOk returns a tuple with the PreChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreChannel

`func (o *RrmEvent) SetPreChannel(v int32)`

SetPreChannel sets PreChannel field to given value.


### GetPrePower

`func (o *RrmEvent) GetPrePower() float32`

GetPrePower returns the PrePower field if non-nil, zero value otherwise.

### GetPrePowerOk

`func (o *RrmEvent) GetPrePowerOk() (*float32, bool)`

GetPrePowerOk returns a tuple with the PrePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePower

`func (o *RrmEvent) SetPrePower(v float32)`

SetPrePower sets PrePower field to given value.


### GetPreUsage

`func (o *RrmEvent) GetPreUsage() string`

GetPreUsage returns the PreUsage field if non-nil, zero value otherwise.

### GetPreUsageOk

`func (o *RrmEvent) GetPreUsageOk() (*string, bool)`

GetPreUsageOk returns a tuple with the PreUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreUsage

`func (o *RrmEvent) SetPreUsage(v string)`

SetPreUsage sets PreUsage field to given value.


### GetTimestamp

`func (o *RrmEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RrmEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RrmEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetUsage

`func (o *RrmEvent) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RrmEvent) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RrmEvent) SetUsage(v string)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


