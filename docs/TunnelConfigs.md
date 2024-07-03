# TunnelConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoProvision** | Pointer to [**TunnelConfigsAutoProvision**](TunnelConfigsAutoProvision.md) |  | [optional] 
**IkeLifetime** | Pointer to **int32** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IkeMode** | Pointer to [**GatewayTemplateTunnelIkeMode**](GatewayTemplateTunnelIkeMode.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELIKEMODE_MAIN]
**IkeProposals** | Pointer to [**[]GatewayTemplateTunnelIkeProposal**](GatewayTemplateTunnelIkeProposal.md) | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IpsecLifetime** | Pointer to **int32** | if &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**IpsecProposals** | Pointer to [**[]GatewayTemplateTunnelIpsecProposal**](GatewayTemplateTunnelIpsecProposal.md) | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**LocalId** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**Mode** | Pointer to [**GatewayTemplateTunnelMode**](GatewayTemplateTunnelMode.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELMODE_ACTIVE_STANDBY]
**Primary** | Pointer to [**GatewayTemplateTunnelNode**](GatewayTemplateTunnelNode.md) |  | [optional] 
**Probe** | Pointer to [**GatewayTemplateTunnelProbe**](GatewayTemplateTunnelProbe.md) |  | [optional] 
**Protocol** | Pointer to [**GatewayTemplateTunnelProtocol**](GatewayTemplateTunnelProtocol.md) |  | [optional] 
**Provider** | Pointer to [**TunnelProviderOptionsName**](TunnelProviderOptionsName.md) |  | [optional] 
**Psk** | Pointer to **string** | Only if: * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;zscaler-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D;&#x60;jse-ipsec&#x60; * &#x60;provider&#x60;&#x3D;&#x3D; &#x60;custom-ipsec&#x60; | [optional] 
**Secondary** | Pointer to [**GatewayTemplateTunnelNode**](GatewayTemplateTunnelNode.md) |  | [optional] 
**Version** | Pointer to [**GatewayTemplateTunnelVersion**](GatewayTemplateTunnelVersion.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELVERSION__2]

## Methods

### NewTunnelConfigs

`func NewTunnelConfigs() *TunnelConfigs`

NewTunnelConfigs instantiates a new TunnelConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigsWithDefaults

`func NewTunnelConfigsWithDefaults() *TunnelConfigs`

NewTunnelConfigsWithDefaults instantiates a new TunnelConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoProvision

`func (o *TunnelConfigs) GetAutoProvision() TunnelConfigsAutoProvision`

GetAutoProvision returns the AutoProvision field if non-nil, zero value otherwise.

### GetAutoProvisionOk

`func (o *TunnelConfigs) GetAutoProvisionOk() (*TunnelConfigsAutoProvision, bool)`

GetAutoProvisionOk returns a tuple with the AutoProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoProvision

`func (o *TunnelConfigs) SetAutoProvision(v TunnelConfigsAutoProvision)`

SetAutoProvision sets AutoProvision field to given value.

### HasAutoProvision

`func (o *TunnelConfigs) HasAutoProvision() bool`

HasAutoProvision returns a boolean if a field has been set.

### GetIkeLifetime

`func (o *TunnelConfigs) GetIkeLifetime() int32`

GetIkeLifetime returns the IkeLifetime field if non-nil, zero value otherwise.

### GetIkeLifetimeOk

`func (o *TunnelConfigs) GetIkeLifetimeOk() (*int32, bool)`

GetIkeLifetimeOk returns a tuple with the IkeLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetime

`func (o *TunnelConfigs) SetIkeLifetime(v int32)`

SetIkeLifetime sets IkeLifetime field to given value.

### HasIkeLifetime

`func (o *TunnelConfigs) HasIkeLifetime() bool`

HasIkeLifetime returns a boolean if a field has been set.

### GetIkeMode

`func (o *TunnelConfigs) GetIkeMode() GatewayTemplateTunnelIkeMode`

GetIkeMode returns the IkeMode field if non-nil, zero value otherwise.

### GetIkeModeOk

`func (o *TunnelConfigs) GetIkeModeOk() (*GatewayTemplateTunnelIkeMode, bool)`

GetIkeModeOk returns a tuple with the IkeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeMode

`func (o *TunnelConfigs) SetIkeMode(v GatewayTemplateTunnelIkeMode)`

SetIkeMode sets IkeMode field to given value.

### HasIkeMode

`func (o *TunnelConfigs) HasIkeMode() bool`

HasIkeMode returns a boolean if a field has been set.

### GetIkeProposals

`func (o *TunnelConfigs) GetIkeProposals() []GatewayTemplateTunnelIkeProposal`

GetIkeProposals returns the IkeProposals field if non-nil, zero value otherwise.

### GetIkeProposalsOk

`func (o *TunnelConfigs) GetIkeProposalsOk() (*[]GatewayTemplateTunnelIkeProposal, bool)`

GetIkeProposalsOk returns a tuple with the IkeProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeProposals

`func (o *TunnelConfigs) SetIkeProposals(v []GatewayTemplateTunnelIkeProposal)`

SetIkeProposals sets IkeProposals field to given value.

### HasIkeProposals

`func (o *TunnelConfigs) HasIkeProposals() bool`

HasIkeProposals returns a boolean if a field has been set.

### GetIpsecLifetime

`func (o *TunnelConfigs) GetIpsecLifetime() int32`

GetIpsecLifetime returns the IpsecLifetime field if non-nil, zero value otherwise.

### GetIpsecLifetimeOk

`func (o *TunnelConfigs) GetIpsecLifetimeOk() (*int32, bool)`

GetIpsecLifetimeOk returns a tuple with the IpsecLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLifetime

`func (o *TunnelConfigs) SetIpsecLifetime(v int32)`

SetIpsecLifetime sets IpsecLifetime field to given value.

### HasIpsecLifetime

`func (o *TunnelConfigs) HasIpsecLifetime() bool`

HasIpsecLifetime returns a boolean if a field has been set.

### GetIpsecProposals

`func (o *TunnelConfigs) GetIpsecProposals() []GatewayTemplateTunnelIpsecProposal`

GetIpsecProposals returns the IpsecProposals field if non-nil, zero value otherwise.

### GetIpsecProposalsOk

`func (o *TunnelConfigs) GetIpsecProposalsOk() (*[]GatewayTemplateTunnelIpsecProposal, bool)`

GetIpsecProposalsOk returns a tuple with the IpsecProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProposals

`func (o *TunnelConfigs) SetIpsecProposals(v []GatewayTemplateTunnelIpsecProposal)`

SetIpsecProposals sets IpsecProposals field to given value.

### HasIpsecProposals

`func (o *TunnelConfigs) HasIpsecProposals() bool`

HasIpsecProposals returns a boolean if a field has been set.

### GetLocalId

`func (o *TunnelConfigs) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *TunnelConfigs) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *TunnelConfigs) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *TunnelConfigs) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetMode

`func (o *TunnelConfigs) GetMode() GatewayTemplateTunnelMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TunnelConfigs) GetModeOk() (*GatewayTemplateTunnelMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TunnelConfigs) SetMode(v GatewayTemplateTunnelMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TunnelConfigs) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPrimary

`func (o *TunnelConfigs) GetPrimary() GatewayTemplateTunnelNode`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *TunnelConfigs) GetPrimaryOk() (*GatewayTemplateTunnelNode, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *TunnelConfigs) SetPrimary(v GatewayTemplateTunnelNode)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *TunnelConfigs) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetProbe

`func (o *TunnelConfigs) GetProbe() GatewayTemplateTunnelProbe`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *TunnelConfigs) GetProbeOk() (*GatewayTemplateTunnelProbe, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *TunnelConfigs) SetProbe(v GatewayTemplateTunnelProbe)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *TunnelConfigs) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### GetProtocol

`func (o *TunnelConfigs) GetProtocol() GatewayTemplateTunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TunnelConfigs) GetProtocolOk() (*GatewayTemplateTunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TunnelConfigs) SetProtocol(v GatewayTemplateTunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *TunnelConfigs) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvider

`func (o *TunnelConfigs) GetProvider() TunnelProviderOptionsName`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TunnelConfigs) GetProviderOk() (*TunnelProviderOptionsName, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TunnelConfigs) SetProvider(v TunnelProviderOptionsName)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TunnelConfigs) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPsk

`func (o *TunnelConfigs) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *TunnelConfigs) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *TunnelConfigs) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *TunnelConfigs) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetSecondary

`func (o *TunnelConfigs) GetSecondary() GatewayTemplateTunnelNode`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *TunnelConfigs) GetSecondaryOk() (*GatewayTemplateTunnelNode, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *TunnelConfigs) SetSecondary(v GatewayTemplateTunnelNode)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *TunnelConfigs) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetVersion

`func (o *TunnelConfigs) GetVersion() GatewayTemplateTunnelVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TunnelConfigs) GetVersionOk() (*GatewayTemplateTunnelVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TunnelConfigs) SetVersion(v GatewayTemplateTunnelVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TunnelConfigs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


