# ApTemplateMatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchModel** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortConfig** | Pointer to [**map[string]ApPortConfig**](ApPortConfig.md) | Property key is the interface(s) name (e.g. \&quot;eth1,eth2\&quot;) | [optional] 

## Methods

### NewApTemplateMatchingRule

`func NewApTemplateMatchingRule() *ApTemplateMatchingRule`

NewApTemplateMatchingRule instantiates a new ApTemplateMatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApTemplateMatchingRuleWithDefaults

`func NewApTemplateMatchingRuleWithDefaults() *ApTemplateMatchingRule`

NewApTemplateMatchingRuleWithDefaults instantiates a new ApTemplateMatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchModel

`func (o *ApTemplateMatchingRule) GetMatchModel() string`

GetMatchModel returns the MatchModel field if non-nil, zero value otherwise.

### GetMatchModelOk

`func (o *ApTemplateMatchingRule) GetMatchModelOk() (*string, bool)`

GetMatchModelOk returns a tuple with the MatchModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchModel

`func (o *ApTemplateMatchingRule) SetMatchModel(v string)`

SetMatchModel sets MatchModel field to given value.

### HasMatchModel

`func (o *ApTemplateMatchingRule) HasMatchModel() bool`

HasMatchModel returns a boolean if a field has been set.

### GetName

`func (o *ApTemplateMatchingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApTemplateMatchingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApTemplateMatchingRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApTemplateMatchingRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfig

`func (o *ApTemplateMatchingRule) GetPortConfig() map[string]ApPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *ApTemplateMatchingRule) GetPortConfigOk() (*map[string]ApPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *ApTemplateMatchingRule) SetPortConfig(v map[string]ApPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *ApTemplateMatchingRule) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


