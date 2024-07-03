# GatewayMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]GatewayMatchingRule**](GatewayMatchingRule.md) |  | [optional] 

## Methods

### NewGatewayMatching

`func NewGatewayMatching() *GatewayMatching`

NewGatewayMatching instantiates a new GatewayMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMatchingWithDefaults

`func NewGatewayMatchingWithDefaults() *GatewayMatching`

NewGatewayMatchingWithDefaults instantiates a new GatewayMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayMatching) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayMatching) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayMatching) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayMatching) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRules

`func (o *GatewayMatching) GetRules() []GatewayMatchingRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GatewayMatching) GetRulesOk() (*[]GatewayMatchingRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GatewayMatching) SetRules(v []GatewayMatchingRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GatewayMatching) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


