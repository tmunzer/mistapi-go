# JunosOobIpConfigNode1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | optional, the network to be used for mgmt | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | if supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to false]
**UseMgmtVrfForHostOut** | Pointer to **bool** | whether to use &#x60;mgmt_junos&#x60; for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]
**VlanId** | Pointer to **int32** | optional, if different from parent | [optional] 

## Methods

### NewJunosOobIpConfigNode1

`func NewJunosOobIpConfigNode1() *JunosOobIpConfigNode1`

NewJunosOobIpConfigNode1 instantiates a new JunosOobIpConfigNode1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosOobIpConfigNode1WithDefaults

`func NewJunosOobIpConfigNode1WithDefaults() *JunosOobIpConfigNode1`

NewJunosOobIpConfigNode1WithDefaults instantiates a new JunosOobIpConfigNode1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *JunosOobIpConfigNode1) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosOobIpConfigNode1) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosOobIpConfigNode1) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosOobIpConfigNode1) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosOobIpConfigNode1) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosOobIpConfigNode1) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosOobIpConfigNode1) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosOobIpConfigNode1) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosOobIpConfigNode1) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosOobIpConfigNode1) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosOobIpConfigNode1) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosOobIpConfigNode1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *JunosOobIpConfigNode1) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosOobIpConfigNode1) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosOobIpConfigNode1) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosOobIpConfigNode1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *JunosOobIpConfigNode1) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *JunosOobIpConfigNode1) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *JunosOobIpConfigNode1) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *JunosOobIpConfigNode1) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.

### GetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigNode1) GetUseMgmtVrfForHostOut() bool`

GetUseMgmtVrfForHostOut returns the UseMgmtVrfForHostOut field if non-nil, zero value otherwise.

### GetUseMgmtVrfForHostOutOk

`func (o *JunosOobIpConfigNode1) GetUseMgmtVrfForHostOutOk() (*bool, bool)`

GetUseMgmtVrfForHostOutOk returns a tuple with the UseMgmtVrfForHostOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigNode1) SetUseMgmtVrfForHostOut(v bool)`

SetUseMgmtVrfForHostOut sets UseMgmtVrfForHostOut field to given value.

### HasUseMgmtVrfForHostOut

`func (o *JunosOobIpConfigNode1) HasUseMgmtVrfForHostOut() bool`

HasUseMgmtVrfForHostOut returns a boolean if a field has been set.

### GetVlanId

`func (o *JunosOobIpConfigNode1) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *JunosOobIpConfigNode1) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *JunosOobIpConfigNode1) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *JunosOobIpConfigNode1) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


