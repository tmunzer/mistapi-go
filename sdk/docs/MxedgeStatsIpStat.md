# MxedgeStatsIpStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **map[string]string** | Property key is the interface name. IPs for each net interface | [optional] 
**Macs** | Pointer to **map[string]string** | Property key is the interface name. MAC for each net interface | [optional] 

## Methods

### NewMxedgeStatsIpStat

`func NewMxedgeStatsIpStat() *MxedgeStatsIpStat`

NewMxedgeStatsIpStat instantiates a new MxedgeStatsIpStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsIpStatWithDefaults

`func NewMxedgeStatsIpStatWithDefaults() *MxedgeStatsIpStat`

NewMxedgeStatsIpStatWithDefaults instantiates a new MxedgeStatsIpStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *MxedgeStatsIpStat) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MxedgeStatsIpStat) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MxedgeStatsIpStat) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MxedgeStatsIpStat) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIps

`func (o *MxedgeStatsIpStat) GetIps() map[string]string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *MxedgeStatsIpStat) GetIpsOk() (*map[string]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *MxedgeStatsIpStat) SetIps(v map[string]string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *MxedgeStatsIpStat) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetMacs

`func (o *MxedgeStatsIpStat) GetMacs() map[string]string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *MxedgeStatsIpStat) GetMacsOk() (*map[string]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *MxedgeStatsIpStat) SetMacs(v map[string]string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *MxedgeStatsIpStat) HasMacs() bool`

HasMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


