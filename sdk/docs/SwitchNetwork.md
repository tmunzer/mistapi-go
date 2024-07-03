# SwitchNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isolation** | Pointer to **bool** | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required) NOTE: this features requires uplink device to also a be Juniper device and &#x60;inter_switch_link&#x60; to be set | [optional] [default to false]
**IsolationVlanId** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **string** | optional for pure switching, required when L3 / routing features are used | [optional] 
**VlanId** | **int32** |  | 

## Methods

### NewSwitchNetwork

`func NewSwitchNetwork(vlanId int32, ) *SwitchNetwork`

NewSwitchNetwork instantiates a new SwitchNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchNetworkWithDefaults

`func NewSwitchNetworkWithDefaults() *SwitchNetwork`

NewSwitchNetworkWithDefaults instantiates a new SwitchNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolation

`func (o *SwitchNetwork) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *SwitchNetwork) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *SwitchNetwork) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *SwitchNetwork) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetIsolationVlanId

`func (o *SwitchNetwork) GetIsolationVlanId() string`

GetIsolationVlanId returns the IsolationVlanId field if non-nil, zero value otherwise.

### GetIsolationVlanIdOk

`func (o *SwitchNetwork) GetIsolationVlanIdOk() (*string, bool)`

GetIsolationVlanIdOk returns a tuple with the IsolationVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationVlanId

`func (o *SwitchNetwork) SetIsolationVlanId(v string)`

SetIsolationVlanId sets IsolationVlanId field to given value.

### HasIsolationVlanId

`func (o *SwitchNetwork) HasIsolationVlanId() bool`

HasIsolationVlanId returns a boolean if a field has been set.

### GetSubnet

`func (o *SwitchNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *SwitchNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *SwitchNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *SwitchNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVlanId

`func (o *SwitchNetwork) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SwitchNetwork) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SwitchNetwork) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


