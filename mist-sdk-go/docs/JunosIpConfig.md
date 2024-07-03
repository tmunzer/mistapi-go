# JunosIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** |  | [optional] 
**DnsSuffix** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] 
**Network** | Pointer to **string** | the network where this mgmt IP reside, this will be used as default network for outbound-ssh, dns, ntp, dns, tacplus, radius, syslog, snmp | [optional] 
**Type** | Pointer to [**IpConfigType**](IpConfigType.md) |  | [optional] [default to IPCONFIGTYPE_DYNAMIC]
**UseMgmtVrf** | Pointer to **bool** | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP) | [optional] [default to false]

## Methods

### NewJunosIpConfig

`func NewJunosIpConfig() *JunosIpConfig`

NewJunosIpConfig instantiates a new JunosIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJunosIpConfigWithDefaults

`func NewJunosIpConfigWithDefaults() *JunosIpConfig`

NewJunosIpConfigWithDefaults instantiates a new JunosIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *JunosIpConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *JunosIpConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *JunosIpConfig) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *JunosIpConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *JunosIpConfig) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *JunosIpConfig) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *JunosIpConfig) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *JunosIpConfig) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *JunosIpConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JunosIpConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JunosIpConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JunosIpConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIp

`func (o *JunosIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *JunosIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *JunosIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *JunosIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *JunosIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *JunosIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *JunosIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *JunosIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *JunosIpConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *JunosIpConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *JunosIpConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *JunosIpConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *JunosIpConfig) GetType() IpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JunosIpConfig) GetTypeOk() (*IpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JunosIpConfig) SetType(v IpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *JunosIpConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseMgmtVrf

`func (o *JunosIpConfig) GetUseMgmtVrf() bool`

GetUseMgmtVrf returns the UseMgmtVrf field if non-nil, zero value otherwise.

### GetUseMgmtVrfOk

`func (o *JunosIpConfig) GetUseMgmtVrfOk() (*bool, bool)`

GetUseMgmtVrfOk returns a tuple with the UseMgmtVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMgmtVrf

`func (o *JunosIpConfig) SetUseMgmtVrf(v bool)`

SetUseMgmtVrf sets UseMgmtVrf field to given value.

### HasUseMgmtVrf

`func (o *JunosIpConfig) HasUseMgmtVrf() bool`

HasUseMgmtVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


