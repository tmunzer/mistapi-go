# WxlanTunnelIpsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether ipsec is enabled, requires DMVPN be enabled | [optional] [default to false]
**Psk** | **string** | ipsec pre-shared key | 

## Methods

### NewWxlanTunnelIpsec

`func NewWxlanTunnelIpsec(psk string, ) *WxlanTunnelIpsec`

NewWxlanTunnelIpsec instantiates a new WxlanTunnelIpsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTunnelIpsecWithDefaults

`func NewWxlanTunnelIpsecWithDefaults() *WxlanTunnelIpsec`

NewWxlanTunnelIpsecWithDefaults instantiates a new WxlanTunnelIpsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WxlanTunnelIpsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WxlanTunnelIpsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WxlanTunnelIpsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WxlanTunnelIpsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPsk

`func (o *WxlanTunnelIpsec) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *WxlanTunnelIpsec) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *WxlanTunnelIpsec) SetPsk(v string)`

SetPsk sets Psk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


