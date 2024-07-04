# ModelSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclPolicies** | Pointer to [**[]AclPolicy**](AclPolicy.md) |  | [optional] 
**AclTags** | Pointer to [**map[string]AclTag**](AclTag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**DhcpConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**DhcpSnooping** | Pointer to [**DhcpSnooping**](DhcpSnooping.md) |  | [optional] 
**DisableAutoConfig** | Pointer to **bool** | for a claimed switch, we control the configs by default. This option (disables the behavior) | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**EvpnConfig** | Pointer to [**EvpnConfig**](EvpnConfig.md) |  | [optional] 
**ExtraRoutes** | Pointer to [**map[string]ExtraRouteProperties**](ExtraRouteProperties.md) |  | [optional] 
**ExtraRoutes6** | Pointer to [**map[string]ExtraRoute6Properties**](ExtraRoute6Properties.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Image1Url** | Pointer to **NullableString** |  | [optional] 
**Image2Url** | Pointer to **NullableString** |  | [optional] 
**Image3Url** | Pointer to **NullableString** |  | [optional] 
**IpConfig** | Pointer to [**JunosIpConfig**](JunosIpConfig.md) |  | [optional] 
**Managed** | Pointer to **bool** | for an adopted switch, we don’t overwrite their existing configs automatically | [optional] [default to false]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to [**map[string]SwitchNetwork**](SwitchNetwork.md) | Property key is network name | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] 
**OobIpConfig** | Pointer to [**JunosOobIpConfigs**](JunosOobIpConfigs.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**OspfConfig** | Pointer to [**OspfConfig**](OspfConfig.md) |  | [optional] 
**OtherIpConfigs** | Pointer to [**map[string]JunosOtherIpConfigs**](JunosOtherIpConfigs.md) | Property key is the network name | [optional] 
**PortConfig** | Pointer to [**map[string]JunosPortConfig**](JunosPortConfig.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PortMirroring** | Pointer to [**map[string]SwitchPortMirroring**](SwitchPortMirroring.md) | Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. | [optional] 
**PortUsages** | Pointer to [**map[string]SwitchPortUsage**](SwitchPortUsage.md) |  | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**Role** | Pointer to [**SwitchRole**](SwitchRole.md) |  | [optional] [default to SWITCHROLE_ACCESS]
**RouterId** | Pointer to **string** | used for OSPF / BGP / EVPN | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**StpConfig** | Pointer to [**SwitchStpConfig**](SwitchStpConfig.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchSwitchMgmt**](SwitchSwitchMgmt.md) |  | [optional] 
**UseRouterIdAsSourceIp** | Pointer to **bool** | whether to use it for snmp / syslog / tacplus / radius | [optional] [default to false]
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**VirtualChassis** | Pointer to [**SwitchVirtualChassis**](SwitchVirtualChassis.md) |  | [optional] 
**VrfConfig** | Pointer to [**VrfConfig**](VrfConfig.md) |  | [optional] 
**VrrpConfig** | Pointer to [**VrrpConfig**](VrrpConfig.md) |  | [optional] 

## Methods

### NewModelSwitch

`func NewModelSwitch() *ModelSwitch`

NewModelSwitch instantiates a new ModelSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSwitchWithDefaults

`func NewModelSwitchWithDefaults() *ModelSwitch`

NewModelSwitchWithDefaults instantiates a new ModelSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclPolicies

`func (o *ModelSwitch) GetAclPolicies() []AclPolicy`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *ModelSwitch) GetAclPoliciesOk() (*[]AclPolicy, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *ModelSwitch) SetAclPolicies(v []AclPolicy)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *ModelSwitch) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *ModelSwitch) GetAclTags() map[string]AclTag`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *ModelSwitch) GetAclTagsOk() (*map[string]AclTag, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *ModelSwitch) SetAclTags(v map[string]AclTag)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *ModelSwitch) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *ModelSwitch) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *ModelSwitch) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *ModelSwitch) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *ModelSwitch) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ModelSwitch) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ModelSwitch) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ModelSwitch) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ModelSwitch) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *ModelSwitch) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *ModelSwitch) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *ModelSwitch) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *ModelSwitch) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *ModelSwitch) GetDhcpConfig() DhcpdConfigs`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *ModelSwitch) GetDhcpConfigOk() (*DhcpdConfigs, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *ModelSwitch) SetDhcpConfig(v DhcpdConfigs)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *ModelSwitch) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *ModelSwitch) GetDhcpSnooping() DhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *ModelSwitch) GetDhcpSnoopingOk() (*DhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *ModelSwitch) SetDhcpSnooping(v DhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *ModelSwitch) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDisableAutoConfig

`func (o *ModelSwitch) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *ModelSwitch) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *ModelSwitch) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *ModelSwitch) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetDnsServers

`func (o *ModelSwitch) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ModelSwitch) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ModelSwitch) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ModelSwitch) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *ModelSwitch) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *ModelSwitch) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *ModelSwitch) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *ModelSwitch) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEvpnConfig

`func (o *ModelSwitch) GetEvpnConfig() EvpnConfig`

GetEvpnConfig returns the EvpnConfig field if non-nil, zero value otherwise.

### GetEvpnConfigOk

`func (o *ModelSwitch) GetEvpnConfigOk() (*EvpnConfig, bool)`

GetEvpnConfigOk returns a tuple with the EvpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnConfig

`func (o *ModelSwitch) SetEvpnConfig(v EvpnConfig)`

SetEvpnConfig sets EvpnConfig field to given value.

### HasEvpnConfig

`func (o *ModelSwitch) HasEvpnConfig() bool`

HasEvpnConfig returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *ModelSwitch) GetExtraRoutes() map[string]ExtraRouteProperties`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *ModelSwitch) GetExtraRoutesOk() (*map[string]ExtraRouteProperties, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *ModelSwitch) SetExtraRoutes(v map[string]ExtraRouteProperties)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *ModelSwitch) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *ModelSwitch) GetExtraRoutes6() map[string]ExtraRoute6Properties`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *ModelSwitch) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *ModelSwitch) SetExtraRoutes6(v map[string]ExtraRoute6Properties)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *ModelSwitch) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetId

`func (o *ModelSwitch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSwitch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSwitch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSwitch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *ModelSwitch) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *ModelSwitch) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *ModelSwitch) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *ModelSwitch) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### SetImage1UrlNil

`func (o *ModelSwitch) SetImage1UrlNil(b bool)`

 SetImage1UrlNil sets the value for Image1Url to be an explicit nil

### UnsetImage1Url
`func (o *ModelSwitch) UnsetImage1Url()`

UnsetImage1Url ensures that no value is present for Image1Url, not even an explicit nil
### GetImage2Url

`func (o *ModelSwitch) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *ModelSwitch) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *ModelSwitch) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *ModelSwitch) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### SetImage2UrlNil

`func (o *ModelSwitch) SetImage2UrlNil(b bool)`

 SetImage2UrlNil sets the value for Image2Url to be an explicit nil

### UnsetImage2Url
`func (o *ModelSwitch) UnsetImage2Url()`

UnsetImage2Url ensures that no value is present for Image2Url, not even an explicit nil
### GetImage3Url

`func (o *ModelSwitch) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *ModelSwitch) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *ModelSwitch) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *ModelSwitch) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### SetImage3UrlNil

`func (o *ModelSwitch) SetImage3UrlNil(b bool)`

 SetImage3UrlNil sets the value for Image3Url to be an explicit nil

### UnsetImage3Url
`func (o *ModelSwitch) UnsetImage3Url()`

UnsetImage3Url ensures that no value is present for Image3Url, not even an explicit nil
### GetIpConfig

`func (o *ModelSwitch) GetIpConfig() JunosIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ModelSwitch) GetIpConfigOk() (*JunosIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ModelSwitch) SetIpConfig(v JunosIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ModelSwitch) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetManaged

`func (o *ModelSwitch) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ModelSwitch) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ModelSwitch) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ModelSwitch) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetModifiedTime

`func (o *ModelSwitch) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *ModelSwitch) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *ModelSwitch) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *ModelSwitch) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *ModelSwitch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSwitch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSwitch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSwitch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *ModelSwitch) GetNetworks() map[string]SwitchNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ModelSwitch) GetNetworksOk() (*map[string]SwitchNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ModelSwitch) SetNetworks(v map[string]SwitchNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ModelSwitch) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNotes

`func (o *ModelSwitch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ModelSwitch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ModelSwitch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ModelSwitch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpServers

`func (o *ModelSwitch) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ModelSwitch) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ModelSwitch) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *ModelSwitch) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *ModelSwitch) GetOobIpConfig() JunosOobIpConfigs`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *ModelSwitch) GetOobIpConfigOk() (*JunosOobIpConfigs, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *ModelSwitch) SetOobIpConfig(v JunosOobIpConfigs)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *ModelSwitch) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelSwitch) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelSwitch) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelSwitch) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelSwitch) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOspfConfig

`func (o *ModelSwitch) GetOspfConfig() OspfConfig`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *ModelSwitch) GetOspfConfigOk() (*OspfConfig, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *ModelSwitch) SetOspfConfig(v OspfConfig)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *ModelSwitch) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.

### GetOtherIpConfigs

`func (o *ModelSwitch) GetOtherIpConfigs() map[string]JunosOtherIpConfigs`

GetOtherIpConfigs returns the OtherIpConfigs field if non-nil, zero value otherwise.

### GetOtherIpConfigsOk

`func (o *ModelSwitch) GetOtherIpConfigsOk() (*map[string]JunosOtherIpConfigs, bool)`

GetOtherIpConfigsOk returns a tuple with the OtherIpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIpConfigs

`func (o *ModelSwitch) SetOtherIpConfigs(v map[string]JunosOtherIpConfigs)`

SetOtherIpConfigs sets OtherIpConfigs field to given value.

### HasOtherIpConfigs

`func (o *ModelSwitch) HasOtherIpConfigs() bool`

HasOtherIpConfigs returns a boolean if a field has been set.

### GetPortConfig

`func (o *ModelSwitch) GetPortConfig() map[string]JunosPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *ModelSwitch) GetPortConfigOk() (*map[string]JunosPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *ModelSwitch) SetPortConfig(v map[string]JunosPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *ModelSwitch) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPortMirroring

`func (o *ModelSwitch) GetPortMirroring() map[string]SwitchPortMirroring`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *ModelSwitch) GetPortMirroringOk() (*map[string]SwitchPortMirroring, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *ModelSwitch) SetPortMirroring(v map[string]SwitchPortMirroring)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *ModelSwitch) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetPortUsages

`func (o *ModelSwitch) GetPortUsages() map[string]SwitchPortUsage`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *ModelSwitch) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *ModelSwitch) SetPortUsages(v map[string]SwitchPortUsage)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *ModelSwitch) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *ModelSwitch) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *ModelSwitch) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *ModelSwitch) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *ModelSwitch) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRole

`func (o *ModelSwitch) GetRole() SwitchRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelSwitch) GetRoleOk() (*SwitchRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelSwitch) SetRole(v SwitchRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModelSwitch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRouterId

`func (o *ModelSwitch) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *ModelSwitch) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *ModelSwitch) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *ModelSwitch) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetSiteId

`func (o *ModelSwitch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ModelSwitch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ModelSwitch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ModelSwitch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStpConfig

`func (o *ModelSwitch) GetStpConfig() SwitchStpConfig`

GetStpConfig returns the StpConfig field if non-nil, zero value otherwise.

### GetStpConfigOk

`func (o *ModelSwitch) GetStpConfigOk() (*SwitchStpConfig, bool)`

GetStpConfigOk returns a tuple with the StpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpConfig

`func (o *ModelSwitch) SetStpConfig(v SwitchStpConfig)`

SetStpConfig sets StpConfig field to given value.

### HasStpConfig

`func (o *ModelSwitch) HasStpConfig() bool`

HasStpConfig returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *ModelSwitch) GetSwitchMgmt() SwitchSwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *ModelSwitch) GetSwitchMgmtOk() (*SwitchSwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *ModelSwitch) SetSwitchMgmt(v SwitchSwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *ModelSwitch) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetUseRouterIdAsSourceIp

`func (o *ModelSwitch) GetUseRouterIdAsSourceIp() bool`

GetUseRouterIdAsSourceIp returns the UseRouterIdAsSourceIp field if non-nil, zero value otherwise.

### GetUseRouterIdAsSourceIpOk

`func (o *ModelSwitch) GetUseRouterIdAsSourceIpOk() (*bool, bool)`

GetUseRouterIdAsSourceIpOk returns a tuple with the UseRouterIdAsSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRouterIdAsSourceIp

`func (o *ModelSwitch) SetUseRouterIdAsSourceIp(v bool)`

SetUseRouterIdAsSourceIp sets UseRouterIdAsSourceIp field to given value.

### HasUseRouterIdAsSourceIp

`func (o *ModelSwitch) HasUseRouterIdAsSourceIp() bool`

HasUseRouterIdAsSourceIp returns a boolean if a field has been set.

### GetVars

`func (o *ModelSwitch) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *ModelSwitch) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *ModelSwitch) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *ModelSwitch) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetVirtualChassis

`func (o *ModelSwitch) GetVirtualChassis() SwitchVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *ModelSwitch) GetVirtualChassisOk() (*SwitchVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *ModelSwitch) SetVirtualChassis(v SwitchVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *ModelSwitch) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### GetVrfConfig

`func (o *ModelSwitch) GetVrfConfig() VrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *ModelSwitch) GetVrfConfigOk() (*VrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *ModelSwitch) SetVrfConfig(v VrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *ModelSwitch) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrrpConfig

`func (o *ModelSwitch) GetVrrpConfig() VrrpConfig`

GetVrrpConfig returns the VrrpConfig field if non-nil, zero value otherwise.

### GetVrrpConfigOk

`func (o *ModelSwitch) GetVrrpConfigOk() (*VrrpConfig, bool)`

GetVrrpConfigOk returns a tuple with the VrrpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpConfig

`func (o *ModelSwitch) SetVrrpConfig(v VrrpConfig)`

SetVrrpConfig sets VrrpConfig field to given value.

### HasVrrpConfig

`func (o *ModelSwitch) HasVrrpConfig() bool`

HasVrrpConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


