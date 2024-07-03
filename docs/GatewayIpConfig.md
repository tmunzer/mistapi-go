# GatewayIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** | except for out-of_band interface (vme/em0/fxp0) | [optional] 
**DnsSuffix** | Pointer to **[]string** | except for out-of_band interface (vme/em0/fxp0) | [optional] 
**Gateway** | Pointer to **string** | except for out-of_band interface (vme/em0/fxp0) | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**PoserPassword** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;pppoe&#x60; | [optional] 
**PppoeAuth** | Pointer to [**GatewayWanPpoeAuth**](GatewayWanPpoeAuth.md) |  | [optional] [default to GATEWAYWANPPOEAUTH_NONE]
**PppoeUsername** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;pppoe&#x60; | [optional] 
**Type** | Pointer to [**GatewayWanType**](GatewayWanType.md) |  | [optional] [default to GATEWAYWANTYPE_DHCP]

## Methods

### NewGatewayIpConfig

`func NewGatewayIpConfig() *GatewayIpConfig`

NewGatewayIpConfig instantiates a new GatewayIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayIpConfigWithDefaults

`func NewGatewayIpConfigWithDefaults() *GatewayIpConfig`

NewGatewayIpConfigWithDefaults instantiates a new GatewayIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *GatewayIpConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *GatewayIpConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *GatewayIpConfig) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *GatewayIpConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *GatewayIpConfig) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *GatewayIpConfig) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *GatewayIpConfig) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *GatewayIpConfig) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *GatewayIpConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GatewayIpConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GatewayIpConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GatewayIpConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *GatewayIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *GatewayIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GatewayIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GatewayIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GatewayIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *GatewayIpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayIpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayIpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayIpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPoserPassword

`func (o *GatewayIpConfig) GetPoserPassword() string`

GetPoserPassword returns the PoserPassword field if non-nil, zero value otherwise.

### GetPoserPasswordOk

`func (o *GatewayIpConfig) GetPoserPasswordOk() (*string, bool)`

GetPoserPasswordOk returns a tuple with the PoserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoserPassword

`func (o *GatewayIpConfig) SetPoserPassword(v string)`

SetPoserPassword sets PoserPassword field to given value.

### HasPoserPassword

`func (o *GatewayIpConfig) HasPoserPassword() bool`

HasPoserPassword returns a boolean if a field has been set.

### GetPppoeAuth

`func (o *GatewayIpConfig) GetPppoeAuth() GatewayWanPpoeAuth`

GetPppoeAuth returns the PppoeAuth field if non-nil, zero value otherwise.

### GetPppoeAuthOk

`func (o *GatewayIpConfig) GetPppoeAuthOk() (*GatewayWanPpoeAuth, bool)`

GetPppoeAuthOk returns a tuple with the PppoeAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoeAuth

`func (o *GatewayIpConfig) SetPppoeAuth(v GatewayWanPpoeAuth)`

SetPppoeAuth sets PppoeAuth field to given value.

### HasPppoeAuth

`func (o *GatewayIpConfig) HasPppoeAuth() bool`

HasPppoeAuth returns a boolean if a field has been set.

### GetPppoeUsername

`func (o *GatewayIpConfig) GetPppoeUsername() string`

GetPppoeUsername returns the PppoeUsername field if non-nil, zero value otherwise.

### GetPppoeUsernameOk

`func (o *GatewayIpConfig) GetPppoeUsernameOk() (*string, bool)`

GetPppoeUsernameOk returns a tuple with the PppoeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoeUsername

`func (o *GatewayIpConfig) SetPppoeUsername(v string)`

SetPppoeUsername sets PppoeUsername field to given value.

### HasPppoeUsername

`func (o *GatewayIpConfig) HasPppoeUsername() bool`

HasPppoeUsername returns a boolean if a field has been set.

### GetType

`func (o *GatewayIpConfig) GetType() GatewayWanType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayIpConfig) GetTypeOk() (*GatewayWanType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayIpConfig) SetType(v GatewayWanType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


