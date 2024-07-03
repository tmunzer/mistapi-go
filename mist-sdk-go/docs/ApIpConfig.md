# ApIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**DnsSuffix** | Pointer to **[]string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Gateway** | Pointer to **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Ip6** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Netmask** | Pointer to **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Netmask6** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**IpType**](IpType.md) |  | [optional] [default to IPTYPE_DHCP]
**Type6** | Pointer to [**IpType6**](IpType6.md) |  | [optional] [default to IPTYPE6_DISABLED]
**VlanId** | Pointer to **int32** | management vlan id, default is 1 (untagged) | [optional] [default to 1]

## Methods

### NewApIpConfig

`func NewApIpConfig() *ApIpConfig`

NewApIpConfig instantiates a new ApIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApIpConfigWithDefaults

`func NewApIpConfigWithDefaults() *ApIpConfig`

NewApIpConfigWithDefaults instantiates a new ApIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *ApIpConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ApIpConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ApIpConfig) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *ApIpConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *ApIpConfig) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *ApIpConfig) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *ApIpConfig) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *ApIpConfig) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *ApIpConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ApIpConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ApIpConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ApIpConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *ApIpConfig) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *ApIpConfig) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *ApIpConfig) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *ApIpConfig) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetIp

`func (o *ApIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *ApIpConfig) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *ApIpConfig) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *ApIpConfig) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *ApIpConfig) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetMtu

`func (o *ApIpConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *ApIpConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *ApIpConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *ApIpConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetmask

`func (o *ApIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ApIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ApIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ApIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *ApIpConfig) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *ApIpConfig) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *ApIpConfig) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *ApIpConfig) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetType

`func (o *ApIpConfig) GetType() IpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApIpConfig) GetTypeOk() (*IpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApIpConfig) SetType(v IpType)`

SetType sets Type field to given value.

### HasType

`func (o *ApIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType6

`func (o *ApIpConfig) GetType6() IpType6`

GetType6 returns the Type6 field if non-nil, zero value otherwise.

### GetType6Ok

`func (o *ApIpConfig) GetType6Ok() (*IpType6, bool)`

GetType6Ok returns a tuple with the Type6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType6

`func (o *ApIpConfig) SetType6(v IpType6)`

SetType6 sets Type6 field to given value.

### HasType6

`func (o *ApIpConfig) HasType6() bool`

HasType6 returns a boolean if a field has been set.

### GetVlanId

`func (o *ApIpConfig) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ApIpConfig) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ApIpConfig) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ApIpConfig) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


