# ModuleStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupVersion** | Pointer to **string** |  | [optional] 
**BiosVersion** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ModuleStatErrorsItems**](ModuleStatErrorsItems.md) | used to report all error states the device node is running into. An error should always have &#x60;type&#x60; and &#x60;since&#x60; fields, and could have some other fields specific to that type. | [optional] 
**Fans** | Pointer to [**[]ModuleStatFansItems**](ModuleStatFansItems.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Poe** | Pointer to [**ModuleStatPoe**](ModuleStatPoe.md) |  | [optional] 
**Psus** | Pointer to [**[]ModuleStatPsusItems**](ModuleStatPsusItems.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Temperatures** | Pointer to [**[]ModuleStatTemperaturesItems**](ModuleStatTemperaturesItems.md) |  | [optional] 
**VcLinks** | Pointer to [**[]ModuleStatVcLinksItems**](ModuleStatVcLinksItems.md) |  | [optional] 
**VcMode** | Pointer to **string** |  | [optional] 
**VcRole** | Pointer to **string** | master / backup / linecard | [optional] 
**VcState** | Pointer to **string** |  | [optional] 

## Methods

### NewModuleStat

`func NewModuleStat() *ModuleStat`

NewModuleStat instantiates a new ModuleStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleStatWithDefaults

`func NewModuleStatWithDefaults() *ModuleStat`

NewModuleStatWithDefaults instantiates a new ModuleStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupVersion

`func (o *ModuleStat) GetBackupVersion() string`

GetBackupVersion returns the BackupVersion field if non-nil, zero value otherwise.

### GetBackupVersionOk

`func (o *ModuleStat) GetBackupVersionOk() (*string, bool)`

GetBackupVersionOk returns a tuple with the BackupVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupVersion

`func (o *ModuleStat) SetBackupVersion(v string)`

SetBackupVersion sets BackupVersion field to given value.

### HasBackupVersion

`func (o *ModuleStat) HasBackupVersion() bool`

HasBackupVersion returns a boolean if a field has been set.

### GetBiosVersion

`func (o *ModuleStat) GetBiosVersion() string`

GetBiosVersion returns the BiosVersion field if non-nil, zero value otherwise.

### GetBiosVersionOk

`func (o *ModuleStat) GetBiosVersionOk() (*string, bool)`

GetBiosVersionOk returns a tuple with the BiosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVersion

`func (o *ModuleStat) SetBiosVersion(v string)`

SetBiosVersion sets BiosVersion field to given value.

### HasBiosVersion

`func (o *ModuleStat) HasBiosVersion() bool`

HasBiosVersion returns a boolean if a field has been set.

### GetErrors

`func (o *ModuleStat) GetErrors() []ModuleStatErrorsItems`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModuleStat) GetErrorsOk() (*[]ModuleStatErrorsItems, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModuleStat) SetErrors(v []ModuleStatErrorsItems)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ModuleStat) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFans

`func (o *ModuleStat) GetFans() []ModuleStatFansItems`

GetFans returns the Fans field if non-nil, zero value otherwise.

### GetFansOk

`func (o *ModuleStat) GetFansOk() (*[]ModuleStatFansItems, bool)`

GetFansOk returns a tuple with the Fans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFans

`func (o *ModuleStat) SetFans(v []ModuleStatFansItems)`

SetFans sets Fans field to given value.

### HasFans

`func (o *ModuleStat) HasFans() bool`

HasFans returns a boolean if a field has been set.

### GetModel

`func (o *ModuleStat) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModuleStat) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModuleStat) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModuleStat) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPoe

`func (o *ModuleStat) GetPoe() ModuleStatPoe`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *ModuleStat) GetPoeOk() (*ModuleStatPoe, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *ModuleStat) SetPoe(v ModuleStatPoe)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *ModuleStat) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetPsus

`func (o *ModuleStat) GetPsus() []ModuleStatPsusItems`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *ModuleStat) GetPsusOk() (*[]ModuleStatPsusItems, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *ModuleStat) SetPsus(v []ModuleStatPsusItems)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *ModuleStat) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### GetSerial

`func (o *ModuleStat) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ModuleStat) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ModuleStat) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ModuleStat) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTemperatures

`func (o *ModuleStat) GetTemperatures() []ModuleStatTemperaturesItems`

GetTemperatures returns the Temperatures field if non-nil, zero value otherwise.

### GetTemperaturesOk

`func (o *ModuleStat) GetTemperaturesOk() (*[]ModuleStatTemperaturesItems, bool)`

GetTemperaturesOk returns a tuple with the Temperatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatures

`func (o *ModuleStat) SetTemperatures(v []ModuleStatTemperaturesItems)`

SetTemperatures sets Temperatures field to given value.

### HasTemperatures

`func (o *ModuleStat) HasTemperatures() bool`

HasTemperatures returns a boolean if a field has been set.

### GetVcLinks

`func (o *ModuleStat) GetVcLinks() []ModuleStatVcLinksItems`

GetVcLinks returns the VcLinks field if non-nil, zero value otherwise.

### GetVcLinksOk

`func (o *ModuleStat) GetVcLinksOk() (*[]ModuleStatVcLinksItems, bool)`

GetVcLinksOk returns a tuple with the VcLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcLinks

`func (o *ModuleStat) SetVcLinks(v []ModuleStatVcLinksItems)`

SetVcLinks sets VcLinks field to given value.

### HasVcLinks

`func (o *ModuleStat) HasVcLinks() bool`

HasVcLinks returns a boolean if a field has been set.

### GetVcMode

`func (o *ModuleStat) GetVcMode() string`

GetVcMode returns the VcMode field if non-nil, zero value otherwise.

### GetVcModeOk

`func (o *ModuleStat) GetVcModeOk() (*string, bool)`

GetVcModeOk returns a tuple with the VcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcMode

`func (o *ModuleStat) SetVcMode(v string)`

SetVcMode sets VcMode field to given value.

### HasVcMode

`func (o *ModuleStat) HasVcMode() bool`

HasVcMode returns a boolean if a field has been set.

### GetVcRole

`func (o *ModuleStat) GetVcRole() string`

GetVcRole returns the VcRole field if non-nil, zero value otherwise.

### GetVcRoleOk

`func (o *ModuleStat) GetVcRoleOk() (*string, bool)`

GetVcRoleOk returns a tuple with the VcRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcRole

`func (o *ModuleStat) SetVcRole(v string)`

SetVcRole sets VcRole field to given value.

### HasVcRole

`func (o *ModuleStat) HasVcRole() bool`

HasVcRole returns a boolean if a field has been set.

### GetVcState

`func (o *ModuleStat) GetVcState() string`

GetVcState returns the VcState field if non-nil, zero value otherwise.

### GetVcStateOk

`func (o *ModuleStat) GetVcStateOk() (*string, bool)`

GetVcStateOk returns a tuple with the VcState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcState

`func (o *ModuleStat) SetVcState(v string)`

SetVcState sets VcState field to given value.

### HasVcState

`func (o *ModuleStat) HasVcState() bool`

HasVcState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


