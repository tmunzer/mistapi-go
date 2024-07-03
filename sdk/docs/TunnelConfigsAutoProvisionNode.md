# TunnelConfigsAutoProvisionNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumHosts** | Pointer to **string** |  | [optional] 
**WanNames** | Pointer to **[]string** | optional, only needed if &#x60;vars_only&#x60;&#x3D;&#x3D;&#x60;false&#x60; | [optional] 

## Methods

### NewTunnelConfigsAutoProvisionNode

`func NewTunnelConfigsAutoProvisionNode() *TunnelConfigsAutoProvisionNode`

NewTunnelConfigsAutoProvisionNode instantiates a new TunnelConfigsAutoProvisionNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigsAutoProvisionNodeWithDefaults

`func NewTunnelConfigsAutoProvisionNodeWithDefaults() *TunnelConfigsAutoProvisionNode`

NewTunnelConfigsAutoProvisionNodeWithDefaults instantiates a new TunnelConfigsAutoProvisionNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumHosts

`func (o *TunnelConfigsAutoProvisionNode) GetNumHosts() string`

GetNumHosts returns the NumHosts field if non-nil, zero value otherwise.

### GetNumHostsOk

`func (o *TunnelConfigsAutoProvisionNode) GetNumHostsOk() (*string, bool)`

GetNumHostsOk returns a tuple with the NumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHosts

`func (o *TunnelConfigsAutoProvisionNode) SetNumHosts(v string)`

SetNumHosts sets NumHosts field to given value.

### HasNumHosts

`func (o *TunnelConfigsAutoProvisionNode) HasNumHosts() bool`

HasNumHosts returns a boolean if a field has been set.

### GetWanNames

`func (o *TunnelConfigsAutoProvisionNode) GetWanNames() []string`

GetWanNames returns the WanNames field if non-nil, zero value otherwise.

### GetWanNamesOk

`func (o *TunnelConfigsAutoProvisionNode) GetWanNamesOk() (*[]string, bool)`

GetWanNamesOk returns a tuple with the WanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanNames

`func (o *TunnelConfigsAutoProvisionNode) SetWanNames(v []string)`

SetWanNames sets WanNames field to given value.

### HasWanNames

`func (o *TunnelConfigsAutoProvisionNode) HasWanNames() bool`

HasWanNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


