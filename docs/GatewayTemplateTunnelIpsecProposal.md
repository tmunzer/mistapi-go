# GatewayTemplateTunnelIpsecProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthAlgo** | Pointer to [**TunnelConfigsAuthAlgo**](TunnelConfigsAuthAlgo.md) |  | [optional] 
**DhGroup** | Pointer to [**TunnelConfigsDhGroup**](TunnelConfigsDhGroup.md) |  | [optional] [default to TUNNELCONFIGSDHGROUP__14]
**EncAlgo** | Pointer to [**NullableTunnelConfigsEncAlgo**](TunnelConfigsEncAlgo.md) |  | [optional] [default to TUNNELCONFIGSENCALGO_AES256]

## Methods

### NewGatewayTemplateTunnelIpsecProposal

`func NewGatewayTemplateTunnelIpsecProposal() *GatewayTemplateTunnelIpsecProposal`

NewGatewayTemplateTunnelIpsecProposal instantiates a new GatewayTemplateTunnelIpsecProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelIpsecProposalWithDefaults

`func NewGatewayTemplateTunnelIpsecProposalWithDefaults() *GatewayTemplateTunnelIpsecProposal`

NewGatewayTemplateTunnelIpsecProposalWithDefaults instantiates a new GatewayTemplateTunnelIpsecProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgo() TunnelConfigsAuthAlgo`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgoOk() (*TunnelConfigsAuthAlgo, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) SetAuthAlgo(v TunnelConfigsAuthAlgo)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroup() TunnelConfigsDhGroup`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroupOk() (*TunnelConfigsDhGroup, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) SetDhGroup(v TunnelConfigsDhGroup)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *GatewayTemplateTunnelIpsecProposal) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgo() TunnelConfigsEncAlgo`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgoOk() (*TunnelConfigsEncAlgo, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgo(v TunnelConfigsEncAlgo)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *GatewayTemplateTunnelIpsecProposal) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### SetEncAlgoNil

`func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgoNil(b bool)`

 SetEncAlgoNil sets the value for EncAlgo to be an explicit nil

### UnsetEncAlgo
`func (o *GatewayTemplateTunnelIpsecProposal) UnsetEncAlgo()`

UnsetEncAlgo ensures that no value is present for EncAlgo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


