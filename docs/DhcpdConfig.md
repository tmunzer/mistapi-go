# DhcpdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, if not defined, system one will be used | [optional] 
**DnsSuffix** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, if not defined, system one will be used | [optional] 
**FixedBindings** | Pointer to [**map[string]DhcpdConfigFixedBinding**](DhcpdConfigFixedBinding.md) | Property key is the MAC Address | [optional] 
**Gateway** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; - optional, &#x60;ip&#x60; will be used if not provided | [optional] 
**IpEnd** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**IpEnd6** | Pointer to **string** | if &#x60;type6&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**IpStart** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**IpStart6** | Pointer to **string** | if &#x60;type6&#x60;&#x3D;&#x3D;&#x60;local&#x60; | [optional] 
**LeaseTime** | Pointer to **int32** | in seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day] | [optional] [default to 86400]
**Options** | Pointer to [**map[string]DhcpdConfigOption**](DhcpdConfigOption.md) | Property key is the DHCP option number | [optional] 
**ServerIdOverride** | Pointer to **bool** | &#x60;server_id_override&#x60;&#x3D;&#x3D;&#x60;true&#x60; means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,  should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address. | [optional] [default to false]
**Servers** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;relay&#x60; | [optional] 
**Servers6** | Pointer to **[]string** | if &#x60;type6&#x60;&#x3D;&#x3D;&#x60;relay&#x60; | [optional] 
**Type** | Pointer to [**DhcpdConfigType**](DhcpdConfigType.md) |  | [optional] [default to DHCPDCONFIGTYPE_LOCAL]
**Type6** | Pointer to [**DhcpdConfigType6**](DhcpdConfigType6.md) |  | [optional] [default to DHCPDCONFIGTYPE6_NONE]
**VendorEncapulated** | Pointer to [**map[string]DhcpdConfigVendorOption**](DhcpdConfigVendorOption.md) | Property key is &lt;enterprise number&gt;:&lt;sub option code&gt;, with * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers) * sub option code: 1-255, sub-option code | [optional] 

## Methods

### NewDhcpdConfig

`func NewDhcpdConfig() *DhcpdConfig`

NewDhcpdConfig instantiates a new DhcpdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpdConfigWithDefaults

`func NewDhcpdConfigWithDefaults() *DhcpdConfig`

NewDhcpdConfigWithDefaults instantiates a new DhcpdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *DhcpdConfig) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *DhcpdConfig) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *DhcpdConfig) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *DhcpdConfig) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *DhcpdConfig) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *DhcpdConfig) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *DhcpdConfig) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *DhcpdConfig) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetFixedBindings

`func (o *DhcpdConfig) GetFixedBindings() map[string]DhcpdConfigFixedBinding`

GetFixedBindings returns the FixedBindings field if non-nil, zero value otherwise.

### GetFixedBindingsOk

`func (o *DhcpdConfig) GetFixedBindingsOk() (*map[string]DhcpdConfigFixedBinding, bool)`

GetFixedBindingsOk returns a tuple with the FixedBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBindings

`func (o *DhcpdConfig) SetFixedBindings(v map[string]DhcpdConfigFixedBinding)`

SetFixedBindings sets FixedBindings field to given value.

### HasFixedBindings

`func (o *DhcpdConfig) HasFixedBindings() bool`

HasFixedBindings returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpdConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpdConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpdConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpdConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpEnd

`func (o *DhcpdConfig) GetIpEnd() string`

GetIpEnd returns the IpEnd field if non-nil, zero value otherwise.

### GetIpEndOk

`func (o *DhcpdConfig) GetIpEndOk() (*string, bool)`

GetIpEndOk returns a tuple with the IpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEnd

`func (o *DhcpdConfig) SetIpEnd(v string)`

SetIpEnd sets IpEnd field to given value.

### HasIpEnd

`func (o *DhcpdConfig) HasIpEnd() bool`

HasIpEnd returns a boolean if a field has been set.

### GetIpEnd6

`func (o *DhcpdConfig) GetIpEnd6() string`

GetIpEnd6 returns the IpEnd6 field if non-nil, zero value otherwise.

### GetIpEnd6Ok

`func (o *DhcpdConfig) GetIpEnd6Ok() (*string, bool)`

GetIpEnd6Ok returns a tuple with the IpEnd6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEnd6

`func (o *DhcpdConfig) SetIpEnd6(v string)`

SetIpEnd6 sets IpEnd6 field to given value.

### HasIpEnd6

`func (o *DhcpdConfig) HasIpEnd6() bool`

HasIpEnd6 returns a boolean if a field has been set.

### GetIpStart

`func (o *DhcpdConfig) GetIpStart() string`

GetIpStart returns the IpStart field if non-nil, zero value otherwise.

### GetIpStartOk

`func (o *DhcpdConfig) GetIpStartOk() (*string, bool)`

GetIpStartOk returns a tuple with the IpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStart

`func (o *DhcpdConfig) SetIpStart(v string)`

SetIpStart sets IpStart field to given value.

### HasIpStart

`func (o *DhcpdConfig) HasIpStart() bool`

HasIpStart returns a boolean if a field has been set.

### GetIpStart6

`func (o *DhcpdConfig) GetIpStart6() string`

GetIpStart6 returns the IpStart6 field if non-nil, zero value otherwise.

### GetIpStart6Ok

`func (o *DhcpdConfig) GetIpStart6Ok() (*string, bool)`

GetIpStart6Ok returns a tuple with the IpStart6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStart6

`func (o *DhcpdConfig) SetIpStart6(v string)`

SetIpStart6 sets IpStart6 field to given value.

### HasIpStart6

`func (o *DhcpdConfig) HasIpStart6() bool`

HasIpStart6 returns a boolean if a field has been set.

### GetLeaseTime

`func (o *DhcpdConfig) GetLeaseTime() int32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *DhcpdConfig) GetLeaseTimeOk() (*int32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *DhcpdConfig) SetLeaseTime(v int32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *DhcpdConfig) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpdConfig) GetOptions() map[string]DhcpdConfigOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpdConfig) GetOptionsOk() (*map[string]DhcpdConfigOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpdConfig) SetOptions(v map[string]DhcpdConfigOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpdConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetServerIdOverride

`func (o *DhcpdConfig) GetServerIdOverride() bool`

GetServerIdOverride returns the ServerIdOverride field if non-nil, zero value otherwise.

### GetServerIdOverrideOk

`func (o *DhcpdConfig) GetServerIdOverrideOk() (*bool, bool)`

GetServerIdOverrideOk returns a tuple with the ServerIdOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIdOverride

`func (o *DhcpdConfig) SetServerIdOverride(v bool)`

SetServerIdOverride sets ServerIdOverride field to given value.

### HasServerIdOverride

`func (o *DhcpdConfig) HasServerIdOverride() bool`

HasServerIdOverride returns a boolean if a field has been set.

### GetServers

`func (o *DhcpdConfig) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DhcpdConfig) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DhcpdConfig) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DhcpdConfig) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetServers6

`func (o *DhcpdConfig) GetServers6() []string`

GetServers6 returns the Servers6 field if non-nil, zero value otherwise.

### GetServers6Ok

`func (o *DhcpdConfig) GetServers6Ok() (*[]string, bool)`

GetServers6Ok returns a tuple with the Servers6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers6

`func (o *DhcpdConfig) SetServers6(v []string)`

SetServers6 sets Servers6 field to given value.

### HasServers6

`func (o *DhcpdConfig) HasServers6() bool`

HasServers6 returns a boolean if a field has been set.

### GetType

`func (o *DhcpdConfig) GetType() DhcpdConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpdConfig) GetTypeOk() (*DhcpdConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpdConfig) SetType(v DhcpdConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpdConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType6

`func (o *DhcpdConfig) GetType6() DhcpdConfigType6`

GetType6 returns the Type6 field if non-nil, zero value otherwise.

### GetType6Ok

`func (o *DhcpdConfig) GetType6Ok() (*DhcpdConfigType6, bool)`

GetType6Ok returns a tuple with the Type6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType6

`func (o *DhcpdConfig) SetType6(v DhcpdConfigType6)`

SetType6 sets Type6 field to given value.

### HasType6

`func (o *DhcpdConfig) HasType6() bool`

HasType6 returns a boolean if a field has been set.

### GetVendorEncapulated

`func (o *DhcpdConfig) GetVendorEncapulated() map[string]DhcpdConfigVendorOption`

GetVendorEncapulated returns the VendorEncapulated field if non-nil, zero value otherwise.

### GetVendorEncapulatedOk

`func (o *DhcpdConfig) GetVendorEncapulatedOk() (*map[string]DhcpdConfigVendorOption, bool)`

GetVendorEncapulatedOk returns a tuple with the VendorEncapulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorEncapulated

`func (o *DhcpdConfig) SetVendorEncapulated(v map[string]DhcpdConfigVendorOption)`

SetVendorEncapulated sets VendorEncapulated field to given value.

### HasVendorEncapulated

`func (o *DhcpdConfig) HasVendorEncapulated() bool`

HasVendorEncapulated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


