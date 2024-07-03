# UpgradeFpgaMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** | list of device id to upgrade bios | [optional] 
**Models** | Pointer to **[]string** | list of device model to upgrade bios | [optional] 
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed | [optional] [default to false]
**Version** | Pointer to **string** | specific FPGA version | [optional] 

## Methods

### NewUpgradeFpgaMulti

`func NewUpgradeFpgaMulti() *UpgradeFpgaMulti`

NewUpgradeFpgaMulti instantiates a new UpgradeFpgaMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeFpgaMultiWithDefaults

`func NewUpgradeFpgaMultiWithDefaults() *UpgradeFpgaMulti`

NewUpgradeFpgaMultiWithDefaults instantiates a new UpgradeFpgaMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *UpgradeFpgaMulti) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UpgradeFpgaMulti) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UpgradeFpgaMulti) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *UpgradeFpgaMulti) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetModels

`func (o *UpgradeFpgaMulti) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *UpgradeFpgaMulti) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *UpgradeFpgaMulti) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *UpgradeFpgaMulti) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetReboot

`func (o *UpgradeFpgaMulti) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeFpgaMulti) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeFpgaMulti) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeFpgaMulti) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeFpgaMulti) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeFpgaMulti) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeFpgaMulti) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeFpgaMulti) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


