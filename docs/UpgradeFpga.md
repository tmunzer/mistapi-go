# UpgradeFpga

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed | [optional] [default to false]
**Version** | Pointer to **string** | specific fpga version | [optional] 

## Methods

### NewUpgradeFpga

`func NewUpgradeFpga() *UpgradeFpga`

NewUpgradeFpga instantiates a new UpgradeFpga object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeFpgaWithDefaults

`func NewUpgradeFpgaWithDefaults() *UpgradeFpga`

NewUpgradeFpgaWithDefaults instantiates a new UpgradeFpga object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *UpgradeFpga) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeFpga) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeFpga) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeFpga) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeFpga) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeFpga) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeFpga) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeFpga) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


