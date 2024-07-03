# UtilsClearBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Neighbor** | **string** | neighbor ip-address or &#39;all&#39; | 
**Type** | [**UtilsClearBgpType**](UtilsClearBgpType.md) |  | [default to UTILSCLEARBGPTYPE_HARD]
**Vrf** | Pointer to **string** | vrf name | [optional] 

## Methods

### NewUtilsClearBgp

`func NewUtilsClearBgp(neighbor string, type_ UtilsClearBgpType, ) *UtilsClearBgp`

NewUtilsClearBgp instantiates a new UtilsClearBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsClearBgpWithDefaults

`func NewUtilsClearBgpWithDefaults() *UtilsClearBgp`

NewUtilsClearBgpWithDefaults instantiates a new UtilsClearBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeighbor

`func (o *UtilsClearBgp) GetNeighbor() string`

GetNeighbor returns the Neighbor field if non-nil, zero value otherwise.

### GetNeighborOk

`func (o *UtilsClearBgp) GetNeighborOk() (*string, bool)`

GetNeighborOk returns a tuple with the Neighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbor

`func (o *UtilsClearBgp) SetNeighbor(v string)`

SetNeighbor sets Neighbor field to given value.


### GetType

`func (o *UtilsClearBgp) GetType() UtilsClearBgpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UtilsClearBgp) GetTypeOk() (*UtilsClearBgpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UtilsClearBgp) SetType(v UtilsClearBgpType)`

SetType sets Type field to given value.


### GetVrf

`func (o *UtilsClearBgp) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsClearBgp) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsClearBgp) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsClearBgp) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


