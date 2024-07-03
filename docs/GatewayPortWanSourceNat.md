# GatewayPortWanSourceNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | or to disable the source-nat | [optional] [default to false]
**NatPool** | Pointer to **string** | if alternative nat_pool is desired | [optional] 

## Methods

### NewGatewayPortWanSourceNat

`func NewGatewayPortWanSourceNat() *GatewayPortWanSourceNat`

NewGatewayPortWanSourceNat instantiates a new GatewayPortWanSourceNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortWanSourceNatWithDefaults

`func NewGatewayPortWanSourceNatWithDefaults() *GatewayPortWanSourceNat`

NewGatewayPortWanSourceNatWithDefaults instantiates a new GatewayPortWanSourceNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *GatewayPortWanSourceNat) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayPortWanSourceNat) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayPortWanSourceNat) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayPortWanSourceNat) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetNatPool

`func (o *GatewayPortWanSourceNat) GetNatPool() string`

GetNatPool returns the NatPool field if non-nil, zero value otherwise.

### GetNatPoolOk

`func (o *GatewayPortWanSourceNat) GetNatPoolOk() (*string, bool)`

GetNatPoolOk returns a tuple with the NatPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPool

`func (o *GatewayPortWanSourceNat) SetNatPool(v string)`

SetNatPool sets NatPool field to given value.

### HasNatPool

`func (o *GatewayPortWanSourceNat) HasNatPool() bool`

HasNatPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


