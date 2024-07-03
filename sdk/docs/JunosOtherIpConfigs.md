# JunosOtherIpConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpnAnycast** | Pointer to **bool** | for EVPN, if anycast is desired | [optional] [default to false]
**Gateway6** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Ip6** | Pointer to **string** | required if &#x60;type6&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] 
**Netmask** | Pointer to **string** | optional, &#x60;subnet&#x60; from &#x60;network&#x60; definition will be used if defined | [optional] 
**Netmask6** | Pointer to **string** | optional, &#x60;subnet&#x60; from &#x60;network&#x60; definition will be used if defined | [optional] 
**Type** | Pointer to [**IpType**](IpType.md) |  | [optional] [default to IPTYPE_DHCP]
**Type6** | Pointer to [**IpType6**](IpType6.md) |  | [optional] [default to IPTYPE6_DISABLED]

## Methods

### NewJunosOtherIpConfigs

`func NewJunosOtherIpConfigs() *JunosOtherIpConfigs`

NewJunosOtherIpConfigs instantiates a new JunosOtherIpConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOtherIpConfigsWithDefaults

`func NewJunosOtherIpConfigsWithDefaults() *JunosOtherIpConfigs`

NewJunosOtherIpConfigsWithDefaults instantiates a new JunosOtherIpConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpnAnycast

`func (o *JunosOtherIpConfigs) GetEvpnAnycast() bool`

GetEvpnAnycast returns the EvpnAnycast field if non-nil, zero value otherwise.

### GetEvpnAnycastOk

`func (o *JunosOtherIpConfigs) GetEvpnAnycastOk() (*bool, bool)`

GetEvpnAnycastOk returns a tuple with the EvpnAnycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpnAnycast

`func (o *JunosOtherIpConfigs) SetEvpnAnycast(v bool)`

SetEvpnAnycast sets EvpnAnycast field to given value.

### HasEvpnAnycast

`func (o *JunosOtherIpConfigs) HasEvpnAnycast() bool`

HasEvpnAnycast returns a boolean if a field has been set.

### GetGateway6

`func (o *JunosOtherIpConfigs) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *JunosOtherIpConfigs) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *JunosOtherIpConfigs) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *JunosOtherIpConfigs) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetIp

`func (o *JunosOtherIpConfigs) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOtherIpConfigs) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOtherIpConfigs) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOtherIpConfigs) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *JunosOtherIpConfigs) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *JunosOtherIpConfigs) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *JunosOtherIpConfigs) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *JunosOtherIpConfigs) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOtherIpConfigs) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOtherIpConfigs) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOtherIpConfigs) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOtherIpConfigs) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *JunosOtherIpConfigs) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *JunosOtherIpConfigs) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *JunosOtherIpConfigs) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *JunosOtherIpConfigs) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetType

`func (o *JunosOtherIpConfigs) GetType() IpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOtherIpConfigs) GetTypeOk() (*IpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOtherIpConfigs) SetType(v IpType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOtherIpConfigs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType6

`func (o *JunosOtherIpConfigs) GetType6() IpType6`

GetType6 returns the Type6 field if non-nil, zero value otherwise.

### GetType6Ok

`func (o *JunosOtherIpConfigs) GetType6Ok() (*IpType6, bool)`

GetType6Ok returns a tuple with the Type6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType6

`func (o *JunosOtherIpConfigs) SetType6(v IpType6)`

SetType6 sets Type6 field to given value.

### HasType6

`func (o *JunosOtherIpConfigs) HasType6() bool`

HasType6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


