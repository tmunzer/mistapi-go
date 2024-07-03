# GatewayTemplateTunnelIkeProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthAlgo** | Pointer to [**TunnelConfigsAuthAlgo**](TunnelConfigsAuthAlgo.md) |  | [optional] 
**DhGroup** | Pointer to [**GatewayTemplateTunnelIkeDhGroup**](GatewayTemplateTunnelIkeDhGroup.md) |  | [optional] [default to GATEWAYTEMPLATETUNNELIKEDHGROUP__14]
**EncAlgo** | Pointer to [**NullableTunnelConfigsEncAlgo**](TunnelConfigsEncAlgo.md) |  | [optional] [default to TUNNELCONFIGSENCALGO_AES256]

## Methods

### NewGatewayTemplateTunnelIkeProposal

`func NewGatewayTemplateTunnelIkeProposal() *GatewayTemplateTunnelIkeProposal`

NewGatewayTemplateTunnelIkeProposal instantiates a new GatewayTemplateTunnelIkeProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateTunnelIkeProposalWithDefaults

`func NewGatewayTemplateTunnelIkeProposalWithDefaults() *GatewayTemplateTunnelIkeProposal`

NewGatewayTemplateTunnelIkeProposalWithDefaults instantiates a new GatewayTemplateTunnelIkeProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) GetAuthAlgo() TunnelConfigsAuthAlgo`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *GatewayTemplateTunnelIkeProposal) GetAuthAlgoOk() (*TunnelConfigsAuthAlgo, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) SetAuthAlgo(v TunnelConfigsAuthAlgo)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *GatewayTemplateTunnelIkeProposal) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) GetDhGroup() GatewayTemplateTunnelIkeDhGroup`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *GatewayTemplateTunnelIkeProposal) GetDhGroupOk() (*GatewayTemplateTunnelIkeDhGroup, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) SetDhGroup(v GatewayTemplateTunnelIkeDhGroup)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *GatewayTemplateTunnelIkeProposal) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) GetEncAlgo() TunnelConfigsEncAlgo`

GetEncAlgo returns the EncAlgo field if non-nil, zero value otherwise.

### GetEncAlgoOk

`func (o *GatewayTemplateTunnelIkeProposal) GetEncAlgoOk() (*TunnelConfigsEncAlgo, bool)`

GetEncAlgoOk returns a tuple with the EncAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) SetEncAlgo(v TunnelConfigsEncAlgo)`

SetEncAlgo sets EncAlgo field to given value.

### HasEncAlgo

`func (o *GatewayTemplateTunnelIkeProposal) HasEncAlgo() bool`

HasEncAlgo returns a boolean if a field has been set.

### SetEncAlgoNil

`func (o *GatewayTemplateTunnelIkeProposal) SetEncAlgoNil(b bool)`

 SetEncAlgoNil sets the value for EncAlgo to be an explicit nil

### UnsetEncAlgo
`func (o *GatewayTemplateTunnelIkeProposal) UnsetEncAlgo()`

UnsetEncAlgo ensures that no value is present for EncAlgo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


