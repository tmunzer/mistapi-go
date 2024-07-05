# GatewayOobIpConfigNode1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | if supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to false]
**UseMgmtVrfForHostOut** | Pointer to **bool** | whether to use &#x60;mgmt_junos&#x60; for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]

## Methods

### NewGatewayOobIpConfigNode1

`func NewGatewayOobIpConfigNode1() *GatewayOobIpConfigNode1`

NewGatewayOobIpConfigNode1 instantiates a new GatewayOobIpConfigNode1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayOobIpConfigNode1WithDefaults

`func NewGatewayOobIpConfigNode1WithDefaults() *GatewayOobIpConfigNode1`

NewGatewayOobIpConfigNode1WithDefaults instantiates a new GatewayOobIpConfigNode1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GatewayOobIpConfigNode1) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayOobIpConfigNode1) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayOobIpConfigNode1) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayOobIpConfigNode1) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *GatewayOobIpConfigNode1) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GatewayOobIpConfigNode1) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GatewayOobIpConfigNode1) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GatewayOobIpConfigNode1) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *GatewayOobIpConfigNode1) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayOobIpConfigNode1) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayOobIpConfigNode1) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayOobIpConfigNode1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *GatewayOobIpConfigNode1) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayOobIpConfigNode1) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayOobIpConfigNode1) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayOobIpConfigNode1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *GatewayOobIpConfigNode1) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *GatewayOobIpConfigNode1) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *GatewayOobIpConfigNode1) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *GatewayOobIpConfigNode1) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.

### GetUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfigNode1) GetUseMgmtVrfForHostOut() bool`

GetUseMgmtVrfForHostOut returns the UseMgmtVrfForHostOut field if non-nil, zero value otherwise.

### GetUseMgmtVrfForHostOutOk

`func (o *GatewayOobIpConfigNode1) GetUseMgmtVrfForHostOutOk() (*bool, bool)`

GetUseMgmtVrfForHostOutOk returns a tuple with the UseMgmtVrfForHostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfigNode1) SetUseMgmtVrfForHostOut(v bool)`

SetUseMgmtVrfForHostOut sets UseMgmtVrfForHostOut field to given value.

### HasUseMgmtVrfForHostOut

`func (o *GatewayOobIpConfigNode1) HasUseMgmtVrfForHostOut() bool`

HasUseMgmtVrfForHostOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


