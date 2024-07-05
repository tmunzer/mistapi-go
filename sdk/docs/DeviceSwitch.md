# DeviceSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclPolicies** | Pointer to [**[]AclPolicy**](AclPolicy.md) |  | [optional] 
**AclTags** | Pointer to [**map[string]AclTag**](AclTag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileId** | Pointer to **string** |  | [optional] [readonly] 
**DhcpSnooping** | Pointer to [**DhcpSnooping**](DhcpSnooping.md) |  | [optional] 
**DhcpdConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
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
**MapId** | Pointer to **string** | map where the device belongs to | [optional] 
**MistNac** | Pointer to [**SwitchMistNac**](SwitchMistNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to [**map[string]SwitchNetwork**](SwitchNetwork.md) | Property key is network name | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] 
**OobIpConfig** | Pointer to [**SwitchOobIpConfig**](SwitchOobIpConfig.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**OspfConfig** | Pointer to [**OspfConfig**](OspfConfig.md) |  | [optional] 
**OtherIpConfigs** | Pointer to [**map[string]JunosOtherIpConfig**](JunosOtherIpConfig.md) | Property key is the network name | [optional] 
**PortConfig** | Pointer to [**map[string]JunosPortConfig**](JunosPortConfig.md) | Property key is the port name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**PortMirroring** | Pointer to [**map[string]SwitchPortMirroringProperty**](SwitchPortMirroringProperty.md) | Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. | [optional] 
**PortUsages** | Pointer to [**map[string]SwitchPortUsage**](SwitchPortUsage.md) |  | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**RemoteSyslog** | Pointer to [**RemoteSyslog**](RemoteSyslog.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**RouterId** | Pointer to **string** | used for OSPF / BGP / EVPN | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SnmpConfig** | Pointer to [**SnmpConfig**](SnmpConfig.md) |  | [optional] 
**StpConfig** | Pointer to [**SwitchStpConfig**](SwitchStpConfig.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchMgmt**](SwitchMgmt.md) |  | [optional] 
**UseRouterIdAsSourceIp** | Pointer to **bool** | whether to use it for snmp / syslog / tacplus / radius | [optional] [default to false]
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 
**VirtualChassis** | Pointer to [**SwitchVirtualChassis**](SwitchVirtualChassis.md) |  | [optional] 
**VrfConfig** | Pointer to [**VrfConfig**](VrfConfig.md) |  | [optional] 
**VrfInstances** | Pointer to [**map[string]VrfInstance**](VrfInstance.md) | Property key is the network name | [optional] 
**VrrpConfig** | Pointer to [**VrrpConfig**](VrrpConfig.md) |  | [optional] 
**X** | Pointer to **float64** | x in pixel | [optional] 
**Y** | Pointer to **float64** | y in pixel | [optional] 

## Methods

### NewDeviceSwitch

`func NewDeviceSwitch() *DeviceSwitch`

NewDeviceSwitch instantiates a new DeviceSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSwitchWithDefaults

`func NewDeviceSwitchWithDefaults() *DeviceSwitch`

NewDeviceSwitchWithDefaults instantiates a new DeviceSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclPolicies

`func (o *DeviceSwitch) GetAclPolicies() []AclPolicy`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *DeviceSwitch) GetAclPoliciesOk() (*[]AclPolicy, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *DeviceSwitch) SetAclPolicies(v []AclPolicy)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *DeviceSwitch) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *DeviceSwitch) GetAclTags() map[string]AclTag`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *DeviceSwitch) GetAclTagsOk() (*map[string]AclTag, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *DeviceSwitch) SetAclTags(v map[string]AclTag)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *DeviceSwitch) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *DeviceSwitch) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *DeviceSwitch) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *DeviceSwitch) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *DeviceSwitch) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DeviceSwitch) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DeviceSwitch) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DeviceSwitch) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DeviceSwitch) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileId

`func (o *DeviceSwitch) GetDeviceprofileId() string`

GetDeviceprofileId returns the DeviceprofileId field if non-nil, zero value otherwise.

### GetDeviceprofileIdOk

`func (o *DeviceSwitch) GetDeviceprofileIdOk() (*string, bool)`

GetDeviceprofileIdOk returns a tuple with the DeviceprofileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileId

`func (o *DeviceSwitch) SetDeviceprofileId(v string)`

SetDeviceprofileId sets DeviceprofileId field to given value.

### HasDeviceprofileId

`func (o *DeviceSwitch) HasDeviceprofileId() bool`

HasDeviceprofileId returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *DeviceSwitch) GetDhcpSnooping() DhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *DeviceSwitch) GetDhcpSnoopingOk() (*DhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *DeviceSwitch) SetDhcpSnooping(v DhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *DeviceSwitch) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *DeviceSwitch) GetDhcpdConfig() DhcpdConfigs`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *DeviceSwitch) GetDhcpdConfigOk() (*DhcpdConfigs, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *DeviceSwitch) SetDhcpdConfig(v DhcpdConfigs)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *DeviceSwitch) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetDisableAutoConfig

`func (o *DeviceSwitch) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *DeviceSwitch) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *DeviceSwitch) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *DeviceSwitch) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetDnsServers

`func (o *DeviceSwitch) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *DeviceSwitch) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *DeviceSwitch) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *DeviceSwitch) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *DeviceSwitch) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *DeviceSwitch) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *DeviceSwitch) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *DeviceSwitch) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEvpnConfig

`func (o *DeviceSwitch) GetEvpnConfig() EvpnConfig`

GetEvpnConfig returns the EvpnConfig field if non-nil, zero value otherwise.

### GetEvpnConfigOk

`func (o *DeviceSwitch) GetEvpnConfigOk() (*EvpnConfig, bool)`

GetEvpnConfigOk returns a tuple with the EvpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnConfig

`func (o *DeviceSwitch) SetEvpnConfig(v EvpnConfig)`

SetEvpnConfig sets EvpnConfig field to given value.

### HasEvpnConfig

`func (o *DeviceSwitch) HasEvpnConfig() bool`

HasEvpnConfig returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *DeviceSwitch) GetExtraRoutes() map[string]ExtraRouteProperties`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *DeviceSwitch) GetExtraRoutesOk() (*map[string]ExtraRouteProperties, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *DeviceSwitch) SetExtraRoutes(v map[string]ExtraRouteProperties)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *DeviceSwitch) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *DeviceSwitch) GetExtraRoutes6() map[string]ExtraRoute6Properties`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *DeviceSwitch) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *DeviceSwitch) SetExtraRoutes6(v map[string]ExtraRoute6Properties)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *DeviceSwitch) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetId

`func (o *DeviceSwitch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceSwitch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceSwitch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceSwitch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage1Url

`func (o *DeviceSwitch) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *DeviceSwitch) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *DeviceSwitch) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *DeviceSwitch) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### SetImage1UrlNil

`func (o *DeviceSwitch) SetImage1UrlNil(b bool)`

 SetImage1UrlNil sets the value for Image1Url to be an explicit nil

### UnsetImage1Url
`func (o *DeviceSwitch) UnsetImage1Url()`

UnsetImage1Url ensures that no value is present for Image1Url, not even an explicit nil
### GetImage2Url

`func (o *DeviceSwitch) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *DeviceSwitch) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *DeviceSwitch) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *DeviceSwitch) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.

### SetImage2UrlNil

`func (o *DeviceSwitch) SetImage2UrlNil(b bool)`

 SetImage2UrlNil sets the value for Image2Url to be an explicit nil

### UnsetImage2Url
`func (o *DeviceSwitch) UnsetImage2Url()`

UnsetImage2Url ensures that no value is present for Image2Url, not even an explicit nil
### GetImage3Url

`func (o *DeviceSwitch) GetImage3Url() string`

GetImage3Url returns the Image3Url field if non-nil, zero value otherwise.

### GetImage3UrlOk

`func (o *DeviceSwitch) GetImage3UrlOk() (*string, bool)`

GetImage3UrlOk returns a tuple with the Image3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage3Url

`func (o *DeviceSwitch) SetImage3Url(v string)`

SetImage3Url sets Image3Url field to given value.

### HasImage3Url

`func (o *DeviceSwitch) HasImage3Url() bool`

HasImage3Url returns a boolean if a field has been set.

### SetImage3UrlNil

`func (o *DeviceSwitch) SetImage3UrlNil(b bool)`

 SetImage3UrlNil sets the value for Image3Url to be an explicit nil

### UnsetImage3Url
`func (o *DeviceSwitch) UnsetImage3Url()`

UnsetImage3Url ensures that no value is present for Image3Url, not even an explicit nil
### GetIpConfig

`func (o *DeviceSwitch) GetIpConfig() JunosIpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *DeviceSwitch) GetIpConfigOk() (*JunosIpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *DeviceSwitch) SetIpConfig(v JunosIpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *DeviceSwitch) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetManaged

`func (o *DeviceSwitch) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceSwitch) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceSwitch) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeviceSwitch) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetMapId

`func (o *DeviceSwitch) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *DeviceSwitch) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *DeviceSwitch) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *DeviceSwitch) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetMistNac

`func (o *DeviceSwitch) GetMistNac() SwitchMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *DeviceSwitch) GetMistNacOk() (*SwitchMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *DeviceSwitch) SetMistNac(v SwitchMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *DeviceSwitch) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *DeviceSwitch) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *DeviceSwitch) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *DeviceSwitch) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *DeviceSwitch) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *DeviceSwitch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceSwitch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceSwitch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceSwitch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *DeviceSwitch) GetNetworks() map[string]SwitchNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DeviceSwitch) GetNetworksOk() (*map[string]SwitchNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DeviceSwitch) SetNetworks(v map[string]SwitchNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *DeviceSwitch) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNotes

`func (o *DeviceSwitch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeviceSwitch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeviceSwitch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DeviceSwitch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNtpServers

`func (o *DeviceSwitch) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *DeviceSwitch) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *DeviceSwitch) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *DeviceSwitch) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *DeviceSwitch) GetOobIpConfig() SwitchOobIpConfig`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *DeviceSwitch) GetOobIpConfigOk() (*SwitchOobIpConfig, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *DeviceSwitch) SetOobIpConfig(v SwitchOobIpConfig)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *DeviceSwitch) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *DeviceSwitch) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DeviceSwitch) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DeviceSwitch) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DeviceSwitch) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOspfConfig

`func (o *DeviceSwitch) GetOspfConfig() OspfConfig`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *DeviceSwitch) GetOspfConfigOk() (*OspfConfig, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *DeviceSwitch) SetOspfConfig(v OspfConfig)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *DeviceSwitch) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.

### GetOtherIpConfigs

`func (o *DeviceSwitch) GetOtherIpConfigs() map[string]JunosOtherIpConfig`

GetOtherIpConfigs returns the OtherIpConfigs field if non-nil, zero value otherwise.

### GetOtherIpConfigsOk

`func (o *DeviceSwitch) GetOtherIpConfigsOk() (*map[string]JunosOtherIpConfig, bool)`

GetOtherIpConfigsOk returns a tuple with the OtherIpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIpConfigs

`func (o *DeviceSwitch) SetOtherIpConfigs(v map[string]JunosOtherIpConfig)`

SetOtherIpConfigs sets OtherIpConfigs field to given value.

### HasOtherIpConfigs

`func (o *DeviceSwitch) HasOtherIpConfigs() bool`

HasOtherIpConfigs returns a boolean if a field has been set.

### GetPortConfig

`func (o *DeviceSwitch) GetPortConfig() map[string]JunosPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *DeviceSwitch) GetPortConfigOk() (*map[string]JunosPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *DeviceSwitch) SetPortConfig(v map[string]JunosPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *DeviceSwitch) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetPortMirroring

`func (o *DeviceSwitch) GetPortMirroring() map[string]SwitchPortMirroringProperty`

GetPortMirroring returns the PortMirroring field if non-nil, zero value otherwise.

### GetPortMirroringOk

`func (o *DeviceSwitch) GetPortMirroringOk() (*map[string]SwitchPortMirroringProperty, bool)`

GetPortMirroringOk returns a tuple with the PortMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirroring

`func (o *DeviceSwitch) SetPortMirroring(v map[string]SwitchPortMirroringProperty)`

SetPortMirroring sets PortMirroring field to given value.

### HasPortMirroring

`func (o *DeviceSwitch) HasPortMirroring() bool`

HasPortMirroring returns a boolean if a field has been set.

### GetPortUsages

`func (o *DeviceSwitch) GetPortUsages() map[string]SwitchPortUsage`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *DeviceSwitch) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *DeviceSwitch) SetPortUsages(v map[string]SwitchPortUsage)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *DeviceSwitch) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *DeviceSwitch) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *DeviceSwitch) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *DeviceSwitch) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *DeviceSwitch) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRemoteSyslog

`func (o *DeviceSwitch) GetRemoteSyslog() RemoteSyslog`

GetRemoteSyslog returns the RemoteSyslog field if non-nil, zero value otherwise.

### GetRemoteSyslogOk

`func (o *DeviceSwitch) GetRemoteSyslogOk() (*RemoteSyslog, bool)`

GetRemoteSyslogOk returns a tuple with the RemoteSyslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslog

`func (o *DeviceSwitch) SetRemoteSyslog(v RemoteSyslog)`

SetRemoteSyslog sets RemoteSyslog field to given value.

### HasRemoteSyslog

`func (o *DeviceSwitch) HasRemoteSyslog() bool`

HasRemoteSyslog returns a boolean if a field has been set.

### GetRole

`func (o *DeviceSwitch) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DeviceSwitch) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DeviceSwitch) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DeviceSwitch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRouterId

`func (o *DeviceSwitch) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *DeviceSwitch) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *DeviceSwitch) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *DeviceSwitch) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetSiteId

`func (o *DeviceSwitch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceSwitch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceSwitch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DeviceSwitch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSnmpConfig

`func (o *DeviceSwitch) GetSnmpConfig() SnmpConfig`

GetSnmpConfig returns the SnmpConfig field if non-nil, zero value otherwise.

### GetSnmpConfigOk

`func (o *DeviceSwitch) GetSnmpConfigOk() (*SnmpConfig, bool)`

GetSnmpConfigOk returns a tuple with the SnmpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpConfig

`func (o *DeviceSwitch) SetSnmpConfig(v SnmpConfig)`

SetSnmpConfig sets SnmpConfig field to given value.

### HasSnmpConfig

`func (o *DeviceSwitch) HasSnmpConfig() bool`

HasSnmpConfig returns a boolean if a field has been set.

### GetStpConfig

`func (o *DeviceSwitch) GetStpConfig() SwitchStpConfig`

GetStpConfig returns the StpConfig field if non-nil, zero value otherwise.

### GetStpConfigOk

`func (o *DeviceSwitch) GetStpConfigOk() (*SwitchStpConfig, bool)`

GetStpConfigOk returns a tuple with the StpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpConfig

`func (o *DeviceSwitch) SetStpConfig(v SwitchStpConfig)`

SetStpConfig sets StpConfig field to given value.

### HasStpConfig

`func (o *DeviceSwitch) HasStpConfig() bool`

HasStpConfig returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *DeviceSwitch) GetSwitchMgmt() SwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *DeviceSwitch) GetSwitchMgmtOk() (*SwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *DeviceSwitch) SetSwitchMgmt(v SwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *DeviceSwitch) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetUseRouterIdAsSourceIp

`func (o *DeviceSwitch) GetUseRouterIdAsSourceIp() bool`

GetUseRouterIdAsSourceIp returns the UseRouterIdAsSourceIp field if non-nil, zero value otherwise.

### GetUseRouterIdAsSourceIpOk

`func (o *DeviceSwitch) GetUseRouterIdAsSourceIpOk() (*bool, bool)`

GetUseRouterIdAsSourceIpOk returns a tuple with the UseRouterIdAsSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRouterIdAsSourceIp

`func (o *DeviceSwitch) SetUseRouterIdAsSourceIp(v bool)`

SetUseRouterIdAsSourceIp sets UseRouterIdAsSourceIp field to given value.

### HasUseRouterIdAsSourceIp

`func (o *DeviceSwitch) HasUseRouterIdAsSourceIp() bool`

HasUseRouterIdAsSourceIp returns a boolean if a field has been set.

### GetVars

`func (o *DeviceSwitch) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *DeviceSwitch) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *DeviceSwitch) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *DeviceSwitch) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetVirtualChassis

`func (o *DeviceSwitch) GetVirtualChassis() SwitchVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *DeviceSwitch) GetVirtualChassisOk() (*SwitchVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *DeviceSwitch) SetVirtualChassis(v SwitchVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *DeviceSwitch) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### GetVrfConfig

`func (o *DeviceSwitch) GetVrfConfig() VrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *DeviceSwitch) GetVrfConfigOk() (*VrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *DeviceSwitch) SetVrfConfig(v VrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *DeviceSwitch) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrfInstances

`func (o *DeviceSwitch) GetVrfInstances() map[string]VrfInstance`

GetVrfInstances returns the VrfInstances field if non-nil, zero value otherwise.

### GetVrfInstancesOk

`func (o *DeviceSwitch) GetVrfInstancesOk() (*map[string]VrfInstance, bool)`

GetVrfInstancesOk returns a tuple with the VrfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfInstances

`func (o *DeviceSwitch) SetVrfInstances(v map[string]VrfInstance)`

SetVrfInstances sets VrfInstances field to given value.

### HasVrfInstances

`func (o *DeviceSwitch) HasVrfInstances() bool`

HasVrfInstances returns a boolean if a field has been set.

### GetVrrpConfig

`func (o *DeviceSwitch) GetVrrpConfig() VrrpConfig`

GetVrrpConfig returns the VrrpConfig field if non-nil, zero value otherwise.

### GetVrrpConfigOk

`func (o *DeviceSwitch) GetVrrpConfigOk() (*VrrpConfig, bool)`

GetVrrpConfigOk returns a tuple with the VrrpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpConfig

`func (o *DeviceSwitch) SetVrrpConfig(v VrrpConfig)`

SetVrrpConfig sets VrrpConfig field to given value.

### HasVrrpConfig

`func (o *DeviceSwitch) HasVrrpConfig() bool`

HasVrrpConfig returns a boolean if a field has been set.

### GetX

`func (o *DeviceSwitch) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DeviceSwitch) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DeviceSwitch) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *DeviceSwitch) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *DeviceSwitch) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DeviceSwitch) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DeviceSwitch) SetY(v float64)`

SetY sets Y field to given value.

### HasY

`func (o *DeviceSwitch) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


