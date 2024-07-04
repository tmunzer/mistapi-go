# GatewayMatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalConfigCmds** | Pointer to **[]string** | additional CLI commands to append to the generated Junos config  **Note**: no check is done | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortConfig** | Pointer to [**map[string]JunosPortConfig**](JunosPortConfig.md) |  | [optional] 

## Methods

### NewGatewayMatchingRule

`func NewGatewayMatchingRule() *GatewayMatchingRule`

NewGatewayMatchingRule instantiates a new GatewayMatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMatchingRuleWithDefaults

`func NewGatewayMatchingRuleWithDefaults() *GatewayMatchingRule`

NewGatewayMatchingRuleWithDefaults instantiates a new GatewayMatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalConfigCmds

`func (o *GatewayMatchingRule) GetAdditionalConfigCmds() []string`

GetAdditionalConfigCmds returns the AdditionalConfigCmds field if non-nil, zero value otherwise.

### GetAdditionalConfigCmdsOk

`func (o *GatewayMatchingRule) GetAdditionalConfigCmdsOk() (*[]string, bool)`

GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigCmds

`func (o *GatewayMatchingRule) SetAdditionalConfigCmds(v []string)`

SetAdditionalConfigCmds sets AdditionalConfigCmds field to given value.

### HasAdditionalConfigCmds

`func (o *GatewayMatchingRule) HasAdditionalConfigCmds() bool`

HasAdditionalConfigCmds returns a boolean if a field has been set.

### GetName

`func (o *GatewayMatchingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayMatchingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayMatchingRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayMatchingRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfig

`func (o *GatewayMatchingRule) GetPortConfig() map[string]JunosPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *GatewayMatchingRule) GetPortConfigOk() (*map[string]JunosPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *GatewayMatchingRule) SetPortConfig(v map[string]JunosPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *GatewayMatchingRule) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


