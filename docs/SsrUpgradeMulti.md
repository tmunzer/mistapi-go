# SsrUpgradeMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**SsrUpgradeChannel**](SsrUpgradeChannel.md) |  | [optional] [default to SSRUPGRADECHANNEL_STABLE]
**DeviceIds** | **[]string** | list of 128T device IDs to upgrade | 
**RebootAt** | Pointer to **int32** | reboot start time in epoch seconds, default is start_time, -1 disables reboot | [optional] 
**StartTime** | Pointer to **int32** | 128T firmware download start time in epoch seconds, default is now, -1 disables download | [optional] 
**Strategy** | Pointer to [**SsrUpgradeStrategy**](SsrUpgradeStrategy.md) |  | [optional] [default to SSRUPGRADESTRATEGY_BIG_BANG]
**Version** | Pointer to **string** | 128T firmware version to upgrade (e.g. 5.3.0-93) | [optional] 

## Methods

### NewSsrUpgradeMulti

`func NewSsrUpgradeMulti(deviceIds []string, ) *SsrUpgradeMulti`

NewSsrUpgradeMulti instantiates a new SsrUpgradeMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsrUpgradeMultiWithDefaults

`func NewSsrUpgradeMultiWithDefaults() *SsrUpgradeMulti`

NewSsrUpgradeMultiWithDefaults instantiates a new SsrUpgradeMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SsrUpgradeMulti) GetChannel() SsrUpgradeChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SsrUpgradeMulti) GetChannelOk() (*SsrUpgradeChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SsrUpgradeMulti) SetChannel(v SsrUpgradeChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SsrUpgradeMulti) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDeviceIds

`func (o *SsrUpgradeMulti) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SsrUpgradeMulti) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *SsrUpgradeMulti) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.


### GetRebootAt

`func (o *SsrUpgradeMulti) GetRebootAt() int32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *SsrUpgradeMulti) GetRebootAtOk() (*int32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *SsrUpgradeMulti) SetRebootAt(v int32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *SsrUpgradeMulti) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetStartTime

`func (o *SsrUpgradeMulti) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SsrUpgradeMulti) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SsrUpgradeMulti) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SsrUpgradeMulti) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStrategy

`func (o *SsrUpgradeMulti) GetStrategy() SsrUpgradeStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *SsrUpgradeMulti) GetStrategyOk() (*SsrUpgradeStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *SsrUpgradeMulti) SetStrategy(v SsrUpgradeStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *SsrUpgradeMulti) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetVersion

`func (o *SsrUpgradeMulti) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SsrUpgradeMulti) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SsrUpgradeMulti) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SsrUpgradeMulti) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


