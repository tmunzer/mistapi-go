# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**DhcpdConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**ExtraRoutes** | Pointer to [**map[string]GatewayExtraRoute**](GatewayExtraRoute.md) |  | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Image1Url** | Pointer to **string** |  | [optional] 
**Image2Url** | Pointer to **string** |  | [optional] 
**Image3Url** | Pointer to **string** |  | [optional] 
**IpConfig** | Pointer to [**map[string]GatewayTemplateIpConfig**](GatewayTemplateIpConfig.md) | Property key is the network name | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to [**map[string]GatewayNetwork**](GatewayNetwork.md) | Property key is the network name or a CIDR | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**OobIpConfig** | Pointer to [**JunosOobIpConfigs**](JunosOobIpConfigs.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortConfig** | Pointer to [**map[string]GatewayPortConfig**](GatewayPortConfig.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PortMirroring** | Pointer to [**GatewayPortMirroring**](GatewayPortMirroring.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 

## Methods

### NewGateway

`func NewGateway() *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalConfigCmds

`func (o *Gateway) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *Gateway) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *Gateway) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *Gateway) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Gateway) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Gateway) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Gateway) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Gateway) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *Gateway) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *Gateway) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *Gateway) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *Gateway) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *Gateway) GetDhcpdConfig() DhcpdConfigs`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *Gateway) GetDhcpdConfigOk() (*DhcpdConfigs, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *Gateway) SetDhcpdConfig(v DhcpdConfigs)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *Gateway) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *Gateway) GetExtraRoutes() map[string]GatewayExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *Gateway) GetExtraRoutesOk() (*map[string]GatewayExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *Gateway) SetExtraRoutes(v map[string]GatewayExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *Gateway) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetForSite

`func (o *Gateway) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Gateway) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Gateway) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Gateway) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *Gateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Gateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *Gateway) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *Gateway) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *Gateway) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *Gateway) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### GetImage2Url

`func (o *Gateway) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *Gateway) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *Gateway) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *Gateway) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### GetImage3Url

`func (o *Gateway) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *Gateway) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *Gateway) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *Gateway) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### GetIpConfig

`func (o *Gateway) GetIpConfig() map[string]GatewayTemplateIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *Gateway) GetIpConfigOk() (*map[string]GatewayTemplateIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *Gateway) SetIpConfig(v map[string]GatewayTemplateIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *Gateway) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetManaged

`func (o *Gateway) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Gateway) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Gateway) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Gateway) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Gateway) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Gateway) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Gateway) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Gateway) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *Gateway) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *Gateway) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *Gateway) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *Gateway) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *Gateway) GetNetworks() map[string]GatewayNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Gateway) GetNetworksOk() (*map[string]GatewayNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Gateway) SetNetworks(v map[string]GatewayNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *Gateway) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNtpServers

`func (o *Gateway) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *Gateway) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *Gateway) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *Gateway) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *Gateway) GetOobIpConfig() JunosOobIpConfigs`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *Gateway) GetOobIpConfigOk() (*JunosOobIpConfigs, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *Gateway) SetOobIpConfig(v JunosOobIpConfigs)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *Gateway) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *Gateway) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Gateway) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Gateway) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Gateway) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortConfig

`func (o *Gateway) GetPortConfig() map[string]GatewayPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *Gateway) GetPortConfigOk() (*map[string]GatewayPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *Gateway) SetPortConfig(v map[string]GatewayPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *Gateway) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPortMirroring

`func (o *Gateway) GetPortMirroring() GatewayPortMirroring`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *Gateway) GetPortMirroringOk() (*GatewayPortMirroring, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *Gateway) SetPortMirroring(v GatewayPortMirroring)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *Gateway) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetSiteId

`func (o *Gateway) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Gateway) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Gateway) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Gateway) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVars

`func (o *Gateway) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *Gateway) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *Gateway) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *Gateway) HasVars() bool`

HasVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


