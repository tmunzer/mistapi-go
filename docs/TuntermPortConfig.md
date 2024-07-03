# TuntermPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownstreamPorts** | Pointer to **[]string** | list of ports to be used for downstream (to AP) purpose | [optional] 
**SeparateUpstreamDownstream** | Pointer to **bool** | weather to separate upstream / downstream ports. default is false where all ports will be used. | [optional] [default to false]
**UpstreamPortVlanId** | Pointer to **int32** | native VLAN id for upstream ports | [optional] [default to 1]
**UpstreamPorts** | Pointer to **[]string** | list of ports to be used for upstrea purpose (to LAN) | [optional] 

## Methods

### NewTuntermPortConfig

`func NewTuntermPortConfig() *TuntermPortConfig`

NewTuntermPortConfig instantiates a new TuntermPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTuntermPortConfigWithDefaults

`func NewTuntermPortConfigWithDefaults() *TuntermPortConfig`

NewTuntermPortConfigWithDefaults instantiates a new TuntermPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownstreamPorts

`func (o *TuntermPortConfig) GetDownstreamPorts() []string`

GetDownstreamPorts returns the DownstreamPorts field if non-nil, zero value otherwise.

### GetDownstreamPortsOk

`func (o *TuntermPortConfig) GetDownstreamPortsOk() (*[]string, bool)`

GetDownstreamPortsOk returns a tuple with the DownstreamPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamPorts

`func (o *TuntermPortConfig) SetDownstreamPorts(v []string)`

SetDownstreamPorts sets DownstreamPorts field to given value.

### HasDownstreamPorts

`func (o *TuntermPortConfig) HasDownstreamPorts() bool`

HasDownstreamPorts returns a boolean if a field has been set.

### GetSeparateUpstreamDownstream

`func (o *TuntermPortConfig) GetSeparateUpstreamDownstream() bool`

GetSeparateUpstreamDownstream returns the SeparateUpstreamDownstream field if non-nil, zero value otherwise.

### GetSeparateUpstreamDownstreamOk

`func (o *TuntermPortConfig) GetSeparateUpstreamDownstreamOk() (*bool, bool)`

GetSeparateUpstreamDownstreamOk returns a tuple with the SeparateUpstreamDownstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateUpstreamDownstream

`func (o *TuntermPortConfig) SetSeparateUpstreamDownstream(v bool)`

SetSeparateUpstreamDownstream sets SeparateUpstreamDownstream field to given value.

### HasSeparateUpstreamDownstream

`func (o *TuntermPortConfig) HasSeparateUpstreamDownstream() bool`

HasSeparateUpstreamDownstream returns a boolean if a field has been set.

### GetUpstreamPortVlanId

`func (o *TuntermPortConfig) GetUpstreamPortVlanId() int32`

GetUpstreamPortVlanId returns the UpstreamPortVlanId field if non-nil, zero value otherwise.

### GetUpstreamPortVlanIdOk

`func (o *TuntermPortConfig) GetUpstreamPortVlanIdOk() (*int32, bool)`

GetUpstreamPortVlanIdOk returns a tuple with the UpstreamPortVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPortVlanId

`func (o *TuntermPortConfig) SetUpstreamPortVlanId(v int32)`

SetUpstreamPortVlanId sets UpstreamPortVlanId field to given value.

### HasUpstreamPortVlanId

`func (o *TuntermPortConfig) HasUpstreamPortVlanId() bool`

HasUpstreamPortVlanId returns a boolean if a field has been set.

### GetUpstreamPorts

`func (o *TuntermPortConfig) GetUpstreamPorts() []string`

GetUpstreamPorts returns the UpstreamPorts field if non-nil, zero value otherwise.

### GetUpstreamPortsOk

`func (o *TuntermPortConfig) GetUpstreamPortsOk() (*[]string, bool)`

GetUpstreamPortsOk returns a tuple with the UpstreamPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPorts

`func (o *TuntermPortConfig) SetUpstreamPorts(v []string)`

SetUpstreamPorts sets UpstreamPorts field to given value.

### HasUpstreamPorts

`func (o *TuntermPortConfig) HasUpstreamPorts() bool`

HasUpstreamPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


