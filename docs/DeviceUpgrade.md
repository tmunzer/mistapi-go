# DeviceUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed (Available on Junos OS devices) | [optional] [default to false]
**RebootAt** | Pointer to **int32** | reboot start time in epoch | [optional] 
**Snapshot** | Pointer to **bool** | Perform recovery snapshot after device is rebooted (Available on Junos OS devices) | [optional] [default to false]
**StartTime** | Pointer to **float32** | firmware download start time in epoch | [optional] 
**Version** | **string** | specific version / &#x60;stable&#x60;, default is to use the latest | [default to "stable"]

## Methods

### NewDeviceUpgrade

`func NewDeviceUpgrade(version string, ) *DeviceUpgrade`

NewDeviceUpgrade instantiates a new DeviceUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpgradeWithDefaults

`func NewDeviceUpgradeWithDefaults() *DeviceUpgrade`

NewDeviceUpgradeWithDefaults instantiates a new DeviceUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *DeviceUpgrade) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *DeviceUpgrade) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *DeviceUpgrade) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *DeviceUpgrade) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRebootAt

`func (o *DeviceUpgrade) GetRebootAt() int32`

GetRebootAt returns the RebootAt field if non-nil, zero value otherwise.

### GetRebootAtOk

`func (o *DeviceUpgrade) GetRebootAtOk() (*int32, bool)`

GetRebootAtOk returns a tuple with the RebootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAt

`func (o *DeviceUpgrade) SetRebootAt(v int32)`

SetRebootAt sets RebootAt field to given value.

### HasRebootAt

`func (o *DeviceUpgrade) HasRebootAt() bool`

HasRebootAt returns a boolean if a field has been set.

### GetSnapshot

`func (o *DeviceUpgrade) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *DeviceUpgrade) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *DeviceUpgrade) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *DeviceUpgrade) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStartTime

`func (o *DeviceUpgrade) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DeviceUpgrade) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DeviceUpgrade) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DeviceUpgrade) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceUpgrade) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceUpgrade) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceUpgrade) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


