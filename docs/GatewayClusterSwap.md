# GatewayClusterSwap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | when &#x60;op&#x60; &#x3D;&#x3D;&#x60;replacement_nodeX&#x60;, new node1 &#39;s MAC, the device has to be standalone and assigned to the same site | [optional] 
**Op** | [**GatewayClusterSwapOp**](GatewayClusterSwapOp.md) |  | [default to GATEWAYCLUSTERSWAPOP_SWAP]

## Methods

### NewGatewayClusterSwap

`func NewGatewayClusterSwap(op GatewayClusterSwapOp, ) *GatewayClusterSwap`

NewGatewayClusterSwap instantiates a new GatewayClusterSwap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayClusterSwapWithDefaults

`func NewGatewayClusterSwapWithDefaults() *GatewayClusterSwap`

NewGatewayClusterSwapWithDefaults instantiates a new GatewayClusterSwap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GatewayClusterSwap) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GatewayClusterSwap) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GatewayClusterSwap) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GatewayClusterSwap) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetOp

`func (o *GatewayClusterSwap) GetOp() GatewayClusterSwapOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GatewayClusterSwap) GetOpOk() (*GatewayClusterSwapOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GatewayClusterSwap) SetOp(v GatewayClusterSwapOp)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


