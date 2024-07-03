# MxedgeOobIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoconf6** | Pointer to **bool** |  | [optional] [default to true]
**Dhcp6** | Pointer to **bool** |  | [optional] [default to true]
**Dns** | Pointer to **[]string** | IPv4 ignored if &#x60;type&#x60;!&#x3D;&#x60;static&#x60; IPv6 ignored if &#x60;type6&#x60;!&#x3D;&#x60;static&#x60; | [optional] [default to ["8.8.8.8","8.8.4.4","2001:4860:4860::8888","2001:4860:4860::8844"]]
**Gateway** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] 
**Ip6** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] 
**Netmask6** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**IpType**](IpType.md) |  | [optional] [default to IPTYPE_DHCP]
**Type6** | Pointer to [**IpType**](IpType.md) |  | [optional] [default to IPTYPE_DHCP]

## Methods

### NewMxedgeOobIpConfig

`func NewMxedgeOobIpConfig() *MxedgeOobIpConfig`

NewMxedgeOobIpConfig instantiates a new MxedgeOobIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeOobIpConfigWithDefaults

`func NewMxedgeOobIpConfigWithDefaults() *MxedgeOobIpConfig`

NewMxedgeOobIpConfigWithDefaults instantiates a new MxedgeOobIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoconf6

`func (o *MxedgeOobIpConfig) GetAutoconf6() bool`

GetAutoconf6 returns the Autoconf6 field if non-nil, zero value otherwise.

### GetAutoconf6Ok

`func (o *MxedgeOobIpConfig) GetAutoconf6Ok() (*bool, bool)`

GetAutoconf6Ok returns a tuple with the Autoconf6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoconf6

`func (o *MxedgeOobIpConfig) SetAutoconf6(v bool)`

SetAutoconf6 sets Autoconf6 field to given value.

### HasAutoconf6

`func (o *MxedgeOobIpConfig) HasAutoconf6() bool`

HasAutoconf6 returns a boolean if a field has been set.

### GetDhcp6

`func (o *MxedgeOobIpConfig) GetDhcp6() bool`

GetDhcp6 returns the Dhcp6 field if non-nil, zero value otherwise.

### GetDhcp6Ok

`func (o *MxedgeOobIpConfig) GetDhcp6Ok() (*bool, bool)`

GetDhcp6Ok returns a tuple with the Dhcp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp6

`func (o *MxedgeOobIpConfig) SetDhcp6(v bool)`

SetDhcp6 sets Dhcp6 field to given value.

### HasDhcp6

`func (o *MxedgeOobIpConfig) HasDhcp6() bool`

HasDhcp6 returns a boolean if a field has been set.

### GetDns

`func (o *MxedgeOobIpConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *MxedgeOobIpConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *MxedgeOobIpConfig) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *MxedgeOobIpConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetGateway

`func (o *MxedgeOobIpConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MxedgeOobIpConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MxedgeOobIpConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MxedgeOobIpConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *MxedgeOobIpConfig) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *MxedgeOobIpConfig) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *MxedgeOobIpConfig) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *MxedgeOobIpConfig) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetIp

`func (o *MxedgeOobIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MxedgeOobIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MxedgeOobIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MxedgeOobIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *MxedgeOobIpConfig) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *MxedgeOobIpConfig) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *MxedgeOobIpConfig) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *MxedgeOobIpConfig) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetNetmask

`func (o *MxedgeOobIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *MxedgeOobIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *MxedgeOobIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *MxedgeOobIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *MxedgeOobIpConfig) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *MxedgeOobIpConfig) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *MxedgeOobIpConfig) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *MxedgeOobIpConfig) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetType

`func (o *MxedgeOobIpConfig) GetType() IpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MxedgeOobIpConfig) GetTypeOk() (*IpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MxedgeOobIpConfig) SetType(v IpType)`

SetType sets Type field to given value.

### HasType

`func (o *MxedgeOobIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType6

`func (o *MxedgeOobIpConfig) GetType6() IpType`

GetType6 returns the Type6 field if non-nil, zero value otherwise.

### GetType6Ok

`func (o *MxedgeOobIpConfig) GetType6Ok() (*IpType, bool)`

GetType6Ok returns a tuple with the Type6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType6

`func (o *MxedgeOobIpConfig) SetType6(v IpType)`

SetType6 sets Type6 field to given value.

### HasType6

`func (o *MxedgeOobIpConfig) HasType6() bool`

HasType6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


