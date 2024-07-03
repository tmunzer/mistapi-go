# Mxedge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Magic** | Pointer to **string** |  | [optional] [readonly] 
**Model** | **string** |  | 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MxagentRegistered** | Pointer to **bool** |  | [optional] [readonly] 
**MxclusterId** | Pointer to **string** | MxCluster this MxEdge belongs to | [optional] 
**MxedgeMgmt** | Pointer to [**MxedgeMgmt**](MxedgeMgmt.md) |  | [optional] 
**Name** | **string** |  | 
**Note** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OobIpConfig** | Pointer to [**MxedgeOobIpConfig**](MxedgeOobIpConfig.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Proxy** | Pointer to [**Proxy**](Proxy.md) |  | [optional] 
**Services** | Pointer to **[]string** | list of services to run, tunterm only for now | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**TuntermDhcpdConfig** | Pointer to [**MxedgeTuntermDhcpdConfigs**](MxedgeTuntermDhcpdConfigs.md) |  | [optional] 
**TuntermExtraRoutes** | Pointer to [**map[string]MxedgeTuntermExtraRoute**](MxedgeTuntermExtraRoute.md) | Property key is a CIDR | [optional] 
**TuntermIgmpSnoopingConfig** | Pointer to [**MxedgeTuntermIgmpSnoopingConfig**](MxedgeTuntermIgmpSnoopingConfig.md) |  | [optional] 
**TuntermIpConfig** | Pointer to [**MxedgeTuntermIpConfig**](MxedgeTuntermIpConfig.md) |  | [optional] 
**TuntermMonitoring** | Pointer to [**[][]TuntermMonitoringItem**]([]TuntermMonitoringItem.md) |  | [optional] 
**TuntermMulticastConfig** | Pointer to [**MxedgeTuntermMulticastConfig**](MxedgeTuntermMulticastConfig.md) |  | [optional] 
**TuntermOtherIpConfigs** | Pointer to [**map[string]MxedgeTuntermOtherIpConfig**](MxedgeTuntermOtherIpConfig.md) | ip configs by VLAN ID. Property key is the VLAN ID | [optional] 
**TuntermPortConfig** | Pointer to [**TuntermPortConfig**](TuntermPortConfig.md) |  | [optional] 
**TuntermRegistered** | Pointer to **bool** |  | [optional] [readonly] 
**TuntermSwitchConfig** | Pointer to [**MxedgeTuntermSwitchConfigs**](MxedgeTuntermSwitchConfigs.md) |  | [optional] 
**Versions** | Pointer to [**MxedgeVersions**](MxedgeVersions.md) |  | [optional] 

## Methods

### NewMxedge

`func NewMxedge(model string, name string, ) *Mxedge`

NewMxedge instantiates a new Mxedge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeWithDefaults

`func NewMxedgeWithDefaults() *Mxedge`

NewMxedgeWithDefaults instantiates a new Mxedge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Mxedge) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Mxedge) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Mxedge) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Mxedge) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *Mxedge) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Mxedge) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Mxedge) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Mxedge) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *Mxedge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mxedge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mxedge) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mxedge) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMagic

`func (o *Mxedge) GetMagic() string`

GetMagic returns the Magic field if non-nil, zero value otherwise.

### GetMagicOk

`func (o *Mxedge) GetMagicOk() (*string, bool)`

GetMagicOk returns a tuple with the Magic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagic

`func (o *Mxedge) SetMagic(v string)`

SetMagic sets Magic field to given value.

### HasMagic

`func (o *Mxedge) HasMagic() bool`

HasMagic returns a boolean if a field has been set.

### GetModel

`func (o *Mxedge) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Mxedge) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Mxedge) SetModel(v string)`

SetModel sets Model field to given value.


### GetModifiedTime

`func (o *Mxedge) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Mxedge) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Mxedge) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Mxedge) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMxagentRegistered

`func (o *Mxedge) GetMxagentRegistered() bool`

GetMxagentRegistered returns the MxagentRegistered field if non-nil, zero value otherwise.

### GetMxagentRegisteredOk

`func (o *Mxedge) GetMxagentRegisteredOk() (*bool, bool)`

GetMxagentRegisteredOk returns a tuple with the MxagentRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxagentRegistered

`func (o *Mxedge) SetMxagentRegistered(v bool)`

SetMxagentRegistered sets MxagentRegistered field to given value.

### HasMxagentRegistered

`func (o *Mxedge) HasMxagentRegistered() bool`

HasMxagentRegistered returns a boolean if a field has been set.

### GetMxclusterId

`func (o *Mxedge) GetMxclusterId() string`

GetMxclusterId returns the MxclusterId field if non-nil, zero value otherwise.

### GetMxclusterIdOk

`func (o *Mxedge) GetMxclusterIdOk() (*string, bool)`

GetMxclusterIdOk returns a tuple with the MxclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterId

`func (o *Mxedge) SetMxclusterId(v string)`

SetMxclusterId sets MxclusterId field to given value.

### HasMxclusterId

`func (o *Mxedge) HasMxclusterId() bool`

HasMxclusterId returns a boolean if a field has been set.

### GetMxedgeMgmt

`func (o *Mxedge) GetMxedgeMgmt() MxedgeMgmt`

GetMxedgeMgmt returns the MxedgeMgmt field if non-nil, zero value otherwise.

### GetMxedgeMgmtOk

`func (o *Mxedge) GetMxedgeMgmtOk() (*MxedgeMgmt, bool)`

GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeMgmt

`func (o *Mxedge) SetMxedgeMgmt(v MxedgeMgmt)`

SetMxedgeMgmt sets MxedgeMgmt field to given value.

### HasMxedgeMgmt

`func (o *Mxedge) HasMxedgeMgmt() bool`

HasMxedgeMgmt returns a boolean if a field has been set.

### GetName

`func (o *Mxedge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mxedge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mxedge) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *Mxedge) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Mxedge) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Mxedge) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Mxedge) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNtpServers

`func (o *Mxedge) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *Mxedge) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *Mxedge) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *Mxedge) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *Mxedge) GetOobIpConfig() MxedgeOobIpConfig`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *Mxedge) GetOobIpConfigOk() (*MxedgeOobIpConfig, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *Mxedge) SetOobIpConfig(v MxedgeOobIpConfig)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *Mxedge) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *Mxedge) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Mxedge) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Mxedge) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Mxedge) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProxy

`func (o *Mxedge) GetProxy() Proxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Mxedge) GetProxyOk() (*Proxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Mxedge) SetProxy(v Proxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Mxedge) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetServices

`func (o *Mxedge) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Mxedge) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Mxedge) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Mxedge) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSiteId

`func (o *Mxedge) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Mxedge) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Mxedge) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Mxedge) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTuntermDhcpdConfig

`func (o *Mxedge) GetTuntermDhcpdConfig() MxedgeTuntermDhcpdConfigs`

GetTuntermDhcpdConfig returns the TuntermDhcpdConfig field if non-nil, zero value otherwise.

### GetTuntermDhcpdConfigOk

`func (o *Mxedge) GetTuntermDhcpdConfigOk() (*MxedgeTuntermDhcpdConfigs, bool)`

GetTuntermDhcpdConfigOk returns a tuple with the TuntermDhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermDhcpdConfig

`func (o *Mxedge) SetTuntermDhcpdConfig(v MxedgeTuntermDhcpdConfigs)`

SetTuntermDhcpdConfig sets TuntermDhcpdConfig field to given value.

### HasTuntermDhcpdConfig

`func (o *Mxedge) HasTuntermDhcpdConfig() bool`

HasTuntermDhcpdConfig returns a boolean if a field has been set.

### GetTuntermExtraRoutes

`func (o *Mxedge) GetTuntermExtraRoutes() map[string]MxedgeTuntermExtraRoute`

GetTuntermExtraRoutes returns the TuntermExtraRoutes field if non-nil, zero value otherwise.

### GetTuntermExtraRoutesOk

`func (o *Mxedge) GetTuntermExtraRoutesOk() (*map[string]MxedgeTuntermExtraRoute, bool)`

GetTuntermExtraRoutesOk returns a tuple with the TuntermExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermExtraRoutes

`func (o *Mxedge) SetTuntermExtraRoutes(v map[string]MxedgeTuntermExtraRoute)`

SetTuntermExtraRoutes sets TuntermExtraRoutes field to given value.

### HasTuntermExtraRoutes

`func (o *Mxedge) HasTuntermExtraRoutes() bool`

HasTuntermExtraRoutes returns a boolean if a field has been set.

### GetTuntermIgmpSnoopingConfig

`func (o *Mxedge) GetTuntermIgmpSnoopingConfig() MxedgeTuntermIgmpSnoopingConfig`

GetTuntermIgmpSnoopingConfig returns the TuntermIgmpSnoopingConfig field if non-nil, zero value otherwise.

### GetTuntermIgmpSnoopingConfigOk

`func (o *Mxedge) GetTuntermIgmpSnoopingConfigOk() (*MxedgeTuntermIgmpSnoopingConfig, bool)`

GetTuntermIgmpSnoopingConfigOk returns a tuple with the TuntermIgmpSnoopingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermIgmpSnoopingConfig

`func (o *Mxedge) SetTuntermIgmpSnoopingConfig(v MxedgeTuntermIgmpSnoopingConfig)`

SetTuntermIgmpSnoopingConfig sets TuntermIgmpSnoopingConfig field to given value.

### HasTuntermIgmpSnoopingConfig

`func (o *Mxedge) HasTuntermIgmpSnoopingConfig() bool`

HasTuntermIgmpSnoopingConfig returns a boolean if a field has been set.

### GetTuntermIpConfig

`func (o *Mxedge) GetTuntermIpConfig() MxedgeTuntermIpConfig`

GetTuntermIpConfig returns the TuntermIpConfig field if non-nil, zero value otherwise.

### GetTuntermIpConfigOk

`func (o *Mxedge) GetTuntermIpConfigOk() (*MxedgeTuntermIpConfig, bool)`

GetTuntermIpConfigOk returns a tuple with the TuntermIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermIpConfig

`func (o *Mxedge) SetTuntermIpConfig(v MxedgeTuntermIpConfig)`

SetTuntermIpConfig sets TuntermIpConfig field to given value.

### HasTuntermIpConfig

`func (o *Mxedge) HasTuntermIpConfig() bool`

HasTuntermIpConfig returns a boolean if a field has been set.

### GetTuntermMonitoring

`func (o *Mxedge) GetTuntermMonitoring() [][]TuntermMonitoringItem`

GetTuntermMonitoring returns the TuntermMonitoring field if non-nil, zero value otherwise.

### GetTuntermMonitoringOk

`func (o *Mxedge) GetTuntermMonitoringOk() (*[][]TuntermMonitoringItem, bool)`

GetTuntermMonitoringOk returns a tuple with the TuntermMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMonitoring

`func (o *Mxedge) SetTuntermMonitoring(v [][]TuntermMonitoringItem)`

SetTuntermMonitoring sets TuntermMonitoring field to given value.

### HasTuntermMonitoring

`func (o *Mxedge) HasTuntermMonitoring() bool`

HasTuntermMonitoring returns a boolean if a field has been set.

### GetTuntermMulticastConfig

`func (o *Mxedge) GetTuntermMulticastConfig() MxedgeTuntermMulticastConfig`

GetTuntermMulticastConfig returns the TuntermMulticastConfig field if non-nil, zero value otherwise.

### GetTuntermMulticastConfigOk

`func (o *Mxedge) GetTuntermMulticastConfigOk() (*MxedgeTuntermMulticastConfig, bool)`

GetTuntermMulticastConfigOk returns a tuple with the TuntermMulticastConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMulticastConfig

`func (o *Mxedge) SetTuntermMulticastConfig(v MxedgeTuntermMulticastConfig)`

SetTuntermMulticastConfig sets TuntermMulticastConfig field to given value.

### HasTuntermMulticastConfig

`func (o *Mxedge) HasTuntermMulticastConfig() bool`

HasTuntermMulticastConfig returns a boolean if a field has been set.

### GetTuntermOtherIpConfigs

`func (o *Mxedge) GetTuntermOtherIpConfigs() map[string]MxedgeTuntermOtherIpConfig`

GetTuntermOtherIpConfigs returns the TuntermOtherIpConfigs field if non-nil, zero value otherwise.

### GetTuntermOtherIpConfigsOk

`func (o *Mxedge) GetTuntermOtherIpConfigsOk() (*map[string]MxedgeTuntermOtherIpConfig, bool)`

GetTuntermOtherIpConfigsOk returns a tuple with the TuntermOtherIpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermOtherIpConfigs

`func (o *Mxedge) SetTuntermOtherIpConfigs(v map[string]MxedgeTuntermOtherIpConfig)`

SetTuntermOtherIpConfigs sets TuntermOtherIpConfigs field to given value.

### HasTuntermOtherIpConfigs

`func (o *Mxedge) HasTuntermOtherIpConfigs() bool`

HasTuntermOtherIpConfigs returns a boolean if a field has been set.

### GetTuntermPortConfig

`func (o *Mxedge) GetTuntermPortConfig() TuntermPortConfig`

GetTuntermPortConfig returns the TuntermPortConfig field if non-nil, zero value otherwise.

### GetTuntermPortConfigOk

`func (o *Mxedge) GetTuntermPortConfigOk() (*TuntermPortConfig, bool)`

GetTuntermPortConfigOk returns a tuple with the TuntermPortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermPortConfig

`func (o *Mxedge) SetTuntermPortConfig(v TuntermPortConfig)`

SetTuntermPortConfig sets TuntermPortConfig field to given value.

### HasTuntermPortConfig

`func (o *Mxedge) HasTuntermPortConfig() bool`

HasTuntermPortConfig returns a boolean if a field has been set.

### GetTuntermRegistered

`func (o *Mxedge) GetTuntermRegistered() bool`

GetTuntermRegistered returns the TuntermRegistered field if non-nil, zero value otherwise.

### GetTuntermRegisteredOk

`func (o *Mxedge) GetTuntermRegisteredOk() (*bool, bool)`

GetTuntermRegisteredOk returns a tuple with the TuntermRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermRegistered

`func (o *Mxedge) SetTuntermRegistered(v bool)`

SetTuntermRegistered sets TuntermRegistered field to given value.

### HasTuntermRegistered

`func (o *Mxedge) HasTuntermRegistered() bool`

HasTuntermRegistered returns a boolean if a field has been set.

### GetTuntermSwitchConfig

`func (o *Mxedge) GetTuntermSwitchConfig() MxedgeTuntermSwitchConfigs`

GetTuntermSwitchConfig returns the TuntermSwitchConfig field if non-nil, zero value otherwise.

### GetTuntermSwitchConfigOk

`func (o *Mxedge) GetTuntermSwitchConfigOk() (*MxedgeTuntermSwitchConfigs, bool)`

GetTuntermSwitchConfigOk returns a tuple with the TuntermSwitchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermSwitchConfig

`func (o *Mxedge) SetTuntermSwitchConfig(v MxedgeTuntermSwitchConfigs)`

SetTuntermSwitchConfig sets TuntermSwitchConfig field to given value.

### HasTuntermSwitchConfig

`func (o *Mxedge) HasTuntermSwitchConfig() bool`

HasTuntermSwitchConfig returns a boolean if a field has been set.

### GetVersions

`func (o *Mxedge) GetVersions() MxedgeVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Mxedge) GetVersionsOk() (*MxedgeVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Mxedge) SetVersions(v MxedgeVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Mxedge) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


