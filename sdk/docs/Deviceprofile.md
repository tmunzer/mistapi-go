# Deviceprofile

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
**Name** | **string** |  | 
**NtpServers** | Pointer to **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PoePassthrough** | Pointer to **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | Pointer to [**map[string]GatewayPortConfig**](GatewayPortConfig.md) | Property key is the port(s) name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PwrConfig** | Pointer to [**ApPwrConfig**](ApPwrConfig.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SwitchConfig** | Pointer to [**ApSwitch**](ApSwitch.md) |  | [optional] 
**Type** | Pointer to [**GatewayTemplateType**](GatewayTemplateType.md) |  | [optional] [default to GATEWAYTEMPLATETYPE_STANDALONE]
**UplinkPortConfig** | Pointer to [**ApUplinkPortConfig**](ApUplinkPortConfig.md) |  | [optional] 
**UsbConfig** | Pointer to [**ApUsb**](ApUsb.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**BgpConfig** | Pointer to [**map[string]BgpConfig**](BgpConfig.md) |  | [optional] 
**DhcpdConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**DnsOverride** | Pointer to **bool** |  | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**ExtraRoutes** | Pointer to [**map[string]GatewayExtraRoute**](GatewayExtraRoute.md) |  | [optional] 
**GatewayMatching** | Pointer to [**GatewayMatching**](GatewayMatching.md) |  | [optional] 
**IdpProfiles** | Pointer to [**map[string]IdpProfile**](IdpProfile.md) | Property key is the profile name | [optional] 
**IpConfigs** | Pointer to [**map[string]GatewayTemplateIpConfig**](GatewayTemplateIpConfig.md) | Property key is the network name | [optional] 
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**NtpOverride** | Pointer to **bool** |  | [optional] [default to false]
**OobIpConfig** | Pointer to [**JunosOobIpConfigs**](JunosOobIpConfigs.md) |  | [optional] 
**PathPreferences** | Pointer to [**map[string]GatewayTemplatePathPreferences**](GatewayTemplatePathPreferences.md) | Property key is the path name | [optional] 
**RouterId** | Pointer to **string** | auto assigned if not set | [optional] 
**RoutingPolicies** | Pointer to [**map[string]RoutingPolicy**](RoutingPolicy.md) | Property key is the routing policy name | [optional] 
**ServicePolicies** | Pointer to [**[]ServicePolicy**](ServicePolicy.md) |  | [optional] 
**TunnelConfigs** | Pointer to [**map[string]TunnelConfigs**](TunnelConfigs.md) | Property key is the tunnel name | [optional] 
**TunnelProviderOptions** | Pointer to [**TunnelProviderOptions**](TunnelProviderOptions.md) |  | [optional] 

## Methods

### NewDeviceprofile

`func NewDeviceprofile(name string, ) *Deviceprofile`

NewDeviceprofile instantiates a new Deviceprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceprofileWithDefaults

`func NewDeviceprofileWithDefaults() *Deviceprofile`

NewDeviceprofileWithDefaults instantiates a new Deviceprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeroscout

`func (o *Deviceprofile) GetAeroscout() ApAeroscout`

GetAeroscout returns the Aeroscout field if non-nil, zero value otherwise.

### GetAeroscoutOk

`func (o *Deviceprofile) GetAeroscoutOk() (*ApAeroscout, bool)`

GetAeroscoutOk returns a tuple with the Aeroscout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeroscout

`func (o *Deviceprofile) SetAeroscout(v ApAeroscout)`

SetAeroscout sets Aeroscout field to given value.

### HasAeroscout

`func (o *Deviceprofile) HasAeroscout() bool`

HasAeroscout returns a boolean if a field has been set.

### GetBleConfig

`func (o *Deviceprofile) GetBleConfig() BleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *Deviceprofile) GetBleConfigOk() (*BleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *Deviceprofile) SetBleConfig(v BleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *Deviceprofile) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Deviceprofile) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Deviceprofile) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Deviceprofile) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Deviceprofile) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisableEth1

`func (o *Deviceprofile) GetDisableEth1() bool`

GetDisableEth1 returns the DisableEth1 field if non-nil, zero value otherwise.

### GetDisableEth1Ok

`func (o *Deviceprofile) GetDisableEth1Ok() (*bool, bool)`

GetDisableEth1Ok returns a tuple with the DisableEth1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth1

`func (o *Deviceprofile) SetDisableEth1(v bool)`

SetDisableEth1 sets DisableEth1 field to given value.

### HasDisableEth1

`func (o *Deviceprofile) HasDisableEth1() bool`

HasDisableEth1 returns a boolean if a field has been set.

### GetDisableEth2

`func (o *Deviceprofile) GetDisableEth2() bool`

GetDisableEth2 returns the DisableEth2 field if non-nil, zero value otherwise.

### GetDisableEth2Ok

`func (o *Deviceprofile) GetDisableEth2Ok() (*bool, bool)`

GetDisableEth2Ok returns a tuple with the DisableEth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth2

`func (o *Deviceprofile) SetDisableEth2(v bool)`

SetDisableEth2 sets DisableEth2 field to given value.

### HasDisableEth2

`func (o *Deviceprofile) HasDisableEth2() bool`

HasDisableEth2 returns a boolean if a field has been set.

### GetDisableEth3

`func (o *Deviceprofile) GetDisableEth3() bool`

GetDisableEth3 returns the DisableEth3 field if non-nil, zero value otherwise.

### GetDisableEth3Ok

`func (o *Deviceprofile) GetDisableEth3Ok() (*bool, bool)`

GetDisableEth3Ok returns a tuple with the DisableEth3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth3

`func (o *Deviceprofile) SetDisableEth3(v bool)`

SetDisableEth3 sets DisableEth3 field to given value.

### HasDisableEth3

`func (o *Deviceprofile) HasDisableEth3() bool`

HasDisableEth3 returns a boolean if a field has been set.

### GetDisableModule

`func (o *Deviceprofile) GetDisableModule() bool`

GetDisableModule returns the DisableModule field if non-nil, zero value otherwise.

### GetDisableModuleOk

`func (o *Deviceprofile) GetDisableModuleOk() (*bool, bool)`

GetDisableModuleOk returns a tuple with the DisableModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModule

`func (o *Deviceprofile) SetDisableModule(v bool)`

SetDisableModule sets DisableModule field to given value.

### HasDisableModule

`func (o *Deviceprofile) HasDisableModule() bool`

HasDisableModule returns a boolean if a field has been set.

### GetEslConfig

`func (o *Deviceprofile) GetEslConfig() ApEslConfig`

GetEslConfig returns the EslConfig field if non-nil, zero value otherwise.

### GetEslConfigOk

`func (o *Deviceprofile) GetEslConfigOk() (*ApEslConfig, bool)`

GetEslConfigOk returns a tuple with the EslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslConfig

`func (o *Deviceprofile) SetEslConfig(v ApEslConfig)`

SetEslConfig sets EslConfig field to given value.

### HasEslConfig

`func (o *Deviceprofile) HasEslConfig() bool`

HasEslConfig returns a boolean if a field has been set.

### GetForSite

`func (o *Deviceprofile) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Deviceprofile) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Deviceprofile) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Deviceprofile) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *Deviceprofile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deviceprofile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deviceprofile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deviceprofile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIotConfig

`func (o *Deviceprofile) GetIotConfig() ApIot`

GetIotConfig returns the IotConfig field if non-nil, zero value otherwise.

### GetIotConfigOk

`func (o *Deviceprofile) GetIotConfigOk() (*ApIot, bool)`

GetIotConfigOk returns a tuple with the IotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotConfig

`func (o *Deviceprofile) SetIotConfig(v ApIot)`

SetIotConfig sets IotConfig field to given value.

### HasIotConfig

`func (o *Deviceprofile) HasIotConfig() bool`

HasIotConfig returns a boolean if a field has been set.

### GetIpConfig

`func (o *Deviceprofile) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *Deviceprofile) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *Deviceprofile) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *Deviceprofile) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLed

`func (o *Deviceprofile) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *Deviceprofile) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *Deviceprofile) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *Deviceprofile) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetMesh

`func (o *Deviceprofile) GetMesh() ApMesh`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *Deviceprofile) GetMeshOk() (*ApMesh, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *Deviceprofile) SetMesh(v ApMesh)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *Deviceprofile) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Deviceprofile) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Deviceprofile) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Deviceprofile) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Deviceprofile) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Deviceprofile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deviceprofile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deviceprofile) SetName(v string)`

SetName sets Name field to given value.


### GetNtpServers

`func (o *Deviceprofile) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *Deviceprofile) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *Deviceprofile) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *Deviceprofile) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *Deviceprofile) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Deviceprofile) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Deviceprofile) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Deviceprofile) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPoePassthrough

`func (o *Deviceprofile) GetPoePassthrough() bool`

GetPoePassthrough returns the PoePassthrough field if non-nil, zero value otherwise.

### GetPoePassthroughOk

`func (o *Deviceprofile) GetPoePassthroughOk() (*bool, bool)`

GetPoePassthroughOk returns a tuple with the PoePassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePassthrough

`func (o *Deviceprofile) SetPoePassthrough(v bool)`

SetPoePassthrough sets PoePassthrough field to given value.

### HasPoePassthrough

`func (o *Deviceprofile) HasPoePassthrough() bool`

HasPoePassthrough returns a boolean if a field has been set.

### GetPortConfig

`func (o *Deviceprofile) GetPortConfig() map[string]GatewayPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *Deviceprofile) GetPortConfigOk() (*map[string]GatewayPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *Deviceprofile) SetPortConfig(v map[string]GatewayPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *Deviceprofile) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPwrConfig

`func (o *Deviceprofile) GetPwrConfig() ApPwrConfig`

GetPwrConfig returns the PwrConfig field if non-nil, zero value otherwise.

### GetPwrConfigOk

`func (o *Deviceprofile) GetPwrConfigOk() (*ApPwrConfig, bool)`

GetPwrConfigOk returns a tuple with the PwrConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrConfig

`func (o *Deviceprofile) SetPwrConfig(v ApPwrConfig)`

SetPwrConfig sets PwrConfig field to given value.

### HasPwrConfig

`func (o *Deviceprofile) HasPwrConfig() bool`

HasPwrConfig returns a boolean if a field has been set.

### GetRadioConfig

`func (o *Deviceprofile) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *Deviceprofile) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *Deviceprofile) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *Deviceprofile) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *Deviceprofile) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Deviceprofile) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Deviceprofile) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Deviceprofile) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSwitchConfig

`func (o *Deviceprofile) GetSwitchConfig() ApSwitch`

GetSwitchConfig returns the SwitchConfig field if non-nil, zero value otherwise.

### GetSwitchConfigOk

`func (o *Deviceprofile) GetSwitchConfigOk() (*ApSwitch, bool)`

GetSwitchConfigOk returns a tuple with the SwitchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchConfig

`func (o *Deviceprofile) SetSwitchConfig(v ApSwitch)`

SetSwitchConfig sets SwitchConfig field to given value.

### HasSwitchConfig

`func (o *Deviceprofile) HasSwitchConfig() bool`

HasSwitchConfig returns a boolean if a field has been set.

### GetType

`func (o *Deviceprofile) GetType() GatewayTemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deviceprofile) GetTypeOk() (*GatewayTemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deviceprofile) SetType(v GatewayTemplateType)`

SetType sets Type field to given value.

### HasType

`func (o *Deviceprofile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkPortConfig

`func (o *Deviceprofile) GetUplinkPortConfig() ApUplinkPortConfig`

GetUplinkPortConfig returns the UplinkPortConfig field if non-nil, zero value otherwise.

### GetUplinkPortConfigOk

`func (o *Deviceprofile) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool)`

GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortConfig

`func (o *Deviceprofile) SetUplinkPortConfig(v ApUplinkPortConfig)`

SetUplinkPortConfig sets UplinkPortConfig field to given value.

### HasUplinkPortConfig

`func (o *Deviceprofile) HasUplinkPortConfig() bool`

HasUplinkPortConfig returns a boolean if a field has been set.

### GetUsbConfig

`func (o *Deviceprofile) GetUsbConfig() ApUsb`

GetUsbConfig returns the UsbConfig field if non-nil, zero value otherwise.

### GetUsbConfigOk

`func (o *Deviceprofile) GetUsbConfigOk() (*ApUsb, bool)`

GetUsbConfigOk returns a tuple with the UsbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbConfig

`func (o *Deviceprofile) SetUsbConfig(v ApUsb)`

SetUsbConfig sets UsbConfig field to given value.

### HasUsbConfig

`func (o *Deviceprofile) HasUsbConfig() bool`

HasUsbConfig returns a boolean if a field has been set.

### GetVars

`func (o *Deviceprofile) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *Deviceprofile) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *Deviceprofile) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *Deviceprofile) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *Deviceprofile) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *Deviceprofile) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *Deviceprofile) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *Deviceprofile) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetBgpConfig

`func (o *Deviceprofile) GetBgpConfig() map[string]BgpConfig`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *Deviceprofile) GetBgpConfigOk() (*map[string]BgpConfig, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *Deviceprofile) SetBgpConfig(v map[string]BgpConfig)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *Deviceprofile) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *Deviceprofile) GetDhcpdConfig() DhcpdConfigs`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *Deviceprofile) GetDhcpdConfigOk() (*DhcpdConfigs, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *Deviceprofile) SetDhcpdConfig(v DhcpdConfigs)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *Deviceprofile) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetDnsOverride

`func (o *Deviceprofile) GetDnsOverride() bool`

GetDnsOverride returns the DnsOverride field if non-nil, zero value otherwise.

### GetDnsOverrideOk

`func (o *Deviceprofile) GetDnsOverrideOk() (*bool, bool)`

GetDnsOverrideOk returns a tuple with the DnsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOverride

`func (o *Deviceprofile) SetDnsOverride(v bool)`

SetDnsOverride sets DnsOverride field to given value.

### HasDnsOverride

`func (o *Deviceprofile) HasDnsOverride() bool`

HasDnsOverride returns a boolean if a field has been set.

### GetDnsServers

`func (o *Deviceprofile) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *Deviceprofile) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *Deviceprofile) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *Deviceprofile) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *Deviceprofile) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *Deviceprofile) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *Deviceprofile) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *Deviceprofile) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *Deviceprofile) GetExtraRoutes() map[string]GatewayExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *Deviceprofile) GetExtraRoutesOk() (*map[string]GatewayExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *Deviceprofile) SetExtraRoutes(v map[string]GatewayExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *Deviceprofile) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetGatewayMatching

`func (o *Deviceprofile) GetGatewayMatching() GatewayMatching`

GetGatewayMatching returns the GatewayMatching field if non-nil, zero value otherwise.

### GetGatewayMatchingOk

`func (o *Deviceprofile) GetGatewayMatchingOk() (*GatewayMatching, bool)`

GetGatewayMatchingOk returns a tuple with the GatewayMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMatching

`func (o *Deviceprofile) SetGatewayMatching(v GatewayMatching)`

SetGatewayMatching sets GatewayMatching field to given value.

### HasGatewayMatching

`func (o *Deviceprofile) HasGatewayMatching() bool`

HasGatewayMatching returns a boolean if a field has been set.

### GetIdpProfiles

`func (o *Deviceprofile) GetIdpProfiles() map[string]IdpProfile`

GetIdpProfiles returns the IdpProfiles field if non-nil, zero value otherwise.

### GetIdpProfilesOk

`func (o *Deviceprofile) GetIdpProfilesOk() (*map[string]IdpProfile, bool)`

GetIdpProfilesOk returns a tuple with the IdpProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProfiles

`func (o *Deviceprofile) SetIdpProfiles(v map[string]IdpProfile)`

SetIdpProfiles sets IdpProfiles field to given value.

### HasIdpProfiles

`func (o *Deviceprofile) HasIdpProfiles() bool`

HasIdpProfiles returns a boolean if a field has been set.

### GetIpConfigs

`func (o *Deviceprofile) GetIpConfigs() map[string]GatewayTemplateIpConfig`

GetIpConfigs returns the IpConfigs field if non-nil, zero value otherwise.

### GetIpConfigsOk

`func (o *Deviceprofile) GetIpConfigsOk() (*map[string]GatewayTemplateIpConfig, bool)`

GetIpConfigsOk returns a tuple with the IpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfigs

`func (o *Deviceprofile) SetIpConfigs(v map[string]GatewayTemplateIpConfig)`

SetIpConfigs sets IpConfigs field to given value.

### HasIpConfigs

`func (o *Deviceprofile) HasIpConfigs() bool`

HasIpConfigs returns a boolean if a field has been set.

### GetNetworks

`func (o *Deviceprofile) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Deviceprofile) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Deviceprofile) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *Deviceprofile) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNtpOverride

`func (o *Deviceprofile) GetNtpOverride() bool`

GetNtpOverride returns the NtpOverride field if non-nil, zero value otherwise.

### GetNtpOverrideOk

`func (o *Deviceprofile) GetNtpOverrideOk() (*bool, bool)`

GetNtpOverrideOk returns a tuple with the NtpOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpOverride

`func (o *Deviceprofile) SetNtpOverride(v bool)`

SetNtpOverride sets NtpOverride field to given value.

### HasNtpOverride

`func (o *Deviceprofile) HasNtpOverride() bool`

HasNtpOverride returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *Deviceprofile) GetOobIpConfig() JunosOobIpConfigs`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *Deviceprofile) GetOobIpConfigOk() (*JunosOobIpConfigs, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *Deviceprofile) SetOobIpConfig(v JunosOobIpConfigs)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *Deviceprofile) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetPathPreferences

`func (o *Deviceprofile) GetPathPreferences() map[string]GatewayTemplatePathPreferences`

GetPathPreferences returns the PathPreferences field if non-nil, zero value otherwise.

### GetPathPreferencesOk

`func (o *Deviceprofile) GetPathPreferencesOk() (*map[string]GatewayTemplatePathPreferences, bool)`

GetPathPreferencesOk returns a tuple with the PathPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPreferences

`func (o *Deviceprofile) SetPathPreferences(v map[string]GatewayTemplatePathPreferences)`

SetPathPreferences sets PathPreferences field to given value.

### HasPathPreferences

`func (o *Deviceprofile) HasPathPreferences() bool`

HasPathPreferences returns a boolean if a field has been set.

### GetRouterId

`func (o *Deviceprofile) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *Deviceprofile) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *Deviceprofile) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *Deviceprofile) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetRoutingPolicies

`func (o *Deviceprofile) GetRoutingPolicies() map[string]RoutingPolicy`

GetRoutingPolicies returns the RoutingPolicies field if non-nil, zero value otherwise.

### GetRoutingPoliciesOk

`func (o *Deviceprofile) GetRoutingPoliciesOk() (*map[string]RoutingPolicy, bool)`

GetRoutingPoliciesOk returns a tuple with the RoutingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicies

`func (o *Deviceprofile) SetRoutingPolicies(v map[string]RoutingPolicy)`

SetRoutingPolicies sets RoutingPolicies field to given value.

### HasRoutingPolicies

`func (o *Deviceprofile) HasRoutingPolicies() bool`

HasRoutingPolicies returns a boolean if a field has been set.

### GetServicePolicies

`func (o *Deviceprofile) GetServicePolicies() []ServicePolicy`

GetServicePolicies returns the ServicePolicies field if non-nil, zero value otherwise.

### GetServicePoliciesOk

`func (o *Deviceprofile) GetServicePoliciesOk() (*[]ServicePolicy, bool)`

GetServicePoliciesOk returns a tuple with the ServicePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicies

`func (o *Deviceprofile) SetServicePolicies(v []ServicePolicy)`

SetServicePolicies sets ServicePolicies field to given value.

### HasServicePolicies

`func (o *Deviceprofile) HasServicePolicies() bool`

HasServicePolicies returns a boolean if a field has been set.

### GetTunnelConfigs

`func (o *Deviceprofile) GetTunnelConfigs() map[string]TunnelConfigs`

GetTunnelConfigs returns the TunnelConfigs field if non-nil, zero value otherwise.

### GetTunnelConfigsOk

`func (o *Deviceprofile) GetTunnelConfigsOk() (*map[string]TunnelConfigs, bool)`

GetTunnelConfigsOk returns a tuple with the TunnelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConfigs

`func (o *Deviceprofile) SetTunnelConfigs(v map[string]TunnelConfigs)`

SetTunnelConfigs sets TunnelConfigs field to given value.

### HasTunnelConfigs

`func (o *Deviceprofile) HasTunnelConfigs() bool`

HasTunnelConfigs returns a boolean if a field has been set.

### GetTunnelProviderOptions

`func (o *Deviceprofile) GetTunnelProviderOptions() TunnelProviderOptions`

GetTunnelProviderOptions returns the TunnelProviderOptions field if non-nil, zero value otherwise.

### GetTunnelProviderOptionsOk

`func (o *Deviceprofile) GetTunnelProviderOptionsOk() (*TunnelProviderOptions, bool)`

GetTunnelProviderOptionsOk returns a tuple with the TunnelProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelProviderOptions

`func (o *Deviceprofile) SetTunnelProviderOptions(v TunnelProviderOptions)`

SetTunnelProviderOptions sets TunnelProviderOptions field to given value.

### HasTunnelProviderOptions

`func (o *Deviceprofile) HasTunnelProviderOptions() bool`

HasTunnelProviderOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


