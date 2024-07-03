# SwitchMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]SwitchMatchingRule**](SwitchMatchingRule.md) |  | [optional] 

## Methods

### NewSwitchMatching

`func NewSwitchMatching() *SwitchMatching`

NewSwitchMatching instantiates a new SwitchMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchMatchingWithDefaults

`func NewSwitchMatchingWithDefaults() *SwitchMatching`

NewSwitchMatchingWithDefaults instantiates a new SwitchMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SwitchMatching) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SwitchMatching) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SwitchMatching) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SwitchMatching) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRules

`func (o *SwitchMatching) GetRules() []SwitchMatchingRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SwitchMatching) GetRulesOk() (*[]SwitchMatchingRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SwitchMatching) SetRules(v []SwitchMatchingRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SwitchMatching) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


