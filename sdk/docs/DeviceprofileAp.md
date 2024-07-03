# DeviceprofileAp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aeroscout** | Pointer to [**ApAeroscout**](ApAeroscout.md) |  | [optional] 
**BleConfig** | Pointer to [**BleConfig**](BleConfig.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DisableEth1** | Pointer to **bool** | whether to disable eth1 port | [optional] [default to false]
**DisableEth2** | Pointer to **bool** | whether to disable eth2 port | [optional] [default to false]
**DisableEth3** | Pointer to **bool** | whether to disable eth3 port | [optional] [default to false]
**DisableModule** | Pointer to **bool** | whether to disable module port | [optional] [default to false]
**EslConfig** | Pointer to [**ApEslConfig**](ApEslConfig.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IotConfig** | Pointer to [**ApIot**](ApIot.md) |  | [optional] 
**IpConfig** | Pointer to [**ApIpConfig**](ApIpConfig.md) |  | [optional] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**Mesh** | Pointer to [**ApMesh**](ApMesh.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PoePassthrough** | Pointer to **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | Pointer to [**map[string]ApPortConfig**](ApPortConfig.md) | Property key is the interface(s) name (e.g. \&quot;eth1,eth2\&quot;) | [optional] 
**PwrConfig** | Pointer to [**ApPwrConfig**](ApPwrConfig.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SwitchConfig** | Pointer to [**ApSwitch**](ApSwitch.md) |  | [optional] 
**Type** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] [default to DEVICETYPE_AP]
**UplinkPortConfig** | Pointer to [**ApUplinkPortConfig**](ApUplinkPortConfig.md) |  | [optional] 
**UsbConfig** | Pointer to [**ApUsb**](ApUsb.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 

## Methods

### NewDeviceprofileAp

`func NewDeviceprofileAp() *DeviceprofileAp`

NewDeviceprofileAp instantiates a new DeviceprofileAp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceprofileApWithDefaults

`func NewDeviceprofileApWithDefaults() *DeviceprofileAp`

NewDeviceprofileApWithDefaults instantiates a new DeviceprofileAp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeroscout

`func (o *DeviceprofileAp) GetAeroscout() ApAeroscout`

GetAeroscout returns the Aeroscout field if non-nil, zero value otherwise.

### GetAeroscoutOk

`func (o *DeviceprofileAp) GetAeroscoutOk() (*ApAeroscout, bool)`

GetAeroscoutOk returns a tuple with the Aeroscout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeroscout

`func (o *DeviceprofileAp) SetAeroscout(v ApAeroscout)`

SetAeroscout sets Aeroscout field to given value.

### HasAeroscout

`func (o *DeviceprofileAp) HasAeroscout() bool`

HasAeroscout returns a boolean if a field has been set.

### GetBleConfig

`func (o *DeviceprofileAp) GetBleConfig() BleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *DeviceprofileAp) GetBleConfigOk() (*BleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *DeviceprofileAp) SetBleConfig(v BleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *DeviceprofileAp) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DeviceprofileAp) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DeviceprofileAp) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DeviceprofileAp) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DeviceprofileAp) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisableEth1

`func (o *DeviceprofileAp) GetDisableEth1() bool`

GetDisableEth1 returns the DisableEth1 field if non-nil, zero value otherwise.

### GetDisableEth1Ok

`func (o *DeviceprofileAp) GetDisableEth1Ok() (*bool, bool)`

GetDisableEth1Ok returns a tuple with the DisableEth1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth1

`func (o *DeviceprofileAp) SetDisableEth1(v bool)`

SetDisableEth1 sets DisableEth1 field to given value.

### HasDisableEth1

`func (o *DeviceprofileAp) HasDisableEth1() bool`

HasDisableEth1 returns a boolean if a field has been set.

### GetDisableEth2

`func (o *DeviceprofileAp) GetDisableEth2() bool`

GetDisableEth2 returns the DisableEth2 field if non-nil, zero value otherwise.

### GetDisableEth2Ok

`func (o *DeviceprofileAp) GetDisableEth2Ok() (*bool, bool)`

GetDisableEth2Ok returns a tuple with the DisableEth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth2

`func (o *DeviceprofileAp) SetDisableEth2(v bool)`

SetDisableEth2 sets DisableEth2 field to given value.

### HasDisableEth2

`func (o *DeviceprofileAp) HasDisableEth2() bool`

HasDisableEth2 returns a boolean if a field has been set.

### GetDisableEth3

`func (o *DeviceprofileAp) GetDisableEth3() bool`

GetDisableEth3 returns the DisableEth3 field if non-nil, zero value otherwise.

### GetDisableEth3Ok

`func (o *DeviceprofileAp) GetDisableEth3Ok() (*bool, bool)`

GetDisableEth3Ok returns a tuple with the DisableEth3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth3

`func (o *DeviceprofileAp) SetDisableEth3(v bool)`

SetDisableEth3 sets DisableEth3 field to given value.

### HasDisableEth3

`func (o *DeviceprofileAp) HasDisableEth3() bool`

HasDisableEth3 returns a boolean if a field has been set.

### GetDisableModule

`func (o *DeviceprofileAp) GetDisableModule() bool`

GetDisableModule returns the DisableModule field if non-nil, zero value otherwise.

### GetDisableModuleOk

`func (o *DeviceprofileAp) GetDisableModuleOk() (*bool, bool)`

GetDisableModuleOk returns a tuple with the DisableModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModule

`func (o *DeviceprofileAp) SetDisableModule(v bool)`

SetDisableModule sets DisableModule field to given value.

### HasDisableModule

`func (o *DeviceprofileAp) HasDisableModule() bool`

HasDisableModule returns a boolean if a field has been set.

### GetEslConfig

`func (o *DeviceprofileAp) GetEslConfig() ApEslConfig`

GetEslConfig returns the EslConfig field if non-nil, zero value otherwise.

### GetEslConfigOk

`func (o *DeviceprofileAp) GetEslConfigOk() (*ApEslConfig, bool)`

GetEslConfigOk returns a tuple with the EslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslConfig

`func (o *DeviceprofileAp) SetEslConfig(v ApEslConfig)`

SetEslConfig sets EslConfig field to given value.

### HasEslConfig

`func (o *DeviceprofileAp) HasEslConfig() bool`

HasEslConfig returns a boolean if a field has been set.

### GetForSite

`func (o *DeviceprofileAp) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *DeviceprofileAp) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *DeviceprofileAp) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *DeviceprofileAp) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *DeviceprofileAp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceprofileAp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceprofileAp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceprofileAp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIotConfig

`func (o *DeviceprofileAp) GetIotConfig() ApIot`

GetIotConfig returns the IotConfig field if non-nil, zero value otherwise.

### GetIotConfigOk

`func (o *DeviceprofileAp) GetIotConfigOk() (*ApIot, bool)`

GetIotConfigOk returns a tuple with the IotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotConfig

`func (o *DeviceprofileAp) SetIotConfig(v ApIot)`

SetIotConfig sets IotConfig field to given value.

### HasIotConfig

`func (o *DeviceprofileAp) HasIotConfig() bool`

HasIotConfig returns a boolean if a field has been set.

### GetIpConfig

`func (o *DeviceprofileAp) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *DeviceprofileAp) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *DeviceprofileAp) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *DeviceprofileAp) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLed

`func (o *DeviceprofileAp) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *DeviceprofileAp) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *DeviceprofileAp) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *DeviceprofileAp) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetMesh

`func (o *DeviceprofileAp) GetMesh() ApMesh`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *DeviceprofileAp) GetMeshOk() (*ApMesh, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *DeviceprofileAp) SetMesh(v ApMesh)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *DeviceprofileAp) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetModifiedTime

`func (o *DeviceprofileAp) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *DeviceprofileAp) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *DeviceprofileAp) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *DeviceprofileAp) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *DeviceprofileAp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceprofileAp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceprofileAp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceprofileAp) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceprofileAp) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceprofileAp) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNtpServers

`func (o *DeviceprofileAp) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *DeviceprofileAp) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *DeviceprofileAp) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *DeviceprofileAp) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *DeviceprofileAp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DeviceprofileAp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DeviceprofileAp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DeviceprofileAp) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPoePassthrough

`func (o *DeviceprofileAp) GetPoePassthrough() bool`

GetPoePassthrough returns the PoePassthrough field if non-nil, zero value otherwise.

### GetPoePassthroughOk

`func (o *DeviceprofileAp) GetPoePassthroughOk() (*bool, bool)`

GetPoePassthroughOk returns a tuple with the PoePassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePassthrough

`func (o *DeviceprofileAp) SetPoePassthrough(v bool)`

SetPoePassthrough sets PoePassthrough field to given value.

### HasPoePassthrough

`func (o *DeviceprofileAp) HasPoePassthrough() bool`

HasPoePassthrough returns a boolean if a field has been set.

### GetPortConfig

`func (o *DeviceprofileAp) GetPortConfig() map[string]ApPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *DeviceprofileAp) GetPortConfigOk() (*map[string]ApPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *DeviceprofileAp) SetPortConfig(v map[string]ApPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *DeviceprofileAp) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPwrConfig

`func (o *DeviceprofileAp) GetPwrConfig() ApPwrConfig`

GetPwrConfig returns the PwrConfig field if non-nil, zero value otherwise.

### GetPwrConfigOk

`func (o *DeviceprofileAp) GetPwrConfigOk() (*ApPwrConfig, bool)`

GetPwrConfigOk returns a tuple with the PwrConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrConfig

`func (o *DeviceprofileAp) SetPwrConfig(v ApPwrConfig)`

SetPwrConfig sets PwrConfig field to given value.

### HasPwrConfig

`func (o *DeviceprofileAp) HasPwrConfig() bool`

HasPwrConfig returns a boolean if a field has been set.

### GetRadioConfig

`func (o *DeviceprofileAp) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *DeviceprofileAp) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *DeviceprofileAp) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *DeviceprofileAp) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *DeviceprofileAp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceprofileAp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceprofileAp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DeviceprofileAp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSwitchConfig

`func (o *DeviceprofileAp) GetSwitchConfig() ApSwitch`

GetSwitchConfig returns the SwitchConfig field if non-nil, zero value otherwise.

### GetSwitchConfigOk

`func (o *DeviceprofileAp) GetSwitchConfigOk() (*ApSwitch, bool)`

GetSwitchConfigOk returns a tuple with the SwitchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchConfig

`func (o *DeviceprofileAp) SetSwitchConfig(v ApSwitch)`

SetSwitchConfig sets SwitchConfig field to given value.

### HasSwitchConfig

`func (o *DeviceprofileAp) HasSwitchConfig() bool`

HasSwitchConfig returns a boolean if a field has been set.

### GetType

`func (o *DeviceprofileAp) GetType() DeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceprofileAp) GetTypeOk() (*DeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceprofileAp) SetType(v DeviceType)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceprofileAp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkPortConfig

`func (o *DeviceprofileAp) GetUplinkPortConfig() ApUplinkPortConfig`

GetUplinkPortConfig returns the UplinkPortConfig field if non-nil, zero value otherwise.

### GetUplinkPortConfigOk

`func (o *DeviceprofileAp) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool)`

GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortConfig

`func (o *DeviceprofileAp) SetUplinkPortConfig(v ApUplinkPortConfig)`

SetUplinkPortConfig sets UplinkPortConfig field to given value.

### HasUplinkPortConfig

`func (o *DeviceprofileAp) HasUplinkPortConfig() bool`

HasUplinkPortConfig returns a boolean if a field has been set.

### GetUsbConfig

`func (o *DeviceprofileAp) GetUsbConfig() ApUsb`

GetUsbConfig returns the UsbConfig field if non-nil, zero value otherwise.

### GetUsbConfigOk

`func (o *DeviceprofileAp) GetUsbConfigOk() (*ApUsb, bool)`

GetUsbConfigOk returns a tuple with the UsbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbConfig

`func (o *DeviceprofileAp) SetUsbConfig(v ApUsb)`

SetUsbConfig sets UsbConfig field to given value.

### HasUsbConfig

`func (o *DeviceprofileAp) HasUsbConfig() bool`

HasUsbConfig returns a boolean if a field has been set.

### GetVars

`func (o *DeviceprofileAp) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *DeviceprofileAp) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *DeviceprofileAp) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *DeviceprofileAp) HasVars() bool`

HasVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


