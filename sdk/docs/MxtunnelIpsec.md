# MxtunnelIpsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | Pointer to **[]string** |  | [optional] 
**DnsSuffix** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExtraRoutes** | Pointer to [**[]MxtunnelIpsecExtraRoute**](MxtunnelIpsecExtraRoute.md) |  | [optional] 
**SplitTunnel** | Pointer to **bool** |  | [optional] 
**UseMxedge** | Pointer to **bool** |  | [optional] 

## Methods

### NewMxtunnelIpsec

`func NewMxtunnelIpsec() *MxtunnelIpsec`

NewMxtunnelIpsec instantiates a new MxtunnelIpsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxtunnelIpsecWithDefaults

`func NewMxtunnelIpsecWithDefaults() *MxtunnelIpsec`

NewMxtunnelIpsecWithDefaults instantiates a new MxtunnelIpsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *MxtunnelIpsec) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *MxtunnelIpsec) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *MxtunnelIpsec) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *MxtunnelIpsec) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *MxtunnelIpsec) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *MxtunnelIpsec) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetDnsSuffix

`func (o *MxtunnelIpsec) GetDnsSuffix() []string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *MxtunnelIpsec) GetDnsSuffixOk() (*[]string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *MxtunnelIpsec) SetDnsSuffix(v []string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *MxtunnelIpsec) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### GetEnabled

`func (o *MxtunnelIpsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxtunnelIpsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxtunnelIpsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxtunnelIpsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *MxtunnelIpsec) GetExtraRoutes() []MxtunnelIpsecExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *MxtunnelIpsec) GetExtraRoutesOk() (*[]MxtunnelIpsecExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *MxtunnelIpsec) SetExtraRoutes(v []MxtunnelIpsecExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *MxtunnelIpsec) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetSplitTunnel

`func (o *MxtunnelIpsec) GetSplitTunnel() bool`

GetSplitTunnel returns the SplitTunnel field if non-nil, zero value otherwise.

### GetSplitTunnelOk

`func (o *MxtunnelIpsec) GetSplitTunnelOk() (*bool, bool)`

GetSplitTunnelOk returns a tuple with the SplitTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitTunnel

`func (o *MxtunnelIpsec) SetSplitTunnel(v bool)`

SetSplitTunnel sets SplitTunnel field to given value.

### HasSplitTunnel

`func (o *MxtunnelIpsec) HasSplitTunnel() bool`

HasSplitTunnel returns a boolean if a field has been set.

### GetUseMxedge

`func (o *MxtunnelIpsec) GetUseMxedge() bool`

GetUseMxedge returns the UseMxedge field if non-nil, zero value otherwise.

### GetUseMxedgeOk

`func (o *MxtunnelIpsec) GetUseMxedgeOk() (*bool, bool)`

GetUseMxedgeOk returns a tuple with the UseMxedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMxedge

`func (o *MxtunnelIpsec) SetUseMxedge(v bool)`

SetUseMxedge sets UseMxedge field to given value.

### HasUseMxedge

`func (o *MxtunnelIpsec) HasUseMxedge() bool`

HasUseMxedge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


