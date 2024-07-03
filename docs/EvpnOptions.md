# EvpnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLoopbackSubnet** | Pointer to **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides | [optional] 
**AutoLoopbackSubnet6** | Pointer to **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides | [optional] 
**AutoRouterIdSubnet** | Pointer to **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] 
**AutoRouterIdSubnet6** | Pointer to **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] [default to "fd31:5700:1::/64"]
**CoreAsBorder** | Pointer to **bool** | optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway when &#x60;routed_at&#x60; !&#x3D; &#x60;core&#x60;, whether to do virtual-gateway at core as well | [optional] [default to false]
**Overlay** | Pointer to [**EvpnOptionsOverlay**](EvpnOptionsOverlay.md) |  | [optional] 
**PerVlanVgaV4Mac** | Pointer to **bool** | by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address&#39;s v4_mac if enabled, 00-00-5e-00-XX-YY will be used (where XX&#x3D;vlan_id/256, YY&#x3D;vlan_id%256) | [optional] [default to false]
**RoutedAt** | Pointer to [**EvpnOptionsRoutedAt**](EvpnOptionsRoutedAt.md) |  | [optional] [default to EVPNOPTIONSROUTEDAT_EDGE]
**Underlay** | Pointer to [**EvpnOptionsUnderlay**](EvpnOptionsUnderlay.md) |  | [optional] 
**VsInstances** | Pointer to [**map[string]EvpnOptionsVsInstance**](EvpnOptionsVsInstance.md) | optional, for EX9200 only to seggregate virtual-switches | [optional] 

## Methods

### NewEvpnOptions

`func NewEvpnOptions() *EvpnOptions`

NewEvpnOptions instantiates a new EvpnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvpnOptionsWithDefaults

`func NewEvpnOptionsWithDefaults() *EvpnOptions`

NewEvpnOptionsWithDefaults instantiates a new EvpnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLoopbackSubnet

`func (o *EvpnOptions) GetAutoLoopbackSubnet() string`

GetAutoLoopbackSubnet returns the AutoLoopbackSubnet field if non-nil, zero value otherwise.

### GetAutoLoopbackSubnetOk

`func (o *EvpnOptions) GetAutoLoopbackSubnetOk() (*string, bool)`

GetAutoLoopbackSubnetOk returns a tuple with the AutoLoopbackSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoopbackSubnet

`func (o *EvpnOptions) SetAutoLoopbackSubnet(v string)`

SetAutoLoopbackSubnet sets AutoLoopbackSubnet field to given value.

### HasAutoLoopbackSubnet

`func (o *EvpnOptions) HasAutoLoopbackSubnet() bool`

HasAutoLoopbackSubnet returns a boolean if a field has been set.

### GetAutoLoopbackSubnet6

`func (o *EvpnOptions) GetAutoLoopbackSubnet6() string`

GetAutoLoopbackSubnet6 returns the AutoLoopbackSubnet6 field if non-nil, zero value otherwise.

### GetAutoLoopbackSubnet6Ok

`func (o *EvpnOptions) GetAutoLoopbackSubnet6Ok() (*string, bool)`

GetAutoLoopbackSubnet6Ok returns a tuple with the AutoLoopbackSubnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoopbackSubnet6

`func (o *EvpnOptions) SetAutoLoopbackSubnet6(v string)`

SetAutoLoopbackSubnet6 sets AutoLoopbackSubnet6 field to given value.

### HasAutoLoopbackSubnet6

`func (o *EvpnOptions) HasAutoLoopbackSubnet6() bool`

HasAutoLoopbackSubnet6 returns a boolean if a field has been set.

### GetAutoRouterIdSubnet

`func (o *EvpnOptions) GetAutoRouterIdSubnet() string`

GetAutoRouterIdSubnet returns the AutoRouterIdSubnet field if non-nil, zero value otherwise.

### GetAutoRouterIdSubnetOk

`func (o *EvpnOptions) GetAutoRouterIdSubnetOk() (*string, bool)`

GetAutoRouterIdSubnetOk returns a tuple with the AutoRouterIdSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterIdSubnet

`func (o *EvpnOptions) SetAutoRouterIdSubnet(v string)`

SetAutoRouterIdSubnet sets AutoRouterIdSubnet field to given value.

### HasAutoRouterIdSubnet

`func (o *EvpnOptions) HasAutoRouterIdSubnet() bool`

HasAutoRouterIdSubnet returns a boolean if a field has been set.

### GetAutoRouterIdSubnet6

`func (o *EvpnOptions) GetAutoRouterIdSubnet6() string`

GetAutoRouterIdSubnet6 returns the AutoRouterIdSubnet6 field if non-nil, zero value otherwise.

### GetAutoRouterIdSubnet6Ok

`func (o *EvpnOptions) GetAutoRouterIdSubnet6Ok() (*string, bool)`

GetAutoRouterIdSubnet6Ok returns a tuple with the AutoRouterIdSubnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouterIdSubnet6

`func (o *EvpnOptions) SetAutoRouterIdSubnet6(v string)`

SetAutoRouterIdSubnet6 sets AutoRouterIdSubnet6 field to given value.

### HasAutoRouterIdSubnet6

`func (o *EvpnOptions) HasAutoRouterIdSubnet6() bool`

HasAutoRouterIdSubnet6 returns a boolean if a field has been set.

### GetCoreAsBorder

`func (o *EvpnOptions) GetCoreAsBorder() bool`

GetCoreAsBorder returns the CoreAsBorder field if non-nil, zero value otherwise.

### GetCoreAsBorderOk

`func (o *EvpnOptions) GetCoreAsBorderOk() (*bool, bool)`

GetCoreAsBorderOk returns a tuple with the CoreAsBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAsBorder

`func (o *EvpnOptions) SetCoreAsBorder(v bool)`

SetCoreAsBorder sets CoreAsBorder field to given value.

### HasCoreAsBorder

`func (o *EvpnOptions) HasCoreAsBorder() bool`

HasCoreAsBorder returns a boolean if a field has been set.

### GetOverlay

`func (o *EvpnOptions) GetOverlay() EvpnOptionsOverlay`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *EvpnOptions) GetOverlayOk() (*EvpnOptionsOverlay, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *EvpnOptions) SetOverlay(v EvpnOptionsOverlay)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *EvpnOptions) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetPerVlanVgaV4Mac

`func (o *EvpnOptions) GetPerVlanVgaV4Mac() bool`

GetPerVlanVgaV4Mac returns the PerVlanVgaV4Mac field if non-nil, zero value otherwise.

### GetPerVlanVgaV4MacOk

`func (o *EvpnOptions) GetPerVlanVgaV4MacOk() (*bool, bool)`

GetPerVlanVgaV4MacOk returns a tuple with the PerVlanVgaV4Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVlanVgaV4Mac

`func (o *EvpnOptions) SetPerVlanVgaV4Mac(v bool)`

SetPerVlanVgaV4Mac sets PerVlanVgaV4Mac field to given value.

### HasPerVlanVgaV4Mac

`func (o *EvpnOptions) HasPerVlanVgaV4Mac() bool`

HasPerVlanVgaV4Mac returns a boolean if a field has been set.

### GetRoutedAt

`func (o *EvpnOptions) GetRoutedAt() EvpnOptionsRoutedAt`

GetRoutedAt returns the RoutedAt field if non-nil, zero value otherwise.

### GetRoutedAtOk

`func (o *EvpnOptions) GetRoutedAtOk() (*EvpnOptionsRoutedAt, bool)`

GetRoutedAtOk returns a tuple with the RoutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutedAt

`func (o *EvpnOptions) SetRoutedAt(v EvpnOptionsRoutedAt)`

SetRoutedAt sets RoutedAt field to given value.

### HasRoutedAt

`func (o *EvpnOptions) HasRoutedAt() bool`

HasRoutedAt returns a boolean if a field has been set.

### GetUnderlay

`func (o *EvpnOptions) GetUnderlay() EvpnOptionsUnderlay`

GetUnderlay returns the Underlay field if non-nil, zero value otherwise.

### GetUnderlayOk

`func (o *EvpnOptions) GetUnderlayOk() (*EvpnOptionsUnderlay, bool)`

GetUnderlayOk returns a tuple with the Underlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlay

`func (o *EvpnOptions) SetUnderlay(v EvpnOptionsUnderlay)`

SetUnderlay sets Underlay field to given value.

### HasUnderlay

`func (o *EvpnOptions) HasUnderlay() bool`

HasUnderlay returns a boolean if a field has been set.

### GetVsInstances

`func (o *EvpnOptions) GetVsInstances() map[string]EvpnOptionsVsInstance`

GetVsInstances returns the VsInstances field if non-nil, zero value otherwise.

### GetVsInstancesOk

`func (o *EvpnOptions) GetVsInstancesOk() (*map[string]EvpnOptionsVsInstance, bool)`

GetVsInstancesOk returns a tuple with the VsInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsInstances

`func (o *EvpnOptions) SetVsInstances(v map[string]EvpnOptionsVsInstance)`

SetVsInstances sets VsInstances field to given value.

### HasVsInstances

`func (o *EvpnOptions) HasVsInstances() bool`

HasVsInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


