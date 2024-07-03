# OspfAreas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeLoopback** | Pointer to **bool** |  | [optional] [default to false]
**Networks** | Pointer to [**map[string]OspfAreasNetwork**](OspfAreasNetwork.md) | Networks to participate in an OSPF area.  Property key is the network name | [optional] 
**Type** | Pointer to [**OspfAreasType**](OspfAreasType.md) |  | [optional] [default to OSPFAREASTYPE_DEFAULT]

## Methods

### NewOspfAreas

`func NewOspfAreas() *OspfAreas`

NewOspfAreas instantiates a new OspfAreas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfAreasWithDefaults

`func NewOspfAreasWithDefaults() *OspfAreas`

NewOspfAreasWithDefaults instantiates a new OspfAreas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeLoopback

`func (o *OspfAreas) GetIncludeLoopback() bool`

GetIncludeLoopback returns the IncludeLoopback field if non-nil, zero value otherwise.

### GetIncludeLoopbackOk

`func (o *OspfAreas) GetIncludeLoopbackOk() (*bool, bool)`

GetIncludeLoopbackOk returns a tuple with the IncludeLoopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLoopback

`func (o *OspfAreas) SetIncludeLoopback(v bool)`

SetIncludeLoopback sets IncludeLoopback field to given value.

### HasIncludeLoopback

`func (o *OspfAreas) HasIncludeLoopback() bool`

HasIncludeLoopback returns a boolean if a field has been set.

### GetNetworks

`func (o *OspfAreas) GetNetworks() map[string]OspfAreasNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *OspfAreas) GetNetworksOk() (*map[string]OspfAreasNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *OspfAreas) SetNetworks(v map[string]OspfAreasNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *OspfAreas) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetType

`func (o *OspfAreas) GetType() OspfAreasType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OspfAreas) GetTypeOk() (*OspfAreasType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OspfAreas) SetType(v OspfAreasType)`

SetType sets Type field to given value.

### HasType

`func (o *OspfAreas) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


