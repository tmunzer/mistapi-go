# NetworkTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclPolicies** | Pointer to [**[]AclPolicy**](AclPolicy.md) |  | [optional] 
**AclTags** | Pointer to [**map[string]AclTag**](AclTag.md) | ACL Tags to identify traffic source or destination. Key name is the tag name | [optional] 
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DhcpSnooping** | Pointer to [**DhcpSnooping**](DhcpSnooping.md) |  | [optional] 
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**ExtraRoutes** | Pointer to [**map[string]ExtraRouteProperties**](ExtraRouteProperties.md) |  | [optional] 
**ExtraRoutes6** | Pointer to [**map[string]ExtraRoute6Properties**](ExtraRoute6Properties.md) | Property key is the destination CIDR (e.g. \&quot;2a02:1234:420a:10c9::/64\&quot;) | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ImportOrgNetworks** | Pointer to **[]string** | Org Networks that we&#39;d like to import | [optional] 
**MistNac** | Pointer to [**NetworkTemplateMistNac**](NetworkTemplateMistNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to [**map[string]SwitchNetwork**](SwitchNetwork.md) | Property key is network name | [optional] 
**NtpServers** | Pointer to **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PortMirrorings** | Pointer to [**map[string]SwitchPortMirroring**](SwitchPortMirroring.md) | Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. | [optional] 
**PortUsages** | Pointer to [**map[string]SwitchPortUsage**](SwitchPortUsage.md) |  | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**RemoteSyslog** | Pointer to [**RemoteSyslog**](RemoteSyslog.md) |  | [optional] 
**SnmpConfig** | Pointer to [**SnmpConfig**](SnmpConfig.md) |  | [optional] 
**SwitchMatching** | Pointer to [**SwitchMatching**](SwitchMatching.md) |  | [optional] 
**SwitchMgmt** | Pointer to [**SwitchMgmt**](SwitchMgmt.md) |  | [optional] 
**VrfConfig** | Pointer to [**VrfConfig**](VrfConfig.md) |  | [optional] 
**VrfInstances** | Pointer to [**map[string]VrfInstance**](VrfInstance.md) | Property key is the network name | [optional] 

## Methods

### NewNetworkTemplate

`func NewNetworkTemplate() *NetworkTemplate`

NewNetworkTemplate instantiates a new NetworkTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTemplateWithDefaults

`func NewNetworkTemplateWithDefaults() *NetworkTemplate`

NewNetworkTemplateWithDefaults instantiates a new NetworkTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclPolicies

`func (o *NetworkTemplate) GetAclPolicies() []AclPolicy`

GetAclPolicies returns the AclPolicies field if non-nil, zero value otherwise.

### GetAclPoliciesOk

`func (o *NetworkTemplate) GetAclPoliciesOk() (*[]AclPolicy, bool)`

GetAclPoliciesOk returns a tuple with the AclPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclPolicies

`func (o *NetworkTemplate) SetAclPolicies(v []AclPolicy)`

SetAclPolicies sets AclPolicies field to given value.

### HasAclPolicies

`func (o *NetworkTemplate) HasAclPolicies() bool`

HasAclPolicies returns a boolean if a field has been set.

### GetAclTags

`func (o *NetworkTemplate) GetAclTags() map[string]AclTag`

GetAclTags returns the AclTags field if non-nil, zero value otherwise.

### GetAclTagsOk

`func (o *NetworkTemplate) GetAclTagsOk() (*map[string]AclTag, bool)`

GetAclTagsOk returns a tuple with the AclTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTags

`func (o *NetworkTemplate) SetAclTags(v map[string]AclTag)`

SetAclTags sets AclTags field to given value.

### HasAclTags

`func (o *NetworkTemplate) HasAclTags() bool`

HasAclTags returns a boolean if a field has been set.

### GetAdditionalConfigCmds

`func (o *NetworkTemplate) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *NetworkTemplate) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *NetworkTemplate) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *NetworkTemplate) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NetworkTemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NetworkTemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NetworkTemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NetworkTemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDhcpSnooping

`func (o *NetworkTemplate) GetDhcpSnooping() DhcpSnooping`

GetDhcpSnooping returns the DhcpSnooping field if non-nil, zero value otherwise.

### GetDhcpSnoopingOk

`func (o *NetworkTemplate) GetDhcpSnoopingOk() (*DhcpSnooping, bool)`

GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSnooping

`func (o *NetworkTemplate) SetDhcpSnooping(v DhcpSnooping)`

SetDhcpSnooping sets DhcpSnooping field to given value.

### HasDhcpSnooping

`func (o *NetworkTemplate) HasDhcpSnooping() bool`

HasDhcpSnooping returns a boolean if a field has been set.

### GetDnsServers

`func (o *NetworkTemplate) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *NetworkTemplate) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *NetworkTemplate) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *NetworkTemplate) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *NetworkTemplate) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *NetworkTemplate) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *NetworkTemplate) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *NetworkTemplate) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *NetworkTemplate) GetExtraRoutes() map[string]ExtraRouteProperties`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *NetworkTemplate) GetExtraRoutesOk() (*map[string]ExtraRouteProperties, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *NetworkTemplate) SetExtraRoutes(v map[string]ExtraRouteProperties)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *NetworkTemplate) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetExtraRoutes6

`func (o *NetworkTemplate) GetExtraRoutes6() map[string]ExtraRoute6Properties`

GetExtraRoutes6 returns the ExtraRoutes6 field if non-nil, zero value otherwise.

### GetExtraRoutes6Ok

`func (o *NetworkTemplate) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool)`

GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes6

`func (o *NetworkTemplate) SetExtraRoutes6(v map[string]ExtraRoute6Properties)`

SetExtraRoutes6 sets ExtraRoutes6 field to given value.

### HasExtraRoutes6

`func (o *NetworkTemplate) HasExtraRoutes6() bool`

HasExtraRoutes6 returns a boolean if a field has been set.

### GetId

`func (o *NetworkTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImportOrgNetworks

`func (o *NetworkTemplate) GetImportOrgNetworks() []string`

GetImportOrgNetworks returns the ImportOrgNetworks field if non-nil, zero value otherwise.

### GetImportOrgNetworksOk

`func (o *NetworkTemplate) GetImportOrgNetworksOk() (*[]string, bool)`

GetImportOrgNetworksOk returns a tuple with the ImportOrgNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOrgNetworks

`func (o *NetworkTemplate) SetImportOrgNetworks(v []string)`

SetImportOrgNetworks sets ImportOrgNetworks field to given value.

### HasImportOrgNetworks

`func (o *NetworkTemplate) HasImportOrgNetworks() bool`

HasImportOrgNetworks returns a boolean if a field has been set.

### GetMistNac

`func (o *NetworkTemplate) GetMistNac() NetworkTemplateMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *NetworkTemplate) GetMistNacOk() (*NetworkTemplateMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *NetworkTemplate) SetMistNac(v NetworkTemplateMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *NetworkTemplate) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *NetworkTemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *NetworkTemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *NetworkTemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *NetworkTemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *NetworkTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *NetworkTemplate) GetNetworks() map[string]SwitchNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkTemplate) GetNetworksOk() (*map[string]SwitchNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkTemplate) SetNetworks(v map[string]SwitchNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkTemplate) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNtpServers

`func (o *NetworkTemplate) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *NetworkTemplate) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *NetworkTemplate) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *NetworkTemplate) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOrgId

`func (o *NetworkTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NetworkTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NetworkTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NetworkTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPortMirrorings

`func (o *NetworkTemplate) GetPortMirrorings() map[string]SwitchPortMirroring`

GetPortMirrorings returns the PortMirrorings field if non-nil, zero value otherwise.

### GetPortMirroringsOk

`func (o *NetworkTemplate) GetPortMirroringsOk() (*map[string]SwitchPortMirroring, bool)`

GetPortMirroringsOk returns a tuple with the PortMirrorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMirrorings

`func (o *NetworkTemplate) SetPortMirrorings(v map[string]SwitchPortMirroring)`

SetPortMirrorings sets PortMirrorings field to given value.

### HasPortMirrorings

`func (o *NetworkTemplate) HasPortMirrorings() bool`

HasPortMirrorings returns a boolean if a field has been set.

### GetPortUsages

`func (o *NetworkTemplate) GetPortUsages() map[string]SwitchPortUsage`

GetPortUsages returns the PortUsages field if non-nil, zero value otherwise.

### GetPortUsagesOk

`func (o *NetworkTemplate) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool)`

GetPortUsagesOk returns a tuple with the PortUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsages

`func (o *NetworkTemplate) SetPortUsages(v map[string]SwitchPortUsage)`

SetPortUsages sets PortUsages field to given value.

### HasPortUsages

`func (o *NetworkTemplate) HasPortUsages() bool`

HasPortUsages returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *NetworkTemplate) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *NetworkTemplate) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *NetworkTemplate) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *NetworkTemplate) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRemoteSyslog

`func (o *NetworkTemplate) GetRemoteSyslog() RemoteSyslog`

GetRemoteSyslog returns the RemoteSyslog field if non-nil, zero value otherwise.

### GetRemoteSyslogOk

`func (o *NetworkTemplate) GetRemoteSyslogOk() (*RemoteSyslog, bool)`

GetRemoteSyslogOk returns a tuple with the RemoteSyslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSyslog

`func (o *NetworkTemplate) SetRemoteSyslog(v RemoteSyslog)`

SetRemoteSyslog sets RemoteSyslog field to given value.

### HasRemoteSyslog

`func (o *NetworkTemplate) HasRemoteSyslog() bool`

HasRemoteSyslog returns a boolean if a field has been set.

### GetSnmpConfig

`func (o *NetworkTemplate) GetSnmpConfig() SnmpConfig`

GetSnmpConfig returns the SnmpConfig field if non-nil, zero value otherwise.

### GetSnmpConfigOk

`func (o *NetworkTemplate) GetSnmpConfigOk() (*SnmpConfig, bool)`

GetSnmpConfigOk returns a tuple with the SnmpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpConfig

`func (o *NetworkTemplate) SetSnmpConfig(v SnmpConfig)`

SetSnmpConfig sets SnmpConfig field to given value.

### HasSnmpConfig

`func (o *NetworkTemplate) HasSnmpConfig() bool`

HasSnmpConfig returns a boolean if a field has been set.

### GetSwitchMatching

`func (o *NetworkTemplate) GetSwitchMatching() SwitchMatching`

GetSwitchMatching returns the SwitchMatching field if non-nil, zero value otherwise.

### GetSwitchMatchingOk

`func (o *NetworkTemplate) GetSwitchMatchingOk() (*SwitchMatching, bool)`

GetSwitchMatchingOk returns a tuple with the SwitchMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMatching

`func (o *NetworkTemplate) SetSwitchMatching(v SwitchMatching)`

SetSwitchMatching sets SwitchMatching field to given value.

### HasSwitchMatching

`func (o *NetworkTemplate) HasSwitchMatching() bool`

HasSwitchMatching returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *NetworkTemplate) GetSwitchMgmt() SwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *NetworkTemplate) GetSwitchMgmtOk() (*SwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *NetworkTemplate) SetSwitchMgmt(v SwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *NetworkTemplate) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetVrfConfig

`func (o *NetworkTemplate) GetVrfConfig() VrfConfig`

GetVrfConfig returns the VrfConfig field if non-nil, zero value otherwise.

### GetVrfConfigOk

`func (o *NetworkTemplate) GetVrfConfigOk() (*VrfConfig, bool)`

GetVrfConfigOk returns a tuple with the VrfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfConfig

`func (o *NetworkTemplate) SetVrfConfig(v VrfConfig)`

SetVrfConfig sets VrfConfig field to given value.

### HasVrfConfig

`func (o *NetworkTemplate) HasVrfConfig() bool`

HasVrfConfig returns a boolean if a field has been set.

### GetVrfInstances

`func (o *NetworkTemplate) GetVrfInstances() map[string]VrfInstance`

GetVrfInstances returns the VrfInstances field if non-nil, zero value otherwise.

### GetVrfInstancesOk

`func (o *NetworkTemplate) GetVrfInstancesOk() (*map[string]VrfInstance, bool)`

GetVrfInstancesOk returns a tuple with the VrfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfInstances

`func (o *NetworkTemplate) SetVrfInstances(v map[string]VrfInstance)`

SetVrfInstances sets VrfInstances field to given value.

### HasVrfInstances

`func (o *NetworkTemplate) HasVrfInstances() bool`

HasVrfInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


