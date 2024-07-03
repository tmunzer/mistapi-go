# ApPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] [default to false]
**DynamicVlan** | Pointer to [**ApPortConfigDynamicVlan**](ApPortConfigDynamicVlan.md) |  | [optional] 
**EnableMacAuth** | Pointer to **bool** |  | [optional] [default to false]
**Forwarding** | Pointer to [**ApPortConfigForwarding**](ApPortConfigForwarding.md) |  | [optional] [default to APPORTCONFIGFORWARDING_ALL]
**MacAuthProtocol** | Pointer to [**ApPortConfigMacAuthProtocol**](ApPortConfigMacAuthProtocol.md) |  | [optional] [default to APPORTCONFIGMACAUTHPROTOCOL_PAP]
**MistNac** | Pointer to [**WlanMistNac**](WlanMistNac.md) |  | [optional] 
**MxTunnelId** | Pointer to **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;mxtunnel&#x60;, vlan_ids comes from mxtunnel | [optional] 
**MxtunnelName** | Pointer to **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;site_mxedge&#x60;, vlan_ids comes from site_mxedge (&#x60;mxtunnels&#x60; under site setting) | [optional] [default to ""]
**PortAuth** | Pointer to [**ApPortConfigPortAuth**](ApPortConfigPortAuth.md) |  | [optional] [default to APPORTCONFIGPORTAUTH_NONE]
**PortVlanId** | Pointer to **int32** | if &#x60;forwrding&#x60;&#x3D;&#x3D;&#x60;limited&#x60; | [optional] 
**RadiusConfig** | Pointer to [**RadiusConfig**](RadiusConfig.md) |  | [optional] 
**Radsec** | Pointer to [**Radsec**](Radsec.md) |  | [optional] 
**VlanId** | Pointer to **int32** | optional to specify the vlan id for a tunnel if forwarding is for &#x60;wxtunnel&#x60;, &#x60;mxtunnel&#x60; or &#x60;site_mxedge&#x60;. * if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel. * if forwarding &#x3D;&#x3D; site_mxedge, vlan_ids comes from site_mxedge (&#x60;mxtunnels&#x60; under site setting) | [optional] 
**VlandIds** | Pointer to **[]int32** | if &#x60;forwrding&#x60;&#x3D;&#x3D;&#x60;limited&#x60; | [optional] 
**WxtunnelId** | Pointer to **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;, the port is bridged to the vlan of the session | [optional] 
**WxtunnelRemoteId** | Pointer to **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;, the port is bridged to the vlan of the session | [optional] [default to ""]

## Methods

### NewApPortConfig

`func NewApPortConfig() *ApPortConfig`

NewApPortConfig instantiates a new ApPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApPortConfigWithDefaults

`func NewApPortConfigWithDefaults() *ApPortConfig`

NewApPortConfigWithDefaults instantiates a new ApPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *ApPortConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApPortConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApPortConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApPortConfig) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDynamicVlan

`func (o *ApPortConfig) GetDynamicVlan() ApPortConfigDynamicVlan`

GetDynamicVlan returns the DynamicVlan field if non-nil, zero value otherwise.

### GetDynamicVlanOk

`func (o *ApPortConfig) GetDynamicVlanOk() (*ApPortConfigDynamicVlan, bool)`

GetDynamicVlanOk returns a tuple with the DynamicVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVlan

`func (o *ApPortConfig) SetDynamicVlan(v ApPortConfigDynamicVlan)`

SetDynamicVlan sets DynamicVlan field to given value.

### HasDynamicVlan

`func (o *ApPortConfig) HasDynamicVlan() bool`

HasDynamicVlan returns a boolean if a field has been set.

### GetEnableMacAuth

`func (o *ApPortConfig) GetEnableMacAuth() bool`

GetEnableMacAuth returns the EnableMacAuth field if non-nil, zero value otherwise.

### GetEnableMacAuthOk

`func (o *ApPortConfig) GetEnableMacAuthOk() (*bool, bool)`

GetEnableMacAuthOk returns a tuple with the EnableMacAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMacAuth

`func (o *ApPortConfig) SetEnableMacAuth(v bool)`

SetEnableMacAuth sets EnableMacAuth field to given value.

### HasEnableMacAuth

`func (o *ApPortConfig) HasEnableMacAuth() bool`

HasEnableMacAuth returns a boolean if a field has been set.

### GetForwarding

`func (o *ApPortConfig) GetForwarding() ApPortConfigForwarding`

GetForwarding returns the Forwarding field if non-nil, zero value otherwise.

### GetForwardingOk

`func (o *ApPortConfig) GetForwardingOk() (*ApPortConfigForwarding, bool)`

GetForwardingOk returns a tuple with the Forwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarding

`func (o *ApPortConfig) SetForwarding(v ApPortConfigForwarding)`

SetForwarding sets Forwarding field to given value.

### HasForwarding

`func (o *ApPortConfig) HasForwarding() bool`

HasForwarding returns a boolean if a field has been set.

### GetMacAuthProtocol

`func (o *ApPortConfig) GetMacAuthProtocol() ApPortConfigMacAuthProtocol`

GetMacAuthProtocol returns the MacAuthProtocol field if non-nil, zero value otherwise.

### GetMacAuthProtocolOk

`func (o *ApPortConfig) GetMacAuthProtocolOk() (*ApPortConfigMacAuthProtocol, bool)`

GetMacAuthProtocolOk returns a tuple with the MacAuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAuthProtocol

`func (o *ApPortConfig) SetMacAuthProtocol(v ApPortConfigMacAuthProtocol)`

SetMacAuthProtocol sets MacAuthProtocol field to given value.

### HasMacAuthProtocol

`func (o *ApPortConfig) HasMacAuthProtocol() bool`

HasMacAuthProtocol returns a boolean if a field has been set.

### GetMistNac

`func (o *ApPortConfig) GetMistNac() WlanMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *ApPortConfig) GetMistNacOk() (*WlanMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *ApPortConfig) SetMistNac(v WlanMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *ApPortConfig) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetMxTunnelId

`func (o *ApPortConfig) GetMxTunnelId() string`

GetMxTunnelId returns the MxTunnelId field if non-nil, zero value otherwise.

### GetMxTunnelIdOk

`func (o *ApPortConfig) GetMxTunnelIdOk() (*string, bool)`

GetMxTunnelIdOk returns a tuple with the MxTunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxTunnelId

`func (o *ApPortConfig) SetMxTunnelId(v string)`

SetMxTunnelId sets MxTunnelId field to given value.

### HasMxTunnelId

`func (o *ApPortConfig) HasMxTunnelId() bool`

HasMxTunnelId returns a boolean if a field has been set.

### GetMxtunnelName

`func (o *ApPortConfig) GetMxtunnelName() string`

GetMxtunnelName returns the MxtunnelName field if non-nil, zero value otherwise.

### GetMxtunnelNameOk

`func (o *ApPortConfig) GetMxtunnelNameOk() (*string, bool)`

GetMxtunnelNameOk returns a tuple with the MxtunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelName

`func (o *ApPortConfig) SetMxtunnelName(v string)`

SetMxtunnelName sets MxtunnelName field to given value.

### HasMxtunnelName

`func (o *ApPortConfig) HasMxtunnelName() bool`

HasMxtunnelName returns a boolean if a field has been set.

### GetPortAuth

`func (o *ApPortConfig) GetPortAuth() ApPortConfigPortAuth`

GetPortAuth returns the PortAuth field if non-nil, zero value otherwise.

### GetPortAuthOk

`func (o *ApPortConfig) GetPortAuthOk() (*ApPortConfigPortAuth, bool)`

GetPortAuthOk returns a tuple with the PortAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortAuth

`func (o *ApPortConfig) SetPortAuth(v ApPortConfigPortAuth)`

SetPortAuth sets PortAuth field to given value.

### HasPortAuth

`func (o *ApPortConfig) HasPortAuth() bool`

HasPortAuth returns a boolean if a field has been set.

### GetPortVlanId

`func (o *ApPortConfig) GetPortVlanId() int32`

GetPortVlanId returns the PortVlanId field if non-nil, zero value otherwise.

### GetPortVlanIdOk

`func (o *ApPortConfig) GetPortVlanIdOk() (*int32, bool)`

GetPortVlanIdOk returns a tuple with the PortVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanId

`func (o *ApPortConfig) SetPortVlanId(v int32)`

SetPortVlanId sets PortVlanId field to given value.

### HasPortVlanId

`func (o *ApPortConfig) HasPortVlanId() bool`

HasPortVlanId returns a boolean if a field has been set.

### GetRadiusConfig

`func (o *ApPortConfig) GetRadiusConfig() RadiusConfig`

GetRadiusConfig returns the RadiusConfig field if non-nil, zero value otherwise.

### GetRadiusConfigOk

`func (o *ApPortConfig) GetRadiusConfigOk() (*RadiusConfig, bool)`

GetRadiusConfigOk returns a tuple with the RadiusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusConfig

`func (o *ApPortConfig) SetRadiusConfig(v RadiusConfig)`

SetRadiusConfig sets RadiusConfig field to given value.

### HasRadiusConfig

`func (o *ApPortConfig) HasRadiusConfig() bool`

HasRadiusConfig returns a boolean if a field has been set.

### GetRadsec

`func (o *ApPortConfig) GetRadsec() Radsec`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *ApPortConfig) GetRadsecOk() (*Radsec, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *ApPortConfig) SetRadsec(v Radsec)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *ApPortConfig) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.

### GetVlanId

`func (o *ApPortConfig) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ApPortConfig) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ApPortConfig) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ApPortConfig) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlandIds

`func (o *ApPortConfig) GetVlandIds() []int32`

GetVlandIds returns the VlandIds field if non-nil, zero value otherwise.

### GetVlandIdsOk

`func (o *ApPortConfig) GetVlandIdsOk() (*[]int32, bool)`

GetVlandIdsOk returns a tuple with the VlandIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlandIds

`func (o *ApPortConfig) SetVlandIds(v []int32)`

SetVlandIds sets VlandIds field to given value.

### HasVlandIds

`func (o *ApPortConfig) HasVlandIds() bool`

HasVlandIds returns a boolean if a field has been set.

### GetWxtunnelId

`func (o *ApPortConfig) GetWxtunnelId() string`

GetWxtunnelId returns the WxtunnelId field if non-nil, zero value otherwise.

### GetWxtunnelIdOk

`func (o *ApPortConfig) GetWxtunnelIdOk() (*string, bool)`

GetWxtunnelIdOk returns a tuple with the WxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtunnelId

`func (o *ApPortConfig) SetWxtunnelId(v string)`

SetWxtunnelId sets WxtunnelId field to given value.

### HasWxtunnelId

`func (o *ApPortConfig) HasWxtunnelId() bool`

HasWxtunnelId returns a boolean if a field has been set.

### GetWxtunnelRemoteId

`func (o *ApPortConfig) GetWxtunnelRemoteId() string`

GetWxtunnelRemoteId returns the WxtunnelRemoteId field if non-nil, zero value otherwise.

### GetWxtunnelRemoteIdOk

`func (o *ApPortConfig) GetWxtunnelRemoteIdOk() (*string, bool)`

GetWxtunnelRemoteIdOk returns a tuple with the WxtunnelRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxtunnelRemoteId

`func (o *ApPortConfig) SetWxtunnelRemoteId(v string)`

SetWxtunnelRemoteId sets WxtunnelRemoteId field to given value.

### HasWxtunnelRemoteId

`func (o *ApPortConfig) HasWxtunnelRemoteId() bool`

HasWxtunnelRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


