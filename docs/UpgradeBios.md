# UpgradeBios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to **bool** | Reboot device immediately after upgrade is completed | [optional] [default to false]
**Version** | Pointer to **string** | specific bios version | [optional] 

## Methods

### NewUpgradeBios

`func NewUpgradeBios() *UpgradeBios`

NewUpgradeBios instantiates a new UpgradeBios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeBiosWithDefaults

`func NewUpgradeBiosWithDefaults() *UpgradeBios`

NewUpgradeBiosWithDefaults instantiates a new UpgradeBios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *UpgradeBios) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpgradeBios) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpgradeBios) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpgradeBios) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeBios) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeBios) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeBios) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeBios) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


