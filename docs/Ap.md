# Ap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aeroscout** | Pointer to [**ApAeroscout**](ApAeroscout.md) |  | [optional] 
**BleConfig** | Pointer to [**BleConfig**](BleConfig.md) |  | [optional] 
**Centrak** | Pointer to [**ApCentrak**](ApCentrak.md) |  | [optional] 
**ClientBridge** | Pointer to [**ApClientBridge**](ApClientBridge.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **NullableString** |  | [optional] 
**DisableEth1** | Pointer to **bool** | whether to disable eth1 port | [optional] [default to false]
**DisableEth2** | Pointer to **bool** | whether to disable eth2 port | [optional] [default to false]
**DisableEth3** | Pointer to **bool** | whether to disable eth3 port | [optional] [default to false]
**DisableModule** | Pointer to **bool** | whether to disable module port | [optional] [default to false]
**EslConfig** | Pointer to [**ApEslConfig**](ApEslConfig.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Height** | Pointer to **float32** | height, in meters, optional | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Image1Url** | Pointer to **NullableString** |  | [optional] 
**Image2Url** | Pointer to **NullableString** |  | [optional] 
**Image3Url** | Pointer to **NullableString** |  | [optional] 
**IotConfig** | Pointer to [**ApIot**](ApIot.md) |  | [optional] 
**IpConfig** | Pointer to [**ApIpConfig**](ApIpConfig.md) |  | [optional] 
**Led** | Pointer to [**ApLed**](ApLed.md) |  | [optional] 
**Locked** | Pointer to **bool** | whether this map is considered locked down | [optional] 
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**Mesh** | Pointer to [**ApMesh**](ApMesh.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** | any notes about this AP | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Orientation** | Pointer to **int32** | orientation, 0-359, in degrees, up is 0, right is 90. | [optional] 
**PoePassthrough** | Pointer to **bool** | whether to enable power out through module port (for APH) or eth1 (for APL/BT11) | [optional] [default to false]
**PortConfig** | Pointer to [**map[string]ApPortConfig**](ApPortConfig.md) | eth0 is not allowed here.  Property key is the interface(s) name (e.g. \&quot;eth1\&quot; or\&quot;eth1,eth2\&quot;) | [optional] 
**PwrConfig** | Pointer to [**ApPwrConfig**](ApPwrConfig.md) |  | [optional] 
**RadioConfig** | Pointer to [**ApRadio**](ApRadio.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**UplinkPortConfig** | Pointer to [**ApUplinkPortConfig**](ApUplinkPortConfig.md) |  | [optional] 
**UsbConfig** | Pointer to [**ApUsb**](ApUsb.md) |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**X** | Pointer to **float32** | x in pixel | [optional] 
**Y** | Pointer to **float32** | y in pixel | [optional] 

## Methods

### NewAp

`func NewAp() *Ap`

NewAp instantiates a new Ap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApWithDefaults

`func NewApWithDefaults() *Ap`

NewApWithDefaults instantiates a new Ap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeroscout

`func (o *Ap) GetAeroscout() ApAeroscout`

GetAeroscout returns the Aeroscout field if non-nil, zero value otherwise.

### GetAeroscoutOk

`func (o *Ap) GetAeroscoutOk() (*ApAeroscout, bool)`

GetAeroscoutOk returns a tuple with the Aeroscout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeroscout

`func (o *Ap) SetAeroscout(v ApAeroscout)`

SetAeroscout sets Aeroscout field to given value.

### HasAeroscout

`func (o *Ap) HasAeroscout() bool`

HasAeroscout returns a boolean if a field has been set.

### GetBleConfig

`func (o *Ap) GetBleConfig() BleConfig`

GetBleConfig returns the BleConfig field if non-nil, zero value otherwise.

### GetBleConfigOk

`func (o *Ap) GetBleConfigOk() (*BleConfig, bool)`

GetBleConfigOk returns a tuple with the BleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleConfig

`func (o *Ap) SetBleConfig(v BleConfig)`

SetBleConfig sets BleConfig field to given value.

### HasBleConfig

`func (o *Ap) HasBleConfig() bool`

HasBleConfig returns a boolean if a field has been set.

### GetCentrak

`func (o *Ap) GetCentrak() ApCentrak`

GetCentrak returns the Centrak field if non-nil, zero value otherwise.

### GetCentrakOk

`func (o *Ap) GetCentrakOk() (*ApCentrak, bool)`

GetCentrakOk returns a tuple with the Centrak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrak

`func (o *Ap) SetCentrak(v ApCentrak)`

SetCentrak sets Centrak field to given value.

### HasCentrak

`func (o *Ap) HasCentrak() bool`

HasCentrak returns a boolean if a field has been set.

### GetClientBridge

`func (o *Ap) GetClientBridge() ApClientBridge`

GetClientBridge returns the ClientBridge field if non-nil, zero value otherwise.

### GetClientBridgeOk

`func (o *Ap) GetClientBridgeOk() (*ApClientBridge, bool)`

GetClientBridgeOk returns a tuple with the ClientBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBridge

`func (o *Ap) SetClientBridge(v ApClientBridge)`

SetClientBridge sets ClientBridge field to given value.

### HasClientBridge

`func (o *Ap) HasClientBridge() bool`

HasClientBridge returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Ap) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Ap) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Ap) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Ap) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *Ap) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *Ap) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *Ap) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *Ap) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### SetDeviceprofileIdNil

`func (o *Ap) SetDeviceprofileIdNil(b bool)`

 SetDeviceprofileIdNil sets the value for DeviceprofileId to be an explicit nil

### UnsetDeviceprofileId
`func (o *Ap) UnsetDeviceprofileId()`

UnsetDeviceprofileId ensures that no value is present for DeviceprofileId, not even an explicit nil
### GetDisableEth1

`func (o *Ap) GetDisableEth1() bool`

GetDisableEth1 returns the DisableEth1 field if non-nil, zero value otherwise.

### GetDisableEth1Ok

`func (o *Ap) GetDisableEth1Ok() (*bool, bool)`

GetDisableEth1Ok returns a tuple with the DisableEth1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth1

`func (o *Ap) SetDisableEth1(v bool)`

SetDisableEth1 sets DisableEth1 field to given value.

### HasDisableEth1

`func (o *Ap) HasDisableEth1() bool`

HasDisableEth1 returns a boolean if a field has been set.

### GetDisableEth2

`func (o *Ap) GetDisableEth2() bool`

GetDisableEth2 returns the DisableEth2 field if non-nil, zero value otherwise.

### GetDisableEth2Ok

`func (o *Ap) GetDisableEth2Ok() (*bool, bool)`

GetDisableEth2Ok returns a tuple with the DisableEth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth2

`func (o *Ap) SetDisableEth2(v bool)`

SetDisableEth2 sets DisableEth2 field to given value.

### HasDisableEth2

`func (o *Ap) HasDisableEth2() bool`

HasDisableEth2 returns a boolean if a field has been set.

### GetDisableEth3

`func (o *Ap) GetDisableEth3() bool`

GetDisableEth3 returns the DisableEth3 field if non-nil, zero value otherwise.

### GetDisableEth3Ok

`func (o *Ap) GetDisableEth3Ok() (*bool, bool)`

GetDisableEth3Ok returns a tuple with the DisableEth3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEth3

`func (o *Ap) SetDisableEth3(v bool)`

SetDisableEth3 sets DisableEth3 field to given value.

### HasDisableEth3

`func (o *Ap) HasDisableEth3() bool`

HasDisableEth3 returns a boolean if a field has been set.

### GetDisableModule

`func (o *Ap) GetDisableModule() bool`

GetDisableModule returns the DisableModule field if non-nil, zero value otherwise.

### GetDisableModuleOk

`func (o *Ap) GetDisableModuleOk() (*bool, bool)`

GetDisableModuleOk returns a tuple with the DisableModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModule

`func (o *Ap) SetDisableModule(v bool)`

SetDisableModule sets DisableModule field to given value.

### HasDisableModule

`func (o *Ap) HasDisableModule() bool`

HasDisableModule returns a boolean if a field has been set.

### GetEslConfig

`func (o *Ap) GetEslConfig() ApEslConfig`

GetEslConfig returns the EslConfig field if non-nil, zero value otherwise.

### GetEslConfigOk

`func (o *Ap) GetEslConfigOk() (*ApEslConfig, bool)`

GetEslConfigOk returns a tuple with the EslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEslConfig

`func (o *Ap) SetEslConfig(v ApEslConfig)`

SetEslConfig sets EslConfig field to given value.

### HasEslConfig

`func (o *Ap) HasEslConfig() bool`

HasEslConfig returns a boolean if a field has been set.

### GetForSite

`func (o *Ap) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Ap) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Ap) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Ap) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHeight

`func (o *Ap) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Ap) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Ap) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Ap) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *Ap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ap) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *Ap) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *Ap) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *Ap) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *Ap) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### SetImage1UrlNil

`func (o *Ap) SetImage1UrlNil(b bool)`

 SetImage1UrlNil sets the value for Image1Url to be an explicit nil

### UnsetImage1Url
`func (o *Ap) UnsetImage1Url()`

UnsetImage1Url ensures that no value is present for Image1Url, not even an explicit nil
### GetImage2Url

`func (o *Ap) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *Ap) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *Ap) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *Ap) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### SetImage2UrlNil

`func (o *Ap) SetImage2UrlNil(b bool)`

 SetImage2UrlNil sets the value for Image2Url to be an explicit nil

### UnsetImage2Url
`func (o *Ap) UnsetImage2Url()`

UnsetImage2Url ensures that no value is present for Image2Url, not even an explicit nil
### GetImage3Url

`func (o *Ap) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *Ap) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *Ap) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *Ap) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### SetImage3UrlNil

`func (o *Ap) SetImage3UrlNil(b bool)`

 SetImage3UrlNil sets the value for Image3Url to be an explicit nil

### UnsetImage3Url
`func (o *Ap) UnsetImage3Url()`

UnsetImage3Url ensures that no value is present for Image3Url, not even an explicit nil
### GetIotConfig

`func (o *Ap) GetIotConfig() ApIot`

GetIotConfig returns the IotConfig field if non-nil, zero value otherwise.

### GetIotConfigOk

`func (o *Ap) GetIotConfigOk() (*ApIot, bool)`

GetIotConfigOk returns a tuple with the IotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotConfig

`func (o *Ap) SetIotConfig(v ApIot)`

SetIotConfig sets IotConfig field to given value.

### HasIotConfig

`func (o *Ap) HasIotConfig() bool`

HasIotConfig returns a boolean if a field has been set.

### GetIpConfig

`func (o *Ap) GetIpConfig() ApIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *Ap) GetIpConfigOk() (*ApIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *Ap) SetIpConfig(v ApIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *Ap) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLed

`func (o *Ap) GetLed() ApLed`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *Ap) GetLedOk() (*ApLed, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *Ap) SetLed(v ApLed)`

SetLed sets Led field to given value.

### HasLed

`func (o *Ap) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLocked

`func (o *Ap) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Ap) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Ap) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Ap) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMapId

`func (o *Ap) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *Ap) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *Ap) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *Ap) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMesh

`func (o *Ap) GetMesh() ApMesh`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *Ap) GetMeshOk() (*ApMesh, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *Ap) SetMesh(v ApMesh)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *Ap) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Ap) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Ap) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Ap) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Ap) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Ap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *Ap) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Ap) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Ap) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Ap) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpServers

`func (o *Ap) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *Ap) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *Ap) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *Ap) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *Ap) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Ap) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Ap) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Ap) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrientation

`func (o *Ap) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Ap) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Ap) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Ap) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPoePassthrough

`func (o *Ap) GetPoePassthrough() bool`

GetPoePassthrough returns the PoePassthrough field if non-nil, zero value otherwise.

### GetPoePassthroughOk

`func (o *Ap) GetPoePassthroughOk() (*bool, bool)`

GetPoePassthroughOk returns a tuple with the PoePassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePassthrough

`func (o *Ap) SetPoePassthrough(v bool)`

SetPoePassthrough sets PoePassthrough field to given value.

### HasPoePassthrough

`func (o *Ap) HasPoePassthrough() bool`

HasPoePassthrough returns a boolean if a field has been set.

### GetPortConfig

`func (o *Ap) GetPortConfig() map[string]ApPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *Ap) GetPortConfigOk() (*map[string]ApPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *Ap) SetPortConfig(v map[string]ApPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *Ap) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPwrConfig

`func (o *Ap) GetPwrConfig() ApPwrConfig`

GetPwrConfig returns the PwrConfig field if non-nil, zero value otherwise.

### GetPwrConfigOk

`func (o *Ap) GetPwrConfigOk() (*ApPwrConfig, bool)`

GetPwrConfigOk returns a tuple with the PwrConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrConfig

`func (o *Ap) SetPwrConfig(v ApPwrConfig)`

SetPwrConfig sets PwrConfig field to given value.

### HasPwrConfig

`func (o *Ap) HasPwrConfig() bool`

HasPwrConfig returns a boolean if a field has been set.

### GetRadioConfig

`func (o *Ap) GetRadioConfig() ApRadio`

GetRadioConfig returns the RadioConfig field if non-nil, zero value otherwise.

### GetRadioConfigOk

`func (o *Ap) GetRadioConfigOk() (*ApRadio, bool)`

GetRadioConfigOk returns a tuple with the RadioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioConfig

`func (o *Ap) SetRadioConfig(v ApRadio)`

SetRadioConfig sets RadioConfig field to given value.

### HasRadioConfig

`func (o *Ap) HasRadioConfig() bool`

HasRadioConfig returns a boolean if a field has been set.

### GetSiteId

`func (o *Ap) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Ap) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Ap) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Ap) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUplinkPortConfig

`func (o *Ap) GetUplinkPortConfig() ApUplinkPortConfig`

GetUplinkPortConfig returns the UplinkPortConfig field if non-nil, zero value otherwise.

### GetUplinkPortConfigOk

`func (o *Ap) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool)`

GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPortConfig

`func (o *Ap) SetUplinkPortConfig(v ApUplinkPortConfig)`

SetUplinkPortConfig sets UplinkPortConfig field to given value.

### HasUplinkPortConfig

`func (o *Ap) HasUplinkPortConfig() bool`

HasUplinkPortConfig returns a boolean if a field has been set.

### GetUsbConfig

`func (o *Ap) GetUsbConfig() ApUsb`

GetUsbConfig returns the UsbConfig field if non-nil, zero value otherwise.

### GetUsbConfigOk

`func (o *Ap) GetUsbConfigOk() (*ApUsb, bool)`

GetUsbConfigOk returns a tuple with the UsbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbConfig

`func (o *Ap) SetUsbConfig(v ApUsb)`

SetUsbConfig sets UsbConfig field to given value.

### HasUsbConfig

`func (o *Ap) HasUsbConfig() bool`

HasUsbConfig returns a boolean if a field has been set.

### GetVars

`func (o *Ap) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *Ap) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *Ap) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *Ap) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetX

`func (o *Ap) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Ap) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Ap) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *Ap) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *Ap) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Ap) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Ap) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *Ap) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


