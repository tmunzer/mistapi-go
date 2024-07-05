# JunosOtherIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpnAnycast** | Pointer to **bool** | for EVPN, if anycast is desired | [optional] [default to false]
**Ip** | Pointer to **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Ip6** | Pointer to **string** | required if &#x60;type6&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Netmask** | Pointer to **string** | optional, &#x60;subnet&#x60; from &#x60;network&#x60; definition will be used if defined | [optional] 
**Netmask6** | Pointer to **string** | optional, &#x60;subnet&#x60; from &#x60;network&#x60; definition will be used if defined | [optional] 
**Type** | Pointer to [**IpType**](IpType.md) |  | [optional] [default to IPTYPE_DHCP]
**Type6** | Pointer to [**IpType6**](IpType6.md) |  | [optional] [default to IPTYPE6_DISABLED]

## Methods

### NewJunosOtherIpConfig

`func NewJunosOtherIpConfig() *JunosOtherIpConfig`

NewJunosOtherIpConfig instantiates a new JunosOtherIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOtherIpConfigWithDefaults

`func NewJunosOtherIpConfigWithDefaults() *JunosOtherIpConfig`

NewJunosOtherIpConfigWithDefaults instantiates a new JunosOtherIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpnAnycast

`func (o *JunosOtherIpConfig) GetEvpnAnycast() bool`

GetEvpnAnycast returns the EvpnAnycast field if non-nil, zero value otherwise.

### GetEvpnAnycastOk

`func (o *JunosOtherIpConfig) GetEvpnAnycastOk() (*bool, bool)`

GetEvpnAnycastOk returns a tuple with the EvpnAnycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnAnycast

`func (o *JunosOtherIpConfig) SetEvpnAnycast(v bool)`

SetEvpnAnycast sets EvpnAnycast field to given value.

### HasEvpnAnycast

`func (o *JunosOtherIpConfig) HasEvpnAnycast() bool`

HasEvpnAnycast returns a boolean if a field has been set.

### GetIp

`func (o *JunosOtherIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOtherIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOtherIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOtherIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *JunosOtherIpConfig) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *JunosOtherIpConfig) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *JunosOtherIpConfig) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *JunosOtherIpConfig) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOtherIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOtherIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOtherIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOtherIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *JunosOtherIpConfig) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *JunosOtherIpConfig) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *JunosOtherIpConfig) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *JunosOtherIpConfig) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetType

`func (o *JunosOtherIpConfig) GetType() IpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOtherIpConfig) GetTypeOk() (*IpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOtherIpConfig) SetType(v IpType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOtherIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType6

`func (o *JunosOtherIpConfig) GetType6() IpType6`

GetType6 returns the Type6 field if non-nil, zero value otherwise.

### GetType6Ok

`func (o *JunosOtherIpConfig) GetType6Ok() (*IpType6, bool)`

GetType6Ok returns a tuple with the Type6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType6

`func (o *JunosOtherIpConfig) SetType6(v IpType6)`

SetType6 sets Type6 field to given value.

### HasType6

`func (o *JunosOtherIpConfig) HasType6() bool`

HasType6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


