# JunosOobIpConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** |  | [optional] 
**DnsSuffix** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Node1** | Pointer to [**JunosOobIpConfigNode1**](JunosOobIpConfigNode1.md) |  | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | if supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to true]
**UseMgmtVrfForHostOut** | Pointer to **bool** | whether to use &#x60;mgmt_junos&#x60; for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewJunosOobIpConfigs

`func NewJunosOobIpConfigs() *JunosOobIpConfigs`

NewJunosOobIpConfigs instantiates a new JunosOobIpConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOobIpConfigsWithDefaults

`func NewJunosOobIpConfigsWithDefaults() *JunosOobIpConfigs`

NewJunosOobIpConfigsWithDefaults instantiates a new JunosOobIpConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *JunosOobIpConfigs) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *JunosOobIpConfigs) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *JunosOobIpConfigs) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *JunosOobIpConfigs) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosOobIpConfigs) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosOobIpConfigs) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosOobIpConfigs) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosOobIpConfigs) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *JunosOobIpConfigs) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosOobIpConfigs) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosOobIpConfigs) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosOobIpConfigs) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *JunosOobIpConfigs) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOobIpConfigs) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOobIpConfigs) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOobIpConfigs) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOobIpConfigs) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOobIpConfigs) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOobIpConfigs) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOobIpConfigs) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosOobIpConfigs) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosOobIpConfigs) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosOobIpConfigs) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosOobIpConfigs) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNode1

`func (o *JunosOobIpConfigs) GetNode1() JunosOobIpConfigNode1`

GetNode1 returns the Node1 field if non-nil, zero value otherwise.

### GetNode1Ok

`func (o *JunosOobIpConfigs) GetNode1Ok() (*JunosOobIpConfigNode1, bool)`

GetNode1Ok returns a tuple with the Node1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode1

`func (o *JunosOobIpConfigs) SetNode1(v JunosOobIpConfigNode1)`

SetNode1 sets Node1 field to given value.

### HasNode1

`func (o *JunosOobIpConfigs) HasNode1() bool`

HasNode1 returns a boolean if a field has been set.

### GetType

`func (o *JunosOobIpConfigs) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOobIpConfigs) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOobIpConfigs) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOobIpConfigs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *JunosOobIpConfigs) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *JunosOobIpConfigs) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *JunosOobIpConfigs) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *JunosOobIpConfigs) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.

### GetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigs) GetUseMgmtVrfForHostOut() bool`

GetUseMgmtVrfForHostOut returns the UseMgmtVrfForHostOut field if non-nil, zero value otherwise.

### GetUseMgmtVrfForHostOutOk

`func (o *JunosOobIpConfigs) GetUseMgmtVrfForHostOutOk() (*bool, bool)`

GetUseMgmtVrfForHostOutOk returns a tuple with the UseMgmtVrfForHostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigs) SetUseMgmtVrfForHostOut(v bool)`

SetUseMgmtVrfForHostOut sets UseMgmtVrfForHostOut field to given value.

### HasUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigs) HasUseMgmtVrfForHostOut() bool`

HasUseMgmtVrfForHostOut returns a boolean if a field has been set.

### GetVlanId

`func (o *JunosOobIpConfigs) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *JunosOobIpConfigs) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *JunosOobIpConfigs) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *JunosOobIpConfigs) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


