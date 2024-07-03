# IpStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **[]string** |  | [optional] 
**DnsSuffix** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ip6** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **map[string]string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Netmask6** | Pointer to **string** |  | [optional] 

## Methods

### NewIpStat

`func NewIpStat() *IpStat`

NewIpStat instantiates a new IpStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpStatWithDefaults

`func NewIpStatWithDefaults() *IpStat`

NewIpStatWithDefaults instantiates a new IpStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *IpStat) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *IpStat) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *IpStat) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *IpStat) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDnsSuffix

`func (o *IpStat) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *IpStat) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *IpStat) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *IpStat) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetGateway

`func (o *IpStat) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpStat) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpStat) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IpStat) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *IpStat) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *IpStat) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *IpStat) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *IpStat) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetIp

`func (o *IpStat) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpStat) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpStat) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IpStat) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *IpStat) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *IpStat) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *IpStat) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *IpStat) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetIps

`func (o *IpStat) GetIps() map[string]string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IpStat) GetIpsOk() (*map[string]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IpStat) SetIps(v map[string]string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IpStat) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetNetmask

`func (o *IpStat) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IpStat) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IpStat) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IpStat) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *IpStat) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *IpStat) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *IpStat) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *IpStat) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


