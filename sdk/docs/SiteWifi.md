# SiteWifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiscoEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Disable11k** | Pointer to **bool** | whether to disable 11k | [optional] [default to false]
**DisableRadiosWhenPowerConstrained** | Pointer to **bool** |  | [optional] [default to false]
**EnableArpSpoofCheck** | Pointer to **bool** | when proxy_arp is enabled, check for arp spoofing. | [optional] [default to false]
**EnableSharedRadioScanning** | Pointer to **bool** |  | [optional] [default to true]
**Enabled** | Pointer to **bool** | enable WIFI feature (using SUB-MAN license) | [optional] [default to true]
**LocateConnected** | Pointer to **bool** | whether to locate connected clients | [optional] [default to true]
**LocateUnconnected** | Pointer to **bool** | whether to locate unconnected clients | [optional] [default to false]
**MeshAllowDfs** | Pointer to **bool** | whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will distrupt the connection. If roaming is desired, keep it disabled. | [optional] [default to false]
**MeshEnableCrm** | Pointer to **bool** | used to enable/disable CRM | [optional] [default to false]
**MeshEnabled** | Pointer to **bool** | whether to enable Mesh feature for the site | [optional] [default to false]
**MeshPsk** | Pointer to **NullableString** | optional passphrase of mesh networking, default is generated randomly | [optional] 
**MeshSsid** | Pointer to **NullableString** | optional ssid of mesh networking, default is based on site_id | [optional] 
**ProxyArp** | Pointer to [**NullableSiteWifiProxyArp**](SiteWifiProxyArp.md) |  | [optional] 

## Methods

### NewSiteWifi

`func NewSiteWifi() *SiteWifi`

NewSiteWifi instantiates a new SiteWifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWifiWithDefaults

`func NewSiteWifiWithDefaults() *SiteWifi`

NewSiteWifiWithDefaults instantiates a new SiteWifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiscoEnabled

`func (o *SiteWifi) GetCiscoEnabled() bool`

GetCiscoEnabled returns the CiscoEnabled field if non-nil, zero value otherwise.

### GetCiscoEnabledOk

`func (o *SiteWifi) GetCiscoEnabledOk() (*bool, bool)`

GetCiscoEnabledOk returns a tuple with the CiscoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoEnabled

`func (o *SiteWifi) SetCiscoEnabled(v bool)`

SetCiscoEnabled sets CiscoEnabled field to given value.

### HasCiscoEnabled

`func (o *SiteWifi) HasCiscoEnabled() bool`

HasCiscoEnabled returns a boolean if a field has been set.

### GetDisable11k

`func (o *SiteWifi) GetDisable11k() bool`

GetDisable11k returns the Disable11k field if non-nil, zero value otherwise.

### GetDisable11kOk

`func (o *SiteWifi) GetDisable11kOk() (*bool, bool)`

GetDisable11kOk returns a tuple with the Disable11k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable11k

`func (o *SiteWifi) SetDisable11k(v bool)`

SetDisable11k sets Disable11k field to given value.

### HasDisable11k

`func (o *SiteWifi) HasDisable11k() bool`

HasDisable11k returns a boolean if a field has been set.

### GetDisableRadiosWhenPowerConstrained

`func (o *SiteWifi) GetDisableRadiosWhenPowerConstrained() bool`

GetDisableRadiosWhenPowerConstrained returns the DisableRadiosWhenPowerConstrained field if non-nil, zero value otherwise.

### GetDisableRadiosWhenPowerConstrainedOk

`func (o *SiteWifi) GetDisableRadiosWhenPowerConstrainedOk() (*bool, bool)`

GetDisableRadiosWhenPowerConstrainedOk returns a tuple with the DisableRadiosWhenPowerConstrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRadiosWhenPowerConstrained

`func (o *SiteWifi) SetDisableRadiosWhenPowerConstrained(v bool)`

SetDisableRadiosWhenPowerConstrained sets DisableRadiosWhenPowerConstrained field to given value.

### HasDisableRadiosWhenPowerConstrained

`func (o *SiteWifi) HasDisableRadiosWhenPowerConstrained() bool`

HasDisableRadiosWhenPowerConstrained returns a boolean if a field has been set.

### GetEnableArpSpoofCheck

`func (o *SiteWifi) GetEnableArpSpoofCheck() bool`

GetEnableArpSpoofCheck returns the EnableArpSpoofCheck field if non-nil, zero value otherwise.

### GetEnableArpSpoofCheckOk

`func (o *SiteWifi) GetEnableArpSpoofCheckOk() (*bool, bool)`

GetEnableArpSpoofCheckOk returns a tuple with the EnableArpSpoofCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArpSpoofCheck

`func (o *SiteWifi) SetEnableArpSpoofCheck(v bool)`

SetEnableArpSpoofCheck sets EnableArpSpoofCheck field to given value.

### HasEnableArpSpoofCheck

`func (o *SiteWifi) HasEnableArpSpoofCheck() bool`

HasEnableArpSpoofCheck returns a boolean if a field has been set.

### GetEnableSharedRadioScanning

`func (o *SiteWifi) GetEnableSharedRadioScanning() bool`

GetEnableSharedRadioScanning returns the EnableSharedRadioScanning field if non-nil, zero value otherwise.

### GetEnableSharedRadioScanningOk

`func (o *SiteWifi) GetEnableSharedRadioScanningOk() (*bool, bool)`

GetEnableSharedRadioScanningOk returns a tuple with the EnableSharedRadioScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedRadioScanning

`func (o *SiteWifi) SetEnableSharedRadioScanning(v bool)`

SetEnableSharedRadioScanning sets EnableSharedRadioScanning field to given value.

### HasEnableSharedRadioScanning

`func (o *SiteWifi) HasEnableSharedRadioScanning() bool`

HasEnableSharedRadioScanning returns a boolean if a field has been set.

### GetEnabled

`func (o *SiteWifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteWifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteWifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteWifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocateConnected

`func (o *SiteWifi) GetLocateConnected() bool`

GetLocateConnected returns the LocateConnected field if non-nil, zero value otherwise.

### GetLocateConnectedOk

`func (o *SiteWifi) GetLocateConnectedOk() (*bool, bool)`

GetLocateConnectedOk returns a tuple with the LocateConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateConnected

`func (o *SiteWifi) SetLocateConnected(v bool)`

SetLocateConnected sets LocateConnected field to given value.

### HasLocateConnected

`func (o *SiteWifi) HasLocateConnected() bool`

HasLocateConnected returns a boolean if a field has been set.

### GetLocateUnconnected

`func (o *SiteWifi) GetLocateUnconnected() bool`

GetLocateUnconnected returns the LocateUnconnected field if non-nil, zero value otherwise.

### GetLocateUnconnectedOk

`func (o *SiteWifi) GetLocateUnconnectedOk() (*bool, bool)`

GetLocateUnconnectedOk returns a tuple with the LocateUnconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateUnconnected

`func (o *SiteWifi) SetLocateUnconnected(v bool)`

SetLocateUnconnected sets LocateUnconnected field to given value.

### HasLocateUnconnected

`func (o *SiteWifi) HasLocateUnconnected() bool`

HasLocateUnconnected returns a boolean if a field has been set.

### GetMeshAllowDfs

`func (o *SiteWifi) GetMeshAllowDfs() bool`

GetMeshAllowDfs returns the MeshAllowDfs field if non-nil, zero value otherwise.

### GetMeshAllowDfsOk

`func (o *SiteWifi) GetMeshAllowDfsOk() (*bool, bool)`

GetMeshAllowDfsOk returns a tuple with the MeshAllowDfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshAllowDfs

`func (o *SiteWifi) SetMeshAllowDfs(v bool)`

SetMeshAllowDfs sets MeshAllowDfs field to given value.

### HasMeshAllowDfs

`func (o *SiteWifi) HasMeshAllowDfs() bool`

HasMeshAllowDfs returns a boolean if a field has been set.

### GetMeshEnableCrm

`func (o *SiteWifi) GetMeshEnableCrm() bool`

GetMeshEnableCrm returns the MeshEnableCrm field if non-nil, zero value otherwise.

### GetMeshEnableCrmOk

`func (o *SiteWifi) GetMeshEnableCrmOk() (*bool, bool)`

GetMeshEnableCrmOk returns a tuple with the MeshEnableCrm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshEnableCrm

`func (o *SiteWifi) SetMeshEnableCrm(v bool)`

SetMeshEnableCrm sets MeshEnableCrm field to given value.

### HasMeshEnableCrm

`func (o *SiteWifi) HasMeshEnableCrm() bool`

HasMeshEnableCrm returns a boolean if a field has been set.

### GetMeshEnabled

`func (o *SiteWifi) GetMeshEnabled() bool`

GetMeshEnabled returns the MeshEnabled field if non-nil, zero value otherwise.

### GetMeshEnabledOk

`func (o *SiteWifi) GetMeshEnabledOk() (*bool, bool)`

GetMeshEnabledOk returns a tuple with the MeshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshEnabled

`func (o *SiteWifi) SetMeshEnabled(v bool)`

SetMeshEnabled sets MeshEnabled field to given value.

### HasMeshEnabled

`func (o *SiteWifi) HasMeshEnabled() bool`

HasMeshEnabled returns a boolean if a field has been set.

### GetMeshPsk

`func (o *SiteWifi) GetMeshPsk() string`

GetMeshPsk returns the MeshPsk field if non-nil, zero value otherwise.

### GetMeshPskOk

`func (o *SiteWifi) GetMeshPskOk() (*string, bool)`

GetMeshPskOk returns a tuple with the MeshPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshPsk

`func (o *SiteWifi) SetMeshPsk(v string)`

SetMeshPsk sets MeshPsk field to given value.

### HasMeshPsk

`func (o *SiteWifi) HasMeshPsk() bool`

HasMeshPsk returns a boolean if a field has been set.

### SetMeshPskNil

`func (o *SiteWifi) SetMeshPskNil(b bool)`

 SetMeshPskNil sets the value for MeshPsk to be an explicit nil

### UnsetMeshPsk
`func (o *SiteWifi) UnsetMeshPsk()`

UnsetMeshPsk ensures that no value is present for MeshPsk, not even an explicit nil
### GetMeshSsid

`func (o *SiteWifi) GetMeshSsid() string`

GetMeshSsid returns the MeshSsid field if non-nil, zero value otherwise.

### GetMeshSsidOk

`func (o *SiteWifi) GetMeshSsidOk() (*string, bool)`

GetMeshSsidOk returns a tuple with the MeshSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshSsid

`func (o *SiteWifi) SetMeshSsid(v string)`

SetMeshSsid sets MeshSsid field to given value.

### HasMeshSsid

`func (o *SiteWifi) HasMeshSsid() bool`

HasMeshSsid returns a boolean if a field has been set.

### SetMeshSsidNil

`func (o *SiteWifi) SetMeshSsidNil(b bool)`

 SetMeshSsidNil sets the value for MeshSsid to be an explicit nil

### UnsetMeshSsid
`func (o *SiteWifi) UnsetMeshSsid()`

UnsetMeshSsid ensures that no value is present for MeshSsid, not even an explicit nil
### GetProxyArp

`func (o *SiteWifi) GetProxyArp() SiteWifiProxyArp`

GetProxyArp returns the ProxyArp field if non-nil, zero value otherwise.

### GetProxyArpOk

`func (o *SiteWifi) GetProxyArpOk() (*SiteWifiProxyArp, bool)`

GetProxyArpOk returns a tuple with the ProxyArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyArp

`func (o *SiteWifi) SetProxyArp(v SiteWifiProxyArp)`

SetProxyArp sets ProxyArp field to given value.

### HasProxyArp

`func (o *SiteWifi) HasProxyArp() bool`

HasProxyArp returns a boolean if a field has been set.

### SetProxyArpNil

`func (o *SiteWifi) SetProxyArpNil(b bool)`

 SetProxyArpNil sets the value for ProxyArp to be an explicit nil

### UnsetProxyArp
`func (o *SiteWifi) UnsetProxyArp()`

UnsetProxyArp ensures that no value is present for ProxyArp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


