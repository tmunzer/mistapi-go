# WxlanTunnelDmvpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether DMVPN is enabled | [optional] [default to false]
**HoldingTime** | Pointer to **int32** | optional; the holding time for NHRP ‘registration requests’ and ‘resolution replies’ sent from the Mist AP (in seconds); default 600 | [optional] 
**HostRoutes** | Pointer to **[]string** | optional; list of IPv4 DMVPN peer host ip-addresses to which traffic is forwarded | [optional] 

## Methods

### NewWxlanTunnelDmvpn

`func NewWxlanTunnelDmvpn() *WxlanTunnelDmvpn`

NewWxlanTunnelDmvpn instantiates a new WxlanTunnelDmvpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTunnelDmvpnWithDefaults

`func NewWxlanTunnelDmvpnWithDefaults() *WxlanTunnelDmvpn`

NewWxlanTunnelDmvpnWithDefaults instantiates a new WxlanTunnelDmvpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WxlanTunnelDmvpn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WxlanTunnelDmvpn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WxlanTunnelDmvpn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WxlanTunnelDmvpn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHoldingTime

`func (o *WxlanTunnelDmvpn) GetHoldingTime() int32`

GetHoldingTime returns the HoldingTime field if non-nil, zero value otherwise.

### GetHoldingTimeOk

`func (o *WxlanTunnelDmvpn) GetHoldingTimeOk() (*int32, bool)`

GetHoldingTimeOk returns a tuple with the HoldingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingTime

`func (o *WxlanTunnelDmvpn) SetHoldingTime(v int32)`

SetHoldingTime sets HoldingTime field to given value.

### HasHoldingTime

`func (o *WxlanTunnelDmvpn) HasHoldingTime() bool`

HasHoldingTime returns a boolean if a field has been set.

### GetHostRoutes

`func (o *WxlanTunnelDmvpn) GetHostRoutes() []string`

GetHostRoutes returns the HostRoutes field if non-nil, zero value otherwise.

### GetHostRoutesOk

`func (o *WxlanTunnelDmvpn) GetHostRoutesOk() (*[]string, bool)`

GetHostRoutesOk returns a tuple with the HostRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoutes

`func (o *WxlanTunnelDmvpn) SetHostRoutes(v []string)`

SetHostRoutes sets HostRoutes field to given value.

### HasHostRoutes

`func (o *WxlanTunnelDmvpn) HasHostRoutes() bool`

HasHostRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


