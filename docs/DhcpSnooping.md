# DhcpSnooping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllNetworks** | Pointer to **bool** |  | [optional] 
**EnableArpSpoofCheck** | Pointer to **bool** | Enable for dynamic ARP inspection check | [optional] 
**EnableIpSourceGuard** | Pointer to **bool** | Enable for check for forging source IP address | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Networks** | Pointer to **[]string** | if &#x60;all_networks&#x60;&#x3D;&#x3D;&#x60;false&#x60;, list of network with DHCP snooping enabled | [optional] 

## Methods

### NewDhcpSnooping

`func NewDhcpSnooping() *DhcpSnooping`

NewDhcpSnooping instantiates a new DhcpSnooping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSnoopingWithDefaults

`func NewDhcpSnoopingWithDefaults() *DhcpSnooping`

NewDhcpSnoopingWithDefaults instantiates a new DhcpSnooping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllNetworks

`func (o *DhcpSnooping) GetAllNetworks() bool`

GetAllNetworks returns the AllNetworks field if non-nil, zero value otherwise.

### GetAllNetworksOk

`func (o *DhcpSnooping) GetAllNetworksOk() (*bool, bool)`

GetAllNetworksOk returns a tuple with the AllNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllNetworks

`func (o *DhcpSnooping) SetAllNetworks(v bool)`

SetAllNetworks sets AllNetworks field to given value.

### HasAllNetworks

`func (o *DhcpSnooping) HasAllNetworks() bool`

HasAllNetworks returns a boolean if a field has been set.

### GetEnableArpSpoofCheck

`func (o *DhcpSnooping) GetEnableArpSpoofCheck() bool`

GetEnableArpSpoofCheck returns the EnableArpSpoofCheck field if non-nil, zero value otherwise.

### GetEnableArpSpoofCheckOk

`func (o *DhcpSnooping) GetEnableArpSpoofCheckOk() (*bool, bool)`

GetEnableArpSpoofCheckOk returns a tuple with the EnableArpSpoofCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArpSpoofCheck

`func (o *DhcpSnooping) SetEnableArpSpoofCheck(v bool)`

SetEnableArpSpoofCheck sets EnableArpSpoofCheck field to given value.

### HasEnableArpSpoofCheck

`func (o *DhcpSnooping) HasEnableArpSpoofCheck() bool`

HasEnableArpSpoofCheck returns a boolean if a field has been set.

### GetEnableIpSourceGuard

`func (o *DhcpSnooping) GetEnableIpSourceGuard() bool`

GetEnableIpSourceGuard returns the EnableIpSourceGuard field if non-nil, zero value otherwise.

### GetEnableIpSourceGuardOk

`func (o *DhcpSnooping) GetEnableIpSourceGuardOk() (*bool, bool)`

GetEnableIpSourceGuardOk returns a tuple with the EnableIpSourceGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpSourceGuard

`func (o *DhcpSnooping) SetEnableIpSourceGuard(v bool)`

SetEnableIpSourceGuard sets EnableIpSourceGuard field to given value.

### HasEnableIpSourceGuard

`func (o *DhcpSnooping) HasEnableIpSourceGuard() bool`

HasEnableIpSourceGuard returns a boolean if a field has been set.

### GetEnabled

`func (o *DhcpSnooping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DhcpSnooping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DhcpSnooping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DhcpSnooping) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetworks

`func (o *DhcpSnooping) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DhcpSnooping) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DhcpSnooping) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *DhcpSnooping) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


