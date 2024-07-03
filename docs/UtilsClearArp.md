# UtilsClearArp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IP address for which to clear an ARP entry. port_id must be specified. Both vlan and ip cannot be specified. | [optional] 
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**PortId** | Pointer to **string** | The device interface on which to clear the ARP cache. | [optional] 
**Vlan** | Pointer to **int32** | The VLAN on which to clear the ARP cache. port_id must be specified. Both vlan and ip cannot be specified. | [optional] 
**Vrf** | Pointer to **string** | The vrf for which to clear an ARP entry. applicable for switch. | [optional] 

## Methods

### NewUtilsClearArp

`func NewUtilsClearArp() *UtilsClearArp`

NewUtilsClearArp instantiates a new UtilsClearArp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsClearArpWithDefaults

`func NewUtilsClearArpWithDefaults() *UtilsClearArp`

NewUtilsClearArpWithDefaults instantiates a new UtilsClearArp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *UtilsClearArp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UtilsClearArp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UtilsClearArp) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UtilsClearArp) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNode

`func (o *UtilsClearArp) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsClearArp) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsClearArp) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsClearArp) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPortId

`func (o *UtilsClearArp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *UtilsClearArp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *UtilsClearArp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *UtilsClearArp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVlan

`func (o *UtilsClearArp) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UtilsClearArp) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UtilsClearArp) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *UtilsClearArp) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVrf

`func (o *UtilsClearArp) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsClearArp) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsClearArp) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsClearArp) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


