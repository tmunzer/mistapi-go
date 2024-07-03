# SsrUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**SsrUpgradeChannel**](SsrUpgradeChannel.md) |  | [optional] [default to SSRUPGRADECHANNEL_STABLE]
**RebootAt** | Pointer to **int32** | eboot start time in epoch seconds, default is start_time, -1 disables reboot | [optional] 
**StartTime** | Pointer to **int32** | 128T firmware download start time in epoch seconds, default is now, -1 disables download | [optional] 
**Version** | **string** | 128T firmware version to upgrade (e.g. 5.3.0-93) | [default to "stable"]

## Methods

### NewSsrUpgrade

`func NewSsrUpgrade(version string, ) *SsrUpgrade`

NewSsrUpgrade instantiates a new SsrUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsrUpgradeWithDefaults

`func NewSsrUpgradeWithDefaults() *SsrUpgrade`

NewSsrUpgradeWithDefaults instantiates a new SsrUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SsrUpgrade) GetChannel() SsrUpgradeChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SsrUpgrade) GetChannelOk() (*SsrUpgradeChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SsrUpgrade) SetChannel(v SsrUpgradeChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SsrUpgrade) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRebootAt

`func (o *SsrUpgrade) GetRebootAt() int32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *SsrUpgrade) GetRebootAtOk() (*int32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *SsrUpgrade) SetRebootAt(v int32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *SsrUpgrade) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetStartTime

`func (o *SsrUpgrade) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SsrUpgrade) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SsrUpgrade) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SsrUpgrade) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *SsrUpgrade) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SsrUpgrade) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SsrUpgrade) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


