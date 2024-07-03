# GatewayTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**BgpConfig** | Pointer to [**map[string]BgpConfig**](BgpConfig.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DhcpdConfig** | Pointer to [**DhcpdConfigs**](DhcpdConfigs.md) |  | [optional] 
**DnsOverride** | Pointer to **bool** |  | [optional] [default to false]
**DnsServers** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**DnsSuffix** | Pointer to **[]string** | Global dns settings. To keep compatibility, dns settings in &#x60;ip_config&#x60; and &#x60;oob_ip_config&#x60; will overwrite this setting | [optional] 
**ExtraRoutes** | Pointer to [**map[string]GatewayExtraRoute**](GatewayExtraRoute.md) |  | [optional] 
**GatewayMatching** | Pointer to [**GatewayMatching**](GatewayMatching.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IdpProfiles** | Pointer to [**map[string]IdpProfile**](IdpProfile.md) | Property key is the profile name | [optional] 
**IpConfigs** | Pointer to [**map[string]GatewayTemplateIpConfig**](GatewayTemplateIpConfig.md) | Property key is the network name | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**NtpOverride** | Pointer to **bool** |  | [optional] [default to false]
**NtpServers** | Pointer to **[]string** | list of NTP servers specific to this device. By default, those in Site Settings will be used | [optional] 
**OobIpConfig** | Pointer to [**JunosOobIpConfigs**](JunosOobIpConfigs.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PathPreferences** | Pointer to [**map[string]GatewayTemplatePathPreferences**](GatewayTemplatePathPreferences.md) | Property key is the path name | [optional] 
**PortConfig** | Pointer to [**map[string]GatewayPortConfig**](GatewayPortConfig.md) | Property key is the port(s) name or range (e.g. \&quot;ge-0/0/0-10\&quot;) | [optional] 
**RouterId** | Pointer to **string** | auto assigned if not set | [optional] 
**RoutingPolicies** | Pointer to [**map[string]RoutingPolicy**](RoutingPolicy.md) | Property key is the routing policy name | [optional] 
**ServicePolicies** | Pointer to [**[]ServicePolicy**](ServicePolicy.md) |  | [optional] 
**TunnelConfigs** | Pointer to [**map[string]TunnelConfigs**](TunnelConfigs.md) | Property key is the tunnel name | [optional] 
**TunnelProviderOptions** | Pointer to [**TunnelProviderOptions**](TunnelProviderOptions.md) |  | [optional] 
**Type** | Pointer to [**GatewayTemplateType**](GatewayTemplateType.md) |  | [optional] [default to GATEWAYTEMPLATETYPE_STANDALONE]

## Methods

### NewGatewayTemplate

`func NewGatewayTemplate(name string, ) *GatewayTemplate`

NewGatewayTemplate instantiates a new GatewayTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateWithDefaults

`func NewGatewayTemplateWithDefaults() *GatewayTemplate`

NewGatewayTemplateWithDefaults instantiates a new GatewayTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalConfigCmds

`func (o *GatewayTemplate) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *GatewayTemplate) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *GatewayTemplate) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *GatewayTemplate) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetBgpConfig

`func (o *GatewayTemplate) GetBgpConfig() map[string]BgpConfig`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *GatewayTemplate) GetBgpConfigOk() (*map[string]BgpConfig, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *GatewayTemplate) SetBgpConfig(v map[string]BgpConfig)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *GatewayTemplate) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetCreatedTime

`func (o *GatewayTemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GatewayTemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GatewayTemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *GatewayTemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDhcpdConfig

`func (o *GatewayTemplate) GetDhcpdConfig() DhcpdConfigs`

GetDhcpdConfig returns the DhcpdConfig field if non-nil, zero value otherwise.

### GetDhcpdConfigOk

`func (o *GatewayTemplate) GetDhcpdConfigOk() (*DhcpdConfigs, bool)`

GetDhcpdConfigOk returns a tuple with the DhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpdConfig

`func (o *GatewayTemplate) SetDhcpdConfig(v DhcpdConfigs)`

SetDhcpdConfig sets DhcpdConfig field to given value.

### HasDhcpdConfig

`func (o *GatewayTemplate) HasDhcpdConfig() bool`

HasDhcpdConfig returns a boolean if a field has been set.

### GetDnsOverride

`func (o *GatewayTemplate) GetDnsOverride() bool`

GetDnsOverride returns the DnsOverride field if non-nil, zero value otherwise.

### GetDnsOverrideOk

`func (o *GatewayTemplate) GetDnsOverrideOk() (*bool, bool)`

GetDnsOverrideOk returns a tuple with the DnsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOverride

`func (o *GatewayTemplate) SetDnsOverride(v bool)`

SetDnsOverride sets DnsOverride field to given value.

### HasDnsOverride

`func (o *GatewayTemplate) HasDnsOverride() bool`

HasDnsOverride returns a boolean if a field has been set.

### GetDnsServers

`func (o *GatewayTemplate) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *GatewayTemplate) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *GatewayTemplate) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *GatewayTemplate) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *GatewayTemplate) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *GatewayTemplate) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *GatewayTemplate) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *GatewayTemplate) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *GatewayTemplate) GetExtraRoutes() map[string]GatewayExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *GatewayTemplate) GetExtraRoutesOk() (*map[string]GatewayExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *GatewayTemplate) SetExtraRoutes(v map[string]GatewayExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *GatewayTemplate) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetGatewayMatching

`func (o *GatewayTemplate) GetGatewayMatching() GatewayMatching`

GetGatewayMatching returns the GatewayMatching field if non-nil, zero value otherwise.

### GetGatewayMatchingOk

`func (o *GatewayTemplate) GetGatewayMatchingOk() (*GatewayMatching, bool)`

GetGatewayMatchingOk returns a tuple with the GatewayMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMatching

`func (o *GatewayTemplate) SetGatewayMatching(v GatewayMatching)`

SetGatewayMatching sets GatewayMatching field to given value.

### HasGatewayMatching

`func (o *GatewayTemplate) HasGatewayMatching() bool`

HasGatewayMatching returns a boolean if a field has been set.

### GetId

`func (o *GatewayTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpProfiles

`func (o *GatewayTemplate) GetIdpProfiles() map[string]IdpProfile`

GetIdpProfiles returns the IdpProfiles field if non-nil, zero value otherwise.

### GetIdpProfilesOk

`func (o *GatewayTemplate) GetIdpProfilesOk() (*map[string]IdpProfile, bool)`

GetIdpProfilesOk returns a tuple with the IdpProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProfiles

`func (o *GatewayTemplate) SetIdpProfiles(v map[string]IdpProfile)`

SetIdpProfiles sets IdpProfiles field to given value.

### HasIdpProfiles

`func (o *GatewayTemplate) HasIdpProfiles() bool`

HasIdpProfiles returns a boolean if a field has been set.

### GetIpConfigs

`func (o *GatewayTemplate) GetIpConfigs() map[string]GatewayTemplateIpConfig`

GetIpConfigs returns the IpConfigs field if non-nil, zero value otherwise.

### GetIpConfigsOk

`func (o *GatewayTemplate) GetIpConfigsOk() (*map[string]GatewayTemplateIpConfig, bool)`

GetIpConfigsOk returns a tuple with the IpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfigs

`func (o *GatewayTemplate) SetIpConfigs(v map[string]GatewayTemplateIpConfig)`

SetIpConfigs sets IpConfigs field to given value.

### HasIpConfigs

`func (o *GatewayTemplate) HasIpConfigs() bool`

HasIpConfigs returns a boolean if a field has been set.

### GetModifiedTime

`func (o *GatewayTemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *GatewayTemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *GatewayTemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *GatewayTemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *GatewayTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *GatewayTemplate) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GatewayTemplate) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GatewayTemplate) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GatewayTemplate) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNtpOverride

`func (o *GatewayTemplate) GetNtpOverride() bool`

GetNtpOverride returns the NtpOverride field if non-nil, zero value otherwise.

### GetNtpOverrideOk

`func (o *GatewayTemplate) GetNtpOverrideOk() (*bool, bool)`

GetNtpOverrideOk returns a tuple with the NtpOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpOverride

`func (o *GatewayTemplate) SetNtpOverride(v bool)`

SetNtpOverride sets NtpOverride field to given value.

### HasNtpOverride

`func (o *GatewayTemplate) HasNtpOverride() bool`

HasNtpOverride returns a boolean if a field has been set.

### GetNtpServers

`func (o *GatewayTemplate) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *GatewayTemplate) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *GatewayTemplate) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *GatewayTemplate) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOobIpConfig

`func (o *GatewayTemplate) GetOobIpConfig() JunosOobIpConfigs`

GetOobIpConfig returns the OobIpConfig field if non-nil, zero value otherwise.

### GetOobIpConfigOk

`func (o *GatewayTemplate) GetOobIpConfigOk() (*JunosOobIpConfigs, bool)`

GetOobIpConfigOk returns a tuple with the OobIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpConfig

`func (o *GatewayTemplate) SetOobIpConfig(v JunosOobIpConfigs)`

SetOobIpConfig sets OobIpConfig field to given value.

### HasOobIpConfig

`func (o *GatewayTemplate) HasOobIpConfig() bool`

HasOobIpConfig returns a boolean if a field has been set.

### GetOrgId

`func (o *GatewayTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GatewayTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GatewayTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GatewayTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPathPreferences

`func (o *GatewayTemplate) GetPathPreferences() map[string]GatewayTemplatePathPreferences`

GetPathPreferences returns the PathPreferences field if non-nil, zero value otherwise.

### GetPathPreferencesOk

`func (o *GatewayTemplate) GetPathPreferencesOk() (*map[string]GatewayTemplatePathPreferences, bool)`

GetPathPreferencesOk returns a tuple with the PathPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPreferences

`func (o *GatewayTemplate) SetPathPreferences(v map[string]GatewayTemplatePathPreferences)`

SetPathPreferences sets PathPreferences field to given value.

### HasPathPreferences

`func (o *GatewayTemplate) HasPathPreferences() bool`

HasPathPreferences returns a boolean if a field has been set.

### GetPortConfig

`func (o *GatewayTemplate) GetPortConfig() map[string]GatewayPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *GatewayTemplate) GetPortConfigOk() (*map[string]GatewayPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *GatewayTemplate) SetPortConfig(v map[string]GatewayPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *GatewayTemplate) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.

### GetRouterId

`func (o *GatewayTemplate) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *GatewayTemplate) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *GatewayTemplate) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *GatewayTemplate) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetRoutingPolicies

`func (o *GatewayTemplate) GetRoutingPolicies() map[string]RoutingPolicy`

GetRoutingPolicies returns the RoutingPolicies field if non-nil, zero value otherwise.

### GetRoutingPoliciesOk

`func (o *GatewayTemplate) GetRoutingPoliciesOk() (*map[string]RoutingPolicy, bool)`

GetRoutingPoliciesOk returns a tuple with the RoutingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicies

`func (o *GatewayTemplate) SetRoutingPolicies(v map[string]RoutingPolicy)`

SetRoutingPolicies sets RoutingPolicies field to given value.

### HasRoutingPolicies

`func (o *GatewayTemplate) HasRoutingPolicies() bool`

HasRoutingPolicies returns a boolean if a field has been set.

### GetServicePolicies

`func (o *GatewayTemplate) GetServicePolicies() []ServicePolicy`

GetServicePolicies returns the ServicePolicies field if non-nil, zero value otherwise.

### GetServicePoliciesOk

`func (o *GatewayTemplate) GetServicePoliciesOk() (*[]ServicePolicy, bool)`

GetServicePoliciesOk returns a tuple with the ServicePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePolicies

`func (o *GatewayTemplate) SetServicePolicies(v []ServicePolicy)`

SetServicePolicies sets ServicePolicies field to given value.

### HasServicePolicies

`func (o *GatewayTemplate) HasServicePolicies() bool`

HasServicePolicies returns a boolean if a field has been set.

### GetTunnelConfigs

`func (o *GatewayTemplate) GetTunnelConfigs() map[string]TunnelConfigs`

GetTunnelConfigs returns the TunnelConfigs field if non-nil, zero value otherwise.

### GetTunnelConfigsOk

`func (o *GatewayTemplate) GetTunnelConfigsOk() (*map[string]TunnelConfigs, bool)`

GetTunnelConfigsOk returns a tuple with the TunnelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConfigs

`func (o *GatewayTemplate) SetTunnelConfigs(v map[string]TunnelConfigs)`

SetTunnelConfigs sets TunnelConfigs field to given value.

### HasTunnelConfigs

`func (o *GatewayTemplate) HasTunnelConfigs() bool`

HasTunnelConfigs returns a boolean if a field has been set.

### GetTunnelProviderOptions

`func (o *GatewayTemplate) GetTunnelProviderOptions() TunnelProviderOptions`

GetTunnelProviderOptions returns the TunnelProviderOptions field if non-nil, zero value otherwise.

### GetTunnelProviderOptionsOk

`func (o *GatewayTemplate) GetTunnelProviderOptionsOk() (*TunnelProviderOptions, bool)`

GetTunnelProviderOptionsOk returns a tuple with the TunnelProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelProviderOptions

`func (o *GatewayTemplate) SetTunnelProviderOptions(v TunnelProviderOptions)`

SetTunnelProviderOptions sets TunnelProviderOptions field to given value.

### HasTunnelProviderOptions

`func (o *GatewayTemplate) HasTunnelProviderOptions() bool`

HasTunnelProviderOptions returns a boolean if a field has been set.

### GetType

`func (o *GatewayTemplate) GetType() GatewayTemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTemplate) GetTypeOk() (*GatewayTemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTemplate) SetType(v GatewayTemplateType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayTemplate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


