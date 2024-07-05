# GatewayOobIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Node1** | Pointer to [**GatewayOobIpConfigNode1**](GatewayOobIpConfigNode1.md) |  | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | f supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to false]
**UseMgmtVrfForHostOut** | Pointer to **bool** | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired, | [optional] [default to false]

## Methods

### NewGatewayOobIpConfig

`func NewGatewayOobIpConfig() *GatewayOobIpConfig`

NewGatewayOobIpConfig instantiates a new GatewayOobIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayOobIpConfigWithDefaults

`func NewGatewayOobIpConfigWithDefaults() *GatewayOobIpConfig`

NewGatewayOobIpConfigWithDefaults instantiates a new GatewayOobIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GatewayOobIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayOobIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayOobIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayOobIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *GatewayOobIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GatewayOobIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GatewayOobIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GatewayOobIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *GatewayOobIpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayOobIpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayOobIpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayOobIpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNode1

`func (o *GatewayOobIpConfig) GetNode1() GatewayOobIpConfigNode1`

GetNode1 returns the Node1 field if non-nil, zero value otherwise.

### GetNode1Ok

`func (o *GatewayOobIpConfig) GetNode1Ok() (*GatewayOobIpConfigNode1, bool)`

GetNode1Ok returns a tuple with the Node1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode1

`func (o *GatewayOobIpConfig) SetNode1(v GatewayOobIpConfigNode1)`

SetNode1 sets Node1 field to given value.

### HasNode1

`func (o *GatewayOobIpConfig) HasNode1() bool`

HasNode1 returns a boolean if a field has been set.

### GetType

`func (o *GatewayOobIpConfig) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayOobIpConfig) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayOobIpConfig) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayOobIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *GatewayOobIpConfig) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *GatewayOobIpConfig) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *GatewayOobIpConfig) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *GatewayOobIpConfig) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.

### GetUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfig) GetUseMgmtVrfForHostOut() bool`

GetUseMgmtVrfForHostOut returns the UseMgmtVrfForHostOut field if non-nil, zero value otherwise.

### GetUseMgmtVrfForHostOutOk

`func (o *GatewayOobIpConfig) GetUseMgmtVrfForHostOutOk() (*bool, bool)`

GetUseMgmtVrfForHostOutOk returns a tuple with the UseMgmtVrfForHostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfig) SetUseMgmtVrfForHostOut(v bool)`

SetUseMgmtVrfForHostOut sets UseMgmtVrfForHostOut field to given value.

### HasUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfig) HasUseMgmtVrfForHostOut() bool`

HasUseMgmtVrfForHostOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


