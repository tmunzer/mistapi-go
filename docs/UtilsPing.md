# UtilsPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [default to 10]
**EgressInterface** | Pointer to **string** | Interface through which packet needs to egress | [optional] 
**Host** | **string** |  | 
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] [default to 56]

## Methods

### NewUtilsPing

`func NewUtilsPing(host string, ) *UtilsPing`

NewUtilsPing instantiates a new UtilsPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsPingWithDefaults

`func NewUtilsPingWithDefaults() *UtilsPing`

NewUtilsPingWithDefaults instantiates a new UtilsPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UtilsPing) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UtilsPing) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UtilsPing) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UtilsPing) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEgressInterface

`func (o *UtilsPing) GetEgressInterface() string`

GetEgressInterface returns the EgressInterface field if non-nil, zero value otherwise.

### GetEgressInterfaceOk

`func (o *UtilsPing) GetEgressInterfaceOk() (*string, bool)`

GetEgressInterfaceOk returns a tuple with the EgressInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressInterface

`func (o *UtilsPing) SetEgressInterface(v string)`

SetEgressInterface sets EgressInterface field to given value.

### HasEgressInterface

`func (o *UtilsPing) HasEgressInterface() bool`

HasEgressInterface returns a boolean if a field has been set.

### GetHost

`func (o *UtilsPing) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UtilsPing) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UtilsPing) SetHost(v string)`

SetHost sets Host field to given value.


### GetNode

`func (o *UtilsPing) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsPing) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsPing) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsPing) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetSize

`func (o *UtilsPing) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UtilsPing) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UtilsPing) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UtilsPing) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


