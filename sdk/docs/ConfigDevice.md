# ConfigDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aeroscout** | Pointer to [**ApAeroscout**](ApAeroscout.md) |  | [optional] 
**BleConfig** | Pointer to [**BleConfig**](BleConfig.md) |  | [optional] 
**Centrak** | Pointer to [**ApCentrak**](ApCentrak.md) |  | [optional] 
**ClientBridge** | Pointer to [**ApClientBridge**](ApClientBridge.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**DisableEth1** | Pointer to **bool** | whether to disable eth1 port | [optional] [default to false]
**DisableEth2** | Pointer to **bool** | whether to disable eth2 port | [optional] [default to false]
**DisableEth3** | Pointer to **bool** | whether to disable eth3 port | [optional] [default to false]
**DisableModule** | Pointer to **bool** | whether to disable module port | [optional] [default to false]
**EslConfig** | Pointer to [**ApEslConfig**](ApEslConfig.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Height** | Pointer to **float32** | height, in meters, optional | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Image1Url** | Pointer to **string** |  | [optional] 
**Image2Url** | Pointer to **string** |  | [optional] 
**Image3Url** | Pointer to **string** |  | [optional] 
**IotConfig** | Pointer to [**ApIot**](ApIot.md) |  | [optional] 
**IpConfig** | Pointer to [**map[string]GatewayTemplateIpConfig**](GatewayTemplateIpConfig.md) | Property key is the network name | [optional] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**Locked** | Pointer to **bool** | whether this map is considered locked down | [optional] 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**Mesh** | Pointer to [**ApMesh**](ApMesh.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Orientation** | Pointer to **int32** | orientation, 0-359, in degrees, up is 0, right is 90. | [optional] 
**PoePassthrough** | Pointer to **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | Pointer to [**map[string]GatewayPortConfig**](GatewayPortConfig.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PwrConfig** | Pointer to [**ApPwrConfig**](ApPwrConfig.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**UplinkPortConfig** | Pointer to [**ApUplinkPortConfig**](ApUplinkPortConfig.md) |  | [optional] 
**UsbConfig** | Pointer to [**ApUsb**](ApUsb.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 
**AclPolicies** | Pointer to [**[]AclPolicy**](AclPolicy.md) |  | [optional] 
**AclTags** | Pointer to [**map[string]AclTag**](AclTag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**DhcpConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**DhcpSnooping** | Pointer to [**DhcpSnooping**](DhcpSnooping.md) |  | [optional] 
**DisableAutoConfig** | Pointer to **bool** | for a claimed switch, we control the configs by default. This option (disables the behavior) | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**EvpnConfig** | Pointer to [**EvpnConfig**](EvpnConfig.md) |  | [optional] 
**ExtraRoutes** | Pointer to [**map[string]GatewayExtraRoute**](GatewayExtraRoute.md) |  | [optional] 
**ExtraRoutes6** | Pointer to [**map[string]ExtraRoute6Properties**](ExtraRoute6Properties.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Networks** | Pointer to [**map[string]GatewayNetwork**](GatewayNetwork.md) | Property key is the network name or a CIDR | [optional] 
**OobIpConfig** | Pointer to [**JunosOobIpConfigs**](JunosOobIpConfigs.md) |  | [optional] 
**OspfConfig** | Pointer to [**OspfConfig**](OspfConfig.md) |  | [optional] 
**OtherIpConfigs** | Pointer to [**map[string]JunosOtherIpConfigs**](JunosOtherIpConfigs.md) | Property key is the network name | [optional] 
**PortMirroring** | Pointer to [**GatewayPortMirroring**](GatewayPortMirroring.md) |  | [optional] 
**PortUsages** | Pointer to [**map[string]SwitchPortUsage**](SwitchPortUsage.md) |  | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**Role** | Pointer to [**SwitchRole**](SwitchRole.md) |  | [optional] [default to SWITCHROLE_ACCESS]
**RouterId** | Pointer to **string** | used for OSPF / BGP / EVPN | [optional] 
**StpConfig** | Pointer to [**SwitchStpConfig**](SwitchStpConfig.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchSwitchMgmt**](SwitchSwitchMgmt.md) |  | [optional] 
**UseRouterIdAsSourceIp** | Pointer to **bool** | whether to use it for snmp / syslog / tacplus / radius | [optional] [default to false]
**VirtualChassis** | Pointer to [**SwitchVirtualChassis**](SwitchVirtualChassis.md) |  | [optional] 
**VrfConfig** | Pointer to [**VrfConfig**](VrfConfig.md) |  | [optional] 
**VrrpConfig** | Pointer to [**VrrpConfig**](VrrpConfig.md) |  | [optional] 
**DhcpdConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewConfigDevice

`func NewConfigDevice() *ConfigDevice`

NewConfigDevice instantiates a new ConfigDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDeviceWithDefaults

`func NewConfigDeviceWithDefaults() *ConfigDevice`

NewConfigDeviceWithDefaults instantiates a new ConfigDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeroscout

`func (o *ConfigDevice) GetAeroscout() ApAeroscout`

GetAeroscout returns the Aeroscout field if non-nil, zero value otherwise.

### GetAeroscoutOk

`func (o *ConfigDevice) GetAeroscoutOk() (*ApAeroscout, bool)`

GetAeroscoutOk returns a tuple with the Aeroscout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeroscout

`func (o *ConfigDevice) SetAeroscout(v ApAeroscout)`

SetAeroscout sets Aeroscout field to given value.

### HasAeroscout

`func (o *ConfigDevice) HasAeroscout() bool`

HasAeroscout returns a boolean if a field has been set.

### GetBleConfig

`func (o *ConfigDevice) GetBleConfig() BleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *ConfigDevice) GetBleConfigOk() (*BleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *ConfigDevice) SetBleConfig(v BleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *ConfigDevice) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetCentrak

`func (o *ConfigDevice) GetCentrak() ApCentrak`

GetCentrak returns the Centrak field if non-nil, zero value otherwise.

### GetCentrakOk

`func (o *ConfigDevice) GetCentrakOk() (*ApCentrak, bool)`

GetCentrakOk returns a tuple with the Centrak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrak

`func (o *ConfigDevice) SetCentrak(v ApCentrak)`

SetCentrak sets Centrak field to given value.

### HasCentrak

`func (o *ConfigDevice) HasCentrak() bool`

HasCentrak returns a boolean if a field has been set.

### GetClientBridge

`func (o *ConfigDevice) GetClientBridge() ApClientBridge`

GetClientBridge returns the ClientBridge field if non-nil, zero value otherwise.

### GetClientBridgeOk

`func (o *ConfigDevice) GetClientBridgeOk() (*ApClientBridge, bool)`

GetClientBridgeOk returns a tuple with the ClientBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBridge

`func (o *ConfigDevice) SetClientBridge(v ApClientBridge)`

SetClientBridge sets ClientBridge field to given value.

### HasClientBridge

`func (o *ConfigDevice) HasClientBridge() bool`

HasClientBridge returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ConfigDevice) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ConfigDevice) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ConfigDevice) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ConfigDevice) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *ConfigDevice) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *ConfigDevice) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *ConfigDevice) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *ConfigDevice) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDisableEth1

`func (o *ConfigDevice) GetDisableEth1() bool`

GetDisableEth1 returns the DisableEth1 field if non-nil, zero value otherwise.

### GetDisableEth1Ok

`func (o *ConfigDevice) GetDisableEth1Ok() (*bool, bool)`

GetDisableEth1Ok returns a tuple with the DisableEth1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth1

`func (o *ConfigDevice) SetDisableEth1(v bool)`

SetDisableEth1 sets DisableEth1 field to given value.

### HasDisableEth1

`func (o *ConfigDevice) HasDisableEth1() bool`

HasDisableEth1 returns a boolean if a field has been set.

### GetDisableEth2

`func (o *ConfigDevice) GetDisableEth2() bool`

GetDisableEth2 returns the DisableEth2 field if non-nil, zero value otherwise.

### GetDisableEth2Ok

`func (o *ConfigDevice) GetDisableEth2Ok() (*bool, bool)`

GetDisableEth2Ok returns a tuple with the DisableEth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth2

`func (o *ConfigDevice) SetDisableEth2(v bool)`

SetDisableEth2 sets DisableEth2 field to given value.

### HasDisableEth2

`func (o *ConfigDevice) HasDisableEth2() bool`

HasDisableEth2 returns a boolean if a field has been set.

### GetDisableEth3

`func (o *ConfigDevice) GetDisableEth3() bool`

GetDisableEth3 returns the DisableEth3 field if non-nil, zero value otherwise.

### GetDisableEth3Ok

`func (o *ConfigDevice) GetDisableEth3Ok() (*bool, bool)`

GetDisableEth3Ok returns a tuple with the DisableEth3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth3

`func (o *ConfigDevice) SetDisableEth3(v bool)`

SetDisableEth3 sets DisableEth3 field to given value.

### HasDisableEth3

`func (o *ConfigDevice) HasDisableEth3() bool`

HasDisableEth3 returns a boolean if a field has been set.

### GetDisableModule

`func (o *ConfigDevice) GetDisableModule() bool`

GetDisableModule returns the DisableModule field if non-nil, zero value otherwise.

### GetDisableModuleOk

`func (o *ConfigDevice) GetDisableModuleOk() (*bool, bool)`

GetDisableModuleOk returns a tuple with the DisableModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModule

`func (o *ConfigDevice) SetDisableModule(v bool)`

SetDisableModule sets DisableModule field to given value.

### HasDisableModule

`func (o *ConfigDevice) HasDisableModule() bool`

HasDisableModule returns a boolean if a field has been set.

### GetEslConfig

`func (o *ConfigDevice) GetEslConfig() ApEslConfig`

GetEslConfig returns the EslConfig field if non-nil, zero value otherwise.

### GetEslConfigOk

`func (o *ConfigDevice) GetEslConfigOk() (*ApEslConfig, bool)`

GetEslConfigOk returns a tuple with the EslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslConfig

`func (o *ConfigDevice) SetEslConfig(v ApEslConfig)`

SetEslConfig sets EslConfig field to given value.

### HasEslConfig

`func (o *ConfigDevice) HasEslConfig() bool`

HasEslConfig returns a boolean if a field has been set.

### GetForSite

`func (o *ConfigDevice) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ConfigDevice) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ConfigDevice) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ConfigDevice) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeight

`func (o *ConfigDevice) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ConfigDevice) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ConfigDevice) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ConfigDevice) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ConfigDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *ConfigDevice) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *ConfigDevice) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *ConfigDevice) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *ConfigDevice) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### GetImage2Url

`func (o *ConfigDevice) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *ConfigDevice) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *ConfigDevice) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *ConfigDevice) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### GetImage3Url

`func (o *ConfigDevice) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *ConfigDevice) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *ConfigDevice) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *ConfigDevice) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### GetIotConfig

`func (o *ConfigDevice) GetIotConfig() ApIot`

GetIotConfig returns the IotConfig field if non-nil, zero value otherwise.

### GetIotConfigOk

`func (o *ConfigDevice) GetIotConfigOk() (*ApIot, bool)`

GetIotConfigOk returns a tuple with the IotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotConfig

`func (o *ConfigDevice) SetIotConfig(v ApIot)`

SetIotConfig sets IotConfig field to given value.

### HasIotConfig

`func (o *ConfigDevice) HasIotConfig() bool`

HasIotConfig returns a boolean if a field has been set.

### GetIpConfig

`func (o *ConfigDevice) GetIpConfig() map[string]GatewayTemplateIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ConfigDevice) GetIpConfigOk() (*map[string]GatewayTemplateIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ConfigDevice) SetIpConfig(v map[string]GatewayTemplateIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ConfigDevice) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLed

`func (o *ConfigDevice) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *ConfigDevice) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *ConfigDevice) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *ConfigDevice) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLocked

`func (o *ConfigDevice) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ConfigDevice) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ConfigDevice) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ConfigDevice) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMapId

`func (o *ConfigDevice) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ConfigDevice) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ConfigDevice) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *ConfigDevice) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMesh

`func (o *ConfigDevice) GetMesh() ApMesh`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *ConfigDevice) GetMeshOk() (*ApMesh, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *ConfigDevice) SetMesh(v ApMesh)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *ConfigDevice) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetModifiedTime

`func (o *ConfigDevice) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *ConfigDevice) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *ConfigDevice) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *ConfigDevice) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *ConfigDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *ConfigDevice) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ConfigDevice) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ConfigDevice) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ConfigDevice) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpServers

`func (o *ConfigDevice) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ConfigDevice) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ConfigDevice) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *ConfigDevice) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *ConfigDevice) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConfigDevice) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConfigDevice) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConfigDevice) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrientation

`func (o *ConfigDevice) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ConfigDevice) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ConfigDevice) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ConfigDevice) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPoePassthrough

`func (o *ConfigDevice) GetPoePassthrough() bool`

GetPoePassthrough returns the PoePassthrough field if non-nil, zero value otherwise.

### GetPoePassthroughOk

`func (o *ConfigDevice) GetPoePassthroughOk() (*bool, bool)`

GetPoePassthroughOk returns a tuple with the PoePassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePassthrough

`func (o *ConfigDevice) SetPoePassthrough(v bool)`

SetPoePassthrough sets PoePassthrough field to given value.

### HasPoePassthrough

`func (o *ConfigDevice) HasPoePassthrough() bool`

HasPoePassthrough returns a boolean if a field has been set.

### GetPortConfig

`func (o *ConfigDevice) GetPortConfig() map[string]GatewayPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *ConfigDevice) GetPortConfigOk() (*map[string]GatewayPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *ConfigDevice) SetPortConfig(v map[string]GatewayPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *ConfigDevice) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPwrConfig

`func (o *ConfigDevice) GetPwrConfig() ApPwrConfig`

GetPwrConfig returns the PwrConfig field if non-nil, zero value otherwise.

### GetPwrConfigOk

`func (o *ConfigDevice) GetPwrConfigOk() (*ApPwrConfig, bool)`

GetPwrConfigOk returns a tuple with the PwrConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrConfig

`func (o *ConfigDevice) SetPwrConfig(v ApPwrConfig)`

SetPwrConfig sets PwrConfig field to given value.

### HasPwrConfig

`func (o *ConfigDevice) HasPwrConfig() bool`

HasPwrConfig returns a boolean if a field has been set.

### GetRadioConfig

`func (o *ConfigDevice) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *ConfigDevice) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *ConfigDevice) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *ConfigDevice) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *ConfigDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConfigDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConfigDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ConfigDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUplinkPortConfig

`func (o *ConfigDevice) GetUplinkPortConfig() ApUplinkPortConfig`

GetUplinkPortConfig returns the UplinkPortConfig field if non-nil, zero value otherwise.

### GetUplinkPortConfigOk

`func (o *ConfigDevice) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool)`

GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortConfig

`func (o *ConfigDevice) SetUplinkPortConfig(v ApUplinkPortConfig)`

SetUplinkPortConfig sets UplinkPortConfig field to given value.

### HasUplinkPortConfig

`func (o *ConfigDevice) HasUplinkPortConfig() bool`

HasUplinkPortConfig returns a boolean if a field has been set.

### GetUsbConfig

`func (o *ConfigDevice) GetUsbConfig() ApUsb`

GetUsbConfig returns the UsbConfig field if non-nil, zero value otherwise.

### GetUsbConfigOk

`func (o *ConfigDevice) GetUsbConfigOk() (*ApUsb, bool)`

GetUsbConfigOk returns a tuple with the UsbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbConfig

`func (o *ConfigDevice) SetUsbConfig(v ApUsb)`

SetUsbConfig sets UsbConfig field to given value.

### HasUsbConfig

`func (o *ConfigDevice) HasUsbConfig() bool`

HasUsbConfig returns a boolean if a field has been set.

### GetVars

`func (o *ConfigDevice) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *ConfigDevice) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *ConfigDevice) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *ConfigDevice) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetX

`func (o *ConfigDevice) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ConfigDevice) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ConfigDevice) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ConfigDevice) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ConfigDevice) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ConfigDevice) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ConfigDevice) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ConfigDevice) HasY() bool`

HasY returns a boolean if a field has been set.

### GetAclPolicies

`func (o *ConfigDevice) GetAclPolicies() []AclPolicy`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *ConfigDevice) GetAclPoliciesOk() (*[]AclPolicy, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *ConfigDevice) SetAclPolicies(v []AclPolicy)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *ConfigDevice) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *ConfigDevice) GetAclTags() map[string]AclTag`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *ConfigDevice) GetAclTagsOk() (*map[string]AclTag, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *ConfigDevice) SetAclTags(v map[string]AclTag)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *ConfigDevice) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *ConfigDevice) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *ConfigDevice) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *ConfigDevice) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *ConfigDevice) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *ConfigDevice) GetDhcpConfig() DhcpdConfigs`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *ConfigDevice) GetDhcpConfigOk() (*DhcpdConfigs, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *ConfigDevice) SetDhcpConfig(v DhcpdConfigs)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *ConfigDevice) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *ConfigDevice) GetDhcpSnooping() DhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *ConfigDevice) GetDhcpSnoopingOk() (*DhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *ConfigDevice) SetDhcpSnooping(v DhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *ConfigDevice) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDisableAutoConfig

`func (o *ConfigDevice) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *ConfigDevice) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *ConfigDevice) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *ConfigDevice) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetDnsServers

`func (o *ConfigDevice) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ConfigDevice) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ConfigDevice) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ConfigDevice) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *ConfigDevice) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *ConfigDevice) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *ConfigDevice) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *ConfigDevice) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEvpnConfig

`func (o *ConfigDevice) GetEvpnConfig() EvpnConfig`

GetEvpnConfig returns the EvpnConfig field if non-nil, zero value otherwise.

### GetEvpnConfigOk

`func (o *ConfigDevice) GetEvpnConfigOk() (*EvpnConfig, bool)`

GetEvpnConfigOk returns a tuple with the EvpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnConfig

`func (o *ConfigDevice) SetEvpnConfig(v EvpnConfig)`

SetEvpnConfig sets EvpnConfig field to given value.

### HasEvpnConfig

`func (o *ConfigDevice) HasEvpnConfig() bool`

HasEvpnConfig returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *ConfigDevice) GetExtraRoutes() map[string]GatewayExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *ConfigDevice) GetExtraRoutesOk() (*map[string]GatewayExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *ConfigDevice) SetExtraRoutes(v map[string]GatewayExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *ConfigDevice) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *ConfigDevice) GetExtraRoutes6() map[string]ExtraRoute6Properties`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *ConfigDevice) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *ConfigDevice) SetExtraRoutes6(v map[string]ExtraRoute6Properties)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *ConfigDevice) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetManaged

`func (o *ConfigDevice) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ConfigDevice) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ConfigDevice) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ConfigDevice) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNetworks

`func (o *ConfigDevice) GetNetworks() map[string]GatewayNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ConfigDevice) GetNetworksOk() (*map[string]GatewayNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ConfigDevice) SetNetworks(v map[string]GatewayNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ConfigDevice) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *ConfigDevice) GetOobIpConfig() JunosOobIpConfigs`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *ConfigDevice) GetOobIpConfigOk() (*JunosOobIpConfigs, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *ConfigDevice) SetOobIpConfig(v JunosOobIpConfigs)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *ConfigDevice) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOspfConfig

`func (o *ConfigDevice) GetOspfConfig() OspfConfig`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *ConfigDevice) GetOspfConfigOk() (*OspfConfig, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *ConfigDevice) SetOspfConfig(v OspfConfig)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *ConfigDevice) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.

### GetOtherIpConfigs

`func (o *ConfigDevice) GetOtherIpConfigs() map[string]JunosOtherIpConfigs`

GetOtherIpConfigs returns the OtherIpConfigs field if non-nil, zero value otherwise.

### GetOtherIpConfigsOk

`func (o *ConfigDevice) GetOtherIpConfigsOk() (*map[string]JunosOtherIpConfigs, bool)`

GetOtherIpConfigsOk returns a tuple with the OtherIpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIpConfigs

`func (o *ConfigDevice) SetOtherIpConfigs(v map[string]JunosOtherIpConfigs)`

SetOtherIpConfigs sets OtherIpConfigs field to given value.

### HasOtherIpConfigs

`func (o *ConfigDevice) HasOtherIpConfigs() bool`

HasOtherIpConfigs returns a boolean if a field has been set.

### GetPortMirroring

`func (o *ConfigDevice) GetPortMirroring() GatewayPortMirroring`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *ConfigDevice) GetPortMirroringOk() (*GatewayPortMirroring, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *ConfigDevice) SetPortMirroring(v GatewayPortMirroring)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *ConfigDevice) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetPortUsages

`func (o *ConfigDevice) GetPortUsages() map[string]SwitchPortUsage`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *ConfigDevice) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *ConfigDevice) SetPortUsages(v map[string]SwitchPortUsage)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *ConfigDevice) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *ConfigDevice) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *ConfigDevice) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *ConfigDevice) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *ConfigDevice) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRole

`func (o *ConfigDevice) GetRole() SwitchRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConfigDevice) GetRoleOk() (*SwitchRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConfigDevice) SetRole(v SwitchRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConfigDevice) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRouterId

`func (o *ConfigDevice) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *ConfigDevice) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *ConfigDevice) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *ConfigDevice) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetStpConfig

`func (o *ConfigDevice) GetStpConfig() SwitchStpConfig`

GetStpConfig returns the StpConfig field if non-nil, zero value otherwise.

### GetStpConfigOk

`func (o *ConfigDevice) GetStpConfigOk() (*SwitchStpConfig, bool)`

GetStpConfigOk returns a tuple with the StpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpConfig

`func (o *ConfigDevice) SetStpConfig(v SwitchStpConfig)`

SetStpConfig sets StpConfig field to given value.

### HasStpConfig

`func (o *ConfigDevice) HasStpConfig() bool`

HasStpConfig returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *ConfigDevice) GetSwitchMgmt() SwitchSwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *ConfigDevice) GetSwitchMgmtOk() (*SwitchSwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *ConfigDevice) SetSwitchMgmt(v SwitchSwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *ConfigDevice) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetUseRouterIdAsSourceIp

`func (o *ConfigDevice) GetUseRouterIdAsSourceIp() bool`

GetUseRouterIdAsSourceIp returns the UseRouterIdAsSourceIp field if non-nil, zero value otherwise.

### GetUseRouterIdAsSourceIpOk

`func (o *ConfigDevice) GetUseRouterIdAsSourceIpOk() (*bool, bool)`

GetUseRouterIdAsSourceIpOk returns a tuple with the UseRouterIdAsSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRouterIdAsSourceIp

`func (o *ConfigDevice) SetUseRouterIdAsSourceIp(v bool)`

SetUseRouterIdAsSourceIp sets UseRouterIdAsSourceIp field to given value.

### HasUseRouterIdAsSourceIp

`func (o *ConfigDevice) HasUseRouterIdAsSourceIp() bool`

HasUseRouterIdAsSourceIp returns a boolean if a field has been set.

### GetVirtualChassis

`func (o *ConfigDevice) GetVirtualChassis() SwitchVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *ConfigDevice) GetVirtualChassisOk() (*SwitchVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *ConfigDevice) SetVirtualChassis(v SwitchVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *ConfigDevice) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### GetVrfConfig

`func (o *ConfigDevice) GetVrfConfig() VrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *ConfigDevice) GetVrfConfigOk() (*VrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *ConfigDevice) SetVrfConfig(v VrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *ConfigDevice) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrrpConfig

`func (o *ConfigDevice) GetVrrpConfig() VrrpConfig`

GetVrrpConfig returns the VrrpConfig field if non-nil, zero value otherwise.

### GetVrrpConfigOk

`func (o *ConfigDevice) GetVrrpConfigOk() (*VrrpConfig, bool)`

GetVrrpConfigOk returns a tuple with the VrrpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpConfig

`func (o *ConfigDevice) SetVrrpConfig(v VrrpConfig)`

SetVrrpConfig sets VrrpConfig field to given value.

### HasVrrpConfig

`func (o *ConfigDevice) HasVrrpConfig() bool`

HasVrrpConfig returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *ConfigDevice) GetDhcpdConfig() DhcpdConfigs`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *ConfigDevice) GetDhcpdConfigOk() (*DhcpdConfigs, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *ConfigDevice) SetDhcpdConfig(v DhcpdConfigs)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *ConfigDevice) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetMspId

`func (o *ConfigDevice) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *ConfigDevice) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *ConfigDevice) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *ConfigDevice) HasMspId() bool`

HasMspId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


