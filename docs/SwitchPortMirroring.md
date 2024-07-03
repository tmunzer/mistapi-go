# SwitchPortMirroring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputNetworksIngress** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**InputPortIdsEgress** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**InputPortIdsIngress** | Pointer to **[]string** | at least one of the &#x60;ingress_port_ids&#x60;, &#x60;egress_port_ids&#x60; or &#x60;ingress_networks &#x60; should be specified | [optional] 
**OutputNetwork** | Pointer to **string** |  | [optional] 
**OutputPortId** | Pointer to **string** | exaclty on of the &#x60;output_port_id&#x60; or &#x60;output_network&#x60; should be provided | [optional] 

## Methods

### NewSwitchPortMirroring

`func NewSwitchPortMirroring() *SwitchPortMirroring`

NewSwitchPortMirroring instantiates a new SwitchPortMirroring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortMirroringWithDefaults

`func NewSwitchPortMirroringWithDefaults() *SwitchPortMirroring`

NewSwitchPortMirroringWithDefaults instantiates a new SwitchPortMirroring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputNetworksIngress

`func (o *SwitchPortMirroring) GetInputNetworksIngress() []string`

GetInputNetworksIngress returns the InputNetworksIngress field if non-nil, zero value otherwise.

### GetInputNetworksIngressOk

`func (o *SwitchPortMirroring) GetInputNetworksIngressOk() (*[]string, bool)`

GetInputNetworksIngressOk returns a tuple with the InputNetworksIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputNetworksIngress

`func (o *SwitchPortMirroring) SetInputNetworksIngress(v []string)`

SetInputNetworksIngress sets InputNetworksIngress field to given value.

### HasInputNetworksIngress

`func (o *SwitchPortMirroring) HasInputNetworksIngress() bool`

HasInputNetworksIngress returns a boolean if a field has been set.

### GetInputPortIdsEgress

`func (o *SwitchPortMirroring) GetInputPortIdsEgress() []string`

GetInputPortIdsEgress returns the InputPortIdsEgress field if non-nil, zero value otherwise.

### GetInputPortIdsEgressOk

`func (o *SwitchPortMirroring) GetInputPortIdsEgressOk() (*[]string, bool)`

GetInputPortIdsEgressOk returns a tuple with the InputPortIdsEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortIdsEgress

`func (o *SwitchPortMirroring) SetInputPortIdsEgress(v []string)`

SetInputPortIdsEgress sets InputPortIdsEgress field to given value.

### HasInputPortIdsEgress

`func (o *SwitchPortMirroring) HasInputPortIdsEgress() bool`

HasInputPortIdsEgress returns a boolean if a field has been set.

### GetInputPortIdsIngress

`func (o *SwitchPortMirroring) GetInputPortIdsIngress() []string`

GetInputPortIdsIngress returns the InputPortIdsIngress field if non-nil, zero value otherwise.

### GetInputPortIdsIngressOk

`func (o *SwitchPortMirroring) GetInputPortIdsIngressOk() (*[]string, bool)`

GetInputPortIdsIngressOk returns a tuple with the InputPortIdsIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortIdsIngress

`func (o *SwitchPortMirroring) SetInputPortIdsIngress(v []string)`

SetInputPortIdsIngress sets InputPortIdsIngress field to given value.

### HasInputPortIdsIngress

`func (o *SwitchPortMirroring) HasInputPortIdsIngress() bool`

HasInputPortIdsIngress returns a boolean if a field has been set.

### GetOutputNetwork

`func (o *SwitchPortMirroring) GetOutputNetwork() string`

GetOutputNetwork returns the OutputNetwork field if non-nil, zero value otherwise.

### GetOutputNetworkOk

`func (o *SwitchPortMirroring) GetOutputNetworkOk() (*string, bool)`

GetOutputNetworkOk returns a tuple with the OutputNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputNetwork

`func (o *SwitchPortMirroring) SetOutputNetwork(v string)`

SetOutputNetwork sets OutputNetwork field to given value.

### HasOutputNetwork

`func (o *SwitchPortMirroring) HasOutputNetwork() bool`

HasOutputNetwork returns a boolean if a field has been set.

### GetOutputPortId

`func (o *SwitchPortMirroring) GetOutputPortId() string`

GetOutputPortId returns the OutputPortId field if non-nil, zero value otherwise.

### GetOutputPortIdOk

`func (o *SwitchPortMirroring) GetOutputPortIdOk() (*string, bool)`

GetOutputPortIdOk returns a tuple with the OutputPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortId

`func (o *SwitchPortMirroring) SetOutputPortId(v string)`

SetOutputPortId sets OutputPortId field to given value.

### HasOutputPortId

`func (o *SwitchPortMirroring) HasOutputPortId() bool`

HasOutputPortId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


