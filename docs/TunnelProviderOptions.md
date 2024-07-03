# TunnelProviderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jse** | Pointer to [**TunnelProviderOptionsJse**](TunnelProviderOptionsJse.md) |  | [optional] 
**Zscaler** | Pointer to [**TunnelProviderOptionsZscaler**](TunnelProviderOptionsZscaler.md) |  | [optional] 

## Methods

### NewTunnelProviderOptions

`func NewTunnelProviderOptions() *TunnelProviderOptions`

NewTunnelProviderOptions instantiates a new TunnelProviderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelProviderOptionsWithDefaults

`func NewTunnelProviderOptionsWithDefaults() *TunnelProviderOptions`

NewTunnelProviderOptionsWithDefaults instantiates a new TunnelProviderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJse

`func (o *TunnelProviderOptions) GetJse() TunnelProviderOptionsJse`

GetJse returns the Jse field if non-nil, zero value otherwise.

### GetJseOk

`func (o *TunnelProviderOptions) GetJseOk() (*TunnelProviderOptionsJse, bool)`

GetJseOk returns a tuple with the Jse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJse

`func (o *TunnelProviderOptions) SetJse(v TunnelProviderOptionsJse)`

SetJse sets Jse field to given value.

### HasJse

`func (o *TunnelProviderOptions) HasJse() bool`

HasJse returns a boolean if a field has been set.

### GetZscaler

`func (o *TunnelProviderOptions) GetZscaler() TunnelProviderOptionsZscaler`

GetZscaler returns the Zscaler field if non-nil, zero value otherwise.

### GetZscalerOk

`func (o *TunnelProviderOptions) GetZscalerOk() (*TunnelProviderOptionsZscaler, bool)`

GetZscalerOk returns a tuple with the Zscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZscaler

`func (o *TunnelProviderOptions) SetZscaler(v TunnelProviderOptionsZscaler)`

SetZscaler sets Zscaler field to given value.

### HasZscaler

`func (o *TunnelProviderOptions) HasZscaler() bool`

HasZscaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


