# UpgradeBiosMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** | list of device id to upgrade bios | [optional] 
**Models** | Pointer to **[]string** | list of device model to upgrade bios | [optional] 
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed | [optional] [default to false]
**Version** | Pointer to **string** | specific bios version | [optional] 

## Methods

### NewUpgradeBiosMulti

`func NewUpgradeBiosMulti() *UpgradeBiosMulti`

NewUpgradeBiosMulti instantiates a new UpgradeBiosMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeBiosMultiWithDefaults

`func NewUpgradeBiosMultiWithDefaults() *UpgradeBiosMulti`

NewUpgradeBiosMultiWithDefaults instantiates a new UpgradeBiosMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *UpgradeBiosMulti) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UpgradeBiosMulti) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UpgradeBiosMulti) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *UpgradeBiosMulti) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetModels

`func (o *UpgradeBiosMulti) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UpgradeBiosMulti) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UpgradeBiosMulti) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *UpgradeBiosMulti) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetReboot

`func (o *UpgradeBiosMulti) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeBiosMulti) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeBiosMulti) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeBiosMulti) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeBiosMulti) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeBiosMulti) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeBiosMulti) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeBiosMulti) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


